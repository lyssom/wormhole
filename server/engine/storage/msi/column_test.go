package msi

import (
	"bytes"
	"testing"
)

func TestInt32ColumnRoundTrip(t *testing.T) {
	data := make([]int32, 1000)
	for i := range data {
		data[i] = int32(i * 10)
	}

	var buf bytes.Buffer
	col := &ColumnMeta{Type: TypeInt32, Name: "id", Offset: 0, Compression: CompressNone}
	if err := WriteInt32Column(&buf, data, col); err != nil {
		t.Fatalf("WriteInt32Column failed: %v", err)
	}

	readBack, err := ReadInt32Column(&buf, col)
	if err != nil {
		t.Fatalf("ReadInt32Column failed: %v", err)
	}

	if len(readBack) != len(data) {
		t.Errorf("length mismatch: got %d, want %d", len(readBack), len(data))
	}
	for i := range data {
		if readBack[i] != data[i] {
			t.Errorf("data[%d]: got %d, want %d", i, readBack[i], data[i])
		}
	}
}