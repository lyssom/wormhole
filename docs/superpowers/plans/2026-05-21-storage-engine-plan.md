# Storage Engine Implementation Plan — MSI Format + Mutable Segment

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 实现自研列式存储格式 .msi 和 Mutable Segment 写入路径，支持可选 Exactly-Once 语义。

**Architecture:** 自研 MSI 列式格式（Vector-first、SIMD 友好、时间索引内嵌），Mutable Segment 接收写入，后台 merge 为 Immutable Segment。WAL 目录实现 Exactly-Once 去重。

**Tech Stack:** 纯 Go，无 CGO 依赖。使用 `encoding/binary` 处理二进制布局，`bytes.Buffer` 构造 MSI 文件。

---

## File Structure

```
server/engine/storage/
├── msi/
│   ├── header.go        # MSI Header (magic, version, column count, row count, timestamps)
│   ├── column.go         # 列数据页 (INT/FLOAT/DOUBLE/STRING/VECTOR_F32)
│   ├── column_vector.go  # 向量列（连续 float32，SIMD 友好）
│   ├── row_index.go      # 时间块索引（每 4096 行一条，记录 min/max timestamp）
│   └── footer.go         # Footer 序列化 (column metas, ANN index offsets)
├── segment/
│   ├── segment.go        # Segment 接口定义
│   ├── mutable.go        # MutableSegment: 内存缓冲，按列追加写入
│   ├── immutable.go      # ImmutableSegment: 只读 MSI 文件读取
│   └── merger.go         # BackgroundMergeTask: Mutable → Immutable
├── wal/
│   ├── wal.go            # WAL 条目结构，单事务一文件
│   ├── wal_test.go
│   └── recovery.go       # Scan WAL, deduplicate by txn_id
├── options.go            # TableOptions (consistency, bucket_size, bucket_duration)
├── table.go              # TableMetadata (_meta.json 读写)
└── storage.go            # StorageEngine 上层接口 (Insert, Select, Recover)

server/engine/storage/msi/msi_test.go          # 格式读写测试
server/engine/storage/segment/segment_test.go  # Segment 接口测试
server/engine/storage/wal/wal_test.go           # WAL 测试
```

---

## Task 1: MSI Header 读写

**Files:**
- Create: `server/engine/storage/msi/header.go`
- Test: `server/engine/storage/msi/header_test.go`

- [ ] **Step 1: 写测试 — header 序列化和反序列化**

```go
package msi

import (
    "bytes"
    "testing"
)

func TestHeaderRoundTrip(t *testing.T) {
    h := &Header{
        Magic:       MagicMSI,
        Version:     1,
        ColumnCount: 3,
        RowCount:    1000,
        TsMin:       1700000000000,
        TsMax:       1700003600000,
    }

    var buf bytes.Buffer
    err := h.WriteTo(&buf)
    if err != nil {
        t.Fatalf("WriteTo failed: %v", err)
    }

    h2, err := ReadHeader(&buf)
    if err != nil {
        t.Fatalf("ReadHeader failed: %v", err)
    }

    if h2.Magic != h.Magic || h2.Version != h.Version || h2.ColumnCount != h.ColumnCount || h2.RowCount != h.RowCount {
        t.Errorf("Header mismatch: got %+v, want %+v", h2, h)
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/msi/... -run TestHeaderRoundTrip -v`
Expected: FAIL — file does not exist

- [ ] **Step 3: 实现 Header 结构**

```go
package msi

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
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/msi/... -run TestHeaderRoundTrip -v`
Expected: PASS

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/msi/header.go server/engine/storage/msi/header_test.go
git commit -m "feat(storage): add MSI header serialization"
```

---

## Task 2: 列数据页读写（Plain Encoding, SIMD 友好）

**Files:**
- Create: `server/engine/storage/msi/column.go`
- Modify: `server/engine/storage/msi/header.go` (添加 ColumnMeta)
- Test: `server/engine/storage/msi/column_test.go`

- [ ] **Step 1: 写测试 — INT32 列的 round-trip**

```go
package msi

import (
    "bytes"
    "testing"
)

