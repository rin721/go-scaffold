# TIME_SLICES.md

## 当前合法时间切片

- Time Slice ID：NONE
- Task ID：NONE
- Status：COMPLETED
- Summary：`dev.tmp/new-plugin.md` 主线已完成；TASK-P2-005 至 TASK-P2-010 均已通过验证。TASK-P2-004 Docker build 验证因缺少 Docker CLI 保留为 `ISSUE-P2-005`。

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

### TS-INFRA-003：TASK-P1-017 后状态一致性修复

- Status：COMPLETED
- Task ID：TASK-INFRA-003
- Scope：
  - 生成状态诊断报告。
  - 修复 `ARCHITECTURE.md`、`MODULES.md`、`PROJECT_BRIEF.md` 和 `ROADMAP.md` 中 TASK-P1-016 前的旧待办表述。
  - 更新状态、验收、测试报告、变更记录、问题记录、风险和交接说明。
- Allowed Files：
  - 根目录项目状态文档。
  - `docs/reports/status_diagnostics/*`
- Forbidden：
  - 不修改 Go 源码或测试文件。
  - 不修改 `go.mod`、`go.sum`。
  - 不修改数据库 schema、部署配置或真实密钥。
- Verification：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- Exit Conditions：
  - 背景文档不再把 TASK-P1-016 已覆盖路径描述为待补范围。
  - 当前合法下一步明确为 `NONE / COMPLETED`。

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

- Status：COMPLETED
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
  - [CONFIRMED] 用户回复 `b`，选择 B：提升 app/router/middleware 等集成测试。
  - [CONFIRMED] 状态文档记录新的唯一合法任务 TASK-P1-015 / TS-P1-015。

### TS-P1-015：app/router/middleware 最小集成测试

- Status：COMPLETED
- Task ID：TASK-P1-015
- Matrix：BL-002、TM-P0-005、TM-P0-006
- Purpose：用 `httptest` 验证 demo Todo HTTP 路由注册和 handler/service/repository 集成路径，并固定 TraceID、CORS、Recovery 中间件链路的最小语义。
- Inputs：
  - 用户回复 `b`
  - `BL-002`
  - `MODULES.md` 中的 `internal/middleware`、`internal/transport/http`、`internal/modules/demo` 测试缺口
  - `internal/transport/http`、`internal/middleware`、`internal/modules/demo` 源码和既有测试
- Allowed Files：
  - `internal/transport/http/**/*_test.go`
  - `internal/middleware/**/*_test.go`
  - `internal/modules/demo/**/*_test.go`
  - 必要时限当前范围实现文件：`internal/transport/http/*.go`、`internal/middleware/*.go`、`internal/modules/demo/handler/*.go`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `pkg/**/*`
  - `types/**/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不启动真实 HTTP server，不绑定固定端口。
  - 不依赖真实外部数据库、Redis、第三方网络服务或生产配置。
  - 不重构 router/middleware/demo 分层或公共 API。
  - 不进入 Phase 6 收尾。
- Execution Steps：
  1. 阅读当前允许范围源码和既有测试。
  2. 使用临时 SQLite 构建真实 demo repository/service/handler，并注入 `NewRouter`。
  3. 新增最小集成测试覆盖 demo Todo HTTP 成功路径、错误路径或路由注册，以及 TraceID/CORS/Recovery 链路。
  4. 运行格式化和验证命令。
  5. 更新状态、验收、测试报告、变更、风险、问题和交接文档。
- Verification Commands：
  - `gofmt -w internal/transport/http/*_test.go internal/middleware/*_test.go internal/modules/demo/**/*_test.go`
  - `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] demo Todo router/handler/service/repository HTTP 集成路径被覆盖。
  - [CONFIRMED] TraceID、CORS、Recovery 至少有路由级链路断言。
  - [CONFIRMED] 测试隔离且不依赖生产配置或外部服务。
- Failure Handling：
  - 同一问题最多修复 3 轮，每轮记录失败现象、原因假设、修改内容、验证命令和结果。
