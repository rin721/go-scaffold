# ACCEPTANCE.md

## 验收状态

- Project：go-scaffold
- Phase：Phase 6 收尾完成
- Status：COMPLETED
- Last Updated：2026-05-26

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
- Agent 基础设施补齐：COMPLETED
- Agent 基础设施一致性修复：COMPLETED
- 代码实现：COMPLETED，TASK-P1-015 已完成并通过验证

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
