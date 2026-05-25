# TIME_SLICES.md

## 当前合法时间切片

- Time Slice ID：TS-NEXT-SCOPE-007
- Task ID：TASK-NEXT-SCOPE-007
- Status：PENDING_USER_CONFIRMATION
- Summary：TASK-P1-014 已完成，等待用户确认下一范围。

## 时间切片列表

### TS-OPT-001：生成中文项目优化启动材料

- Status：COMPLETED
- Task ID：TASK-OPT-001
- Verification：
  - `go test ./... -count=1`：PASS

### TS-OPT-002：确认项目优化路线和边界

- Status：COMPLETED
- Task ID：TASK-OPT-002
- Confirmed：
  - 治理优先。
  - `pkg/*` 混合策略。
  - demo 长期标准示例。
  - 迁移 dev-prod 分层。
  - 中文化根文档和模板优先。
  - auth/JWT 延后处理。

### TS-OPT-003：生成模块边界清单

- Status：COMPLETED
- Task ID：TASK-OPT-003
- Files Changed：
  - `MODULES.md`
  - `REQUIREMENTS.md`
  - `ACCEPTANCE.md`
  - `ROADMAP.md`
  - `BACKLOG.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `STATUS.md`
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `AGENT_HANDOFF.md`
- Verification：
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - 模块边界清单存在。
  - 测试矩阵草案存在。
  - P1 优化候选项进入 Backlog 或任务草案。

### TS-OPT-004：生成正式测试矩阵和任务拆分草案

- Status：COMPLETED
- Task ID：TASK-OPT-004
- Scope：
  - 将 `MODULES.md` 的测试矩阵草案整理为正式 `TEST_MATRIX.md`。
  - 将 P1 优化候选项拆成任务草案。
  - 为每个任务定义允许文件范围、验证命令和退出条件。
  - 更新状态、验收和交接文件。
- Allowed Files：
  - 项目文档和状态文件。
  - 可新增 `TEST_MATRIX.md`。
- Forbidden：
  - 不修改 Go 代码。
  - 不新增 Go 测试文件。
  - 不实现 Backlog 项。
- Verification：
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - 正式测试矩阵存在。
  - P1 任务和时间切片草案存在。
  - 下一合法任务明确。

### TS-OPT-005：确认正式测试矩阵和 P1 执行顺序

- Status：COMPLETED
- Task ID：TASK-OPT-005
- Scope：
  - 确认 `TEST_MATRIX.md`。
  - 确认 P1 任务顺序。
  - 确认是否进入首个代码切片。
- Recommended Default：
  - 用户再次发送“下一步”时，按推荐顺序进入 TS-P1-001。
- Allowed Files：
  - 项目文档和状态文件。
- Forbidden：
  - 不修改 Go 代码。
  - 不新增 Go 测试文件。
- Exit Conditions：
  - 用户确认，或用“下一步”接受推荐默认顺序。
  - 当前合法切片推进到 TS-P1-001。
- Completion Evidence：
  - 用户再次发送“下一步”，按推荐默认顺序确认。
  - 已进入并完成 TS-P1-001。

### TS-INFRA-001：补齐 Prompt 全量 Agent 基础设施

- Status：COMPLETED
- Task ID：TASK-INFRA-001
- Scope：
  - 新增跨 Agent 入口与规则文件。
  - 新增缺失模板。
  - 新增 reports/specs、跨工具目录和项目 skills。
  - 更新状态、验收、路线图、测试报告、变更记录和交接说明。
- Allowed Files：
  - `AGENTS.md`
  - `CLAUDE.md`
  - `AGENT_RULES.md`
  - `SKILLS.md`
  - `docs/templates/*`
  - `docs/reports/*`
  - `docs/specs/*`
  - `skills/*/SKILL.md`
  - `.agents/*`
  - `.cursor/*`
  - `.kiro/*`
  - `.codex/*`
  - 项目状态文档
- Forbidden：
  - 不修改 Go 代码。
  - 不执行 TASK-P1-002。
- Verification：
  - Prompt 全量产物存在性核对：PASS。
  - `go test ./... -count=1`：PASS。
- Exit Conditions：
  - Prompt 全量产物存在。
  - 当前合法下一步恢复到 TS-P1-002。

### TS-INFRA-002：Agent 基础设施一致性修复

- Status：COMPLETED
- Task ID：TASK-INFRA-002
- Scope：
  - 新增缺失的 `AGENTS.md`。
  - 统一跨工具入口引用。
  - 补齐 14 个 canonical skills 的 frontmatter 和完整结构。
  - 新增 14 个 `.agents` 轻量适配器。
  - 标准化 `docs/templates/*` 为可复用模板。
  - 生成状态诊断报告并更新状态、测试、变更和交接文件。
- Allowed Files：
  - `AGENTS.md`
  - `CLAUDE.md`
  - `AGENT_RULES.md`
  - `SKILLS.md`
  - `docs/templates/*`
  - `docs/reports/status_diagnostics/*`
  - `skills/*/SKILL.md`
  - `.agents/skills/*/SKILL.md`
  - `.cursor/*`
  - `.kiro/*`
  - `.codex/*`
  - 项目状态文档
- Forbidden：
  - 不修改 Go 代码。
  - 不执行 TASK-P1-002。
  - 不安装依赖、不提交 git、不部署。
- Verification：
  - Agent 基础设施文件存在性核对：PASS。
  - `quick_validate.py` 验证 canonical 和 adapter skills：PASS。
  - 跨工具入口引用一致性检查：PASS。
  - `go test ./... -count=1`：PASS。
- Exit Conditions：
  - `AGENTS.md` 存在。
  - 所有 required skills 可验证。
  - 状态冲突已记录并关闭。
  - 当前合法下一步恢复到 TS-P1-002。

## P1 时间切片草案

### TS-P1-001：配置 copy/update 测试与修复

- Status：COMPLETED
- Task ID：TASK-P1-001
- Matrix：TM-P0-001、TM-P0-002、TM-P0-006
- Allowed Files：
  - `internal/config/*`
  - 必要 testdata
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - 配置复制关键字段完整。
  - 测试和状态记录完成。
- Files Changed：
  - `internal/config/manager.go`
  - `internal/config/manager_test.go`
  - 项目状态文档

### TS-P1-002：配置环境变量策略收拢

- Status：COMPLETED
- Task ID：TASK-P1-002
- Matrix：TM-P0-001、TM-P0-006
- Allowed Files：
  - `internal/config/*`
  - `.env.example`
  - 配置相关文档
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - 环境变量策略一致或差异被文档化。
  - `.env.example` 与实现一致。
- Files Changed：
  - `.env.example`
  - `internal/config/app_database.go`
  - `internal/config/constants.go`
  - `internal/config/manager_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS

### TS-P1-003：HTTP health/ready smoke test

- Status：COMPLETED
- Task ID：TASK-P1-003
- Matrix：TM-P0-003、TM-P0-006
- Allowed Files：
  - `internal/transport/http/*_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/transport/http -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - health/ready 状态码和响应语义被测试固定。
- Files Changed：
  - `internal/transport/http/router_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/transport/http -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - `/health` HTTP 200 和成功响应语义已固定。
  - `/ready` 数据库缺失、ping 失败、ping 成功路径已固定。
  - 未启动真实 HTTP server。

### TS-P1-004：demo CRUD 测试基线

- Status：COMPLETED
- Task ID：TASK-P1-004
- Matrix：TM-P0-005、TM-P0-006
- Allowed Files：
  - `internal/modules/demo/**/*_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/modules/demo/... -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - demo Todo CRUD 关键路径有隔离测试。
