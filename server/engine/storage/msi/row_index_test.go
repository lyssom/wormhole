package msi

import (
	"bytes"
	"testing"
)

func TestRowIndexRoundTrip(t *testing.T) {
	blocks := []*RowIndexBlock{
		{RowOffset: 0, TsMin: 1700000000000, TsMax: 1700001000000},
		{RowOffset: 4096, TsMin: 1700001000001, TsMax: 1700002000000},
	}

	var buf bytes.Buffer
	err := WriteRowIndex(&buf, blocks)
	if err != nil {
		t.Fatalf("WriteRowIndex failed: %v", err)
	}

	readBack, err := ReadRowIndex(&buf)
	if err != nil {
		t.Fatalf("ReadRowIndex failed: %v", err)
	}

	if len(readBack) != len(blocks) {
		t.Errorf("block count: got %d, want %d", len(readBack), len(blocks))
	}
	for i, b := range blocks {
		if readBack[i].RowOffset != b.RowOffset || readBack[i].TsMin != b.TsMin {
			t.Errorf("block %d mismatch: got %+v, want %+v", i, readBack[i], b)
		}
	}
}