- Exit Conditions：
  - [CONFIRMED] 集成测试存在并通过。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`internal/transport/http/router_integration_test.go`、项目状态文档。
  - 命令：`gofmt -w internal/transport/http/router_integration_test.go`；`go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 验证结论：demo Todo router/handler/service/repository HTTP 集成和 TraceID/CORS/Recovery 链路已有最小测试覆盖。
  - 修复记录：前两次相关包测试失败来自测试构造问题，固定 `httptest` Host 后通过。
- Next Slice Entry Conditions：
  - 用户确认进入 Phase 6 收尾、继续 app 装配/reload/config 等剩余集成测试，或结束本轮。

### TS-NEXT-SCOPE-008：确认 app/router/middleware 集成测试后的后续范围

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-008
- Matrix：BL-002、TM-P0-004、TM-P0-006
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
  - [CONFIRMED] 用户选择 A：进入 Phase 6 收尾。
  - [CONFIRMED] 状态文档记录新的唯一合法任务和切片 TASK-PHASE6-001 / TS-PHASE6-001。

### TS-PHASE6-001：Phase 6 收尾与交接

- Status：COMPLETED
- Task ID：TASK-PHASE6-001
- Matrix：TM-P0-006
- Purpose：冻结本轮项目优化成果，完成最终状态、验收、测试报告、变更记录、风险/Backlog 和交接更新。
- Inputs：
  - 用户最新回复 `a`
  - TASK-P1-015 完成证据
  - 当前项目状态文档
- Allowed Files：
  - 根目录项目状态文档
  - `AGENT_HANDOFF.md`
- Forbidden Files：
  - Go 源码和测试文件。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Verification Commands：
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] Phase 6 收尾文档更新完成。
  - [CONFIRMED] 最终验证命令已执行并记录：`go test ./... -count=1` 与 `git diff --check` 均通过。
  - [CONFIRMED] 无自动下一实现任务，后续工作需要用户重新确认。
- Evidence：
  - 本切片未修改 Go 源码或测试文件。
  - 项目状态、验收、变更、测试报告、问题和交接文档已同步到收尾完成状态。
  - app 装配、reload/config 等剩余集成测试未在本切片继续实现；该历史后续范围已由 TS-P1-016 覆盖并关闭。

### TS-P1-016：app 装配与 reload/config 剩余集成测试

- Status：COMPLETED
- Task ID：TASK-P1-016
- Matrix：BL-002、TM-P0-004、TM-P0-006、TM-P1-012、RISK-008
- Purpose：补齐 app 装配、配置变更 hook 与 reload 分发的剩余集成测试，固定真实 `app.New` server/initdb 装配语义和 reload 组件分发语义。
- Inputs：
  - 用户明确要求实施 TASK-P1-016 计划。
  - `BL-002` 剩余 app 装配/reload/config 测试范围。
  - `TASK-PHASE6-001` 收尾记录中的后续确认范围。
  - `internal/app`、`internal/app/reloadapp`、`internal/config` 现有源码和测试。
- Allowed Files：
  - `internal/app/app_integration_test.go`
  - `internal/app/reloadapp/reload_test.go`
  - 必要时限当前范围修复文件：`internal/app/**`、`internal/config/**`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `pkg/**/*`
  - `types/**/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不新增功能或导出 API。
  - 不启动真实 HTTP server，不占用固定端口。
  - 不依赖 Redis/MySQL/Postgres/外部网络或生产配置。
  - 不重构装配或 reload 结构。
- Execution Steps：
  1. 将当前唯一合法任务/切片切到 TASK-P1-016 / TS-P1-016。
  2. 用临时 YAML 和临时 SQLite 编写真实 `app.New` 装配测试。
  3. 用 fake cache/database/logger/executor/httpserver/storage 编写 reload 分发测试。
  4. 运行格式化、`internal/app` 包测试、全量回归和 diff 检查。
  5. 更新状态、验收、测试报告、变更记录、Backlog、风险、问题和交接说明。
