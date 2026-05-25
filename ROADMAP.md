# ROADMAP.md

## 路线图状态

- 项目：go-scaffold
- 状态：ACTIVE
- 最后更新：2026-05-25
- 路线：治理优先

## Phase 0：项目重新启动

- 目标：按 `docs/ai/prompt.md` 建立中文启动材料，并把当前主线切回全项目优化。
- 状态：COMPLETED
- 证据：
  - `docs/templates/*` 六个模板已中文化。
  - `STATUS.md` 当前合法任务不再指向插件系统扩展。
  - `go test ./... -count=1` 通过。

## Phase 1：需求与高层架构确认

- 目标：确认默认优化路线和关键边界。
- 状态：COMPLETED
- 已确认：
  - 治理优先。
  - `pkg/*` 混合策略。
  - demo 作为长期标准示例。
  - 迁移采用 dev-prod 分层。
  - 中文化先覆盖根文档和模板。
  - auth/JWT 暂不进入实现。

## Phase 2：模块边界清单

- 目标：逐模块分析当前优缺点、边界风险和优化方向。
- 状态：COMPLETED
- 输出：
  - `MODULES.md` 已生成。
  - 模块职责清单已生成。
  - 设计边界冲突清单已生成。
  - P1 优化候选项已生成。
  - 进入代码优化前必须补齐的测试矩阵草案已生成。
- 当前候选模块：
  - `cmd/server`
  - `internal/app`
  - `internal/config`
  - `internal/transport/http`
  - `internal/modules/demo`
  - `pkg/*`
  - `types/*`

## Phase 3：测试矩阵与验收门禁

- 目标：为后续代码优化建立可验证基线。
- 状态：NOT_STARTED
- 输出：
  - app 启动 smoke test。
  - `/health`、`/ready` smoke test。
  - demo CRUD 集成测试。
  - 配置加载/环境覆盖/热更新测试。
  - 迁移策略验证。

## Phase 4：任务拆分与时间切片

- 目标：把确认后的优化项拆成可执行、可验证、可恢复的最小切片。
- 状态：NOT_STARTED
- 输出：
  - `TASKS.md` 的 P1/P2 任务清单。
  - `TIME_SLICES.md` 的唯一合法执行顺序。
  - 每个切片的允许文件范围、测试命令和退出条件。

## Phase 5：代码优化

- 目标：按唯一合法时间切片逐步优化项目。
- 状态：BLOCKED
- 进入条件：
  - Phase 2 模块边界清单完成。
  - Phase 3 测试矩阵确认。
  - Phase 4 时间切片确认。
- 禁止：
  - 不做未授权重构。
  - 不实现 Backlog 项。
  - 不绕过测试和状态更新。

## Phase 6：收尾与交接

- 目标：完成文档、测试、变更、风险和交接更新。
- 状态：NOT_STARTED
- 输出：
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `AGENT_HANDOFF.md`
  - 下一合法任务说明
