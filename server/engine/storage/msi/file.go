package msi

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

// TableMetadata contains table-level metadata (columns and row count).
type TableMetadata struct {
	Columns  []*ColumnMeta
	RowCount uint64
}

// WriteMSI writes a complete MSI file to w.
// Layout: [Header 32B] → [Column 1 data] → [Column 2 data] → ... → [RowIndex] → [Footer]
func WriteMSI(w io.Writer, meta *TableMetadata, columns map[string]interface{}) error {
	// We need to know column offsets before writing the Header.
	// Strategy: collect all column data into buffers first, then write everything in order.

	colBuffers := make([]*bytes.Buffer, len(meta.Columns))

	// Collect column data into buffers to determine offsets
	var currentOffset uint64 = 32 // After header

	for i, colMeta := range meta.Columns {
		buf := &bytes.Buffer{}
		colBuffers[i] = buf

		switch col := columns[colMeta.Name].(type) {
		case []int64:
			colMeta.ValuesCount = uint64(len(col))
			if err := WriteInt64Column(buf, col, colMeta); err != nil {
				return fmt.Errorf("writing column %s: %w", colMeta.Name, err)
			}
		case []int32:
			colMeta.ValuesCount = uint64(len(col))
			if err := WriteInt32Column(buf, col, colMeta); err != nil {
				return fmt.Errorf("writing column %s: %w", colMeta.Name, err)
			}
		case []float32:
			colMeta.ValuesCount = uint64(len(col))
			if err := WriteFloat32Column(buf, col, colMeta); err != nil {
				return fmt.Errorf("writing column %s: %w", colMeta.Name, err)
			}
		case []float64:
			colMeta.ValuesCount = uint64(len(col))
			if err := WriteFloat64Column(buf, col, colMeta); err != nil {
				return fmt.Errorf("writing column %s: %w", colMeta.Name, err)
			}
		case [][]float32:
			colMeta.ValuesCount = uint64(len(col))
			if err := WriteVectorColumn(buf, col, colMeta); err != nil {
				return fmt.Errorf("writing column %s: %w", colMeta.Name, err)
			}
		default:
			return fmt.Errorf("unsupported column type for %s: %T", colMeta.Name, col)
		}

		colMeta.Offset = currentOffset
		currentOffset += uint64(buf.Len())
	}

	// Calculate TsMin and TsMax from the "ts" column if present
	tsMin, tsMax := computeTsRange(columns)

	// Write Header (32 bytes)
	header := &Header{
		Magic:       MagicMSI,
		Version:     1,
		ColumnCount: uint16(len(meta.Columns)),
		RowCount:    meta.RowCount,
		TsMin:       tsMin,
		TsMax:       tsMax,
	}
	if err := header.WriteTo(w); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	// Write each column's buffered data
	for i, buf := range colBuffers {
		if _, err := w.Write(buf.Bytes()); err != nil {
			return fmt.Errorf("writing column %d data: %w", i, err)
		}
	}

	// Write RowIndex (1 block for simplicity)
	blocks := []*RowIndexBlock{
		{RowOffset: 0, TsMin: tsMin, TsMax: tsMax},
	}
	if err := WriteRowIndex(w, blocks); err != nil {
		return fmt.Errorf("writing row index: %w", err)
	}

	// Write Footer with ColumnMetas (includes offsets)
	footer := &Footer{
		ColumnMetas:     meta.Columns,
		AnnIndexOffsets: make([]int64, len(meta.Columns)), // no ANN index for now
	}
	if err := footer.WriteTo(w); err != nil {
		return fmt.Errorf("writing footer: %w", err)
	}

	return nil
}

