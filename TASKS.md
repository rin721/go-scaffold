# TASKS.md

## 当前合法任务

- Task ID：NONE
- Status：PENDING_USER_CONFIRMATION
- Time Slice：NONE
- Summary：[ACCEPT] 用户纠正当前项目仍未开发完整，不应发布第一版。`dev.tmp/new-plugin.md` 设计已完成，TASK-P2-004 的 Docker build 验证也已解除阻塞；这些只代表已确认切片完成，不代表项目整体完成或 v1 可发布。当前无自动下一实现任务；下一阶段开发范围或第一版发布验收清单需由用户重新确认。

## 任务列表

### TASK-OPT-001：生成中文项目优化启动材料

- Status：COMPLETED
- Verification：
  - `go test ./... -count=1`
- Completion Evidence：
  - 六个中文模板已生成。
  - 当前主线已切换为项目优化启动确认。
  - `go test ./... -count=1` 通过。

### TASK-OPT-002：确认项目优化路线和关键边界

- Status：COMPLETED
- Confirmed Decisions：
  - 优化路线：治理优先。
  - `pkg/*` 定位：混合策略。
  - demo 模块定位：长期标准示例。
  - 迁移策略：dev-prod 分层策略。
  - 中文化范围：根文档和模板优先，历史文档与包 README 分阶段处理。
  - auth/JWT：先延后处理，不进入当前实现范围。
- Completion Evidence：
  - 用户发送“下一步”，按当前文档推荐默认值视为确认。
  - 决策已写入 `DECISIONS.md`。

### TASK-OPT-003：生成模块边界清单和优化路线明细

- Status：COMPLETED
- Goal：逐模块分析当前优缺点、职责边界、设计冲突、测试缺口和优化优先级。
- Requirements Covered：
  - REQ-OPT-P1-001
  - REQ-OPT-P1-002
  - REQ-OPT-P1-003
- Allowed Files：
  - 项目文档和状态文件。
  - `MODULES.md`。
- Non-Goals：
  - 不写 Go 代码。
  - 不重构模块。
  - 不实现 Backlog。
  - 不扩展插件系统。
- Verification：
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - `MODULES.md` 已生成。
  - 模块职责清单、设计边界冲突清单、测试矩阵草案、P1 优化候选项已生成。
  - 状态、测试报告、变更日志和交接文档已更新到 TASK-OPT-004。

### TASK-OPT-004：生成正式测试矩阵和任务拆分草案

- Status：COMPLETED
- Goal：把 `MODULES.md` 的测试矩阵草案和 P1 优化候选项转成正式任务与时间切片草案。
- Requirements Covered：
  - REQ-OPT-P1-004
- Allowed Files：
  - 项目文档和状态文件。
  - 可新增 `TEST_MATRIX.md`。
  - 可更新 `TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`。
  - 不允许修改 Go 实现文件。
- Expected Output：
  - 正式测试矩阵。
  - P1 任务拆分。
  - P1 时间切片草案。
  - 每个任务的允许文件范围和验证命令。
- Non-Goals：
  - 不写 Go 测试代码。
  - 不修复已发现问题。
  - 不实现 Backlog。
- Verification：
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - `TEST_MATRIX.md` 已生成。
  - P1 任务草案已写入本文件。
  - P1 时间切片草案已写入 `TIME_SLICES.md`。

### TASK-OPT-005：确认正式测试矩阵和 P1 执行顺序

- Status：COMPLETED
- Goal：确认 `TEST_MATRIX.md` 中的推荐执行顺序，允许后续进入首个代码优化切片。
- Requirements Covered：
  - REQ-OPT-P1-004
- Recommended Default：
  - 接受 `TEST_MATRIX.md` 的推荐顺序。
  - 首个代码切片为 TASK-P1-001。
- Allowed Files：
  - 项目文档和状态文件。
- Non-Goals：
  - 不写 Go 代码。
  - 不新增 Go 测试文件。
- Completion Rule：
  - 用户明确确认，或再次发送“下一步”按推荐默认顺序视为确认。
- Completion Evidence：
  - 用户再次发送“下一步”，按推荐默认顺序视为确认。
  - 当前已进入并完成 TASK-P1-001。

### TASK-INFRA-001：补齐 Prompt 全量 Agent 基础设施

- Status：COMPLETED
- Goal：补齐 `docs/ai/prompt.md` 要求但当前缺失的 Agent 入口、规则、skills、模板和目录基础设施。
- Reason：
  - 用户确认“prompt 产物生产是否完整”的检查结论，并要求按计划实施。
- Allowed Files：
  - Agent 入口与规则文件。
  - `docs/templates/*`。
  - `docs/reports/*`、`docs/specs/*`。
  - `skills/*/SKILL.md`。
  - `.agents/*`、`.cursor/*`、`.kiro/*`、`.codex/*`。
  - 项目状态文档。
- Non-Goals：
  - 不修改 Go 代码。
  - 不继续 TASK-P1-002。
  - 不修改业务逻辑。
- Verification：
  - Prompt 全量产物存在性核对：PASS。
  - `go test ./... -count=1`：PASS。
- Completion Evidence：
  - `AGENTS.md`、`CLAUDE.md`、`AGENT_RULES.md`、`SKILLS.md` 已新增。
  - 缺失的任务拆分和时间切片模板已新增。
  - reports/specs、跨工具目录和 14 个项目 skills 已新增。
  - 当前合法下一步恢复为 TASK-P1-002。

### TASK-INFRA-002：修复 Agent 基础设施一致性冲突

- Status：COMPLETED
- Goal：修复 `TASK-INFRA-001` 记录与实际文件系统不一致的问题，补齐 `AGENTS.md`、规范化项目 skills、建立 `.agents` 适配器并记录诊断证据。
- Reason：
  - 用户要求实施“Agent 基础设施一致性修复计划”。
  - `STATUS.md` 等文档声称 `AGENTS.md` 已补齐，但仓库根目录实际缺失。
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
- Forbidden Files：
  - Go 源码和测试文件。
  - `go.mod`、`go.sum`。
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - Agent 基础设施文件存在性核对：PASS。
  - canonical 和 adapter skills frontmatter 验证：PASS。
  - 跨工具入口引用一致性检查：PASS。
  - `go test ./... -count=1`：PASS。
- Completion Evidence：
  - `AGENTS.md` 已新增。
  - `AGENT_RULES.md`、`CLAUDE.md`、Cursor、Kiro、Codex 入口已统一引用 `AGENTS.md` 和 `docs/ai/prompt.md`。
  - 14 个 canonical `skills/*/SKILL.md` 已补齐 frontmatter 和完整执行结构。
  - 14 个 `.agents/skills/*/SKILL.md` 轻量适配器已新增。
  - `docs/templates/*` 已标准化为可复用模板。
  - 状态诊断报告已写入 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md`。
  - 当前合法下一步恢复为 TASK-P1-002。

### TASK-INFRA-003：修复 TASK-P1-017 后背景文档状态漂移

- Status：COMPLETED
- Goal：修复用户发送“下一步”时发现的状态漂移：核心状态已确认 TASK-P1-016 和 TASK-P1-017 完成，但部分背景文档仍保留 app 装配、reload/config 后续待补的旧表述。
- Reason：
  - `AGENTS.md` 要求“下一步”前读取状态文档，并在文档冲突时先生成状态诊断报告并修复 Agent 基础设施。
  - `ARCHITECTURE.md`、`MODULES.md`、`PROJECT_BRIEF.md` 和 `ROADMAP.md` 中存在 TASK-P1-016 前的残留表述。
- Allowed Files：
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - `PROJECT_BRIEF.md`
  - `ROADMAP.md`
  - `STATUS.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `ACCEPTANCE.md`
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `ISSUES.md`
  - `RISK_REGISTER.md`
  - `AGENT_HANDOFF.md`
  - `docs/reports/status_diagnostics/*`
- Forbidden Files：
  - Go 源码和测试文件。
  - `go.mod`、`go.sum`。
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- Completion Evidence：
  - 新增 `docs/reports/status_diagnostics/2026-05-26-task-p1-017-post-completion-doc-drift.md`。
  - 背景文档已同步 TASK-P1-016 完成事实，不再把 app 装配、配置变更 hook 或 reload/config 剩余集成测试描述为待补范围。
  - `pkg/i18n` 已补测试事实已同步到架构表述。
  - 当前合法下一步恢复为 `NONE / COMPLETED`。

## P1 任务草案

### TASK-P1-001：修复 `copyConfig` 字段覆盖并补配置测试

