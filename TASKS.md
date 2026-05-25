# TASKS.md

## 当前合法任务

- Task ID：TASK-P1-004
- Status：NOT_STARTED
- Time Slice：TS-P1-004
- Summary：HTTP health/ready smoke test 已完成。当前合法下一步为 TASK-P1-004。

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

- Status：NOT_STARTED
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

### TASK-P1-005：明确 demo 迁移边界

- Status：NOT_STARTED
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
  - dev/demo 与生产/bootstrap 迁移职责被明确记录。
  - 自动迁移触发点可验证。

### TASK-P1-006：收拢 `cmd/server tests` 命令语义

- Status：NOT_STARTED
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
  - 命令不再误导为 Go test。
  - 命令注册或行为有最小测试。

### TASK-P1-007：完成 `pkg/*` 公共/内部分类

- Status：NOT_STARTED
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
  - 每个 `pkg/*` 包均有定位。
  - 破坏性重构仍需单独任务确认。

### TASK-P1-008：标注 `pkg/sqlgen` 未实现能力

- Status：NOT_STARTED
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
  - 未实现能力不再暗示可用。
  - 如涉及代码行为，必须有测试覆盖。

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
