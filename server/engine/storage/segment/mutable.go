package segment

import (
	"fmt"
	"os"
	"sync"
	"time"

	"wormhole/server/engine/storage/msi"
)

// MutableSegment buffers rows in memory and seals to an ImmutableSegment when threshold is reached.
// It is NOT safe for concurrent use - a single goroutine should handle writes.
type MutableSegment struct {
	tableName string
	columns   map[string][]interface{} // column name -> values (column-oriented storage)
	rowCount  uint64
	maxRows   uint64
	maxAge    time.Duration
	createdAt time.Time
	mu        sync.Mutex // optional mutex for safety, though single-threaded usage is expected
}

// MutableSegmentOption configures a MutableSegment.
type MutableSegmentOption func(*MutableSegment)

// WithMaxRows sets the maximum rows before sealing.
func WithMaxRows(max uint64) MutableSegmentOption {
	return func(ms *MutableSegment) {
		ms.maxRows = max
	}
}

// WithMaxAge sets the maximum age before sealing.
func WithMaxAge(max time.Duration) MutableSegmentOption {
	return func(ms *MutableSegment) {
		ms.maxAge = max
	}
}

// NewMutableSegment creates a new MutableSegment.
func NewMutableSegment(tableName string, opts ...MutableSegmentOption) *MutableSegment {
	ms := &MutableSegment{
		tableName: tableName,
		columns:   make(map[string][]interface{}),
		maxRows:   10000, // default
		createdAt: time.Now(),
	}
	for _, opt := range opts {
		opt(ms)
	}
	return ms
}

// Append adds a batch of rows. The map values must be slices of the same length.
// Supported slice types: []int64, []int32, []float32, []float64, []interface{}
func (ms *MutableSegment) Append(batch map[string]interface{}) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	// Determine batch size from first column
	var batchSize int
	for _, v := range batch {
		batchSize = getSliceLen(v)
		break
	}

	if batchSize == 0 {
		return nil
	}

	// Ensure all columns have the same length
	for colName, v := range batch {
		sz := getSliceLen(v)
		if sz != batchSize {
			return fmt.Errorf("Append: column %q has %d elements, expected %d", colName, sz, batchSize)
		}
	}

	// Append to each column
	for colName, rawValues := range batch {
		values := normalizeToInterfaceSlice(rawValues)
		if values == nil {
			return fmt.Errorf("Append: unsupported type for column %q: %T", colName, rawValues)
		}
		col, ok := ms.columns[colName]
		if !ok {
			// New column - allocate slice
			ms.columns[colName] = values
		} else {
			// Append to existing column
			ms.columns[colName] = append(col, values...)
		}
	}

	ms.rowCount += uint64(batchSize)
	return nil
}

// getSliceLen returns the length of a slice value, regardless of the concrete type.
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
	default:
		return 0
	}
}

// normalizeToInterfaceSlice converts typed slices to []interface{} for uniform storage.
func normalizeToInterfaceSlice(v interface{}) []interface{} {
	switch s := v.(type) {
	case []interface{}:
		return s
	case []int64:
		result := make([]interface{}, len(s))
		for i := range s {
			result[i] = s[i]
		}
		return result
	case []int32:
		result := make([]interface{}, len(s))
		for i := range s {
			result[i] = s[i]
		}
		return result
	case []float32:
		result := make([]interface{}, len(s))
		for i := range s {
			result[i] = s[i]
		}
		return result
	case []float64:
		result := make([]interface{}, len(s))
		for i := range s {
			result[i] = s[i]
		}
		return result
	default:
		return nil
	}
}

// RowCount returns the current number of rows in the segment.
func (ms *MutableSegment) RowCount() uint64 {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	return ms.rowCount
}

// ShouldSeal returns true if the segment should be sealed (row count reached or time window expired).
func (ms *MutableSegment) ShouldSeal() bool {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if ms.maxRows > 0 && ms.rowCount >= ms.maxRows {
		return true
	}
	if ms.maxAge > 0 && time.Since(ms.createdAt) >= ms.maxAge {
		return true
	}
	return false
}