- Verification Commands：
  - `gofmt -w internal/app/app_integration_test.go internal/app/reloadapp/reload_test.go`
  - `go test ./internal/app/... -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] server 模式 Core / Infra / Modules / Transport 最小装配链路通过真实 app 构建验证。
  - [CONFIRMED] initdb 模式仅初始化数据库并创建 demo schema，不装配 HTTP transport。
  - [CONFIRMED] 配置变更 hook 能更新 `Core.Config`，且未启动真实 HTTP server。
  - [CONFIRMED] reload 分发覆盖配置未变、单组件变化、可选组件关闭和 database reload 不隐式迁移。
- Failure Handling：
  - 同一问题最多修复 3 轮，每轮记录失败现象、原因假设、修改内容、验证命令和结果。
- Exit Conditions：
  - [CONFIRMED] 新增测试存在并通过。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接说明已更新。
- Evidence：
  - 修改文件：`internal/app/app_integration_test.go`、`internal/app/reloadapp/reload_test.go`、项目状态文档。
  - 命令：`gofmt -w internal/app/app_integration_test.go internal/app/reloadapp/reload_test.go`；`go test ./internal/app/... -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 结论：app 装配、reload/config 剩余集成测试完成；当前无自动下一实现任务。

### TS-P1-017：`pkg/*` README 第一阶段中文化

- Status：COMPLETED
- Task ID：TASK-P1-017
- Matrix：BL-006、RISK-005、TM-P1-013、TM-P0-006
- Purpose：把 `pkg/*/README.md` 的主要读者文本统一为中文，降低包文档中英混杂和交接成本。
- Inputs：
  - 用户选择 A
  - `BL-006`
  - `RISK-005`
  - 现有 `pkg/*/README.md`
- Allowed Files：
  - `pkg/*/README.md`
  - `REQUIREMENTS.md`
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - 项目状态文档
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - `pkg/**/*` 中除 `README.md` 以外的文件
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Strict Non-Goals：
  - 不修改 Go 代码或测试文件。
  - 不新增功能或导出 API。
  - 不修改配置、路由、数据库 schema 或依赖。
  - 不强行翻译 Go 标识符、代码示例、协议名、环境变量名和必要技术名词。
- Execution Steps：
  1. 识别 `pkg/*/README.md` 中标题、段落、风险说明、许可证和边界说明的中英混杂点。
  2. 中文化读者可见文本，并同步已完成测试后的明显过期风险表述。
  3. 运行 `go test ./... -count=1` 和 `git diff --check`。
  4. 更新状态、验收、测试报告、变更、风险、Backlog、问题和交接文档。
- Verification Commands：
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] 13 个 `pkg/*/README.md` 的主要说明已中文化或确认无需修改。
  - [CONFIRMED] 文档不新增未实现能力承诺。
  - [CONFIRMED] 验证命令通过。
- Exit Conditions：
  - [CONFIRMED] README 中文化完成。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接说明已更新。
