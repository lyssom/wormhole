package segment

import (
	"bytes"
	"os"
	"testing"

	"wormhole/server/storage/msi"
)

func TestImmutableSegmentRead(t *testing.T) {
	// 构造 MSI 文件
	meta := &msi.TableMetadata{
		Columns: []*msi.ColumnMeta{
			{Name: "id", Type: msi.TypeInt64},
			{Name: "vec", Type: msi.TypeVectorF32, VectorDim: 128},
		},
		RowCount: 50,
	}
	ids := make([]int64, 50)
	vectors := make([][]float32, 50)
	for i := range ids {
		ids[i] = int64(i)
		vectors[i] = make([]float32, 128)
		vectors[i][0] = float32(i)
	}

	var buf bytes.Buffer
	msi.WriteMSI(&buf, meta, map[string]interface{}{
		"id":  ids,
		"vec": vectors,
	})

	tmp, _ := os.CreateTemp("", "test-*.msi")
	tmp.Write(buf.Bytes())
	tmp.Close()
	defer os.Remove(tmp.Name())

	seg, err := OpenImmutableSegment(tmp.Name())
	if err != nil {
		t.Fatalf("OpenImmutableSegment failed: %v", err)
	}

	cols, err := seg.ReadColumns([]string{"id", "vec"})
	if err != nil {
		t.Fatalf("ReadColumns failed: %v", err)
	}

	if len(cols) != 2 {
		t.Errorf("column count: got %d, want 2", len(cols))
	}

	// Verify id column data
	ids, ok := cols["id"].([]int64)
	if !ok {
		t.Fatalf("id column type: got %T, want []int64", cols["id"])
	}
	if len(ids) != 50 {
		t.Errorf("id column length: got %d, want 50", len(ids))
	}
	for i, id := range ids {
		if id != int64(i) {
			t.Errorf("id[%d]: got %d, want %d", i, id, int64(i))
		}
	}

	// Verify vec column data
	vecs, ok := cols["vec"].([][]float32)
	if !ok {
		t.Fatalf("vec column type: got %T, want [][]float32", cols["vec"])
	}
	if len(vecs) != 50 {
		t.Errorf("vec column length: got %d, want 50", len(vecs))
	}
	for i, v := range vecs {
		if len(v) != 128 {
			t.Errorf("vec[%d] dim: got %d, want 128", i, len(v))
		}
		if v[0] != float32(i) {
			t.Errorf("vec[%d][0]: got %f, want %f", i, v[0], float32(i))
		}
	}

	// Verify RowCount
	if seg.RowCount() != 50 {
		t.Errorf("RowCount: got %d, want 50", seg.RowCount())
	}
}