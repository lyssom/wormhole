package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"wormhole/server/engine/storage/msi"
	"wormhole/server/engine/storage/segment"
)

// TableMetadata contains table-level metadata (columns and row count).
type TableMetadata = msi.TableMetadata

// Table represents an open table with its mutable segment.
type Table struct {
	name       string
	db         string
	metadata   *TableMetadata
	mutable    *segment.MutableSegment
	immutables []*segment.ImmutableSegment
	opts       *TableOptions
	mu         sync.RWMutex

	// Staging rows for reading before seal (column-oriented)
	staging map[string][]interface{}
}

// newTable creates a new Table.
func newTable(db, name string, meta *TableMetadata, opts *TableOptions) *Table {
	if opts == nil {
		opts = DefaultTableOptions()
	}
	return &Table{
		name:       name,
		db:         db,
		metadata:   meta,
		mutable:    segment.NewMutableSegment(name),
		immutables: make([]*segment.ImmutableSegment, 0),
		opts:       opts,
		staging:    make(map[string][]interface{}),
	}
}

// Insert inserts rows into the table's mutable segment.
func (t *Table) Insert(rows []map[string]interface{}) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Convert rows to column-oriented format and stage for reading
	for _, row := range rows {
		for colName, val := range row {
			t.staging[colName] = append(t.staging[colName], val)
		}
	}

	// Only append new rows to mutable segment for persistence
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

	return t.mutable.Append(batch)
}

// Seal seals the mutable segment if it should be sealed.
// Returns true if sealing occurred, false otherwise.
func (t *Table) Seal(dir string) (bool, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.mutable.ShouldSeal() {
		return false, nil
	}

	if t.mutable.RowCount() == 0 {
		return false, nil
	}

	// Create file path for the sealed segment
	segNum := len(t.immutables)
	filePath := filepath.Join(dir, fmt.Sprintf("%d.msi", segNum))

	// Seal to file
	_, err := t.mutable.SealToFile(filePath)
	if err != nil {
		return false, fmt.Errorf("Seal: SealToFile failed: %w", err)
	}

	// Open as immutable segment
	imm, err := segment.OpenImmutableSegment(filePath)
	if err != nil {
		return false, fmt.Errorf("Seal: OpenImmutableSegment failed: %w", err)
	}
	t.immutables = append(t.immutables, imm)

	// Clear staging after successful seal (data now in immutable)
	t.staging = make(map[string][]interface{})

	return true, nil
}

// SelectColumns returns column data for the requested columns.
func (t *Table) SelectColumns(columns []string) (map[string]interface{}, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	result := make(map[string]interface{})

	// Read from staging (mutable, unsealed data)
	for _, col := range columns {
		if colData, ok := t.staging[col]; ok {
			result[col] = colData
		}
	}

	// Read from immutable segments
	for _, imm := range t.immutables {
		cols, err := imm.ReadColumns(columns)
		if err != nil {
			return nil, err
		}
		// Merge into result - prepend staging data before immutable data
		for colName, colData := range cols {
			if _, ok := result[colName]; !ok {
				result[colName] = colData
			} else {
				// Prepend immutable data after staging data
				result[colName] = append(result[colName].([]interface{}), colData.([]interface{})...)
			}
		}
	}

	return result, nil
}

// RowCount returns the total number of rows (staging + immutable).
func (t *Table) RowCount() uint64 {
	t.mu.RLock()
	defer t.mu.RUnlock()

	count := uint64(0)
	for _, col := range t.staging {
		if len(col) > int(count) {
			count = uint64(len(col))
		}
	}
	for _, imm := range t.immutables {
		count += imm.RowCount()
	}
	return count
}

// LoadImmutableSegments loads existing MSI files from dir as immutable segments.
func (t *Table) LoadImmutableSegments(dir string) error {
	entries, err := os.ReadDir(dir)
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
		path := filepath.Join(dir, entry.Name())
		imm, err := segment.OpenImmutableSegment(path)
		if err != nil {
			continue // skip corrupted files
		}
		t.immutables = append(t.immutables, imm)
	}
	return nil
}