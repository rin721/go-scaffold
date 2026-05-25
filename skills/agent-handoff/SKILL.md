---
name: agent-handoff
description: Create or update AGENT_HANDOFF with the latest project snapshot, changed files, commands, tests, blockers, decisions, risks, and legal next step so another agent can continue without chat context.
---

# Skill: agent-handoff

## Purpose

这个 skill 用来生成或更新 Agent 交接说明，让任何新 Agent 可以不依赖聊天记录接手。

## When to Use

在以下情况必须使用：

- 每个时间切片结束。
- 工具切换或会话结束前。
- 任务进入阻塞、返工或待验证。
- 用户要求交接或总结当前状态。

## Inputs

必须读取：

- `AGENTS.md`
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `TEST_REPORT.md`
- `CHANGELOG.md`
- `ISSUES.md`

可选读取：

- `RISK_REGISTER.md`
- `BACKLOG.md`
- `DECISIONS.md`
- 最近 git diff。

## Outputs

必须写入或更新：

- `AGENT_HANDOFF.md`
- 必要时写入 `docs/reports/handoff_snapshots/`

## Preconditions

执行前必须满足：

- 当前任务状态已判断。
- 已知道最近修改文件、命令和测试结果。

## Procedure

1. 读取当前状态、任务、切片和测试报告。
2. 记录项目快照、最近完成、变更文件、命令和测试状态。
3. 记录当前阻塞、待验证、重要决策、风险和 Backlog 提醒。
4. 写清合法下一步、进入条件和禁止事项。
5. 更新恢复说明。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 新 Agent 能从 `AGENT_HANDOFF.md` 判断当前合法任务。
- 最近变更、命令、测试和风险都有摘要。
- 禁止事项和恢复步骤明确。

## Completion Decision

- `COMPLETED`：交接说明足够恢复。
- `PENDING_VERIFICATION`：交接已写但测试状态未确认。
- `BLOCKED`：无法确定当前合法任务。
- `REWORK_REQUIRED`：交接内容与状态文件冲突。

## Failure Handling

如果失败：

1. 生成状态诊断报告。
2. 标明缺失事实。
3. 更新状态为阻塞或待确认。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 最近命令。
- 测试结果。
- 当前合法任务。
- 下一步。