- Status：COMPLETED
- Matrix：TM-P0-001、TM-P0-002、TM-P0-006
- Goal：为配置复制和更新建立测试，并修复字段遗漏问题。
- Allowed Files：
  - `internal/config/*`
  - 必要 testdata
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - `copyConfig` 不丢失已确认关键字段。
  - 配置 Update/copy 测试通过。
  - 状态、测试报告、变更日志和交接文档更新。
- Completion Evidence：
  - `internal/config/manager.go` 已修复为结构体整体复制并深拷贝 slice。
  - `internal/config/manager_test.go` 已覆盖完整字段复制、slice 深拷贝和 `Update` 保留未修改字段。

### TASK-P1-002：统一配置环境变量策略

- Status：COMPLETED
- Matrix：TM-P0-001、TM-P0-006
- Goal：收拢 `.env.example` 与配置覆盖实现的环境变量命名策略。
- Allowed Files：
  - `internal/config/*`
  - `.env.example`
  - 配置相关文档
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - 数据库、server、logger、i18n 等环境变量策略一致或差异被明确记录。
  - `.env.example` 与实现不再冲突。
- Completion Evidence：
  - 数据库配置覆盖现在优先读取 `DB_*`，旧 `REI_APP_DB_*` 保留为兼容 fallback，且 `DB_*` 优先级更高。
  - `.env.example` 已移除未实现的 JWT 示例，并列出 DB、Redis、Server、Logger、I18n、Storage、CORS 的实际环境变量策略。
  - `internal/config/manager_test.go` 已覆盖数据库 `DB_*` 主策略、旧前缀 fallback、非数据库变量覆盖。
  - `go test ./internal/config -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-P1-003：增加 health/ready 与 router smoke test

- Status：COMPLETED
- Matrix：TM-P0-003、TM-P0-006
- Goal：用 `httptest` 固定 `/health`、`/ready` 的状态码和响应语义。
- Allowed Files：
  - `internal/transport/http/*_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/transport/http -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - `/health`、`/ready` 正常和异常路径均有断言。
  - 不启动真实 HTTP server。
- Completion Evidence：
  - `internal/transport/http/router_test.go` 已新增。
  - `/health` 断言 HTTP 200、响应 `code=0`、`message=success`、`data.status=ok`。
  - `/ready` 覆盖数据库缺失、ping 失败、ping 成功三条路径，并断言 HTTP 状态、`data.status` 和 `data.checks.database`。
  - `go test ./internal/transport/http -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-P1-004：增加 demo CRUD 测试基线

- Status：COMPLETED
- Matrix：TM-P0-005、TM-P0-006
- Goal：为 demo Todo 示例建立 Create/List/Get/Update/Delete 测试基线。
- Allowed Files：
  - `internal/modules/demo/**/*_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/modules/demo/... -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - 使用临时 SQLite 或等价隔离数据库。
  - 不依赖真实外部服务。
- Completion Evidence：
  - `internal/modules/demo/service/todo_test.go` 已新增。
  - 测试使用临时 SQLite 和真实 repository/service，不依赖真实外部服务。
  - `TestTodoServiceCRUD` 覆盖 Create/List/Get/Update/Delete 成功路径和软删除后的不可见语义。
  - `TestTodoServiceValidationAndNotFound` 覆盖空标题校验、缺失 Get/Update/Delete 的 not found 语义。
  - `go test ./internal/modules/demo/... -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-P1-005：明确 demo 迁移边界

- Status：COMPLETED
- Matrix：TM-P1-001、TM-P0-006
- Goal：收拢 demo `AutoMigrate`、`initdb` 和 reload 的触发策略。
- Allowed Files：
  - `internal/app/**/*`
  - 迁移边界文档
  - 项目状态文档
- Verification：
  - `go test ./internal/app/... -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] dev/demo 与生产/bootstrap 迁移职责已记录到 `docs/specs/demo_migration_boundary.md`。
  - [CONFIRMED] 自动迁移触发点已由 `internal/app/initapp/demo_migration_test.go` 验证。
