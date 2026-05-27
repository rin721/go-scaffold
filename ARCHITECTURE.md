# ARCHITECTURE.md

## 架构状态

- 项目：go-scaffold
- 当前焦点：项目治理与优化路线 v1
- 状态：COMPLETED
- 最后更新：2026-05-26

## 已确认架构原则

- [CONFIRMED] 采用治理优先路线：先稳定文档、边界、任务和测试矩阵，再改代码。
- [CONFIRMED] 保持当前依赖方向：`cmd -> internal/app -> internal/transport + internal/modules -> pkg`。
- [CONFIRMED] `internal/app` 是组合根和生命周期边界。
- [CONFIRMED] `internal/modules/demo` 暂定为长期标准示例。
- [CONFIRMED] `pkg/*` 采用混合策略，后续必须逐包标注公共 API 或内部支撑定位。
- [CONFIRMED] 迁移采用 dev-prod 分层策略：server-start/initdb 可执行 demo `AutoMigrate`，reload 不执行隐式 schema 变更，生产/bootstrap 迁移框架延后。

## 目标依赖方向

```text
cmd/server
  -> internal/app
      -> internal/config
      -> internal/modules/*
      -> internal/transport/http
      -> pkg/*
types/*
```

## 已确认边界

| 边界 | 结论 | 后续动作 | 状态 |
|---|---|---|---|
| `cmd/server` | 进程入口、CLI 命令、信号处理 | 检查是否有业务逻辑外溢 | [CONFIRMED] |
| `internal/app` | 组合根、生命周期、跨层装配 | 生成装配链路清单 | [CONFIRMED] |
| `internal/transport/http` | HTTP router、middleware、health/ready、API 注册 | [CONFIRMED] health/ready、demo Todo HTTP 集成测试和 app 装配级路径均已补最小测试；TASK-P1-016 覆盖真实 app server/initdb 装配与 reload/config 分发 | [CONFIRMED] |
| `internal/modules/demo` | 长期标准示例 | 生成 demo 分层验收和测试路线 | [CONFIRMED] |
| `pkg/*` | 混合策略 | [CONFIRMED] TASK-P1-007 已逐包分类为公共基础设施 API、公共工具 API 或内部支撑工具包 | [CONFIRMED] |
| 数据库迁移 | dev-prod 分层 | [CONFIRMED] TASK-P1-005 已明确 demo `AutoMigrate`、`initdb`、reload 职责；生产迁移框架仍延后 | [CONFIRMED] |
| 插件系统 | v1 local/http 保留；注册责任已收拢为被动 registry/runtime；TASK-P2-005 至 TASK-P2-007 已增加 hooks、HTTP server helper 和 `RemoteHook` | [CONFIRMED] rpc/ws/discovery、插件发现和 Go `.so` 插件仍留在 Backlog | [CONFIRMED] |
| IAM/auth/JWT | `pkg/iam` 公共接口与 memory 实现已完成；JWT 中间件和业务 RBAC 仍不实现 | 后续如需业务登录、HTTP 中间件或数据库版权限，必须单独提升任务 | [CONFIRMED] |
| CI/CD 与部署 | 先建立非生产质量门禁、手动部署说明、远程部署显式参数契约、手动远程部署 workflow、Docker production 制品和远程 Linux 统一 `deploy.sh` 入口；真实生产运行仍需单独确认 | [CONFIRMED] TASK-P2-001 已新增 CI workflow 和部署说明；TASK-P2-002 已新增 `deploy.sh` / `script/install.sh` 显式参数契约；TASK-P2-003 已新增手动 staging 远程部署 workflow；TASK-P2-004 已补 Dockerfile、production Compose 示例、统一 `deploy.sh` 部署入口、手动 production 闸门并完成 Docker build 验证；镜像发布和真实 production 运行仍需单独确认 | [CONFIRMED] |

## 需要详细分析的模块

