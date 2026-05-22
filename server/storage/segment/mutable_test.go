package segment

import (
	"os"
	"testing"
)

func TestMutableSegmentAppendAndSeal(t *testing.T) {
	ms := NewMutableSegment("test_table", WithMaxRows(1000))

	// Append 500 rows
	err := ms.Append(map[string]interface{}{
		"id": []int64{1, 2, 3},
	})
	if err != nil {
		t.Fatalf("Append failed: %v", err)
	}

	if ms.RowCount() != 3 {
		t.Errorf("RowCount: got %d, want 3", ms.RowCount())
	}

	// Append to threshold, ShouldSeal should be true
	for i := 0; i < 1000; i++ {
		ms.Append(map[string]interface{}{"id": []int64{int64(i)}})
	}

	if !ms.ShouldSeal() {
		t.Errorf("ShouldSeal: got false, want true")
	}
}

func TestMutableSegmentSeal(t *testing.T) {
	ms := NewMutableSegment("test_table", WithMaxRows(1000))

	// Append some rows
	err := ms.Append(map[string]interface{}{
		"id":   []int64{1, 2, 3},
		"val":  []float64{1.0, 2.0, 3.0},
	})
	if err != nil {
		t.Fatalf("Append failed: %v", err)
	}

	// Seal the segment
	tmpDir := t.TempDir()
	filePath := tmpDir + "/test.msi"
	f, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Create file failed: %v", err)
	}
	defer f.Close()

	si, err := ms.Seal(f)
	if err != nil {
		t.Fatalf("Seal failed: %v", err)
	}

	if si == nil {
		t.Fatalf("Seal returned nil SegmentInstance")
	}

	if si.Path() != filePath {
		t.Errorf("Path: got %s, want %s", si.Path(), filePath)
	}

	if si.RowCount() != 3 {
		t.Errorf("RowCount: got %d, want 3", si.RowCount())
	}
}