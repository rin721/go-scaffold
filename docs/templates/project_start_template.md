# 项目启动模板

## 1. 项目目标

- [CONFIRMED] 项目名称：`go-scaffold`。
- [CONFIRMED] 项目类型：已有 Go 后端脚手架，不是空项目。
- [CONFIRMED] 本轮目标：分析当前项目的优缺点，统一不合理或不一致的设计边界，并制定可验证、可交接的详细优化路线。
- [CONFIRMED] 本轮不写 Go 代码、不做重构、不新增业务功能。
- [INFERRED] 默认推进路线：先治理文档和状态，再确认架构边界，再拆任务和时间切片，最后进入代码优化。

## 2. 当前项目事实

| 分类 | 内容 | 状态 |
|---|---|---|
| 启动链路 | `cmd/server` 作为进程入口，`internal/app` 负责组合装配 | [CONFIRMED] |
| HTTP | 使用 Gin，`internal/transport/http` 提供 `/health`、`/ready` 和 demo API | [CONFIRMED] |
| 示例模块 | `internal/modules/demo` 提供 Todo CRUD 示例 | [CONFIRMED] |
| 基础设施 | 已有 database、logger、httpserver、cache、i18n、storage、executor、sqlgen、plugin、utils 等包 | [CONFIRMED] |
| 测试基线 | `go test ./... -count=1` 当前通过 | [CONFIRMED] |
| 当前文档 | 根文档和模板已存在，但主线曾停留在插件系统 v1 收尾 | [CONFIRMED] |

## 3. 当前优势

- [CONFIRMED] 项目可以通过单一命令启动服务：`go run ./cmd/server server`。
- [CONFIRMED] README 已描述目录边界、启动方式、Demo Todo API 和测试命令。
- [CONFIRMED] 应用装配链路已经有清晰分层：核心配置、基础设施、模块、传输层。
- [CONFIRMED] demo 模块展示了 handler、service、repository、model 的基础分层。
- [CONFIRMED] 当前全量 Go 测试通过，可作为后续优化前的回归基线。

## 4. 当前问题

- [RISK] 文档语言不统一，根文档、模板和部分包 README 混用中文和英文。
- [RISK] 当前状态主线曾偏向“插件系统下一方向选择”，不符合本轮“全项目治理与优化路线”的目标。
- [RISK] `pkg/*` 看起来像可复用公共库，但尚未确认公共 API、内部支撑或混合策略。
- [RISK] 数据库迁移边界不清，demo `AutoMigrate`、`initdb` 命令和 SQL 脚本并存。
- [RISK] 测试覆盖偏包级，多个 app/router/demo/config reload 路径显示 `[no test files]`。
- [RISK] `.env.example` 包含 JWT 示例，而 README 明确当前暂不实现 auth/rbac，存在范围暗示漂移。

## 5. 非目标

- [CONFIRMED] 不立即重写项目。
- [CONFIRMED] 不新增业务功能。
- [CONFIRMED] 不实现 auth/rbac。
- [CONFIRMED] 不部署、不执行生产命令、不执行不可逆数据库迁移。
- [CONFIRMED] 不扩展插件系统的 rpc/ws/discovery 能力，除非后续单独确认。
- [CONFIRMED] 不把未验证的代码任务标记为 `COMPLETED`。

## 6. 默认路线

| 阶段 | 目标 | 退出条件 | 状态 |
|---|---|---|---|
| Phase 0 项目重新启动 | 生成中文启动模板，收拢当前主线 | 模板齐全，状态进入 `PENDING_USER_CONFIRMATION` | [CONFIRMED] |
| Phase 1 需求确认 | 确认优化方向、中文化范围、包 API 策略、迁移策略 | `REQUIREMENTS.md` 与 `ACCEPTANCE.md` 被确认 | [NEEDS_CONFIRMATION] |
| Phase 2 架构确认 | 明确目录、模块、包、迁移、插件、配置热更新边界 | `ARCHITECTURE.md` 与 `DECISIONS.md` 被确认 | [NEEDS_CONFIRMATION] |
| Phase 3 任务拆分 | 将优化路线拆成任务和时间切片 | `TASKS.md`、`TIME_SLICES.md` 唯一合法任务明确 | [NEEDS_CONFIRMATION] |
| Phase 4 代码优化 | 只按已确认时间切片修改代码 | 每个切片有测试证据和交接记录 | [DEFERRED] |

## 7. 待确认事项

| ID | 问题 | 默认值 | 影响 | 状态 |
|---|---|---|---|---|
| START-Q001 | 是否采用“治理优先”作为正式优化路线？ | 是 | 决定路线图和任务顺序 | [NEEDS_CONFIRMATION] |
| START-Q002 | `pkg/*` 应按公共复用库、内部支撑或混合策略维护？ | 混合策略 | 决定兼容性和重构边界 | [NEEDS_CONFIRMATION] |
| START-Q003 | demo 模块是长期标准示例、临时占位还是后续移除？ | 长期标准示例 | 决定测试和文档样板 | [NEEDS_CONFIRMATION] |
| START-Q004 | 迁移策略以 `AutoMigrate`、SQL 脚本还是 dev-prod 分层为准？ | dev-prod 分层 | 决定数据库初始化边界 | [NEEDS_CONFIRMATION] |
| START-Q005 | 中文化范围到哪里？ | 根文档加模板优先 | 决定文档工作量和验收 | [NEEDS_CONFIRMATION] |

## 8. 当前完成判断

- [CONFIRMED] 本模板用于项目重新启动阶段。
- [CONFIRMED] 后续下一步应等待用户确认，而不是直接进入代码实现。
- [CONFIRMED] 插件系统 v1 保留为历史内容和 Backlog，不作为当前主线继续扩展。