| 模块 | 主要问题 | 下一步 |
|---|---|---|
| `internal/app` | 装配、reload、mode、lifecycle 边界需形成清单 | TASK-OPT-003 |
| `internal/config` | 环境覆盖、热更新和默认值职责需统一说明 | TASK-OPT-003 |
| `internal/transport/http` | health/ready、demo 路由和 app 装配级 server/initdb 路径均已补最小测试 | 继续保持不启动真实 HTTP server 的测试边界 |
| `internal/modules/demo` | 示例职责与生产约束需分离 | TASK-OPT-003 |
| `pkg/*` | 公共/内部定位已在 TASK-P1-007 标记，`pkg/sqlgen` unsupported 边界已在 TASK-P1-008 标记 | 后续破坏性重构或新能力实现仍需单独确认 |
| `types/*` | 跨层公共契约、HTTP/Gin 响应契约和有限聚合入口 | [CONFIRMED] TASK-P1-009 已明确 `types/result`、`types/errors`、`types/constants` 和根 `types` 边界 |

## 代码变更门禁

- [CONFIRMED] 未完成模块边界清单前，不进入代码优化。
- [CONFIRMED] 未确认 P1 执行顺序前，不修改 app/router/demo/config/migration 核心路径。
- [CONFIRMED] 包 API 分类已完成；任何 `pkg/*` 破坏性重构仍需单独任务确认。

## `pkg/*` API 分类

| 包 | 分类 | 稳定边界 | 当前风险 | 状态 |
|---|---|---|---|---|
| `pkg/cache` | 公共基础设施 API | `Cache`、`Config`、`DefaultConfig`、`NewRedis` | [CONFIRMED] TASK-P1-013 已覆盖配置、读写、批量、计数器、过期、缺失键和 reload 隔离行为 | [CONFIRMED] |
| `pkg/cli` | 公共工具 API | `App`、`Command`、`Context`、`Flag`、错误类型、`GetExitCode` | [CONFIRMED] flag parser、help 输出和错误包装已有最小包级测试 | [CONFIRMED] |
| `pkg/crypto` | 公共基础设施 API | `Crypto`、bcrypt 实现、`Config`、Option 配置函数 | 当前稳定实现仅覆盖 bcrypt | [CONFIRMED] |
| `pkg/database` | 公共基础设施 API | `Database`、`Reloader`、事务接口、`Config`、`New`、`NewWithHooks` | Hook、Reload、多驱动路径覆盖有限 | [CONFIRMED] |
| `pkg/executor` | 公共基础设施 API | `Manager`、`Config`、`PoolName`、`NewManager` | [CONFIRMED] TASK-P1-012 已覆盖配置校验、任务执行、缺失池、过载、关闭、失败 reload 和 panic handler | [CONFIRMED] |
| `pkg/httpserver` | 公共基础设施 API | `HTTPServer`、`Config`、`Handler`、`New`、错误类型 | [CONFIRMED] TASK-P1-012 已覆盖构造、默认配置、配置错误、停止态 reload/shutdown 和已运行 start 拒绝路径 | [CONFIRMED] |
| `pkg/i18n` | 公共基础设施 API | `I18n`、`Config`、`New`、`Default`、语言常量 | [CONFIRMED] TASK-P1-011 已覆盖 JSON/YAML 加载、模板渲染、fallback、`MustT` panic 和加载错误路径 | [CONFIRMED] |
| `pkg/logger` | 公共基础设施 API | `Logger`、`Reloader`、`Config`、`New`、`Default` | 文件输出和轮转路径覆盖有限 | [CONFIRMED] |
| `pkg/plugin` | 公共基础设施 API | v1 local/http runtime、被动 `Manager.Register` registry、`Plugin`、`Request`、`Response`、`Definition`、`NewHTTP`、`hooks`、HTTP server helper、`RemoteHook` | [CONFIRMED] hook-aware runtime 已完成；rpc/ws/discovery、插件发现和 Go `.so` 插件延后 | [CONFIRMED] |
| `pkg/iam` | 公共基础设施 API | `Principal`、`Credential`、`Policy`、`Authenticator`、`Authorizer`、`Service`、context helper、memory service | [CONFIRMED] 当前是公共接口和内存实现，不是 JWT 中间件或完整业务 RBAC | [CONFIRMED] |
| `pkg/sqlgen` | 公共工具 API | 当前测试覆盖的 SQL 构建、解析、事务和模板能力；unsupported 路径显式返回 `ErrCodeUnsupportedOperation` 或在 README 标注 partial | 高级查询、批量删除、DB reverse 和部分 rollback 能力不属于当前稳定能力 | [CONFIRMED] |
| `pkg/storage` | 公共基础设施 API | `Storage`、`Config`、`New`、文件读写、复制、监听和 MIME/媒体辅助能力 | [CONFIRMED] TASK-P1-012 已覆盖内存文件系统读写、复制、MIME、Excel、图片和配置错误路径 | [CONFIRMED] |
| `pkg/utils` | 内部支撑工具包 | 当前供 `internal/*` 和少量 `types/*` 使用的 ID、地址、端口、设备 ID、i18n helper | [CONFIRMED] TASK-P1-014 已覆盖最小行为；默认 Snowflake panic 策略保持不变 | [CONFIRMED] |
| `pkg/yaml2go` | 公共工具 API | `Converter`、`Config`、`New`、`Convert` 返回结构 | [CONFIRMED] 包自身已有最小行为测试；文件写入和 CLI 管理仍非目标 | [CONFIRMED] |

