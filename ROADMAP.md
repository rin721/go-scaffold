# ROADMAP.md

## 路线图状态

- 项目：go-scaffold
- 状态：COMPLETED
- 最后更新：2026-05-26
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

## Phase 1.5：Agent 基础设施

- 目标：补齐 `docs/ai/prompt.md` 要求的 Agent 入口、规则、skills、模板、reports/specs 和跨工具目录。
- 状态：COMPLETED
- 输出：
  - `AGENTS.md`、`CLAUDE.md`、`AGENT_RULES.md`、`SKILLS.md` 已补齐并一致引用。
  - `docs/templates/*` 已标准化为可复用模板。
  - reports/specs、跨工具目录、14 个 canonical skills 和 14 个 `.agents` adapters 已新增或修复。
  - `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` 已记录并关闭入口缺失冲突。
  - 当前合法下一步已恢复到 TASK-P1-002。

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
- 状态：COMPLETED
- 输出：
  - `TEST_MATRIX.md` 已生成。
  - app 启动、health/ready、demo CRUD、配置加载/环境覆盖/热更新、迁移策略均有矩阵项。
  - 每个矩阵项已包含建议文件范围、验证命令和退出条件。

## Phase 4：任务拆分与时间切片

- 目标：把确认后的优化项拆成可执行、可验证、可恢复的最小切片。
- 状态：COMPLETED
- 输出：
  - `TASKS.md` 已新增 P1 任务草案。
  - `TIME_SLICES.md` 已新增 P1 时间切片草案。
  - 推荐执行顺序已确认。

## Phase 5：代码优化

- 目标：按唯一合法时间切片逐步优化项目。
- 状态：COMPLETED
- 进入条件：
  - Phase 2 模块边界清单完成。
  - Phase 3 测试矩阵确认。
  - Phase 4 时间切片执行顺序确认。
- 禁止：
  - 不做未授权重构。
  - 不实现 Backlog 项。
  - 不绕过测试和状态更新。
