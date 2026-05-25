# TIME_SLICES.md

## 当前合法时间切片

- Time Slice ID：TS-OPT-004
- Task ID：TASK-OPT-004
- Status：NOT_STARTED
- Summary：生成正式测试矩阵和任务拆分草案。

## 时间切片列表

### TS-OPT-001：生成中文项目优化启动材料

- Status：COMPLETED
- Task ID：TASK-OPT-001
- Verification：
  - `go test ./... -count=1`：PASS

### TS-OPT-002：确认项目优化路线和边界

- Status：COMPLETED
- Task ID：TASK-OPT-002
- Confirmed：
  - 治理优先。
  - `pkg/*` 混合策略。
  - demo 长期标准示例。
  - 迁移 dev-prod 分层。
  - 中文化根文档和模板优先。
  - auth/JWT 延后处理。

### TS-OPT-003：生成模块边界清单

- Status：COMPLETED
- Task ID：TASK-OPT-003
- Files Changed：
  - `MODULES.md`
  - `REQUIREMENTS.md`
  - `ACCEPTANCE.md`
  - `ROADMAP.md`
  - `BACKLOG.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `STATUS.md`
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `AGENT_HANDOFF.md`
- Verification：
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - 模块边界清单存在。
  - 测试矩阵草案存在。
  - P1 优化候选项进入 Backlog 或任务草案。

### TS-OPT-004：生成正式测试矩阵和任务拆分草案

- Status：NOT_STARTED
- Task ID：TASK-OPT-004
- Scope：
  - 将 `MODULES.md` 的测试矩阵草案整理为正式 `TEST_MATRIX.md`。
  - 将 P1 优化候选项拆成任务草案。
  - 为每个任务定义允许文件范围、验证命令和退出条件。
  - 更新状态、验收和交接文件。
- Allowed Files：
  - 项目文档和状态文件。
  - 可新增 `TEST_MATRIX.md`。
- Forbidden：
  - 不修改 Go 代码。
  - 不新增 Go 测试文件。
  - 不实现 Backlog 项。
- Verification：
  - `go test ./... -count=1`
- Exit Conditions：
  - 正式测试矩阵存在。
  - P1 任务和时间切片草案存在。
  - 下一合法任务明确。

## 历史时间切片

### TS-HIST-PLUGIN-001：实现插件系统 v1

- Status：COMPLETED
- Summary：历史切片，完成 `pkg/plugin` local/http 支持。

### TS-HIST-PLUGIN-002：确认插件系统 v1 API

- Status：COMPLETED
- Summary：历史切片，确认 v1 local/http 边界；后续扩展留在 Backlog。
