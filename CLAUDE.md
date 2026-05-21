# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go build ./...                    # Build all packages
go run ./cmd/server/              # Start the Thrift RPC server (default: 10.67.0.15:1320)
go run ./cmd/client/              # Start the TUI SQL client
```

No Makefile, tests, or linting configuration exist yet.

## Architecture

**wormhole** is a SQL database with a Thrift RPC frontend, an ANTLR-based SQL parser, and Apache Parquet as the storage engine.

### RPC Layer (Apache Thrift)

The RPC contract is defined in `rpc/common.thrift` — a single `RPCService` with `ping()` and `executeQueryStatement(QueryReq): QueryResp`. Generated Go code lives in `rpc/gen-go/common/`. The server uses `TCompactProtocol` and `TSimpleServer`.

### SQL Query Pipeline

1. Client sends SQL string via Thrift to `server.ExecuteQueryStatement`
2. ANTLR lexer/parser (`server/query/parser/`) tokenizes and parses the SQL into a parse tree. Grammars: `SqlLexer.g4` and `SqlParser.g4`.
3. `MyVisitor` (`server/query/sqlVisitor.go`) walks the parse tree and produces typed statement structs from `server/query/statement/` (implementing `Statement` interface with `GetKind() string`)
4. `server.Execute()` dispatches based on `GetKind()` return value to the `Interpreter` in `server/query/interpreter/`

### Supported SQL Statements

| Kind string | Statement type | Interpreter method |
|---|---|---|
| `"select"` | `QueryStatement` | `ExecuteFetchColumns` |
| `"create database"` | `CreateDatabaseStatement` | `ExecuteCreateDatabase` |
| `"create table"` | `CreateTableStatement` | `ExecuteCreateTable` |
| `"insert"` | `InsertStatement` | `ExecuteInsert` |
| `"use"` | `UseStatement` | `ExecuteUse` |

### Storage Engine

Data is stored as Parquet files on disk, rooted at a `data/` directory resolved relative to the interpreter source file via `runtime.Caller`. The layout:

```
data/<database>/<table>/table.json     # Column schema (JSON-encoded CreateTableStatement)
data/<database>/<table>/1.parquet      # Column data in Parquet format
```

The Parquet engine under `server/engine/parquet/` is a fork/variant of the Apache Arrow Go Parquet library. Key modules: `file/` (readers/writers), `schema/`, `metadata/`, `compress/` (brotli, snappy, gzip, zstd), `memory/`.

### Client

A terminal UI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) (`client/tui.go`). Presents a `sql: ` prompt in a loop, sends each input to the server and prints the result.

### Code Generation

- **ANTLR**: The `.g4` grammar files in `server/query/` generate the parser/lexer in `server/query/parser/`. Generated with `antlr` CLI tool.
- **Thrift**: `rpc/common.thrift` generates the RPC client/server stubs in `rpc/gen-go/common/`. Generated with `thrift` CLI tool (version 0.18.1).

## Skill routing

When the user's request matches an available skill, invoke it via the Skill tool. When in doubt, invoke the skill.

Key routing rules:
- Product ideas/brainstorming → invoke /office-hours
- Strategy/scope → invoke /plan-ceo-review
- Architecture → invoke /plan-eng-review
- Design system/plan review → invoke /design-consultation or /plan-design-review
- Full review pipeline → invoke /autoplan
- Bugs/errors → invoke /investigate
- QA/testing site behavior → invoke /qa or /qa-only
- Code review/diff check → invoke /review
- Visual polish → invoke /design-review
- Ship/deploy/PR → invoke /ship or /land-and-deploy
- Save progress → invoke /context-save
- Resume context → invoke /context-restore