# Demo Migration Boundary

## 状态

- 任务：TASK-P1-005
- 时间切片：TS-P1-005
- 最后更新：2026-05-25
- 状态：COMPLETED

## 目标

- [CONFIRMED] demo 模块是长期标准示例，可以在开发和本地示例路径中自动准备自身 schema。
- [CONFIRMED] 生产/bootstrap 迁移仍采用 dev-prod 分层策略，后续应倾向显式 SQL 或独立迁移流程。
- [CONFIRMED] reload 是运行期配置更新路径，不应隐式执行 demo schema 迁移。

## 触发策略

| 触发点 | 入口 | 是否执行 demo AutoMigrate | 原因 | 状态 |
|---|---|---|---|---|
| server-start | `initapp.NewModules` | 是 | 保持本地 demo server 启动即可使用 Todo 示例 | [CONFIRMED] |
| initdb | `modeapp.BuildInitDB` | 是 | `initdb` 是显式 demo schema bootstrap 命令 | [CONFIRMED] |
| reload | `reloadapp.reloadDatabase` | 否 | reload 不应成为运行期隐式 schema 变更点 | [CONFIRMED] |

## 非目标

- [CONFIRMED] 本切片不引入生产迁移框架。
- [CONFIRMED] 本切片不修改数据库 schema、配置结构或依赖版本。
- [CONFIRMED] 本切片不改变 demo Todo model 或 repository 行为。

## 验证

- `DemoMigrationPolicyFor` 固定 server-start、initdb、reload 三类触发点策略。
- `MigrateDemoSchemaForTrigger` 使用隔离 SQLite 验证 server-start/initdb 会创建 Todo 表，reload 不创建 Todo 表。
- `go test ./internal/app/... -count=1` 与 `go test ./... -count=1` 必须通过。
