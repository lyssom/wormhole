package index

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// TableMetadata interface for MSI table metadata (avoids import cycle)
type TableMetadata interface {
	GetRowCount() uint64
	GetColumns() []ColumnMetaInfo
}

// ColumnMetaInfo interface for MSI column metadata
type ColumnMetaInfo interface {
	GetName() string
	GetType() int
}

// ColumnStats holds min/max statistics for a column within a row index block.
type ColumnStats struct {
	Name string

	// For int64 columns
	MinInt64, MaxInt64 *int64

	// For float64 columns
	MinFloat64, MaxFloat64 *float64

	// For float32 columns
	MinFloat32, MaxFloat32 *float32

	// For int32 columns
	MinInt32, MaxInt32 *int32
}

// SkipIndex provides min/max statistics for each row index block to enable index-based filtering.
type SkipIndex struct {
	Blocks []BlockStats
}

// BlockStats holds statistics for a single row index block.
type BlockStats struct {
	RowOffset uint64
	TsMin     uint64
	TsMax     uint64
	ColumnStats map[string]*ColumnStats
}

// BuildFromMSI builds a SkipIndex from MSI metadata and column data.
// It divides rows into blocks and computes min/max for each column per block.
func BuildFromMSI(meta TableMetadata, columns map[string]interface{}, blockSize int) *SkipIndex {
	if blockSize <= 0 {
		blockSize = 1000 // default block size
	}

	rowCount := int(meta.GetRowCount())
	blockCount := (rowCount + blockSize - 1) / blockSize

	blocks := make([]BlockStats, 0, blockCount)

	for blockIdx := 0; blockIdx < blockCount; blockIdx++ {
		startRow := blockIdx * blockSize
		endRow := startRow + blockSize
		if endRow > rowCount {
			endRow = rowCount
		}

		blockStats := BlockStats{
			RowOffset:   uint64(startRow),
			TsMin:       math.MaxUint64,
			TsMax:       0,
			ColumnStats: make(map[string]*ColumnStats),
		}

		// Compute stats for each column in this block
		for _, colMeta := range meta.GetColumns() {
			colData := columns[colMeta.GetName()]
			stats := computeColumnStats(colMeta.GetName(), colData, startRow, endRow)
			blockStats.ColumnStats[colMeta.GetName()] = stats

			// Update TsMin/TsMax if this column is "ts"
			if colMeta.GetName() == "ts" {
				if stats.MinInt64 != nil && uint64(*stats.MinInt64) < blockStats.TsMin {
					blockStats.TsMin = uint64(*stats.MinInt64)
				}
				if stats.MaxInt64 != nil && uint64(*stats.MaxInt64) > blockStats.TsMax {
					blockStats.TsMax = uint64(*stats.MaxInt64)
				}
			}
		}

		blocks = append(blocks, blockStats)
	}

	return &SkipIndex{Blocks: blocks}
}

// computeColumnStats computes min/max for a column slice [startRow, endRow).
func computeColumnStats(name string, colData interface{}, startRow, endRow int) *ColumnStats {
	stats := &ColumnStats{Name: name}

	switch data := colData.(type) {
	case []int64:
		if startRow >= len(data) {
			return stats
		}
		if endRow > len(data) {
			endRow = len(data)
		}
		if startRow >= endRow {
			return stats
		}
		minVal := data[startRow]
		maxVal := data[startRow]
		for i := startRow + 1; i < endRow; i++ {
			if data[i] < minVal {
				minVal = data[i]
			}
			if data[i] > maxVal {
				maxVal = data[i]
			}
		}
		stats.MinInt64 = &minVal
		stats.MaxInt64 = &maxVal

	case []int32:
		if startRow >= len(data) {
			return stats
		}
		if endRow > len(data) {
			endRow = len(data)
		}
		if startRow >= endRow {
			return stats
		}
		minVal := data[startRow]
		maxVal := data[startRow]
		for i := startRow + 1; i < endRow; i++ {
			if data[i] < minVal {
				minVal = data[i]
			}
			if data[i] > maxVal {
				maxVal = data[i]
			}
		}
		stats.MinInt32 = &minVal
		stats.MaxInt32 = &maxVal

	case []float64:
		if startRow >= len(data) {
			return stats
		}
		if endRow > len(data) {
			endRow = len(data)
		}
		if startRow >= endRow {
			return stats
		}
		minVal := data[startRow]
		maxVal := data[startRow]
		for i := startRow + 1; i < endRow; i++ {
			if data[i] < minVal {
				minVal = data[i]
			}
			if data[i] > maxVal {
				maxVal = data[i]
			}
		}
		stats.MinFloat64 = &minVal
		stats.MaxFloat64 = &maxVal

	case []float32:
		if startRow >= len(data) {
			return stats
		}
		if endRow > len(data) {
			endRow = len(data)
		}
		if startRow >= endRow {
			return stats
		}
		minVal := data[startRow]
		maxVal := data[startRow]
		for i := startRow + 1; i < endRow; i++ {
			if data[i] < minVal {
				minVal = data[i]
			}
			if data[i] > maxVal {
				maxVal = data[i]
			}
		}
		stats.MinFloat32 = &minVal
		stats.MaxFloat32 = &maxVal

	case [][]float32:
		// Vector column - can't do min/max comparison for skip
		// For now, we just mark that stats are not available
	}

	return stats
}

