# ACCEPTANCE.md

## 验收状态

- Project：go-scaffold
- Phase：HTTP health/ready smoke test
- Status：IN_PROGRESS
- Last Updated：2026-05-25

## 本轮启动验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-OPT-001 | 中文项目启动模板已生成 | 检查 `docs/templates/project_start_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-002 | 中文需求澄清模板已生成 | 检查 `docs/templates/requirements_clarification_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-003 | 中文技术方案模板已生成 | 检查 `docs/templates/technical_options_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-004 | 中文架构约束模板已生成 | 检查 `docs/templates/architecture_constraints_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-005 | 中文验收模板已生成 | 检查 `docs/templates/acceptance_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-006 | 中文风险确认模板已生成 | 检查 `docs/templates/risk_confirmation_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-007 | 当前主线切换到项目优化启动确认 | 检查 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-OPT-008 | 本轮不修改 Go 代码 | 检查 git diff | 是 | [CONFIRMED] |
| ACC-OPT-009 | 全量 Go 测试通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## 文档验收

- [CONFIRMED] 当前输出优先使用中文。
- [CONFIRMED] 历史插件系统内容保留，但不再作为当前合法任务。
- [CONFIRMED] 当前风险和待确认项已进入模板和状态文件。
- [CONFIRMED] 后续“下一步”应先处理用户确认，而不是进入代码实现。

## 后续代码任务验收门禁

后续任何代码优化任务只有同时满足以下条件，才能标记为 `COMPLETED`：

1. [CONFIRMED] 已映射到确认后的需求。
2. [CONFIRMED] 已映射到确认后的架构边界。
3. [CONFIRMED] 已映射到唯一任务和唯一时间切片。
4. [CONFIRMED] 修改范围没有超出时间切片。
5. [CONFIRMED] 相关测试或验证命令已执行。
6. [CONFIRMED] 文档、状态、测试报告、变更记录和交接说明已更新。
7. [CONFIRMED] 下一步合法任务明确。

## 已确认验收

| ID | 验收项 | 结果 | 状态 |
|---|---|---|---|
| ACC-FUTURE-001 | 优化路线被确认 | 治理优先 | [CONFIRMED] |
| ACC-FUTURE-002 | `pkg/*` API 策略被确认 | 混合策略 | [CONFIRMED] |
| ACC-FUTURE-003 | demo 模块定位被确认 | 长期标准示例 | [CONFIRMED] |
| ACC-FUTURE-004 | 迁移策略被确认 | dev-prod 分层 | [CONFIRMED] |
| ACC-FUTURE-005 | 中文化范围被确认 | 根文档和模板优先，历史内容分阶段处理 | [CONFIRMED] |

## 当前完成判断

- 启动模板生成：COMPLETED
- 项目优化启动确认：COMPLETED
- 模块边界清单：COMPLETED
- 测试矩阵与任务拆分：COMPLETED
- P1 执行顺序确认：COMPLETED
- 配置 copy/update 测试与修复：COMPLETED
- 配置环境变量策略收拢：COMPLETED
- HTTP health/ready smoke test：COMPLETED
- Agent 基础设施补齐：COMPLETED
- Agent 基础设施一致性修复：COMPLETED
- 代码实现：IN_PROGRESS

## Prompt 全量产物验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-INFRA-001 | 跨 Agent 入口文件已补齐 | 检查 `AGENTS.md`、`CLAUDE.md` | 是 | [CONFIRMED] |
| ACC-INFRA-002 | Agent 规则与 skills 索引已补齐 | 检查 `AGENT_RULES.md`、`SKILLS.md` | 是 | [CONFIRMED] |
| ACC-INFRA-003 | 任务拆分和时间切片模板已补齐 | 检查 `docs/templates/task_decomposition_template.md`、`docs/templates/time_slice_template.md` | 是 | [CONFIRMED] |
| ACC-INFRA-004 | reports/specs 目录入口已补齐 | 检查 `docs/reports/*`、`docs/specs/*` | 是 | [CONFIRMED] |
| ACC-INFRA-005 | 14 个项目专用 skills 已补齐 | 检查 `skills/*/SKILL.md` | 是 | [CONFIRMED] |
| ACC-INFRA-006 | 跨工具目录入口已补齐 | 检查 `.agents`、`.cursor`、`.kiro`、`.codex` | 是 | [CONFIRMED] |
| ACC-INFRA-007 | 文档基础设施补齐不破坏 Go 测试 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-INFRA-008 | 本任务未新增 Go 代码修改 | 核对变更范围 | 是 | [CONFIRMED] |
| ACC-INFRA-009 | `AGENTS.md` 实际存在且跨工具引用闭合 | 文件存在性核对和引用一致性检查 | 是 | [CONFIRMED] |
| ACC-INFRA-010 | canonical 与 `.agents` skills 均可通过 skill 验证 | `quick_validate.py` 验证 28 个 skill 目录 | 是 | [CONFIRMED] |
| ACC-INFRA-011 | 模板不混入当前项目实例事实 | 检查 `docs/templates/*` | 是 | [CONFIRMED] |
| ACC-INFRA-012 | 状态冲突已形成诊断报告 | 检查 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` | 是 | [CONFIRMED] |

## 测试矩阵与任务拆分验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-OPT-010 | 正式测试矩阵已生成 | 检查 `TEST_MATRIX.md` | 是 | [CONFIRMED] |
| ACC-OPT-011 | P1 任务草案已生成 | 检查 `TASKS.md` 中 TASK-P1-001 至 TASK-P1-008 | 是 | [CONFIRMED] |
| ACC-OPT-012 | P1 时间切片草案已生成 | 检查 `TIME_SLICES.md` 中 TS-P1-001 至 TS-P1-008 | 是 | [CONFIRMED] |
| ACC-OPT-013 | 每个 P1 任务有允许文件范围 | 检查 `TASKS.md` 和 `TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-OPT-014 | 每个 P1 任务有验证命令和退出条件 | 检查 `TASKS.md` 和 `TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-OPT-015 | 当前仍未修改 Go 代码 | 检查 git diff | 是 | [CONFIRMED] |

## 下一步确认验收

- [CONFIRMED] 推荐执行顺序已接受：`TEST_MATRIX.md` 中从 TASK-P1-001 到 TASK-P1-008。
- [CONFIRMED] 用户再次发送“下一步”已按推荐默认顺序视为确认，并已完成 TASK-P1-001。

## TASK-P1-001 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-001 | `copyConfig` 不丢失关键字段 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-002 | `copyConfig` 对 slice 做深拷贝 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-003 | `Update` 保留未修改字段 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-004 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-002 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-005 | 数据库环境变量主策略为 `DB_*` | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-006 | 旧 `REI_APP_DB_*` 仍作为兼容 fallback | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-007 | `DB_*` 优先级高于旧前缀 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-008 | `.env.example` 与实现一致且不再暗示 JWT 已实现 | 人工检查 `.env.example` | 是 | [CONFIRMED] |
| ACC-P1-009 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-003 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-010 | `/health` HTTP 200 和响应语义被固定 | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-011 | `/ready` 数据库缺失路径返回 503 和 `not_ready` | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-012 | `/ready` 数据库 ping 失败路径返回 503 和错误语义 | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-013 | `/ready` 数据库 ping 成功路径返回 200 和 `ready` | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-014 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |
