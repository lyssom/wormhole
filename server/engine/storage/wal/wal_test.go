package wal

import (
	"bytes"
	"testing"
)

func TestWALEntryRoundTrip(t *testing.T) {
	entry := &WALEntry{
		TxnID:     42,
		Rows:      []map[string]interface{}{{"id": int64(1), "name": "test"}},
		Committed: false,
		Timestamp: 1700000000000,
	}

	var buf bytes.Buffer
	err := entry.WriteTo(&buf)
	if err != nil {
		t.Fatalf("WriteTo failed: %v", err)
	}

	entry2, err := ReadWALEntry(&buf)
	if err != nil {
		t.Fatalf("ReadWALEntry failed: %v", err)
	}

	if entry2.TxnID != 42 || entry2.Committed != false {
		t.Errorf("Entry mismatch: got %+v", entry2)
	}

	// Mark committed and rewrite
	entry2.Committed = true
	var buf2 bytes.Buffer
	entry2.WriteTo(&buf2)

	entry3, _ := ReadWALEntry(&buf2)
	if entry3.Committed != true {
		t.Errorf("Committed flag not set: got %v", entry3.Committed)
	}
}