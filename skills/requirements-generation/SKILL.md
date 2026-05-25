---
name: requirements-generation
description: Generate confirmed requirement documents, acceptance criteria, backlog items, and risk updates from clarified project intake material. Use after project intake is confirmed or corrected and before architecture or implementation work.
---

# Skill: requirements-generation

## Purpose

这个 skill 用来把已澄清的项目目标转成正式需求、验收标准、Backlog 和风险登记。

## When to Use

在以下情况必须使用：

- `PROJECT_BRIEF.md` 或启动模板已被用户确认或修正。
- 需要从需求草案进入架构设计。
- 新需求需要判断 P0、P1、P2 或明确不做范围。

## Inputs

必须读取：

- `AGENTS.md`
- `PROJECT_BRIEF.md`
- `STATUS.md`
- `docs/templates/requirements_clarification_template.md`

可选读取：

- `README.md`
- `ARCHITECTURE.md`
- `DECISIONS.md`
- `RISK_REGISTER.md`
- `BACKLOG.md`

## Outputs

必须写入或更新：

- `REQUIREMENTS.md`
- `ACCEPTANCE.md`
- `BACKLOG.md`
- `RISK_REGISTER.md`
- `STATUS.md`

## Preconditions

执行前必须满足：

- 项目启动材料已存在。
- 对用户修正已完成合理性审查。
- 不存在未解决的核心目标冲突。

## Procedure

1. 读取启动材料和用户确认结果。
2. 将需求拆为 P0、P1、P2 和明确不做。
3. 为每个 P0 需求生成验收标准和证据要求。
4. 记录非功能需求：安全、测试、维护、部署、可观测性。
5. 将未确认或非当前主线内容放入 `BACKLOG.md`。
6. 更新风险登记和状态文件。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- P0 需求都有验收标准。
- P1/P2 未混入当前必须实现范围。
- 明确不做范围存在。
- 风险和待确认事项已登记。

## Completion Decision

- `COMPLETED`：需求文档完整且已由用户确认。
- `PENDING_USER_CONFIRMATION`：需求已生成但仍需用户确认。
- `BLOCKED`：需求之间存在未决冲突。
- `REWORK_REQUIRED`：需求优先级、验收或风险登记不一致。

## Failure Handling

如果失败：

1. 标记 `[CONFLICT]` 或 `[UNKNOWN]`。
2. 给出取舍方案。
3. 等待用户确认。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 需求来源。
- P0 验收映射。
- Backlog 和风险更新。
- 下一步。
