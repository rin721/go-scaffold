# MODULES.md

## 模块边界清单状态

- 项目：go-scaffold
- 任务：TASK-OPT-003
- 时间切片：TS-OPT-003
- 状态：COMPLETED
- 最后更新：2026-05-26
- 原则：本文只记录事实、风险和优化候选项，不授权 Go 代码修改。

## 总体依赖方向

```text
cmd/server
  -> internal/app
      -> internal/config
      -> internal/middleware
      -> internal/modules/demo
      -> internal/transport/http
      -> pkg/*
types/*
```

- [CONFIRMED] `cmd/server` 是进程入口和 CLI 边界。
- [CONFIRMED] `internal/app` 是组合根，负责初始化、模式构建、热重载和生命周期。
- [CONFIRMED] `internal/modules/demo` 是当前唯一业务示例模块。
- [CONFIRMED] `pkg/*` 采用混合 API 策略，已逐包标注公共 API、内部支撑或待确认。
- [CONFIRMED] `types/*` 当前承载常量、错误码、响应类型和少量类型别名。

## 模块总览

| 模块 | 当前职责 | 边界定位 | 测试状态 | 主要风险 |
|---|---|---|---|---|
| `cmd/server` | CLI 入口、server/initdb/tests 命令、信号处理 | 进程入口 | tests 命令语义测试已补 | `runApp` 内部 `os.Exit` 仍需后续隔离才能测试启动失败路径 |
| `internal/app` | 组合根、启动模式、生命周期、热重载 | 应用装配层 | demo 迁移策略测试、真实 app 装配测试和 reload/config 分发测试已补 | 更大范围端到端启动仍需单独确认 |
| `internal/config` | 配置结构、加载、环境变量覆盖、热重载 | 配置边界 | 有测试 | TASK-P1-001、TASK-P1-002 和 TASK-P1-016 已收拢 copy/update、环境变量策略和配置变更 hook/reload 分发 |
| `internal/middleware` | Gin 中间件：i18n、CORS、logger、recovery、traceid | HTTP 横切层 | router 级 TraceID/CORS/Recovery 集成测试已补 | i18n/logger 细粒度行为仍可后续增强 |
| `internal/transport/http` | Gin router、health/ready、demo API 注册 | 传输层 | health/ready smoke test、demo HTTP 集成测试和 app 装配级 server/initdb 测试已补 | 真实 HTTP server 启动仍非当前测试目标 |
| `internal/modules/demo` | Todo 示例业务模块 | 标准示例模块 | service/repository CRUD 基线和 handler/router HTTP 集成已补 | 示例与生产约束需继续保持分离 |
| `pkg/*` | 可复用基础设施、公共工具和内部支撑工具 | 混合 API | 主要公共包与内部支撑路径已有最小测试 | [CONFIRMED] TASK-P1-017 已完成第一阶段包 README 中文化 |
| `types/*` | 公共常量、错误码、HTTP 响应结构、类型别名 | 跨层契约和 HTTP/Gin 响应契约 | `types/constants`、`types/errors`、`types/result` 有测试 | auth 错误码仅为预留契约，auth/rbac 不在当前功能范围 |

## `cmd/server`

### 当前职责

- [CONFIRMED] `main.go` 注册 `server`、`initdb`、`tests` 三个 CLI 命令。
- [CONFIRMED] `server` 命令读取配置路径并启动应用。
- [CONFIRMED] `initdb` 命令以 `ModeInitDB` 初始化应用并执行 demo schema 初始化。
- [CONFIRMED] `run.go` 处理 SIGINT/SIGTERM，并在退出时调用应用 shutdown。

### 优势

- [CONFIRMED] 入口职责整体较薄，主要负责 CLI、配置路径和信号。
- [CONFIRMED] 配置路径可通过 `--config` 或 `REI_CONFIG_PATH` 指定。

### 问题和风险

- [CONFIRMED] `tests` 命令已在 TASK-P1-006 改为真实 Go test 入口，默认执行 `go test ./...`。
- [RISK] `runApp` 内部直接 `os.Exit`，后续若要测试启动失败路径，需要额外隔离。
- [RISK] CLI 命令描述仍是英文，与中文项目方向不一致。

### 优化候选

- [CONFIRMED] 将 `tests` 命令改为真实测试入口，并新增最小命令语义测试。
- [DEFERRED] 为 CLI 命令增加最小单元测试或命令注册测试。