## `types/*` 契约分类

| 包 | 分类 | 稳定边界 | 当前风险 | 状态 |
|---|---|---|---|---|
| `types/result` | HTTP API 响应契约 | `Result`、`Success`、`Error`、`ErrorWithTrace`、分页结构和 Gin 响应 helper | 依赖 Gin，不能作为纯领域类型包使用 | [CONFIRMED] |
| `types/errors` | 错误码和业务错误契约 | 错误码分段、`BizError`、错误链 | auth/rbac 错误码是预留契约，不代表当前已实现 auth/rbac | [CONFIRMED] |
| `types/constants` | 跨层运行常量契约 | 应用命令、默认配置路径、关闭超时、cache key、executor pool 名称 | 常量修改会影响 cmd/internal/pkg 使用方 | [CONFIRMED] |
| `types` | 有限聚合入口 | `Crypto` 别名、`CacheInjectable` 接口 | 新增聚合类型需单独确认 | [CONFIRMED] |

## 下一架构任务

- [CONFIRMED] TASK-P1-005 已完成 demo 迁移触发边界收拢。
- [CONFIRMED] TASK-P1-007 已完成 `pkg/*` API 分类。
- [CONFIRMED] TASK-P1-008 已完成 `pkg/sqlgen` unsupported 边界标注。
- [CONFIRMED] TASK-P1-009 已明确 `types/*` 契约边界。
- [CONFIRMED] TASK-P1-010 已收拢 `pkg/plugin` 被动注册边界。
- [CONFIRMED] 用户选择 A，`BL-020` 首批 `pkg/*` 行为测试已完成 TASK-P1-011，第二批已完成 TASK-P1-012，第三批 `pkg/cache` 已完成 TASK-P1-013。
- [CONFIRMED] TASK-P2-005 至 TASK-P2-010 已完成插件 hooks、HTTP 远程插件、IAM 公共接口、配置接入和 app reload/lifecycle 组装。
- [CONFIRMED] 用户选择 B，`BL-023` `pkg/utils` 内部支撑测试已完成 TASK-P1-014。
- [CONFIRMED] 用户选择 B，`BL-002` router/middleware/demo HTTP 集成测试已完成 TASK-P1-015。
- [CONFIRMED] 用户明确要求实施 TASK-P1-016，app 装配、配置变更 hook 与 reload/config 剩余集成测试已完成。
- [CONFIRMED] 用户选择 A，`BL-006` 第一阶段包 README 中文化已完成 TASK-P1-017。
- [CONFIRMED] TASK-INFRA-003 已修复 TASK-P1-016/017 后背景文档中的旧状态漂移。
- [CONFIRMED] 用户选择 D，CI 质量门禁与部署说明首切片已完成 TASK-P2-001。
- [CONFIRMED] 用户选择 C，进入真实 CD / 镜像发布 / 远程部署自动化范围确认；TASK-P2-002 已完成远程部署 `.env` 模板，TASK-P2-003 已完成手动 staging 远程部署 workflow，TASK-P2-004 已新增统一 `deploy.sh` 部署入口。
- [DEFERRED] 生产迁移框架需要单独需求和架构确认，不属于当前切片。
