# wormhole 存储引擎与查询引擎设计

日期: 2026-05-21
状态: 设计阶段

---

## 1. 定位与目标

**定位**: OLAP 分析数据库，向量检索作为核心扩展能力。

**目标层次**:
1. 成为最好的 Golang 数据库
2. 成为最好的 Golang 向量数据库
3. 最终目标：最好的数据库之一

**workload**: 高频实时写入 + 即时查询。写路径和查询路径分开优化。

---

## 2. 整体架构

```
SQL Client (TUI)
       │ Thrift RPC
       ▼
┌──────────────────────────────────┐
│      RPC Layer (保留 Thrift)     │
└──────────┬───────────────────────┘
           │
    ┌──────▼───────────────────────────┐
    │      SQL Parser (保留 ANTLR)       │
    │         MyVisitor → Statement      │
    └──────────────┬─────────────────────┘
                   │
    ┌──────────────▼─────────────────────┐
    │         Query Engine (新设计)       │
    │                                     │
    │  Logical Plan Builder               │
    │       │                             │
    │       ▼                             │
    │  Physical Plan Optimizer             │
    │       │                             │
    │       ▼                             │
    │  DAG Executor (Push + 向量化)         │
    └──────────────┬─────────────────────┘
                   │
    ┌──────────────▼─────────────────────┐
    │       Storage Engine (新设计)        │
    │                                     │
    │  Write Path 🔴     Read Path 🔵     │
    │  Mutable Segment → Immutable Segments │
    └─────────────────────────────────────┘
```

### 2.1 保留的组件

- **Thrift RPC**: `rpc/common.thrift` 不变，继续使用 `TCompactProtocol`
- **ANTLR SQL Parser**: `SqlLexer.g4` / `SqlParser.g4` 扩展，语法兼容现有
- **Client TUI**: `client/tui.go` 不变

### 2.2 重写的组件

- `server/query/interpreter/`: 替换为完整的 Query Engine
- `server/engine/parquet/`: 替换为自研列式存储引擎

---

## 3. 存储引擎设计

### 3.1 目录结构

```
data/
├── db1/
│   └── table1/
│       ├── _meta.json              # 表 schema, 列定义, 向量维度
│       ├── _wal/                   # WAL 目录 (exactly_once 模式)
│       │   ├── 000001.wal         # txn_id=1, committed=true
│       │   └── 000002.wal         # txn_id=2, committed=false
│       ├── mutable/
│       │   └── 1.msy              # Mutable Segment (内存或 mmap)
│       └── immutable/
│           ├── 000001.msi         # Immutable Segment (按时间分桶)
│           └── 000002.msi
│           ├── 000001.msi.col0.ann # 列0的 HNSW 索引
│           └── 000001.msi.col1.ann # 列1的向量索引
```

### 3.2 自研列式格式 (.msi)

**设计原则**:
- Vector-first: 向量列优先，支持固定维度 float32 数组的连续内存布局
- SIMD 友好: INT/FLOAT/DOUBLE 只用 Plain 编码（连续内存），直接 SIMD 批量读取
- 时间有序: 每行数据在文件内按时间戳排序存储，支持二分跳过

**文件格式**:

```
MSI Header:
  Magic: 0xMSI1 (4 bytes)
  Version: uint16
  Column Count: uint16
  Row Count: uint64
  Timestamp Min/Max: uint64 x2
  Column Meta[]: 每列的偏移量、压缩算法、索引位置

Column Data (Per Column):
  - 物理类型: INT32 / INT64 / FLOAT32 / FLOAT64 / BYTE_ARRAY / VECTOR_F32
  - 编码: PLAIN / DICT (低基数字符串)
  - 数据: 连续内存布局，适合 SIMD 扫描

Vector Column (差异化的核心):
  - 固定维度，如 FLOAT[128]
  - 连续 float32 数组，无字典编码，直接 SIMD distance 计算
  - 每列自带 ANN 索引元数据（HNSW 参数、索引偏移量）

Row Index:
  - 每 4096 行一个索引块，记录该块的 timestamp 范围
  - 支持时间范围查询时直接跳转

Footer:
  - Column Meta 数组的偏移量
  - ANN 索引偏移量数组
  - Schema 完整定义
```