## `internal/app`

### 当前职责

- [CONFIRMED] `app.New` 负责加载 core、注册 logger handler、按模式构建 infra/modules/transport。
- [CONFIRMED] `modeapp` 支持 `server` 和 `initdb` 两种模式。
- [CONFIRMED] `initapp` 负责配置、日志、i18n、数据库、缓存、executor、storage、demo 模块和 HTTP server 构建。
- [CONFIRMED] `reloadapp` 响应配置变更并重载 cache、database、logger、executor、HTTP server、storage。
- [CONFIRMED] `lifecycleapp` 负责 HTTP server、storage、executor、cache、database、logger 的关闭顺序。

### 优势

- [CONFIRMED] 组合根集中，依赖装配点明确。
- [CONFIRMED] `Core -> Infrastructure -> Modules -> Transport` 的分层意图在代码注释中明确。
- [CONFIRMED] 关闭流程覆盖主要资源。

### 问题和风险

- [CONFIRMED] demo schema 迁移触发策略已在 TASK-P1-005 收拢：server-start/initdb 执行，reload 跳过。
- [CONFIRMED] `ModeInitDB` 明确作为 demo schema bootstrap 入口；生产/bootstrap 迁移框架仍延后。
- [CONFIRMED] reload/config 分发路径已在 TASK-P1-016 使用 fake 组件覆盖未变化、单组件变化、可选组件关闭和 database reload 不隐式迁移。
- [RISK] `initapp` 同时负责大量映射和构建逻辑，后续新增模块时容易膨胀。

### 优化候选

- [CONFIRMED] TASK-P1-016 已建立 app 装配链路测试，覆盖 server/initdb 模式。
- [CONFIRMED] 将 demo schema 自动迁移标注为 dev/demo 策略，并与 `initdb`/SQL 脚本分层。
- [CONFIRMED] TASK-P1-016 已为 reload 路径增加配置变更和分发集成测试。

## `internal/config`

### 当前职责

- [CONFIRMED] 定义顶层 `Config` 和 server/database/redis/logger/i18n/initdb/executor/storage/cors 子配置。
- [CONFIRMED] 使用 Viper 加载配置，支持 `${VAR:default}` 替换、`.env` 加载、环境变量覆盖和 fsnotify 监听。
- [CONFIRMED] `Manager` 使用 `atomic.Pointer` 存储配置快照。

### 优势

- [CONFIRMED] 配置结构集中且有校验接口。
- [CONFIRMED] 支持 shadow loading，变更配置无效时保留旧配置。

### 问题和风险

- [CONFIRMED] `manager.copyConfig` 字段覆盖问题已在 TASK-P1-001 修复并补测试。
- [CONFIRMED] 数据库环境变量策略已在 TASK-P1-002 收拢为 `DB_*` 优先，旧 `REI_APP_DB_*` 作为兼容 fallback。
- [CONFIRMED] `.env.example` 已移除未实现 JWT 示例，并与实际环境变量覆盖策略对齐。
- [CONFIRMED] 配置 reload 路径已由 TASK-P1-016 使用 fake 组件覆盖分发、关闭配置和 database reload 不隐式迁移。

### 优化候选

- [CONFIRMED] 统一环境变量命名策略，并同步 `.env.example`。
- [CONFIRMED] 修复并测试 `copyConfig`，确保热更新不会丢字段。
- [CONFIRMED] 清理重复 database override 行为，共用同一覆盖实现。
- [DEFERRED] 增加配置加载、环境覆盖、无效变更回滚测试。

## `internal/middleware`

### 当前职责

- [CONFIRMED] 提供 CORS、i18n、logger、recovery、traceid 中间件。
- [CONFIRMED] `internal/transport/http` 在 router 创建时装配中间件链。

### 优势

- [CONFIRMED] 横切能力集中在 middleware 包，HTTP router 不直接展开全部细节。
- [CONFIRMED] trace id 与 `types/result` 的错误响应存在衔接。

### 问题和风险

- [CONFIRMED] TASK-P1-015 已补路由级 TraceID、CORS 和 Recovery 链路测试。
- [RISK] trace id 初始化失败时存在 panic 路径，后续需要确认是否符合服务可用性要求。

### 优化候选

