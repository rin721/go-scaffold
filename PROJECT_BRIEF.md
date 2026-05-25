# PROJECT_BRIEF.md

## 项目身份

- Project：go-scaffold
- Current Phase：pkg/utils 内部支撑测试完成
- Overall Status：PENDING_USER_CONFIRMATION
- Last Updated：2026-05-25
- Source Rule：`docs/ai/prompt.md`

## 摘要

- [CONFIRMED] 本仓库是已有 Go 脚手架项目，不是新建空项目。
- [CONFIRMED] 当前项目包含 `cmd/server`、`internal/app`、`internal/transport/http`、`internal/modules/demo` 和多个可复用 `pkg/*` 基础设施包。
- [CONFIRMED] README 说明当前阶段保留基础设施启动链路和 demo CRUD 示例，暂不实现 auth/rbac。
- [CONFIRMED] `go test ./... -count=1` 当前通过，可作为后续优化前的回归基线。
- [CONFIRMED] 本轮目标是生成和重写中文项目启动材料，重新启动“全项目分析与优化路线”主线。
- [RISK] 旧状态文档曾停留在插件系统 v1 收尾，容易让后续“下一步”偏离当前目标。
- [RISK] 根文档、模板和部分包 README 存在中英文混杂。

## 一句话目标

- [CONFIRMED] 系统分析 `go-scaffold` 当前优缺点，收拢不合理或不一致的设计边界，并制定可确认、可拆分、可验证的详细优化路线。

## 目标用户

| 用户类型 | 使用场景 | 痛点 | 优先级 | 状态 |
|---|---|---|---|---|
| 项目维护者 | 判断接下来优化什么 | 文档主线和边界需要统一 | P0 | [CONFIRMED] |
| 后续 AI Agent | 在上下文丢失后继续工作 | 必须从仓库文档恢复当前合法任务 | P0 | [CONFIRMED] |
| Go 开发者 | 使用脚手架作为项目基线 | 需要稳定扩展规则、示例和测试路径 | P1 | [INFERRED] |

## 当前优势

- [CONFIRMED] 服务启动链路清晰，`cmd/server` 负责入口，`internal/app` 负责装配。
- [CONFIRMED] 已有 health、ready 和 demo Todo CRUD API。
- [CONFIRMED] 基础设施包覆盖数据库、日志、HTTP server、缓存、国际化、存储、执行器、SQL 生成、插件、工具等能力。
- [CONFIRMED] 当前全量测试通过。
- [INFERRED] `internal/app` 的组合根适合成为后续边界治理的中心。

## 当前问题

- [RISK] 文档语言不统一，且旧模板偏英文。
- [RISK] 当前状态主线曾偏向插件系统扩展，而不是全项目优化。
- [RISK] `pkg/*` 公共性和兼容策略不清。
- [RISK] demo `AutoMigrate`、`initdb` 命令和 SQL 脚本职责未统一。
- [RISK] app/router/demo/config reload 等路径缺少集成测试。
- [RISK] JWT/auth 示例与“暂不实现 auth/rbac”的 README 说明存在范围漂移。
- [RISK] `pkg/sqlgen` 存在 TODO/未实现能力，需要确认是 unsupported、Backlog 还是后续实现。

## 当前范围

### P0

- [CONFIRMED] 生成/重写中文项目启动模板。
- [CONFIRMED] 重新设置当前合法任务为“项目优化启动确认”。
- [CONFIRMED] 保留插件系统 v1 为历史内容，不继续扩展。
- [CONFIRMED] 执行 `go test ./... -count=1` 作为文档变更后的回归基线。

### P1

- [CONFIRMED] 优化路线、`pkg/*` 定位、demo 模块定位、迁移策略和中文化范围已按默认值确认。
- [CONFIRMED] 已基于确认结果生成正式需求、架构、路线图、任务和时间切片。
- [CONFIRMED] 模块边界清单、测试矩阵、P1 切片、`types/*` 契约边界和 `pkg/plugin` 被动注册边界已完成。
- [CONFIRMED] TASK-P1-014 已完成，`pkg/utils` 已有最小确定性行为测试。
- [NEEDS_CONFIRMATION] 下一步等待用户选择进入 Phase 6 收尾、补 app/router/middleware 等集成测试或结束本轮。

### P2

- [DEFERRED] auth/rbac、部署流水线、性能测试、多租户、脚手架生成器、插件系统 rpc/ws/discovery 扩展。

## 非目标

- [CONFIRMED] 本轮不写 Go 代码。
- [CONFIRMED] 本轮不做重构。
- [CONFIRMED] 本轮不新增业务功能。
- [CONFIRMED] 本轮不实现 auth/rbac。
- [CONFIRMED] 本轮不部署、不执行生产命令、不执行不可逆迁移。

## 推荐默认

- [INFERRED] 采用“方案 A：治理优先”作为默认路线。
- [INFERRED] 保持当前高层依赖方向：`cmd -> internal/app -> internal/transport + internal/modules -> pkg`。
- [INFERRED] 所有未确认优化先进入 Backlog。

## 已确认事项

1. [CONFIRMED] 正式采用“治理优先”路线。
2. [CONFIRMED] `pkg/*` 采用混合策略。
3. [CONFIRMED] demo 模块暂定为长期标准示例。
4. [CONFIRMED] 迁移策略采用 dev-prod 分层。
5. [CONFIRMED] 中文化范围先覆盖根文档和模板，历史文档与包 README 分阶段处理。
6. [CONFIRMED] auth/JWT 先延后处理，不进入当前实现范围。