- Files Changed：
  - `internal/modules/demo/service/todo_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/modules/demo/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - 使用临时 SQLite，不依赖真实外部服务。
  - Todo Create/List/Get/Update/Delete 成功路径已固定。
  - 空标题校验、not found、软删除后不可见语义已固定。

### TS-P1-005：demo 迁移边界收拢

- Status：COMPLETED
- Task ID：TASK-P1-005
- Matrix：TM-P1-001、TM-P0-006
- Allowed Files：
  - `internal/app/**/*`
  - 迁移边界文档
  - 项目状态文档
- Verification：
  - `go test ./internal/app/... -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] demo 自动迁移触发策略清晰且可验证。
- Files Changed：
  - `internal/app/initapp/modules.go`
  - `internal/app/modeapp/mode.go`
  - `internal/app/reloadapp/reload.go`
  - `internal/app/initapp/demo_migration_test.go`
  - `docs/specs/demo_migration_boundary.md`
  - 项目状态文档
- Verification：
  - `go test ./internal/app/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - server-start 和 initdb 仍显式执行 demo `AutoMigrate`。
  - reload 策略改为跳过 demo `AutoMigrate`，避免运行期隐式 schema 变更。
  - 迁移边界文档已记录 dev/demo 与生产/bootstrap 职责。

### TS-P1-006：CLI tests 命令语义收拢

- Status：COMPLETED
- Task ID：TASK-P1-006
- Matrix：TM-P1-002、TM-P0-006
- Allowed Files：
  - `cmd/server/*`
  - CLI 相关文档
  - 项目状态文档
- Verification：
  - `go test ./cmd/server -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] `tests` 命令不再误导使用者。
- Files Changed：
  - `cmd/server/tests.go`
  - `cmd/server/main.go`
  - `cmd/server/tests_test.go`
  - `docs/specs/cli_tests_command_boundary.md`
  - 项目状态文档
- Verification：
  - `go test ./cmd/server -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - `tests` 命令现在运行 Go tests，而不是 yaml2go 示例转换。
  - 命令支持 `--package/-p` 指定测试范围，默认 `./...`。
  - 命令元信息和 runner 行为已有最小单元测试。

### TS-P1-007：`pkg/*` API 分类

- Status：COMPLETED
- Task ID：TASK-P1-007
- Matrix：TM-P1-003、TM-P0-006
- Allowed Files：
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - 包 README
  - 项目状态文档
- Verification：
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] 每个 `pkg/*` 包定位明确。
- Files Changed：
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - `pkg/cache/README.md`
  - `pkg/cli/README.md`
  - `pkg/crypto/README.md`
  - `pkg/database/README.md`
  - `pkg/executor/README.md`
  - `pkg/httpserver/README.md`
  - `pkg/i18n/README.md`
  - `pkg/logger/README.md`
  - `pkg/plugin/README.md`
  - `pkg/sqlgen/README.md`
  - `pkg/storage/README.md`
  - `pkg/utils/README.md`
  - `pkg/yaml2go/README.md`
  - 项目状态文档
- Verification：
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - 公共基础设施 API、公共工具 API、内部支撑工具包分类已落到包 README。
  - 破坏性 `pkg/*` 重构仍需单独任务确认。

### TS-P1-008：`pkg/sqlgen` unsupported 边界

- Status：COMPLETED
- Task ID：TASK-P1-008
- Matrix：TM-P1-004、TM-P0-006
- Allowed Files：
  - `pkg/sqlgen/*`
  - 包 README
  - 项目状态文档
- Verification：
  - `go test ./pkg/sqlgen -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] 未实现能力边界清晰。
- Files Changed：
  - `pkg/sqlgen/errors.go`
  - `pkg/sqlgen/types.go`
  - `pkg/sqlgen/generator.go`
  - `pkg/sqlgen/query.go`
  - `pkg/sqlgen/update.go`
  - `pkg/sqlgen/delete.go`
  - `pkg/sqlgen/reverse.go`
  - `pkg/sqlgen/doc.go`
  - `pkg/sqlgen/sqlgen_test.go`
  - `pkg/sqlgen/README.md`
  - 项目状态文档
- Verification：
  - `gofmt -w pkg/sqlgen/errors.go pkg/sqlgen/types.go pkg/sqlgen/generator.go pkg/sqlgen/query.go pkg/sqlgen/update.go pkg/sqlgen/delete.go pkg/sqlgen/reverse.go pkg/sqlgen/doc.go pkg/sqlgen/sqlgen_test.go`：PASS
  - `go test ./pkg/sqlgen -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - 高级查询和部分未实现能力已显式返回 `ErrCodeUnsupportedOperation`。
  - README 已标注 unsupported / partial 能力边界。
  - 当前 P1 已拆分任务列表全部完成。

### TS-P1-009：`types/*` 契约边界

- Status：COMPLETED
- Task ID：TASK-P1-009
- Matrix：TM-P1-005、TM-P0-006
- Purpose：明确 `types/*` 作为跨层契约、HTTP 响应契约或聚合入口的边界，避免 `types/result` 的 Gin 依赖被误解为纯领域类型。
- Inputs：
  - `BL-021`
  - `TM-P1-005`
  - `MODULES.md` 的 `types/*` 边界说明
  - `ARCHITECTURE.md` 的目标依赖方向
- Outputs：
  - `types/*` 契约边界文档或包级说明
  - 必要的 `types/**/*_test.go`
  - `docs/specs/types_contract_boundary.md`
  - 状态、验收、测试报告、变更日志和交接更新
- Allowed Files：
  - `types/**/*`
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - `TEST_MATRIX.md`
  - `ACCEPTANCE.md`
  - `docs/specs/types_contract_boundary.md`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `pkg/**/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不实现 auth/rbac。
  - 不修改 HTTP router、middleware 或 demo handler 行为。
  - 不移动 `types/*` 包或做破坏性 API 重构。
  - 不补 `pkg/*` 行为测试；该方向仍留在 `BL-020`。
