package msi

import (
	"bytes"
	"testing"
)

func TestHeaderRoundTrip(t *testing.T) {
	h := &Header{
		Magic:       MagicMSI,
		Version:     1,
		ColumnCount: 3,
		RowCount:    1000,
		TsMin:       1700000000000,
		TsMax:       1700003600000,
	}

	var buf bytes.Buffer
	err := h.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo failed: %v", err)
	}

	h2, err := ReadHeader(&buf)
	if err != nil {
		t.Fatalf("ReadHeader failed: %v", err)
	}

	if h2.Magic != h.Magic || h2.Version != h.Version || h2.ColumnCount != h.ColumnCount || h2.RowCount != h.RowCount {
		t.Errorf("Header mismatch: got %+v, want %+v", h2, h)
	}
}