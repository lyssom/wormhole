package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// StorageEngine provides the top-level API for inserting and selecting data.
type StorageEngine struct {
	rootDir string
	tables  map[string]*Table // key: "db/table"
	opts    *Options
	mu      sync.RWMutex
}

// NewStorageEngine creates a new StorageEngine rooted at rootDir.
func NewStorageEngine(rootDir string, opts *Options) *StorageEngine {
	if opts == nil {
		opts = DefaultOptions()
	}
	return &StorageEngine{
		rootDir: rootDir,
		tables:  make(map[string]*Table),
		opts:    opts,
	}
}

// tableKey returns the internal key for a database.table.
func tableKey(db, table string) string {
	return db + "/" + table
}

// dataPath returns the directory path for a table's data files.
func (e *StorageEngine) dataPath(db, table string) string {
	return filepath.Join(e.rootDir, db, table)
}

// ensureDir creates the directory if it doesn't exist.
func ensureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// CreateTable creates a new table in the given database.
func (e *StorageEngine) CreateTable(db, table string, meta *TableMetadata) error {
	key := tableKey(db, table)
	if _, exists := e.tables[key]; exists {
		return fmt.Errorf("CreateTable: table %s.%s already exists", db, table)
	}

	// Create table directory
	tablePath := e.dataPath(db, table)
	if err := ensureDir(tablePath); err != nil {
		return fmt.Errorf("CreateTable: mkdir failed: %w", err)
	}

	// Write table metadata to table.json
	metaPath := filepath.Join(tablePath, "table.json")
	if err := writeTableMetadata(metaPath, meta); err != nil {
		return fmt.Errorf("CreateTable: write metadata failed: %w", err)
	}

	// Create and register the table
	e.tables[key] = newTable(db, table, meta, nil, e.opts.BucketDurationMs)

	return nil
}

// GetOrCreateTable returns an existing table or creates a new one.
func (e *StorageEngine) GetOrCreateTable(db, table string, meta *TableMetadata) (*Table, error) {
	key := tableKey(db, table)
	e.mu.Lock()
	defer e.mu.Unlock()
	if t, ok := e.tables[key]; ok {
		return t, nil
	}

	// Create table directory
	tablePath := e.dataPath(db, table)
	if err := ensureDir(tablePath); err != nil {
		return nil, fmt.Errorf("GetOrCreateTable: mkdir failed: %w", err)
	}

	// Write table metadata to table.json
	metaPath := filepath.Join(tablePath, "table.json")
	if err := writeTableMetadata(metaPath, meta); err != nil {
		return nil, fmt.Errorf("GetOrCreateTable: write metadata failed: %w", err)
	}

	// Create and register the table
	t := newTable(db, table, meta, nil, e.opts.BucketDurationMs)
	e.tables[key] = t

	return t, nil
}

// Insert inserts rows into the specified table.
// For AtLeastOnce: writes directly to Mutable Segment
// For ExactlyOnce: writes to WAL first, then Mutable Segment
func (e *StorageEngine) Insert(db, table string, rows []map[string]interface{}) error {
	key := tableKey(db, table)
	t, exists := e.tables[key]
	if !exists {
		return fmt.Errorf("Insert: table %s.%s not found", db, table)
	}

	return t.Insert(rows)
}

// SelectColumns reads the specified columns from the table.
// Returns a map from column name to column data.
func (e *StorageEngine) SelectColumns(db, table string, columns []string) (map[string]interface{}, error) {
	key := tableKey(db, table)
	t, exists := e.tables[key]
	if !exists {
		return nil, fmt.Errorf("SelectColumns: table %s.%s not found", db, table)
	}

	return t.SelectColumns(columns)
}

// UpdateRow updates a specific row at index rowIdx in the table.
func (e *StorageEngine) UpdateRow(db, table string, rowIdx int, updates map[string]interface{}) error {
	key := tableKey(db, table)
	t, exists := e.tables[key]
	if !exists {
		return fmt.Errorf("UpdateRow: table %s.%s not found", db, table)
	}

	return t.UpdateRow(rowIdx, updates)
}

// DeleteRow deletes a row at index rowIdx from the table.
func (e *StorageEngine) DeleteRow(db, table string, rowIdx int) error {
	key := tableKey(db, table)
	t, exists := e.tables[key]
	if !exists {
		return fmt.Errorf("DeleteRow: table %s.%s not found", db, table)
	}

	return t.DeleteRow(rowIdx)
}

// writeTableMetadata writes the table metadata to a JSON file.
func writeTableMetadata(path string, meta *TableMetadata) error {
	data, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Tables returns all tables in the engine.
func (e *StorageEngine) Tables() []*Table {
	e.mu.RLock()
	defer e.mu.RUnlock()
	tables := make([]*Table, 0, len(e.tables))
	for _, t := range e.tables {
		tables = append(tables, t)
	}
	return tables
}

// RootDir returns the engine's root directory.
func (e *StorageEngine) RootDir() string {
	return e.rootDir
}