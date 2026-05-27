# ACCEPTANCE.md

## 验收状态

- Project：go-scaffold
- Phase：P2 Linux Docker production 部署制品完成
- Status：COMPLETED
- Last Updated：2026-05-27

## 本轮启动验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-OPT-001 | 中文项目启动模板已生成 | 检查 `docs/templates/project_start_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-002 | 中文需求澄清模板已生成 | 检查 `docs/templates/requirements_clarification_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-003 | 中文技术方案模板已生成 | 检查 `docs/templates/technical_options_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-004 | 中文架构约束模板已生成 | 检查 `docs/templates/architecture_constraints_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-005 | 中文验收模板已生成 | 检查 `docs/templates/acceptance_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-006 | 中文风险确认模板已生成 | 检查 `docs/templates/risk_confirmation_template.md` | 是 | [CONFIRMED] |
| ACC-OPT-007 | 当前主线切换到项目优化启动确认 | 检查 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-OPT-008 | 本轮不修改 Go 代码 | 检查 git diff | 是 | [CONFIRMED] |
| ACC-OPT-009 | 全量 Go 测试通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## 文档验收

- [CONFIRMED] 当前输出优先使用中文。
- [CONFIRMED] 历史插件系统内容保留，但不再作为当前合法任务。
- [CONFIRMED] 当前风险和待确认项已进入模板和状态文件。
- [CONFIRMED] 后续“下一步”应先处理用户确认，而不是进入代码实现。

## 后续代码任务验收门禁

后续任何代码优化任务只有同时满足以下条件，才能标记为 `COMPLETED`：

1. [CONFIRMED] 已映射到确认后的需求。
2. [CONFIRMED] 已映射到确认后的架构边界。
3. [CONFIRMED] 已映射到唯一任务和唯一时间切片。
4. [CONFIRMED] 修改范围没有超出时间切片。
5. [CONFIRMED] 相关测试或验证命令已执行。
6. [CONFIRMED] 文档、状态、测试报告、变更记录和交接说明已更新。
7. [CONFIRMED] 下一步合法任务明确。

## 已确认验收

| ID | 验收项 | 结果 | 状态 |
|---|---|---|---|
| ACC-FUTURE-001 | 优化路线被确认 | 治理优先 | [CONFIRMED] |
| ACC-FUTURE-002 | `pkg/*` API 策略被确认 | 混合策略 | [CONFIRMED] |
| ACC-FUTURE-003 | demo 模块定位被确认 | 长期标准示例 | [CONFIRMED] |
| ACC-FUTURE-004 | 迁移策略被确认 | dev-prod 分层 | [CONFIRMED] |
| ACC-FUTURE-005 | 中文化范围被确认 | 根文档和模板优先，历史内容分阶段处理 | [CONFIRMED] |

## 当前完成判断

