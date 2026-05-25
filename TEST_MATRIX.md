# TEST_MATRIX.md

## 测试矩阵状态

- 项目：go-scaffold
- 任务：NONE
- 时间切片：NONE
- 状态：COMPLETED
- 最后更新：2026-05-26
- 原则：本文定义后续优化的验证边界，不代表测试代码已经实现。

## 验证分层

| 层级 | 目标 | 进入代码优化前是否必须 | 说明 |
|---|---|---|---|
| P0 基线 | 保证当前可运行链路不被破坏 | 是 | 每个后续代码切片至少运行相关包测试和 `go test ./... -count=1` |
| P0 新增测试 | 覆盖 app、router、demo、config 关键路径 | 是 | 优先服务边界收拢，不追求一次性全覆盖 |
| P1 边界测试 | 覆盖迁移、CLI、公共包 API | 否 | 按后续任务逐项补齐 |
| P2 质量工程 | CI、性能、发布前验证 | 否 | 进入 Backlog 或后续阶段 |

## 当前基线

| ID | 范围 | 当前证据 | 命令 | 状态 |
|---|---|---|---|---|
| TM-BASE-001 | 全仓库 Go 测试 | 当前通过；`internal/transport/http` 已覆盖 health/ready 与 demo Todo HTTP 集成，`internal/modules/demo/service` 已覆盖 service/repository CRUD，部分 app/reload 路径仍为 `[no test files]` | `go test ./... -count=1` | [CONFIRMED] |
| TM-BASE-002 | 已有包级测试 | `pkg/crypto`、`pkg/database`、`pkg/logger`、`pkg/plugin`、`pkg/sqlgen`、`types/constants` 当前通过 | `go test ./pkg/crypto ./pkg/database ./pkg/logger ./pkg/plugin ./pkg/sqlgen ./types/constants -count=1` | [CONFIRMED] |

## P0 正式测试矩阵

| ID | 范围 | 验证目标 | 建议文件范围 | 验证命令 | 退出条件 | 关联风险 |
|---|---|---|---|---|---|---|
| TM-P0-001 | `internal/config` | 配置加载、`${VAR:default}`、环境变量覆盖、无效配置报错 | `internal/config/*_test.go`、必要 testdata、`.env.example` 只在策略任务中改 | `go test ./internal/config -count=1` | 配置测试存在；失败场景有断言；不会依赖真实外部服务 | BC-001、BC-002 |
| TM-P0-002 | `internal/config.Manager` | `Update`/copy 后不丢 `InitDB`、`Executor`、`Storage`、`CORS`、`Server.Host` 等字段 | `internal/config/*_test.go`、`internal/config/*.go` | `go test ./internal/config -count=1` | 测试能证明字段完整复制；必要修复已完成 | BC-002 |
| TM-P0-003 | `internal/transport/http` | `/health`、`/ready` 在数据库正常/缺失/失败时的 HTTP 状态和响应语义 | `internal/transport/http/*_test.go` | `go test ./internal/transport/http -count=1` | [CONFIRMED] TASK-P1-003 已用 `httptest` 覆盖；不启动真实 server | BC-006 |
| TM-P0-004 | `internal/app` | `app.New` 在 server/initdb 模式的最小装配链路 | `internal/app/**/*_test.go` | `go test ./internal/app/... -count=1` | 使用临时配置；不依赖真实外部服务；资源可关闭 | BC-003、BC-006 |
| TM-P0-005 | `internal/modules/demo` | demo Todo Create/List/Get/Update/Delete 关键路径 | `internal/modules/demo/**/*_test.go`、`internal/transport/http/*_test.go` | `go test ./internal/modules/demo/... ./internal/transport/http -count=1` | [CONFIRMED] TASK-P1-004 已使用临时 SQLite 覆盖 service/repository 关键行为；TASK-P1-015 已覆盖 handler/router HTTP 集成路径 | BC-003、BC-006 |
| TM-P0-006 | 全仓库回归 | 每个代码切片完成后确认无全局回归 | 不限制，只读验证 | `go test ./... -count=1` | 全量测试 PASS；新增失败进入修复流程 | FIND-001 |

## P1 正式测试矩阵

