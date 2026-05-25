---
name: acceptance-judgement
description: Judge whether a task or time slice is truly complete using acceptance criteria, tests, changed files, status updates, changelog, issues, test report, and handoff evidence. Use before marking work COMPLETED.
---

# Skill: acceptance-judgement

## Purpose

这个 skill 用来判断任务或时间切片是否真的满足完成门禁，防止把未验证或未记录的工作标记为完成。

## When to Use

在以下情况必须使用：

- 准备将任务或切片标记为 `COMPLETED`。
- 验证通过后需要最终判定。
- 状态从 `PENDING_VERIFICATION` 转为完成前。
- 用户询问任务是否已经完成。

## Inputs

必须读取：

- `AGENTS.md`
- `ACCEPTANCE.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `TEST_REPORT.md`
- `STATUS.md`
- `CHANGELOG.md`
- `ISSUES.md`
- `AGENT_HANDOFF.md`

可选读取：

- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- 最近 git diff。

## Outputs

必须写入或更新：

- 完成判定。
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- 必要时更新 `TEST_REPORT.md`、`ISSUES.md` 或 `AGENT_HANDOFF.md`

## Preconditions

执行前必须满足：

- 已有产物或变更可供判定。
- 已执行或明确无法执行验证。

## Procedure

1. 检查产物是否存在。
2. 检查代码或文档是否在允许范围内。
3. 检查测试是否运行并通过。
4. 检查需求、架构、任务和切片是否一致。
5. 检查状态、变更、问题、测试报告和交接是否更新。
6. 判断是否可标记 `COMPLETED`，否则选择更准确状态。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 判定结论明确。
- 证据清单完整。
- 未满足条件不会被忽略。
- 下一步状态明确。

## Completion Decision

- `COMPLETED`：全部完成门禁满足。
- `PENDING_VERIFICATION`：缺少验证证据。
- `BLOCKED`：缺少外部条件或决策。
- `REWORK_REQUIRED`：产物、测试或文档不符合验收。

## Failure Handling

如果失败：

1. 列出未满足条件。
2. 更新状态为更准确的非完成状态。
3. 指定修复或验证下一步。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 执行命令。
- 输出摘要。
- 测试结论。
- 状态更新。
- 下一步。
