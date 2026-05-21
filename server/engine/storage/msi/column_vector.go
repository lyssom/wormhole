package msi

import (
	"encoding/binary"
	"io"
)

// WriteVectorColumn writes count vectors of dimension dim as flat float32.
// Layout: [vec0_dim0, vec0_dim1, ..., vec1_dim0, ...]
// Designed for SIMD dot product / euclidean distance.
func WriteVectorColumn(w io.Writer, vectors [][]float32, meta *ColumnMeta) error {
	meta.ValuesCount = uint64(len(vectors))
	dim := meta.VectorDim
	buf := make([]float32, len(vectors)*dim)
	for i, v := range vectors {
		copy(buf[i*dim:(i+1)*dim], v)
	}
	return binary.Write(w, binary.LittleEndian, buf)
}

// ReadVectorColumn reads vectors written by WriteVectorColumn.
func ReadVectorColumn(r io.Reader, meta *ColumnMeta) ([][]float32, error) {
	dim := meta.VectorDim
	count := int(meta.ValuesCount)
	buf := make([]float32, count*dim)
	for i := 0; i < count*dim; i++ {
		if err := binary.Read(r, binary.LittleEndian, &buf[i]); err != nil {
			return nil, err
		}
	}
	vectors := make([][]float32, count)
	for i := 0; i < count; i++ {
		vectors[i] = buf[i*dim : (i+1)*dim]
	}
	return vectors, nil
}