- Completion Evidence：
  - `DemoMigrationPolicyFor` 固定 `server-start`、`initdb`、`reload` 三类触发策略。
  - `NewModules` 使用 `DemoMigrationTriggerServerStart`，`BuildInitDB` 使用 `DemoMigrationTriggerInitDB`。
  - `reloadDatabase` 使用 `DemoMigrationTriggerReload`，策略为跳过 demo `AutoMigrate`。
  - `go test ./internal/app/... -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-P1-006：收拢 `cmd/server tests` 命令语义

- Status：COMPLETED
- Matrix：TM-P1-002、TM-P0-006
- Goal：让 `tests` 命令名、描述或行为与真实用途一致。
- Allowed Files：
  - `cmd/server/*`
  - CLI 相关文档
  - 项目状态文档
- Verification：
  - `go test ./cmd/server -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] 命令不再误导为 Go test：`tests` 现在执行 `go test`。
  - [CONFIRMED] 命令注册或行为有最小测试。
- Completion Evidence：
  - `cmd/server/tests.go` 已移除 yaml2go 示例转换逻辑。
  - `TestsCommand` 默认执行 `go test ./...`，并支持 `--package/-p` 指定测试范围。
  - `cmd/server/tests_test.go` 已覆盖命令元信息、默认包范围、指定包范围和失败返回。
  - `go test ./cmd/server -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-P1-007：完成 `pkg/*` 公共/内部分类

- Status：COMPLETED
- Matrix：TM-P1-003、TM-P0-006
- Goal：逐包标注公共 API、内部支撑或待确认定位。
- Allowed Files：
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - 包 README
  - 项目状态文档
- Verification：
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] 每个 `pkg/*` 包均有定位。
  - [CONFIRMED] 破坏性重构仍需单独任务确认。
- Completion Evidence：
  - 13 个 `pkg/*/README.md` 已新增 API 分类段。
  - `ARCHITECTURE.md` 已新增 `pkg/*` API 分类表。
  - `MODULES.md` 已更新 `pkg/*` API 分类和风险状态。
  - `go test ./... -count=1` 通过。

### TASK-P1-008：标注 `pkg/sqlgen` 未实现能力

- Status：COMPLETED
- Matrix：TM-P1-004、TM-P0-006
- Goal：把 `pkg/sqlgen` TODO/unsupported 能力边界从隐含状态改为明确状态。
- Allowed Files：
  - `pkg/sqlgen/*`
  - 包 README
  - 项目状态文档
- Verification：
  - `go test ./pkg/sqlgen -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - [CONFIRMED] 未实现能力不再暗示可用。
  - [CONFIRMED] 涉及代码行为的 unsupported 路径已有测试覆盖。
- Completion Evidence：
  - `ErrCodeUnsupportedOperation` 和 `NewUnsupportedError` 已新增。
  - `Or`、`Not`、`Group`、`Having`、`Distinct`、`Joins` 不再静默 no-op，后续生成 SQL 时返回 unsupported 错误。
  - `DeleteInBatches` 不再退化为普通删除，直接返回 unsupported 错误。
  - `ReverseDB(...).Generate`、`GenerateAll`、`GenerateToDir` 返回 unsupported 错误。
  - `pkg/sqlgen/README.md` 已标注 unsupported / partial 能力边界。
  - `go test ./pkg/sqlgen -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-P1-009：明确 `types/*` 契约边界

- Status：COMPLETED
- Matrix：TM-P1-005、TM-P0-006
- Goal：明确 `types/constants`、`types/errors`、`types/result` 和根 `types` 聚合入口的公共契约定位，尤其标注 `types/result` 属于 HTTP/Gin 响应契约而非纯类型包。
- Source：
  - 用户选择 A。
  - `BL-021`：明确 `types/result` 是否属于 HTTP 契约。
  - `TM-P1-005`：`types/result`、错误码、跨层契约边界清晰。
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
- Verification：
  - `go test ./types/... -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] `types/result` 的 HTTP/Gin 契约定位被文档化。
  - [CONFIRMED] `types/errors` 中 auth/rbac 预留错误码与当前未实现范围的关系被标注。
  - [CONFIRMED] `types/constants` 和根 `types` 聚合入口的公共/跨层边界被标注。
  - [CONFIRMED] 如新增或修改行为测试，相关 `types` 包测试和全量回归通过。
  - [CONFIRMED] 状态、测试报告、变更日志、风险、Backlog 和交接文档已更新。
- Completion Evidence：
  - `types/result/result.go` 包注释已明确 `types/result` 是 HTTP API 响应契约，其中 Gin helper 属于 HTTP/Gin 适配层。
  - `types/doc.go`、`types/constants/doc.go`、`types/errors/doc.go` 已标注根 `types`、常量和错误码契约边界。
  - `docs/specs/types_contract_boundary.md` 已记录 `types/*` 包边界、依赖说明、auth/rbac 非目标和验收证据。
  - `types/result/result_test.go` 已覆盖响应结构、分页总页数、Gin helper 的 HTTP 状态码和错误码映射、trace id 提取。
  - `types/errors/error_test.go` 已覆盖 `BizError` 错误链和错误码分段。
  - `go test ./types/... -count=1` 和 `go test ./... -count=1` 均通过。

### TASK-NEXT-SCOPE：确认下一阶段范围

- Status：COMPLETED
- Matrix：BL-021、TM-P1-005
- Goal：确认已完成 P1 列表之后的下一步：提升 Backlog 项、补测试、进入 Phase 6 收尾，或结束本轮。
- Allowed Files：
  - 项目状态文档
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户选择 A：提升 `BL-021` / `TM-P1-005`。
  - [CONFIRMED] 新的任务 TASK-P1-009 和时间切片 TS-P1-009 已写入 `TASKS.md` 和 `TIME_SLICES.md`。
  - [CONFIRMED] 当前合法任务已推进为 TASK-P1-009。

### TASK-NEXT-SCOPE-002：确认 `types/*` 契约边界后的后续范围

- Status：COMPLETED
- Matrix：BL-022、TM-P1-006
- Goal：确认 TASK-P1-009 完成之后的下一步：提升 `BL-020` 补 `pkg/*` 行为测试、进入 Phase 6 收尾，或结束本轮。
- Allowed Files：
  - 项目状态文档
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户提出架构修正：`pkg/plugin` 不应主动注册插件服务，应被动由插件服务注册。
  - [CONFIRMED] 新的任务 TASK-P1-010 和时间切片 TS-P1-010 已写入 `TASKS.md` 和 `TIME_SLICES.md`。

### TASK-P1-010：收拢 `pkg/plugin` 被动注册边界

- Status：COMPLETED
- Matrix：TM-P1-006
- Goal：落实用户修正：`pkg/plugin` 只作为被动 registry/runtime，不主动从配置加载并注册插件服务；插件服务或宿主装配层显式构造插件并调用注册接口。
- Source：
  - 用户修正：`pkg/plugin不应该主动注册插件服务，而是被动由插件服务进行注册`。
  - User Correction Review：ACCEPT_WITH_RISK。
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
- Verification：
  - `go test ./pkg/plugin -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] `Manager` 的公共 API 不再暴露主动配置加载/本地 factory 装配入口。
  - [CONFIRMED] local/http 插件由插件服务或宿主装配层显式创建后调用 `Register` 注册。
  - [CONFIRMED] README、架构、模块清单和决策记录已说明被动注册边界。
  - [CONFIRMED] `pkg/plugin` 包测试和全量回归通过。
- Completion Evidence：
  - `Manager` 接口移除 `Load`、`RegisterLocalFactory` 和 manager option 主动装配公共面。
  - 新增 `NewHTTP` 和 HTTP option，让远程插件可由插件服务构造后注册。
  - local/http 测试已改为先构造插件实例，再调用 `mgr.Register`。
  - `pkg/plugin` README 和 package doc 已说明被动注册边界。
  - `go test ./pkg/plugin -count=1`、`go test ./... -count=1` 和 `git diff --check` 均通过。

### TASK-NEXT-SCOPE-003：确认 `pkg/plugin` 被动注册边界后的后续范围

- Status：COMPLETED
- Matrix：BL-020、TM-P1-003、TM-P1-007
- Goal：确认 TASK-P1-010 完成之后的下一步：提升 `BL-020` 补 `pkg/*` 行为测试、进入 Phase 6 收尾，或结束本轮。
- Allowed Files：
  - 项目状态文档
- Verification：
  - 无需 Go 测试，除非用户确认进入新的代码或测试切片。
- Exit Conditions：
  - [CONFIRMED] 用户明确选择 A：提升 `BL-020` 补 `pkg/*` 行为测试。
  - [CONFIRMED] 首批任务 TASK-P1-011 和时间切片 TS-P1-011 已写入 `TASKS.md` 和 `TIME_SLICES.md`。
  - [CONFIRMED] 当前合法下一步不再是待确认状态。

### TASK-P1-011：补首批无外部服务依赖 `pkg/*` 行为测试

- Status：COMPLETED
- Matrix：TM-P1-003、TM-P1-007、TM-P0-006
- Goal：为当前无包级测试且无外部服务依赖的公共 `pkg/*` 包补最小行为测试，首批覆盖 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go`。
- Source：
  - 用户选择 A。
  - `BL-020`：为无测试的公共 `pkg/*` 包补最小行为测试。
  - `TM-P1-003`：`pkg/*` API 分类与后续测试缺口。
- Priority：P1
- Complexity：Medium
- Conservative Estimate：1 个时间切片；若新增测试暴露实现缺陷，同一问题最多修复 3 轮。
- Dependencies：
  - TASK-P1-007：`pkg/*` API 分类已完成。
  - TASK-P1-010：`pkg/plugin` 被动注册边界已完成。
- Inputs：
  - `ARCHITECTURE.md` 中 `pkg/*` API 分类。
  - `MODULES.md` 中无测试包清单。
  - `TEST_MATRIX.md` 的 `TM-P1-003` 和 `TM-P1-007`。
- Outputs：
  - `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 的最小行为测试。
  - 状态、验收、测试报告、变更日志、问题记录和交接更新。
- Allowed Files：
  - `pkg/cli/**/*_test.go`
  - `pkg/i18n/**/*_test.go`
  - `pkg/yaml2go/**/*_test.go`
  - 若测试暴露当前范围内的真实实现缺陷，可修改 `pkg/cli/*.go`、`pkg/i18n/*.go`、`pkg/yaml2go/*.go`，但必须记录修复原因和验证证据。
  - 项目状态文档。
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他 `pkg/*` 包。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Steps：
  1. 阅读 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 当前公开 API 和 README 分类。
  2. 为每个包选择最小、确定性、无外部服务依赖的行为路径。
  3. 新增包级测试，不改变公共 API。
  4. 运行相关包测试、全量回归和 diff 检查。
  5. 更新状态、验收、测试报告、变更日志、问题记录和交接。
- Test Commands：
  - `gofmt -w pkg/cli/*_test.go pkg/i18n/*_test.go pkg/yaml2go/*_test.go`
  - `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Acceptance Criteria：
  - [CONFIRMED] `pkg/cli` 至少覆盖命令/flag 或错误行为的稳定公共路径。
  - [CONFIRMED] `pkg/i18n` 至少覆盖翻译加载、默认语言或错误路径中的稳定公共行为。
  - [CONFIRMED] `pkg/yaml2go` 至少覆盖一个成功转换路径和一个错误路径。
  - [CONFIRMED] 测试不依赖 Redis、数据库、真实 HTTP server、生产配置或外部网络。
  - [CONFIRMED] 相关包测试和全量回归通过。
- Completion Criteria：
  - 测试文件已新增或更新。
  - 修改范围符合本任务。
  - 所有验证命令已执行并记录。
  - 无新增未记录失败项。
  - 状态文档、测试报告、变更记录和交接说明已更新。
- Failure Handling：
  - 同一失败最多修复 3 轮。
  - 如果失败来自本切片新增测试覆盖的当前范围缺陷，在允许文件内修复并记录。
  - 如果失败来自其他包或历史未确认缺陷，登记 `ISSUES.md` 和 `RISK_REGISTER.md`，不得扩大范围。
- Evidence：
  - 修改文件：`pkg/cli/app_test.go`、`pkg/i18n/i18n_test.go`、`pkg/yaml2go/converter_test.go`、`pkg/yaml2go/converter_impl.go`、`pkg/yaml2go/method_generator.go`、`pkg/yaml2go/utils.go`、项目状态文档。
  - 命令：`gofmt -w pkg/cli/app_test.go pkg/i18n/i18n_test.go pkg/yaml2go/converter_test.go`；`go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 验证结论：`pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 均已有最小行为测试；`pkg/yaml2go` 生成 tag 和方法 import 顺序缺陷已修复。
- Next Task：
  - TASK-NEXT-SCOPE-004：确认继续 `BL-020` 下一批、进入 Phase 6 收尾，或结束本轮。

### TASK-NEXT-SCOPE-004：确认首批 `pkg/*` 行为测试完成后的后续范围

- Status：COMPLETED
- Matrix：BL-020、TM-P1-003、TM-P1-007
- Goal：确认 TASK-P1-011 完成之后的下一步：继续 `BL-020` 后续 `pkg/*` 行为测试批次、进入 Phase 6 收尾，或结束本轮。
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
  - [CONFIRMED] TASK-P1-012 和 TS-P1-012 已写入 `TASKS.md` 与 `TIME_SLICES.md`。
  - [CONFIRMED] 当前合法下一步不再是待确认状态。

### TASK-P1-012：补第二批 `pkg/*` 行为测试

- Status：COMPLETED
- Matrix：BL-020、TM-P1-003、TM-P1-008、TM-P0-006
- Source：
  - 用户发送“下一步”，确认继续 `BL-020`。
  - `MODULES.md`：`pkg/executor`、`pkg/httpserver`、`pkg/storage` 仍无包级测试。
- Priority：P1
- Type：测试；若新增测试暴露当前三包内缺陷，可做最小修复。
- Goal：为不依赖 Redis、数据库或外部网络服务的第二批公共 `pkg/*` 包补最小行为测试，覆盖稳定成功路径和明确错误路径。
- Allowed Files：
  - `pkg/executor/**/*_test.go`
  - `pkg/httpserver/**/*_test.go`
  - `pkg/storage/**/*_test.go`
  - 必要时限当前三包实现文件：`pkg/executor/*.go`、`pkg/httpserver/*.go`、`pkg/storage/*.go`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`、`TEST_MATRIX.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`ISSUES.md`、`RISK_REGISTER.md`、`BACKLOG.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - `pkg/cache/**/*`
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他无关 `pkg/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Non-Goals：
  - 不接入真实 Redis、数据库、第三方网络服务或生产配置。
  - 不补 `pkg/cache` Redis 行为测试。
  - 不重构 HTTP 路由、业务 handler、storage 对外 API 或 executor 公共接口。
  - 不实现 httpserver 文档中尚未落地的 executor 注入能力。
