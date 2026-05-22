package server

import (
	"os"
	"path/filepath"
	"testing"

	"wormhole/server/config"
	"wormhole/server/sql/statement"
)

// Helper to set up a clean test environment
func setupTestDB(t *testing.T) (cleanup func()) {
	tmpDir := filepath.Join(os.TempDir(), "wormhole_e2e_test")
	config.SetGlobal(&config.Config{
		Storage: config.StorageConfig{
			DataPath:         tmpDir,
			MutableMaxRows:   100,
			BucketDurationMs: 3600000,
		},
	})
	return func() {
		os.RemoveAll(tmpDir)
	}
}

// parse is a helper to parse SQL and return the statement
func parse(sql string) statement.Statement {
	return ExecuteParser(sql).(statement.Statement)
}

// TestE2E_CreateDatabase tests CREATE DATABASE end-to-end
func TestE2E_CreateDatabase(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	sql := "CREATE DATABASE testdb"
	stmt := parse(sql)
	dbStmt, ok := stmt.(statement.CreateDatabaseStatement)
	if !ok {
		t.Fatalf("expected CreateDatabaseStatement, got %T", stmt)
	}

	if dbStmt.Database != "testdb" {
		t.Errorf("expected database 'testdb', got '%s'", dbStmt.Database)
	}

	result := Execute(stmt, "")
	if string(result) != "ok" {
		t.Errorf("Execute failed: %s", result)
	}

	// Verify database directory was created
	dataPath := filepath.Join(os.TempDir(), "wormhole_e2e_test", "testdb")
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		t.Errorf("database directory not created: %s", dataPath)
	}
}

// TestE2E_CreateTable tests CREATE TABLE end-to-end
func TestE2E_CreateTable(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// First create the database
	Execute(parse("CREATE DATABASE testdb"), "")

	sql := "CREATE TABLE users (id INT64, name STRING, age INT64)"
	stmt := parse(sql)
	tableStmt, ok := stmt.(statement.CreateTableStatement)
	if !ok {
		t.Fatalf("expected CreateTableStatement, got %T", stmt)
	}

	if tableStmt.TableName != "users" {
		t.Errorf("expected table 'users', got '%s'", tableStmt.TableName)
	}

	if len(tableStmt.Columns) != 3 {
		t.Errorf("expected 3 columns, got %d", len(tableStmt.Columns))
	}

	result := Execute(stmt, "testdb")
	if string(result) != "ok" {
		t.Errorf("Execute failed: %s", result)
	}

	// Verify table metadata file was created
	tablePath := filepath.Join(os.TempDir(), "wormhole_e2e_test", "testdb", "users", "table.json")
	if _, err := os.Stat(tablePath); os.IsNotExist(err) {
		t.Errorf("table.json not created: %s", tablePath)
	}
}

// TestE2E_Select_ParserOnly tests SELECT parsing without execution
// (Full select execution requires table to be registered in storage engine)
func TestE2E_Select_ParserOnly(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// Setup
	Execute(parse("CREATE DATABASE testdb"), "")
	Execute(parse("CREATE TABLE users (id INT64, name STRING, age INT64)"), "testdb")

	// Select with WHERE - parser test
	sql := "SELECT id, name, age FROM users WHERE age > 25"
	stmt := parse(sql)
	qs, ok := stmt.(statement.QueryStatement)
	if !ok {
		t.Fatalf("expected QueryStatement, got %T", stmt)
	}

	if len(qs.SelectComponent.ResultColumns) != 3 {
		t.Errorf("expected 3 result columns, got %d", len(qs.SelectComponent.ResultColumns))
	}

	if qs.FromComponent.FromClause != "users" {
		t.Errorf("expected from 'users', got '%s'", qs.FromComponent.FromClause)
	}

	if qs.WhereCondition.Condition == nil {
		t.Errorf("expected WHERE condition, got nil")
	}

	t.Logf("Parsed SELECT: %+v", qs)
}

// TestE2E_Update_ParserOnly tests UPDATE parsing without execution
func TestE2E_Update_ParserOnly(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	Execute(parse("CREATE DATABASE testdb"), "")
	Execute(parse("CREATE TABLE users (id INT64, name STRING, age INT64)"), "testdb")

	sql := "UPDATE users SET age = 26 WHERE id = 2"
	stmt := parse(sql)
	us, ok := stmt.(statement.UpdateStatement)
	if !ok {
		t.Fatalf("expected UpdateStatement, got %T", stmt)
	}

	if us.TableName != "users" {
		t.Errorf("expected table 'users', got '%s'", us.TableName)
	}

	if len(us.SetClauses) != 1 {
		t.Errorf("expected 1 set clause, got %d", len(us.SetClauses))
	}

	if us.SetClauses[0].Column != "age" {
		t.Errorf("expected column 'age', got '%s'", us.SetClauses[0].Column)
	}

	if us.WhereCondition == nil {
		t.Errorf("expected WHERE condition, got nil")
	}

	t.Logf("Parsed UPDATE: %+v", us)
}

// TestE2E_Delete_ParserOnly tests DELETE parsing without execution
func TestE2E_Delete_ParserOnly(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	Execute(parse("CREATE DATABASE testdb"), "")
	Execute(parse("CREATE TABLE users (id INT64, name STRING, age INT64)"), "testdb")

	sql := "DELETE FROM users WHERE id = 1"
	stmt := parse(sql)
	ds, ok := stmt.(statement.DeleteStatement)
	if !ok {
		t.Fatalf("expected DeleteStatement, got %T", stmt)
	}

	if ds.TableName != "users" {
		t.Errorf("expected table 'users', got '%s'", ds.TableName)
	}

	if ds.WhereCondition == nil {
		t.Errorf("expected WHERE condition, got nil")
	}

	t.Logf("Parsed DELETE: %+v", ds)
}

// TestE2E_Insert_ParserOnly tests INSERT parsing without execution
func TestE2E_Insert_ParserOnly(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	Execute(parse("CREATE DATABASE testdb"), "")
	Execute(parse("CREATE TABLE users (id INT64, name STRING, age INT64)"), "testdb")

	sql := "INSERT INTO users (id, name, age) VALUES (1, 'Alice', 30), (2, 'Bob', 25)"
	stmt := parse(sql)
	is, ok := stmt.(statement.InsertStatement)
	if !ok {
		t.Fatalf("expected InsertStatement, got %T", stmt)
	}

	if is.TableName != "users" {
		t.Errorf("expected table 'users', got '%s'", is.TableName)
	}

	if len(is.Columns) != 3 {
		t.Errorf("expected 3 columns, got %d", len(is.Columns))
	}

	if len(is.Values) != 2 {
		t.Errorf("expected 2 rows, got %d", len(is.Values))
	}

	t.Logf("Parsed INSERT: %+v", is)
}