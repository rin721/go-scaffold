# AGENT_HANDOFF.md

## 最近更新

- 日期：2026-05-25
- Agent：Codex
- 工具：Codex Desktop

## 项目快照

- 项目：go-scaffold
- 阶段：demo CRUD 测试基线
- 模块：internal/modules/demo
- 当前任务：TASK-P1-004
- 当前时间切片：TS-P1-004
- 总体状态：IN_PROGRESS

## 最近完成

- 按用户计划重新启动“全项目分析与优化路线”主线。
- 生成/重写六个中文启动模板。
- 更新核心项目文档和状态文件，避免后续“下一步”继续指向插件系统扩展。
- 保留插件系统 v1 为历史内容和 Backlog，不删除历史。
- 执行 `go test ./... -count=1`，结果通过。
- 用户发送“下一步”后，已按推荐默认值确认项目优化路线和关键边界。
- 新增 `ROADMAP.md`。
- 再次收到“下一步”后，已生成 `MODULES.md`，完成模块职责、边界冲突、测试矩阵草案和 P1 优化候选项。
- 更新任务、时间切片、状态、测试报告、变更日志和交接文档，将合法下一步推进到 TASK-OPT-004。
- 本轮收到“下一步”后，已生成 `TEST_MATRIX.md`，并把 P1 任务草案与 P1 时间切片草案写入 `TASKS.md` 和 `TIME_SLICES.md`。
- 已补齐 `ISSUES.md` 作为问题记录入口。
- 已执行 `go test ./... -count=1`，结果通过；未修改 Go 文件。
- 用户再次发送“下一步”，已按推荐默认顺序确认 P1 执行顺序。
- 已完成 TASK-P1-001：修复 `copyConfig` 字段覆盖，并新增配置 copy/update 测试。
- 已执行 `go test ./internal/config -count=1` 和 `go test ./... -count=1`，结果均通过。
- 用户确认补齐 Prompt 全量 Agent 基础设施，已完成 TASK-INFRA-001。
- 已新增 `AGENTS.md`、`CLAUDE.md`、`AGENT_RULES.md`、`SKILLS.md`、缺失模板、reports/specs、跨工具目录和 14 个项目 skills。
- 已执行 Prompt 全量产物存在性核对和 `go test ./... -count=1`，结果均通过。
- 用户要求实施 Agent 基础设施一致性修复计划，已完成 TASK-INFRA-002。
- 已修复 `AGENTS.md` 实际缺失但状态文件声称存在的冲突。
- 已扩充 14 个 canonical `skills/*/SKILL.md`，新增 14 个 `.agents/skills/*/SKILL.md` 适配器。
- 已标准化 `docs/templates/*`，新增状态诊断报告。
- 已执行文件存在性核对、`quick_validate.py` skill 验证、跨工具入口引用一致性检查和 `go test ./... -count=1`，结果均通过。
- 用户再次发送“下一步”，已完成 TASK-P1-002：统一配置环境变量策略。
- 数据库环境变量现在以 `DB_*` 为主，旧 `REI_APP_DB_*` 保留为兼容 fallback。
- `.env.example` 已移除未实现的 JWT 示例，并补齐 Storage/CORS 示例。
- 已执行 `go test ./internal/config -count=1` 和 `go test ./... -count=1`，结果均通过。
- 用户再次发送“下一步”，已完成 TASK-P1-003：增加 health/ready 与 router smoke test。
- 新增 `internal/transport/http/router_test.go`，覆盖 `/health`、`/ready` 数据库缺失、ping 失败、ping 成功路径。
- 已执行 `go test ./internal/transport/http -count=1` 和 `go test ./... -count=1`，结果均通过。

## 最近变更文件

