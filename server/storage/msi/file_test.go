package msi

import (
	"bytes"
	"testing"
)

func TestFullMSIFileRoundTrip(t *testing.T) {
	// 构造一个包含多列的 MSI 文件
	meta := &TableMetadata{
		Columns: []*ColumnMeta{
			{Name: "id", Type: TypeInt64},
			{Name: "ts", Type: TypeInt64},
			{Name: "vec", Type: TypeVectorF32, VectorDim: 128},
		},
		RowCount: 100,
	}

	ids := make([]int64, 100)
	ts := make([]int64, 100)
	vectors := make([][]float32, 100)
	for i := range ids {
		ids[i] = int64(i)
		ts[i] = 1700000000000 + int64(i)*1000
		vectors[i] = make([]float32, 128)
		for j := range vectors[i] {
			vectors[i][j] = float32(i + j)
		}
	}

	var buf bytes.Buffer
	err := WriteMSI(&buf, meta, map[string]interface{}{
		"id":  ids,
		"ts":  ts,
		"vec": vectors,
	})
	if err != nil {
		t.Fatalf("WriteMSI failed: %v", err)
	}

	// 读回
	m, cols, skipData, err := ReadMSI(&buf)
	if err != nil {
		t.Fatalf("ReadMSI failed: %v", err)
	}
	if skipData == nil {
		t.Errorf("skipData should not be nil after WriteMSI")
	}
	if m.RowCount != 100 {
		t.Errorf("RowCount: got %d, want 100", m.RowCount)
	}
	if len(cols) != 3 {
		t.Errorf("Column count: got %d, want 3", len(cols))
	}

	// 验证各列数据
	readIds := cols["id"].([]int64)
	readTs := cols["ts"].([]int64)
	readVecs := cols["vec"].([][]float32)

	if len(readIds) != 100 {
		t.Errorf("ids length: got %d, want 100", len(readIds))
	}
	if len(readTs) != 100 {
		t.Errorf("ts length: got %d, want 100", len(readTs))
	}
	if len(readVecs) != 100 {
		t.Errorf("vecs length: got %d, want 100", len(readVecs))
	}

	// 验证 id 列数据
	for i := range ids {
		if readIds[i] != ids[i] {
			t.Errorf("ids[%d]: got %d, want %d", i, readIds[i], ids[i])
		}
	}

	// 验证 ts 列数据
	for i := range ts {
		if readTs[i] != ts[i] {
			t.Errorf("ts[%d]: got %d, want %d", i, readTs[i], ts[i])
		}
	}

	// 验证 vec 列数据
	for i := range vectors {
		if len(readVecs[i]) != len(vectors[i]) {
			t.Errorf("vecs[%d] length: got %d, want %d", i, len(readVecs[i]), len(vectors[i]))
			continue
		}
		for j := range vectors[i] {
			if readVecs[i][j] != vectors[i][j] {
				t.Errorf("vecs[%d][%d]: got %f, want %f", i, j, readVecs[i][j], vectors[i][j])
			}
		}
	}
}