# Wormhole

```
    ⭐━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━⭐
         __                                   __
        /  \                                 /  \
       |    |  ⚡ W O R M H O L E ⚡         |    |
        \__/     The Database That Feels      \__/
               Like Opening a Portal
    ⭐━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━⭐
```

**Wormhole** is a lightweight SQL database built for speed and simplicity. It combines an Apache Thrift RPC frontend, an ANTLR-powered SQL parser, and a high-performance storage engine to deliver a database experience that's both powerful and delightful to use.

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-Apache--2.0-blue?style=flat-square)](LICENSE)

---

## English

### Features

- **SQL at Your Fingertips** — Full support for `SELECT`, `INSERT`, `UPDATE`, `DELETE`, `CREATE DATABASE`, `CREATE TABLE`, and `USE` statements
- **ANTLR-Powered Parser** — Rock-solid SQL parsing with ANTLR4, generating parse trees that feed a type-safe statement visitor
- **Thrift RPC Interface** — Clean client-server communication via Apache Thrift with `TCompactProtocol`
- **High-Performance Storage** — MSI (Mutable/Immutable) storage engine with Write-Ahead Logging (WAL) for durability
- **Beautiful TUI Client** — A terminal UI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) that makes querying feel like magic
- **Configuration via YAML** — Simple `config.yaml` for data paths, row limits, and bucket durations

### Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         Client (TUI)                        │
│                    bubble tea + Thrift                      │
└─────────────────────────┬───────────────────────────────────┘
                          │ Thrift RPC
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                     Wormhole Server                         │
│  ┌─────────────┐  ┌─────────────┐  ┌────────────────────┐ │
│  │ ANTLR4      │→ │  MyVisitor  │→ │   Interpreter      │ │
│  │ Lexer/Parser│  │  (AST Gen)  │  │   (Executor)       │ │
│  └─────────────┘  └─────────────┘  └─────────┬──────────┘ │
│                                              │             │
└──────────────────────────────────────────────┼─────────────┘
                                               │ SQL Result
                                               ▼
┌─────────────────────────────────────────────────────────────┐
│                   Storage Engine (MSI)                      │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────┐  │
│  │   Mutable    │→ │  Immutable   │→ │   WAL (Durable)  │  │
│  │   Segment   │  │   Segment    │  │   Write-Ahead    │  │
│  └──────────────┘  └──────────────┘  └──────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

### Quick Start

```bash
# Build
go build ./...

# Start the server (default: 10.67.0.15:1320)
go run ./cmd/server/

# In another terminal, start the TUI client
go run ./cmd/client/
```

### Configuration

Edit `server/config.yaml`:

```yaml
storage:
  data_path: "./data"
  mutable_max_rows: 100000
  bucket_duration_ms: 3600000
```

### SQL Examples

```sql
-- Create a database
CREATE DATABASE mydb;

-- Use it
USE mydb;

-- Create a table
CREATE TABLE users (
  id INT,
  name TEXT,
  email TEXT
);

-- Insert data
INSERT INTO users VALUES (1, 'Alice', 'alice@wormhole.dev');
INSERT INTO users VALUES (2, 'Bob', 'bob@wormhole.dev');

-- Query
SELECT * FROM users WHERE id = 1;

-- Update
UPDATE users SET email = 'alice@newmail.dev' WHERE id = 1;

-- Delete
DELETE FROM users WHERE id = 2;
```

### Project Structure

```
wormhole/
├── cmd/
│   ├── client/          # TUI client entry point
│   └── server/         # Server entry point
├── rpc/
│   └── common.thrift   # Thrift RPC definition
├── server/
│   ├── config/         # Configuration loading
│   ├── storage/        # MSI storage engine
│   └── sql/
│       ├── parser/      # ANTLR4 generated parser
│       ├── interpreter/ # SQL execution engine
│       └── statement/  # Typed SQL statement structs
└── client/
    └── tui.go          # Bubble Tea terminal UI
```

### Tech Stack

