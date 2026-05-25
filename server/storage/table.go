package storage

import (
	"fmt"
	"os"
	"sync"

	"wormhole/server/storage/msi"
	"wormhole/server/storage/partition"
	"wormhole/server/storage/segment"
)

// TableMetadata contains table-level metadata (columns and row count).
type TableMetadata = msi.TableMetadata

// RowLocation tracks where a logical row lives in the storage.
type RowLocation struct {
	PartitionKey string
	IsMutable    bool
	SegmentIndex int   // for immutables, the segment index
	LocalIndex   int   // index within the segment's column arrays
}

// Table represents an open table with its mutable segment.
type Table struct {
	name     string
	db       string
	metadata *TableMetadata
	pm       *partition.PartitionManager
	opts     *TableOptions
	mu       sync.RWMutex

	// Row tracking for UpdateRow/DeleteRow
	nextRowID    int                   // global counter for row IDs
	rowIDToLoc   map[int]RowLocation    // rowID -> location in partitions
	pendingUpd   map[int]map[string]interface{} // rowID -> column updates
	pendingDel   map[int]bool          // rowID -> deleted
}

// newTable creates a new Table.
func newTable(db, name string, meta *TableMetadata, opts *TableOptions, bucketMs int64) *Table {
	if opts == nil {
		opts = DefaultTableOptions()
	}
	if bucketMs == 0 {
		bucketMs = DefaultOptions().BucketDurationMs
	}
	return &Table{
		name:       name,
		db:         db,
		metadata:   meta,
		pm:         partition.NewPartitionManager(bucketMs),
		opts:       opts,
		nextRowID:   0,
		rowIDToLoc:  make(map[int]RowLocation),
		pendingUpd:  make(map[int]map[string]interface{}),
		pendingDel: make(map[int]bool),
	}
}

// Insert inserts rows into the table's partition based on timestamp.
func (t *Table) Insert(rows []map[string]interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(rows) == 0 {
		return nil
	}

	// Extract timestamp from first row to determine partition
	var ts uint64 = 0
	if tsVal, ok := rows[0]["ts"].(int64); ok {
		ts = uint64(tsVal)
	}

	part := t.pm.GetOrCreatePartition(ts)

	// Get starting global row ID before insert
	startRowID := t.nextRowID

	// Append to partition's mutable segment
	if err := part.Insert(rows); err != nil {
		return err
	}

	// Track row locations for each inserted row
	numRows := len(rows)
	for i := 0; i < numRows; i++ {
		t.rowIDToLoc[startRowID+i] = RowLocation{
			PartitionKey: part.Key(),
			IsMutable:    true,
			SegmentIndex: 0,
			LocalIndex:   int(part.NextInsertIndex()) - numRows + i,
		}
	}

	// Increment global row ID
	t.nextRowID += numRows

	return nil
}

// Seal seals all partitions' mutable segments if they should be sealed.
// Returns true if any sealing occurred, false otherwise.
func (t *Table) Seal(dir string) (bool, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	sealed := false
	for _, part := range t.pm.Partitions() {
		if part.ShouldSeal() {
			s, err := part.Seal(dir)
			if err != nil {
				return sealed, fmt.Errorf("Seal: %w", err)
			}
			if s {
				sealed = true
			}
		}
	}
	return sealed, nil
}

// SelectColumns returns column data for the requested columns.
// It applies pending updates and filters deleted rows.
func (t *Table) SelectColumns(columns []string) (map[string]interface{}, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	result := make(map[string]interface{})
	// Initialize result columns as empty slices
	for _, col := range columns {
		result[col] = []interface{}{}
	}

	// Track global row index as we read
	globalRowIdx := 0

	// Read from all partitions (mutable and immutable)
	for _, part := range t.pm.Partitions() {
		// Read from mutable (unsealed data)
		mutableCols, err := part.ReadMutableColumns(columns)
		if err != nil {
			return nil, err
		}

		mutableRowCount := int(part.MutableRowCount())
		for localIdx := 0; localIdx < mutableRowCount; localIdx++ {
			// Check if this row is deleted
			if t.pendingDel[globalRowIdx] {
				globalRowIdx++
				continue
			}

			// Apply pending updates if any
			rowUpdates := t.pendingUpd[globalRowIdx]

			// Add each column value (applying updates if needed)
			for _, colName := range columns {
				var val interface{}
				if colData, ok := mutableCols[colName].([]interface{}); ok && localIdx < len(colData) {
					val = colData[localIdx]
					// Apply pending update for this column if exists
					if rowUpdates != nil {
						if updVal, ok := rowUpdates[colName]; ok {
							val = updVal
						}
					}
				}
				result[colName] = append(result[colName].([]interface{}), val)
			}
			globalRowIdx++
		}

		// Read from immutable segments
		for _, imm := range part.ImmutableSegments() {
			cols, err := imm.ReadColumns(columns)
			if err != nil {
				return nil, err
			}
			immRowCount := int(imm.RowCount())
			for localIdx := 0; localIdx < immRowCount; localIdx++ {
				// Check if this row is deleted
				if t.pendingDel[globalRowIdx] {
					globalRowIdx++
					continue
				}

				// Apply pending updates if any
				rowUpdates := t.pendingUpd[globalRowIdx]

				// Add each column value (applying updates if needed)
				for _, colName := range columns {
					var val interface{}
					if colData, ok := cols[colName].([]interface{}); ok && localIdx < len(colData) {
						val = colData[localIdx]
						// Apply pending update for this column if exists
						if rowUpdates != nil {
							if updVal, ok := rowUpdates[colName]; ok {
								val = updVal
							}
						}
					}
					result[colName] = append(result[colName].([]interface{}), val)
				}
				globalRowIdx++
			}
		}
	}

	return result, nil
}

