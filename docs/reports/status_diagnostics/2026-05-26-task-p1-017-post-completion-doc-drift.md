# TASK-P1-017 后文档状态漂移诊断

## 摘要

- 日期：2026-05-26
- 触发：用户发送“下一步”，按 `AGENTS.md` 要求执行上下文恢复与状态一致性检查。
- 结论：核心状态文件已一致指向 `NONE / COMPLETED`，但部分背景文档仍保留 TASK-P1-016 前的旧表述，可能误导后续 Agent 认为 app 装配、reload/config 仍待补测试。

## 冲突事实

- `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`TEST_REPORT.md` 和 `AGENT_HANDOFF.md` 已记录 TASK-P1-016 完成，并确认 `internal/app` 与 `internal/app/reloadapp` 已覆盖 app 装配、配置变更 hook 与 reload/config 分发路径。
- `ARCHITECTURE.md`、`MODULES.md`、`PROJECT_BRIEF.md` 和 `ROADMAP.md` 仍存在“app 装配/reload/config 后续确认或待覆盖”的旧语句。
- 该漂移不影响当前唯一合法任务判断；当前仍为 `NONE / COMPLETED`，无自动下一实现任务。

## 最小修复

- 将背景文档中的旧待办表述改为 TASK-P1-016 已完成事实。
- 将 `pkg/i18n` 已补测试事实同步到架构风险表述。
- 更新状态、任务、时间切片、验收、测试报告、变更记录、问题记录和交接说明，记录本次 `TASK-INFRA-003 / TS-INFRA-003` 状态一致性修复。

## 验证

- `go test ./... -count=1`
- `git diff --check`

## 下一合法状态

- Task ID：NONE
- Time Slice ID：NONE
- Status：COMPLETED
- 后续工作仍需用户重新确认并建立新的任务/时间切片。
