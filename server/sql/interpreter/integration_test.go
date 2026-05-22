package interpreter

import (
	"os"
	"path/filepath"
	"testing"

	"wormhole/server/config"
	"wormhole/server/storage"
	"wormhole/server/storage/msi"
)

// TestIntegrationStorageCRUD tests the full CRUD pipeline at storage engine level.
func TestIntegrationStorageCRUD(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "wormhole_integration_test")
	defer os.RemoveAll(tmpDir)

	engine := storage.NewStorageEngine(tmpDir, &storage.Options{
		MutableMaxRows:   100,
		BucketDurationMs: 3600000,
	})

	dbName := "testdb"
	tableName := "users"

	t.Run("CreateTable", func(t *testing.T) {
		err := engine.CreateTable(dbName, tableName, &storage.TableMetadata{
			Columns: []*msi.ColumnMeta{
				{Name: "id", Type: msi.TypeInt64},
				{Name: "name", Type: msi.TypeByteArray},
				{Name: "age", Type: msi.TypeInt64},
			},
		})
		if err != nil {
			t.Fatalf("CreateTable failed: %v", err)
		}
	})

	t.Run("Insert", func(t *testing.T) {
		rows := []map[string]interface{}{
			{"id": int64(1), "name": "Alice", "age": int64(30)},
			{"id": int64(2), "name": "Bob", "age": int64(25)},
			{"id": int64(3), "name": "Charlie", "age": int64(35)},
		}
		err := engine.Insert(dbName, tableName, rows)
		if err != nil {
			t.Fatalf("Insert failed: %v", err)
		}
	})

	t.Run("SelectAll", func(t *testing.T) {
		cols, err := engine.SelectColumns(dbName, tableName, []string{"id", "name", "age"})
		if err != nil {
			t.Fatalf("SelectColumns failed: %v", err)
		}

		idCol, ok := cols["id"].([]interface{})
		if !ok {
			t.Fatalf("id column is not []interface{}")
		}
		if len(idCol) != 3 {
			t.Errorf("expected 3 rows, got %d", len(idCol))
		}

		nameCol, ok := cols["name"].([]interface{})
		if !ok {
			t.Fatalf("name column is not []interface{}")
		}
		if len(nameCol) != 3 {
			t.Errorf("expected 3 names, got %d", len(nameCol))
		}
	})

	t.Run("SelectColumns", func(t *testing.T) {
		cols, err := engine.SelectColumns(dbName, tableName, []string{"id", "age"})
		if err != nil {
			t.Fatalf("SelectColumns failed: %v", err)
		}

		if _, ok := cols["name"]; ok {
			t.Error("name column should not be selected")
		}

		idCol, ok := cols["id"].([]interface{})
		if !ok {
			t.Fatalf("id column is not []interface{}")
		}
		if len(idCol) != 3 {
			t.Errorf("expected 3 rows, got %d", len(idCol))
		}
	})

	t.Run("UpdateRow", func(t *testing.T) {
		updates := map[string]interface{}{
			"id":   int64(2),
			"name": "Bob",
			"age":  int64(26),
		}
		err := engine.UpdateRow(dbName, tableName, 1, updates)
		if err != nil {
			t.Fatalf("UpdateRow failed: %v", err)
		}

		cols, err := engine.SelectColumns(dbName, tableName, []string{"age"})
		if err != nil {
			t.Fatalf("SelectColumns after update failed: %v", err)
		}
		ageCol := cols["age"].([]interface{})
		if ageCol[1].(int64) != 26 {
			t.Errorf("expected age 26 at index 1, got %v", ageCol[1])
		}
	})

	t.Run("DeleteRow", func(t *testing.T) {
		err := engine.DeleteRow(dbName, tableName, 0)
		if err != nil {
			t.Fatalf("DeleteRow failed: %v", err)
		}

		cols, err := engine.SelectColumns(dbName, tableName, []string{"id"})
		if err != nil {
			t.Fatalf("SelectColumns after delete failed: %v", err)
		}
		idCol := cols["id"].([]interface{})
		if len(idCol) != 2 {
			t.Errorf("expected 2 rows after delete, got %d", len(idCol))
		}
	})

	t.Run("FinalState", func(t *testing.T) {
		cols, err := engine.SelectColumns(dbName, tableName, []string{"id", "name", "age"})
		if err != nil {
			t.Fatalf("SelectColumns failed: %v", err)
		}

		idCol := cols["id"].([]interface{})
		nameCol := cols["name"].([]interface{})
		ageCol := cols["age"].([]interface{})

		if len(idCol) != 2 {
			t.Errorf("expected 2 rows, got %d", len(idCol))
		}

		if nameCol[0] != "Bob" {
			t.Errorf("expected first name Bob, got %v", nameCol[0])
		}
		if ageCol[0].(int64) != 26 {
			t.Errorf("expected Bob's age 26, got %v", ageCol[0])
		}
	})
}