// appendSlice appends src to dst, handling type conversion to []interface{}.
func appendSlice(dst, src interface{}) []interface{} {
	dstSlice := toInterfaceSlice(dst)
	srcSlice := toInterfaceSlice(src)
	return append(dstSlice, srcSlice...)
}

// toInterfaceSlice converts a typed slice to []interface{}.
func toInterfaceSlice(v interface{}) []interface{} {
	switch s := v.(type) {
	case []interface{}:
		return s
	case []int64:
		result := make([]interface{}, len(s))
		for i, val := range s {
			result[i] = val
		}
		return result
	case []int32:
		result := make([]interface{}, len(s))
		for i, val := range s {
			result[i] = val
		}
		return result
	case []float32:
		result := make([]interface{}, len(s))
		for i, val := range s {
			result[i] = val
		}
		return result
	case []float64:
		result := make([]interface{}, len(s))
		for i, val := range s {
			result[i] = val
		}
		return result
	case [][]float32:
		result := make([]interface{}, len(s))
		for i, val := range s {
			result[i] = val
		}
		return result
	default:
		return nil
	}
}

// RowCount returns the total number of rows in all partitions.
func (t *Table) RowCount() uint64 {
	t.mu.RLock()
	defer t.mu.RUnlock()

	count := uint64(0)
	for _, part := range t.pm.Partitions() {
		count += part.RowCount()
	}
	return count
}

// UpdateRow updates a specific row at the given index.
// Updates are stored as pending and applied during SelectColumns.
// This approach works for both mutable and immutable segments.
func (t *Table) UpdateRow(rowIdx int, updates map[string]interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Just store the pending update - SelectColumns will apply it
	if t.pendingUpd == nil {
		t.pendingUpd = make(map[int]map[string]interface{})
	}
	// Merge with existing pending updates
	if existing, ok := t.pendingUpd[rowIdx]; ok {
		for k, v := range updates {
			existing[k] = v
		}
	} else {
		t.pendingUpd[rowIdx] = updates
	}
	return nil
}

// DeleteRow deletes a row at the given index.
func (t *Table) DeleteRow(rowIdx int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.pendingDel == nil {
		t.pendingDel = make(map[int]bool)
	}
	t.pendingDel[rowIdx] = true
	return nil
}

// LoadImmutableSegments loads existing MSI files from partition directories.
func (t *Table) LoadImmutableSegments(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		// Check if it's a bucket directory (bucket_<timestamp>)
		name := entry.Name()
		if len(name) < 7 || name[:7] != "bucket_" {
			continue
		}

		// Create partition for this bucket
		part := t.pm.GetOrCreatePartition(0) // ts=0 will create bucket_0

		// Load immutable segments from this partition
		if err := part.LoadImmutableSegments(dir); err != nil {
			return err
		}
	}
	return nil
}

// Name returns the table name.
func (t *Table) Name() string {
	return t.name
}

// DB returns the database name.
func (t *Table) DB() string {
	return t.db
}

// ImmutableSegments returns all immutable segments from all partitions.
func (t *Table) ImmutableSegments() []*segment.ImmutableSegment {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var result []*segment.ImmutableSegment
	for _, part := range t.pm.Partitions() {
		result = append(result, part.ImmutableSegments()...)
	}
	return result
}

// SchedulerSegments returns segments grouped by partition for merge scheduling.
func (t *Table) SchedulerSegments() map[string][]*segment.ImmutableSegment {
	t.mu.RLock()
	defer t.mu.RUnlock()
	result := make(map[string][]*segment.ImmutableSegment)
	for _, part := range t.pm.Partitions() {
		result[part.Key()] = part.ImmutableSegments()
	}
	return result
}

// Partitions returns all partitions.
func (t *Table) Partitions() []*partition.Partition {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.pm.Partitions()
}

// ReplaceImmutableSegments replaces old segments with a new one after merge.
// It finds which partition contains the old segments and calls ReplaceImmutables on it.
func (t *Table) ReplaceImmutableSegments(oldPaths []string, newPath string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Build set of old paths
	oldSet := make(map[string]bool)
	for _, p := range oldPaths {
		oldSet[p] = true
	}

	// Find the partition that contains the old segments
	for _, part := range t.pm.Partitions() {
		for _, imm := range part.ImmutableSegments() {
			if oldSet[imm.Path()] {
				return part.ReplaceImmutables(oldPaths, newPath)
			}
		}
	}
	return fmt.Errorf("ReplaceImmutableSegments: no partition found for old paths")
}