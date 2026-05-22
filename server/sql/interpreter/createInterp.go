package interpreter

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"wormhole/server/config"
	"wormhole/server/sql/statement"
)

type Interpreter struct {
	UsedDB string
}

// sourceDir is computed once at package load time using the actual source file location.
// This is used to resolve the config file path relative to the binary.
var sourceDir string

func init() {
	var filename string
	_, filename, _, _ = runtime.Caller(0)
	sourceDir = path.Dir(filename)
}

// getSourceDir returns the directory containing this source file.
func getSourceDir() string {
	return sourceDir
}

// getConfig returns the global config, loading it from the standard location if needed.
func getConfig() *config.Config {
	cfg := config.Global()
	if cfg == nil {
		// Try to load from standard location: <sourceDir>/../config.yaml
		configPath := filepath.Join(getSourceDir(), "..", "config.yaml")
		if cfg2, err := config.Load(configPath); err == nil {
			config.SetGlobal(cfg2)
			return cfg2
		}
		// Fallback to defaults
		cfg = &config.Config{
			Storage: config.StorageConfig{
				DataPath:         "./data",
				MutableMaxRows:   100000,
				BucketDurationMs: 3600000,
			},
		}
		config.SetGlobal(cfg)
	}
	return cfg
}

func (i Interpreter) GetDBabPath() string {
	cfg := getConfig()
	return filepath.Join(cfg.GetDataPath(), i.UsedDB)
}

func (i Interpreter) ExecuteCreateDatabase(dbName string) []byte {
	cfg := getConfig()
	dataPath := filepath.Join(cfg.GetDataPath(), dbName)
	os.MkdirAll(dataPath, os.ModePerm)
	return []byte("ok")
}

func (i Interpreter) ExecuteCreateTable(table statement.CreateTableStatement) []byte {
	dbPath := i.GetDBabPath()
	os.MkdirAll(path.Join(dbPath, table.TableName), os.ModePerm)

	filePtr, err := os.Create(path.Join(dbPath, table.TableName, "table.json"))
	fmt.Println(err)
	defer filePtr.Close()
	encoder := json.NewEncoder(filePtr)
	encoder.Encode(table)
	return []byte("ok")
}