// TestMultipleTables tests working with multiple tables
func TestMultipleTables(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "wormhole_multi_table_test")
	defer os.RemoveAll(tmpDir)

	engine := storage.NewStorageEngine(tmpDir, nil)
	dbName := "testdb"

	err := engine.CreateTable(dbName, "users", &storage.TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "name", Type: msi.TypeByteArray},
		},
	})
	if err != nil {
		t.Fatalf("CreateTable users failed: %v", err)
	}

	err = engine.CreateTable(dbName, "products", &storage.TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "price", Type: msi.TypeFloat64},
		},
	})
	if err != nil {
		t.Fatalf("CreateTable products failed: %v", err)
	}

	err = engine.Insert(dbName, "users", []map[string]interface{}{
		{"id": int64(1), "name": "Alice"},
		{"id": int64(2), "name": "Bob"},
	})
	if err != nil {
		t.Fatalf("Insert users failed: %v", err)
	}

	err = engine.Insert(dbName, "products", []map[string]interface{}{
		{"id": int64(1), "price": float64(9.99)},
		{"id": int64(2), "price": float64(19.99)},
	})
	if err != nil {
		t.Fatalf("Insert products failed: %v", err)
	}

	userCols, err := engine.SelectColumns(dbName, "users", []string{"id", "name"})
	if err != nil {
		t.Fatalf("SelectColumns users failed: %v", err)
	}
	if len(userCols["id"].([]interface{})) != 2 {
		t.Errorf("expected 2 users, got %d", len(userCols["id"].([]interface{})))
	}

	prodCols, err := engine.SelectColumns(dbName, "products", []string{"id", "price"})
	if err != nil {
		t.Fatalf("SelectColumns products failed: %v", err)
	}
	if len(prodCols["id"].([]interface{})) != 2 {
		t.Errorf("expected 2 products, got %d", len(prodCols["id"].([]interface{})))
	}
}

// TestInterpreterPathResolution verifies that GetDBabPath() returns a consistent
// path based on the global config.
func TestInterpreterPathResolution(t *testing.T) {
	// Set up a test config
	testDataPath := filepath.Join(os.TempDir(), "wormhole_test_data")
	defer os.RemoveAll(testDataPath)

	testCfg := &config.Config{
		Storage: config.StorageConfig{
			DataPath:         testDataPath,
			MutableMaxRows:   100000,
			BucketDurationMs: 3600000,
		},
	}
	config.SetGlobal(testCfg)

	// Create a new interpreter with a test DB
	interp := Interpreter{UsedDB: "testdb"}

	// GetDBabPath should return the same value every time
	path1 := interp.GetDBabPath()
	path2 := interp.GetDBabPath()

	if path1 != path2 {
		t.Errorf("GetDBabPath returned inconsistent paths: %s != %s", path1, path2)
	}

	// The path should be based on the config
	expectedPath := filepath.Join(testDataPath, "testdb")
	if path1 != expectedPath {
		t.Errorf("GetDBabPath returned wrong path: got %s, want %s", path1, expectedPath)
	}
}
