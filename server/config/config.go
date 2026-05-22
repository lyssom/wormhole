package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	DataPath         string `yaml:"data_path"`
	MutableMaxRows   int    `yaml:"mutable_max_rows"`
	BucketDurationMs int    `yaml:"bucket_duration_ms"`
}

var globalConfig *Config

// Load reads the config file from the given path.
func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Set defaults
	if cfg.Storage.DataPath == "" {
		cfg.Storage.DataPath = "./data"
	}
	if cfg.Storage.MutableMaxRows == 0 {
		cfg.Storage.MutableMaxRows = 100000
	}
	if cfg.Storage.BucketDurationMs == 0 {
		cfg.Storage.BucketDurationMs = 3600000
	}

	return &cfg, nil
}

// LoadFromDir loads config from server/config.yaml relative to the server binary.
func LoadFromDir(serverDir string) (*Config, error) {
	configPath := filepath.Join(serverDir, "config.yaml")
	return Load(configPath)
}

// Global returns the global config instance.
func Global() *Config {
	return globalConfig
}

// SetGlobal sets the global config instance.
func SetGlobal(cfg *Config) {
	globalConfig = cfg
}

// GetDataPath returns the data storage path.
func (c *Config) GetDataPath() string {
	return c.Storage.DataPath
}

// GetMutableMaxRows returns the mutable segment max rows.
func (c *Config) GetMutableMaxRows() int {
	return c.Storage.MutableMaxRows
}

// GetBucketDurationMs returns the bucket duration in milliseconds.
func (c *Config) GetBucketDurationMs() int {
	return c.Storage.BucketDurationMs
}
