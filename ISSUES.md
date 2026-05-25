# ISSUES.md

## Issue 状态

- 项目：go-scaffold
- 最后更新：2026-05-25
- 规则：失败、返工和阻塞问题记录在本文；风险项仍记录在 `RISK_REGISTER.md`。

## Open Issues

| ID | Linked Task | Severity | Status | Summary | Next Action |
|---|---|---|---|---|---|
|  |  |  |  |  |  |

## Issue Details

- 当前无已确认失败项。
- ISSUE-INFRA-002：`AGENTS.md` 缺失但状态文件声称已补齐。已在 TASK-INFRA-002 中修复，诊断报告见 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md`。
- ISSUE-P1-002：`.env.example` 与数据库环境变量前缀不一致，且 JWT 示例暗示未实现能力。已在 TASK-P1-002 中修复，相关测试通过。
- ISSUE-P1-003：无新增失败项。TASK-P1-003 新增 router smoke test 后，包测试和全量回归均通过。

## 历史说明

- 2026-05-25：记录并关闭 `.env.example` 与数据库环境变量实现不一致问题。
- 2026-05-25：记录 TASK-P1-003 无新增失败项，HTTP router smoke test 和全量回归通过。
- 2026-05-25：记录并关闭 `AGENTS.md` 缺失导致的 Agent 入口冲突。
- 2026-05-25：创建 `ISSUES.md`，补齐 `docs/ai/prompt.md` 要求的项目问题记录入口。
