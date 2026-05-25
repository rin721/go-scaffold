---
name: task-decomposition
description: Break confirmed requirements and architecture into engineering tasks with allowed files, forbidden files, dependencies, verification commands, acceptance criteria, and completion rules. Use before implementation planning or when tasks are ambiguous.
---

# Skill: task-decomposition

## Purpose

这个 skill 用来把需求和模块拆成工程级任务，并确保每个任务可执行、可验证、可恢复。

## When to Use

在以下情况必须使用：

- 架构或模块边界已确认，需要进入实现准备。
- 当前任务过大、不可验收或缺少允许文件范围。
- 需要恢复或重建 `TASKS.md`。

## Inputs

必须读取：

- `AGENTS.md`
- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- `MODULES.md`
- `ROADMAP.md`
- `ACCEPTANCE.md`

可选读取：

- `TEST_MATRIX.md`
- `RISK_REGISTER.md`
- `BACKLOG.md`
- `STATUS.md`

## Outputs

必须写入或更新：

- `TASKS.md`
- `ROADMAP.md`
- `STATUS.md`
- 必要时更新 `RISK_REGISTER.md`

## Preconditions

执行前必须满足：

- 架构已确认或已明确进入任务拆分草案阶段。
- 不存在阻塞任务拆分的核心需求冲突。

## Procedure

1. 按模块、阶段、子模块拆分任务。
2. 为每个任务定义 ID、优先级、复杂度、依赖、输入、输出。
3. 明确允许文件、禁止文件、步骤、验证命令和验收标准。
4. 标记依赖关系和阻塞风险。
5. 指定当前唯一合法任务。
6. 将未确认或非当前主线内容放入 Backlog 或风险登记。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 每个 P0 需求至少映射到一个任务。
- 每个任务有验收标准和验证方式。
- 每个任务有允许和禁止修改范围。
- 当前唯一合法任务明确。

## Completion Decision

- `COMPLETED`：任务拆分已确认。
- `PENDING_USER_CONFIRMATION`：任务顺序或范围待确认。
- `BLOCKED`：需求或架构冲突导致无法拆分。
- `REWORK_REQUIRED`：任务过大、不可验证或范围不清。

## Failure Handling

如果失败：

1. 生成状态诊断。
2. 标出不可拆分原因。
3. 给出更小任务方案。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 需求到任务映射。
- 当前合法任务。
- 阻塞项。
- 下一步。
