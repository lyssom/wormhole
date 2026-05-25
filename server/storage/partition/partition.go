package partition

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"wormhole/server/storage/segment"
)

// Partition represents a time-based partition containing mutable and immutable segments.
type Partition struct {
	key        string
	bucketMs   int64
	mutable    *segment.MutableSegment
	immutables []*segment.ImmutableSegment
	mu         sync.RWMutex

	// Total rows ever inserted (cumulative, even after seal/reset)
	totalRows uint64
}

// NewPartition creates a new Partition.
func NewPartition(key string, bucketMs int64) *Partition {
	return &Partition{
		key:        key,
		bucketMs:   bucketMs,
		mutable:    segment.NewMutableSegment(key),
		immutables: make([]*segment.ImmutableSegment, 0),
	}
}

// Key returns the partition key.
func (p *Partition) Key() string {
	return p.key
}

// BucketDurationMs returns the bucket duration in milliseconds.
func (p *Partition) BucketDurationMs() int64 {
	return p.bucketMs
}

// Insert converts rows to column-oriented format and inserts into the mutable segment.
func (p *Partition) Insert(rows []map[string]interface{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Convert row-oriented to column-oriented format
	batch := make(map[string]interface{})
	for _, row := range rows {
		for colName, val := range row {
			if batch[colName] == nil {
				batch[colName] = []interface{}{val}
			} else {
				batch[colName] = append(batch[colName].([]interface{}), val)
			}
		}
	}

	// Get count before insert to track total
	beforeCount := p.mutable.RowCount()

	if err := p.mutable.Append(batch); err != nil {
		return err
	}

	// Update total rows (persists across seal/reset)
	afterCount := p.mutable.RowCount()
	p.totalRows += (afterCount - beforeCount)

	return nil
}

// ReadMutableColumns reads column data from the mutable segment.
func (p *Partition) ReadMutableColumns(names []string) (map[string]interface{}, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.mutable.ReadColumns(names)
}

// ImmutableSegments returns the list of immutable segments.
func (p *Partition) ImmutableSegments() []*segment.ImmutableSegment {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.immutables
}

// AddImmutable adds an immutable segment to the partition.
func (p *Partition) AddImmutable(imm *segment.ImmutableSegment) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.immutables = append(p.immutables, imm)
}

// RowCount returns the total number of rows (mutable + immutable).
func (p *Partition) RowCount() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	count := p.mutable.RowCount()
	for _, imm := range p.immutables {
		count += imm.RowCount()
	}
	return count
}

// MutableRowCount returns the number of rows in the mutable segment.
func (p *Partition) MutableRowCount() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.mutable.RowCount()
}

// TotalRowCount returns the cumulative number of rows ever inserted (persists across seal).
func (p *Partition) TotalRowCount() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.totalRows
}

// NextInsertIndex returns the index at which the next row will be inserted.
func (p *Partition) NextInsertIndex() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.totalRows
}

// ShouldSeal returns true if the mutable segment should be sealed.
func (p *Partition) ShouldSeal() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.mutable.ShouldSeal()
}

// Seal seals the mutable segment and adds it to immutables.
func (p *Partition) Seal(dir string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.mutable.ShouldSeal() {
		return false, nil
	}

	if p.mutable.RowCount() == 0 {
		return false, nil
	}

	// Create file path
	segNum := len(p.immutables)
	filePath := filepath.Join(dir, p.key, "immutable", fmt.Sprintf("%d.msi", segNum))

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return false, fmt.Errorf("mkdir failed: %w", err)
	}

	// Seal to file
	_, err := p.mutable.SealToFile(filePath)
	if err != nil {
		return false, fmt.Errorf("SealToFile failed: %w", err)
	}

	// Open as immutable segment
	imm, err := segment.OpenImmutableSegment(filePath)
	if err != nil {
		return false, fmt.Errorf("OpenImmutableSegment failed: %w", err)
	}

	p.immutables = append(p.immutables, imm)
	return true, nil
}

// LoadImmutableSegments loads existing MSI files from the partition directory.
func (p *Partition) LoadImmutableSegments(dir string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	partDir := filepath.Join(dir, p.key, "immutable")
	entries, err := os.ReadDir(partDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) != ".msi" {
			continue
		}
		path := filepath.Join(partDir, entry.Name())
		imm, err := segment.OpenImmutableSegment(path)
		if err != nil {
			continue // skip corrupted files
		}
		p.immutables = append(p.immutables, imm)
	}
	return nil
}

// RemoveImmutable removes an immutable segment by path.
func (p *Partition) RemoveImmutable(path string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	var remaining []*segment.ImmutableSegment
	for _, imm := range p.immutables {
		if imm.Path() == path {
			continue
		}
		remaining = append(remaining, imm)
	}
	p.immutables = remaining
	return nil
}

// ReplaceImmutables replaces old segments with a new one after merge.
func (p *Partition) ReplaceImmutables(oldPaths []string, newPath string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Build set of old paths
	oldSet := make(map[string]bool)
	for _, path := range oldPaths {
		oldSet[path] = true
	}

	// Filter out old segments
	var remaining []*segment.ImmutableSegment
	for _, imm := range p.immutables {
		if oldSet[imm.Path()] {
			continue
		}
		remaining = append(remaining, imm)
	}

	// Add new segment
	newImm, err := segment.OpenImmutableSegment(newPath)
	if err != nil {
		return fmt.Errorf("OpenImmutableSegment failed: %w", err)
	}
	remaining = append(remaining, newImm)

	p.immutables = remaining
	return nil
}

// PartitionManager manages time-based partitions.
type PartitionManager struct {
	bucketMs   int64
	partitions map[string]*Partition
	mu         sync.RWMutex
}

// NewPartitionManager creates a new PartitionManager.
func NewPartitionManager(bucketMs int64) *PartitionManager {
	return &PartitionManager{
		bucketMs:   bucketMs,
		partitions: make(map[string]*Partition),
	}
}

// BucketKey returns the partition key for a given timestamp.
func BucketKey(tsMs uint64, bucketMs int64) string {
	bucketID := int64(tsMs) / bucketMs
	return fmt.Sprintf("bucket_%d", bucketID)
}

// GetOrCreatePartition returns the partition for a given timestamp.
func (pm *PartitionManager) GetOrCreatePartition(tsMs uint64) *Partition {
	key := BucketKey(tsMs, pm.bucketMs)

	pm.mu.RLock()
	if p, ok := pm.partitions[key]; ok {
		pm.mu.RUnlock()
		return p
	}
	pm.mu.RUnlock()

	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Double-check after acquiring write lock
	if p, ok := pm.partitions[key]; ok {
		return p
	}

	p := NewPartition(key, pm.bucketMs)
	pm.partitions[key] = p
	return p
}

// Partitions returns all partitions.
func (pm *PartitionManager) Partitions() []*Partition {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	parts := make([]*Partition, 0, len(pm.partitions))
	for _, p := range pm.partitions {
		parts = append(parts, p)
	}
	return parts
}

// BucketDurationMs returns the bucket duration.
func (pm *PartitionManager) BucketDurationMs() int64 {
	return pm.bucketMs
}
