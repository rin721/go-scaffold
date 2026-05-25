---
name: requirements-clarification
description: Clarify ambiguous project ideas into goals, users, scenarios, boundaries, risks, and confirmation questions for go-scaffold. Use during project intake, when requirements are unclear, or when a user introduces a new direction before implementation.
---

# Skill: requirements-clarification

## Purpose

这个 skill 用来从开发者原始想法中提取项目目标、用户、场景、功能边界、风险和待确认项，并把事实写入仓库文档。

## When to Use

在以下情况必须使用：

- 项目首次启动或重新启动。
- 用户只给出模糊想法、目标或方向。
- 用户新增方向但尚未进入正式需求变更。
- 项目状态文件无法说明当前需求来源。

## Inputs

必须读取：

- `AGENTS.md`
- `AGENT_RULES.md`
- `PROJECT_BRIEF.md` 如果存在
- `STATUS.md`
- `docs/ai/prompt.md`

可选读取：

- `README.md`
- `REQUIREMENTS.md`
- `RISK_REGISTER.md`
- `docs/templates/project_start_template.md`
- `docs/templates/requirements_clarification_template.md`

## Outputs

必须写入或更新：

- `PROJECT_BRIEF.md`
- `docs/templates/project_start_template.md`
- `docs/templates/requirements_clarification_template.md`
- `STATUS.md`

## Preconditions

执行前必须满足：

- 已确认当前工作是澄清需求，不是写业务代码。
- 已识别现有项目文件，不能把仓库当空项目处理。

## Procedure

1. 读取必读文件，判断当前项目阶段。
2. 从用户输入和仓库事实中提取目标、用户、场景、功能、约束、风险。
3. 使用 `[CONFIRMED]`、`[INFERRED]`、`[NEEDS_CONFIRMATION]`、`[RISK]`、`[UNKNOWN]`、`[DEFERRED]` 标注事实来源。
4. 生成结构化待确认问题，区分必须确认、可默认但建议确认、可延后。
5. 写入项目启动和需求澄清模板。
6. 将整体状态设为 `PENDING_USER_CONFIRMATION`，除非文档已有明确确认。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 项目目标、用户、场景、非目标和风险均已记录。
- AI 推测没有被写成确认事实。
- 下一步是确认或修正需求，而不是代码实现。
- 状态文件能说明为什么等待确认。

## Completion Decision

- `COMPLETED`：澄清产物完整，状态已更新，下一步明确。
- `PENDING_VERIFICATION`：文档已写入但未核对完整性。
- `BLOCKED`：缺少关键用户意图，无法给出可选方案。
- `REWORK_REQUIRED`：事实标签错误、范围混乱或状态与文档冲突。

## Failure Handling

如果失败：

1. 生成最小可用假设。
2. 标出所有未知项和风险。
3. 给出 3 到 5 个候选方向。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 读取到的关键事实。
- 待确认问题。
- 状态更新。
- 下一步。
