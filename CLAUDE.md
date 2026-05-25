# CLAUDE.md

Claude Code 在本仓库中必须遵循 `AGENTS.md` 和 `AGENT_RULES.md`。长版工程级 Agent Prompt 位于 `docs/ai/prompt.md`。

## 工作入口

1. 先读 `AGENTS.md`。
2. 再读 `AGENT_RULES.md`。
3. 读取 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` 确认唯一合法任务。
4. 读取 `REQUIREMENTS.md`、`ARCHITECTURE.md`、`ACCEPTANCE.md`、`ISSUES.md`、`TEST_REPORT.md`、`AGENT_HANDOFF.md`。
5. 执行前确认允许文件范围。
6. 完成后更新状态、测试报告、变更记录和交接说明。

## 禁止事项

- 不把 Claude memory 或聊天记录当作唯一事实来源。
- 不绕过 `STATUS.md` 指向的任务。
- 不跳过阻塞、待验证或返工状态。
- 不在未验证、未更新文档或未记录证据时标记完成。
- 不自动提交、部署、执行不可逆迁移或泄露敏感信息。
