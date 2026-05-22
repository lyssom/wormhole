package wal

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
)

// WALEntry represents a write-ahead log entry for a transaction.
type WALEntry struct {
	TxnID     uint64
	Rows      []map[string]interface{} // row data encoded as JSON
	Committed bool
	Timestamp uint64
}

// Magic number to identify WAL files
var walMagic = []byte{0x57, 0x41, 0x4C, 0x01} // "WAL" + version 1

// WriteTo writes the WAL entry to the provided writer.
// Format: magic(4) + timestamp(8) + txnID(8) + committed(1) + rows(JSON)
func (e *WALEntry) WriteTo(w io.Writer) error {
	// Write magic bytes
	if _, err := w.Write(walMagic); err != nil {
		return err
	}

	// Write timestamp
	if err := binary.Write(w, binary.BigEndian, e.Timestamp); err != nil {
		return err
	}

	// Write txnID
	if err := binary.Write(w, binary.BigEndian, e.TxnID); err != nil {
		return err
	}

	// Write committed flag
	committedByte := byte(0)
	if e.Committed {
		committedByte = 1
	}
	if _, err := w.Write([]byte{committedByte}); err != nil {
		return err
	}

	// Encode rows as JSON
	rowsJSON, err := json.Marshal(e.Rows)
	if err != nil {
		return err
	}

	// Write rows length and data
	if err := binary.Write(w, binary.BigEndian, uint64(len(rowsJSON))); err != nil {
		return err
	}
	if _, err := w.Write(rowsJSON); err != nil {
		return err
	}
	return nil
}

// ReadWALEntry reads a WAL entry from the provided reader.
func ReadWALEntry(r io.Reader) (*WALEntry, error) {
	// Read and verify magic
	magic := make([]byte, 4)
	if _, err := io.ReadFull(r, magic); err != nil {
		return nil, err
	}
	if string(magic) != string(walMagic) {
		return nil, errors.New("invalid WAL magic bytes")
	}

	entry := &WALEntry{}

	// Read timestamp
	if err := binary.Read(r, binary.BigEndian, &entry.Timestamp); err != nil {
		return nil, err
	}

	// Read txnID
	if err := binary.Read(r, binary.BigEndian, &entry.TxnID); err != nil {
		return nil, err
	}

	// Read committed flag
	committedByte := make([]byte, 1)
	if _, err := io.ReadFull(r, committedByte); err != nil {
		return nil, err
	}
	entry.Committed = committedByte[0] == 1

	// Read rows JSON
	var rowsLen uint64
	if err := binary.Read(r, binary.BigEndian, &rowsLen); err != nil {
		return nil, err
	}

	rowsJSON := make([]byte, rowsLen)
	if _, err := io.ReadFull(r, rowsJSON); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rowsJSON, &entry.Rows); err != nil {
		return nil, err
	}

	return entry, nil
}