- Execution Steps：
  1. 阅读 `pkg/executor`、`pkg/httpserver`、`pkg/storage` 源码和 README，确认可测公共行为。
  2. 新增最小包级测试：executor 覆盖配置校验、任务执行、缺失池、关闭和 panic handler；httpserver 覆盖 `New`、配置默认/校验、停止态 reload/shutdown；storage 覆盖内存文件系统读写、复制、MIME、Excel、图片和配置错误路径。
  3. 若新增测试暴露当前三包内缺陷，只做最小修复。
  4. 运行格式化、相关包测试、全量回归和 diff 检查。
  5. 更新验收、测试报告、变更、问题、风险、Backlog 和交接文档。
- Verification：
  - `gofmt -w pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`
  - `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] `pkg/executor` 有确定性最小行为测试。
  - [CONFIRMED] `pkg/httpserver` 有确定性最小行为测试，且不依赖固定生产端口。
  - [CONFIRMED] `pkg/storage` 有确定性最小行为测试，使用内存文件系统。
  - [CONFIRMED] 本切片验证命令通过，失败已按修复上限记录并修复。
  - [CONFIRMED] 状态、变更、测试报告、问题和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`pkg/executor/executor_test.go`、`pkg/httpserver/httpserver_test.go`、`pkg/storage/storage_test.go`、`pkg/executor/constants.go`、`pkg/executor/manager.go`、`pkg/executor/pool.go`、项目状态文档。
  - 命令：`gofmt -w pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`；`gofmt -w pkg/executor/constants.go pkg/executor/manager.go pkg/executor/pool.go pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`；`go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 修复记录：首次相关包测试暴露 `pkg/executor` sentinel 错误包装和 panic handler 未调用缺陷；第一轮修复后仅测试存在任务完成等待竞态；第二轮修正测试等待后通过。
- Next Task：
  - TASK-NEXT-SCOPE-005：确认继续 `pkg/cache` 等剩余 `BL-020` 范围、进入 Phase 6 收尾，或结束本轮。

### TASK-NEXT-SCOPE-005：确认第二批 `pkg/*` 行为测试完成后的后续范围

- Status：COMPLETED
- Matrix：BL-020、TM-P1-003、TM-P1-008
- Goal：确认 TASK-P1-012 完成之后的下一步：继续 `BL-020` 剩余 `pkg/*` 行为测试、进入 Phase 6 收尾，或结束本轮。
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
  - [CONFIRMED] 用户选择 A：继续 `BL-020` 剩余 `pkg/*` 行为测试。
  - [CONFIRMED] 新任务 TASK-P1-013 和时间切片 TS-P1-013 已写入 `TASKS.md` 与 `TIME_SLICES.md`。
  - [CONFIRMED] 当前合法下一步不再是待确认状态。

### TASK-P1-013：补第三批 `pkg/cache` 隔离行为测试

- Status：COMPLETED
- Matrix：BL-020、TM-P1-003、TM-P1-009、TM-P0-006
- Source：
  - 用户选择 A，确认继续 `BL-020` 剩余范围。
  - `ARCHITECTURE.md` 和 `MODULES.md`：`pkg/cache` 是公共基础设施 API，Redis 依赖路径缺少隔离测试。
