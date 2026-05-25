package scheduler

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"wormhole/server/storage/msi"
	"wormhole/server/storage/segment"
)

// rowData holds a single row's data for sort-merge operations.
type rowData struct {
	ts   uint64
	cols map[string]interface{}
}

// MergeTask executes a single merge operation on a set of immutable segments.
type MergeTask struct {
	tableName  string
	db         string
	partition  string
	segments   []*segment.ImmutableSegment
	rootDir    string
	outputPath string
}

// NewMergeTask creates a new MergeTask.
func NewMergeTask(tableName, db, partition string, segments []*segment.ImmutableSegment, rootDir string) *MergeTask {
	return &MergeTask{
		tableName: tableName,
		db:        db,
		partition: partition,
		segments:  segments,
		rootDir:   rootDir,
	}
}

// Execute performs the sort-merge and writes the new MSI.
func (t *MergeTask) Execute() error {
	// Read all column data from input segments
	mergedColumns, rowCount, err := t.sortMerge()
	if err != nil {
		return fmt.Errorf("sortMerge failed: %w", err)
	}

	// Build metadata
	colMetas := buildColumnMetas(mergedColumns)
	meta := &msi.TableMetadata{
		Columns:  colMetas,
		RowCount: rowCount,
	}

	// Determine output path
	tablePath := filepath.Join(t.rootDir, t.db, t.tableName)
	outputPath := filepath.Join(tablePath, t.partition, "immutable", fmt.Sprintf("%d.msi", len(t.segments)))
	t.outputPath = outputPath

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	// Write MSI file
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Create failed: %w", err)
	}
	defer f.Close()

	if err := msi.WriteMSI(f, meta, mergedColumns); err != nil {
		return fmt.Errorf("WriteMSI failed: %w", err)
	}

	return nil
}

// sortMerge reads all segments, sort-merges by timestamp, returns merged columns and row count.
func (t *MergeTask) sortMerge() (map[string]interface{}, uint64, error) {
	if len(t.segments) == 0 {
		return nil, 0, fmt.Errorf("no segments to merge")
	}

	// Collect all column names and types from first segment's metadata
	meta := t.segments[0].Metadata()
	var colNames []string
	colTypes := make(map[string]msi.ColumnType)
	for _, col := range meta.Columns {
		colNames = append(colNames, col.Name)
		colTypes[col.Name] = col.Type
	}

	// Read all column data from all segments
	var allRows []rowData

	for _, seg := range t.segments {
		cols, err := seg.ReadColumns(colNames)
		if err != nil {
			return nil, 0, fmt.Errorf("ReadColumns failed: %w", err)
		}

		segMeta := seg.Metadata()
		// Get ts column for sorting
		tsCol, ok := cols["ts"].([]int64)
		if !ok {
			// If no ts column, use TsMin from metadata
			tsCol = make([]int64, segMeta.RowCount)
		}

		// Build row map
		for i := uint64(0); i < segMeta.RowCount; i++ {
			row := make(map[string]interface{})
			for name, colData := range cols {
				row[name] = getValueAt(colData, int(i))
			}
			ts := tsCol[int(i)]
			allRows = append(allRows, rowData{ts: uint64(ts), cols: row})
		}
	}

	// Sort by timestamp
	sort.Slice(allRows, func(i, j int) bool {
		return allRows[i].ts < allRows[j].ts
	})

	// Rebuild column-oriented format with proper types
	result := make(map[string]interface{})
	for _, name := range colNames {
		result[name] = buildTypedColumn(colTypes[name], allRows, name)
	}

	return result, uint64(len(allRows)), nil
}

// buildTypedColumn builds a properly typed slice from row data.
func buildTypedColumn(colType msi.ColumnType, rows []rowData, colName string) interface{} {
	switch colType {
	case msi.TypeInt64:
		vals := make([]int64, len(rows))
		for i, row := range rows {
			vals[i] = row.cols[colName].(int64)
		}
		return vals
	case msi.TypeInt32:
		vals := make([]int32, len(rows))
		for i, row := range rows {
			vals[i] = row.cols[colName].(int32)
		}
		return vals
	case msi.TypeFloat64:
		vals := make([]float64, len(rows))
		for i, row := range rows {
			vals[i] = row.cols[colName].(float64)
		}
		return vals
	case msi.TypeFloat32:
		vals := make([]float32, len(rows))
		for i, row := range rows {
			vals[i] = row.cols[colName].(float32)
		}
		return vals
	case msi.TypeVectorF32:
		vals := make([][]float32, len(rows))
		for i, row := range rows {
			vals[i] = row.cols[colName].([]float32)
		}
		return vals
	default:
		// Fallback to []interface{}
		vals := make([]interface{}, len(rows))
		for i, row := range rows {
			vals[i] = row.cols[colName]
		}
		return vals
	}
}

// getValueAt returns the value at index i from a typed slice.
func getValueAt(data interface{}, i int) interface{} {
	switch s := data.(type) {
	case []int64:
		return s[i]
	case []int32:
		return s[i]
	case []float32:
		return s[i]
	case []float64:
		return s[i]
	case [][]float32:
		return s[i]
	case []interface{}:
		return s[i]
	default:
		return nil
	}
}

// buildColumnMetas builds column metadata from merged column data.
func buildColumnMetas(columns map[string]interface{}) []*msi.ColumnMeta {
	var metas []*msi.ColumnMeta
	for name, data := range columns {
		colType := getMSIType(data)
		if colType == msi.TypeByteArray && getSliceLen(data) == 0 {
			continue
		}
		metas = append(metas, &msi.ColumnMeta{
			Name: name,
			Type: colType,
		})
	}
	return metas
}

// getMSIType returns the MSI ColumnType for a typed slice.
func getMSIType(data interface{}) msi.ColumnType {
	switch data.(type) {
	case []int64:
		return msi.TypeInt64
	case []int32:
		return msi.TypeInt32
	case []float64:
		return msi.TypeFloat64
	case []float32:
		return msi.TypeFloat32
	case [][]float32:
		return msi.TypeVectorF32
	default:
		return msi.TypeByteArray
	}
}

// getSliceLen returns the length of a slice value.
func getSliceLen(v interface{}) int {
	switch s := v.(type) {
	case []interface{}:
		return len(s)
	case []int64:
		return len(s)
	case []int32:
		return len(s)
	case []float32:
		return len(s)
	case []float64:
		return len(s)
	case [][]float32:
		return len(s)
	default:
		return 0
	}
}

// inferMSIType infers the MSI ColumnType from a value.
func inferMSIType(value interface{}) msi.ColumnType {
	switch value.(type) {
	case int64:
		return msi.TypeInt64
	case int32:
		return msi.TypeInt32
	case float64:
		return msi.TypeFloat64
	case float32:
		return msi.TypeFloat32
	case []float32:
		return msi.TypeVectorF32
	default:
		return msi.TypeByteArray
	}
}

// OutputPath returns the path of the merged MSI file.
func (t *MergeTask) OutputPath() string {
	return t.outputPath
}

// InputPaths returns the paths of the input segments.
func (t *MergeTask) InputPaths() []string {
	var paths []string
	for _, seg := range t.segments {
		paths = append(paths, seg.Path())
	}
	return paths
}