- Evidence：
  - 修改文件：12 个 `pkg/*/README.md`、`REQUIREMENTS.md`、`ARCHITECTURE.md`、`MODULES.md` 和项目状态文档。
  - `pkg/crypto/README.md` 已检查，当前阶段无需修改。
  - 命令：`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 结论：第一阶段包 README 中文化完成；当前无自动下一实现任务。

## P2 时间切片

### TS-P2-001：CI 质量门禁与部署说明首切片

- Status：COMPLETED
- Task ID：TASK-P2-001
- Matrix：REQ-OPT-P2-003、BL-007、BL-008、TM-P2-001、RISK-016
- Purpose：建立非生产 CI 质量门禁和手动部署说明，让后续发布前检查可恢复、可验证。
- Inputs：
  - 用户选择 D。
  - `REQ-OPT-P2-003`。
  - `BL-007`、`BL-008`。
- Allowed Files：
  - `.github/workflows/ci.yml`
  - `docs/deployment.md`
  - `README.md`
  - 项目状态文档
- Forbidden：
  - 不修改 Go 源码或测试文件。
  - 不修改 `go.mod`、`go.sum`。
  - 不新增生产配置、真实 `.env`、密钥或凭据。
  - 不连接远程环境、不推送镜像、不执行部署。
- Strict Non-Goals：
  - 不实现真实 CD。
  - 不新增 Dockerfile、Kubernetes、systemd 或云平台模板。
  - 不实现生产迁移框架。
- Verification Commands：
  - gofmt 漂移报告（非阻塞）
  - `go test ./... -count=1`
  - `go build -o <temp> ./cmd/server`
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] CI workflow 权限最小化为只读内容权限。
  - [CONFIRMED] CI 不使用 secrets、不推送产物、不部署。
  - [CONFIRMED] 历史 gofmt 漂移不作为本切片硬门禁，已记录到 `BL-025`。
  - [CONFIRMED] 部署说明覆盖配置入口、手动运行、initdb 和未实现项。
  - [CONFIRMED] 本地等价验证通过。
- Evidence：
  - 新增 `.github/workflows/ci.yml`。
  - 新增 `docs/deployment.md`。
  - README 新增 CI 与部署入口。
  - 验证结果：硬门禁 PASS；gofmt 漂移审计为 KNOWN_DRIFT，已记录 `BL-025`；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 结论：CI 质量门禁与部署说明首切片完成；当前无自动下一实现任务。

### TS-NEXT-SCOPE-010：真实 CD / 镜像发布 / 远程部署自动化范围确认

- Status：COMPLETED
- Task ID：TASK-NEXT-SCOPE-010
- Matrix：BL-024、REQ-OPT-P2-003、RISK-016、RISK-017
- Purpose：把用户选择 C 转成可恢复的待确认状态，收集真实 CD 实现所需的最小决策。
- Inputs：
  - 用户回复 `c`
  - 用户补充“使用远程部署”
  - `BL-024`
  - TASK-P2-001 已完成的 CI 质量门禁和部署说明
- Allowed Files：
  - 项目状态文档
- Forbidden：
  - `.github/workflows/*` 实现或修改
  - Go 源码、测试文件、依赖文件
  - Dockerfile、Kubernetes、systemd、云平台模板
  - 真实 `.env`、密钥、部署凭据、生产配置
- Strict Non-Goals：
  - 不推送镜像。
  - 不连接远程服务器。
  - 不部署 staging 或 production。
  - 不读取、输出或假造真实 secrets。
- Verification Commands：
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] 用户选择 C 已被记录。
  - [CONFIRMED] 用户确认使用远程部署。
  - [CONFIRMED] 审查结论为 `NEEDS_USER_DECISION`。
  - [CONFIRMED] 用户确认使用 `.env` 风格文件配置远程部署参数。
- Next Slice Entry Conditions：
  - 已进入并完成 TS-P2-002。

### TS-P2-002：远程部署 env 配置模板

- Status：COMPLETED
- Task ID：TASK-P2-002
- Matrix：BL-024、TM-P2-003、RISK-016、RISK-017
- Purpose：为远程部署提供可提交的 `.env` 风格示例模板，同时避免真实部署配置进入 Git。
- Inputs：
  - 用户要求“远程部署 .env 来配置”
  - TASK-NEXT-SCOPE-010
  - `docs/deployment.md`
- Allowed Files：
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `.gitignore`
  - `README.md`
  - `docs/deployment.md`
  - 项目状态文档
- Forbidden：
  - 真实 `.env` / 显式部署参数
  - `.github/workflows/*` 自动部署实现
  - Go 源码、测试文件、依赖文件
  - Dockerfile、Kubernetes、systemd、云平台模板
  - 真实服务器地址、密钥、token、密码或生产配置
- Strict Non-Goals：
  - 不推送镜像。
  - 不连接远程服务器。
  - 不部署 staging 或 production。
  - 不读取、输出或假造真实 secrets。
- Verification Commands：
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] `deploy.sh` / `script/install.sh` 显式参数契约 存在且只包含占位值。
  - [CONFIRMED] 旧本地部署 env 文件已删除。
  - [CONFIRMED] 部署说明记录远程部署变量边界。
  - [CONFIRMED] 未实现真实 CD workflow、镜像发布或远程连接。

### TS-P2-003：手动远程部署 workflow

- Status：COMPLETED
- Task ID：TASK-P2-003
- Matrix：BL-024、TM-P2-004、RISK-016、RISK-017
- Purpose：在用户确认后新增手动远程部署 GitHub Actions workflow，固定 staging/manual/Secrets/SSH/Docker Compose 的安全边界。
- Inputs：
  - 用户明确回复“确认实现远程部署 workflow”
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `docs/deployment.md`
  - TASK-P2-001 / TASK-P2-002 的 CI 与部署说明
- Allowed Files：
  - `.github/workflows/deploy-remote.yml`
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `README.md`
  - `docs/deployment.md`
  - 项目状态文档
- Forbidden：
  - 真实 `.env` / 显式部署参数
  - Go 源码、测试文件、依赖文件
  - Dockerfile、Kubernetes、systemd、云平台部署模板
  - 真实服务器地址、密钥、token、密码或生产配置
- Strict Non-Goals：
  - 不在当前会话触发 GitHub workflow。
  - 不连接远程服务器、不推送镜像、不执行 staging 或 production 部署。
  - 不新增生产迁移框架。
  - 不把 production 作为默认部署环境。
- Execution Steps：
  1. 新增 `.github/workflows/deploy-remote.yml`。
  2. workflow 使用 `workflow_dispatch` 和确认输入，默认只支持 `staging`。
  3. workflow 从 GitHub Variables/Secrets 组装显式部署参数，校验远程 SSH 输入和安全占位。
  4. workflow 配置 SSH key/known_hosts，通过 SSH 在远程主机执行 `script/install.sh` / `deploy.sh`。
  5. workflow 在远程主机执行 Docker Compose pull/up 和 health/ready 检查。
  6. 更新 `deploy.sh` / `script/install.sh` 显式参数契约、部署说明、README 和状态文档。
  7. 运行验证命令并记录结果。
- Verification Commands：
  - workflow YAML 结构检查
  - `git diff --check`
- Acceptance：
  - [CONFIRMED] workflow 仅手动触发，且需要输入确认词。
  - [CONFIRMED] workflow 不包含真实密钥、真实服务器地址或生产配置。
  - [CONFIRMED] workflow 使用 GitHub Secrets 注入 显式部署参数、SSH key 和可选 registry token。
  - [CONFIRMED] workflow 不构建或推送镜像；远程主机按 `DEPLOY_IMAGE` 拉取既有镜像。
  - [CONFIRMED] 部署说明包含 Secrets 配置、远程主机前置条件和手动触发步骤。
- Evidence：
  - 新增 `.github/workflows/deploy-remote.yml`。
  - 更新 `deploy.sh` / `script/install.sh` 显式参数契约、`README.md`、`docs/deployment.md` 和项目状态文档。
  - 验证：临时 Go YAML 解析 PASS；actionlint PASS；`git diff --check` PASS，仅有 Windows LF/CRLF 转换警告。
  - Go 测试未运行：本切片未修改 Go 代码、依赖、配置 schema、HTTP 路由或数据库 schema。

### TS-P2-004：Linux Docker production 部署制品

- Status：PENDING_VERIFICATION
- Task ID：TASK-P2-004
- Matrix：BL-024、TM-P2-005、RISK-016、RISK-017、RISK-018
- Purpose：为 Linux Docker production 远程部署补齐可提交运行制品、production Compose 示例、统一 `deploy.sh` 部署入口和受控手动 workflow 闸门。
- Inputs：
  - 用户要求“开始，linux、docker、production -> 部署”
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `.github/workflows/deploy-remote.yml`
  - `docs/deployment.md`
  - TASK-P2-003 已完成的 staging 远程部署 workflow
- Allowed Files：
  - `Dockerfile`
  - `.dockerignore`
  - `deploy/docker-compose.production.example.yml`
  - `deploy/config.production.example.yaml`
  - `deploy.sh`
  - `.github/workflows/deploy-remote.yml`
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `README.md`
  - `docs/deployment.md`
  - 项目状态文档
- Forbidden：
  - 真实 `.env` / 显式部署参数
  - Go 源码、测试文件、依赖文件
  - 数据库 schema、真实服务器地址、密钥、token、密码或生产配置
  - 自动触发 production 部署
  - 生产迁移框架
- Strict Non-Goals：
  - 不在当前会话触发 GitHub workflow。
  - 不连接远程服务器、不推送镜像、不执行真实 production。
  - 不新增业务功能、Go API、配置 schema、HTTP 路由或数据库 schema。
- Execution Steps：
  1. 新增 Dockerfile 和 `.dockerignore`，构建 Linux server 镜像。
  2. 新增 production Compose 示例，使用 显式部署参数 中的 `DEPLOY_IMAGE` 和 `APP_PORT`。
  3. 扩展远程部署 workflow 支持 `staging` / `production` 手动选择，并要求 `deploy-staging` 或 `deploy-production` 确认词。
  4. 新增远程 Linux 部署脚本，按参数/环境变量动态生成 `运行期显式部署参数`。
  5. 更新 显式参数部署入口、部署说明、README 和状态文档。
  6. 运行 Docker/Workflow/Go/脚本语法验证和 diff 检查。
- Verification Commands：
  - `docker build -t go-scaffold:local .`：PENDING，当前本机未安装 Docker CLI；`docker`、`podman`、`nerdctl` 均不可用。
  - 临时 Go YAML 解析：PASS。
  - `go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml`：PASS。
  - `bash -n deploy.sh`：FAIL_ENV，本机无可用 bash，WSL 未安装 Linux 发行版。
  - `shfmt` Bash 语法解析：PASS。
  - `go test ./... -count=1`：PASS。
  - `go build -o <temp> ./cmd/server`：PASS。
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告。
- Acceptance：
  - [PENDING_VERIFICATION] Dockerfile 已存在并配置非 root 用户运行 server；镜像构建待具备 Docker 的环境验证。
  - [CONFIRMED] Compose 示例适合 Linux Docker production 远程主机，并保留真实配置外置挂载。
  - [CONFIRMED] production 配置样例绑定 `0.0.0.0:9999`，不包含真实密钥。
  - [CONFIRMED] 远程 Linux 部署脚本会按显式参数注入运行环境，脚本不打印 password/token/secret 值。
  - [CONFIRMED] workflow 只手动触发，production 需要显式环境选择和确认词。
  - [CONFIRMED] 文档说明 production Secrets、GitHub Environment 审批、远程目录权限和不执行真实部署边界。
- Exit Conditions：
  - [CONFIRMED] 制品和文档完成。
  - [PENDING_VERIFICATION] Docker build 不可执行原因已记录，仍需补跑。
  - [CONFIRMED] 状态文档、测试报告和交接说明已更新。
- Evidence：
  - 修改文件：`Dockerfile`、`.dockerignore`、`deploy/docker-compose.production.example.yml`、`deploy/config.production.example.yaml`、`deploy.sh`、`.github/workflows/deploy-remote.yml`、`deploy.sh` / `script/install.sh` 显式参数契约、`README.md`、`docs/deployment.md` 和项目状态文档。
  - 未执行真实部署、未触发 GitHub workflow、未连接远程服务器、未推送镜像、未写入真实 `.env` 或 secrets。
  - 下一步：在具备 Docker CLI/daemon 的 Linux 或 Docker Desktop 环境执行 `docker build -t go-scaffold:local .`。

### TS-P2-005：`pkg/plugin/hooks` 独立钩子引擎

- Status：COMPLETED
- Task ID：TASK-P2-005
- Matrix：TM-P2-006
- Purpose：实现无应用层依赖的钩子注册、排序、执行和停止语义。
- Allowed Files：`pkg/plugin/hooks/**/*`、项目状态文档。
- Verification Commands：`go test ./pkg/plugin/hooks -count=1`；`go test ./pkg/plugin/... -count=1`。
- Acceptance：优先级高到低；执行前复制处理器；context 取消生效；nil handler 被拒绝；停止语义可测试。
- Evidence：新增 `pkg/plugin/hooks` 并通过 `pkg/plugin` 相关测试。

### TS-P2-006：插件管理器钩子化

- Status：COMPLETED
- Task ID：TASK-P2-006
- Matrix：TM-P2-007
- Purpose：扩展 `pkg/plugin.Manager` 钩子能力并保持被动注册 API。
- Allowed Files：`pkg/plugin/**/*`、项目状态文档。
- Verification Commands：`go test ./pkg/plugin/... -count=1`。
- Acceptance：注册、调用、调用错误和关闭钩子行为稳定；`pkg/plugin` 不耦合 IAM、logger、config 或 internal。
- Evidence：`Manager` 已支持 `Hooks()`、`RegisterHook`、`WithHooks` 和标准钩子点；包边界保持解耦。

### TS-P2-007：HTTP 远程插件与远程钩子

- Status：COMPLETED
- Task ID：TASK-P2-007
- Matrix：TM-P2-008
- Purpose：新增 HTTP server helper、标准操作和 `RemoteHook`。
- Allowed Files：`pkg/plugin/**/*`、项目状态文档。
- Verification Commands：`go test ./pkg/plugin/... -count=1`。
- Acceptance：`POST /plugin/v1/invoke`、非法请求、插件错误、响应限制和 `hooks.execute` round trip 均被测试固定。
- Evidence：新增 `NewHTTPServer`、`hooks.execute` 和 `RemoteHook`，HTTP/远程钩子路径测试通过。

### TS-P2-008：IAM 公共 API 与 memory 实现

- Status：COMPLETED
- Task ID：TASK-P2-008
- Matrix：TM-P2-009
- Purpose：新增独立身份认证与授权基础设施包。
- Allowed Files：`pkg/iam/**/*`、项目状态文档。
- Verification Commands：`go test ./pkg/iam/... -count=1`。
- Acceptance：token 认证、策略授权、拒绝优先、通配、过期和 context helper 均有测试。
- Evidence：新增 `pkg/iam` 与 `pkg/iam/memory`，memory 策略与上下文 helper 测试通过。

### TS-P2-009：配置与应用组装接入

- Status：COMPLETED
- Task ID：TASK-P2-009
- Matrix：TM-P2-010
- Purpose：新增 plugin/iam 配置并在 server 模式装配。
- Allowed Files：`internal/config/**/*`、`internal/app/**/*`、`pkg/plugin/**/*`、`pkg/iam/**/*`、项目状态文档。
- Verification Commands：`go test ./internal/config ./internal/app/... ./pkg/plugin/... ./pkg/iam/... -count=1`。
- Acceptance：HTTP 插件适配器、远程钩子绑定、IAM 授权钩子和默认 disabled 配置均可验证。
- Evidence：`plugin` / `iam` 配置默认 disabled，app 组合层装配 IAM、Plugins、RemoteHook 和权限钩子。

### TS-P2-010：reload、生命周期和最终验证

- Status：COMPLETED
- Task ID：TASK-P2-010
- Matrix：TM-P2-011
- Purpose：完成配置重载、关闭顺序、全量验证和状态交接。
- Allowed Files：`internal/app/**/*`、项目状态文档、必要文档。
- Verification Commands：`go test ./internal/config ./internal/app/... -count=1`；`go test ./... -count=1`；`go build -o <temp> ./cmd/server`；`git diff --check`。
- Acceptance：新实例先构建后替换；失败保留旧实例；关闭顺序安全；最终文档可恢复。
- Evidence：reload/lifecycle 已接入并通过目标包测试、全量回归、server build 和 diff 检查。

## 历史时间切片

### TS-HIST-PLUGIN-001：实现插件系统 v1

- Status：COMPLETED
- Summary：历史切片，完成 `pkg/plugin` local/http 支持。

### TS-HIST-PLUGIN-002：确认插件系统 v1 API

- Status：COMPLETED
- Summary：历史切片，确认 v1 local/http 边界；后续扩展留在 Backlog。