- [CONFIRMED] 增加 router middleware smoke test，覆盖 TraceID、CORS 和 Recovery 最小语义。
- [DEFERRED] 明确 trace id 初始化失败策略。

## `internal/transport/http`

### 当前职责

- [CONFIRMED] 创建 Gin engine。
- [CONFIRMED] 注册 `/health` 和 `/ready`。
- [CONFIRMED] 注册 `/api/v1/demo/todos` CRUD 路由。
- [CONFIRMED] `ReadyCheck` 使用 database ping。

### 优势

- [CONFIRMED] 传输层边界较清晰，业务 handler 通过依赖传入。
- [CONFIRMED] `TodoHandler` 为空时 demo 路由不注册，便于最小化装配。

### 问题和风险

- [CONFIRMED] health/ready HTTP smoke test 已在 TASK-P1-003 补齐。
- [CONFIRMED] ready 在数据库缺失时返回 `result.Success` 包裹 `not_ready`，HTTP 状态为 503；该响应语义已被测试固定。
- [CONFIRMED] TASK-P1-015 已补 demo Todo HTTP CRUD、CORS、TraceID 和 Recovery 集成测试。

### 优化候选

- [CONFIRMED] 增加 `/health`、`/ready` smoke test。
- [CONFIRMED] 增加 demo router 注册和 HTTP CRUD 集成测试。

## `internal/modules/demo`

### 当前职责

- [CONFIRMED] `model.Todo` 映射表名 `demo_todos`。
- [CONFIRMED] repository 只做 GORM 数据访问。
- [CONFIRMED] service 做标题校验、事务编排和 not found 归一。
- [CONFIRMED] handler 做请求绑定、ID 解析、HTTP 响应转换和错误映射。

### 优势

- [CONFIRMED] 分层符合 README 中 `model -> repository -> service -> handler` 的示例规则。
- [CONFIRMED] service 层已经避免把空标题写入数据库。
- [CONFIRMED] delete 前会确认记录存在。

### 问题和风险

- [CONFIRMED] demo service/repository CRUD 基线已在 TASK-P1-004 补齐。
- [CONFIRMED] TASK-P1-015 已补 demo handler/router HTTP 集成测试。
- [CONFIRMED] demo 自动迁移与生产迁移触发边界已在 TASK-P1-005 隔离。
- [RISK] demo 作为长期标准示例，需要补齐测试和文档，否则示例质量会影响后续模块。

### 优化候选

- [CONFIRMED] 增加 demo service/repository CRUD 基线测试。
- [CONFIRMED] 增加 demo CRUD HTTP 集成测试。
- [DEFERRED] 为新模块建立基于 demo 的模板规范。

## `pkg/*` API 分类

| 包 | 当前定位 | README 分类 | 测试 | 主要风险 |
|---|---|---|---|---|
| `pkg/cache` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] TASK-P1-013 已覆盖配置、Redis 读写、批量、计数器、TTL/Expire、缺失键和 reload 语义 |
| `pkg/cli` | 公共工具 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] flag parser、help 输出和错误包装已有最小包级测试 |
| `pkg/crypto` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | 当前稳定实现仅 bcrypt |
| `pkg/database` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | Hook/Reload/多驱动路径覆盖不足 |
| `pkg/executor` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] TASK-P1-012 已覆盖配置校验、任务执行、缺失池、过载、关闭、失败 reload 和 panic handler |
| `pkg/httpserver` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] TASK-P1-012 已覆盖构造、默认配置、配置错误、停止态 reload/shutdown 和已运行 start 拒绝路径 |
| `pkg/i18n` | 公共基础设施 API | [CONFIRMED] 已写入 | 无 | MustT panic 和加载错误路径需测试 |
| `pkg/logger` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | 文件输出/轮转路径覆盖有限 |
| `pkg/plugin` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] 被动注册边界、hooks、HTTP server helper 和 `RemoteHook` 已完成；rpc/ws/discovery、插件发现和 Go `.so` 明确延后 |
| `pkg/iam` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] 公共接口与 memory 实现已完成；JWT 中间件、数据库版权限和业务 RBAC 明确延后 |
| `pkg/sqlgen` | 公共工具 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] 高级查询、批量删除、DB reverse 和部分 rollback 边界已显式 unsupported / partial |
| `pkg/storage` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] TASK-P1-012 已覆盖内存文件系统读写、复制、MIME、Excel、图片和配置错误路径 |
| `pkg/utils` | 内部支撑工具包 | [CONFIRMED] 已写入 | 有 | [CONFIRMED] TASK-P1-014 已覆盖 Snowflake、地址校验、端口查找、设备 ID 和 i18n helper 最小行为；默认 Snowflake panic 策略保持不变 |
| `pkg/yaml2go` | 公共工具 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] 多文件生成、错误输入和配置校验已有最小包级测试 |