| 文件 | 用途 |
|---|---|
| `PROJECT_BRIEF.md` | 记录项目当前目标、优势、问题和已确认事项 |
| `REQUIREMENTS.md` | 记录确认后的优化需求 |
| `ARCHITECTURE.md` | 记录确认后的高层架构原则和需分析模块 |
| `ACCEPTANCE.md` | 记录启动阶段和后续代码任务验收门禁 |
| `RISK_REGISTER.md` | 记录范围、文档、迁移、包 API、测试等风险 |
| `BACKLOG.md` | 保留未确认优化项和插件扩展项 |
| `DECISIONS.md` | 保留插件历史决策并新增当前主线切换决策 |
| `TASKS.md` | 当前合法任务推进到 TASK-P1-004，TASK-P1-003 已完成 |
| `TIME_SLICES.md` | 当前合法切片推进到 TS-P1-004，TS-P1-003 已完成 |
| `STATUS.md` | 当前合法下一步切换为 TASK-P1-004 |
| `TEST_REPORT.md` | 记录 TASK-P1-003 HTTP router 测试和全量测试通过 |
| `CHANGELOG.md` | 记录 TASK-P1-003 health/ready smoke test |
| `ROADMAP.md` | 记录治理优先优化路线 |
| `MODULES.md` | 记录模块职责、边界冲突、测试矩阵草案和 P1 优化候选项 |
| `TEST_MATRIX.md` | 记录正式测试矩阵、任务草案和推荐执行顺序 |
| `ISSUES.md` | 补齐失败和阻塞问题记录入口 |
| `AGENTS.md` | 新增实际缺失的跨 Agent 主入口 |
| `internal/config/manager.go` | 修复 `copyConfig` 字段覆盖和 slice 深拷贝 |
| `internal/config/manager_test.go` | 新增配置 copy/update 和环境变量策略测试 |
| `internal/config/app_database.go` | 数据库环境变量改为 `DB_*` 优先，旧前缀 fallback |
| `internal/config/constants.go` | 记录旧前缀兼容策略 |
| `.env.example` | 对齐环境变量策略，移除 JWT 示例，补充 Storage/CORS |
| `internal/transport/http/router_test.go` | 新增 health/ready router smoke test |
| `docs/templates/*` | 标准化为可复用模板 |
| `CLAUDE.md` / `AGENT_RULES.md` / `SKILLS.md` | 统一跨 Agent 入口、规则和 skills 索引 |
| `skills/*/SKILL.md` | 扩充 14 个 canonical 项目专用 skills |
| `.agents/skills/*/SKILL.md` | 新增 14 个轻量适配器 |
| `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` | 记录并关闭 `AGENTS.md` 缺失冲突 |
| `docs/reports/*` / `docs/specs/*` | 新增报告与规格目录入口 |
| `.agents/*` / `.cursor/*` / `.kiro/*` / `.codex/*` | 新增跨工具适配入口 |

## 最近执行命令

| 命令 | 结果 |
|---|---|
| Agent 基础设施文件存在性核对 | PASS |
| `quick_validate.py` 验证 28 个 skill 目录 | PASS |
| 跨工具入口引用一致性检查 | PASS |
| `go test ./internal/config -count=1` | PASS |
| `go test ./internal/transport/http -count=1` | PASS |
| `go test ./... -count=1` | PASS |
| `git diff --check` | PASS，只有 Windows CRLF 转换警告 |

## 测试状态

- [CONFIRMED] `internal/config` 测试通过。
- [CONFIRMED] `internal/transport/http` 测试通过。
- [CONFIRMED] 全量 Go 测试通过。
- [RISK] `cmd/server`、`internal/app`、`internal/modules/demo` 等关键路径仍无测试文件。

## 当前阻塞项

| ID | 阻塞项 | 需要动作 |
|---|---|---|
|  |  |  |

## 重要决策

- [CONFIRMED] 本轮不写 Go 代码。
- [CONFIRMED] 当前主线切回项目治理与优化路线。
- [CONFIRMED] 插件系统 v1 保留为历史记录。
- [CONFIRMED] 已确认“治理优先”路线。
- [CONFIRMED] 已确认 `pkg/*` 混合策略。
- [CONFIRMED] 已确认 demo 长期标准示例。
- [CONFIRMED] 已确认迁移 dev-prod 分层。
- [CONFIRMED] 已确认中文化根文档和模板优先。
- [CONFIRMED] 已确认 P1 推荐执行顺序。
- [CONFIRMED] TASK-P1-001 已完成。
- [CONFIRMED] TASK-INFRA-001 已完成。
- [CONFIRMED] Prompt 全量 Agent 基础设施产物已补齐。
- [CONFIRMED] TASK-INFRA-002 已完成。
- [CONFIRMED] `AGENTS.md` 缺失冲突已修复。
- [CONFIRMED] 28 个 skill 目录已通过 `quick_validate.py`。
- [CONFIRMED] TASK-P1-002 已完成。
- [CONFIRMED] 配置环境变量策略统一为无全局前缀，数据库旧前缀仅兼容。
- [CONFIRMED] TASK-P1-003 已完成。
- [CONFIRMED] `/health` 与 `/ready` 状态码和响应语义已被 `httptest` 固定。

## 合法下一步

- 执行 TASK-P1-004：增加 demo CRUD 测试基线。
- 允许修改 `internal/modules/demo/**/*_test.go` 和项目状态文档。
- 目标是用隔离测试固定 demo Todo Create/List/Get/Update/Delete 关键路径。

## 禁止事项

- 不修改当前时间切片未授权的 Go 代码。
- 不重构项目结构。
- 不实现 Backlog 项。
- 不继续插件系统 rpc/ws/discovery。
- 不执行部署或不可逆迁移。

## 恢复说明

1. 先读 `STATUS.md`。
2. 再读 `TASKS.md` 和 `TIME_SLICES.md`。
3. 如用户只说“下一步”，执行 TASK-P1-004。
4. TASK-P1-004 允许修改 `internal/modules/demo/**/*_test.go` 和状态文档，必须运行 `go test ./internal/modules/demo/... -count=1` 和 `go test ./... -count=1`。
