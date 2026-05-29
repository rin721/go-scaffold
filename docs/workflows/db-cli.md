# DB CLI 工作流

DB CLI 位于 `cmd/main`，通过 `db` 命令调用。它用于生成 SQL 预览、可选 schema 应用和 demo Todo CRUD 检查。

## 命令

```bash
go run ./cmd/main db --operation=schema
go run ./cmd/main db --operation=schema --apply
go run ./cmd/main db --operation=todo-create --title="hello"
go run ./cmd/main db --operation=todo-list
go run ./cmd/main db --operation=todo-get --id=1
go run ./cmd/main db --operation=todo-update --id=1 --title="updated" --done=true
go run ./cmd/main db --operation=todo-delete --id=1
```

使用 `--config` 指定配置文件：

```bash
go run ./cmd/main db --config=configs/config.yaml --operation=schema
```

## 范围

DB CLI 当前聚焦 demo Todo schema 和 CRUD 操作。user schema 应用属于 server 启动装配路径，不在当前 DB CLI 表面内。

## 已移除 InitDB 边界

已移除的 `initdb` 命令和 InitDB 配置段不得在没有新确认任务的情况下恢复。当前 schema 生成应继续通过 `pkg/sqlgen` 和显式 DB CLI 操作完成。

## 验证

修改 DB CLI 行为后，至少运行：

```bash
go test ./cmd/main -count=1
go test ./internal/app/... -count=1
go test ./pkg/sqlgen/... -count=1
```
