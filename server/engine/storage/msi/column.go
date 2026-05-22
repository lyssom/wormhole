package msi

import (
	"encoding/binary"
	"fmt"
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

// WriteInt64Column writes int64 data to the writer using little-endian binary encoding.
func WriteInt64Column(w io.Writer, data []int64, col *ColumnMeta) error {
	col.ValuesCount = uint64(len(data))
	for _, v := range data {
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadInt64Column reads int64 data from the reader using little-endian binary encoding.
func ReadInt64Column(r io.Reader, col *ColumnMeta) ([]int64, error) {
	count := int(col.ValuesCount)
	data := make([]int64, count)
	for i := 0; i < count; i++ {
		if err := binary.Read(r, binary.LittleEndian, &data[i]); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// WriteFloat32Column writes float32 data to the writer using little-endian binary encoding.
func WriteFloat32Column(w io.Writer, data []float32, col *ColumnMeta) error {
	col.ValuesCount = uint64(len(data))
	for _, v := range data {
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadFloat32Column reads float32 data from the reader using little-endian binary encoding.
func ReadFloat32Column(r io.Reader, col *ColumnMeta) ([]float32, error) {
	count := int(col.ValuesCount)
	data := make([]float32, count)
	for i := 0; i < count; i++ {
		if err := binary.Read(r, binary.LittleEndian, &data[i]); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// WriteFloat64Column writes float64 data to the writer using little-endian binary encoding.
func WriteFloat64Column(w io.Writer, data []float64, col *ColumnMeta) error {
	col.ValuesCount = uint64(len(data))
	for _, v := range data {
		if err := binary.Write(w, binary.LittleEndian, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadFloat64Column reads float64 data from the reader using little-endian binary encoding.
func ReadFloat64Column(r io.Reader, col *ColumnMeta) ([]float64, error) {
	count := int(col.ValuesCount)
	data := make([]float64, count)
	for i := 0; i < count; i++ {
		if err := binary.Read(r, binary.LittleEndian, &data[i]); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// ColumnTypeSize returns the byte size of a single value for the given column type.
func ColumnTypeSize(t ColumnType, vectorDim int) int {
	switch t {
	case TypeInt32:
		return 4
	case TypeInt64:
		return 8
	case TypeFloat32:
		return 4
	case TypeFloat64:
		return 8
	case TypeByteArray:
		return 0 // variable size, need offset index
	case TypeVectorF32:
		return 4 * vectorDim
	default:
		return 0
	}
}

func init() {
	// Validate ColumnTypeSize for known types
	if ColumnTypeSize(TypeInt64, 0) != 8 {
		panic(fmt.Sprintf("ColumnTypeSize TypeInt64: got %d, want 8", ColumnTypeSize(TypeInt64, 0)))
	}
	if ColumnTypeSize(TypeVectorF32, 128) != 512 {
		panic(fmt.Sprintf("ColumnTypeSize TypeVectorF32: got %d, want 512", ColumnTypeSize(TypeVectorF32, 128)))
	}
}