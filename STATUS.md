# STATUS.md

## 项目状态

- 项目：go-scaffold
- 当前阶段：P1 后续范围已确认
- 总体状态：IN_PROGRESS
- 最后更新：2026-05-25
- 最近 Agent：Codex
- 最近工具：Codex Desktop

## 当前合法工作

- 当前模块：项目优化路线
- 当前任务 ID：TASK-P1-009
- 当前时间切片 ID：TS-P1-009
- 当前状态：NOT_STARTED
- 为什么这是唯一合法下一步：[CONFIRMED] 用户选择选项 A，提升 `BL-021` / `TM-P1-005`，下一步只能执行 `types/*` 契约边界切片；`TASK-NEXT-SCOPE` 已完成范围确认。

## 阶段状态

| 阶段 | 状态 | 证据 |
|---|---|---|
| 项目启动 | COMPLETED | `PROJECT_BRIEF.md` 和 `docs/templates/*` 已中文化并切回项目优化主线 |
| 需求 | COMPLETED | `REQUIREMENTS.md` 已记录确认结果 |
| 高层架构 | COMPLETED | `ARCHITECTURE.md` 已记录确认边界 |
| 路线图 | COMPLETED | `ROADMAP.md` 已生成 |
| 模块边界清单 | COMPLETED | `MODULES.md` 已生成 |
| 测试矩阵与任务拆分 | COMPLETED | `TEST_MATRIX.md` 已生成，`TASKS.md` 和 `TIME_SLICES.md` 已写入 P1 草案 |
| P1 执行顺序确认 | COMPLETED | 用户再次发送“下一步”，按推荐默认顺序确认 |
| 配置 copy/update 测试与修复 | COMPLETED | `internal/config/manager_test.go` 已新增，`copyConfig` 已修复 |
| 配置环境变量策略收拢 | COMPLETED | `DB_*` 成为数据库主环境变量，`REI_APP_DB_*` 保留兼容 fallback；`.env.example` 与实现一致 |
| HTTP health/ready smoke test | COMPLETED | `internal/transport/http/router_test.go` 已覆盖 `/health`、`/ready` missing/failure/ready 路径 |
| demo CRUD 测试基线 | COMPLETED | `internal/modules/demo/service/todo_test.go` 已用临时 SQLite 覆盖 Todo Create/List/Get/Update/Delete |
| demo 迁移边界收拢 | COMPLETED | `DemoMigrationPolicyFor` 固定 server-start/initdb/reload 策略，reload 不再隐式执行 demo `AutoMigrate` |
| CLI tests 命令语义收拢 | COMPLETED | `cmd/server tests` 现在执行 `go test`，并由 `cmd/server/tests_test.go` 固定命令语义 |
| pkg/* API 分类 | COMPLETED | 13 个 `pkg/*` 包已在各自 README、`ARCHITECTURE.md`、`MODULES.md` 中标注 API 定位 |
| pkg/sqlgen unsupported 边界标注 | COMPLETED | unsupported 链式查询、批量删除和 DB reverse 已显式返回 `ErrCodeUnsupportedOperation`，README 已标注部分能力边界 |
| Agent 基础设施补齐 | COMPLETED | `AGENTS.md`、`AGENT_RULES.md`、`SKILLS.md`、项目 skills、reports/specs 和跨工具目录已补齐 |
| Agent 基础设施一致性修复 | COMPLETED | TASK-INFRA-002 已补齐实际缺失的 `AGENTS.md`，规范化 skills、模板和 `.agents` 适配器 |
| 实现 | IN_PROGRESS | 用户已选择提升 `BL-021` / `TM-P1-005`，当前合法下一步为 TASK-P1-009 |
| 验证 | COMPLETED | TASK-P1-008 已运行 `go test ./pkg/sqlgen -count=1` 和 `go test ./... -count=1`，结果 PASS |
| 交接 | COMPLETED | `AGENT_HANDOFF.md` 已记录 TASK-NEXT-SCOPE 范围确认和 TASK-P1-009 下一步 |

## 当前关键发现

| ID | 发现 | 来源 | 状态 |
|---|---|---|---|
| FIND-001 | 部分关键路径仍无测试文件 | `go test ./... -count=1` 输出和 `rg --files -g '*_test.go'` | [RISK] |
| FIND-002 | `.env.example` 与数据库环境变量前缀不一致 | `MODULES.md` BC-001；TASK-P1-002 已修复 | [CONFIRMED] 已处理 |
| FIND-003 | `manager.copyConfig` 未完整复制配置字段 | `MODULES.md` BC-002；TASK-P1-001 已修复 | [CONFIRMED] 已处理 |
| FIND-004 | demo schema 自动迁移触发点需收拢 | `MODULES.md` BC-003；TASK-P1-005 已固定 server-start/initdb/reload 策略 | [CONFIRMED] 已处理 |
| FIND-005 | `cmd/server tests` 命令语义与行为不一致 | `MODULES.md` BC-004；TASK-P1-006 已改为真实 Go test 入口 | [CONFIRMED] 已处理 |
| FIND-010 | `pkg/*` 公共/内部定位未逐包标记 | `ARCHITECTURE.md`、`MODULES.md`；TASK-P1-007 已完成分类 | [CONFIRMED] 已处理 |
| FIND-011 | `pkg/sqlgen` TODO/unsupported 边界不清 | `pkg/sqlgen` README 和源码；TASK-P1-008 已显式返回 unsupported 或文档化 partial 能力 | [CONFIRMED] 已处理 |
| FIND-012 | `types/result`、错误码和跨层类型边界待明确 | 用户选择 A，提升 `BL-021` / `TM-P1-005` | [CONFIRMED] 已提升为 TASK-P1-009 |
| FIND-006 | P1 执行顺序尚未确认 | `TEST_MATRIX.md`、`RISK_REGISTER.md` RISK-009；用户再次发送“下一步” | [CONFIRMED] 已确认 |
| FIND-007 | `AGENTS.md` 被状态文件声明已补齐但实际缺失 | `Test-Path AGENTS.md`、`docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` | [CONFIRMED] 已修复 |
| FIND-008 | `/health`、`/ready` 路由缺少 smoke test | `TEST_MATRIX.md` TM-P0-003；TASK-P1-003 已补测试 | [CONFIRMED] 已处理 |
| FIND-009 | demo Todo CRUD 缺少测试基线 | `TEST_MATRIX.md` TM-P0-005；TASK-P1-004 已补 service/repository 隔离测试 | [CONFIRMED] 已处理 |

## 待用户确认

| ID | 问题 | 影响 | 选项 | Required By |
|---|---|---|---|
| CONFIRM-NEXT-001 | 选择 P1 后续范围或进入收尾 | 已确认：用户选择 A | A: 提升 `BL-021` / `TM-P1-005` 做 `types/*` 契约边界 | COMPLETED |

## 待验证

| ID | 任务 | 需要验证内容 | 命令/方法 |
|---|---|---|---|
|  |  |  |  |

## 需要返工

| ID | 任务 | 原因 | 下一步 |
|---|---|---|---|
|  |  |  |  |

## 最近执行

- 摘要：用户选择 A，完成 TASK-NEXT-SCOPE / TS-NEXT-SCOPE 的范围确认，并将 `BL-021` / `TM-P1-005` 提升为 TASK-P1-009 / TS-P1-009。
- 变更文件：项目状态、任务、时间切片、测试矩阵、Backlog、风险、验收、变更记录和交接文档。
- 执行命令：状态一致性文本检查；`go test ./types/... -count=1`；`go test ./... -count=1`；`git diff --check`。
- 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
- 完成判断：TASK-NEXT-SCOPE 已完成；当前合法下一步为 TASK-P1-009。

## 下一步

- 合法下一步：执行 TASK-P1-009 / TS-P1-009，明确 `types/*` 契约边界。
- 进入条件：开发者发送“下一步”，Agent 读取状态文件后按 TS-P1-009 的允许范围执行。
- 预期输出：`types/*` 契约边界文档化或测试化，运行 `go test ./types/... -count=1` 和 `go test ./... -count=1`，并更新状态文件。
