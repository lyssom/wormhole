package segment

import (
	"bytes"
	"os"
	"testing"

	"wormhole/server/engine/storage/msi"
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
}