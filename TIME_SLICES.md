# TIME_SLICES.md

## 当前合法时间切片

- Time Slice ID：TS-P1-004
- Task ID：TASK-P1-004
- Status：NOT_STARTED
- Summary：HTTP health/ready smoke test 已完成。当前合法下一步为 TS-P1-004。

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

- Status：COMPLETED
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
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - 正式测试矩阵存在。
  - P1 任务和时间切片草案存在。
  - 下一合法任务明确。

### TS-OPT-005：确认正式测试矩阵和 P1 执行顺序

- Status：COMPLETED
- Task ID：TASK-OPT-005
- Scope：
  - 确认 `TEST_MATRIX.md`。
  - 确认 P1 任务顺序。
  - 确认是否进入首个代码切片。
- Recommended Default：
  - 用户再次发送“下一步”时，按推荐顺序进入 TS-P1-001。
- Allowed Files：
  - 项目文档和状态文件。
- Forbidden：
  - 不修改 Go 代码。
  - 不新增 Go 测试文件。
- Exit Conditions：
  - 用户确认，或用“下一步”接受推荐默认顺序。
  - 当前合法切片推进到 TS-P1-001。
- Completion Evidence：
  - 用户再次发送“下一步”，按推荐默认顺序确认。
  - 已进入并完成 TS-P1-001。

### TS-INFRA-001：补齐 Prompt 全量 Agent 基础设施

- Status：COMPLETED
- Task ID：TASK-INFRA-001
- Scope：
  - 新增跨 Agent 入口与规则文件。
  - 新增缺失模板。
  - 新增 reports/specs、跨工具目录和项目 skills。
  - 更新状态、验收、路线图、测试报告、变更记录和交接说明。
- Allowed Files：
  - `AGENTS.md`
  - `CLAUDE.md`
  - `AGENT_RULES.md`
  - `SKILLS.md`
  - `docs/templates/*`
  - `docs/reports/*`
  - `docs/specs/*`
  - `skills/*/SKILL.md`
  - `.agents/*`
  - `.cursor/*`
  - `.kiro/*`
  - `.codex/*`
  - 项目状态文档
- Forbidden：
  - 不修改 Go 代码。
  - 不执行 TASK-P1-002。
- Verification：
  - Prompt 全量产物存在性核对：PASS。
  - `go test ./... -count=1`：PASS。
- Exit Conditions：
  - Prompt 全量产物存在。
  - 当前合法下一步恢复到 TS-P1-002。

### TS-INFRA-002：Agent 基础设施一致性修复

- Status：COMPLETED
- Task ID：TASK-INFRA-002
- Scope：
  - 新增缺失的 `AGENTS.md`。
  - 统一跨工具入口引用。
  - 补齐 14 个 canonical skills 的 frontmatter 和完整结构。
  - 新增 14 个 `.agents` 轻量适配器。
  - 标准化 `docs/templates/*` 为可复用模板。
  - 生成状态诊断报告并更新状态、测试、变更和交接文件。
- Allowed Files：
  - `AGENTS.md`
  - `CLAUDE.md`
  - `AGENT_RULES.md`
  - `SKILLS.md`
  - `docs/templates/*`
  - `docs/reports/status_diagnostics/*`
  - `skills/*/SKILL.md`
  - `.agents/skills/*/SKILL.md`
  - `.cursor/*`
  - `.kiro/*`
  - `.codex/*`
  - 项目状态文档
- Forbidden：
  - 不修改 Go 代码。
  - 不执行 TASK-P1-002。
  - 不安装依赖、不提交 git、不部署。
- Verification：
  - Agent 基础设施文件存在性核对：PASS。
  - `quick_validate.py` 验证 canonical 和 adapter skills：PASS。
  - 跨工具入口引用一致性检查：PASS。
  - `go test ./... -count=1`：PASS。
- Exit Conditions：
  - `AGENTS.md` 存在。
  - 所有 required skills 可验证。
  - 状态冲突已记录并关闭。
  - 当前合法下一步恢复到 TS-P1-002。

## P1 时间切片草案

### TS-P1-001：配置 copy/update 测试与修复

