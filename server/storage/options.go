package storage

// ConsistencyLevel determines the durability guarantee for inserts.
type ConsistencyLevel int

const (
	AtLeastOnce  ConsistencyLevel = iota // Default: write to mutable segment directly
	ExactlyOnce                          // Write to WAL first, then mutable segment
)

// Options configures the StorageEngine.
type Options struct {
	MutableMaxRows   int   // Default 100000
	BucketDurationMs int64 // Time partition bucket size in ms, default 3600000 (1 hour)
}

// TableOptions configures per-table behavior.
type TableOptions struct {
	Consistency ConsistencyLevel // Default AtLeastOnce
}

// DefaultOptions returns the default engine options.
func DefaultOptions() *Options {
	return &Options{
		MutableMaxRows:   100000,
		BucketDurationMs: 3600000, // 1 hour
	}
}

// DefaultTableOptions returns the default table options.
func DefaultTableOptions() *TableOptions {
	return &TableOptions{
		Consistency: AtLeastOnce,
	}
}