| Layer | Technology |
|-------|------------|
| Language | Go 1.19+ |
| RPC | Apache Thrift 0.18.1 |
| SQL Parser | ANTLR4 |
| TUI | Charm Bubble Tea |
| Storage | Custom MSI Engine + WAL |

---

## 中文

**Wormhole** 是一个轻量级 SQL 数据库，专为速度和简洁而生。它融合了 Apache Thrift RPC 前端、ANTLR 驱动的 SQL 解析器和高性能存储引擎，带来既强大又令人愉悦的数据库体验。

### 核心特性

- **完整 SQL 支持** — 支持 `SELECT`、`INSERT`、`UPDATE`、`DELETE`、`CREATE DATABASE`、`CREATE TABLE` 和 `USE` 语句
- **ANTLR 解析引擎** — 基于 ANTLR4 的可靠 SQL 解析，生成类型安全的 AST 语句
- **Thrift RPC 接口** — 通过 Apache Thrift 和 `TCompactProtocol` 实现简洁的客户端-服务器通信
- **高性能存储** — MSI（可变/不可变）存储引擎，配备预写式日志（WAL）保障持久性
- **精美 TUI 客户端** — 基于 [Bubble Tea](https://github.com/charmbracelet/bubbletea) 的终端界面，让查询如魔法般体验
- **YAML 配置** — 通过 `config.yaml` 轻松配置数据路径、行数限制和桶时长

### 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                        客户端 (TUI)                         │
│                    Bubble Tea + Thrift                      │
└─────────────────────────┬───────────────────────────────────┘
                          │ Thrift RPC
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                       Wormhole 服务器                        │
│  ┌─────────────┐  ┌─────────────┐  ┌────────────────────┐   │
│  │  ANTLR4     │→ │  MyVisitor  │→ │    解释器          │   │
│  │ 词法/解析器  │  │   (AST生成) │  │    (执行器)        │   │
│  └─────────────┘  └─────────────┘  └─────────┬──────────┘   │
│                                              │              │
└──────────────────────────────────────────────┼──────────────┘
                                               │ SQL 结果
                                               ▼
┌─────────────────────────────────────────────────────────────┐
│                     存储引擎 (MSI)                          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────┐   │
│  │   可变段      │→ │  不可变段    │→ │  WAL (预写日志)  │   │
│  │ Mutable      │  │  Immutable   │  │  Write-Ahead     │   │
│  └──────────────┘  └──────────────┘  └──────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

### 快速开始

```bash
# 编译
go build ./...

# 启动服务器（默认: 10.67.0.15:1320）
go run ./cmd/server/

# 另一个终端，启动 TUI 客户端
go run ./cmd/client/
```

### 配置说明

编辑 `server/config.yaml`:

```yaml
storage:
  data_path: "./data"
  mutable_max_rows: 100000
  bucket_duration_ms: 3600000
```

### SQL 示例

```sql
-- 创建数据库
CREATE DATABASE mydb;

-- 使用数据库
USE mydb;

-- 创建表
CREATE TABLE users (
  id INT,
  name TEXT,
  email TEXT
);

-- 插入数据
INSERT INTO users VALUES (1, 'Alice', 'alice@wormhole.dev');
INSERT INTO users VALUES (2, 'Bob', 'bob@wormhole.dev');

-- 查询数据
SELECT * FROM users WHERE id = 1;

-- 更新数据
UPDATE users SET email = 'alice@newmail.dev' WHERE id = 1;

-- 删除数据
DELETE FROM users WHERE id = 2;
```

### 技术栈

| 层级 | 技术 |
|------|------|
| 编程语言 | Go 1.19+ |
| RPC | Apache Thrift 0.18.1 |
| SQL 解析器 | ANTLR4 |
| 终端界面 | Charm Bubble Tea |
| 存储引擎 | 自研 MSI 引擎 + WAL |

---

<p align="center">
  <strong>Wormhole — 打开数据之门</strong><br>
  <em>The Database That Feels Like Opening a Portal</em>
</p>