- Execution Steps：
  1. 阅读 `types/*` 源码和现有调用点，确认 `types/result`、`types/errors`、`types/constants` 和根 `types` 的实际使用边界。
  2. 编写或更新契约边界文档，明确公共契约、HTTP/Gin 契约、预留错误码和非目标。
  3. 如存在可稳定验证的行为，新增最小 `types` 包测试。
  4. 运行 `go test ./types/... -count=1` 和 `go test ./... -count=1`。
  5. 更新状态、验收、测试报告、风险、Backlog、变更日志和交接说明。
- Verification：
  - `go test ./types/... -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] HTTP 契约与纯类型边界被标注。
  - [CONFIRMED] auth/rbac 预留错误码不再暗示当前已实现 auth/rbac。
  - [CONFIRMED] 相关测试通过，或未新增测试的原因被记录并不影响本切片验收。
  - [CONFIRMED] 所有项目状态文件已更新，下一合法任务明确。
- Files Changed：
  - `types/doc.go`
  - `types/constants/doc.go`
  - `types/errors/doc.go`
  - `types/result/result.go`
  - `types/result/result_test.go`
  - `types/errors/error_test.go`
  - `docs/specs/types_contract_boundary.md`
  - 项目状态文档
- Verification：
  - `gofmt -w types/doc.go types/constants/doc.go types/errors/doc.go types/result/result.go types/result/result_test.go types/errors/error_test.go`：PASS
  - `go test ./types/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - HTTP/Gin 契约、错误码预留、跨层常量和根聚合入口边界已文档化。
  - `types/result` 和 `types/errors` 已补最小契约测试。