// ReadMSI reads a complete MSI file.
// Layout: [Header 32B] → [Column 1 data] → [Column 2 data] → ... → [RowIndex] → [Footer]
func ReadMSI(r io.Reader) (*TableMetadata, map[string]interface{}, error) {
	// Read entire file into memory to enable seeking to column offsets
	allData, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, fmt.Errorf("reading all data: %w", err)
	}

	totalLen := int64(len(allData))

	// Read Header
	reader := bytes.NewReader(allData)
	header, err := ReadHeader(reader)
	if err != nil {
		return nil, nil, fmt.Errorf("reading header: %w", err)
	}

	// Footer is at the end. Since we don't have footer offset in header,
	// we search backwards from the end of the file.
	// Maximum reasonable footer size: for typical files with 3 columns and reasonable names,
	// footer is < 500 bytes. Search up to 100KB to be safe.

	// Search from 100KB from end (very conservative)
	maxSearchBack := int64(100 * 1024)
	if maxSearchBack > totalLen-32 {
		maxSearchBack = totalLen - 32
	}

	var footer *Footer
	// Search backwards from near the end, checking for valid footer
	// Start from totalLen - 100 (approximately footer size for small files)
	searchStart := totalLen - 100
	if searchStart < 32 {
		searchStart = 32
	}
	for seekPos := searchStart; seekPos >= totalLen-maxSearchBack && seekPos >= 32; seekPos-- {
		reader.Seek(seekPos, io.SeekStart)
		testFooter, err := ReadFooter(reader)
		if err == nil && len(testFooter.ColumnMetas) == int(header.ColumnCount) {
			footer = testFooter
			break
		}
	}

	if footer == nil {
		return nil, nil, fmt.Errorf("could not find valid footer with %d columns", header.ColumnCount)
	}

	// Build result columns map
	result := make(map[string]interface{})

	// Read each column at its recorded offset
	for _, colMeta := range footer.ColumnMetas {
		switch colMeta.Type {
		case TypeInt64:
			col := &ColumnMeta{
				ValuesCount: colMeta.ValuesCount,
				Type:        colMeta.Type,
				Name:        colMeta.Name,
				Offset:      colMeta.Offset,
				VectorDim:   colMeta.VectorDim,
			}
			colData, err := readColumnAt(reader, int64(colMeta.Offset), int64(colMeta.ValuesCount)*8)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			data, err := ReadInt64Column(bytes.NewReader(colData), col)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			result[colMeta.Name] = data

		case TypeInt32:
			col := &ColumnMeta{
				ValuesCount: colMeta.ValuesCount,
				Type:        colMeta.Type,
				Name:        colMeta.Name,
				Offset:      colMeta.Offset,
				VectorDim:   colMeta.VectorDim,
			}
			colData, err := readColumnAt(reader, int64(colMeta.Offset), int64(colMeta.ValuesCount)*4)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			data, err := ReadInt32Column(bytes.NewReader(colData), col)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			result[colMeta.Name] = data

		case TypeFloat32:
			col := &ColumnMeta{
				ValuesCount: colMeta.ValuesCount,
				Type:        colMeta.Type,
				Name:        colMeta.Name,
				Offset:      colMeta.Offset,
				VectorDim:   colMeta.VectorDim,
			}
			colData, err := readColumnAt(reader, int64(colMeta.Offset), int64(colMeta.ValuesCount)*4)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			data, err := ReadFloat32Column(bytes.NewReader(colData), col)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			result[colMeta.Name] = data

		case TypeFloat64:
			col := &ColumnMeta{
				ValuesCount: colMeta.ValuesCount,
				Type:        colMeta.Type,
				Name:        colMeta.Name,
				Offset:      colMeta.Offset,
				VectorDim:   colMeta.VectorDim,
			}
			colData, err := readColumnAt(reader, int64(colMeta.Offset), int64(colMeta.ValuesCount)*8)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			data, err := ReadFloat64Column(bytes.NewReader(colData), col)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			result[colMeta.Name] = data

		case TypeVectorF32:
			col := &ColumnMeta{
				ValuesCount: colMeta.ValuesCount,
				Type:        colMeta.Type,
				Name:        colMeta.Name,
				Offset:      colMeta.Offset,
				VectorDim:   colMeta.VectorDim,
			}
			colData, err := readColumnAt(reader, int64(colMeta.Offset), int64(colMeta.ValuesCount)*int64(colMeta.VectorDim)*4)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			data, err := ReadVectorColumn(bytes.NewReader(colData), col)
			if err != nil {
				return nil, nil, fmt.Errorf("reading column %s: %w", colMeta.Name, err)
			}
			result[colMeta.Name] = data

		default:
			return nil, nil, fmt.Errorf("unsupported column type: %v", colMeta.Type)
		}
	}

	meta := &TableMetadata{
		Columns:  footer.ColumnMetas,
		RowCount: header.RowCount,
	}

	return meta, result, nil
}

// readColumnAt reads size bytes from r at the given offset.
func readColumnAt(r io.ReadSeeker, offset int64, size int64) ([]byte, error) {
	if _, err := r.Seek(offset, io.SeekStart); err != nil {
		return nil, err
	}
	data := make([]byte, size)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, err
	}
	return data, nil
}

// computeTsRange computes TsMin and TsMax from the "ts" column if present.
func computeTsRange(columns map[string]interface{}) (uint64, uint64) {
	ts, ok := columns["ts"].([]int64)
	if !ok || len(ts) == 0 {
		return 0, 0
	}
	tsMin := uint64(ts[0])
	tsMax := uint64(ts[len(ts)-1])
	// Also handle if data is not sorted - find actual min/max
	for _, t := range ts {
		if uint64(t) < tsMin {
			tsMin = uint64(t)
		}
		if uint64(t) > tsMax {
			tsMax = uint64(t)
		}
	}
	return tsMin, tsMax
}

// float32bits returns the IEEE 754 binary representation of f.
func float32bits(f float32) uint32 {
	return math.Float32bits(f)
}