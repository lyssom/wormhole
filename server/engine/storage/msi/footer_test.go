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

	if len(footer2.ColumnMetas) != 2 || footer2.ColumnMetas[0].Name != "id" {
		t.Errorf("Footer mismatch: got %+v", footer2)
	}
}