### TS-NEXT-SCOPE：确认下一阶段范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE
- Matrix：BL-021、TM-P1-005
- Allowed Files：
  - 项目状态文档
- Verification：
  - 不需要 Go 测试，除非确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户选择 A：提升 `BL-021` / `TM-P1-005`。
  - [CONFIRMED] 新的唯一合法任务和时间切片已写入状态文件。

### TS-NEXT-SCOPE-002：确认 `types/*` 契约边界后的后续范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-002
- Matrix：BL-022、TM-P1-006
- Allowed Files：
  - 项目状态文档
- Verification：
  - 不需要 Go 测试，除非确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户修正并选择 `pkg/plugin` 被动注册边界作为后续范围。
  - [CONFIRMED] 新的唯一合法任务和时间切片已写入状态文件。

### TS-P1-010：`pkg/plugin` 被动注册边界

- Status：COMPLETED
- Task ID：TASK-P1-010
- Matrix：TM-P1-006
- Purpose：让 `pkg/plugin` 保持被动 registry/runtime，插件服务或宿主装配层显式创建 local/http 插件并调用 `Register` 注册。
- Inputs：
  - 用户架构修正
  - `DECISIONS.md` 中历史插件决策
  - `pkg/plugin` 当前 Manager、local/http 和 README
- Outputs：
  - `pkg/plugin` 被动注册 API 与测试
  - README、架构、模块清单、测试矩阵和状态文档更新
- Allowed Files：
  - `pkg/plugin/**/*`
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - `TEST_MATRIX.md`
  - `ACCEPTANCE.md`
  - `DECISIONS.md`
  - `BACKLOG.md`
  - `RISK_REGISTER.md`
  - `docs/reports/status_diagnostics/*`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不实现 rpc/ws/discovery。
  - 不接入 `internal/app`。
  - 不新增真实插件服务示例。
  - 不修改生产配置、数据库 schema 或依赖。
