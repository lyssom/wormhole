package msi

import (
	"encoding/binary"
	"fmt"
	"io"
)

const MagicMSI = 0x4D534931 // "MSI1" in little-endian

type Header struct {
	Magic       uint32
	Version     uint16
	ColumnCount uint16
	RowCount    uint64
	TsMin       uint64 // 毫秒级 Unix 时间戳
	TsMax       uint64
	// ColumnMetas 偏移量在 footer 中
}

// WriteTo serializes the Header to the given writer.
func (h *Header) WriteTo(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, h.Magic); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.ColumnCount); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.RowCount); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.TsMin); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, h.TsMax); err != nil {
		return err
	}
	return nil
}

// ReadHeader deserializes a Header from the given reader.
func ReadHeader(r io.Reader) (*Header, error) {
	h := &Header{}
	if err := binary.Read(r, binary.LittleEndian, &h.Magic); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.ColumnCount); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.RowCount); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.TsMin); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.TsMax); err != nil {
		return nil, err
	}
	if h.Magic != MagicMSI {
		return nil, fmt.Errorf("invalid magic: got 0x%08X, want 0x%08X", h.Magic, MagicMSI)
	}
	if h.Version != 1 {
		return nil, fmt.Errorf("unsupported version: %d", h.Version)
	}
	if h.TsMax < h.TsMin {
		return nil, fmt.Errorf("invalid timestamp range: TsMax (%d) < TsMin (%d)", h.TsMax, h.TsMin)
	}
	return h, nil
}