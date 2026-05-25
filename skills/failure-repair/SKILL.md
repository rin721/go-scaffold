---
name: failure-repair
description: Repair failed tests, runtime errors, or acceptance failures with a maximum of three documented attempts, issue reporting, and status updates. Use when verification fails or current results do not satisfy acceptance criteria.
---

# Skill: failure-repair

## Purpose

这个 skill 用来对测试失败、运行失败或验收失败执行有限次数修复，并防止无限修复循环。

## When to Use

在以下情况必须使用：

- 验证命令失败。
- 当前结果不符合验收标准。
- 状态为 `REWORK_REQUIRED`。
- 同一问题已经反复出现。

## Inputs

必须读取：

- `AGENTS.md`
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `TEST_REPORT.md`
- `ISSUES.md`
- 最近失败命令和输出摘要。

可选读取：

- 最近修改文件。
- 相关源码和测试。
- `RISK_REGISTER.md`

## Outputs

必须写入或更新：

- 修复范围内的允许文件。
- `TEST_REPORT.md`
- `ISSUES.md`
- `STATUS.md`
- 必要时写入 `docs/reports/issue_reports/`

## Preconditions

执行前必须满足：

- 失败属于当前时间切片范围，或已明确记录为范围外问题。
- 当前问题的修复次数尚未超过 3 轮。

## Procedure

1. 记录失败现象、命令和输出摘要。
2. 判断失败是否属于当前切片。
3. 提出本轮原因假设。
4. 只修改允许文件范围内的最小内容。
5. 重新运行验证命令。
6. 失败则进入下一轮，最多 3 轮。
7. 第三轮仍失败，生成 issue report 并标记 `BLOCKED` 或 `REWORK_REQUIRED`。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 每轮修复都有假设、修改、命令和结果。
- 修复次数未无限扩张。
- 成功时测试通过并更新状态。
- 失败时 issue report 和阻塞状态明确。

## Completion Decision

- `COMPLETED`：修复成功，验证通过。
- `PENDING_VERIFICATION`：修复已做但验证未完成。
- `BLOCKED`：第三轮失败或缺少外部条件。
- `REWORK_REQUIRED`：仍不满足验收，需要重新设计。

## Failure Handling

如果失败：

1. 不继续第四轮盲修。
2. 生成问题报告。
3. 给出可选方案和所需用户决策。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 失败命令。
- 每轮假设和修改文件。
- 验证结果。
- issue 或状态更新。
- 下一步。