// CanSkipInt64 returns true if the column's values in this block
// have no overlap with the query range [minVal, maxVal].
func (s *BlockStats) CanSkipInt64(colName string, minVal, maxVal int64) bool {
	stats, ok := s.ColumnStats[colName]
	if !ok || stats.MinInt64 == nil || stats.MaxInt64 == nil {
		return false // can't skip if no stats
	}
	// Skip if block's max < query min OR block's min > query max
	return *stats.MaxInt64 < minVal || *stats.MinInt64 > maxVal
}

// CanSkipFloat64 returns true if the column's values in this block
// have no overlap with the query range [minVal, maxVal].
func (s *BlockStats) CanSkipFloat64(colName string, minVal, maxVal float64) bool {
	stats, ok := s.ColumnStats[colName]
	if !ok || stats.MinFloat64 == nil || stats.MaxFloat64 == nil {
		return false
	}
	return *stats.MaxFloat64 < minVal || *stats.MinFloat64 > maxVal
}

// CanSkip returns true if the block can be skipped for the given filter.
func (s *BlockStats) CanSkip(colName string, minVal, maxVal interface{}) bool {
	switch min := minVal.(type) {
	case int64:
		if max, ok := maxVal.(int64); ok {
			return s.CanSkipInt64(colName, min, max)
		}
	case float64:
		if max, ok := maxVal.(float64); ok {
			return s.CanSkipFloat64(colName, min, max)
		}
	}
	return false // can't determine, don't skip
}

// CanSkip returns true if ALL blocks can be skipped for the given filter.
// This is used for whole-segment skip (when segment has no relevant data).
func (si *SkipIndex) CanSkip(colName string, minVal, maxVal interface{}) bool {
	if len(si.Blocks) == 0 {
		return false
	}
	// If all blocks can be skipped, the whole segment can be skipped
	for i := range si.Blocks {
		if !si.Blocks[i].CanSkip(colName, minVal, maxVal) {
			return false
		}
	}
	return true
}

// BlockCount returns the number of blocks in the skip index.
func (si *SkipIndex) BlockCount() int {
	return len(si.Blocks)
}

const skipIndexMagic = "SKIP"
const skipIndexVersion = 1

// MarshalBinary serializes the SkipIndex to binary format.
func (si *SkipIndex) MarshalBinary() ([]byte, error) {
	// Estimate size: header (12) + blocks * (24 + colstats)
	estSize := 12 + len(si.Blocks)*32
	buf := make([]byte, 0, estSize)

	// Write magic
	buf = append(buf, skipIndexMagic...)
	buf = binary.LittleEndian.AppendUint32(buf, skipIndexVersion)
	buf = binary.LittleEndian.AppendUint32(buf, uint32(len(si.Blocks)))

	for _, block := range si.Blocks {
		buf = binary.LittleEndian.AppendUint64(buf, block.RowOffset)
		buf = binary.LittleEndian.AppendUint64(buf, block.TsMin)
		buf = binary.LittleEndian.AppendUint64(buf, block.TsMax)
		buf = binary.LittleEndian.AppendUint32(buf, uint32(len(block.ColumnStats)))

		for name, stats := range block.ColumnStats {
			nameBytes := []byte(name)
			buf = binary.LittleEndian.AppendUint32(buf, uint32(len(nameBytes)))
			buf = append(buf, nameBytes...)

			// Write stats flags and values
			buf = appendBool(buf, stats.MinInt64 != nil)
			if stats.MinInt64 != nil {
				buf = binary.LittleEndian.AppendUint64(buf, uint64(*stats.MinInt64))
			}
			buf = appendBool(buf, stats.MaxInt64 != nil)
			if stats.MaxInt64 != nil {
				buf = binary.LittleEndian.AppendUint64(buf, uint64(*stats.MaxInt64))
			}

			buf = appendBool(buf, stats.MinFloat64 != nil)
			if stats.MinFloat64 != nil {
				buf = appendFloat64(buf, *stats.MinFloat64)
			}
			buf = appendBool(buf, stats.MaxFloat64 != nil)
			if stats.MaxFloat64 != nil {
				buf = appendFloat64(buf, *stats.MaxFloat64)
			}

			buf = appendBool(buf, stats.MinFloat32 != nil)
			if stats.MinFloat32 != nil {
				buf = appendFloat32(buf, *stats.MinFloat32)
			}
			buf = appendBool(buf, stats.MaxFloat32 != nil)
			if stats.MaxFloat32 != nil {
				buf = appendFloat32(buf, *stats.MaxFloat32)
			}

			buf = appendBool(buf, stats.MinInt32 != nil)
			if stats.MinInt32 != nil {
				buf = binary.LittleEndian.AppendUint32(buf, uint32(*stats.MinInt32))
			}
			buf = appendBool(buf, stats.MaxInt32 != nil)
			if stats.MaxInt32 != nil {
				buf = binary.LittleEndian.AppendUint32(buf, uint32(*stats.MaxInt32))
			}
		}
	}

	return buf, nil
}