// Seal exports the mutable segment to an ImmutableSegment (MSI file on disk).
// After sealing, the MutableSegment is empty and ready for reuse (or can be discarded).
func (ms *MutableSegment) Seal(w interface{ Write([]byte) (int, error) }) (*SegmentInstance, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if ms.rowCount == 0 {
		return nil, fmt.Errorf("Seal: segment is empty")
	}

	// Build column metas
	colMetas := make([]*msi.ColumnMeta, 0, len(ms.columns))
	for name, values := range ms.columns {
		// Determine type from first value
		if len(values) == 0 {
			continue
		}
		colType := inferColumnType(values[0])
		colMetas = append(colMetas, &msi.ColumnMeta{
			Name:        name,
			Type:        colType,
			ValuesCount: uint64(len(values)),
		})
	}

	meta := &msi.TableMetadata{
		Columns:  colMetas,
		RowCount: ms.rowCount,
	}

	// Write MSI file
	// w is io.Writer (from Seal signature), but we need to convert if needed
	writer, ok := w.(interface {
		Write([]byte) (int, error)
	})
	if !ok {
		return nil, fmt.Errorf("Seal: w does not implement Write([]byte) (int, error)")
	}

	// Use WriteMSI - need to convert ms.columns from map[string][]interface{} to map[string]interface{}
	columnsData := make(map[string]interface{})
	for k, v := range ms.columns {
		columnsData[k] = v
	}

	// WriteMSI expects specific types, not []interface{} directly
	// We need to convert []interface{} to proper typed slices
	typedColumns := make(map[string]interface{})
	for colName, values := range ms.columns {
		if len(values) == 0 {
			continue
		}
		// Convert to appropriate typed slice based on type inference
		switch values[0].(type) {
		case int64:
			// Convert []interface{} to []int64
			converted := make([]int64, len(values))
			for i, v := range values {
				if iv, ok := v.(int64); ok {
					converted[i] = iv
				} else {
					return nil, fmt.Errorf("Seal: column %s has mixed types", colName)
				}
			}
			typedColumns[colName] = converted
		case float64:
			converted := make([]float64, len(values))
			for i, v := range values {
				if iv, ok := v.(float64); ok {
					converted[i] = iv
				} else {
					return nil, fmt.Errorf("Seal: column %s has mixed types", colName)
				}
			}
			typedColumns[colName] = converted
		case float32:
			converted := make([]float32, len(values))
			for i, v := range values {
				if iv, ok := v.(float32); ok {
					converted[i] = iv
				} else {
					return nil, fmt.Errorf("Seal: column %s has mixed types", colName)
				}
			}
			typedColumns[colName] = converted
		case int32:
			converted := make([]int32, len(values))
			for i, v := range values {
				if iv, ok := v.(int32); ok {
					converted[i] = iv
				} else {
					return nil, fmt.Errorf("Seal: column %s has mixed types", colName)
				}
			}
			typedColumns[colName] = converted
		default:
			return nil, fmt.Errorf("Seal: unsupported column type for %s: %T", colName, values[0])
		}
	}

	if err := msi.WriteMSI(writer, meta, typedColumns); err != nil {
		return nil, fmt.Errorf("Seal: WriteMSI failed: %w", err)
	}

	// Get the path from the writer (if it has a Name method)
	path := ""
	if f, ok := w.(interface{ Name() string }); ok {
		path = f.Name()
	}

	// Reset the segment after sealing
	ms.columns = make(map[string][]interface{})
	ms.rowCount = 0

	return &SegmentInstance{
		path:     path,
		rowCount: meta.RowCount,
	}, nil
}

// SealToFile seals the segment and writes to a file at the given path.
func (ms *MutableSegment) SealToFile(filePath string) (*SegmentInstance, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("SealToFile: Create failed: %w", err)
	}
	defer f.Close()

	si, err := ms.Seal(f)
	if err != nil {
		return nil, err
	}
	si.path = filePath
	return si, nil
}

// inferColumnType infers the MSI ColumnType from a value.
func inferColumnType(value interface{}) msi.ColumnType {
	switch value.(type) {
	case int64:
		return msi.TypeInt64
	case int32:
		return msi.TypeInt32
	case float64:
		return msi.TypeFloat64
	case float32:
		return msi.TypeFloat32
	default:
		return msi.TypeByteArray
	}
}

func getFirstKey(m map[string]interface{}) string {
	for k := range m {
		return k
	}
	return ""
}

// --- SegmentInstance (ImmutableSegment on disk) ---

// SegmentInstance represents an immutable segment on disk.
type SegmentInstance struct {
	path     string
	rowCount uint64
}

// Path returns the file path of the segment.
func (si *SegmentInstance) Path() string {
	return si.path
}

// RowCount returns the number of rows in the segment.
func (si *SegmentInstance) RowCount() uint64 {
	return si.rowCount
}