- 当前进展：
  - TASK-P1-001 已完成：`copyConfig` 字段覆盖修复并补配置 copy/update 测试。
  - TASK-P1-002 已完成：配置环境变量策略收拢，`.env.example` 与实现对齐。
  - TASK-P1-003 已完成：`/health`、`/ready` router smoke test 已补齐并通过回归。
  - TASK-P1-004 已完成：demo Todo service/repository CRUD 测试基线已补齐并通过回归。
  - TASK-P1-005 已完成：demo 迁移边界已收拢，server-start/initdb/reload 触发策略可验证。
  - TASK-P1-006 已完成：`cmd/server tests` 已改为真实 Go test 入口，并补齐最小命令语义测试。
  - TASK-P1-007 已完成：13 个 `pkg/*` 包已完成公共基础设施 API、公共工具 API、内部支撑工具包分类。
  - TASK-P1-008 已完成：`pkg/sqlgen` TODO/unsupported 能力已显式返回错误或文档化 partial 边界。
  - TASK-NEXT-SCOPE 已完成：用户选择 A，提升 `BL-021` / `TM-P1-005`。
  - TASK-P1-009 已完成：`types/*` 契约边界已明确，并补充 `types/result` 与 `types/errors` 最小测试。
  - TASK-NEXT-SCOPE-002 已完成：用户修正并选择提升 `pkg/plugin` 被动注册边界。
  - TASK-P1-010 已完成：`pkg/plugin` 已收拢为被动 registry/runtime，local/http 插件由服务侧显式注册。
  - TASK-NEXT-SCOPE-003 已完成：用户选择 A，提升 `BL-020` 首批 `pkg/*` 行为测试。
  - TASK-P1-011 已完成：`pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 已补最小行为测试。
  - TASK-NEXT-SCOPE-004 已完成：用户确认继续第二批 `pkg/*` 行为测试。
  - TASK-P1-012 已完成：`pkg/executor`、`pkg/httpserver`、`pkg/storage` 已补最小行为测试。
  - TASK-NEXT-SCOPE-005 已完成：用户选择 A，提升第三批 `pkg/cache` 隔离行为测试。
  - TASK-P1-013 已完成：`pkg/cache` 已补进程内 Redis 隔离行为测试。
  - TASK-NEXT-SCOPE-006 已完成：用户选择 B，提升 `pkg/utils` 内部支撑测试。
  - TASK-P1-014 已完成：`pkg/utils` 已补最小确定性行为测试。
  - TASK-NEXT-SCOPE-007 已完成：用户选择 B，提升 app/router/middleware 等集成测试。
  - TASK-P1-015 已完成：`internal/transport/http/router_integration_test.go` 已覆盖 demo Todo HTTP 集成和 TraceID/CORS/Recovery 链路。
  - TASK-P1-016 已完成：`internal/app/app_integration_test.go` 和 `internal/app/reloadapp/reload_test.go` 已覆盖 app 装配、配置变更 hook 与 reload/config 分发。
  - TASK-P1-017 已完成：第一阶段 `pkg/*/README.md` 中文化已完成。
  - TASK-P2-001 已完成：CI 质量门禁与部署说明首切片已完成，不执行真实部署。
- 收尾决策：
  - [CONFIRMED] Phase 6 收尾已完成；后续用户又明确确认并完成 TASK-P1-016 与 TASK-P1-017。
  - [CONFIRMED] 当前无自动下一实现任务，后续更大范围中文化、生产迁移、auth/rbac、CI/CD 或部署仍需重新确认。

## Phase 6：收尾与交接

- 目标：完成文档、测试、变更、风险和交接更新。
- 状态：COMPLETED
- 输出：
  - `TEST_REPORT.md` 已记录最终回归。
  - `CHANGELOG.md` 已记录 TASK-PHASE6-001、TASK-P1-016、TASK-P1-017 和 TASK-INFRA-003。
  - `AGENT_HANDOFF.md` 已说明当前无自动下一实现任务。
  - 下一状态：追加测试、包 README 中文化和状态一致性修复完成；后续工作需要用户重新确认。

## Phase 7：CI/CD 与部署首切片

- 目标：建立非生产 CI 质量门禁，并记录手动部署边界。
- 状态：COMPLETED
- 输出：
  - `.github/workflows/ci.yml` 已新增。
  - `docs/deployment.md` 已新增。
  - `README.md` 已新增 CI 与部署入口。
  - `BL-007` 和 `BL-008` 已关闭。
  - 真实 CD、镜像发布、远程部署自动化进入 Backlog，需单独确认。

## Phase 8：真实 CD 范围确认

- 目标：确认真实 CD、镜像发布和远程部署自动化的目标平台、安全边界与实现输入；用户已确认使用远程部署。
- 状态：PENDING_VERIFICATION
- 当前切片：
  - TASK-NEXT-SCOPE-010 / TS-NEXT-SCOPE-010。
  - TASK-P2-002 / TS-P2-002。
  - TASK-P2-003 / TS-P2-003。
  - TASK-P2-004 / TS-P2-004。
- 待确认：
  - 镜像仓库、发布环境、触发方式和审批策略。
  - GitHub Secrets 名称和权限边界。
  - Docker build 在具备 Docker 的环境中补跑验证。
  - 镜像发布流水线、真实 production 运行和生产迁移框架是否开始实现。
- 已完成：
  - `.env.deploy.example` 远程部署变量模板。
  - `.env.deploy` Git 忽略规则。
  - 部署说明中的远程变量边界。
  - 手动 staging 远程部署 workflow。
  - Dockerfile、production Compose 示例、production 配置样例、远程 Linux 动态 env 部署脚本和手动 production workflow 闸门。
- 约束：
  - 本会话未执行远程部署、不推送镜像、不连接远程环境、不读取真实 secrets；远程 Linux 脚本仅按参数生成非密钥 `.env.deploy`。
  - 当前本机缺少 Docker CLI，`docker build -t go-scaffold:local .` 待补跑。
  - 镜像发布、真实 production 运行、生产迁移框架仍需单独确认。