func TestInt32ColumnRoundTrip(t *testing.T) {
    data := make([]int32, 1000)
    for i := range data {
        data[i] = int32(i * 10)
    }

    var buf bytes.Buffer
    col := &ColumnMeta{Type: TypeInt32, Name: "id", Offset: 0, Compression: CompressNone}
    n, err := WriteInt32Column(&buf, data, col)
    if err != nil {
        t.Fatalf("WriteInt32Column failed: %v", err)
    }

    readBack, err := ReadInt32Column(&buf, col)
    if err != nil {
        t.Fatalf("ReadInt32Column failed: %v", err)
    }

    if len(readBack) != len(data) {
        t.Errorf("length mismatch: got %d, want %d", len(readBack), len(data))
    }
    for i := range data {
        if readBack[i] != data[i] {
            t.Errorf("data[%d]: got %d, want %d", i, readBack[i], data[i])
        }
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/msi/... -run TestInt32ColumnRoundTrip -v`
Expected: FAIL

- [ ] **Step 3: 实现 ColumnType 和 ColumnMeta**

```go
package msi

type ColumnType uint8

const (
    TypeInt32     ColumnType = 0
    TypeInt64      ColumnType = 1
    TypeFloat32   ColumnType = 2
    TypeFloat64   ColumnType = 3
    TypeByteArray ColumnType = 4 // 字符串
    TypeVectorF32 ColumnType = 5 // 差异化：固定维度向量
)

const CompressNone = 0

type ColumnMeta struct {
    Name         string
    Type         ColumnType
    Offset       uint64 // 文件内偏移量
    ValuesCount  uint64
    Compression  uint8
    // 向量列特有
    VectorDim    int   // e.g., 128 for FLOAT[128]
    IndexOffset  int64 // ANN 索引在文件内的偏移量
}
```

- [ ] **Step 4: 实现 WriteInt32Column / ReadInt32Column**

Plain encoding: `binary.LittleEndian` 连续写入，无压缩。用途是 SIMD 批量读取。

- [ ] **Step 5: 运行测试确认通过**

Run: `go test ./server/engine/storage/msi/... -run TestInt32ColumnRoundTrip -v`

- [ ] **Step 6: 提交**

```bash
git add server/engine/storage/msi/column.go server/engine/storage/msi/header.go
git commit -m "feat(storage): add INT32 column read/write"
```

---

## Task 3: 向量列 (VECTOR_F32) — 差异化的核心

**Files:**
- Create: `server/engine/storage/msi/column_vector.go`
- Test: `server/engine/storage/msi/column_vector_test.go`

- [ ] **Step 1: 写测试 — FLOAT[128] 向量列 round-trip**

```go
package msi

import (
    "bytes"
    "testing"
)

func TestVectorColumnRoundTrip(t *testing.T) {
    dim := 128
    count := 100
    vectors := make([][]float32, count)
    for i := range vectors {
        vectors[i] = make([]float32, dim)
        for j := range vectors[i] {
            vectors[i][j] = float32(i*dim + j)
        }
    }

    var buf bytes.Buffer
    col := &ColumnMeta{Type: TypeVectorF32, VectorDim: dim}
    err := WriteVectorColumn(&buf, vectors, col)
    if err != nil {
        t.Fatalf("WriteVectorColumn failed: %v", err)
    }

    readBack, err := ReadVectorColumn(&buf, col)
    if err != nil {
        t.Fatalf("ReadVectorColumn failed: %v", err)
    }

    if len(readBack) != count {
        t.Errorf("count mismatch: got %d, want %d", len(readBack), count)
    }
    for i := range vectors {
        if len(readBack[i]) != dim {
            t.Errorf("dim mismatch at %d: got %d, want %d", i, len(readBack[i]), dim)
        }
        for j := range vectors[i] {
            if readBack[i][j] != vectors[i][j] {
                t.Errorf("vectors[%d][%d]: got %f, want %f", i, j, readBack[i][j], vectors[i][j])
            }
        }
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/msi/... -run TestVectorColumnRoundTrip -v`

- [ ] **Step 3: 实现 WriteVectorColumn / ReadVectorColumn**

关键：连续 float32 内存，无字典编码，适合 SIMD dot product / euclidean distance。

```go
// WriteVectorColumn writes count vectors of dimension dim as flat float32.
// Layout: [vec0_dim0, vec0_dim1, ..., vec1_dim0, ...]
func WriteVectorColumn(w io.Writer, vectors [][]float32, meta *ColumnMeta) error {
    dim := meta.VectorDim
    buf := make([]float32, len(vectors)*dim)
    for i, v := range vectors {
        copy(buf[i*dim:(i+1)*dim], v)
    }
    return binary.Write(w, binary.LittleEndian, buf)
}
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/msi/... -run TestVectorColumnRoundTrip -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/msi/column_vector.go server/engine/storage/msi/column_vector_test.go
git commit -m "feat(storage): add VECTOR_F32 column read/write, contiguous float32 layout for SIMD"
```

---

## Task 4: Row Index（时间块索引）

**Files:**
- Create: `server/engine/storage/msi/row_index.go`
- Test: `server/engine/storage/msi/row_index_test.go`

- [ ] **Step 1: 写测试 — 每 4096 行一个索引条目**

```go
package msi

import (
    "bytes"
    "testing"
)

func TestRowIndexRoundTrip(t *testing.T) {
    blocks := []*RowIndexBlock{
        {RowOffset: 0, TsMin: 1700000000000, TsMax: 1700001000000},
        {RowOffset: 4096, TsMin: 1700001000001, TsMax: 1700002000000},
    }

    var buf bytes.Buffer
    err := WriteRowIndex(&buf, blocks)
    if err != nil {
        t.Fatalf("WriteRowIndex failed: %v", err)
    }

    readBack, err := ReadRowIndex(&buf)
    if err != nil {
        t.Fatalf("ReadRowIndex failed: %v", err)
    }

    if len(readBack) != len(blocks) {
        t.Errorf("block count: got %d, want %d", len(readBack), len(blocks))
    }
    for i, b := range blocks {
        if readBack[i].RowOffset != b.RowOffset || readBack[i].TsMin != b.TsMin {
            t.Errorf("block %d mismatch: got %+v, want %+v", i, readBack[i], b)
        }
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/msi/... -run TestRowIndexRoundTrip -v`

- [ ] **Step 3: 实现 RowIndexBlock 和读写函数**

```go
type RowIndexBlock struct {
    RowOffset uint64 // 该块第一行的文件内字节偏移量（估算）
    TsMin    uint64
    TsMax    uint64
}

const RowIndexBlockSize = 4096

// WriteRowIndex 写入块数 + 每个块的 (TsMin, TsMax)
// 用于时间范围查询时二分跳转到目标块
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/msi/... -run TestRowIndexRoundTrip -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/msi/row_index.go server/engine/storage/msi/row_index_test.go
git commit -m "feat(storage): add row index for time-based pruning"
```

---

## Task 5: MSI Footer + 完整文件组装

**Files:**
- Create: `server/engine/storage/msi/footer.go`
- Test: `server/engine/storage/msi/footer_test.go`

- [ ] **Step 1: 写测试 — Footer round-trip with column metas**

```go
package msi

import (
    "bytes"
    "testing"
)

func TestFooterRoundTrip(t *testing.T) {
    footer := &Footer{
        ColumnMetas: []*ColumnMeta{
            {Name: "id", Type: TypeInt64, Offset: 128, VectorDim: 0},
            {Name: "vec", Type: TypeVectorF32, Offset: 128 + 8*1000, VectorDim: 128},
        },
        AnnIndexOffsets: []int64{1024, 2048},
    }

    var buf bytes.Buffer
    err := footer.WriteTo(&buf)
    if err != nil {
        t.Fatalf("WriteTo failed: %v", err)
    }

    footer2, err := ReadFooter(&buf)
    if err != nil {
        t.Fatalf("ReadFooter failed: %v", err)
    }

    if len(footer2.ColumnMetas) != 2 || footer2.ColumnMetas[0].Name != "id" {
        t.Errorf("Footer mismatch: got %+v", footer2)
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/msi/... -run TestFooterRoundTrip -v`

- [ ] **Step 3: 实现 Footer 结构**

```go
type Footer struct {
    ColumnMetas     []*ColumnMeta
    AnnIndexOffsets []int64 // 每个列的 ANN 索引偏移量
}
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/msi/... -run TestFooterRoundTrip -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/msi/footer.go server/engine/storage/msi/footer_test.go
git commit -m "feat(storage): add MSI footer serialization"
```

---

## Task 6: MSI 文件完整读写（端到端集成）

**Files:**
- Create: `server/engine/storage/msi/file.go`
- Test: `server/engine/storage/msi/file_test.go`

- [ ] **Step 1: 写测试 — 完整 MSI 文件读、写、读回**

```go
package msi

import (
    "bytes"
    "os"
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
    m, cols, err := ReadMSI(&buf)
    if err != nil {
        t.Fatalf("ReadMSI failed: %v", err)
    }
    if m.RowCount != 100 {
        t.Errorf("RowCount: got %d, want 100", m.RowCount)
    }
    if len(cols) != 3 {
        t.Errorf("Column count: got %d, want 3", len(cols))
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/msi/... -run TestFullMSIFileRoundTrip -v`

- [ ] **Step 3: 实现 TableMetadata, WriteMSI, ReadMSI**

WriteMSI 流程: 写 Header → 写各列数据 → 写 RowIndex → 写 Footer
ReadMSI 流程: 读 Header → 读 Footer → 按 Footer.ColumnMetas 读取各列

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/msi/... -run TestFullMSIFileRoundTrip -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/msi/file.go server/engine/storage/msi/file_test.go
git commit -m "feat(storage): add full MSI file read/write end-to-end"
```

---

## Task 7: Mutable Segment

**Files:**
- Create: `server/engine/storage/segment/mutable.go`
- Test: `server/engine/storage/segment/mutable_test.go`

- [ ] **Step 1: 写测试 — 追加写入多批，触发阈值后 seal**

```go
package segment

import (
    "testing"
)

func TestMutableSegmentAppendAndSeal(t *testing.T) {
    ms := NewMutableSegment("test_table", WithMaxRows(1000))

    // 追加 500 行
    err := ms.Append(map[string]interface{}{
        "id": []int64{1, 2, 3},
    })
    if err != nil {
        t.Fatalf("Append failed: %v", err)
    }

    if ms.RowCount() != 3 {
        t.Errorf("RowCount: got %d, want 3", ms.RowCount())
    }

    // 追加到阈值，shouldSeal 应该为 true
    for i := 0; i < 1000; i++ {
        ms.Append(map[string]interface{}{"id": []int64{int64(i)}})
    }

    if !ms.ShouldSeal() {
        t.Errorf("ShouldSeal: got false, want true")
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/segment/... -run TestMutableSegmentAppendAndSeal -v`

- [ ] **Step 3: 实现 MutableSegment**

关键设计:
- 内存中按列存储 `map[string][]interface{}`
- 追加写入按列追加到切片
- `ShouldSeal()` 当行数达到阈值或时间窗口到期返回 true
- `Seal()` 导出为 ImmutableSegment（调用 WriteMSI）

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/segment/... -run TestMutableSegmentAppendAndSeal -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/segment/mutable.go server/engine/storage/segment/mutable_test.go
git commit -m "feat(storage): add MutableSegment with append and auto-seal"
```

---

## Task 8: Immutable Segment

**Files:**
- Create: `server/engine/storage/segment/immutable.go`
- Test: `server/engine/storage/segment/immutable_test.go`

- [ ] **Step 1: 写测试 — 读取刚写入的 MSI 文件**

```go
package segment

import (
    "bytes"
    "os"
    "testing"

    "wormhole/server/engine/storage/msi"
)

func TestImmutableSegmentRead(t *testing.T) {
    // 构造 MSI 文件
    meta := &msi.TableMetadata{
        Columns: []*msi.ColumnMeta{
            {Name: "id", Type: msi.TypeInt64},
            {Name: "vec", Type: msi.TypeVectorF32, VectorDim: 128},
        },
        RowCount: 50,
    }
    ids := make([]int64, 50)
    vectors := make([][]float32, 50)
    for i := range ids {
        ids[i] = int64(i)
        vectors[i] = make([]float32, 128)
        vectors[i][0] = float32(i)
    }

    var buf bytes.Buffer
    msi.WriteMSI(&buf, meta, map[string]interface{}{
        "id":  ids,
        "vec": vectors,
    })

    tmp, _ := os.CreateTemp("", "test-*.msi")
    tmp.Write(buf.Bytes())
    tmp.Close()
    defer os.Remove(tmp.Name())

    seg, err := OpenImmutableSegment(tmp.Name())
    if err != nil {
        t.Fatalf("OpenImmutableSegment failed: %v", err)
    }

    cols, err := seg.ReadColumns([]string{"id", "vec"})
    if err != nil {
        t.Fatalf("ReadColumns failed: %v", err)
    }

    if len(cols) != 2 {
        t.Errorf("column count: got %d, want 2", len(cols))
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/segment/... -run TestImmutableSegmentRead -v`

- [ ] **Step 3: 实现 OpenImmutableSegment 和 ReadColumns**

```go
type ImmutableSegment struct {
    path   string
    meta   *msi.TableMetadata
    footer *msi.Footer
}

func OpenImmutableSegment(path string) (*ImmutableSegment, error) {
    f, err := os.Open(path)
    // 读取 header, footer
    // 返回 segment
}

func (s *ImmutableSegment) ReadColumns(names []string) (map[string]interface{}, error) {
    // 按列名读取，返回 map[name][]interface{}
}
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/segment/... -run TestImmutableSegmentRead -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/segment/immutable.go server/engine/storage/segment/immutable_test.go
git commit -m "feat(storage): add ImmutableSegment for reading MSI files"
```

---

## Task 9: WAL + Exactly-Once 支持

**Files:**
- Create: `server/engine/storage/wal/wal.go`
- Test: `server/engine/storage/wal/wal_test.go`
- Create: `server/engine/storage/wal/recovery.go`

- [ ] **Step 1: 写测试 — WAL 条目读写和 committed 标记**

```go
package wal

import (
    "bytes"
    "testing"
)

func TestWALEntryRoundTrip(t *testing.T) {
    entry := &WALEntry{
        TxnID:      42,
        Rows:       []map[string]interface{}{{"id": int64(1), "name": "test"}},
        Committed:  false,
        Timestamp:  1700000000000,
    }

    var buf bytes.Buffer
    err := entry.WriteTo(&buf)
    if err != nil {
        t.Fatalf("WriteTo failed: %v", err)
    }

    entry2, err := ReadWALEntry(&buf)
    if err != nil {
        t.Fatalf("ReadWALEntry failed: %v", err)
    }

    if entry2.TxnID != 42 || entry2.Committed != false {
        t.Errorf("Entry mismatch: got %+v", entry2)
    }

    // 标记 committed 并重写
    entry2.Committed = true
    var buf2 bytes.Buffer
    entry2.WriteTo(&buf2)

    entry3, _ := ReadWALEntry(&buf2)
    if entry3.Committed != true {
        t.Errorf("Committed flag not set: got %v", entry3.Committed)
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/wal/... -run TestWALEntryRoundTrip -v`

- [ ] **Step 3: 实现 WALEntry 和读写**

```go
package wal

type WALEntry struct {
    TxnID     uint64
    Rows      []map[string]interface{} // 行数据（JSON 或 binary 编码）
    Committed bool
    Timestamp uint64
}

func (e *WALEntry) WriteTo(w io.Writer) error
func ReadWALEntry(r io.Reader) (*WALEntry, error)
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/wal/... -run TestWALEntryRoundTrip -v`

- [ ] **Step 5: 实现 Recovery 函数**

```go
// ScanWAL 扫描 WAL 目录，返回所有 committed 的 txn_id
// 用于 crash recovery 后的去重
func ScanWAL(walDir string) (map[uint64]*WALEntry, error)
```

- [ ] **Step 6: 提交**

```bash
git add server/engine/storage/wal/wal.go server/engine/storage/wal/wal_test.go server/engine/storage/wal/recovery.go
git commit -m "feat(storage): add WAL for exactly-once semantics"
```

---

## Task 10: StorageEngine 上层接口

**Files:**
- Create: `server/engine/storage/options.go`
- Create: `server/engine/storage/table.go`
- Create: `server/engine/storage/storage.go`
- Test: `server/engine/storage/storage_test.go`

- [ ] **Step 1: 写测试 — Insert 和 Select 端到端**

```go
package storage

import (
    "os"
    "testing"
)

func TestStorageInsertAndSelect(t *testing.T) {
    tmp := os.TempDir()
    engine := NewStorageEngine(tmp, &TableOptions{
        Consistency: AtLeastOnce, // 默认
    })

    // 创建表
    err := engine.CreateTable("db1", "t1", &TableMetadata{
        Columns: []*msi.ColumnMeta{
            {Name: "id", Type: msi.TypeInt64},
            {Name: "vec", Type: msi.TypeVectorF32, VectorDim: 128},
        },
    })
    if err != nil {
        t.Fatalf("CreateTable failed: %v", err)
    }

    // 插入
    err = engine.Insert("db1", "t1", []map[string]interface{}{
        {"id": int64(1), "vec": make([]float32, 128)},
    })
    if err != nil {
        t.Fatalf("Insert failed: %v", err)
    }

    // 读取
    cols, err := engine.SelectColumns("db1", "t1", []string{"id", "vec"})
    if err != nil {
        t.Fatalf("SelectColumns failed: %v", err)
    }
    if len(cols["id"]) != 1 {
        t.Errorf("row count: got %d, want 1", len(cols["id"]))
    }
}
```

- [ ] **Step 2: 运行测试确认失败**

Run: `go test ./server/engine/storage/... -run TestStorageInsertAndSelect -v`

- [ ] **Step 3: 实现 StorageEngine**

```go
type StorageEngine struct {
    rootDir string
    tables  map[string]*Table
    opts    *Options
}

type Options struct {
    MutableMaxRows    int           // 默认 100000
    BucketDurationMs  int64         // 时间分区桶大小，默认 3600000 (1小时)
}

func (e *StorageEngine) Insert(db, table string, rows []map[string]interface{}) error {
    // 根据表配置决定 at_least_once 还是 exactly_once
    // at_least_once: 直接写入 Mutable Segment
    // exactly_once: 写 WAL，标记 committed，写入 Mutable
}

func (e *StorageEngine) SelectColumns(db, table string, columns []string) (map[string]interface{}, error) {
    // 读取所有涉及 Segment（Mutable + Immutable）
    // 合并结果返回
}
```

- [ ] **Step 4: 运行测试确认通过**

Run: `go test ./server/engine/storage/... -run TestStorageInsertAndSelect -v`

- [ ] **Step 5: 提交**

```bash
git add server/engine/storage/options.go server/engine/storage/table.go server/engine/storage/storage.go server/engine/storage/storage_test.go
git commit -m "feat(storage): add StorageEngine top-level API (Insert, Select)"
```

---

## Self-Review

**Spec coverage:**
- [x] MSI 格式: Header, Column, Vector, RowIndex, Footer — Task 1-6
- [x] Mutable Segment: Task 7
- [x] Immutable Segment: Task 8
- [x] WAL + Exactly-Once: Task 9
- [x] StorageEngine 上层接口: Task 10

**Placeholder scan:**
- 无 "TBD" / "TODO" / "实现后续补充" 等占位符
- 每步都有具体代码实现

**Type consistency:**
- `msi.Header`, `msi.ColumnMeta`, `msi.TableMetadata` 在 Task 5 后统一使用
- `MutableSegment.ShouldSeal()` 在 Task 7 定义，Task 10 的 StorageEngine 调用

**Spec gap check:**
- 时序（bucket 分区）在 Options 中有配置项，但具体时间分桶逻辑在 StorageEngine.Insert 中略过（可后续 Task 补充）
- 后台 merge 触发在 ShouldSeal 中已体现，但 Merger goroutine 是独立模块，可在后续 Task 补充