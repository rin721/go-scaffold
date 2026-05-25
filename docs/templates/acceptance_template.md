# 验收标准模板

## 1. 启动阶段验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-START-001 | 六个启动模板存在且为中文 | 检查 `docs/templates/*` | 是 | [CONFIRMED] |
| ACC-START-002 | 模板使用事实标签 | 人工审查 `[CONFIRMED]`、`[INFERRED]`、`[NEEDS_CONFIRMATION]`、`[RISK]`、`[DEFERRED]` | 是 | [CONFIRMED] |
| ACC-START-003 | 当前合法任务切换到项目优化启动确认 | 检查 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-START-004 | 本轮没有 Go 代码变更 | 检查 git diff | 是 | [CONFIRMED] |
| ACC-START-005 | 下一步是用户确认，不是代码实现 | 检查 `STATUS.md` | 是 | [CONFIRMED] |

## 2. 文档验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-DOC-001 | 历史英文状态不得继续作为当前唯一事实 | 当前主线文档改为中文项目优化启动 | 是 | [CONFIRMED] |
| ACC-DOC-002 | 插件系统 v1 内容保留为历史或 Backlog | 检查 `TASKS.md`、`TIME_SLICES.md`、`BACKLOG.md` | 是 | [CONFIRMED] |
| ACC-DOC-003 | 当前风险和待确认项明确 | 检查风险模板和 `STATUS.md` | 是 | [CONFIRMED] |
| ACC-DOC-004 | 中文化范围有待确认选项 | 检查需求澄清模板 | 是 | [CONFIRMED] |

## 3. 后续代码优化验收

后续任何代码任务只有同时满足以下条件，才能标记为 `COMPLETED`：

1. [CONFIRMED] 有已确认需求。
2. [CONFIRMED] 有已确认架构边界。
3. [CONFIRMED] 有明确任务和时间切片。
4. [CONFIRMED] 修改范围未超出当前时间切片。
5. [CONFIRMED] 相关验证命令已执行并记录。
6. [CONFIRMED] `STATUS.md`、`CHANGELOG.md`、`TEST_REPORT.md`、`AGENT_HANDOFF.md` 已更新。
7. [CONFIRMED] 下一步合法任务明确。

## 4. 当前验证命令

| 命令 | 目的 | 当前要求 |
|---|---|---|
| `go test ./... -count=1` | 确认文档变更未影响 Go 测试基线 | 本轮执行并记录 |

## 5. 未来测试矩阵草案

- [NEEDS_CONFIRMATION] app 启动 smoke test。
- [NEEDS_CONFIRMATION] `/health` 与 `/ready` HTTP smoke test。
- [NEEDS_CONFIRMATION] Demo Todo CRUD 集成测试。
- [NEEDS_CONFIRMATION] 配置加载与热更新测试。
- [NEEDS_CONFIRMATION] 数据库迁移策略测试。

## 6. 当前完成判断

- 启动模板生成：`COMPLETED`
- 项目优化路线确认：`PENDING_USER_CONFIRMATION`
- 代码实现：`BLOCKED`，直到用户确认需求和架构路线。