- 启动模板生成：COMPLETED
- 项目优化启动确认：COMPLETED
- 模块边界清单：COMPLETED
- 测试矩阵与任务拆分：COMPLETED
- P1 执行顺序确认：COMPLETED
- 配置 copy/update 测试与修复：COMPLETED
- 配置环境变量策略收拢：COMPLETED
- HTTP health/ready smoke test：COMPLETED
- demo CRUD 测试基线：COMPLETED
- demo 迁移边界收拢：COMPLETED
- CLI tests 命令语义收拢：COMPLETED
- pkg/* API 分类：COMPLETED
- pkg/sqlgen unsupported 边界标注：COMPLETED
- 下一阶段范围确认：COMPLETED，用户选择 A，提升 `BL-021` / `TM-P1-005`
- types/* 契约边界：COMPLETED
- `pkg/plugin` 被动注册边界：COMPLETED
- `pkg/plugin` 后续范围确认：COMPLETED，用户选择 A，提升 `BL-020` 首批行为测试
- 首批 `pkg/*` 行为测试：COMPLETED，`pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 已有最小行为测试
- 第二批 `pkg/*` 行为测试：COMPLETED，`pkg/executor`、`pkg/httpserver`、`pkg/storage` 已有最小行为测试
- 第三批 `pkg/*` 行为测试：COMPLETED，`pkg/cache` 已有基于进程内 Redis 的隔离行为测试
- `pkg/utils` 内部支撑测试：COMPLETED，`pkg/utils/utils_test.go` 已覆盖最小确定性行为
- app/router/middleware 集成测试：COMPLETED，`internal/transport/http/router_integration_test.go` 已覆盖 demo Todo HTTP 集成和 TraceID/CORS/Recovery 链路
- TASK-NEXT-SCOPE-008：COMPLETED，用户选择 A 进入 Phase 6 收尾
- Phase 6 收尾与交接：COMPLETED，最终回归、变更记录、验收、问题记录和交接说明已更新
- app 装配与 reload/config 集成测试：COMPLETED，`internal/app/app_integration_test.go` 和 `internal/app/reloadapp/reload_test.go` 已覆盖真实 app 装配、配置变更 hook 与 reload 分发
- 包 README 第一阶段中文化：COMPLETED，`pkg/*/README.md` 已统一主要中文读者文本并同步过期风险描述
- TASK-INFRA-003 状态一致性修复：COMPLETED，背景文档中的 TASK-P1-016 前旧待办表述已修复
- TASK-P2-001 CI 质量门禁与部署说明：COMPLETED，CI workflow 和手动部署说明已新增
- TASK-NEXT-SCOPE-010 真实 CD 范围确认：COMPLETED，用户已确认使用远程部署和 `.env` 风格配置
- TASK-P2-002 显式参数部署入口：COMPLETED，`deploy.sh` 和 `script/install.sh` 已新增，旧本地部署 env 文件依赖已删除
- TASK-P2-003 手动远程部署 workflow：COMPLETED，staging/manual/Secrets/SSH/Docker Compose 路径已新增，本会话未执行真实部署
- TASK-P2-004 Linux Docker production 部署制品：COMPLETED，Dockerfile、production Compose 示例、统一 `deploy.sh` 部署入口和手动 production 闸门已补齐；用户已在 Linux Docker 环境执行带 `GOPROXY` 的 Docker build 并通过
- TASK-P2-005 至 TASK-P2-010 插件钩子运行时与 IAM 公共接口：COMPLETED，`pkg/plugin/hooks`、hook-aware manager、HTTP 远程插件服务端、`RemoteHook`、`pkg/iam` memory、配置/app/reload/lifecycle 接入均已完成并通过验证
- Agent 基础设施补齐：COMPLETED
- Agent 基础设施一致性修复：COMPLETED
- 代码实现：COMPLETED，TASK-P1-016 已完成并通过验证

## Prompt 全量产物验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-INFRA-001 | 跨 Agent 入口文件已补齐 | 检查 `AGENTS.md`、`CLAUDE.md` | 是 | [CONFIRMED] |
| ACC-INFRA-002 | Agent 规则与 skills 索引已补齐 | 检查 `AGENT_RULES.md`、`SKILLS.md` | 是 | [CONFIRMED] |
| ACC-INFRA-003 | 任务拆分和时间切片模板已补齐 | 检查 `docs/templates/task_decomposition_template.md`、`docs/templates/time_slice_template.md` | 是 | [CONFIRMED] |
| ACC-INFRA-004 | reports/specs 目录入口已补齐 | 检查 `docs/reports/*`、`docs/specs/*` | 是 | [CONFIRMED] |
| ACC-INFRA-005 | 14 个项目专用 skills 已补齐 | 检查 `skills/*/SKILL.md` | 是 | [CONFIRMED] |
| ACC-INFRA-006 | 跨工具目录入口已补齐 | 检查 `.agents`、`.cursor`、`.kiro`、`.codex` | 是 | [CONFIRMED] |
| ACC-INFRA-007 | 文档基础设施补齐不破坏 Go 测试 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-INFRA-008 | 本任务未新增 Go 代码修改 | 核对变更范围 | 是 | [CONFIRMED] |
| ACC-INFRA-009 | `AGENTS.md` 实际存在且跨工具引用闭合 | 文件存在性核对和引用一致性检查 | 是 | [CONFIRMED] |
| ACC-INFRA-010 | canonical 与 `.agents` skills 均可通过 skill 验证 | `quick_validate.py` 验证 28 个 skill 目录 | 是 | [CONFIRMED] |
| ACC-INFRA-011 | 模板不混入当前项目实例事实 | 检查 `docs/templates/*` | 是 | [CONFIRMED] |
| ACC-INFRA-012 | 状态冲突已形成诊断报告 | 检查 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` | 是 | [CONFIRMED] |
| ACC-INFRA-013 | TASK-P1-016/017 后背景文档漂移已形成诊断报告并修复 | 检查 `docs/reports/status_diagnostics/2026-05-26-task-p1-017-post-completion-doc-drift.md` 和状态文档 | 是 | [CONFIRMED] |

## 测试矩阵与任务拆分验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-OPT-010 | 正式测试矩阵已生成 | 检查 `TEST_MATRIX.md` | 是 | [CONFIRMED] |
| ACC-OPT-011 | P1 任务草案已生成 | 检查 `TASKS.md` 中 TASK-P1-001 至 TASK-P1-008 | 是 | [CONFIRMED] |
| ACC-OPT-012 | P1 时间切片草案已生成 | 检查 `TIME_SLICES.md` 中 TS-P1-001 至 TS-P1-008 | 是 | [CONFIRMED] |
| ACC-OPT-013 | 每个 P1 任务有允许文件范围 | 检查 `TASKS.md` 和 `TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-OPT-014 | 每个 P1 任务有验证命令和退出条件 | 检查 `TASKS.md` 和 `TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-OPT-015 | 任务拆分阶段未修改 Go 代码 | 检查当时 git diff | 是 | [CONFIRMED] |

## 下一步确认验收

- [CONFIRMED] 推荐执行顺序已接受：`TEST_MATRIX.md` 中从 TASK-P1-001 到 TASK-P1-008。
- [CONFIRMED] 用户再次发送“下一步”已按推荐默认顺序视为确认，并已完成 TASK-P1-001。

## TASK-P1-001 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-001 | `copyConfig` 不丢失关键字段 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-002 | `copyConfig` 对 slice 做深拷贝 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-003 | `Update` 保留未修改字段 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-004 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-002 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-005 | 数据库环境变量主策略为 `DB_*` | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-006 | 旧 `REI_APP_DB_*` 仍作为兼容 fallback | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-007 | `DB_*` 优先级高于旧前缀 | `go test ./internal/config -count=1` | 是 | [CONFIRMED] |
| ACC-P1-008 | `.env.example` 与实现一致且不再暗示 JWT 已实现 | 人工检查 `.env.example` | 是 | [CONFIRMED] |
| ACC-P1-009 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-003 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-010 | `/health` HTTP 200 和响应语义被固定 | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-011 | `/ready` 数据库缺失路径返回 503 和 `not_ready` | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-012 | `/ready` 数据库 ping 失败路径返回 503 和错误语义 | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-013 | `/ready` 数据库 ping 成功路径返回 200 和 `ready` | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-014 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-004 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-015 | demo Todo Create/List/Get/Update/Delete 成功路径被固定 | `go test ./internal/modules/demo/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-016 | demo Todo 使用临时 SQLite 或等价隔离数据库 | 检查 `internal/modules/demo/service/todo_test.go` | 是 | [CONFIRMED] |
| ACC-P1-017 | demo Todo 不依赖真实外部服务 | `go test ./internal/modules/demo/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-018 | 空标题校验、not found 和软删除后不可见语义被固定 | `go test ./internal/modules/demo/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-019 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-005 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-020 | server-start 触发点会执行 demo `AutoMigrate` | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-021 | `initdb` 触发点会执行 demo `AutoMigrate` | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-022 | reload 触发点不会隐式执行 demo `AutoMigrate` | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-023 | dev/demo 与生产/bootstrap 迁移职责已记录 | 检查 `docs/specs/demo_migration_boundary.md` | 是 | [CONFIRMED] |
| ACC-P1-024 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-006 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-025 | `tests` 命令执行 Go tests 而非 yaml2go 示例 | `go test ./cmd/server -count=1` | 是 | [CONFIRMED] |
| ACC-P1-026 | `tests` 命令描述与行为一致 | 检查 `cmd/server/tests.go` 和 `cmd/server/tests_test.go` | 是 | [CONFIRMED] |
| ACC-P1-027 | `tests` 支持默认 `./...` 和指定 package pattern | `go test ./cmd/server -count=1` | 是 | [CONFIRMED] |
| ACC-P1-028 | CLI 语义边界已记录 | 检查 `docs/specs/cli_tests_command_boundary.md` | 是 | [CONFIRMED] |
| ACC-P1-029 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-007 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-030 | 13 个 `pkg/*` 包均有 API 分类 | 检查 `pkg/*/README.md` | 是 | [CONFIRMED] |
| ACC-P1-031 | 根架构文档包含 `pkg/*` API 分类表 | 检查 `ARCHITECTURE.md` | 是 | [CONFIRMED] |
| ACC-P1-032 | 模块清单与包 README 分类一致 | 检查 `MODULES.md` | 是 | [CONFIRMED] |
| ACC-P1-033 | `pkg/*` 破坏性重构仍需单独任务确认 | 检查 `ARCHITECTURE.md` 和 `MODULES.md` | 是 | [CONFIRMED] |
| ACC-P1-034 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-008 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-035 | `pkg/sqlgen` TODO/unsupported 能力边界被显式标注 | 检查 `pkg/sqlgen/*` 和包 README | 是 | [CONFIRMED] |
| ACC-P1-036 | 如涉及代码行为，unsupported 路径有测试覆盖 | `go test ./pkg/sqlgen -count=1` | 是 | [CONFIRMED] |
| ACC-P1-037 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-001 | 用户已明确选择后续范围 | 用户回复 `a`，对应选项 A | 是 | [CONFIRMED] |
| ACC-NEXT-002 | `BL-021` / `TM-P1-005` 已提升为正式任务 | 检查 `TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md` | 是 | [CONFIRMED] |
| ACC-NEXT-003 | 当前合法下一步不再是待确认状态 | 检查 `STATUS.md` 和 `AGENT_HANDOFF.md` | 是 | [CONFIRMED] |

## TASK-P1-009 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-038 | `types/result` HTTP/Gin 响应契约边界被标注 | 检查 `types/result/result.go` 和 `docs/specs/types_contract_boundary.md` | 是 | [CONFIRMED] |
| ACC-P1-039 | `types/errors` auth/rbac 预留错误码不暗示当前已实现 auth/rbac | 检查 `types/errors/doc.go` 和 `docs/specs/types_contract_boundary.md` | 是 | [CONFIRMED] |
| ACC-P1-040 | `types/constants` 和根 `types` 聚合入口的跨层边界被标注 | 检查 `types/constants/doc.go`、`types/doc.go` 和契约说明 | 是 | [CONFIRMED] |
| ACC-P1-041 | `types` 包测试通过 | `go test ./types/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-042 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-P1-010 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-043 | `pkg/plugin` 公共 manager API 不主动从配置加载并注册插件服务 | 检查 `pkg/plugin/manager.go` 和测试 | 是 | [CONFIRMED] |
| ACC-P1-044 | local 插件由插件服务或宿主装配层显式构造并 `Register` | `go test ./pkg/plugin -count=1` | 是 | [CONFIRMED] |
| ACC-P1-045 | HTTP 插件可由插件服务或宿主装配层显式构造并 `Register` | `go test ./pkg/plugin -count=1` | 是 | [CONFIRMED] |
| ACC-P1-046 | 被动注册边界已记录到 README、架构和决策文档 | 人工检查文档 | 是 | [CONFIRMED] |
| ACC-P1-047 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-003 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-004 | 用户已明确选择 `pkg/plugin` 后续范围 | 用户回复 `A`，对应提升 `BL-020` | 是 | [CONFIRMED] |
| ACC-NEXT-005 | `BL-020` 首批已提升为正式任务 | 检查 `TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md` | 是 | [CONFIRMED] |
| ACC-NEXT-006 | 当前合法下一步不再是待确认状态 | 检查 `STATUS.md` 和 `AGENT_HANDOFF.md` | 是 | [CONFIRMED] |

## TASK-P1-011 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-048 | `pkg/cli` 有最小包级行为测试 | `go test ./pkg/cli -count=1` | 是 | [CONFIRMED] |
| ACC-P1-049 | `pkg/i18n` 有最小包级行为测试 | `go test ./pkg/i18n -count=1` | 是 | [CONFIRMED] |
| ACC-P1-050 | `pkg/yaml2go` 有最小包级行为测试 | `go test ./pkg/yaml2go -count=1` | 是 | [CONFIRMED] |
| ACC-P1-051 | 首批 `pkg/*` 测试不依赖外部服务、生产配置或网络 | 人工检查测试实现 | 是 | [CONFIRMED] |
| ACC-P1-052 | 首批相关包测试通过 | `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1` | 是 | [CONFIRMED] |
| ACC-P1-053 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-004 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-007 | 用户已选择首批 `pkg/*` 行为测试后的后续范围 | 用户发送“下一步”，按选项 A 继续下一批 | 是 | [CONFIRMED] |
| ACC-NEXT-008 | 新的唯一合法任务或收尾状态已写入状态文件 | TASK-P1-012 / TS-P1-012 已写入 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |

## TASK-P1-012 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-054 | `pkg/executor` 有最小包级行为测试 | `go test ./pkg/executor -count=1` | 是 | [CONFIRMED] |
| ACC-P1-055 | `pkg/httpserver` 有最小包级行为测试 | `go test ./pkg/httpserver -count=1` | 是 | [CONFIRMED] |
| ACC-P1-056 | `pkg/storage` 有最小包级行为测试 | `go test ./pkg/storage -count=1` | 是 | [CONFIRMED] |
| ACC-P1-057 | 第二批 `pkg/*` 测试不依赖 Redis、数据库、第三方网络服务或生产配置 | 人工检查测试实现 | 是 | [CONFIRMED] |
| ACC-P1-058 | 第二批相关包测试通过 | `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1` | 是 | [CONFIRMED] |
| ACC-P1-059 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-005 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-009 | 用户已选择第二批 `pkg/*` 行为测试后的后续范围 | 用户回复 `A` | 是 | [CONFIRMED] |
| ACC-NEXT-010 | 新的唯一合法任务或收尾状态已写入状态文件 | TASK-P1-013 / TS-P1-013 已写入 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |

## TASK-P1-013 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-060 | `pkg/cache` 配置默认值和校验路径被固定 | `go test ./pkg/cache -count=1` | 是 | [CONFIRMED] |
| ACC-P1-061 | `pkg/cache` Redis 基本读写、批量、计数器和过期语义被隔离测试覆盖 | `go test ./pkg/cache -count=1` | 是 | [CONFIRMED] |
| ACC-P1-062 | `pkg/cache` reload 失败保持旧连接、成功切换新连接语义被覆盖 | `go test ./pkg/cache -count=1` | 是 | [CONFIRMED] |
| ACC-P1-063 | 第三批 `pkg/*` 测试不依赖真实 Redis、数据库、第三方网络服务或生产配置 | 人工检查测试实现 | 是 | [CONFIRMED] |
| ACC-P1-064 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-006 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-011 | 用户已选择 `pkg/cache` 行为测试后的后续范围 | 用户选择 B，提升 `pkg/utils` 内部支撑测试 | 是 | [CONFIRMED] |
| ACC-NEXT-012 | 新的唯一合法任务或收尾状态已写入状态文件 | TASK-P1-014 / TS-P1-014 已写入 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |

## TASK-P1-014 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-065 | `pkg/utils` Snowflake ID 生成路径被最小测试覆盖 | `go test ./pkg/utils -count=1` | 是 | [CONFIRMED] |
| ACC-P1-066 | `pkg/utils` 地址校验和端口查找路径被最小测试覆盖 | `go test ./pkg/utils -count=1` | 是 | [CONFIRMED] |
| ACC-P1-067 | `pkg/utils` 设备 ID 稳定性和 i18n helper 委托语义被最小测试覆盖 | `go test ./pkg/utils -count=1` | 是 | [CONFIRMED] |
| ACC-P1-068 | `pkg/utils` 测试不依赖真实外部网络服务、固定生产端口、数据库或生产配置 | 人工检查测试实现 | 是 | [CONFIRMED] |
| ACC-P1-069 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-007 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-013 | 用户已选择 `pkg/utils` 内部支撑测试后的后续范围 | 用户回复 `b`，对应选项 B | 是 | [CONFIRMED] |
| ACC-NEXT-014 | 新的唯一合法任务或收尾状态已写入状态文件 | TASK-P1-015 / TS-P1-015 已写入 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |

## TASK-P1-015 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-070 | demo Todo HTTP Create/List/Get/Update/Delete 集成路径被覆盖 | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-071 | TraceID、CORS 和 Recovery 中间件链路有路由级断言 | `go test ./internal/transport/http -count=1` | 是 | [CONFIRMED] |
| ACC-P1-072 | 集成测试不启动真实 HTTP server，不依赖外部数据库、Redis 或生产配置 | 人工检查 `internal/transport/http/router_integration_test.go` | 是 | [CONFIRMED] |
| ACC-P1-073 | 相关包测试通过 | `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-074 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-008 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-NEXT-015 | 用户已选择 app/router/middleware 集成测试后的后续范围 | 用户回复 `a`，选择 A 进入 Phase 6 收尾 | 是 | [CONFIRMED] |
| ACC-NEXT-016 | 新的唯一合法任务或收尾状态已写入状态文件 | 检查 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |

## TASK-PHASE6-001 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-PHASE6-001 | 用户选择 A 已记录为进入 Phase 6 收尾 | 检查 `DECISIONS.md`、`STATUS.md` | 是 | [CONFIRMED] |
| ACC-PHASE6-002 | 收尾文档已更新到本轮完成状态 | 检查 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`CHANGELOG.md`、`AGENT_HANDOFF.md` | 是 | [CONFIRMED] |
| ACC-PHASE6-003 | 最终全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-PHASE6-004 | diff 空白检查通过 | `git diff --check` | 是 | [CONFIRMED] |
| ACC-PHASE6-005 | 后续工作不会自动开始，必须重新确认 | 检查 `AGENT_HANDOFF.md` 和 `STATUS.md` | 是 | [CONFIRMED] |

## TASK-P1-016 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-075 | server 模式真实 app 装配链路被覆盖 | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-076 | initdb 模式仅初始化数据库并创建 demo schema，不装配 HTTP transport | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-077 | `ConfigManager.Update` 触发 app hook 并更新 `Core.Config`，不启动真实 server | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-078 | reload 分发覆盖未变化、单组件变化和 Redis/executor/storage 关闭置空路径 | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-079 | database reload 不触发 demo schema 隐式迁移路径 | `go test ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-080 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-081 | diff 空白检查通过 | `git diff --check` | 是 | [CONFIRMED] |

## TASK-P1-017 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P1-082 | 13 个 `pkg/*/README.md` 已检查并完成第一阶段中文化 | 人工检查 `pkg/*/README.md` | 是 | [CONFIRMED] |
| ACC-P1-083 | README 中与已完成测试状态明显冲突的风险描述已同步 | 人工检查包 README、`ARCHITECTURE.md`、`MODULES.md` | 是 | [CONFIRMED] |
| ACC-P1-084 | 未修改 Go 代码、依赖、配置 schema、HTTP 路由或数据库 schema | 核对 git diff 范围 | 是 | [CONFIRMED] |
| ACC-P1-085 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-P1-086 | diff 空白检查通过 | `git diff --check` | 是 | [CONFIRMED] |

## TASK-P2-001 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P2-001 | CI workflow 已新增且不使用 secrets | 检查 `.github/workflows/ci.yml` | 是 | [CONFIRMED] |
| ACC-P2-002 | CI 包含 gofmt 漂移报告、全量测试、server 构建和空白检查 | 检查 `.github/workflows/ci.yml` | 是 | [CONFIRMED] |
| ACC-P2-003 | 部署说明记录配置入口、手动运行和 initdb 边界 | 检查 `docs/deployment.md` | 是 | [CONFIRMED] |
| ACC-P2-004 | README 有 CI 与部署说明入口 | 检查 `README.md` | 是 | [CONFIRMED] |
| ACC-P2-005 | 未执行真实部署、未推送镜像、未写入密钥 | 核对变更范围 | 是 | [CONFIRMED] |
| ACC-P2-006 | 全量回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-007 | server 构建通过 | `go build -o <temp> ./cmd/server` | 是 | [CONFIRMED] |
| ACC-P2-008 | diff 空白检查通过 | `git diff --check` | 是 | [CONFIRMED] |

## TASK-NEXT-SCOPE-010 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P2-009 | 用户选择 C 已记录 | 检查 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-P2-010 | 真实 CD 缺少的目标平台、环境和 secrets 决策已列出 | 检查 `TASKS.md`、`TIME_SLICES.md` | 是 | [CONFIRMED] |
| ACC-P2-011 | 确认前不实现真实 CD workflow、不推送镜像、不连接远程环境 | 核对变更范围 | 是 | [CONFIRMED] |
| ACC-P2-012 | 用户确认使用远程部署 | 用户回复 | 是 | [CONFIRMED] |
| ACC-P2-013 | 用户确认远程部署 `.env` 风格配置 | 用户回复 | 是 | [CONFIRMED] |

## TASK-P2-002 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P2-014 | `deploy.sh` 和 `script/install.sh` 存在，示例只使用占位值 | 检查文件 | 是 | [CONFIRMED] |
| ACC-P2-015 | 旧本地部署 env 文件已删除且不再作为部署入口 | 检查 git 状态和部署文档 | 是 | [CONFIRMED] |
| ACC-P2-016 | 部署说明记录远程部署变量边界 | 检查 `docs/deployment.md` | 是 | [CONFIRMED] |
| ACC-P2-017 | 未实现真实部署 workflow、未连接服务器、未写入密钥 | 核对变更范围 | 是 | [CONFIRMED] |

## TASK-P2-003 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P2-018 | 远程部署 workflow 仅手动触发并要求确认词 | 检查 `.github/workflows/deploy-remote.yml` | 是 | [CONFIRMED] |
| ACC-P2-019 | workflow 当前只支持 staging，不自动生产发布 | 检查 workflow inputs 和校验步骤 | 是 | [CONFIRMED] |
| ACC-P2-020 | workflow 使用 GitHub Variables/Secrets 组装显式参数，不写真实值 | 检查 workflow 和部署脚本 | 是 | [CONFIRMED] |
| ACC-P2-021 | workflow 校验远程 SSH 输入并通过 SSH 执行 `script/install.sh` / `deploy.sh` 部署路径 | 检查 workflow | 是 | [CONFIRMED] |
| ACC-P2-022 | 部署说明包含 Secrets、远程主机前置条件和手动触发步骤 | 检查 `docs/deployment.md` | 是 | [CONFIRMED] |
| ACC-P2-023 | workflow YAML 可解析且 actionlint 通过 | 临时 Go YAML 解析；actionlint | 是 | [CONFIRMED] |
| ACC-P2-024 | diff 空白检查通过 | `git diff --check` | 是 | [CONFIRMED] |
| ACC-P2-025 | 本会话未执行真实部署、未连接服务器、未推送镜像、未写入密钥 | 核对变更范围和执行命令 | 是 | [CONFIRMED] |

## TASK-P2-004 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P2-026 | Dockerfile 存在且可构建 Linux server 镜像 | `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` | 是 | [CONFIRMED] 用户已在 Linux Docker 环境补跑通过，BuildKit 输出 `23/23 FINISHED`，镜像标记为 `docker.io/library/go-scaffold:local` |
| ACC-P2-027 | production Compose 示例存在并使用外置配置、数据和日志挂载 | 检查 `deploy/docker-compose.production.example.yml` | 是 | [CONFIRMED] |
| ACC-P2-028 | production 配置样例绑定 `0.0.0.0:9999` 且不含真实密钥 | 检查 `deploy/config.production.example.yaml` | 是 | [CONFIRMED] |
| ACC-P2-029 | 远程部署 workflow 支持 staging/production 手动选择并要求环境绑定确认词 | 检查 `.github/workflows/deploy-remote.yml` | 是 | [CONFIRMED] |
| ACC-P2-030 | 部署文档记录 GitHub Environment、production Secrets、目录权限和回滚边界 | 检查 `docs/deployment.md` | 是 | [CONFIRMED] |
| ACC-P2-031 | workflow YAML 与 actionlint 通过 | 临时 Go YAML 解析；`actionlint` | 是 | [CONFIRMED] |
| ACC-P2-032 | 全量 Go 回归通过 | `go test ./... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-033 | diff 空白检查通过 | `git diff --check` | 是 | [CONFIRMED] |
| ACC-P2-034 | 本会话未触发 workflow、未连接服务器、未推送镜像、未执行真实 production | 核对执行命令 | 是 | [CONFIRMED] |
| ACC-P2-035 | 远程 Linux 部署脚本按显式参数注入运行环境且不打印密钥值 | 检查 `deploy.sh`；`shfmt` Bash 语法解析 | 是 | [CONFIRMED] |

## TASK-P2-005 至 TASK-P2-010 验收

| ID | 验收项 | 方法 | 必须 | 状态 |
|---|---|---|---|---|
| ACC-P2-036 | `pkg/plugin/hooks` 提供独立钩子 API、优先级执行、复制快照、context 取消、停止语义和 nil handler 拒绝 | `go test ./pkg/plugin/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-037 | `plugin.Manager` 支持 `Hooks()`、`RegisterHook`、`WithHooks` 和标准钩子点，且保持被动注册模型 | `go test ./pkg/plugin/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-038 | `before_invoke` 可阻止插件调用，`invoke_error` 不覆盖原错误，`after_invoke` 错误返回插件响应和包装错误 | `go test ./pkg/plugin/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-039 | HTTP 远程插件服务端只接受 `POST /plugin/v1/invoke`，`RemoteHook` 可通过 `hooks.execute` 解码 `hooks.Result` | `go test ./pkg/plugin/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-040 | `pkg/iam` 公共 API 与 memory 实现支持 token、精确匹配、`*` 通配、拒绝优先、过期检查和默认拒绝 | `go test ./pkg/iam/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-041 | `plugin` / `iam` 配置默认 disabled，配置创建插件仅限 HTTP adapter，本地插件仍由代码显式注册 | `go test ./internal/config ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-042 | IAM 权限检查钩子只在 `internal/app` 注册，`pkg/plugin` 与 `pkg/iam` 互不导入 | 代码检查；包测试 | 是 | [CONFIRMED] |
| ACC-P2-043 | reload 先构建新 IAM/plugin 基础设施再替换，失败保留旧实例；关闭顺序在 HTTP server 后、cache/database 前关闭插件管理器 | `go test ./internal/config ./internal/app/... -count=1` | 是 | [CONFIRMED] |
| ACC-P2-044 | 全量回归、server build 和 diff 空白检查通过 | `go test ./... -count=1`；`go build -o <temp> ./cmd/server`；`git diff --check` | 是 | [CONFIRMED] |
| ACC-P2-045 | 本轮未实现 JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS、生产部署、镜像发布或密钥管理 | 核对变更范围 | 是 | [CONFIRMED] |
