package index

import (
	"testing"
)

// mockColumnMeta implements ColumnMetaInfo for testing
type mockColumnMeta struct {
	name string
	typ  int
}

func (m *mockColumnMeta) GetName() string { return m.name }
func (m *mockColumnMeta) GetType() int { return m.typ }

// mockTableMetadata implements TableMetadata for testing
type mockTableMetadata struct {
	columns []*mockColumnMeta
	rowCount uint64
}

func (m *mockTableMetadata) GetRowCount() uint64 { return m.rowCount }
func (m *mockTableMetadata) GetColumns() []ColumnMetaInfo {
	result := make([]ColumnMetaInfo, len(m.columns))
	for i, c := range m.columns {
		result[i] = c
	}
	return result
}

func TestBuildFromMSI(t *testing.T) {
	meta := &mockTableMetadata{
		columns: []*mockColumnMeta{
			{name: "id", typ: 1},  // TypeInt64 = 1
			{name: "score", typ: 3}, // TypeFloat64 = 3
		},
		rowCount: 5,
	}
	columns := map[string]interface{}{
		"id":    []int64{1, 2, 3, 4, 5},
		"score": []float64{1.5, 2.5, 3.5, 4.5, 5.5},
	}

	skipIdx := BuildFromMSI(meta, columns, 2) // 2 rows per block

	if skipIdx.BlockCount() != 3 {
		t.Errorf("BlockCount: got %d, want 3", skipIdx.BlockCount())
	}

	// Block 0: rows 0-1 (id: 1-2, score: 1.5-2.5)
	if !skipIdx.Blocks[0].CanSkipInt64("id", 5, 10) {
		t.Errorf("Block 0 should be skippable for id > 5")
	}
	if skipIdx.Blocks[0].CanSkipInt64("id", 1, 3) {
		t.Errorf("Block 0 should NOT be skippable for id 1-3")
	}
}

func TestCanSkipInt64(t *testing.T) {
	stats := &BlockStats{
		RowOffset:   0,
		TsMin:       0,
		TsMax:       100,
		ColumnStats: make(map[string]*ColumnStats),
	}
	stats.ColumnStats["age"] = &ColumnStats{
		Name:     "age",
		MinInt64: int64Ptr(20),
		MaxInt64: int64Ptr(30),
	}

	// Query range [25, 35] overlaps with block range [20, 30] -> don't skip
	if stats.CanSkipInt64("age", 25, 35) {
		t.Errorf("Should NOT skip: query [25,35] overlaps block [20,30]")
	}

	// Query range [35, 45] doesn't overlap with block [20, 30] -> skip
	if !stats.CanSkipInt64("age", 35, 45) {
		t.Errorf("Should skip: query [35,45] doesn't overlap block [20,30]")
	}

	// Query range [10, 15] doesn't overlap with block [20, 30] -> skip
	if !stats.CanSkipInt64("age", 10, 15) {
		t.Errorf("Should skip: query [10,15] doesn't overlap block [20,30]")
	}

	// Query range [15, 25] overlaps with block [20, 30] -> don't skip
	if stats.CanSkipInt64("age", 15, 25) {
		t.Errorf("Should NOT skip: query [15,25] overlaps block [20,30]")
	}
}

func TestCanSkipFloat64(t *testing.T) {
	stats := &BlockStats{
		RowOffset:   0,
		TsMin:       0,
		TsMax:       100,
		ColumnStats: make(map[string]*ColumnStats),
	}
	stats.ColumnStats["score"] = &ColumnStats{
		Name:         "score",
		MinFloat64:  float64Ptr(1.5),
		MaxFloat64:  float64Ptr(3.5),
	}

	// Query range [2.0, 4.0] overlaps with block [1.5, 3.5] -> don't skip
	if stats.CanSkipFloat64("score", 2.0, 4.0) {
		t.Errorf("Should NOT skip: query [2.0,4.0] overlaps block [1.5,3.5]")
	}

	// Query range [4.0, 5.0] doesn't overlap with block [1.5, 3.5] -> skip
	if !stats.CanSkipFloat64("score", 4.0, 5.0) {
		t.Errorf("Should skip: query [4.0,5.0] doesn't overlap block [1.5,3.5]")
	}
}

func TestCanSkipNoStats(t *testing.T) {
	stats := &BlockStats{
		RowOffset:   0,
		TsMin:       0,
		TsMax:       100,
		ColumnStats: make(map[string]*ColumnStats),
	}
	// No stats for "unknown" column -> can't skip
	if stats.CanSkipInt64("unknown", 10, 20) {
		t.Errorf("Should NOT skip if no stats available")
	}
}

func TestSkipIndexCanSkip(t *testing.T) {
	idx := &SkipIndex{
		Blocks: []BlockStats{
			{
				RowOffset:   0,
				TsMin:       0,
				TsMax:       50,
				ColumnStats: map[string]*ColumnStats{
					"id": {Name: "id", MinInt64: int64Ptr(1), MaxInt64: int64Ptr(10)},
				},
			},
			{
				RowOffset:   10,
				TsMin:       50,
				TsMax:       100,
				ColumnStats: map[string]*ColumnStats{
					"id": {Name: "id", MinInt64: int64Ptr(11), MaxInt64: int64Ptr(20)},
				},
			},
		},
	}

	// Query [15, 25]: block 0 can be skipped (max 10 < 15), block 1 can't (11 <= 25)
	if idx.CanSkip("id", int64(15), int64(25)) {
		t.Errorf("Should NOT skip: block 1 overlaps")
	}

	// Query [25, 35]: block 0 can be skipped, block 1 can be skipped (min 11 > 25)
	if !idx.CanSkip("id", int64(25), int64(35)) {
		t.Errorf("Should skip: both blocks outside query range")
	}
}

func int64Ptr(v int64) *int64 {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}
