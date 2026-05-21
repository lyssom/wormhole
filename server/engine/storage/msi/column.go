package msi

import (
	"encoding/binary"
	"io"
)

// WriteInt32Column writes int32 data to the writer using little-endian binary encoding.
// Plain encoding without compression, designed for SIMD batch reading.
func WriteInt32Column(w io.Writer, data []int32, col *ColumnMeta) error {
	// First write the column metadata
	col.ValuesCount = uint64(len(data))

	// Write all int32 values in little-endian order
	for _, v := range data {
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadInt32Column reads int32 data from the reader using little-endian binary encoding.
func ReadInt32Column(r io.Reader, col *ColumnMeta) ([]int32, error) {
	count := int(col.ValuesCount)
	data := make([]int32, count)
	for i := 0; i < count; i++ {
		if err := binary.Read(r, binary.LittleEndian, &data[i]); err != nil {
			return nil, err
		}
	}
	return data, nil
}