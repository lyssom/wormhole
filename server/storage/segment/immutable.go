package segment

import (
	"fmt"
	"os"

	"wormhole/server/storage/index"
	"wormhole/server/storage/msi"
)

// ImmutableSegment provides read access to a sealed MSI file on disk.
type ImmutableSegment struct {
	path      string
	meta      *msi.TableMetadata
	footer    *msi.Footer
	data      map[string]interface{} // column name -> column data
	skipIndex *index.SkipIndex
}

// OpenImmutableSegment opens an MSI file for reading.
func OpenImmutableSegment(path string) (*ImmutableSegment, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("OpenImmutableSegment: %w", err)
	}
	defer f.Close()

	meta, columns, skipIndexData, err := msi.ReadMSI(f)
	if err != nil {
		return nil, fmt.Errorf("OpenImmutableSegment: ReadMSI failed: %w", err)
	}

	// Deserialize skip index from footer data if available
	var skipIdx *index.SkipIndex
	if len(skipIndexData) > 0 {
		skipIdx = &index.SkipIndex{}
		if err := skipIdx.UnmarshalBinary(skipIndexData); err != nil {
			// If skip index deserialization fails, build from columns
			skipIdx = index.BuildFromMSI(meta, columns, 1000)
		}
	} else {
		// Fallback: build skip index from columns
		skipIdx = index.BuildFromMSI(meta, columns, 1000)
	}

	// Build footer from column metas for reference
	footer := &msi.Footer{
		ColumnMetas:     meta.Columns,
		AnnIndexOffsets: make([]int64, len(meta.Columns)),
	}

	return &ImmutableSegment{
		path:      path,
		meta:      meta,
		footer:    footer,
		data:      columns,
		skipIndex: skipIdx,
	}, nil
}

// ReadColumns returns column data for the requested column names.
// Returns a map from column name to the column's data slice.
func (s *ImmutableSegment) ReadColumns(names []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for _, name := range names {
		col, ok := s.data[name]
		if !ok {
			return nil, fmt.Errorf("ReadColumns: column %q not found", name)
		}
		result[name] = col
	}
	return result, nil
}

// Path returns the file path of the segment.
func (s *ImmutableSegment) Path() string {
	return s.path
}

// RowCount returns the number of rows in the segment.
func (s *ImmutableSegment) RowCount() uint64 {
	return s.meta.RowCount
}

// Metadata returns the table metadata.
func (s *ImmutableSegment) Metadata() *msi.TableMetadata {
	return s.meta
}

// SkipIndex returns the segment's skip index for query optimization.
func (s *ImmutableSegment) SkipIndex() *index.SkipIndex {
	return s.skipIndex
}

// CanSkip returns true if this segment can be skipped for the given filter.
// It checks if all blocks in the skip index indicate the filter can't match.
func (s *ImmutableSegment) CanSkip(colName string, minVal, maxVal interface{}) bool {
	if s.skipIndex == nil {
		return false
	}
	return s.skipIndex.CanSkip(colName, minVal, maxVal)
}