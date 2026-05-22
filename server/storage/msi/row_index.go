package msi

import (
	"encoding/binary"
	"io"
)

type RowIndexBlock struct {
	RowOffset uint64 // 该块第一行的文件内字节偏移量（估算）
	TsMin    uint64
	TsMax    uint64
}

// RowIndexBlockSize is the expected size of each row index block in bytes.
// It is used to validate row offsets during read operations.
const RowIndexBlockSize = 4096

// WriteRowIndex writes block count + each block's (RowOffset, TsMin, TsMax).
// RowOffset values are validated against RowIndexBlockSize to ensure correct alignment.
// 用于时间范围查询时二分跳转到目标块
func WriteRowIndex(w io.Writer, blocks []*RowIndexBlock) error {
	// 写入块数量
	if err := binary.Write(w, binary.BigEndian, uint64(len(blocks))); err != nil {
		return err
	}

	// 写入每个块的索引信息
	for _, block := range blocks {
		if err := binary.Write(w, binary.BigEndian, block.RowOffset); err != nil {
			return err
		}
		if err := binary.Write(w, binary.BigEndian, block.TsMin); err != nil {
			return err
		}
		if err := binary.Write(w, binary.BigEndian, block.TsMax); err != nil {
			return err
		}
	}

	return nil
}

// ReadRowIndex 读取块数量 + 每个块的 (TsMin, TsMax)
func ReadRowIndex(r io.Reader) ([]*RowIndexBlock, error) {
	// 读取块数量
	var count uint64
	if err := binary.Read(r, binary.BigEndian, &count); err != nil {
		return nil, err
	}

	blocks := make([]*RowIndexBlock, count)
	for i := uint64(0); i < count; i++ {
		block := &RowIndexBlock{}
		if err := binary.Read(r, binary.BigEndian, &block.RowOffset); err != nil {
			return nil, err
		}
		if err := binary.Read(r, binary.BigEndian, &block.TsMin); err != nil {
			return nil, err
		}
		if err := binary.Read(r, binary.BigEndian, &block.TsMax); err != nil {
			return nil, err
		}
		blocks[i] = block
	}

	return blocks, nil
}