| ID | 范围 | 验证目标 | 建议文件范围 | 验证命令 | 退出条件 | 关联风险 |
|---|---|---|---|---|---|---|
| TM-P1-001 | 迁移策略 | demo `AutoMigrate`、`initdb`、reload 触发边界清晰 | `internal/app/**/*_test.go`、迁移边界文档 | `go test ./internal/app/... -count=1` | [CONFIRMED] TASK-P1-005 已写清 dev/demo 与生产/bootstrap 职责，并验证触发策略 | BC-003 |
| TM-P1-002 | `cmd/server` | `tests` 命令语义与实际行为一致 | `cmd/server/*_test.go`、`cmd/server/*.go` | `go test ./cmd/server -count=1` | [CONFIRMED] TASK-P1-006 已改为真实 Go test 入口并补测试 | BC-004 |
| TM-P1-003 | `pkg/*` API 分类与后续测试缺口 | 先明确公共基础设施 API、公共工具 API、内部支撑工具包边界，再按后续任务补行为测试 | `ARCHITECTURE.md`、`MODULES.md`、`pkg/*/README.md`，后续测试任务再触碰 `pkg/*/*_test.go` | TASK-P1-007：`go test ./... -count=1`；后续行为测试按包执行 | [CONFIRMED] TASK-P1-007 已完成分类；行为测试补齐仍按后续任务或 Backlog 处理 | RISK-004、RISK-008、FIND-001 |
| TM-P1-004 | `pkg/sqlgen` | 未实现能力显式返回 unsupported 或文档化 | `pkg/sqlgen/*`、包 README | `go test ./pkg/sqlgen -count=1` | TODO 能力不再暗示已支持；测试或文档覆盖 unsupported | BC-005 |
| TM-P1-005 | `types/*` | `types/result`、错误码、跨层契约边界清晰 | `types/**/*_test.go`、`ARCHITECTURE.md` 或包 README | `go test ./types/... -count=1` | [CONFIRMED] TASK-P1-009 已标注 HTTP 契约与纯类型边界，并补最小测试 | BL-021 |
| TM-P1-006 | `pkg/plugin` | 插件注册责任为被动 registry/runtime，不由 `pkg/plugin` 主动注册插件服务 | `pkg/plugin/*`、包 README、架构/决策文档 | `go test ./pkg/plugin -count=1` | [CONFIRMED] 插件服务或宿主装配层显式创建并 `Register` local/http 插件；manager 公共 API 不再暴露主动配置加载/服务注册入口 | BL-022、RISK-015 |
| TM-P1-007 | 首批无外部依赖 `pkg/*` 行为测试 | 为 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 补最小包级行为测试，优先覆盖稳定成功路径和明确错误路径 | `pkg/cli/**/*_test.go`、`pkg/i18n/**/*_test.go`、`pkg/yaml2go/**/*_test.go`；必要时限当前包实现文件 | `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 均已有确定性包级行为测试；不依赖外部服务或生产配置 | BL-020、RISK-008 |
| TM-P1-008 | 第二批无外部服务 `pkg/*` 行为测试 | 为 `pkg/executor`、`pkg/httpserver`、`pkg/storage` 补最小包级行为测试，覆盖稳定成功路径和明确错误路径 | `pkg/executor/**/*_test.go`、`pkg/httpserver/**/*_test.go`、`pkg/storage/**/*_test.go`；必要时限当前三包实现文件 | `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`；`go test ./... -count=1` | [CONFIRMED] 三包已有确定性包级行为测试；不依赖 Redis、数据库、第三方网络服务或生产配置 | BL-020、RISK-008 |
| TM-P1-009 | 第三批 `pkg/cache` 隔离行为测试 | 为 `pkg/cache` 补最小包级行为测试，用进程内 Redis 覆盖成功路径和明确错误路径 | `pkg/cache/**/*_test.go`；必要时限当前包实现文件；测试依赖可修改 `go.mod`、`go.sum` | `go test ./pkg/cache -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/cache` 已有确定性隔离行为测试；不依赖真实 Redis、数据库、第三方网络服务或生产配置 | BL-020、RISK-008 |
| TM-P1-010 | `pkg/utils` 内部支撑测试 | 为 `pkg/utils` 补最小确定性行为测试，覆盖 Snowflake、地址校验、端口查找、设备 ID 和 i18n helper 委托 | `pkg/utils/**/*_test.go`；必要时限当前包实现文件 | `go test ./pkg/utils -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/utils` 已有最小确定性行为测试；不依赖真实外部网络服务、固定生产端口、数据库或生产配置 | BL-023、RISK-008 |
| TM-P1-011 | router/middleware/demo HTTP 集成测试 | 用 `httptest` 覆盖 demo Todo HTTP 路由、handler/service/repository 集成，以及 TraceID、CORS、Recovery 中间件链路 | `internal/transport/http/**/*_test.go`、必要时 `internal/middleware/**/*_test.go` 或 `internal/modules/demo/**/*_test.go` | `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`；`go test ./... -count=1` | [CONFIRMED] TASK-P1-015 已覆盖 demo Todo HTTP CRUD、CORS preflight/actual、TraceID round-trip 和 Recovery trace 响应；不启动真实 HTTP server | BL-002、RISK-008 |

## P1 优化任务草案

| 任务 ID | 标题 | 优先级 | 类型 | 允许文件范围 | 验证命令 | 退出条件 | 状态 |
|---|---|---|---|---|---|---|---|
| TASK-OPT-005 | 确认测试矩阵和 P1 执行顺序 | P0 | 确认 | 根文档和状态文件 | 不需要 Go 测试；确认后下一代码切片需测试 | 用户确认或发送“下一步”接受推荐顺序 | COMPLETED |
| TASK-P1-001 | 修复 `copyConfig` 字段覆盖并补配置测试 | P1 | 测试+修复 | `internal/config/*`、必要 testdata、状态文档 | `go test ./internal/config -count=1`；`go test ./... -count=1` | 配置复制不丢字段；测试覆盖关键字段 | COMPLETED |
| TASK-P1-002 | 统一配置环境变量策略 | P1 | 测试+修复+文档 | `internal/config/*`、`.env.example`、配置文档、状态文档 | `go test ./internal/config -count=1`；`go test ./... -count=1` | `.env.example` 与实现一致；前缀策略被记录 | COMPLETED |
| TASK-P1-003 | 增加 health/ready 与 router smoke test | P1 | 测试 | `internal/transport/http/*_test.go`、状态文档 | `go test ./internal/transport/http -count=1`；`go test ./... -count=1` | `/health`、`/ready` 行为被测试固定 | COMPLETED |
| TASK-P1-004 | 增加 demo CRUD 测试基线 | P1 | 测试 | `internal/modules/demo/**/*_test.go`、状态文档 | `go test ./internal/modules/demo/... -count=1`；`go test ./... -count=1` | Todo 关键 CRUD 路径被临时 SQLite 测试覆盖 | COMPLETED |
| TASK-P1-005 | 明确 demo 迁移边界 | P1 | 架构+测试+小修 | `internal/app/**/*`、迁移边界文档、状态文档 | `go test ./internal/app/... -count=1`；`go test ./... -count=1` | server/initdb/reload 迁移职责可解释、可验证 | COMPLETED |
| TASK-P1-006 | 收拢 `cmd/server tests` 命令语义 | P1 | CLI 小修+测试 | `cmd/server/*`、CLI 文档、状态文档 | `go test ./cmd/server -count=1`；`go test ./... -count=1` | 命令名、描述或行为与真实用途一致 | COMPLETED |
| TASK-P1-007 | 完成 `pkg/*` 公共/内部分类 | P1 | 文档 | `ARCHITECTURE.md`、`MODULES.md`、包 README、状态文档 | `go test ./... -count=1` | [CONFIRMED] 每个 `pkg/*` 包定位已标注；破坏性重构仍需单独确认 | COMPLETED |
| TASK-P1-008 | 标注 `pkg/sqlgen` 未实现能力 | P1 | 文档+测试或小修 | `pkg/sqlgen/*`、包 README、状态文档 | `go test ./pkg/sqlgen -count=1`；`go test ./... -count=1` | [CONFIRMED] TODO/unsupported 边界不再误导使用者 | COMPLETED |
| TASK-P1-009 | 明确 `types/*` 契约边界 | P1 | 文档+测试或小修 | `types/**/*`、`ARCHITECTURE.md`、`MODULES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`docs/specs/types_contract_boundary.md`、状态文档 | `go test ./types/... -count=1`；`go test ./... -count=1` | [CONFIRMED] `types/result` HTTP 契约、错误码预留和跨层类型边界已标注 | COMPLETED |
| TASK-P1-010 | 收拢 `pkg/plugin` 被动注册边界 | P1 | API 小修+测试+文档 | `pkg/plugin/*`、包 README、架构/决策/状态文档 | `go test ./pkg/plugin -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/plugin` 不主动注册插件服务，local/http 插件由插件服务显式注册 | COMPLETED |
| TASK-P1-011 | 补首批无外部服务依赖 `pkg/*` 行为测试 | P1 | 测试 | `pkg/cli/**/*_test.go`、`pkg/i18n/**/*_test.go`、`pkg/yaml2go/**/*_test.go`、必要时限当前包实现文件、状态文档 | `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 均有最小行为测试且不依赖外部服务 | COMPLETED |
| TASK-P1-012 | 补第二批 `pkg/*` 行为测试 | P1 | 测试 | `pkg/executor/**/*_test.go`、`pkg/httpserver/**/*_test.go`、`pkg/storage/**/*_test.go`、必要时限当前三包实现文件、状态文档 | `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/executor`、`pkg/httpserver`、`pkg/storage` 均有最小行为测试且不依赖外部服务 | COMPLETED |
| TASK-P1-013 | 补第三批 `pkg/cache` 隔离行为测试 | P1 | 测试 | `pkg/cache/**/*_test.go`、必要时限当前包实现文件、测试依赖、状态文档 | `go test ./pkg/cache -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/cache` 有最小隔离行为测试且不依赖真实 Redis | COMPLETED |
| TASK-P1-014 | 补 `pkg/utils` 内部支撑工具最小行为测试 | P1 | 测试 | `pkg/utils/**/*_test.go`、必要时限当前包实现文件、状态文档 | `go test ./pkg/utils -count=1`；`go test ./... -count=1` | [CONFIRMED] `pkg/utils` 有最小确定性行为测试且不依赖真实外部服务 | COMPLETED |
| TASK-P1-015 | 补 app/router/middleware 最小集成测试 | P1 | 测试 | `internal/transport/http/**/*_test.go`、必要时 `internal/middleware/**/*_test.go`、`internal/modules/demo/**/*_test.go`、状态文档 | `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`；`go test ./... -count=1` | [CONFIRMED] demo Todo HTTP 集成和 TraceID/CORS/Recovery 链路有最小测试覆盖 | COMPLETED |

## 推荐执行顺序

1. TASK-OPT-005：确认测试矩阵和 P1 执行顺序。
2. TASK-P1-001：先处理 `copyConfig`，因为它影响配置热更新正确性。
3. TASK-P1-002：统一环境变量策略，避免配置文档与实现继续漂移。
4. TASK-P1-003：固定 HTTP health/ready 行为。
5. TASK-P1-004：固定 demo CRUD 示例行为。
6. TASK-P1-005：在测试基线后收拢迁移边界。
7. TASK-P1-006：处理 CLI 语义不一致。
8. TASK-P1-007：完成包 API 分类。
9. TASK-P1-008：标注 `pkg/sqlgen` 未实现能力。
10. TASK-P1-011：补首批无外部依赖 `pkg/*` 行为测试。
11. TASK-P1-012：补第二批 `pkg/*` 行为测试。
12. TASK-P1-013：补第三批 `pkg/cache` 隔离行为测试。
13. TASK-P1-014：补 `pkg/utils` 内部支撑工具最小行为测试。
14. TASK-P1-015：补 router/middleware/demo HTTP 集成测试。

当前合法下一项：

- [COMPLETED] 本轮测试矩阵执行与 Phase 6 收尾完成；当前无自动下一实现任务。
- [DEFERRED] app 装配、reload/config 等剩余集成测试仍可后续提升，但必须由用户重新确认并拆成新的任务/时间切片。

## 验收门禁

- [CONFIRMED] 后续任何代码切片必须先声明关联的测试矩阵 ID。
- [CONFIRMED] 后续任何代码切片必须记录允许文件范围、验证命令和退出条件。
- [CONFIRMED] 代码切片完成后必须运行相关包测试和 `go test ./... -count=1`。
- [CONFIRMED] 未执行验证的代码任务不得标记为 `COMPLETED`。
- [CONFIRMED] 如果新增测试暴露既有缺陷，必须在当前任务内修复，或记录为 `REWORK_REQUIRED`/`BLOCKED`，不得静默跳过。

## 非目标

- [CONFIRMED] 本文不要求一次性实现所有测试代码；具体测试按 P1 时间切片逐项落地。
- [CONFIRMED] 本文不修改 Go 代码、配置结构、数据库结构或 HTTP 路由。
- [CONFIRMED] 本文不提升 auth/rbac、插件 rpc/ws/discovery、CI/CD 或部署任务。
