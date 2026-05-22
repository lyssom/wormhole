package wal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// ScanWAL scans the WAL directory and returns all committed transactions.
// The returned map is keyed by txn_id.
// WAL files are named: <txn_id>.wal (e.g., "12345.wal")
func ScanWAL(walDir string) (map[uint64]*WALEntry, error) {
	committed := make(map[uint64]*WALEntry)

	entries, err := os.ReadDir(walDir)
	if err != nil {
		if os.IsNotExist(err) {
			return committed, nil
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(name, ".wal") {
			continue
		}

		txnIDStr := strings.TrimSuffix(name, ".wal")
		txnID, err := strconv.ParseUint(txnIDStr, 10, 64)
		if err != nil {
			continue // skip invalid filenames
		}

		filePath := filepath.Join(walDir, name)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		walEntry, err := ReadWALEntry(file)
		file.Close()
		if err != nil {
			// Log warning and continue, don't abort entire scan
			continue
		}

		if walEntry.Committed {
			committed[txnID] = walEntry
		}
	}

	return committed, nil
}

// ScanWALSorted returns committed transactions sorted by txn_id ascending.
func ScanWALSorted(walDir string) ([]*WALEntry, error) {
	committed, err := ScanWAL(walDir)
	if err != nil {
		return nil, err
	}

	// Sort by txn_id
	ids := make([]uint64, 0, len(committed))
	for id := range committed {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })

	result := make([]*WALEntry, 0, len(ids))
	for _, id := range ids {
		result = append(result, committed[id])
	}

	return result, nil
}

// WALFileName returns the expected WAL filename for a given transaction ID.
func WALFileName(txnID uint64) string {
	return strconv.FormatUint(txnID, 10) + ".wal"
}

// WALFilePath returns the full path to the WAL file for a given transaction.
func WALFilePath(walDir string, txnID uint64) string {
	return filepath.Join(walDir, WALFileName(txnID))
}

// PersistWAL writes a WAL entry to disk atomically.
func PersistWAL(walDir string, entry *WALEntry) error {
	if err := os.MkdirAll(walDir, 0755); err != nil {
		return err
	}

	filePath := WALFilePath(walDir, entry.TxnID)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return entry.WriteTo(file)
}

// CommitWAL marks a transaction as committed by rewriting its WAL entry.
func CommitWAL(walDir string, txnID uint64) error {
	filePath := WALFilePath(walDir, txnID)
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	entry, err := ReadWALEntry(file)
	if err != nil {
		return err
	}
	file.Close()

	entry.Committed = true

	// Rewrite the file with committed flag
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return entry.WriteTo(file)
}

// MarshalJSON implements json.Marshaler for WALEntry.
func (e *WALEntry) MarshalJSON() ([]byte, error) {
	type alias struct {
		TxnID     uint64                   `json:"txn_id"`
		Rows      []map[string]interface{} `json:"rows"`
		Committed bool                     `json:"committed"`
		Timestamp uint64                   `json:"timestamp"`
	}
	return json.Marshal(alias{
		TxnID:     e.TxnID,
		Rows:      e.Rows,
		Committed: e.Committed,
		Timestamp: e.Timestamp,
	})
}