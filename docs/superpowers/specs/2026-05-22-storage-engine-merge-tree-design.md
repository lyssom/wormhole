# Storage Engine Phase 1: MergeTree with Partition

**Goal:** 实现带 Partition 的最小可用 MergeTree，支持时间分区和后台合并。

**Architecture:** 自研列式格式 MSI + Mutable Segment + Immutable Segment，后台 MergeScheduler 负责分区内的 Segments 合并。

**Tech Stack:** 纯 Go，无 CGO 依赖。复用现有 MSI 格式、Segment 架构、WAL 机制。

---

## File Structure

```
server/engine/storage/
├── msi/                          # 已有
├── segment/                      # 已有
│   ├── mutable.go
│   ├── immutable.go
│   └── merger.go                 # 新增：BackgroundMergeTask
├── wal/                          # 已有
├── partition/
│   ├── partition.go              # PartitionManager: 根据 bucket 计算分区路由
│   ├── partition_test.go
│   └── metadata.go              # PartitionMetadata: 分区内 Segments 列表
├── scheduler/
│   ├── scheduler.go             # MergeScheduler: 后台守护，定期扫描合并候选
│   ├── scheduler_test.go
│   └── merge_task.go            # MergeTask: 单次合并任务
├── options.go                    # 修改：添加 bucket_duration_ms
├── table.go                      # 修改：添加分区路由逻辑
└── storage.go                    # 修改：添加 MergeScheduler 集成
```

---

## Data Layout

```
data/<database>/<table>/
├── _partitions/                  # 分区元数据目录
│   └── partition_2024-01.json   # 每个分区一个文件
├── bucket_<timestamp_ms>/        # 分区目录
│   ├── mutable/                  # 当前活跃的 Mutable Segment
│   │   └── seg.msi
│   └── immutable/               # 已 seal 的 Immutable Segments
│       ├── 1.msi
│       ├── 2.msi
│       └── 3.msi
```

**Partition 计算：**
```go
bucketID = ts / bucket_duration_ms
partitionPath = "bucket_" + bucketID
```

---

## Core Components

### 1. PartitionManager

**职责：** 根据 timestamp 计算分区路由，决定 Insert 数据写入哪个分区。

```go
type PartitionManager struct {
    bucketDurationMs int64
}

func (pm *PartitionManager) GetPartitionPath(ts int64) string {
    bucketID := ts / pm.bucketDurationMs
    return fmt.Sprintf("bucket_%d", bucketID)
}
```

### 2. PartitionMetadata

**职责：** 管理每个分区内所有 Segments 的元数据（用于合并决策）。

```go
type PartitionMetadata struct {
    PartitionPath string
    Segments      []SegmentMeta  // 按时间排序
    TotalRows     uint64
    TotalSize     uint64
    LastModified  time.Time
}

type SegmentMeta struct {
    Path       string
    Type       SegmentType  // Mutable / Immutable
    RowCount   uint64
    SizeBytes  uint64
    TsMin      uint64
    TsMax      uint64
}
```

**存储：** `data/<db>/<table>/_partitions/partition_<bucketID>.json`

### 3. MergeScheduler

**职责：** 后台 goroutine，定期扫描所有分区的合并候选，触发合并任务。

```go
type MergeScheduler struct {
    interval    time.Duration
    mergeThresh  int           // 超过 N 个 Segments 就合并
    partitions   map[string]*PartitionMetadata
    mu          sync.RWMutex
}

func (s *MergeScheduler) Run(ctx context.Context) {
    ticker := time.NewTicker(s.interval)
    for {
        select {
        case <-ticker.C:
            s.scanAndMerge()
        case <-ctx.Done():
            return
        }
    }
}
```

**合并触发条件：**
- 同一分区内 Immutable Segments 数量 >= mergeThresh (默认 3)
- 新数据时间窗口与旧数据有明显边界（可选优化）

### 4. MergeTask

**职责：** 执行单次合并任务 — 读取多个 Segments，合并为一个新 Segment。

