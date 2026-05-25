# TIME_SLICES.md

## 当前合法时间切片

- Time Slice ID：TS-P1-009
- Task ID：TASK-P1-009
- Status：NOT_STARTED
- Summary：用户选择 A，`BL-021` / `TM-P1-005` 已提升；当前唯一合法切片是明确 `types/*` 契约边界。

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

- Status：NOT_STARTED
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

## 历史时间切片

### TS-HIST-PLUGIN-001：实现插件系统 v1

- Status：COMPLETED
- Summary：历史切片，完成 `pkg/plugin` local/http 支持。

### TS-HIST-PLUGIN-002：确认插件系统 v1 API

- Status：COMPLETED
- Summary：历史切片，确认 v1 local/http 边界；后续扩展留在 Backlog。