- Execution Steps：
  1. 移除或收窄 `Manager` 的主动配置加载/本地 factory 装配公共面。
  2. 暴露 local/http 插件由服务侧构造后注册的路径。
  3. 更新 `pkg/plugin` 测试，证明注册由服务侧显式调用 `Register` 完成。
  4. 运行 `go test ./pkg/plugin -count=1`、`go test ./... -count=1` 和 `git diff --check`。
  5. 更新状态、测试报告、变更日志和交接说明。
- Verification：
  - `go test ./pkg/plugin -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] `pkg/plugin` 不再主动从配置加载并注册插件服务。
  - [CONFIRMED] local/http 插件均可由插件服务或宿主装配层显式注册。
  - [CONFIRMED] 文档记录该边界和兼容风险。
  - [CONFIRMED] 相关验证通过。
- Files Changed：
  - `pkg/plugin/manager.go`
  - `pkg/plugin/http.go`
  - `pkg/plugin/local.go`
  - `pkg/plugin/errors.go`
  - `pkg/plugin/config.go`
  - `pkg/plugin/constants.go`
  - `pkg/plugin/doc.go`
  - `pkg/plugin/plugin_test.go`
  - `pkg/plugin/README.md`
  - 项目状态文档
- Verification：
  - `gofmt -w pkg/plugin/manager.go pkg/plugin/http.go pkg/plugin/constants.go pkg/plugin/errors.go pkg/plugin/doc.go pkg/plugin/config.go pkg/plugin/local.go pkg/plugin/plugin_test.go`：PASS
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- Completion Evidence：
  - manager 公共 API 不再包含 `Load` 或 local factory 注册入口。
  - `NewHTTP` 提供 HTTP 插件显式构造路径。
  - local/http 测试均通过服务侧构造插件后 `Register` 的路径验证。

### TS-NEXT-SCOPE-003：确认 `pkg/plugin` 被动注册边界后的后续范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-003
- Matrix：BL-020、TM-P1-003、TM-P1-007
- Allowed Files：
  - 项目状态文档
- Verification：
  - 不需要 Go 测试，除非确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户选择 A：提升 `BL-020` 补 `pkg/*` 行为测试。
  - [CONFIRMED] 首批任务 TASK-P1-011 / TS-P1-011 已写入状态文件。
  - [CONFIRMED] 当前合法时间切片不再是待确认状态。

### TS-P1-011：首批 `pkg/*` 行为测试

- Status：COMPLETED
- Task ID：TASK-P1-011
- Matrix：TM-P1-003、TM-P1-007、TM-P0-006
- Purpose：为无外部服务依赖且当前无包级测试的 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 补最小行为测试。
- Inputs：
  - `BL-020`
  - `ARCHITECTURE.md` 的 `pkg/*` API 分类
  - `MODULES.md` 的 `pkg/*` 测试缺口
  - `TEST_MATRIX.md` 的 `TM-P1-003` 和 `TM-P1-007`
- Outputs：
  - `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 测试文件
  - 状态、验收、测试报告、变更日志和交接更新
- Allowed Files：
  - `pkg/cli/**/*_test.go`
  - `pkg/i18n/**/*_test.go`
  - `pkg/yaml2go/**/*_test.go`
  - 必要时仅限当前三个包内的实现文件，用于修复新增测试暴露的当前范围缺陷
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他 `pkg/*` 包
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不补 `pkg/cache`、`pkg/executor`、`pkg/httpserver`、`pkg/storage` 测试；这些进入后续批次。
  - 不接入真实 Redis、数据库、HTTP server 或外部网络。
  - 不改变公共 API。
  - 不进行包结构重构或中文化 README。
- Execution Steps：
  1. 阅读三个包的源码和 README，确认稳定公共行为。
  2. 设计每包 1 到 3 个最小行为测试，优先覆盖成功路径和明确错误路径。
  3. 新增测试文件并运行 gofmt。
  4. 运行相关包测试、全量回归和 `git diff --check`。
  5. 根据验证结果更新状态、问题、风险、测试报告、变更日志和交接。
- Test Commands：
  - `gofmt -w pkg/cli/*_test.go pkg/i18n/*_test.go pkg/yaml2go/*_test.go`
  - `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Verification Method：
  - 包级 Go 测试、全量回归、diff 检查。
- Acceptance Criteria：
  - `pkg/cli` 有最小包级行为测试。
  - `pkg/i18n` 有最小包级行为测试。
  - `pkg/yaml2go` 有最小包级行为测试。
  - 测试不依赖外部服务或生产配置。
  - 验证命令通过，或失败被记录为 `REWORK_REQUIRED` / `BLOCKED`。
- Completion Criteria：
  - 本切片输出存在。
  - 修改范围不超出允许文件。
  - 验证证据已记录。
  - 下一步进入条件明确。
- Failure Handling：
  - 同一失败最多修复 3 轮。
  - 仍失败则生成 issue report，并将任务标记为 `BLOCKED` 或 `REWORK_REQUIRED`。
- Max Repair Attempts：3
- Evidence：
  - 修改文件：`pkg/cli/app_test.go`、`pkg/i18n/i18n_test.go`、`pkg/yaml2go/converter_test.go`、`pkg/yaml2go/converter_impl.go`、`pkg/yaml2go/method_generator.go`、`pkg/yaml2go/utils.go`、项目状态文档。
  - 命令：`gofmt -w pkg/cli/app_test.go pkg/i18n/i18n_test.go pkg/yaml2go/converter_test.go`；`go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 验证结论：首批三包行为测试完成；新增测试暴露的 `pkg/yaml2go` 当前范围缺陷已修复。
- Next Slice Entry Conditions：
  - 用户确认继续 `pkg/cache`、`pkg/executor`、`pkg/httpserver`、`pkg/storage` 等后续批次，或进入 Phase 6 收尾，或结束本轮。

### TS-NEXT-SCOPE-004：确认首批 `pkg/*` 行为测试后的后续范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-004
- Matrix：BL-020、TM-P1-003、TM-P1-007
- Allowed Files：
  - 项目状态文档
- Forbidden Files：
  - Go 源码和测试文件，除非用户确认进入新的代码或测试切片。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户发送“下一步”，按选项 A 继续 `BL-020` 下一批。
  - [CONFIRMED] 状态文档记录新的唯一合法任务 TASK-P1-012 / TS-P1-012。

### TS-P1-012：补第二批 `pkg/*` 行为测试

- Status：COMPLETED
- Task ID：TASK-P1-012
- Matrix：BL-020、TM-P1-003、TM-P1-008、TM-P0-006
- Purpose：为 `pkg/executor`、`pkg/httpserver`、`pkg/storage` 补最小包级行为测试，并在测试暴露当前包缺陷时做最小修复。
- Inputs：
  - `BL-020`
  - `MODULES.md` 的 `pkg/*` 测试缺口
  - `pkg/executor`、`pkg/httpserver`、`pkg/storage` README 和源码
- Allowed Files：
  - `pkg/executor/**/*_test.go`
  - `pkg/httpserver/**/*_test.go`
  - `pkg/storage/**/*_test.go`
  - 必要时限当前三包实现文件：`pkg/executor/*.go`、`pkg/httpserver/*.go`、`pkg/storage/*.go`
  - 项目状态文档
- Forbidden Files：
  - `pkg/cache/**/*`
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他无关 `pkg/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不测试或接入真实 Redis、数据库、第三方网络服务或生产配置。
  - 不补 `pkg/cache` 行为测试。
  - 不启动依赖固定端口的 HTTP 服务；需要 HTTP 生命周期验证时必须使用确定性隔离策略。
  - 不实现 httpserver 文档中尚未落地的 executor 注入能力。