```go
type MergeTask struct {
    partitionPath string
    inputs        []*ImmutableSegment
    outputPath    string
}

func (t *MergeTask) Execute() error {
    // 1. 读取所有 input segments
    // 2. 按主键/时间排序合并
    // 3. 写入新的 MSI 文件
    // 4. 更新 PartitionMetadata
    // 5. 删除旧的 Segments
}
```

**合并策略（一期简化）：**
- 只合并时间上相邻的 Segments
- 不做全量排序（按 ts 顺序追加即可）

---

## Modified Components

### 1. MutableSegment (已有 Seal 方法)

Seal 后需要：
1. 更新 PartitionMetadata，标记该 Segment 为 Immutable
2. 通知 MergeScheduler 有新的合并候选

### 2. Table.Insert

路由到正确分区：
```go
func (t *Table) Insert(rows []map[string]interface{}) error {
    ts := extractTimestamp(rows)
    partitionPath := t.partMgr.GetPartitionPath(ts)
    seg := t.getOrCreateMutableSegment(partitionPath)
    return seg.Append(rows)
}
```

### 3. StorageEngine

管理 MergeScheduler 生命周期：
```go
type StorageEngine struct {
    scheduler *MergeScheduler
    // ...
}

func (e *StorageEngine) Start() {
    ctx, cancel := context.WithCancel(context.Background())
    e.scheduler.Run(ctx)
}

func (e *StorageEngine) Stop() {
    // 通知 scheduler 停止
}
```

---

## WAL + Exactly-Once (已有，复用)

WAL 机制不变，用于 crash recovery 后的去重。
Merge 操作本身是幂等的（合并结果唯一），不需要 WAL。

---

## Workflow: Write → Seal → Merge

```
Insert(ts=1700000000000)
  → PartitionManager 计算 bucket=1700000000
  → MutableSegment.append(batch)
  → MutableSegment.rowCount >= maxRows
  → ShouldSeal() = true
  → Seal() 生成 ImmutableSegment
  → 更新 PartitionMetadata
  → MergeScheduler 检测到候选

[后台] MergeScheduler.scanAndMerge()
  → 读取 partition_<bucket>.json
  → 发现 3 个 ImmutableSegments
  → 满足 mergeThresh
  → 创建 MergeTask
  → 读取 3 个 MSI，排序合并
  → 生成 merged.msi
  → 更新 PartitionMetadata
  → 删除旧 Segments
```

---

## Testing Strategy

### 1. 单元测试

- `partition_test.go`: bucket 计算、分区路由
- `metadata_test.go`: PartitionMetadata 读写
- `scheduler_test.go`: 合并触发条件、候选选择

### 2. 集成测试

- `seal_integration_test.go`: Mutable → Immutable 完整流程
- `merge_integration_test.go`: 多 Segments 合并

### 3. 端到端演示

```go
func DemoLogAnalysis() {
    // 1. 创建表，bucket_duration = 1小时
    engine.CreateTable("logs", "access_log", opts)

    // 2. 模拟写入 3 小时数据（触发 Seal）
    for i := 0; i < 3; i++ {
        engine.Insert("logs", "access_log", generateBatch(i))
        time.Sleep(100 * time.Millisecond)
    }

    // 3. 等待后台合并完成
    time.Sleep(2 * time.Second)

    // 4. 查询验证
    result := engine.SelectColumns("logs", "access_log", columns)
    // 验证数据完整性
}
```

---

## Decisions & Open Questions

| 决策点 | 选择 | 理由 |
|-------|------|------|
| Partition 粒度 | 用户自定义 bucket_duration_ms | 灵活，对标 Milvus |
| 合并触发 | 超过 3 个 Segments | 平衡合并频率和资源消耗 |
| 合并策略 | 简单时间顺序合并 | 一期不做全量排序 |
| Mutable Seal | 阈值触发 (maxRows) | 已有实现 |

**未解决：**
- 按主键排序还是按时间排序？一期按 ts 顺序
- 合并时如何处理 Vector 列？和普通列一样合并
- 合并失败如何回滚？一期简化处理

---

## Scope (Phase 1 不包含)

- 主键稀疏索引（Phase 2）
- Mutation (DELETE/UPDATE)（Phase 2）
- TTL（Phase 2）
- 向量索引（Phase 3）
- 跨分区查询优化（Phase 2+）