- Priority：P1
- Type：测试；若新增测试暴露 `pkg/cache` 当前范围内缺陷，可做最小修复。
- Goal：为 `pkg/cache` 补最小隔离行为测试，覆盖配置校验、Redis 基本操作、批量操作、计数器、缺失键、重载失败保持旧连接和重载成功切换连接。
- Allowed Files：
  - `pkg/cache/**/*_test.go`
  - 必要时限当前包实现文件：`pkg/cache/*.go`
  - 如隔离 Redis 需要纯测试依赖，可修改 `go.mod`、`go.sum`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`、`TEST_MATRIX.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`ISSUES.md`、`RISK_REGISTER.md`、`BACKLOG.md`、`DECISIONS.md`、`ROADMAP.md`、`PROJECT_BRIEF.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他无关 `pkg/*`
  - 数据库 schema、部署配置、真实密钥文件
- Non-Goals：
  - 不连接真实 Redis 服务。
  - 不引入生产运行依赖或修改生产配置。
  - 不重构 `Cache` 公共接口。
  - 不实现 Memcached、本地缓存或分布式锁语义扩展。
- Steps：
  1. 阅读 `pkg/cache` 源码和 README，确认稳定公共行为。
  2. 使用进程内隔离 Redis 测试服务或等价策略覆盖成功路径，不依赖外部 Redis。
  3. 新增配置、输入错误和 Redis 行为测试；必要时做当前包内最小修复。
  4. 运行格式化、`pkg/cache` 包测试、全量回归和 diff 检查。
  5. 更新验收、测试报告、变更、问题、风险、Backlog、决策和交接文档。
- Verification：
  - `gofmt -w pkg/cache/*_test.go`
  - `go test ./pkg/cache -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] `pkg/cache` 有确定性最小行为测试。
  - [CONFIRMED] 测试不依赖真实 Redis、外部网络服务、生产配置或数据库。
  - [CONFIRMED] 相关包测试和全量回归通过。
  - [CONFIRMED] 状态、变更、测试报告、问题和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`pkg/cache/cache_test.go`、`go.mod`、`go.sum`、项目状态文档。
  - 命令：`go get github.com/alicebob/miniredis/v2@latest`；`gofmt -w pkg/cache/cache_test.go`；`go test ./pkg/cache -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 修复记录：首次包测试为测试代码编译失败，原因是误读 `miniredis.Get` 返回值；修正测试断言后通过。
- Next Task：
  - TASK-NEXT-SCOPE-006：确认进入 Phase 6 收尾、提升 `pkg/utils` 等内部支撑测试，或结束本轮。

### TASK-NEXT-SCOPE-006：确认 `pkg/cache` 行为测试完成后的后续范围

- Status：COMPLETED
- Matrix：BL-020、TM-P1-003、TM-P1-009
- Goal：确认 TASK-P1-013 完成之后的下一步：进入 Phase 6 收尾、补 `pkg/utils` 等内部支撑测试，或结束本轮。
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
  - [CONFIRMED] 用户选择 B：提升 `pkg/utils` 等内部支撑测试。
  - [CONFIRMED] 新任务 TASK-P1-014 和时间切片 TS-P1-014 已写入 `TASKS.md` 与 `TIME_SLICES.md`。

### TASK-P1-014：补 `pkg/utils` 内部支撑工具最小行为测试

- Status：COMPLETED
- Matrix：BL-023、TM-P1-010、TM-P0-006
- Source：
  - 用户选择 B，确认提升内部支撑测试。
  - `BL-023`：为 `pkg/utils` 内部支撑工具补最小测试。
  - `MODULES.md`：`pkg/utils` 被分类为内部支撑工具包，能力较杂，默认 Snowflake panic 策略需确认。
- Priority：P1
- Type：测试；若新增测试暴露 `pkg/utils` 当前范围内缺陷，可做最小修复。
- Goal：为 `pkg/utils` 补最小确定性行为测试，覆盖 Snowflake、监听地址校验、端口查找、设备 ID 稳定性和 i18n helper 默认语言委托语义。
- Allowed Files：
  - `pkg/utils/**/*_test.go`
  - 必要时限当前包实现文件：`pkg/utils/*.go`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`、`TEST_MATRIX.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`ISSUES.md`、`RISK_REGISTER.md`、`BACKLOG.md`、`DECISIONS.md`、`ROADMAP.md`、`PROJECT_BRIEF.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - 其他无关 `pkg/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Non-Goals：
  - 不修改 `pkg/utils` 公共 API。
  - 不改变默认 Snowflake panic 策略。
  - 不依赖真实外部网络服务、固定端口、生产配置或机器专属断言。
  - 不补 `internal/*`、middleware、router 或 app 集成测试。
- Steps：
  1. 阅读 `pkg/utils` 源码和 README，确认确定性可测行为。
  2. 新增最小包级测试，优先覆盖无外部依赖和可隔离网络绑定路径。
  3. 若新增测试暴露当前包缺陷，只做最小修复并记录。
  4. 运行格式化、`pkg/utils` 包测试、全量回归和 diff 检查。
  5. 更新验收、测试报告、变更、问题、风险、Backlog、决策和交接文档。
- Verification：
  - `gofmt -w pkg/utils/*_test.go`
  - `go test ./pkg/utils -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] `pkg/utils` 有确定性最小行为测试。
  - [CONFIRMED] 测试不依赖真实外部网络服务、固定生产端口、数据库或生产配置。
  - [CONFIRMED] 相关包测试和全量回归通过。
  - [CONFIRMED] 状态、变更、测试报告、问题和交接文档已更新。
  - [CONFIRMED] 下一合法任务明确。
- Evidence：
  - 修改文件：`pkg/utils/utils_test.go`、项目状态文档。
  - 覆盖内容：Snowflake 生成和非法 node、监听地址校验、端口范围和 exclude、设备 ID 稳定/盐值、i18n helper 默认语言转发。
  - 命令：`gofmt -w pkg/utils/utils_test.go`；`go test ./pkg/utils -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 修复记录：初始测试中占用端口断言受 Windows/Go 绑定语义影响不稳定；改为确定性无效地址、端口范围和 exclude 断言后通过。
- Next Task：
  - TASK-NEXT-SCOPE-007：确认进入 Phase 6 收尾、提升 app/router/middleware 等集成测试，或结束本轮。

### TASK-NEXT-SCOPE-007：确认 `pkg/utils` 内部支撑测试完成后的后续范围

- Status：COMPLETED
- Matrix：BL-023、TM-P1-010、TM-P0-006
- Goal：确认 TASK-P1-014 完成之后的下一步：进入 Phase 6 收尾、继续补 app/router/middleware 等集成测试，或结束本轮。
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
  - [CONFIRMED] 新的唯一合法任务 TASK-P1-015 和时间切片 TS-P1-015 已写入状态文件。

### TASK-P1-015：补 app/router/middleware 最小集成测试

- Status：COMPLETED
- Matrix：BL-002、TM-P0-005、TM-P0-006
- Source：
  - 用户回复 `b`，确认提升 app/router/middleware 等集成测试。
  - `BL-002`：增加 app/router/demo 集成测试。
  - `MODULES.md`：`internal/middleware` 中间件链路无测试，`internal/transport/http` demo 路由缺少 integration 测试，`internal/modules/demo` handler/router 集成仍未覆盖。
- Priority：P1
- Type：测试；若新增测试暴露当前 router/middleware/demo handler 范围内缺陷，可做最小修复。
- Goal：用 `httptest` 固定 demo Todo 路由注册、handler/service/repository HTTP 关键路径，以及 TraceID、CORS、Recovery 等中间件链路，不启动真实 HTTP server。
- Allowed Files：
  - `internal/transport/http/**/*_test.go`
  - `internal/middleware/**/*_test.go`
  - `internal/modules/demo/**/*_test.go`
  - 必要时限当前范围实现文件：`internal/transport/http/*.go`、`internal/middleware/*.go`、`internal/modules/demo/handler/*.go`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`、`TEST_MATRIX.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`ISSUES.md`、`RISK_REGISTER.md`、`BACKLOG.md`、`DECISIONS.md`、`ROADMAP.md`、`PROJECT_BRIEF.md`、`MODULES.md`、`ARCHITECTURE.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - `cmd/**/*`
  - `pkg/**/*`，除既有接口只读使用外不得修改
  - `types/**/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Non-Goals：
  - 不启动真实 HTTP server 或占用固定端口。
  - 不接入真实外部数据库、Redis、第三方服务或生产配置。
  - 不重构 app 装配、router、middleware、demo handler 或 service API。
  - 不实现 auth/rbac、生产迁移框架、CI/CD 或 Phase 6 收尾。
- Steps：
  1. 阅读 router、middleware、demo handler/service/repository 源码，确认最小可测链路。
  2. 新增或扩展 `internal/transport/http` 集成测试，使用临时 SQLite 和真实 demo repository/service/handler。
  3. 覆盖 demo route create/list/get/update/delete 或最小关键路径，并断言 trace id、CORS 和 recovery 响应语义。
  4. 若测试暴露当前范围缺陷，只做最小修复并记录。
  5. 运行格式化、相关包测试、全量回归和 diff 检查。
  6. 更新验收、测试报告、变更、问题、风险、Backlog、决策和交接文档。
- Verification：
  - `gofmt -w internal/transport/http/*_test.go internal/middleware/*_test.go internal/modules/demo/**/*_test.go`
  - `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] demo Todo router/handler/service/repository HTTP 集成路径被 `httptest` 覆盖。
  - [CONFIRMED] TraceID、CORS 和 Recovery 中间件链路有最小路由级断言。
  - [CONFIRMED] 测试不依赖真实 HTTP server、固定生产端口、外部数据库、Redis 或生产配置。
  - [CONFIRMED] 相关包测试、全量回归和 diff 检查通过。
  - [CONFIRMED] 状态、变更、测试报告、问题和交接文档已更新。
- Evidence：
  - 修改文件：`internal/transport/http/router_integration_test.go`、项目状态文档。
  - 覆盖内容：demo Todo HTTP Create/List/Get/Update/Delete、删除后 404、CORS preflight/actual origin header、TraceID header round-trip、Recovery 500 响应 traceId 和 logger 调用。
  - 命令：`gofmt -w internal/transport/http/router_integration_test.go`；`go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - 修复记录：前两次相关包测试失败来自测试构造问题：`httptest.NewRequest` 默认 Host 与 Origin 同源，导致 CORS 中间件跳过；固定测试 Host 为 `api.local` 后通过。
- Next Task：
  - TASK-NEXT-SCOPE-008：确认进入 Phase 6 收尾、继续 app 装配/reload/config 等剩余集成测试，或结束本轮。

### TASK-NEXT-SCOPE-008：确认 app/router/middleware 集成测试后的后续范围

- Status：COMPLETED
- Matrix：BL-002、TM-P0-004、TM-P0-006
- Goal：确认 TASK-P1-015 完成之后的下一步：进入 Phase 6 收尾、继续补 app 装配/reload/config 等剩余集成测试，或结束本轮。
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
  - [CONFIRMED] 新的唯一合法任务和切片 TASK-PHASE6-001 / TS-PHASE6-001 已写入状态文件。

### TASK-PHASE6-001：Phase 6 收尾与交接

- Status：COMPLETED
- Matrix：TM-P0-006
- Source：
  - 用户最新回复 `a`，确认 TASK-NEXT-SCOPE-008 选项 A。
  - TASK-P1-015 已完成并通过验证。