## `types/*` 边界

| 包 | 当前职责 | 风险 |
|---|---|---|
| `types/constants` | 应用名、命令名、超时、executor 池名等常量 | 与 `pkg/executor` 有依赖，属于跨层公共常量 |
| `types/errors` | 应用错误码和 `BizError` | auth/rbac 未实现，认证/授权错误码只是预留契约 |
| `types/result` | JSON 响应结构、Gin 响应 helper、分页结构、trace id 错误响应 | 依赖 Gin，属于 HTTP API 响应契约而不是纯类型包 |
| `types` | `Crypto` 类型别名、`CacheInjectable` 接口 | 依赖 `pkg/cache`、`pkg/crypto`，是有限聚合入口 |

- [CONFIRMED] TASK-P1-009 已完成 `types/*` 契约边界标注，详见 `docs/specs/types_contract_boundary.md`。

## `pkg/plugin` 注册边界

- [ACCEPT_WITH_RISK] 用户修正：`pkg/plugin` 不应主动注册插件服务，而应被动由插件服务进行注册。
- [CONFIRMED] 目标边界：`pkg/plugin` 只提供插件接口、local/http 插件实现和被动 registry/runtime；插件服务或宿主装配层负责构造插件并调用 `Register`。
- [CONFIRMED] TASK-P2-005 至 TASK-P2-007 已在该边界内扩展 hooks、标准钩子点、HTTP server helper 和 `RemoteHook`；`pkg/plugin` 不导入 IAM、配置、日志或 `internal/*`。
- [CONFIRMED] 历史 v1 API 中的 `Manager.Load(config)` 和 local factory 装配公共面已在 TASK-P1-010 移除，避免误解为由 `pkg/plugin` 主动发现并注册服务。
- [CONFIRMED] TASK-P1-010 已收拢该边界，rpc/ws/discovery 仍不进入当前范围。

## 设计边界冲突清单

| ID | 冲突 | 影响 | 建议 |
|---|---|---|---|
| BC-001 | `.env.example` 使用 `DB_*`，数据库 override 使用 `REI_APP_DB_*` | 环境变量文档与实现不一致 | [CONFIRMED] TASK-P1-002 已统一为 `DB_*` 优先，旧前缀兼容 |
| BC-002 | `copyConfig` 未覆盖完整配置字段 | 配置 Update/热更新可能丢字段 | [CONFIRMED] TASK-P1-001 已修复并补测试 |
| BC-003 | demo schema 在 server/initdb/reload 路径自动迁移 | dev/prod 迁移职责混乱 | [CONFIRMED] TASK-P1-005 已明确 server-start/initdb 执行、reload 跳过 |
| BC-004 | `cmd/server tests` 不运行测试 | CLI 语义误导 | [CONFIRMED] TASK-P1-006 已改为真实测试入口 |
| BC-005 | `pkg/sqlgen` README/代码存在未实现能力 | 公共工具 API 期望不稳定 | [CONFIRMED] TASK-P1-008 已显式 unsupported 或文档化 partial 能力 |
| BC-006 | 多个关键路径无测试 | 后续重构回归风险高 | [CONFIRMED] P1 已补配置、router、demo、app 装配、reload/config 和主要 `pkg/*` 最小测试 |

## 测试矩阵草案

