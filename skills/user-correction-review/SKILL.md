---
name: user-correction-review
description: Review user-proposed corrections to requirements, architecture, tasks, time slices, acceptance criteria, or process rules before applying them. Use when a user changes scope, priority, technical direction, or asks to bypass the current legal task.
---

# Skill: user-correction-review

## Purpose

这个 skill 用来审查用户修正是否符合已确认需求、架构、任务、验收和 Agent 稳定推进规则。

## When to Use

在以下情况必须使用：

- 用户提出改变需求、架构、任务顺序、验收标准或流程。
- 用户要求跳过当前合法任务。
- 用户提出可能扩大范围、降低测试或破坏架构的方案。

## Inputs

必须读取：

- `AGENTS.md`
- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `STATUS.md`
- `RISK_REGISTER.md`

可选读取：

- `ACCEPTANCE.md`
- `DECISIONS.md`
- `BACKLOG.md`
- `ISSUES.md`

## Outputs

必须写入或更新：

- 修正审查结论。
- `DECISIONS.md`，如果产生决策。
- `RISK_REGISTER.md`，如果新增风险。
- `BACKLOG.md`，如果延后处理。
- `STATUS.md`，如果影响当前状态。

## Preconditions

执行前必须满足：

- 已复述用户意图。
- 已确定修正类型：需求、架构、范围、优先级、验收、任务拆分、技术偏好或临时优化。

## Procedure

1. 对照已确认需求和架构检查冲突。
2. 判断是否影响当前任务、时间切片、测试和交接。
3. 判断是否存在死亡优化、范围膨胀或上下文不可恢复风险。
4. 给出 `ACCEPT`、`ACCEPT_WITH_RISK`、`REJECT`、`DEFER_TO_BACKLOG` 或 `NEEDS_USER_DECISION`。
5. 如接受，更新相关文档；如拒绝或延后，记录原因和替代方案。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 审查结论明确。
- 原因、后果、替代方案和推荐方案已说明。
- 需要用户确认时状态为 `PENDING_USER_CONFIRMATION`。
- 接受的修正已写入相关文档。

## Completion Decision

- `COMPLETED`：审查和必要文档更新完成。
- `PENDING_USER_CONFIRMATION`：需要用户做取舍。
- `BLOCKED`：修正会破坏核心约束且无安全替代方案。
- `REWORK_REQUIRED`：已错误采纳修正，需要回到审查流程。

## Failure Handling

如果失败：

1. 停止执行修正。
2. 标记冲突和影响。
3. 给出最小安全替代方案。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 用户意图。
- 冲突检查。
- 风险和替代方案。
- 决策或状态更新。
- 下一步。