**与 Parquet 的差异化**:
| 特性 | Parquet | wormhole .msi |
|------|---------|---------------|
| 向量列 | 通用类型，无特殊优化 | VECTOR_F32，连续内存，SIMD 友好 |
| 时间索引 | 无文件内时间索引 | 行级时间戳 + 块索引，二分跳转 |
| SIMD | 依赖解码后处理 | Plain 编码，直接 SIMD |
| 索引内嵌 | 无 | 每列独立 ANN 索引在同一文件内 |
| Exactly-Once | 无 | WAL 层支持 |

### 3.3 Mutable Segment

- 写入缓冲: 接收 INSERT 数据，内存中按列存储
- 可选 mmap: 数据量大时使用 memory-mapped file，减少 GC 压力
- 后台 merge: 当 Mutable 超过阈值（行数或时间），合并为 Immutable Segment
- 无索引: 只支持追加写入，不建 ANN 索引（索引在 Immutable 阶段构建）

### 3.4 Immutable Segment

- 只读列文件，已按时间排序
- 每个 Segment 对应的列数据建 ANN 索引
- 支持并行扫描: 多个 Immutable Segment 可被 DAG 并行读取

### 3.5 写入语义

**表创建时配置**:
```sql
CREATE TABLE events (
  id BIGINT,
  payload STRING,
  vec FLOAT[128]
) WITH (consistency = 'exactly_once')   -- 默认 'at_least_once'
```

| 特性 | `at_least_once` (默认) | `exactly_once` |
|------|----------------------|----------------|
| 写入路径 | 直接写 Mutable | 写 WAL → Mutable → 标记 committed |
| fsync | 可选批量异步 | 每次事务同步刷 WAL |
| 重复写入 | 可能重复 | 幂等去重 (txn_id 过滤) |
| 性能 | 高吞吐 | 约 50% 下降 |
| 适用 | 日志、时序 | 金融、订单 |

**WAL 设计**:
- 单独文件: `data/<db>/<table>/_wal/<txn_id>.wal`
- 每事务一个文件，文件名 = txn_id（方便二分查找）
- 内容: `{txn_id, rows, committed, timestamp}`
- Recovery: 扫描所有 WAL，保留 committed=true 的记录

---

## 4. 查询引擎设计

### 4.1 三层转换

```
SQL
  │
  ▼
Logical Plan (DAG of Operators)
  │ Filter(ts > '2025-01-01')
  │   └─ TableScan → [columns: id, vec, ts]
  │
  ▼
Physical Plan (Stage-level)
  │ Stage 0: TableScan → Filter → [Chunk]
  │ Stage 1: KnnSearch(vec, query, k=10) → [ids, distances]
  │ Stage 2: TopK(SORT BY distance ASC, LIMIT 10) → [output]
  │
  ▼
DAG Execution (并行)
  │ goroutine per Stage, channel 桥接
  │ Push-based streaming, 向量化列扫描
```

### 4.2 Logical Plan

SQL AST 通过 MyVisitor 生成 Operator DAG:
- `TableScan`: 读取指定列，指定 Segment 列表
- `Filter`: 布尔表达式求值，支持 SIMD 批量处理
- `KnnSearch`: 向量 ANN 搜索
- `Projection`: 列投影、表达式计算
- `Sort`: 排序
- `Limit`: 取前 N 行（可提前终止上游）

### 4.3 Physical Plan

按 Stage 划分，Stage 之间并行:
- Stage 0: Source — 读取列数据
- Stage 1: 计算密集型 — Filter、KnnSearch、Projection
- Stage 2: 结果汇总 — Sort、Limit、Aggregate

Stage 划分依据: 数据依赖和并行度。

### 4.4 DAG Executor

**Push + 向量化混合模式**:
- 行级 Push: Source 通过 channel 向上游推送 Chunk（128行/批）
- 列级 SIMD: Filter 在 Chunk 内用 SIMD 批量求值
- 背压机制: channel 满时 Source 停止推送，下游消费后自动恢复

**Chunk 结构**:
```go
type Chunk struct {
    Cols   [][]interface{}   // 按列存储，128行/批
    Rows   int               // 实际行数
    Stats  ColumnStats       // min/max，用于剪枝
}
```