| ID | 范围 | 建议测试 | 优先级 | 状态 |
|---|---|---|---|---|
| TM-001 | 全仓库 | `go test ./... -count=1` | P0 | [CONFIRMED] 当前通过 |
| TM-002 | app 启动 | 使用测试配置构建 `app.New`，验证 core/infra/modules/transport 非空 | P0 | [NOT_STARTED] |
| TM-003 | health/ready | 使用 `httptest` 验证 `/health`、数据库正常/缺失/失败时 `/ready` | P0 | [CONFIRMED] TASK-P1-003 已覆盖 |
| TM-004 | demo CRUD | 使用 SQLite 临时库跑 Create/List/Get/Update/Delete | P0 | [CONFIRMED] TASK-P1-004 已覆盖 service/repository 基线；TASK-P1-015 已覆盖 HTTP handler/router 集成 |
| TM-005 | config load/override | 验证 YAML、`${VAR:default}`、环境变量覆盖、无效配置报错 | P0 | [CONFIRMED] 环境覆盖已补测试 |
| TM-006 | config update/copy | 验证 `Update` 后不丢失 InitDB/Executor/Storage/CORS 等字段 | P0 | [CONFIRMED] TASK-P1-001 已覆盖 |
| TM-007 | migration boundary | 验证 server/initdb/reload 对 demo schema 的触发策略 | P1 | [CONFIRMED] TASK-P1-005 已覆盖 |
| TM-008 | pkg API | 为无测试的公共包补最小行为测试 | P1 | [CONFIRMED] `BL-020` 首批 TASK-P1-011、第二批 TASK-P1-012 和第三批 TASK-P1-013 已覆盖公共 `pkg/*` 行为测试；`pkg/utils` 内部支撑测试已由 TASK-P1-014 覆盖 |

## P1 优化候选项

| ID | 标题 | 依据 | 建议下一步 |
|---|---|---|---|
| OPT-P1-001 | 建立 app/router/demo/config 测试矩阵 | 多个关键路径无测试 | 生成测试任务和时间切片 |
| OPT-P1-002 | 统一配置环境变量策略 | `.env.example` 与实现不一致 | [CONFIRMED] 已完成 |
| OPT-P1-003 | 修复 `copyConfig` 字段覆盖 | 热更新可能丢配置 | [CONFIRMED] 已完成 |
| OPT-P1-004 | 明确迁移策略 | `AutoMigrate`、`initdb`、SQL 脚本职责冲突 | [CONFIRMED] 已编写迁移边界文档和测试 |
| OPT-P1-005 | 处理 `cmd/server tests` 命令语义 | 命令名与行为不符 | [CONFIRMED] 已改造为真实测试入口 |
| OPT-P1-006 | 为 `pkg/*` 完成公共/内部分类 | 混合 API 策略需要落地 | [CONFIRMED] 已更新包级 README、`ARCHITECTURE.md` 和本文 |
| OPT-P1-007 | 标注 `pkg/sqlgen` 未实现能力 | TODO/unsupported 边界不清 | [CONFIRMED] 已完成 |

## 当前完成判断

- [CONFIRMED] 模块职责清单已生成。
- [CONFIRMED] 设计边界冲突清单已生成。
- [CONFIRMED] TASK-P1-014 已完成，`pkg/utils` 内部支撑工具有最小确定性行为测试。
- [CONFIRMED] TASK-P1-015 已完成，router/middleware/demo HTTP 集成路径有最小测试覆盖。
- [CONFIRMED] TASK-PHASE6-001 已完成，Phase 6 收尾与交接已完成。
- [CONFIRMED] TASK-P1-016 已完成，app 装配、配置变更 hook 与 reload/config 剩余集成路径有最小测试覆盖。
- [CONFIRMED] TASK-P2-004 已补齐 Dockerfile、production Compose 示例、production 配置样例、统一 `deploy.sh` 部署入口和手动 production workflow 闸门；Docker build 已由用户在 Linux Docker 环境验证通过。
- [CONFIRMED] 测试矩阵草案已生成。
- [CONFIRMED] P1 优化候选项已生成。
- [CONFIRMED] 已按 P1 切片进行受控 Go 测试代码修改或文档分类：TASK-P1-001、TASK-P1-002、TASK-P1-003、TASK-P1-004、TASK-P1-005、TASK-P1-006、TASK-P1-007、TASK-P1-008 均已完成并验证通过。
- [CONFIRMED] `types/*` 契约边界已完成。
- [CONFIRMED] 用户已选择 A，`BL-020` 首批 `pkg/*` 行为测试已完成 TASK-P1-011 / TS-P1-011，第二批已完成 TASK-P1-012 / TS-P1-012，第三批 `pkg/cache` 已完成 TASK-P1-013 / TS-P1-013。
- [CONFIRMED] 用户选择 A，`BL-006` 第一阶段包 README 中文化已完成 TASK-P1-017 / TS-P1-017。
