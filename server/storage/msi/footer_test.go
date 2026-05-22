package msi

import (
	"bytes"
	"testing"
)

func TestFooterRoundTrip(t *testing.T) {
	footer := &Footer{
		ColumnMetas: []*ColumnMeta{
			{Name: "id", Type: TypeInt64, Offset: 128, VectorDim: 0},
			{Name: "vec", Type: TypeVectorF32, Offset: 128 + 8*1000, VectorDim: 128},
		},
		AnnIndexOffsets: []int64{1024, 2048},
	}

	var buf bytes.Buffer
	err := footer.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo failed: %v", err)
	}

	footer2, err := ReadFooter(&buf)
	if err != nil {
		t.Fatalf("ReadFooter failed: %v", err)
	}

	if len(footer2.ColumnMetas) != 2 {
		t.Errorf("ColumnMeta count: got %d, want 2", len(footer2.ColumnMetas))
	}
	if footer2.ColumnMetas[0].Name != "id" || footer2.ColumnMetas[0].Type != TypeInt64 {
		t.Errorf("ColumnMeta[0] mismatch: got %+v", footer2.ColumnMetas[0])
	}
	if footer2.ColumnMetas[1].Name != "vec" || footer2.ColumnMetas[1].Type != TypeVectorF32 || footer2.ColumnMetas[1].VectorDim != 128 {
		t.Errorf("ColumnMeta[1] mismatch: got %+v", footer2.ColumnMetas[1])
	}
	if len(footer2.AnnIndexOffsets) != 2 || footer2.AnnIndexOffsets[0] != 1024 || footer2.AnnIndexOffsets[1] != 2048 {
		t.Errorf("AnnIndexOffsets mismatch: got %+v", footer2.AnnIndexOffsets)
	}
}