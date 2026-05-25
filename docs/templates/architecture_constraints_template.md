# 架构约束模板

## 1. 约束状态

- 当前状态：`PENDING_USER_CONFIRMATION`
- 适用范围：后续需求、架构设计、任务拆分、代码优化和验收。
- 来源：仓库现状、README、`docs/ai/prompt.md`、本轮项目优化启动计划。

## 2. 依赖方向

- [CONFIRMED] `cmd/server` 只作为进程入口，负责 CLI 命令、参数和信号处理。
- [CONFIRMED] `internal/app` 是组合根和生命周期边界，负责装配核心配置、基础设施、模块和传输层。
- [CONFIRMED] `internal/transport/http` 负责 HTTP router、health/ready 和 API 路由注册。
- [CONFIRMED] `internal/modules/*` 承载业务模块，当前 demo 模块只是示例业务。
- [CONFIRMED] `pkg/*` 不应依赖 `internal/*`。
- [NEEDS_CONFIRMATION] `pkg/*` 是公共复用 API、项目内部支撑还是混合策略。

## 3. 分层规则

- [CONFIRMED] handler 层只处理参数绑定、HTTP 状态码和响应转换。
- [CONFIRMED] service 层处理业务校验、事务编排和 repository 调用。
- [CONFIRMED] repository 层处理 GORM 数据访问，不写业务判断。
- [NEEDS_CONFIRMATION] 上述 demo 分层规则是否作为所有未来模块的强制模板。

## 4. 配置边界

- [INFERRED] 配置结构、加载、环境变量覆盖和校验应继续集中在 `internal/config`。
- [INFERRED] 从 app 配置映射到 `pkg/*` 配置的逻辑应继续留在 `internal/app/initapp`。
- [RISK] 配置默认值、环境覆盖、热更新和包级 README 示例可能存在漂移，需要统一说明。

## 5. 数据库迁移边界

- [CONFIRMED] demo 当前在模块装配时执行 `AutoMigrate(&model.Todo{})`。
- [CONFIRMED] 项目同时存在 `cmd/server initdb` 相关命令和 `scripts/initdb` SQL 脚本。
- [RISK] `AutoMigrate`、`initdb` 和 SQL 脚本职责未统一，可能导致开发态和生产态边界混乱。
- [NEEDS_CONFIRMATION] 推荐默认：dev-prod 分层策略，开发/demo 可用 `AutoMigrate`，生产/bootstrap 使用显式 SQL 或迁移流程。

## 6. Demo 边界

- [CONFIRMED] `internal/modules/demo` 当前提供 Todo CRUD 示例。
- [INFERRED] demo 可作为新业务模块分层样板。
- [RISK] demo 的便利性实现不应无确认地变成生产模块强约束。
- [NEEDS_CONFIRMATION] demo 是长期标准示例、临时占位还是后续移除。

## 7. 包 API 边界

- [CONFIRMED] `pkg/plugin` v1 local/http 已完成并通过测试。
- [CONFIRMED] 插件系统 rpc/ws/discovery 扩展暂不属于当前主线。
- [RISK] `pkg/sqlgen` 存在 TODO/未实现能力，需要明确 unsupported、Backlog 或后续任务。
- [RISK] 包 README 中英混杂，与“中文项目”的要求不完全一致。
- [NEEDS_CONFIRMATION] 需要为 `pkg/*` 建立公共 API、内部支撑或混合策略。

## 8. 安全和范围约束

- [CONFIRMED] 不输出真实 `.env` 密钥。
- [CONFIRMED] 不自动部署。
- [CONFIRMED] 不自动执行不可逆迁移。
- [CONFIRMED] 不在本轮实现 auth/rbac。
- [RISK] `.env.example` 中 JWT 示例可能暗示当前未实现的能力，需后续清理、保留占位或提升为正式需求。

## 9. 后续变更门禁

- [CONFIRMED] 任何代码变更必须映射到已确认需求、架构边界、任务和时间切片。
- [CONFIRMED] 任何代码任务必须记录验证命令和结果。
- [CONFIRMED] 未确认的新优化默认进入 `BACKLOG.md`。