- Priority：P0
- Type：收尾、验证、交接。
- Goal：冻结本轮项目优化成果，更新状态、验收、测试报告、变更记录、风险/Backlog 和交接说明，并运行最终验证命令。
- Allowed Files：
  - `STATUS.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `ACCEPTANCE.md`
  - `TEST_MATRIX.md`
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `ISSUES.md`
  - `RISK_REGISTER.md`
  - `BACKLOG.md`
  - `DECISIONS.md`
  - `ROADMAP.md`
  - `PROJECT_BRIEF.md`
  - `MODULES.md`
  - `ARCHITECTURE.md`
  - `AGENT_HANDOFF.md`
- Forbidden Files：
  - Go 源码和测试文件。
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件。
- Verification：
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] Phase 6 收尾状态写入项目状态文档。
  - [CONFIRMED] 最终验证命令已执行并记录：`go test ./... -count=1` 与 `git diff --check` 均通过。
  - [CONFIRMED] `AGENT_HANDOFF.md` 明确无自动下一实现任务；后续工作需要用户重新确认。
- Evidence：
  - 修改文件：项目状态文档、验收、测试报告、变更记录、风险/Backlog、决策记录和交接说明。
  - 验证：`go test ./... -count=1` PASS；`git diff --check` PASS，仅有 Windows LF/CRLF 转换警告。
  - 结论：Phase 6 收尾完成；当时 app 装配、reload/config 等剩余集成路径保留为后续确认范围，后续已由 TASK-P1-016 覆盖并关闭。

### TASK-P1-016：补 app 装配与 reload/config 剩余集成测试

- Status：COMPLETED
- Matrix：BL-002、TM-P0-004、TM-P0-006、TM-P1-012、RISK-008
- Source：
  - 用户明确要求实施 TASK-P1-016 计划。
  - `BL-002` 剩余范围中 app 装配、配置变更 hook 与 reload 路径尚未覆盖。
  - `TASK-PHASE6-001` 已记录该范围保留为后续确认事项。
- Priority：P1
- Type：测试；若新增测试暴露当前 `internal/app/**` 或必要 `internal/config/**` 范围内缺陷，只做最小修复。
- Goal：新增 app 装配与 reload/config 集成测试，覆盖 server/initdb 模式最小装配链路、配置变更 hook、reload 分发与关闭配置路径。
- Allowed Files：
  - `internal/app/app_integration_test.go`
  - `internal/app/reloadapp/reload_test.go`
  - 必要时限当前范围修复文件：`internal/app/**`、`internal/config/**`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`ISSUES.md`、`DECISIONS.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - `cmd/**/*`
  - `pkg/**/*`，除既有接口只读使用外不得修改
  - `types/**/*`
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Non-Goals：
  - 不新增导出业务 API、配置 schema、HTTP 路由或数据库 schema。
  - 不启动真实 HTTP server。
  - 不依赖 Redis/MySQL/Postgres/外部网络。
  - 不重构 app 装配、reload、config 或 demo 模块。
- Steps：
  1. 更新当前任务与时间切片状态，确认唯一合法范围。
  2. 新增 `internal/app/app_integration_test.go`，用临时 YAML、临时 SQLite 和真实 `app.New` 覆盖 server/initdb/config hook 路径。
  3. 新增 `internal/app/reloadapp/reload_test.go`，用 fake 组件覆盖 reload 分发、关闭配置和 database reload 不隐式迁移路径。
  4. 运行格式化、相关包测试、全量回归和 diff 检查。
  5. 更新验收、测试报告、变更、风险、Backlog、问题和交接文档。
- Verification：
  - `gofmt -w internal/app/app_integration_test.go internal/app/reloadapp/reload_test.go`
  - `go test ./internal/app/... -count=1`
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] server 模式真实 app 装配链路有测试覆盖，且资源可 `Shutdown`。
  - [CONFIRMED] initdb 模式只初始化数据库并创建 demo schema，不装配 HTTP transport。
  - [CONFIRMED] `ConfigManager.Update` 能触发 app hook 并更新 `Core.Config`，不启动真实 server。
  - [CONFIRMED] reload 分发逻辑覆盖未变化、变化、关闭 Redis/executor/storage 和 database reload 不隐式迁移路径。
  - [CONFIRMED] 相关包测试、全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接说明已更新。
- Evidence：
  - 修改文件：`internal/app/app_integration_test.go`、`internal/app/reloadapp/reload_test.go`、项目状态文档。
  - 覆盖内容：真实 `app.New` server/initdb 装配、demo schema 创建、app 配置变更 hook、reload 组件分发、可选组件关闭置空、database reload 不触发 demo schema 隐式迁移。
  - 命令：`gofmt -w internal/app/app_integration_test.go internal/app/reloadapp/reload_test.go`；`go test ./internal/app/... -count=1`；`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - Next Task：NONE；后续工作需用户重新确认并建立新的任务/时间切片。

### TASK-P1-017：分阶段中文化 `pkg/*` README

- Status：COMPLETED
- Matrix：BL-006、RISK-005、TM-P1-013、TM-P0-006
- Source：
  - TASK-P1-016 完成后，用户选择 A。
  - `BL-006`：分阶段中文化包 README。
  - `RISK-005`：包 README 中英混杂。
- Priority：P1
- Type：文档；不新增功能、不修改代码。
- Goal：将第一阶段包 README 中文化范围限定为 `pkg/*/README.md`，统一包标题、说明、边界、风险和许可证等面向阅读者的中文表达，同时保留 Go 标识符、配置键、命令、代码示例和必要英文技术名词。
- Allowed Files：
  - `pkg/*/README.md`
  - `REQUIREMENTS.md`
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`ISSUES.md`、`DECISIONS.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - `cmd/**/*`
  - `internal/**/*`
  - `types/**/*`
  - `pkg/**/*` 中除 `README.md` 以外的文件
  - `go.mod`
  - `go.sum`
  - 数据库 schema、部署配置、真实密钥文件
- Non-Goals：
  - 不改 Go 代码、配置 schema、HTTP 路由、数据库 schema 或依赖。
  - 不补新的行为测试。
  - 不重写全部历史文档或模板。
  - 不把 API 标识符、代码示例、协议名、环境变量名强行翻译。
- Steps：
  1. 审查当前 `pkg/*/README.md` 的中英文混杂点和过期风险描述。
  2. 翻译或收拢标题、段落、边界说明、风险说明和许可证等读者可见文本。
  3. 保留代码示例和 API 名称，避免文档与实现不一致。
  4. 运行全量回归和 diff 空白检查。
  5. 更新验收、测试报告、变更、Backlog、风险、问题和交接文档。
- Verification：
  - `go test ./... -count=1`
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] 13 个 `pkg/*/README.md` 已检查；当前阶段主要读者文本已中文化或保留为必要技术名词。
  - [CONFIRMED] README 中与已完成测试状态明显冲突的风险描述已同步。
  - [CONFIRMED] 未修改 Go 代码、依赖、配置 schema、HTTP 路由或数据库 schema。
  - [CONFIRMED] 全量回归和 diff 检查通过。
  - [CONFIRMED] 状态文档、测试报告和交接说明已更新。
- Evidence：
  - 修改文件：`pkg/cache/README.md`、`pkg/cli/README.md`、`pkg/database/README.md`、`pkg/executor/README.md`、`pkg/httpserver/README.md`、`pkg/i18n/README.md`、`pkg/logger/README.md`、`pkg/plugin/README.md`、`pkg/sqlgen/README.md`、`pkg/storage/README.md`、`pkg/utils/README.md`、`pkg/yaml2go/README.md`、需求/架构/模块和项目状态文档。
  - `pkg/crypto/README.md` 已检查，当前标题和主体已符合第一阶段中文化要求，未产生内容修改。
  - 命令：`go test ./... -count=1`；`git diff --check`。
  - 测试结果：PASS；`git diff --check` 仅有 Windows LF/CRLF 转换警告。
  - Next Task：NONE；后续工作需用户重新确认并建立新的任务/时间切片。

## P2 任务

### TASK-P2-001：补 CI 质量门禁与部署说明

- Status：COMPLETED
- Matrix：REQ-OPT-P2-003、BL-007、BL-008、TM-P2-001、RISK-016
- Source：
  - 用户选择 D，确认进入 CI/CD 与部署方向。
  - `BL-007`：增加 CI 质量门禁。
  - `BL-008`：增加部署说明。
- Priority：P2
- Type：质量工程+文档；不执行真实部署。
- Goal：新增非生产 CI 质量门禁和手动部署说明，固定发布前最小检查、配置入口和当前部署边界。
- Allowed Files：
  - `.github/workflows/ci.yml`
  - `docs/deployment.md`
  - `README.md`
  - 项目状态文档：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`REQUIREMENTS.md`、`ARCHITECTURE.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`ISSUES.md`、`DECISIONS.md`、`ROADMAP.md`、`PROJECT_BRIEF.md`、`AGENT_HANDOFF.md`
- Forbidden Files：
  - Go 源码和测试文件。
  - `go.mod`、`go.sum`。
  - 数据库 schema、真实生产配置、真实 `.env`、密钥或部署凭据。
  - 会连接远程环境、推送镜像或执行部署的脚本。
- Non-Goals：
  - 不实现真实 CD。
  - 不新增 Dockerfile、镜像仓库发布、Kubernetes、systemd 或云平台部署模板。
  - 不处理生产密钥。
  - 不执行生产命令或远程部署。