func appendBool(buf []byte, v bool) []byte {
	if v {
		return append(buf, 1)
	}
	return append(buf, 0)
}

func appendFloat64(buf []byte, v float64) []byte {
	return binary.LittleEndian.AppendUint64(buf, math.Float64bits(v))
}

func appendFloat32(buf []byte, v float32) []byte {
	return binary.LittleEndian.AppendUint32(buf, math.Float32bits(v))
}

// UnmarshalBinary deserializes a SkipIndex from binary data.
func (si *SkipIndex) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	// Read and verify magic
	magic := make([]byte, 4)
	if _, err := r.Read(magic); err != nil {
		return err
	}
	if string(magic) != skipIndexMagic {
		return fmt.Errorf("invalid magic: %q", string(magic))
	}

	var version uint32
	if err := binary.Read(r, binary.LittleEndian, &version); err != nil {
		return err
	}
	if version != skipIndexVersion {
		return fmt.Errorf("unsupported version: %d", version)
	}

	var blockCount uint32
	if err := binary.Read(r, binary.LittleEndian, &blockCount); err != nil {
		return err
	}
	si.Blocks = make([]BlockStats, blockCount)

	for i := uint32(0); i < blockCount; i++ {
		block := BlockStats{
			ColumnStats: make(map[string]*ColumnStats),
		}
		if err := binary.Read(r, binary.LittleEndian, &block.RowOffset); err != nil {
			return err
		}
		if err := binary.Read(r, binary.LittleEndian, &block.TsMin); err != nil {
			return err
		}
		if err := binary.Read(r, binary.LittleEndian, &block.TsMax); err != nil {
			return err
		}

		var colStatCount uint32
		if err := binary.Read(r, binary.LittleEndian, &colStatCount); err != nil {
			return err
		}
		for j := uint32(0); j < colStatCount; j++ {
			var nameLen uint32
			if err := binary.Read(r, binary.LittleEndian, &nameLen); err != nil {
				return err
			}
			nameBytes := make([]byte, nameLen)
			if _, err := r.Read(nameBytes); err != nil {
				return err
			}
			name := string(nameBytes)

			stats := &ColumnStats{Name: name}

			var hasMin, hasMax byte
			var v64 uint64
			var v32 uint32

			if err := binary.Read(r, binary.LittleEndian, &hasMin); err != nil {
				return err
			}
			if hasMin == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v64); err != nil {
					return err
				}
				v := int64(v64)
				stats.MinInt64 = &v
			}
			if err := binary.Read(r, binary.LittleEndian, &hasMax); err != nil {
				return err
			}
			if hasMax == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v64); err != nil {
					return err
				}
				v := int64(v64)
				stats.MaxInt64 = &v
			}

			if err := binary.Read(r, binary.LittleEndian, &hasMin); err != nil {
				return err
			}
			if hasMin == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v64); err != nil {
					return err
				}
				v := math.Float64frombits(v64)
				stats.MinFloat64 = &v
			}
			if err := binary.Read(r, binary.LittleEndian, &hasMax); err != nil {
				return err
			}
			if hasMax == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v64); err != nil {
					return err
				}
				v := math.Float64frombits(v64)
				stats.MaxFloat64 = &v
			}

			if err := binary.Read(r, binary.LittleEndian, &hasMin); err != nil {
				return err
			}
			if hasMin == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v32); err != nil {
					return err
				}
				v := math.Float32frombits(v32)
				stats.MinFloat32 = &v
			}
			if err := binary.Read(r, binary.LittleEndian, &hasMax); err != nil {
				return err
			}
			if hasMax == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v32); err != nil {
					return err
				}
				v := math.Float32frombits(v32)
				stats.MaxFloat32 = &v
			}

			if err := binary.Read(r, binary.LittleEndian, &hasMin); err != nil {
				return err
			}
			if hasMin == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v32); err != nil {
					return err
				}
				v := int32(v32)
				stats.MinInt32 = &v
			}
			if err := binary.Read(r, binary.LittleEndian, &hasMax); err != nil {
				return err
			}
			if hasMax == 1 {
				if err := binary.Read(r, binary.LittleEndian, &v32); err != nil {
					return err
				}
				v := int32(v32)
				stats.MaxInt32 = &v
			}

			block.ColumnStats[name] = stats
		}

		si.Blocks[i] = block
	}

	return nil
}
