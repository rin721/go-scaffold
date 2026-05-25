# AGENT_RULES.md

## Core Rule

项目事实必须沉淀在仓库文件中，不依赖聊天上下文。长版规则见 `docs/ai/prompt.md`，跨工具入口见 `AGENTS.md`。

## Startup Rule

执行任何任务前必须读取：

1. `AGENTS.md`
2. `AGENT_RULES.md`
3. `STATUS.md`
4. `TASKS.md`
5. `TIME_SLICES.md`
6. `REQUIREMENTS.md`
7. `ARCHITECTURE.md`
8. `ACCEPTANCE.md`
9. `ISSUES.md`
10. `TEST_REPORT.md`
11. `AGENT_HANDOFF.md`

如果必读文件缺失或互相冲突，先生成 `docs/reports/status_diagnostics/` 诊断报告，不得直接写业务代码。

## Next Step Rule

用户发送“下一步”时：

1. 从 `STATUS.md` 找当前任务和时间切片。
2. 用 `TASKS.md` 和 `TIME_SLICES.md` 交叉确认。
3. 如果状态是 `BLOCKED`，先处理阻塞或请求决策。
4. 如果状态是 `PENDING_VERIFICATION`，先验证。
5. 如果状态是 `REWORK_REQUIRED`，先返工。
6. 如果状态是 `PENDING_USER_CONFIRMATION`，只按已记录的默认规则推进，或等待用户确认。
7. 如果状态是 `COMPLETED`，推进到下一个明确的合法切片。
8. 否则只执行当前合法时间切片。

## Completion Rule

不得在缺少产物、验证、状态更新、测试报告、变更记录或交接说明时标记 `COMPLETED`。测试未运行只能是 `PENDING_VERIFICATION` 或 `BLOCKED`；测试失败只能是 `REWORK_REQUIRED` 或 `BLOCKED`。

## Scope Rule

只修改当前时间切片允许的文件。发现额外问题时，记录到 `BACKLOG.md`、`ISSUES.md` 或 `RISK_REGISTER.md`，不得插队实现。

## Repair Rule

同一问题最多修复 3 轮。仍失败时记录问题报告，并标记 `BLOCKED` 或 `REWORK_REQUIRED`。

## Documentation Rule

每次任务状态变化后至少更新：

- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `CHANGELOG.md`
- `TEST_REPORT.md`
- `AGENT_HANDOFF.md`

如果出现问题，更新 `ISSUES.md`；如果出现风险，更新 `RISK_REGISTER.md`；如果出现新想法，更新 `BACKLOG.md`。

## Skill Rule

项目专用技能以 `skills/*/SKILL.md` 为 canonical source。`.agents/skills/*/SKILL.md` 只作为轻量适配入口，避免双份长内容漂移。

## Handoff Rule

每个任务结束时，`AGENT_HANDOFF.md` 必须足以让新 Agent 无聊天上下文接手。
