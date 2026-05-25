---
name: anti-death-optimization
description: Prevent endless optimization, scope creep, unauthorized refactors, repeated repairs, and deviations from the current legal time slice. Use when new ideas, refactor urges, broad fixes, or non-critical improvements appear.
---

# Skill: anti-death-optimization

## Purpose

这个 skill 用来防止无限优化、范围膨胀、无授权重构和偏离当前时间切片。

## When to Use

在以下情况必须使用：

- 出现新想法但不属于当前切片。
- 想顺手重构或美化。
- 当前功能已满足验收但仍想继续优化。
- 同一问题反复修复且接近 3 轮。
- 用户提出明显超出当前阶段的扩展。

## Inputs

必须读取：

- `AGENTS.md`
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `ACCEPTANCE.md`

可选读取：

- `BACKLOG.md`
- `RISK_REGISTER.md`
- `ISSUES.md`

## Outputs

必须写入或更新：

- `BACKLOG.md`，用于非关键优化。
- `RISK_REGISTER.md`，用于范围或成本风险。
- `ISSUES.md`，用于影响当前验收的问题。
- 必要时更新 `STATUS.md`。

## Preconditions

执行前必须满足：

- 已判断建议是否属于当前切片验收所必需。

## Procedure

1. 对照当前切片目标和非目标。
2. 判断建议是否影响运行、测试、安全、需求一致性或核心验收。
3. 非必要内容进入 Backlog。
4. 必要但超范围内容登记风险或创建后续任务。
5. 若用户坚持高风险方向，进入 `user-correction-review`。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 当前切片范围未扩大。
- 非关键优化已记录而非插队实现。
- 继续执行或等待确认的理由明确。

## Completion Decision

- `COMPLETED`：范围已收敛，当前任务可继续。
- `PENDING_USER_CONFIRMATION`：用户需要确认范围变更。
- `BLOCKED`：用户要求与核心约束冲突。
- `REWORK_REQUIRED`：已发生范围越界，需要回退到任务边界。

## Failure Handling

如果失败：

1. 停止继续优化。
2. 记录偏离点。
3. 恢复到当前合法切片。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 触发的新想法或优化。
- 范围判断。
- Backlog、风险或 issue 更新。
- 当前任务是否继续。
- 下一步。
