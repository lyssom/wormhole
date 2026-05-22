package storage

import (
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

	// Also append to mutable segment for persistence
	batch := make(map[string]interface{})
	for colName, vals := range t.staging {
		batch[colName] = vals
	}

	return t.mutable.Append(batch)
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
		// Merge into result
		for colName, colData := range cols {
			if _, ok := result[colName]; !ok {
				result[colName] = colData
			} else {
				// Merge - append immutable data to staging data
				result[colName] = colData
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