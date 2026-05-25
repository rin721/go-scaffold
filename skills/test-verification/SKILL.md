---
name: test-verification
description: Run and record verification for the current task or time slice, including test command discovery, result classification, issue updates, and completion eligibility. Use after changes, for pending verification, or before marking work complete.
---

# Skill: test-verification

## Purpose

这个 skill 用来运行验证命令，判断当前任务是否满足测试和验收要求。

## When to Use

在以下情况必须使用：

- 代码或文档变更后。
- 当前状态为 `PENDING_VERIFICATION`。
- 准备将任务标记为 `COMPLETED` 前。
- 需要发现项目最小验证命令。

## Inputs

必须读取：

- `AGENTS.md`
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `ACCEPTANCE.md`
- `TEST_REPORT.md`
- `ISSUES.md`

可选读取：

- `package.json`
- `go.mod`
- `Makefile`
- `README.md`
- 最近变更文件。

## Outputs

必须写入或更新：

- `TEST_REPORT.md`
- `STATUS.md`
- `ISSUES.md`，如果失败。
- `RISK_REGISTER.md`，如果发现风险。

## Preconditions

执行前必须满足：

- 当前任务、切片或待验证项明确。
- 知道本次验证范围和期望结果。

## Procedure

1. 读取当前切片的验证命令。
2. 如果命令缺失，按项目规则从清单中发现最小足够命令。
3. 执行验证并记录命令、结果、耗时和摘要。
4. 判断失败是否属于当前切片。
5. 更新测试报告和状态。
6. 需要修复时进入 `failure-repair`。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 验证命令已执行，或无法执行的原因已记录。
- 结果已写入 `TEST_REPORT.md`。
- 失败项已写入 `ISSUES.md` 或风险登记。
- 是否可标记完成的结论明确。

## Completion Decision

- `COMPLETED`：验证通过且文档状态已更新。
- `PENDING_VERIFICATION`：验证尚未执行或证据不足。
- `BLOCKED`：环境、依赖或权限导致无法验证。
- `REWORK_REQUIRED`：验证失败且属于当前范围。

## Failure Handling

如果失败：

1. 记录失败现象和命令。
2. 判断是否当前切片范围。
3. 当前范围内进入有限修复；范围外登记 issue。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 命令。
- 输出摘要。
- 测试结果。
- 失败项。
- 状态更新。
- 下一步。