- Execution Steps：
  1. 阅读当前三包源码，确定不依赖外部服务的可测公共行为。
  2. 新增 `pkg/executor/executor_test.go`、`pkg/httpserver/httpserver_test.go`、`pkg/storage/storage_test.go`。
  3. 若测试暴露当前三包内缺陷，按最小修复处理并记录。
  4. 运行格式化和验证命令。
  5. 更新状态、验收、测试报告、变更、风险、问题和交接文档。
- Verification Commands：
  - `gofmt -w pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`
  - `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] `pkg/executor` 覆盖配置校验、任务执行、缺失池、关闭、过载、失败 reload 和 panic handler。
  - [CONFIRMED] `pkg/httpserver` 覆盖构造、默认配置、配置错误、停止态 reload/shutdown 和已运行 start 拒绝路径。
  - [CONFIRMED] `pkg/storage` 覆盖内存文件系统读写、复制、MIME、Excel、图片和配置错误路径。
  - [CONFIRMED] 验证命令通过。
- Failure Handling：
  - 同一问题最多修复 3 轮，每轮记录失败现象、原因假设、修改内容、验证命令和结果。
- Exit Conditions：
  - [CONFIRMED] 三包测试存在并通过。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`pkg/executor/executor_test.go`、`pkg/httpserver/httpserver_test.go`、`pkg/storage/storage_test.go`、`pkg/executor/constants.go`、`pkg/executor/manager.go`、`pkg/executor/pool.go`、项目状态文档。
  - 命令：`gofmt -w pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`；`go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 验证结论：第二批三包行为测试完成；新增测试暴露的 `pkg/executor` 当前范围缺陷已修复。
