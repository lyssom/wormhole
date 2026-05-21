package msi

import (
	"bytes"
	"testing"
)

func TestVectorColumnRoundTrip(t *testing.T) {
	dim := 128
	count := 100
	vectors := make([][]float32, count)
	for i := range vectors {
		vectors[i] = make([]float32, dim)
		for j := range vectors[i] {
			vectors[i][j] = float32(i*dim + j)
		}
	}

	var buf bytes.Buffer
	col := &ColumnMeta{Type: TypeVectorF32, VectorDim: dim}
	err := WriteVectorColumn(&buf, vectors, col)
	if err != nil {
		t.Fatalf("WriteVectorColumn failed: %v", err)
	}

	readBack, err := ReadVectorColumn(&buf, col)
	if err != nil {
		t.Fatalf("ReadVectorColumn failed: %v", err)
	}

	if len(readBack) != count {
		t.Errorf("count mismatch: got %d, want %d", len(readBack), count)
	}
	for i := range vectors {
		if len(readBack[i]) != dim {
			t.Errorf("dim mismatch at %d: got %d, want %d", i, len(readBack[i]), dim)
		}
		for j := range vectors[i] {
			if readBack[i][j] != vectors[i][j] {
				t.Errorf("vectors[%d][%d]: got %f, want %f", i, j, readBack[i][j], vectors[i][j])
			}
		}
	}
}