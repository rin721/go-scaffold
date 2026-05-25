# CHANGELOG.md

## 最新变更

### 2026-05-25 - TASK-OPT-003 - TS-OPT-003

- 变更：新增 `MODULES.md`，记录模块职责、依赖方向、边界冲突、测试矩阵草案和 P1 优化候选项。
- 变更：确认 `.env.example` 与数据库环境变量前缀不一致、`copyConfig` 字段复制不完整、demo 自动迁移触发点分散、`cmd/server tests` 语义不一致等为优先收口风险。
- 变更：更新需求、验收、路线图、Backlog、任务、时间切片、状态、测试报告和交接文档。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-003 COMPLETED；TASK-OPT-004 NOT_STARTED。

## 历史变更

### 2026-05-25 - TASK-OPT-002 - TS-OPT-002

- 变更：按用户“下一步”确认推荐默认值。
- 变更：确认治理优先、`pkg/*` 混合策略、demo 长期标准示例、迁移 dev-prod 分层、中文化根文档和模板优先。
- 变更：新增 `ROADMAP.md`。
- 变更：更新需求、架构、决策、任务、时间切片、状态、验收、风险和交接文档。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-002 COMPLETED；TASK-OPT-003 NOT_STARTED。

### 2026-05-25 - TASK-OPT-001 - TS-OPT-001

- 变更：重新启动全项目治理与优化路线主线。
- 变更：将六个启动模板重写为中文：
  - `docs/templates/project_start_template.md`
  - `docs/templates/requirements_clarification_template.md`
  - `docs/templates/technical_options_template.md`
  - `docs/templates/architecture_constraints_template.md`
  - `docs/templates/acceptance_template.md`
  - `docs/templates/risk_confirmation_template.md`
- 变更：更新根目录项目文档和状态文件，使当前合法任务从插件系统扩展切换为项目优化启动确认。
- 变更：将插件系统 v1 内容保留为历史记录和 Backlog，不作为当前主线继续扩展。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-001 COMPLETED；TASK-OPT-002 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-HIST-PLUGIN-002 - TS-HIST-PLUGIN-002

- 历史：接受并关闭 `pkg/plugin` v1 local/http API 边界。
- 测试：
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：COMPLETED。

### 2026-05-25 - TASK-HIST-PLUGIN-001 - TS-HIST-PLUGIN-001

- 历史：实现独立 `pkg/plugin` 包，支持 local 和 HTTP 协议。
- 测试：
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：COMPLETED。