- Next Slice Entry Conditions：
  - 用户确认继续 `pkg/cache` 等剩余 `BL-020` 范围，或进入 Phase 6 收尾，或结束本轮。

### TS-NEXT-SCOPE-005：确认第二批 `pkg/*` 行为测试后的后续范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-005
- Matrix：BL-020、TM-P1-003、TM-P1-008
- Allowed Files：
  - 项目状态文档
- Forbidden Files：
  - Go 源码和测试文件，除非用户确认进入新的代码或测试切片。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户选择 A：继续 `BL-020` 剩余范围。
  - [CONFIRMED] 状态文档记录新的唯一合法任务 TASK-P1-013 / TS-P1-013。

### TS-P1-013：`pkg/cache` 隔离行为测试

- Status：COMPLETED
- Task ID：TASK-P1-013
- Matrix：BL-020、TM-P1-003、TM-P1-009、TM-P0-006
- Purpose：为 `pkg/cache` 补最小包级行为测试，用进程内隔离 Redis 或等价策略覆盖 Redis 成功路径和错误路径。
- Inputs：
  - 用户选择 A
  - `BL-020`
  - `pkg/cache` README 和源码
  - `ARCHITECTURE.md` / `MODULES.md` 中的 `pkg/cache` 风险说明
- Allowed Files：
  - `pkg/cache/**/*_test.go`
  - 必要时限当前包实现文件：`pkg/cache/*.go`
  - 如隔离 Redis 需要纯测试依赖，可修改 `go.mod`、`go.sum`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他无关 `pkg/*`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不连接真实 Redis 服务。
  - 不修改生产配置。
  - 不重构 `Cache` 公共接口。
  - 不实现新的缓存后端、分布式锁语义或业务 key 约定。
- Execution Steps：
  1. 阅读 `pkg/cache` 源码并确认可测公共行为。
  2. 新增 `pkg/cache` 包级测试，覆盖配置校验、连接、基本读写、批量、计数器、过期、缺失键、重载失败和重载成功。
  3. 若测试暴露当前包缺陷，只做最小修复并记录。
  4. 运行格式化和验证命令。
  5. 更新状态、验收、测试报告、变更、风险、问题和交接文档。
- Verification Commands：
  - `gofmt -w pkg/cache/*_test.go`
  - `go test ./pkg/cache -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] `pkg/cache` 配置默认值和校验路径被测试固定。
  - [CONFIRMED] Redis 基本操作、批量操作、计数器、TTL/Expire 和缺失键语义被隔离测试覆盖。
  - [CONFIRMED] `Reload` 失败保持旧连接、成功切换新连接的语义被测试覆盖。
  - [CONFIRMED] 测试不依赖真实 Redis、外部网络服务、数据库或生产配置。
- Failure Handling：
  - 同一问题最多修复 3 轮，每轮记录失败现象、原因假设、修改内容、验证命令和结果。
- Exit Conditions：
  - [CONFIRMED] `pkg/cache` 测试存在并通过。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`pkg/cache/cache_test.go`、`go.mod`、`go.sum`、项目状态文档。
  - 命令：`go get github.com/alicebob/miniredis/v2@latest`；`gofmt -w pkg/cache/cache_test.go`；`go test ./pkg/cache -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 验证结论：`pkg/cache` Redis 行为已通过进程内隔离 Redis 测试覆盖。
