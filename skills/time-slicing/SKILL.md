---
name: time-slicing
description: Split engineering tasks into small legal time slices with scope limits, strict non-goals, verification commands, evidence requirements, and next-slice entry conditions. Use after task decomposition and before execution.
---

# Skill: time-slicing

## Purpose

这个 skill 用来把任务拆成可执行、可验证、可恢复的小时间切片。

## When to Use

在以下情况必须使用：

- `TASKS.md` 已有任务，但缺少可执行切片。
- 当前任务过大，需要限制一次工作范围。
- 需要确定“下一步”的唯一合法切片。

## Inputs

必须读取：

- `AGENTS.md`
- `TASKS.md`
- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- `ACCEPTANCE.md`

可选读取：

- `TEST_MATRIX.md`
- `STATUS.md`
- `RISK_REGISTER.md`

## Outputs

必须写入或更新：

- `TIME_SLICES.md`
- `TASKS.md`
- `STATUS.md`

## Preconditions

执行前必须满足：

- 任务已存在且目标清晰。
- 每个切片能绑定明确验证方式。

## Procedure

1. 读取任务和验收标准。
2. 将任务拆成小而完整的切片。
3. 为每个切片定义输入、输出、允许文件、禁止文件和严格非目标。
4. 指定执行步骤、测试命令、验收标准、失败处理和最大修复次数。
5. 标记当前唯一合法时间切片。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 每个切片只做一种工作类型。
- 每个切片有验证命令或人工验证方法。
- 当前唯一合法切片明确。
- 下一切片进入条件明确。

## Completion Decision

- `COMPLETED`：切片拆分已确认。
- `PENDING_USER_CONFIRMATION`：切片顺序或范围待确认。
- `BLOCKED`：任务缺少足够信息，无法切片。
- `REWORK_REQUIRED`：切片过大、混合多类工作或不可验证。

## Failure Handling

如果失败：

1. 标出过大或不可验证的切片。
2. 重新拆分为更小切片。
3. 更新状态。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 切片列表。
- 当前合法切片。
- 验证命令。
- 下一步。
