package scheduler

import (
	"os"
	"testing"
	"time"

	"wormhole/server/storage"
	"wormhole/server/storage/msi"
	"wormhole/server/storage/segment"
)

func TestMergeSchedulerStartStop(t *testing.T) {
	tmp := t.TempDir()
	engine := storage.NewStorageEngine(tmp, nil)
	sched := NewMergeScheduler(engine, nil)

	sched.Start()
	time.Sleep(100 * time.Millisecond)

	if err := sched.Stop(); err != nil {
		t.Fatalf("Stop failed: %v", err)
	}

	// Start again should work
	sched.Start()
	time.Sleep(100 * time.Millisecond)
	if err := sched.Stop(); err != nil {
		t.Fatalf("Stop failed: %v", err)
	}
}

func TestMergeSchedulerDoubleStart(t *testing.T) {
	tmp := t.TempDir()
	engine := storage.NewStorageEngine(tmp, nil)
	sched := NewMergeScheduler(engine, nil)

	sched.Start()
	sched.Start() // Should be no-op
	time.Sleep(100 * time.Millisecond)

	if err := sched.Stop(); err != nil {
		t.Fatalf("Stop failed: %v", err)
	}
}

func TestMergeSchedulerDoubleStop(t *testing.T) {
	tmp := t.TempDir()
	engine := storage.NewStorageEngine(tmp, nil)
	sched := NewMergeScheduler(engine, nil)

	sched.Start()
	time.Sleep(100 * time.Millisecond)

	if err := sched.Stop(); err != nil {
		t.Fatalf("First Stop failed: %v", err)
	}
	if err := sched.Stop(); err != nil {
		t.Fatalf("Second Stop failed: %v", err)
	}
}

func TestMergeTaskSortMerge(t *testing.T) {
	tmp := t.TempDir()

	// Create 2 immutable segments with overlapping timestamps
	seg1 := createTestSegment(t, tmp, "seg1.msi", []int64{1, 3, 5}, []int64{100, 300, 500})
	seg2 := createTestSegment(t, tmp, "seg2.msi", []int64{2, 4, 6}, []int64{200, 400, 600})

	task := NewMergeTask("test_table", "test_db", "default", []*segment.ImmutableSegment{seg1, seg2}, tmp)
	if err := task.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Verify output file exists
	if task.OutputPath() == "" {
		t.Errorf("OutputPath is empty")
	}
	if _, err := os.Stat(task.OutputPath()); os.IsNotExist(err) {
		t.Errorf("Output file does not exist: %s", task.OutputPath())
	}

	// Verify input paths
	inputs := task.InputPaths()
	if len(inputs) != 2 {
		t.Errorf("InputPaths: got %d, want 2", len(inputs))
	}
}

func TestMergeTaskSortMergeMultipleSegments(t *testing.T) {
	tmp := t.TempDir()

	// Create 3 immutable segments
	seg1 := createTestSegment(t, tmp, "seg1.msi", []int64{1, 4}, []int64{100, 400})
	seg2 := createTestSegment(t, tmp, "seg2.msi", []int64{2, 5}, []int64{200, 500})
	seg3 := createTestSegment(t, tmp, "seg3.msi", []int64{3, 6}, []int64{300, 600})

	task := NewMergeTask("test_table", "test_db", "default", []*segment.ImmutableSegment{seg1, seg2, seg3}, tmp)
	if err := task.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Read merged segment and verify
	mergedSeg, err := segment.OpenImmutableSegment(task.OutputPath())
	if err != nil {
		t.Fatalf("OpenImmutableSegment failed: %v", err)
	}

	if mergedSeg.RowCount() != 6 {
		t.Errorf("RowCount: got %d, want 6", mergedSeg.RowCount())
	}

	cols, err := mergedSeg.ReadColumns([]string{"id", "ts"})
	if err != nil {
		t.Fatalf("ReadColumns failed: %v", err)
	}

	idCol := cols["id"].([]int64)
	tsCol := cols["ts"].([]int64)

	// Verify sorted by ts
	for i := 1; i < len(tsCol); i++ {
		if tsCol[i] < tsCol[i-1] {
			t.Errorf("Not sorted by ts: ts[%d]=%d < ts[%d]=%d", i, tsCol[i], i-1, tsCol[i-1])
		}
	}

	// Verify IDs are in correct order
	expectedIDs := []int64{1, 2, 3, 4, 5, 6}
	for i, expected := range expectedIDs {
		if idCol[i] != expected {
			t.Errorf("id[%d]: got %d, want %d", i, idCol[i], expected)
		}
	}
}

func TestMergeTaskInputPaths(t *testing.T) {
	tmp := t.TempDir()

	seg1 := createTestSegment(t, tmp, "seg1.msi", []int64{1}, []int64{100})
	seg2 := createTestSegment(t, tmp, "seg2.msi", []int64{2}, []int64{200})

	task := NewMergeTask("test_table", "test_db", "default", []*segment.ImmutableSegment{seg1, seg2}, tmp)

	paths := task.InputPaths()
	if len(paths) != 2 {
		t.Errorf("len(InputPaths): got %d, want 2", len(paths))
	}
}

// createTestSegment creates a test MSI file with id and ts columns.
func createTestSegment(t *testing.T, dir string, filename string, ids, ts []int64) *segment.ImmutableSegment {
	meta := &msi.TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "ts", Type: msi.TypeInt64},
		},
		RowCount: uint64(len(ids)),
	}

	columns := map[string]interface{}{
		"id": ids,
		"ts": ts,
	}

	path := dir + "/" + filename
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	defer f.Close()

	if err := msi.WriteMSI(f, meta, columns); err != nil {
		t.Fatalf("WriteMSI failed: %v", err)
	}

	seg, err := segment.OpenImmutableSegment(path)
	if err != nil {
		t.Fatalf("OpenImmutableSegment failed: %v", err)
	}

	return seg
}
