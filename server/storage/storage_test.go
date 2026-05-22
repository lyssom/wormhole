package storage

import (
	"os"
	"testing"

	"wormhole/server/storage/msi"
)

func TestStorageInsertAndSelect(t *testing.T) {
	tmp := os.TempDir()
	engine := NewStorageEngine(tmp, &Options{
		MutableMaxRows:   100000,
		BucketDurationMs: 3600000,
	})

	// Create table
	err := engine.CreateTable("db1", "t1", &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "vec", Type: msi.TypeVectorF32, VectorDim: 128},
		},
	})
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}

	// Insert
	err = engine.Insert("db1", "t1", []map[string]interface{}{
		{"id": int64(1), "vec": make([]float32, 128)},
	})
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	// Read
	cols, err := engine.SelectColumns("db1", "t1", []string{"id", "vec"})
	if err != nil {
		t.Fatalf("SelectColumns failed: %v", err)
	}
	idCol, ok := cols["id"].([]interface{})
	if !ok {
		t.Fatalf("cols[id] is not []interface{}")
	}
	if len(idCol) != 1 {
		t.Errorf("row count: got %d, want 1", len(idCol))
	}
}