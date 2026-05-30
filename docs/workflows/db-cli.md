# DB CLI 工作流

`db` 命令提供基于 sqlgen 的数据库 DDL 和 Demo Todo 操作。

## 示例

```bash
go run ./cmd/main db --operation=schema
go run ./cmd/main db --operation=schema --apply
go run ./cmd/main db --operation=todo-create --title="交付文档"
go run ./cmd/main db --operation=todo-list
```

## 范围

该命令当前聚焦 Demo Todo 表结构和 CRUD 操作，不管理身份认证、访问控制或账号数据。

## 维护提示

- CLI 解析保持在 `cmd/main`。
- SQL 生成和执行行为保持在 `internal/app/dbapp`。
- flag 行为的测试放在命令层附近，SQL 行为的测试放在 `dbapp` 附近。