- Verification：
  - gofmt 漂移审计：KNOWN_DRIFT（历史 Go 文件格式漂移，未在本切片批量格式化，已进入 `BL-025`）
  - `go test ./... -count=1`：PASS
  - `go build -o <temp> ./cmd/server`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- Exit Conditions：
  - [CONFIRMED] CI workflow 只执行只读质量门禁，不使用 secrets、不推送镜像、不连接远程环境。
  - [CONFIRMED] 部署说明记录配置入口、发布前检查、手动启动、initdb 边界和未实现的真实 CD 项。
  - [CONFIRMED] README 有 CI 与部署说明入口。
  - [CONFIRMED] 本地等价验证通过。
  - [CONFIRMED] 状态、变更、测试报告、问题和交接文档已更新。
- Evidence：
  - 修改文件：`.github/workflows/ci.yml`、`docs/deployment.md`、`README.md` 和项目状态文档。
  - CI 内容：checkout、setup-go、gofmt 漂移报告（非阻塞）、`go test ./... -count=1 -mod=readonly`、server 构建、`git diff --check`。
  - 部署说明内容：发布前检查、配置入口、手动运行、initdb 边界、手动发布步骤和未实现项。
  - Next Task：NONE；真实 CD、镜像发布和远程部署自动化需用户重新确认并建立新的任务/时间切片。


### TASK-NEXT-SCOPE-010：确认真实 CD / 镜像发布 / 远程部署自动化边界

- Status：COMPLETED
- Time Slice：TS-NEXT-SCOPE-010
- Source：用户选择 C；`BL-024`；`REQ-OPT-P2-003` 后续范围。
- Goal：在不实现真实部署、不读取密钥、不连接远程环境的前提下，确认后续真实 CD 自动化的安全边界和实现输入。
- Review Result：NEEDS_USER_DECISION。
- Required Decisions：
  - 镜像仓库与镜像命名策略。
  - 远程部署已确认；仍需确认是否采用 SSH 到 Linux 服务器 + Docker Compose 作为默认运行方式。
  - 发布环境、触发方式和审批策略。
  - GitHub Secrets 名称和权限边界。
  - 是否允许首个实现切片只做 staging/manual dry-run，而不碰 production。
