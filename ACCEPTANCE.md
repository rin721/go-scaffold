# ACCEPTANCE.md

## 验收状态

- Project：go-scaffold
- Phase：测试矩阵与任务拆分
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
- 测试矩阵与任务拆分：NOT_STARTED
- 代码实现：BLOCKED
