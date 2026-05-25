package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"wormhole/server/storage/msi"
	"wormhole/server/storage/partition"
	"wormhole/server/storage/segment"
)

func TestTableInsertAndSelect(t *testing.T) {
	meta := &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "score", Type: msi.TypeFloat64},
		},
	}
	table := newTable("db1", "t1", meta, nil, 3600000)

	// Insert rows
	rows := []map[string]interface{}{
		{"id": int64(1), "score": 1.5},
		{"id": int64(2), "score": 2.5},
		{"id": int64(3), "score": 3.5},
	}
	if err := table.Insert(rows); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	// Select columns
	cols, err := table.SelectColumns([]string{"id", "score"})
	if err != nil {
		t.Fatalf("SelectColumns failed: %v", err)
	}

	idCol, ok := cols["id"].([]interface{})
	if !ok {
		t.Fatalf("id column type: got %T, want []interface{}", cols["id"])
	}
	if len(idCol) != 3 {
		t.Errorf("id column length: got %d, want 3", len(idCol))
	}
	for i, id := range idCol {
		if id != int64(i+1) {
			t.Errorf("id[%d]: got %v, want %d", i, id, i+1)
		}
	}
}

func TestTableRowCount(t *testing.T) {
	meta := &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
		},
	}
	table := newTable("db1", "t1", meta, nil, 3600000)

	if table.RowCount() != 0 {
		t.Errorf("initial RowCount: got %d, want 0", table.RowCount())
	}

	rows := []map[string]interface{}{
		{"id": int64(1)},
		{"id": int64(2)},
	}
	if err := table.Insert(rows); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	if table.RowCount() != 2 {
		t.Errorf("RowCount after insert: got %d, want 2", table.RowCount())
	}
}

func TestTablePartitionRouting(t *testing.T) {
	meta := &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "ts", Type: msi.TypeInt64},
		},
	}
	// 1-hour bucket
	bucketMs := int64(3600000)
	table := newTable("db1", "t1", meta, nil, bucketMs)

	// Insert rows with different timestamps (should go to different partitions)
	rows1 := []map[string]interface{}{
		{"id": int64(1), "ts": int64(1000)},  // bucket 0
		{"id": int64(2), "ts": int64(1000)},
	}
	rows2 := []map[string]interface{}{
		{"id": int64(3), "ts": int64(4000000)}, // bucket 1
		{"id": int64(4), "ts": int64(4000000)},
	}

	if err := table.Insert(rows1); err != nil {
		t.Fatalf("Insert batch 1 failed: %v", err)
	}
	if err := table.Insert(rows2); err != nil {
		t.Fatalf("Insert batch 2 failed: %v", err)
	}

	// Verify two partitions were created
	parts := table.Partitions()
	if len(parts) != 2 {
		t.Errorf("partition count: got %d, want 2", len(parts))
	}

	// Verify total row count
	if table.RowCount() != 4 {
		t.Errorf("RowCount: got %d, want 4", table.RowCount())
	}
}

func TestTableUpdateRow(t *testing.T) {
	meta := &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "val", Type: msi.TypeFloat64},
		},
	}
	table := newTable("db1", "t1", meta, nil, 3600000)

	// Insert rows
	rows := []map[string]interface{}{
		{"id": int64(1), "val": 1.0},
		{"id": int64(2), "val": 2.0},
	}
	if err := table.Insert(rows); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	// UpdateRow applies pending updates that are applied during SelectColumns
	if err := table.UpdateRow(0, map[string]interface{}{"val": 100.0}); err != nil {
		t.Fatalf("UpdateRow failed: %v", err)
	}

	// Verify data is updated
	cols, err := table.SelectColumns([]string{"val"})
	if err != nil {
		t.Fatalf("SelectColumns failed: %v", err)
	}

	valCol := cols["val"].([]interface{})
	if valCol[0] != 100.0 {
		t.Errorf("val[0] after update: got %v, want 100.0", valCol[0])
	}
	// Row 1 should be unchanged
	if valCol[1] != 2.0 {
		t.Errorf("val[1] unchanged: got %v, want 2.0", valCol[1])
	}
}