- Status：COMPLETED
- Task ID：TASK-P1-001
- Matrix：TM-P0-001、TM-P0-002、TM-P0-006
- Allowed Files：
  - `internal/config/*`
  - 必要 testdata
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Exit Conditions：
  - 配置复制关键字段完整。
  - 测试和状态记录完成。
- Files Changed：
  - `internal/config/manager.go`
  - `internal/config/manager_test.go`
  - 项目状态文档

### TS-P1-002：配置环境变量策略收拢

- Status：COMPLETED
- Task ID：TASK-P1-002
- Matrix：TM-P0-001、TM-P0-006
- Allowed Files：
  - `internal/config/*`
  - `.env.example`
  - 配置相关文档
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - 环境变量策略一致或差异被文档化。
  - `.env.example` 与实现一致。
- Files Changed：
  - `.env.example`
  - `internal/config/app_database.go`
  - `internal/config/constants.go`
  - `internal/config/manager_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS

### TS-P1-003：HTTP health/ready smoke test

- Status：COMPLETED
- Task ID：TASK-P1-003
- Matrix：TM-P0-003、TM-P0-006
- Allowed Files：
  - `internal/transport/http/*_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/transport/http -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - health/ready 状态码和响应语义被测试固定。
- Files Changed：
  - `internal/transport/http/router_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/transport/http -count=1`：PASS
  - `go test ./... -count=1`：PASS
- Completion Evidence：
  - `/health` HTTP 200 和成功响应语义已固定。
  - `/ready` 数据库缺失、ping 失败、ping 成功路径已固定。
  - 未启动真实 HTTP server。

### TS-P1-004：demo CRUD 测试基线

- Status：NOT_STARTED
- Task ID：TASK-P1-004
- Matrix：TM-P0-005、TM-P0-006
- Allowed Files：
  - `internal/modules/demo/**/*_test.go`
  - 项目状态文档
- Verification：
  - `go test ./internal/modules/demo/... -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - demo Todo CRUD 关键路径有隔离测试。

### TS-P1-005：demo 迁移边界收拢

- Status：NOT_STARTED
- Task ID：TASK-P1-005
- Matrix：TM-P1-001、TM-P0-006
- Allowed Files：
  - `internal/app/**/*`
  - 迁移边界文档
  - 项目状态文档
- Verification：
  - `go test ./internal/app/... -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - demo 自动迁移触发策略清晰且可验证。

### TS-P1-006：CLI tests 命令语义收拢

- Status：NOT_STARTED
- Task ID：TASK-P1-006
- Matrix：TM-P1-002、TM-P0-006
- Allowed Files：
  - `cmd/server/*`
  - CLI 相关文档
  - 项目状态文档
- Verification：
  - `go test ./cmd/server -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - `tests` 命令不再误导使用者。

### TS-P1-007：`pkg/*` API 分类

- Status：NOT_STARTED
- Task ID：TASK-P1-007
- Matrix：TM-P1-003、TM-P0-006
- Allowed Files：
  - `ARCHITECTURE.md`
  - `MODULES.md`
  - 包 README
  - 项目状态文档
- Verification：
  - `go test ./... -count=1`
- Exit Conditions：
  - 每个 `pkg/*` 包定位明确。

### TS-P1-008：`pkg/sqlgen` unsupported 边界

- Status：NOT_STARTED
- Task ID：TASK-P1-008
- Matrix：TM-P1-004、TM-P0-006
- Allowed Files：
  - `pkg/sqlgen/*`
  - 包 README
  - 项目状态文档
- Verification：
  - `go test ./pkg/sqlgen -count=1`
  - `go test ./... -count=1`
- Exit Conditions：
  - 未实现能力边界清晰。

## 历史时间切片

### TS-HIST-PLUGIN-001：实现插件系统 v1

- Status：COMPLETED
- Summary：历史切片，完成 `pkg/plugin` local/http 支持。

### TS-HIST-PLUGIN-002：确认插件系统 v1 API

- Status：COMPLETED
- Summary：历史切片，确认 v1 local/http 边界；后续扩展留在 Backlog。
