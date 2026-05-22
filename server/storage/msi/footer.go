package msi

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Footer stores column metadata and ANN index offsets.
// It comes after the row index in the MSI file format.
type Footer struct {
	ColumnMetas     []*ColumnMeta
	AnnIndexOffsets []int64 // 每个列的 ANN 索引偏移量
}

// WriteTo serializes the Footer to the given writer.
func (f *Footer) WriteTo(w io.Writer) error {
	// Write number of column metas
	if err := binary.Write(w, binary.LittleEndian, uint32(len(f.ColumnMetas))); err != nil {
		return err
	}

	// Write each column meta
	for _, col := range f.ColumnMetas {
		// Write name length and name bytes
		nameBytes := []byte(col.Name)
		if err := binary.Write(w, binary.LittleEndian, uint32(len(nameBytes))); err != nil {
			return err
		}
		if _, err := w.Write(nameBytes); err != nil {
			return err
		}
		// Write column meta fields
		if err := binary.Write(w, binary.LittleEndian, col.Type); err != nil {
			return err
		}
		if err := binary.Write(w, binary.LittleEndian, col.Offset); err != nil {
			return err
		}
		if err := binary.Write(w, binary.LittleEndian, int32(col.VectorDim)); err != nil {
			return err
		}
		if err := binary.Write(w, binary.LittleEndian, col.ValuesCount); err != nil {
			return err
		}
	}

	// Write number of ANN index offsets
	if err := binary.Write(w, binary.LittleEndian, uint32(len(f.AnnIndexOffsets))); err != nil {
		return err
	}

	// Write each ANN index offset
	for _, offset := range f.AnnIndexOffsets {
		if err := binary.Write(w, binary.LittleEndian, offset); err != nil {
			return err
		}
	}

	return nil
}

// ReadFooter deserializes a Footer from the given reader.
func ReadFooter(r io.Reader) (*Footer, error) {
	f := &Footer{}

	// Read number of column metas
	var colMetaCount uint32
	if err := binary.Read(r, binary.LittleEndian, &colMetaCount); err != nil {
		return nil, err
	}

	// Sanity check: limit column count to prevent allocation attacks
	if colMetaCount > 1000 {
		return nil, fmt.Errorf("column meta count too large: %d", colMetaCount)
	}

	// Read each column meta
	f.ColumnMetas = make([]*ColumnMeta, colMetaCount)
	for i := uint32(0); i < colMetaCount; i++ {
		col := &ColumnMeta{}

		// Read name length
		var nameLen uint32
		if err := binary.Read(r, binary.LittleEndian, &nameLen); err != nil {
			return nil, err
		}

		// Limit name length to prevent allocation attacks
		if nameLen > 1024 {
			return nil, fmt.Errorf("column name too long: %d", nameLen)
		}

		// Read name bytes
		nameBytes := make([]byte, nameLen)
		if _, err := io.ReadFull(r, nameBytes); err != nil {
			return nil, err
		}
		col.Name = string(nameBytes)

		// Read column meta fields
		if err := binary.Read(r, binary.LittleEndian, &col.Type); err != nil {
			return nil, err
		}
		if err := binary.Read(r, binary.LittleEndian, &col.Offset); err != nil {
			return nil, err
		}
		var vectorDim int32
		if err := binary.Read(r, binary.LittleEndian, &vectorDim); err != nil {
			return nil, err
		}
		col.VectorDim = int(vectorDim)
		if err := binary.Read(r, binary.LittleEndian, &col.ValuesCount); err != nil {
			return nil, err
		}

		f.ColumnMetas[i] = col
	}

	// Read number of ANN index offsets
	var annIndexCount uint32
	if err := binary.Read(r, binary.LittleEndian, &annIndexCount); err != nil {
		return nil, err
	}

	// Validate: ann index count must match column count
	if int(annIndexCount) != len(f.ColumnMetas) && annIndexCount > 0 {
		return nil, errors.New("ANN index count mismatch with column count")
	}

	// Read each ANN index offset
	f.AnnIndexOffsets = make([]int64, annIndexCount)
	for i := uint32(0); i < annIndexCount; i++ {
		if err := binary.Read(r, binary.LittleEndian, &f.AnnIndexOffsets[i]); err != nil {
			return nil, err
		}
	}

	return f, nil
}