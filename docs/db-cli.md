# DB CLI 概述、使用与扩展说明

## 目标

`cmd/server db` 是当前仓库唯一的数据库命令入口。它负责暴露数据库创建 SQL、demo schema SQL 以及 demo Todo CRUD 操作，但不直接拼接业务 SQL，也不恢复旧的 `initdb`、SQL 脚本目录或运行时 `AutoMigrate` 路径。

当前约束：

- DDL 和 CRUD SQL 必须通过 `pkg/sqlgen` 的链式 API 生成。
- CLI 只负责参数解析、调用应用服务和打印结果。
- SQL 生成和执行边界位于 `internal/app/dbapp`。
- `--apply` 只用于显式执行 sqlgen 生成的 DDL。
- 生产 schema 变更仍需要单独确认的迁移流程，不由当前 DB CLI 自动承担。

## 快速使用

预览 demo schema SQL：

```bash
go run ./cmd/server db --config=configs/config.yaml --operation=schema
```

执行 demo schema SQL：

```bash
go run ./cmd/server db --config=configs/config.yaml --operation=schema --apply
```

预览数据库创建 SQL：

```bash
go run ./cmd/server db --config=configs/config.yaml --operation=database
```

执行数据库创建 SQL：

```bash
go run ./cmd/server db --config=configs/config.yaml --operation=database --apply
```

demo Todo CRUD 示例：

```bash
go run ./cmd/server db --config=configs/config.yaml --operation=todo-create --title="demo" --description="created by db cli"
go run ./cmd/server db --config=configs/config.yaml --operation=todo-list --limit=20 --offset=0
go run ./cmd/server db --config=configs/config.yaml --operation=todo-get --id=1
go run ./cmd/server db --config=configs/config.yaml --operation=todo-update --id=1 --title="updated" --completed=true
go run ./cmd/server db --config=configs/config.yaml --operation=todo-delete --id=1
```

如果需要同时查看 CRUD 生成的 SQL，可追加：

```bash
--print-sql
```

## 操作列表

| Operation | 行为 | 默认是否执行数据库连接 |
|---|---|---|
| `database` | 通过 `sqlgen.DatabaseIfNotExists` 生成数据库创建 SQL | 否，只有 `--apply` 时连接 |
| `schema` | 通过 `sqlgen.TableIfNotExists` 生成 demo schema SQL | 否，只有 `--apply` 时连接 |
| `todo-create` | 通过 sqlgen create 链路写入 demo Todo | 是 |
| `todo-list` | 通过 sqlgen query 链路分页读取 demo Todo | 是 |
| `todo-get` | 通过 sqlgen query 链路按 ID 读取 demo Todo | 是 |
| `todo-update` | 通过 sqlgen update 链路更新 demo Todo | 是 |
| `todo-delete` | 通过 sqlgen delete 链路删除 demo Todo | 是 |

## 常用参数

| 参数 | 默认值 | 说明 |
|---|---|---|
| `--config` | `configs/config.yaml` | 配置文件路径，也可由 `RIN_CONFIG_PATH` 指定 |
| `--operation` | `schema` | 数据库操作类型 |
| `--apply` | `false` | 对 `database` 或 `schema` 执行生成的 DDL |
| `--id` | `0` | Todo ID，供 get/update/delete 使用 |
| `--title` | 空 | Todo 标题，供 create/update 使用 |
| `--description` | 空 | Todo 描述，供 create/update 使用 |
| `--completed` | 空 | Todo 完成状态，接受 `true` 或 `false` |
| `--limit` | `dbapp.DefaultTodoLimit` | list 分页大小 |
| `--offset` | `0` | list 分页偏移 |
| `--print-sql` | `false` | CRUD 执行后打印 sqlgen 生成的 SQL |

## 分层边界

`cmd/server/db.go`：

- 定义命令名、参数、操作分发和输出格式。
- 对 `database` / `schema` 的只读预览保持无副作用，不初始化数据库连接。
- 不直接维护建库、建表或 CRUD SQL 文本。

`internal/app/dbapp`：

- 根据配置 driver 选择 sqlgen dialect。
- 封装数据库创建、demo schema、Todo CRUD 的应用服务函数。
- 负责把 sqlgen 生成结果交给 `pkg/database` 执行。

`pkg/sqlgen`：

- 提供数据库、表结构、查询、创建、更新、删除的链式 SQL 生成能力。
- 新增 SQL 能力时应优先在这里补生成器和测试，而不是在 CLI 或应用层拼接字符串。

## 扩展流程

新增 DB CLI 操作时，按下面顺序推进：

1. 在 `cmd/server/db.go` 增加 operation 常量、参数读取和 switch 分支。
2. 在 `internal/app/dbapp` 增加输入结构和应用服务函数。
3. 在应用服务函数中只使用 `pkg/sqlgen` 链式 API 生成 SQL。
4. 如果 sqlgen 缺少能力，先扩展 `pkg/sqlgen`，并补充生成器测试。
5. 在 `cmd/server` 或 `internal/app/dbapp` 增加覆盖命令语义和 SQL 生成边界的测试。
6. 更新本文档和状态文档。

禁止事项：

- 不新增手写 SQL 脚本作为建库、建表或 CRUD 的主路径。
- 不恢复 `cmd/server initdb`。
- 不恢复 InitDB 配置段。
- 不把 demo schema 的 `AutoMigrate` 挂回 server start 或 reload。
- 不在未经确认的环境中执行 production schema 变更。

## 验证建议

修改 DB CLI 或 dbapp 后至少执行：

```bash
go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1
git diff --check
```

如果变更影响公共配置、装配或运行时路径，再执行：

```bash
go test ./... -count=1
```
