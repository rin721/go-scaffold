---
name: context-recovery
description: Recover go-scaffold project state after context loss, new session, tool switch, or conflicting status files by reading required documents, diagnosing blockers, and identifying the legal next step.
---

# Skill: context-recovery

## Purpose

这个 skill 用来在上下文丢失、换 Agent、换工具或状态冲突时恢复项目当前状态。

## When to Use

在以下情况必须使用：

- 新会话开始。
- 用户说“继续”或“下一步”但当前上下文不足。
- 必读文件缺失、冲突或过期。
- 新 Agent 接手项目。

## Inputs

必须读取：

- `AGENTS.md`
- `AGENT_RULES.md`
- `PROJECT_BRIEF.md`
- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- `ROADMAP.md`
- `MODULES.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `STATUS.md`
- `ISSUES.md`
- `TEST_REPORT.md`
- `DECISIONS.md`
- `CHANGELOG.md`
- `AGENT_HANDOFF.md`
- `BACKLOG.md`
- `RISK_REGISTER.md`

可选读取：

- 最近 git diff。
- 相关源码。

## Outputs

必须写入或更新：

- 恢复摘要。
- 必要时写入 `docs/reports/status_diagnostics/`。
- 必要时更新 `STATUS.md`。

## Preconditions

执行前必须满足：

- 不直接写业务代码。
- 先用文档恢复事实。

## Procedure

1. 读取所有必读文件。
2. 判断当前阶段、模块、任务和时间切片。
3. 检查阻塞、待验证、返工和未完成修复。
4. 判断下一步是否合法且唯一。
5. 输出恢复摘要。
6. 如无法恢复，生成状态诊断报告并标记阻塞。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 当前任务和时间切片明确，或诊断报告明确说明无法确定。
- 阻塞、待验证和返工状态已检查。
- 下一步和进入条件已记录。

## Completion Decision

- `COMPLETED`：恢复成功，可继续当前合法任务。
- `PENDING_VERIFICATION`：恢复后发现待验证项。
- `BLOCKED`：状态冲突或缺失无法安全恢复。
- `REWORK_REQUIRED`：恢复发现上次工作不符合验收。

## Failure Handling

如果失败：

1. 不继续写代码。
2. 生成状态诊断报告。
3. 给出最小修复方案。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 读取文件。
- 当前阶段、任务、切片。
- 冲突或缺失。
- 恢复结论。
- 下一步。