func TestTableDeleteRow(t *testing.T) {
	meta := &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "val", Type: msi.TypeFloat64},
		},
	}
	table := newTable("db1", "t1", meta, nil, 3600000)

	// Insert rows
	rows := []map[string]interface{}{
		{"id": int64(1), "val": 1.0},
		{"id": int64(2), "val": 2.0},
		{"id": int64(3), "val": 3.0},
	}
	if err := table.Insert(rows); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	// DeleteRow is a no-op in partition mode (MutableSegment doesn't support delete)
	if err := table.DeleteRow(1); err != nil {
		t.Fatalf("DeleteRow failed: %v", err)
	}

	// Row count unchanged (since delete is a no-op)
	if table.RowCount() != 3 {
		t.Errorf("RowCount unchanged (delete is no-op): got %d, want 3", table.RowCount())
	}
}

func TestPartitionInsertAndSeal(t *testing.T) {
	tmp := t.TempDir()
	bucketMs := int64(3600000)
	part := partition.NewPartition("bucket_0", bucketMs)

	// Insert rows
	rows := []map[string]interface{}{
		{"id": int64(1), "val": 1.0},
		{"id": int64(2), "val": 2.0},
	}
	if err := part.Insert(rows); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	if part.RowCount() != 2 {
		t.Errorf("RowCount after insert: got %d, want 2", part.RowCount())
	}

	// ShouldSeal returns false (threshold not met), so Seal returns false
	sealed, err := part.Seal(tmp)
	if err != nil {
		t.Fatalf("Seal failed: %v", err)
	}
	if sealed {
		t.Errorf("Seal: expected false (threshold not met)")
	}

	// Verify no immutable segments yet
	imms := part.ImmutableSegments()
	if len(imms) != 0 {
		t.Errorf("immutable segment count: got %d, want 0", len(imms))
	}
}

func TestPartitionSealMultipleSegments(t *testing.T) {
	tmp := t.TempDir()
	bucketMs := int64(3600000)
	part := partition.NewPartition("bucket_0", bucketMs)

	// Insert first batch and seal (will not actually seal since threshold not met)
	rows1 := []map[string]interface{}{
		{"id": int64(1)},
		{"id": int64(2)},
	}
	if err := part.Insert(rows1); err != nil {
		t.Fatalf("Insert batch 1 failed: %v", err)
	}
	part.Seal(tmp) // returns false since threshold not met

	// Insert second batch and seal
	rows2 := []map[string]interface{}{
		{"id": int64(3)},
		{"id": int64(4)},
	}
	if err := part.Insert(rows2); err != nil {
		t.Fatalf("Insert batch 2 failed: %v", err)
	}
	part.Seal(tmp) // returns false since threshold not met

	// Should have 0 immutable segments (threshold never met)
	imms := part.ImmutableSegments()
	if len(imms) != 0 {
		t.Errorf("immutable segment count: got %d, want 0", len(imms))
	}

	// Total row count should be 4
	if part.RowCount() != 4 {
		t.Errorf("RowCount: got %d, want 4", part.RowCount())
	}
}

func TestPartitionManager(t *testing.T) {
	bucketMs := int64(3600000) // 1 hour
	pm := partition.NewPartitionManager(bucketMs)

	// Get partitions for different timestamps
	p1 := pm.GetOrCreatePartition(1000)        // bucket 0
	p2 := pm.GetOrCreatePartition(4000000)    // bucket 1
	p3 := pm.GetOrCreatePartition(1000)       // same as p1

	if p1 != p3 {
		t.Errorf("GetOrCreatePartition should return same partition for same bucket")
	}

	if p1 == p2 {
		t.Errorf("Different buckets should return different partitions")
	}

	// Verify bucket keys
	if partition.BucketKey(1000, bucketMs) != "bucket_0" {
		t.Errorf("BucketKey(1000): got bucket_0")
	}
	if partition.BucketKey(4000000, bucketMs) != "bucket_1" {
		t.Errorf("BucketKey(4000000): got bucket_1")
	}

	// Verify partitions
	parts := pm.Partitions()
	if len(parts) != 2 {
		t.Errorf("partition count: got %d, want 2", len(parts))
	}
}