- Next Slice Entry Conditions：
  - 用户确认进入 Phase 6 收尾、提升 `pkg/utils` 等内部支撑测试，或结束本轮。

### TS-NEXT-SCOPE-006：确认 `pkg/cache` 行为测试后的后续范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-006
- Matrix：BL-020、TM-P1-003、TM-P1-009
- Allowed Files：
  - 项目状态文档
- Forbidden Files：
  - Go 源码和测试文件，除非用户确认进入新的代码或测试切片。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户选择 B：提升内部支撑测试。
  - [CONFIRMED] 状态文档记录新的唯一合法任务 TASK-P1-014 / TS-P1-014。

### TS-P1-014：`pkg/utils` 内部支撑测试

- Status：COMPLETED
- Task ID：TASK-P1-014
- Matrix：BL-023、TM-P1-010、TM-P0-006
- Purpose：为 `pkg/utils` 补最小确定性行为测试，覆盖内部支撑工具的稳定成功路径和明确错误路径。
- Inputs：
  - 用户选择 B
  - `BL-023`
  - `pkg/utils` README 和源码
  - `ARCHITECTURE.md` / `MODULES.md` 中的 `pkg/utils` 内部支撑定位
- Allowed Files：
  - `pkg/utils/**/*_test.go`
  - 必要时限当前包实现文件：`pkg/utils/*.go`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他无关 `pkg/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不修改 `pkg/utils` 公共 API。
  - 不改变默认 Snowflake panic 策略。
  - 不依赖真实外部网络服务、固定端口、生产配置或机器专属断言。
  - 不补 `internal/*` 集成测试。
- Execution Steps：
  1. 阅读 `pkg/utils` 源码并确认可测公共行为。
  2. 新增 `pkg/utils` 包级测试，覆盖 Snowflake、地址校验、端口查找、设备 ID 稳定性和 i18n helper 委托。
  3. 若测试暴露当前包缺陷，只做最小修复并记录。
  4. 运行格式化和验证命令。
  5. 更新状态、验收、测试报告、变更、风险、问题和交接文档。
- Verification Commands：
  - `gofmt -w pkg/utils/*_test.go`
  - `go test ./pkg/utils -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] `pkg/utils` Snowflake、地址校验、端口查找、设备 ID 和 i18n helper 有最小行为测试。
  - [CONFIRMED] 测试不依赖真实外部网络服务、固定生产端口、数据库或生产配置。
  - [CONFIRMED] 验证命令通过。
- Failure Handling：
  - 同一问题最多修复 3 轮，每轮记录失败现象、原因假设、修改内容、验证命令和结果。
- Exit Conditions：
  - [CONFIRMED] `pkg/utils` 测试存在并通过。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`pkg/utils/utils_test.go`、项目状态文档。
  - 命令：`gofmt -w pkg/utils/utils_test.go`；`go test ./pkg/utils -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 验证结论：`pkg/utils` 内部支撑工具已有最小确定性行为测试。
- Next Slice Entry Conditions：
  - 用户确认进入 Phase 6 收尾、提升 app/router/middleware 等集成测试，或结束本轮。

### TS-NEXT-SCOPE-007：确认 `pkg/utils` 内部支撑测试后的后续范围

- Status：PENDING_USER_CONFIRMATION
- Task ID：TASK-NEXT-SCOPE-007
- Matrix：BL-023、TM-P1-010、TM-P0-006
- Allowed Files：
  - 项目状态文档
- Forbidden Files：
  - Go 源码和测试文件，除非用户确认进入新的代码或测试切片。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [PENDING] 用户选择后续范围。
  - [PENDING] 状态文档记录新的唯一合法任务或收尾状态。

## 历史时间切片

### TS-HIST-PLUGIN-001：实现插件系统 v1

- Status：COMPLETED
- Summary：历史切片，完成 `pkg/plugin` local/http 支持。

### TS-HIST-PLUGIN-002：确认插件系统 v1 API

- Status：COMPLETED
- Summary：历史切片，确认 v1 local/http 边界；后续扩展留在 Backlog。
