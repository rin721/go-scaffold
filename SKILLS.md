# SKILLS.md

## 项目专用 Skills 索引

本仓库按 `docs/ai/prompt.md` 提供以下项目级 skills。每个 skill 的执行细则位于对应 `skills/*/SKILL.md`。

`skills/*/SKILL.md` 是 canonical source。`.agents/skills/*/SKILL.md` 仅作为通用 Agent 目录适配入口，必须指向 canonical skill，避免双份长内容漂移。

| Skill | 用途 |
|---|---|
| `requirements-clarification` | 从模糊想法提取目标、边界、风险和待确认项 |
| `requirements-generation` | 将确认内容转成正式需求、验收和 Backlog |
| `user-correction-review` | 审查用户修正是否与现有工程约束冲突 |
| `architecture-decomposition` | 生成架构方案、模块边界和技术取舍 |
| `task-decomposition` | 把模块拆成可执行任务 |
| `time-slicing` | 把任务拆成可验证、可恢复的小切片 |
| `code-execution` | 按唯一合法时间切片执行代码修改 |
| `test-verification` | 执行测试并更新验证结论 |
| `failure-repair` | 有限次数修复失败并记录问题 |
| `status-update` | 维护 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` |
| `context-recovery` | 在上下文丢失后恢复当前状态 |
| `agent-handoff` | 生成或更新交接说明 |
| `anti-death-optimization` | 防止无限优化和范围膨胀 |
| `acceptance-judgement` | 判断任务是否满足完成门禁 |

## 验证要求

- 每个 canonical skill 必须包含 YAML frontmatter：`name` 和 `description`。
- 每个 canonical skill 必须包含 Purpose、When to Use、Inputs、Outputs、Preconditions、Procedure、Acceptance Criteria、Completion Decision、Failure Handling、Evidence Required。
- 每个 `.agents` adapter 必须能通过 `quick_validate.py`，并指向对应 canonical skill。