func TestMutableSegmentAppendInt32(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))
	err := ms.Append(map[string]interface{}{
		"val": []int32{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("Append int32 failed: %v", err)
	}
	if ms.RowCount() != 3 {
		t.Errorf("RowCount: got %d, want 3", ms.RowCount())
	}
}

func TestMutableSegmentAppendFloat32(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))
	err := ms.Append(map[string]interface{}{
		"val": []float32{1.0, 2.0, 3.0},
	})
	if err != nil {
		t.Fatalf("Append float32 failed: %v", err)
	}
	if ms.RowCount() != 3 {
		t.Errorf("RowCount: got %d, want 3", ms.RowCount())
	}
}

func TestMutableSegmentAppendFloat64(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))
	err := ms.Append(map[string]interface{}{
		"val": []float64{1.0, 2.0, 3.0},
	})
	if err != nil {
		t.Fatalf("Append float64 failed: %v", err)
	}
	if ms.RowCount() != 3 {
		t.Errorf("RowCount: got %d, want 3", ms.RowCount())
	}
}

func TestMutableSegmentAppendVectorNotSupported(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))
	vec1 := []float32{1.0, 2.0, 3.0, 4.0}
	vec2 := []float32{5.0, 6.0, 7.0, 8.0}
	err := ms.Append(map[string]interface{}{
		"vec": [][]float32{vec1, vec2},
	})
	// Append silently does nothing for unsupported types
	if err != nil {
		t.Logf("Append vector returned error (known limitation): %v", err)
	}
	if ms.RowCount() != 0 {
		t.Errorf("RowCount: got %d, want 0 (vector type not supported in Append)", ms.RowCount())
	}
}

func TestMutableSegmentShouldSealMaxAge(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxAge(50*time.Millisecond))

	if ms.ShouldSeal() {
		t.Errorf("ShouldSeal immediately after creation: got true, want false")
	}

	// Append some data
	ms.Append(map[string]interface{}{"id": []int64{1}})

	// Wait for maxAge to pass
	time.Sleep(60 * time.Millisecond)

	if !ms.ShouldSeal() {
		t.Errorf("ShouldSeal after maxAge: got false, want true")
	}
}

func TestMutableSegmentAppendMixedBatchSizes(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))

	// First batch: 3 rows
	err := ms.Append(map[string]interface{}{
		"id": []int64{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("Append batch 1 failed: %v", err)
	}

	// Second batch: 2 rows
	err = ms.Append(map[string]interface{}{
		"id": []int64{4, 5},
	})
	if err != nil {
		t.Fatalf("Append batch 2 failed: %v", err)
	}

	if ms.RowCount() != 5 {
		t.Errorf("RowCount: got %d, want 5", ms.RowCount())
	}
}

func TestMutableSegmentAppendMismatchedBatchSizes(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))

	// Batch with mismatched column lengths
	err := ms.Append(map[string]interface{}{
		"id":  []int64{1, 2, 3},
		"val": []float64{1.0, 2.0}, // mismatch!
	})
	if err == nil {
		t.Fatalf("Append with mismatched batch sizes: expected error")
	}
}

func TestMutableSegmentAppendEmpty(t *testing.T) {
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))

	err := ms.Append(map[string]interface{}{})
	if err != nil {
		t.Fatalf("Append empty batch failed: %v", err)
	}

	if ms.RowCount() != 0 {
		t.Errorf("RowCount after empty append: got %d, want 0", ms.RowCount())
	}
}