**并发模型**:
- 每个 Stage 启动 `parallelism` 个 goroutine worker
- Stage 内多个 worker 共享 input channel（无锁）
- Stage 间用 channel 桥接，无共享状态

### 4.5 向量检索融入 DAG

```sql
SELECT id, vec, distance(vec, query_vec) AS dist
FROM embeddings
WHERE ts > '2025-01-01'
ORDER BY dist ASC
LIMIT 10
```

物理计划:
```
Stage 0: TableScan(ts, id, vec) → Filter(ts > '2025-01-01') → [col_chunks]
Stage 1: KnnSearch(vec, query_vec, k=10) → [ids, distances]
Stage 2: Projection(id, dist) → TopK(SORT dist ASC, LIMIT 10) → [output]
```

向量索引查找不改变 Stage 划分，只替换 Stage 1 的 operator 实现。

---

## 5. 向量索引设计

### 5.1 插件式接口

```go
type VectorIndex interface {
    Build(vectors [][]float32, opts IndexOptions) error
    Search(query []float32, k int, opts SearchOptions) ([]int, []float32, error)
    Serialize(w io.Writer) error
    Load(r io.Reader) error
}

type IndexOptions struct {
    Dim            int
    M             int   // HNSW M 参数
    EfConstruction int  // 构建时搜索宽度
}

type SearchOptions struct {
    EfSearch       int   // 查询时搜索宽度
}
```

### 5.2 默认实现: HNSW (纯 Go)

**为何选 HNSW**:
- 工业主流: Milvus、Qdrant、Weaviate 都在用
- 纯 Go 无 CGO 依赖，交叉编译无障碍
- 内存占用可控，比 IVF-PQ 简单

**实现优势**:
- **mmap 索引**: 索引文件 mmap 到内存，query 时按需 pageFault，不加载整个索引到堆
- **Level-0 优化**: <10k vectors 不建跳表，直接暴力扫描 + 裁剪，性能反而更好
- **批量构建**: 多条向量 INSERT 时批量 build HNSW 图，避免图结构碎片化
- **ef_construction 可配置**: 用户在索引质量和构建速度间 trade-off

**与 Milvus 的差异**:
- Milvus 的 HNSW 是 C++ 实现 + gRPC，本项目是纯 Go 内嵌
- mmap 策略减少 GC 压力，适合 long-running 查询服务

---

## 6. 权限控制设计

### 6.1 实现方式: 内嵌 DAG

权限检查本质是 filter，不做独立服务，理由:
- 权限 filter 和普通 WHERE 条件一样走 DAG 的 Filter operator
- 单节点 Go 服务，无 sidecar 架构复杂度
- 性能开销可忽略

### 6.2 行级安全策略 (RLS)

```sql
CREATE TABLE events (
  id BIGINT,
  user_id BIGINT,
  payload STRING
) WITH (rls = 'user_id = current_user()');
```

TableScan operator 在执行时自动注入 RLS 表达式。

### 6.3 列级权限

```sql
GRANT SELECT(id, payload) ON events TO analyst;
GRANT SELECT(vec) ON events TO vector_role;
```

列级权限在 TableScan 时裁剪列，敏感列（如 password）对未授权角色不可见。

### 6.4 数据库/表级权限

标准 RBAC: `GRANT SELECT ON database.table TO user`。

---

## 7. 时序语义

### 7.1 时间有序保证

- 写入时自动附加隐式 `timestamp` 字段（毫秒级 Unix 时间戳）
- Mutable Segment 和 Immutable Segment 内数据按时间排序存储
- 支持用户显式时间列（如果 schema 包含 ts 字段）

### 7.2 时间分区

Immutable Segment 按时间分桶:
- 桶大小可配置（1小时/1天/1周）
- 查询时 DAG 根据时间范围裁剪不需要的 Segment

### 7.3 Exactly-Once 写入

见 3.5 节。通过 WAL + txn_id 去重实现。

---

## 8. 待确认 / 未来工作

- [ ] 列文件格式的详细二进制布局（Header、Page、Footer 结构）
- [ ] WAL 的清理策略（committed 后多久删除）
- [ ] HNSW 纯 Go 实现的具体参数选择
- [ ] 时间分区桶大小的默认值和配置方式
- [ ] 向量距离函数的可扩展性（DISTANCE EUCLIDEAN / COSINE / DOT）