# Agent Project Driver

本仓库使用工程级 AI Agent 项目驱动流程。任何 Agent 在执行任务前，都必须把仓库文档作为事实来源，而不是依赖聊天上下文。

长版规范位于 `docs/ai/prompt.md`。本文件是跨工具主入口，适用于 Codex、Claude Code、Cursor、Kiro、Antigravity 以及其他 coding agent。

## Required Reading Before Work

执行任何实现、修复、验证或状态推进前，必须读取：

1. `AGENT_RULES.md`
2. `STATUS.md`
3. `TASKS.md`
4. `TIME_SLICES.md`
5. `REQUIREMENTS.md`
6. `ARCHITECTURE.md`
7. `ACCEPTANCE.md`
8. `ISSUES.md`
9. `TEST_REPORT.md`
10. `AGENT_HANDOFF.md`

如果上述文件缺失、互相冲突或无法确定当前唯一合法任务，不得编造下一步；必须先生成状态诊断报告并修复 Agent 基础设施。

## Current Task Discipline

- 只执行 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` 共同指向的当前唯一合法任务。
- 不跳过 `BLOCKED`、`PENDING_VERIFICATION`、`REWORK_REQUIRED` 或 `PENDING_USER_CONFIRMATION`。
- 不扩大当前时间切片范围，不顺手重构，不实现未确认功能。
- 当前时间切片未授权的文件不得修改。
- 新想法默认进入 `BACKLOG.md`，问题进入 `ISSUES.md`，风险进入 `RISK_REGISTER.md`。

## Completion Gate

任务只有在以下条件全部满足时才可标记为 `COMPLETED`：

- 产物已按任务要求实现。
- 修改范围符合当前时间切片。
- 关键路径能运行。
- 相关验证命令已执行并通过，或失败已明确记录且不影响本切片验收。
- 没有明显运行时错误或已知回归。
- 需求、架构、任务和时间切片约束均已对齐。
- `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`CHANGELOG.md`、`TEST_REPORT.md`、`ISSUES.md`、`AGENT_HANDOFF.md` 已按需更新。
- 下一合法任务或下一状态明确。

任一条件不满足，只能标记为 `PENDING_VERIFICATION`、`BLOCKED` 或 `REWORK_REQUIRED`。

## Repair Limit

同一问题最多修复 3 轮。每轮必须记录失败现象、原因假设、修改内容、验证命令和结果。第三轮仍失败时，创建 issue report，并将任务标记为 `BLOCKED` 或 `REWORK_REQUIRED`。

## User Correction Review

当用户提出需求、架构、任务、验收或流程修正时，先审查再执行：

1. 复述用户意图。
2. 对照已确认需求、架构、任务和当前时间切片。
3. 判断范围、成本、测试、上下文可恢复性和死亡优化风险。
4. 给出 `ACCEPT`、`ACCEPT_WITH_RISK`、`REJECT`、`DEFER_TO_BACKLOG` 或 `NEEDS_USER_DECISION`。
5. 需要确认时等待用户取舍。

## Next Step Protocol

当用户只说“下一步”：

1. 读取必读文件。
2. 确定当前阶段、任务和时间切片。
3. 如果当前任务阻塞，先处理阻塞。
4. 如果待验证，先验证。
5. 如果需要返工，进入失败修复流程。
6. 如果已完成，推进到下一个合法时间切片。
7. 否则只执行当前合法时间切片。
8. 运行验证并更新状态、测试报告、变更记录和交接说明。

## Safety

不得自动提交 git、推送远程、部署生产、执行不可逆迁移、删除重要文件、泄露密钥或把真实 `.env` 值写入聊天或文档。涉及生产、付费服务、密钥、认证权限或数据库 schema 的操作必须先获得明确确认。
