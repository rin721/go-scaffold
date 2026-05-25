---
name: status-update
description: Maintain STATUS, TASKS, and TIME_SLICES consistency for go-scaffold, including legal current work, blockers, pending verification, rework, completion evidence, and next step recovery.
---

# Skill: status-update

## Purpose

这个 skill 用来维护项目状态文件一致性，确保任何 Agent 都能无聊天上下文恢复当前合法任务。

## When to Use

在以下情况必须使用：

- 任务开始、完成、阻塞、返工或待验证。
- 当前任务或时间切片推进。
- 状态文件冲突、缺失或过期。
- 用户发送“下一步”前后。

## Inputs

必须读取：

- `AGENTS.md`
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `TEST_REPORT.md`
- `AGENT_HANDOFF.md`

可选读取：

- `CHANGELOG.md`
- `ISSUES.md`
- `RISK_REGISTER.md`

## Outputs

必须写入或更新：

- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- 必要时更新 `CHANGELOG.md`

## Preconditions

执行前必须满足：

- 已知道本次状态变化原因。
- 已有证据支持状态变更。

## Procedure

1. 读取当前状态和任务切片。
2. 检查当前任务是否唯一。
3. 根据证据更新任务、切片和项目阶段状态。
4. 记录阻塞、待验证、返工或完成证据。
5. 明确下一合法任务和进入条件。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` 对当前任务一致。
- 不存在文件声称产物存在但实际缺失的冲突。
- 下一步可由文档恢复。

## Completion Decision

- `COMPLETED`：状态一致且下一步明确。
- `PENDING_VERIFICATION`：状态更新后仍需验证。
- `BLOCKED`：状态冲突无法自动修复。
- `REWORK_REQUIRED`：状态错误导致执行偏离。

## Failure Handling

如果失败：

1. 生成状态诊断报告。
2. 标出冲突文件和影响。
3. 给出最小修复方案。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 状态变化原因。
- 支撑证据。
- 当前合法任务。
- 下一步。