- Allowed Files：
  - `STATUS.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `REQUIREMENTS.md`
  - `ARCHITECTURE.md`
  - `TEST_MATRIX.md`
  - `ACCEPTANCE.md`
  - `BACKLOG.md`
  - `RISK_REGISTER.md`
  - `DECISIONS.md`
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `ISSUES.md`
  - `AGENT_HANDOFF.md`
- Forbidden Files：
  - `.github/workflows/*`（确认前不得实现真实 CD）
  - Go 源码和测试文件
  - `go.mod`、`go.sum`
  - Dockerfile、Kubernetes、systemd、云平台部署模板
  - 真实 `.env`、密钥、部署凭据或生产配置
- Verification：
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] 用户确认使用远程部署，并确认使用 `.env` 风格文件配置。
  - [CONFIRMED] 已创建并完成 TASK-P2-002 / TS-P2-002。
  - [CONFIRMED] 确认前不实现镜像推送、远程部署、生产发布或密钥读取。

### TASK-P2-002：补远程部署 env 配置模板

- Status：COMPLETED
- Time Slice：TS-P2-002
- Source：用户要求“远程部署 .env 来配置”。
- Goal：提供可提交的远程部署显式参数契约和说明，同时确保真实 显式部署参数 不会进入 Git。
- Allowed Files：
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `.gitignore`
  - `README.md`
  - `docs/deployment.md`
  - 项目状态文档
- Forbidden Files：
  - 真实 `.env` / 显式部署参数
  - `.github/workflows/*` 自动部署实现
  - Go 源码和测试文件
  - `go.mod`、`go.sum`
  - Dockerfile、Kubernetes、systemd、云平台部署模板
  - 真实密钥、服务器地址、部署凭据或生产配置
- Verification：
  - `git diff --check`
- Evidence：
  - 新增 `deploy.sh` / `script/install.sh` 显式参数契约，仅包含占位值和说明。
  - 删除旧本地部署 env 文件依赖，部署配置改由显式参数传入。
  - `docs/deployment.md` 增加远程部署变量说明。
  - `README.md` 增加模板入口。
  - 未实现真实部署 workflow、未连接服务器、未读取 secrets。

### TASK-P2-003：实现手动远程部署 workflow

- Status：COMPLETED
- Time Slice：TS-P2-003
- Matrix：BL-024、TM-P2-004、RISK-016、RISK-017
- Source：用户明确回复“确认实现远程部署 workflow”。
- Priority：P2
- Type：CI/CD 配置与文档；不执行真实部署。
- Goal：新增手动触发的 GitHub Actions 远程部署 workflow，读取 GitHub Secrets 中的 显式部署参数 内容和 SSH 密钥，通过 SSH 到远程 Linux 主机执行 Docker Compose 拉取镜像、重启服务和健康检查。
- Allowed Files：
  - `.github/workflows/deploy-remote.yml`
  - `deploy.sh` / `script/install.sh` 显式参数契约
  - `README.md`
  - `docs/deployment.md`
  - 项目状态文档
- Forbidden Files：
  - 真实 `.env` / 显式部署参数
  - Go 源码和测试文件
  - `go.mod`、`go.sum`
  - 数据库 schema、真实服务器地址、真实 token、SSH 私钥、密码或生产配置
  - Dockerfile、Kubernetes、systemd、云平台部署模板
- Non-Goals：
  - 不在本地或当前会话执行 GitHub workflow。
  - 不连接远程服务器、不推送镜像、不执行真实部署。
  - 不实现生产数据库迁移框架。
  - 不新增业务功能或修改 Go API。
- Verification：
  - workflow YAML 语法/结构检查
  - `git diff --check`
- Exit Conditions：
  - [CONFIRMED] workflow 只通过 `workflow_dispatch` 手动触发，默认 staging，不自动生产发布。
  - [CONFIRMED] workflow 使用 GitHub Secrets，不在仓库写入真实部署值。
  - [CONFIRMED] workflow 能从 显式部署参数 secret 解析 显式部署参数 形状，并校验必需变量。
  - [CONFIRMED] workflow 使用 SSH 执行 Docker Compose pull/up 和 health/ready 检查。
  - [CONFIRMED] 部署说明记录需要配置的 Secrets、远程主机前置条件和手动触发步骤。
  - [CONFIRMED] 验证命令通过，状态文档、测试报告和交接说明已更新。
- Evidence：
  - 新增 `.github/workflows/deploy-remote.yml`。
  - workflow 使用 `workflow_dispatch`，需要 `confirm=deploy`，且当前只允许 `staging`。
  - workflow 读取 显式部署参数、`DEPLOY_SSH_KEY`、可选 `DEPLOY_SSH_KNOWN_HOSTS`、可选 `GHCR_USERNAME` / `GHCR_TOKEN`。
  - workflow 从 GitHub Variables/Secrets 组装显式部署参数，通过 SSH 执行 `script/install.sh` / `deploy.sh`，并由远程脚本执行 Docker build/pull、Compose up 与 health/ready 检查。
  - `docs/deployment.md` 已补 Secrets、远程主机前置条件和手动触发步骤。
  - 验证：临时 Go YAML 解析 PASS；`go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml` PASS；`git diff --check` PASS。
  - 未执行真实部署、未连接远程服务器、未写入真实密钥、未推送镜像、未修改 Go 代码。

### TASK-P2-004：补 Linux Docker production 部署制品

- Status：COMPLETED
- Time Slice：TS-P2-004
- Matrix：BL-024、TM-P2-005、RISK-016、RISK-017、RISK-018
- Source：用户要求“开始，linux、docker、production -> 部署”。
- Priority：P2
- Type：发布工程配置+文档；不执行真实 production 部署。
- Goal：新增可提交的 Linux Docker 运行制品、production Compose 示例和统一 `deploy.sh` 部署入口，并把远程部署 workflow 从 staging-only 扩展为手动 staging/production，但 production 必须显式选择 GitHub Environment 并输入环境绑定确认词。
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
- Forbidden Files：
  - 真实 `.env` / 显式部署参数
  - Go 源码和测试文件
  - `go.mod`、`go.sum`
  - 数据库 schema、真实服务器地址、真实 token、SSH 私钥、密码或生产配置
  - 自动触发 production 部署的 workflow
  - 生产迁移框架
- Non-Goals：
  - 不在当前会话触发 GitHub workflow。
  - 不连接远程服务器、不推送镜像、不执行 staging 或 production 部署。
  - 不新增业务功能或修改 Go API。
  - 不引入 Kubernetes、systemd 或云平台模板。
- Verification：
  - `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .`：PASS_REMOTE，用户在 Linux Docker 环境补跑通过，BuildKit 输出 `23/23 FINISHED`，镜像标记为 `docker.io/library/go-scaffold:local`。
  - 临时 Go YAML 解析：PASS。
  - `go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml`：PASS。
  - `bash -n deploy.sh`：FAIL_ENV，本机无可用 bash，WSL 未安装 Linux 发行版。
  - `shfmt` Bash 语法解析：PASS。
  - `go test ./... -count=1`：PASS。
  - `go build -o <temp> ./cmd/server`：PASS。
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告。
- Exit Conditions：
  - [CONFIRMED] Dockerfile 镜像构建已在 Linux Docker 环境验证通过。
  - [CONFIRMED] production Compose 示例使用 `DEPLOY_IMAGE`、只挂载示例路径，并包含 healthcheck。
  - [CONFIRMED] production 配置样例绑定 `0.0.0.0:9999`，不包含真实密钥。
  - [CONFIRMED] 远程 Linux 部署脚本会按参数/环境变量动态生成 `运行期显式部署参数`，不要求提交真实 显式部署参数。
  - [CONFIRMED] workflow 仍只支持 `workflow_dispatch`，staging/production 均需要环境绑定确认词。
  - [CONFIRMED] production 发布依赖 GitHub Environment `production` 和 Secrets，不保存真实值。
  - [CONFIRMED] 部署说明记录 Linux 主机、Docker Compose、目录权限、production 手动触发和回滚边界。
  - [CONFIRMED] 除本机 `bash -n` 因环境缺失不可执行外，Docker build、脚本 Bash 语法解析、YAML 解析、actionlint、Go 回归、server build 和 diff 检查均已通过或已有等价验证记录。
- Evidence：
  - 修改文件：`Dockerfile`、`.dockerignore`、`deploy/docker-compose.production.example.yml`、`deploy/config.production.example.yaml`、`deploy.sh`、`.github/workflows/deploy-remote.yml`、`deploy.sh` / `script/install.sh` 显式参数契约、README、部署说明和项目状态文档。
  - Docker build：用户在 Linux Docker 环境执行 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .`，构建 214.8s 完成，BuildKit 输出 `23/23 FINISHED`，镜像 sha256 为 `4df5520bcf1c45a922be8db2e6c5e58ae8fc025f34bea5f1d4bf33f0b2301785`。
  - 其他验证：脚本 Bash 语法解析 PASS；YAML 解析 PASS；actionlint PASS；`go test ./... -count=1` PASS；server build PASS；`git diff --check` PASS。
  - Next Task：当前无自动下一实现任务。真实 production 运行、镜像发布流水线和生产迁移框架仍需单独确认。

### TASK-P2-005：实现 `pkg/plugin/hooks` 独立钩子引擎

- Status：COMPLETED
- Time Slice：TS-P2-005
- Source：用户要求实现 `dev.tmp/new-plugin.md` 设计；用户修正审查结论为 `ACCEPT_WITH_RISK`。
- Priority：P2
- Type：公共基础设施 API + 测试
- Goal：新增不依赖 logger/config/IAM/internal 的可复用钩子引擎，为插件运行时、远程钩子和 app 组合根接入提供基础。
- Allowed Files：`pkg/plugin/hooks/**/*`、`pkg/plugin/hooks/*`、项目状态文档。
- Forbidden Files：真实 `.env`、密钥、部署凭据、数据库 schema、JWT 中间件、生产权限系统。
- Verification：`go test ./pkg/plugin/hooks -count=1`；`go test ./pkg/plugin/... -count=1`。
- Exit Conditions：处理器按优先级从高到低执行；执行前复制处理器列表；每个处理器前检查 context；支持停止语义；拒绝 nil handler。
- Evidence：新增 `pkg/plugin/hooks`，包含 `Point`、`Event`、`Result`、`Handler`、`HandlerFunc`、`Registry` 和服务查找能力；相关测试通过。

### TASK-P2-006：让 `pkg/plugin.Manager` 支持钩子

- Status：COMPLETED
- Time Slice：TS-P2-006
- Goal：扩展 manager 钩子 API，同时保持被动注册边界和既有 `NewManager()` 兼容。
- Allowed Files：`pkg/plugin/**/*`、项目状态文档。
- Verification：`go test ./pkg/plugin/... -count=1`。
- Exit Conditions：`before_invoke` 可阻止插件调用；`invoke_error` 不覆盖原错误；`after_invoke` 错误以包装错误返回；`pkg/plugin` 不导入 `pkg/iam` 或 `internal/*`。
- Evidence：`Manager` 新增 `Hooks()`、`RegisterHook` 和 `WithHooks`；标准钩子点已加入，调用错误钩子按尽力通知处理，包边界保持与 IAM/internal 解耦。

### TASK-P2-007：实现 HTTP 远程插件服务端和 `RemoteHook`

- Status：COMPLETED
- Time Slice：TS-P2-007
- Goal：沿用现有 JSON `Request`/`Response` 协议，新增 HTTP server helper、标准操作和远程钩子适配器。
- Allowed Files：`pkg/plugin/**/*`、项目状态文档。
- Verification：`go test ./pkg/plugin/... -count=1`。
- Exit Conditions：只接受 `POST /plugin/v1/invoke`；非法请求、插件错误和响应大小限制有测试；`hooks.execute` 可解码为 `hooks.Result`。
- Evidence：新增 `NewHTTPServer(plugin Plugin) http.Handler`、`OperationHooksExecute` 和 `RemoteHook`；沿用 JSON `Request`/`Response` 协议并覆盖错误路径。

### TASK-P2-008：实现 `pkg/iam` 与内存实现

- Status：COMPLETED
- Time Slice：TS-P2-008
- Goal：新增独立 IAM 公共 API，覆盖主体、凭证、策略、认证、授权、上下文和 memory service。
- Allowed Files：`pkg/iam/**/*`、项目状态文档。
- Verification：`go test ./pkg/iam/... -count=1`。
- Exit Conditions：内存实现支持 token 凭证、精确和 `*` 通配策略、拒绝优先、过期检查、启用后默认拒绝。
- Evidence：新增 `pkg/iam` 公共类型、上下文 helper 与 `pkg/iam/memory`；memory 测试覆盖 token、通配、拒绝优先、过期和默认拒绝。

### TASK-P2-009：接入配置与应用组装

- Status：COMPLETED
- Time Slice：TS-P2-009
- Goal：新增 `plugin`、`iam` 配置，server 模式装配 IAM 和插件管理器，配置创建的插件仅为 HTTP 适配器。
- Allowed Files：`internal/config/**/*`、`internal/app/**/*`、`pkg/plugin/**/*`、`pkg/iam/**/*`、项目状态文档。
- Verification：`go test ./internal/config ./internal/app/... ./pkg/plugin/... ./pkg/iam/... -count=1`。
- Exit Conditions：本地插件仍由代码显式注册；钩子绑定创建 `RemoteHook`；权限检查钩子只在 `internal/app` 注册。
- Evidence：配置新增 `plugin` 与 `iam`，默认 disabled；`internal/app/initapp.Infrastructure` 新增 `IAM` 与 `Plugins`，配置插件仅创建 HTTP adapter，IAM 授权钩子只在 app 组合层注册。

### TASK-P2-010：实现 reload、生命周期与收尾验证

- Status：COMPLETED
- Time Slice：TS-P2-010
- Goal：在配置重载和应用关闭中安全处理 IAM/plugin 基础设施，并完成全量验证和交接。
- Allowed Files：`internal/app/**/*`、项目状态文档、必要文档。
- Verification：`go test ./internal/config ./internal/app/... -count=1`；`go test ./... -count=1`；`go build -o <temp> ./cmd/server`；`git diff --check`。
- Exit Conditions：重载先构建新实例再替换，失败保留旧实例；关闭时 HTTP server 停止后、cache/database 前关闭插件管理器；最终状态和交接文档一致。
- Evidence：reload 先构建新 IAM/plugin 基础设施再替换；关闭顺序已在 HTTP server 后、cache/database 前关闭插件管理器；目标包测试、全量回归、server build 和 diff 检查通过。

## 历史任务

### TASK-HIST-PLUGIN-001：实现独立插件系统 v1

- Status：COMPLETED
- Summary：历史任务，已完成 `pkg/plugin` local/http 能力。
- Verification：
  - `go test ./pkg/plugin -count=1`
  - `go test ./... -count=1`

### TASK-HIST-PLUGIN-002：确认插件系统 v1 API 边界

- Status：COMPLETED
- Summary：历史任务，v1 local/http API 已接受。
- Follow-up：
  - rpc/ws/discovery/examples 保留在 Backlog，不属于当前主线。
