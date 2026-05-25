package partition

import (
	"time"
)

// PartitionMetadata tracks metadata about a partition.
type PartitionMetadata struct {
	Key         string    `json:"key"`
	TotalRows   uint64    `json:"total_rows"`
	TotalSize   int64     `json:"total_size"`
	SegmentCount int      `json:"segment_count"`
	LastModified time.Time `json:"last_modified"`
}

// SegmentMeta holds information about a single segment.
type SegmentMeta struct {
	Path       string `json:"path"`
	RowCount   uint64 `json:"row_count"`
	SizeBytes  int64  `json:"size_bytes"`
	TsMin      uint64 `json:"ts_min"`
	TsMax      uint64 `json:"ts_max"`
}
