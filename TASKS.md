# TASKS.md

## 当前合法任务

- Task ID：TASK-OPT-004
- Status：NOT_STARTED
- Time Slice：TS-OPT-004
- Summary：生成正式测试矩阵和任务拆分草案。

## 任务列表

### TASK-OPT-001：生成中文项目优化启动材料

- Status：COMPLETED
- Verification：
  - `go test ./... -count=1`
- Completion Evidence：
  - 六个中文模板已生成。
  - 当前主线已切换为项目优化启动确认。
  - `go test ./... -count=1` 通过。

### TASK-OPT-002：确认项目优化路线和关键边界

- Status：COMPLETED
- Confirmed Decisions：
  - 优化路线：治理优先。
  - `pkg/*` 定位：混合策略。
  - demo 模块定位：长期标准示例。
  - 迁移策略：dev-prod 分层策略。
  - 中文化范围：根文档和模板优先，历史文档与包 README 分阶段处理。
  - auth/JWT：先延后处理，不进入当前实现范围。
- Completion Evidence：
  - 用户发送“下一步”，按当前文档推荐默认值视为确认。
  - 决策已写入 `DECISIONS.md`。

### TASK-OPT-003：生成模块边界清单和优化路线明细

- Status：COMPLETED
- Goal：逐模块分析当前优缺点、职责边界、设计冲突、测试缺口和优化优先级。
- Requirements Covered：
  - REQ-OPT-P1-001
  - REQ-OPT-P1-002
  - REQ-OPT-P1-003
- Allowed Files：
  - 项目文档和状态文件。
  - `MODULES.md`。
- Non-Goals：
  - 不写 Go 代码。
  - 不重构模块。
  - 不实现 Backlog。
  - 不扩展插件系统。
- Verification：
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - `MODULES.md` 已生成。
  - 模块职责清单、设计边界冲突清单、测试矩阵草案、P1 优化候选项已生成。
  - 状态、测试报告、变更日志和交接文档已更新到 TASK-OPT-004。

### TASK-OPT-004：生成正式测试矩阵和任务拆分草案

- Status：NOT_STARTED
- Goal：把 `MODULES.md` 的测试矩阵草案和 P1 优化候选项转成正式任务与时间切片草案。
- Requirements Covered：
  - REQ-OPT-P1-004
- Allowed Files：
  - 项目文档和状态文件。
  - 可新增 `TEST_MATRIX.md`。
  - 可更新 `TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`。
  - 不允许修改 Go 实现文件。
- Expected Output：
  - 正式测试矩阵。
  - P1 任务拆分。
  - P1 时间切片草案。
  - 每个任务的允许文件范围和验证命令。
- Non-Goals：
  - 不写 Go 测试代码。
  - 不修复已发现问题。
  - 不实现 Backlog。
- Verification：
  - `go test ./... -count=1`

## 历史任务

### TASK-HIST-PLUGIN-001：实现独立插件系统 v1

- Status：COMPLETED
- Summary：历史任务，已完成 `pkg/plugin` local/http 能力。
- Verification：
  - `go test ./pkg/plugin -count=1`
  - `go test ./... -count=1`

### TASK-HIST-PLUGIN-002：确认插件系统 v1 API 边界

- Status：COMPLETED
- Summary：历史任务，v1 local/http API 已接受。
- Follow-up：
  - rpc/ws/discovery/examples 保留在 Backlog，不属于当前主线。