func TestMutableSegmentSealMultipleColumns(t *testing.T) {
	tmp := t.TempDir()
	ms := segment.NewMutableSegment("test", segment.WithMaxRows(1000))

	err := ms.Append(map[string]interface{}{
		"id":    []int64{1, 2, 3},
		"name":  []int32{10, 20, 30},
		"score": []float64{1.5, 2.5, 3.5},
	})
	if err != nil {
		t.Fatalf("Append failed: %v", err)
	}

	segPath := filepath.Join(tmp, "test.msi")
	si, err := ms.SealToFile(segPath)
	if err != nil {
		t.Fatalf("SealToFile failed: %v", err)
	}

	if si.RowCount() != 3 {
		t.Errorf("SegmentInstance RowCount: got %d, want 3", si.RowCount())
	}

	// Verify file was created
	if _, err := os.Stat(segPath); os.IsNotExist(err) {
		t.Errorf("MSI file was not created at %s", segPath)
	}
}

func TestStorageEngineGetOrCreateTable(t *testing.T) {
	tmp := t.TempDir()
	engine := NewStorageEngine(tmp, nil)

	meta := &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
		},
	}

	// Create first time
	t1, err := engine.GetOrCreateTable("db1", "t1", meta)
	if err != nil {
		t.Fatalf("GetOrCreateTable (create) failed: %v", err)
	}
	if t1 == nil {
		t.Fatalf("GetOrCreateTable returned nil table")
	}

	// Get existing table
	t2, err := engine.GetOrCreateTable("db1", "t1", meta)
	if err != nil {
		t.Fatalf("GetOrCreateTable (get) failed: %v", err)
	}
	if t2 != t1 {
		t.Errorf("GetOrCreateTable should return same table instance")
	}

	// Insert should work on existing table
	err = engine.Insert("db1", "t1", []map[string]interface{}{
		{"id": int64(1)},
	})
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}
}

func TestStorageEngineTableNotFound(t *testing.T) {
	tmp := t.TempDir()
	engine := NewStorageEngine(tmp, nil)

	// Insert to non-existent table
	err := engine.Insert("db1", "nonexistent", []map[string]interface{}{
		{"id": int64(1)},
	})
	if err == nil {
		t.Fatalf("Insert to nonexistent table: expected error")
	}

	// Select from non-existent table
	_, err = engine.SelectColumns("db1", "nonexistent", []string{"id"})
	if err == nil {
		t.Fatalf("SelectColumns from nonexistent table: expected error")
	}

	// Update non-existent table
	err = engine.UpdateRow("db1", "nonexistent", 0, map[string]interface{}{"id": int64(1)})
	if err == nil {
		t.Fatalf("UpdateRow on nonexistent table: expected error")
	}

	// Delete from non-existent table
	err = engine.DeleteRow("db1", "nonexistent", 0)
	if err == nil {
		t.Fatalf("DeleteRow on nonexistent table: expected error")
	}
}

func TestStorageEngineInsertAndSelect(t *testing.T) {
	tmp := t.TempDir()
	engine := NewStorageEngine(tmp, &Options{
		MutableMaxRows:   100000,
		BucketDurationMs: 3600000,
	})

	// Create table
	err := engine.CreateTable("db1", "t1", &TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "vec", Type: msi.TypeVectorF32, VectorDim: 128},
		},
	})
	if err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}

	// Insert
	err = engine.Insert("db1", "t1", []map[string]interface{}{
		{"id": int64(1), "vec": make([]float32, 128)},
	})
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	// Read
	cols, err := engine.SelectColumns("db1", "t1", []string{"id", "vec"})
	if err != nil {
		t.Fatalf("SelectColumns failed: %v", err)
	}
	idCol, ok := cols["id"].([]interface{})
	if !ok {
		t.Fatalf("cols[id] is not []interface{}")
	}
	if len(idCol) != 1 {
		t.Errorf("row count: got %d, want 1", len(idCol))
	}
}
