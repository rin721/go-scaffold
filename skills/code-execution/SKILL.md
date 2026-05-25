---
name: code-execution
description: Execute code changes only for the current legal time slice, respecting allowed and forbidden files, architecture constraints, tests, documentation updates, and handoff requirements. Use when STATUS points to an implementation task.
---

# Skill: code-execution

## Purpose

这个 skill 用来按当前唯一合法时间切片执行最小必要代码修改。

## When to Use

在以下情况必须使用：

- `STATUS.md` 指向实现类任务。
- 当前切片允许修改代码。
- 用户发送“下一步”且当前任务不是阻塞、待验证或返工。

## Inputs

必须读取：

- `AGENTS.md`
- `AGENT_RULES.md`
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- `ACCEPTANCE.md`
- `ISSUES.md`
- `TEST_REPORT.md`
- `AGENT_HANDOFF.md`

可选读取：

- 任务允许范围内的源码。
- `go.mod`
- `README.md`

## Outputs

必须写入或更新：

- 允许范围内的代码文件。
- `STATUS.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `CHANGELOG.md`
- `TEST_REPORT.md`
- `AGENT_HANDOFF.md`
- 必要时更新 `ISSUES.md` 或 `RISK_REGISTER.md`

## Preconditions

执行前必须满足：

- 当前任务和时间切片唯一。
- 允许文件和禁止文件范围明确。
- 不存在未处理的阻塞、待验证或返工状态。

## Procedure

1. 确认当前切片和允许修改范围。
2. 读取相关源码和测试。
3. 制定最小实现方案。
4. 只修改必要文件，不顺手重构。
5. 运行当前切片指定测试。
6. 测试失败时进入 `failure-repair`。
7. 测试通过后更新状态、变更、测试和交接文档。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 代码变更只落在允许范围。
- 验证命令已执行并记录。
- 文档和状态已更新。
- 下一合法步骤明确。

## Completion Decision

- `COMPLETED`：代码、验证、文档和状态均满足完成门禁。
- `PENDING_VERIFICATION`：代码已改但验证证据不足。
- `BLOCKED`：缺少依赖、权限、环境或用户决策。
- `REWORK_REQUIRED`：测试失败或验收不通过。

## Failure Handling

如果失败：

1. 记录失败命令和输出摘要。
2. 判断是否属于当前切片范围。
3. 属于范围则最多修复 3 轮；否则记录 issue 或 backlog。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 执行命令。
- 输出摘要。
- 测试结果。
- 状态更新。
- 下一步。
