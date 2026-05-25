# MODULES.md

## 模块边界清单状态

- 项目：go-scaffold
- 任务：TASK-OPT-003
- 时间切片：TS-OPT-003
- 状态：COMPLETED
- 最后更新：2026-05-25
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
- [CONFIRMED] `pkg/*` 采用混合 API 策略，需逐包标注公共 API、内部支撑或待确认。
- [CONFIRMED] `types/*` 当前承载常量、错误码、响应类型和少量类型别名。

## 模块总览

| 模块 | 当前职责 | 边界定位 | 测试状态 | 主要风险 |
|---|---|---|---|---|
| `cmd/server` | CLI 入口、server/initdb/tests 命令、信号处理 | 进程入口 | tests 命令语义测试已补 | `runApp` 内部 `os.Exit` 仍需后续隔离才能测试启动失败路径 |
| `internal/app` | 组合根、启动模式、生命周期、热重载 | 应用装配层 | demo 迁移策略测试已补 | reload 其他副作用测试不足 |
| `internal/config` | 配置结构、加载、环境变量覆盖、热重载 | 配置边界 | 有测试 | TASK-P1-001 和 TASK-P1-002 已收拢 copy/update 与环境变量策略；reload 路径仍需后续覆盖 |
| `internal/middleware` | Gin 中间件：i18n、CORS、logger、recovery、traceid | HTTP 横切层 | [no test files] | 中间件链路缺少路由级验证 |
| `internal/transport/http` | Gin router、health/ready、demo API 注册 | 传输层 | 有 health/ready smoke test | demo 路由缺少 integration 测试 |
| `internal/modules/demo` | Todo 示例业务模块 | 标准示例模块 | service/repository CRUD 基线已补 | handler/router 集成仍需后续覆盖，示例与生产约束需分离 |
| `pkg/*` | 可复用基础设施、公共工具和内部支撑工具 | 混合 API | 部分有测试 | [CONFIRMED] TASK-P1-007 已逐包分类；包 README 中英混杂仍待后续中文化 |
| `types/*` | 公共常量、错误码、响应结构、类型别名 | 跨层契约 | `types/constants` 有测试 | auth 错误码存在但 auth/rbac 不在当前功能范围 |

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
- [RISK] 热重载路径没有测试，尤其是数据库、HTTP server 和 storage 的重载副作用。
- [RISK] `initapp` 同时负责大量映射和构建逻辑，后续新增模块时容易膨胀。

### 优化候选

- [DEFERRED] 建立 app 装配链路测试，覆盖 server/initdb 模式。
- [CONFIRMED] 将 demo schema 自动迁移标注为 dev/demo 策略，并与 `initdb`/SQL 脚本分层。
- [DEFERRED] 为 reload 路径增加配置变更单元测试或集成测试。

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
- [RISK] 配置 reload 路径仍缺少测试文件。

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

- [RISK] 中间件链路没有路由级测试，无法确认 trace id、错误响应和 CORS 配置是否符合预期。
- [RISK] trace id 初始化失败时存在 panic 路径，后续需要确认是否符合服务可用性要求。

### 优化候选

- [DEFERRED] 增加 router middleware smoke test。
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
- [RISK] demo 路由注册和中间件链路没有集成测试。

### 优化候选

- [CONFIRMED] 增加 `/health`、`/ready` smoke test。
- [DEFERRED] 增加 demo router 注册测试。

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
- [RISK] demo handler/router 集成测试仍未覆盖。
- [CONFIRMED] demo 自动迁移与生产迁移触发边界已在 TASK-P1-005 隔离。
- [RISK] demo 作为长期标准示例，需要补齐测试和文档，否则示例质量会影响后续模块。

### 优化候选

- [CONFIRMED] 增加 demo service/repository CRUD 基线测试。
- [DEFERRED] 增加 demo CRUD 集成测试。
- [DEFERRED] 为新模块建立基于 demo 的模板规范。

## `pkg/*` API 分类

| 包 | 当前定位 | README 分类 | 测试 | 主要风险 |
|---|---|---|---|---|
| `pkg/cache` | 公共基础设施 API | [CONFIRMED] 已写入 | 无 | Redis 依赖路径无测试 |
| `pkg/cli` | 公共工具 API | [CONFIRMED] 已写入 | `cmd/server` 已覆盖命令使用；包自身无测试 | flag parser 和 help 输出无包级测试 |
| `pkg/crypto` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | 当前稳定实现仅 bcrypt |
| `pkg/database` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | Hook/Reload/多驱动路径覆盖不足 |
| `pkg/executor` | 公共基础设施 API | [CONFIRMED] 已写入 | 无 | reload/shutdown/overload/panic handler 无包级测试 |
| `pkg/httpserver` | 公共基础设施 API | [CONFIRMED] 已写入 | 无 | start/reload/shutdown 无包级测试 |
| `pkg/i18n` | 公共基础设施 API | [CONFIRMED] 已写入 | 无 | MustT panic 和加载错误路径需测试 |
| `pkg/logger` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | 文件输出/轮转路径覆盖有限 |
| `pkg/plugin` | 公共基础设施 API | [CONFIRMED] 已写入 | 有 | rpc/ws/discovery 明确延后 |
| `pkg/sqlgen` | 公共工具 API | [CONFIRMED] 已写入 | 有 | [CONFIRMED] 高级查询、批量删除、DB reverse 和部分 rollback 边界已显式 unsupported / partial |
| `pkg/storage` | 公共基础设施 API | [CONFIRMED] 已写入 | 无 | 文件监听、Excel、图片处理、复制边界无包级测试 |
| `pkg/utils` | 内部支撑工具包 | [CONFIRMED] 已写入 | 无 | snowflake 默认实例 panic 策略需确认；新增工具需谨慎 |
| `pkg/yaml2go` | 公共工具 API | [CONFIRMED] 已写入 | 无 | 原 `cmd/server tests` 演示依赖已移除；包自身仍无测试 |

## `types/*` 边界

| 包 | 当前职责 | 风险 |
|---|---|---|
| `types/constants` | 应用名、命令名、超时、executor 池名等常量 | 与 `pkg/executor` 有依赖，属于跨层公共常量 |
| `types/errors` | 应用错误码，含认证/授权段 | auth/rbac 未实现，但错误码已预留 |
| `types/result` | Gin 响应 helper、分页结构、trace id 错误响应 | 依赖 Gin，属于 HTTP 契约而不是纯类型包 |
| `types` | `Crypto` 类型别名、`CacheInjectable` 接口 | 依赖 `pkg/cache`、`pkg/crypto`，需确认是否继续作为公共聚合入口 |

- [CONFIRMED] 用户选择 A，`types/*` 边界已提升为 TASK-P1-009 / TS-P1-009。

## 设计边界冲突清单

| ID | 冲突 | 影响 | 建议 |
|---|---|---|---|
| BC-001 | `.env.example` 使用 `DB_*`，数据库 override 使用 `REI_APP_DB_*` | 环境变量文档与实现不一致 | [CONFIRMED] TASK-P1-002 已统一为 `DB_*` 优先，旧前缀兼容 |
| BC-002 | `copyConfig` 未覆盖完整配置字段 | 配置 Update/热更新可能丢字段 | [CONFIRMED] TASK-P1-001 已修复并补测试 |
| BC-003 | demo schema 在 server/initdb/reload 路径自动迁移 | dev/prod 迁移职责混乱 | [CONFIRMED] TASK-P1-005 已明确 server-start/initdb 执行、reload 跳过 |
| BC-004 | `cmd/server tests` 不运行测试 | CLI 语义误导 | [CONFIRMED] TASK-P1-006 已改为真实测试入口 |
| BC-005 | `pkg/sqlgen` README/代码存在未实现能力 | 公共工具 API 期望不稳定 | [CONFIRMED] TASK-P1-008 已显式 unsupported 或文档化 partial 能力 |
| BC-006 | 多个关键路径无测试 | 后续重构回归风险高 | P1 先建测试矩阵 |

## 测试矩阵草案

| ID | 范围 | 建议测试 | 优先级 | 状态 |
|---|---|---|---|---|
| TM-001 | 全仓库 | `go test ./... -count=1` | P0 | [CONFIRMED] 当前通过 |
| TM-002 | app 启动 | 使用测试配置构建 `app.New`，验证 core/infra/modules/transport 非空 | P0 | [NOT_STARTED] |
| TM-003 | health/ready | 使用 `httptest` 验证 `/health`、数据库正常/缺失/失败时 `/ready` | P0 | [CONFIRMED] TASK-P1-003 已覆盖 |
| TM-004 | demo CRUD | 使用 SQLite 临时库跑 Create/List/Get/Update/Delete | P0 | [CONFIRMED] TASK-P1-004 已覆盖 service/repository 基线 |
| TM-005 | config load/override | 验证 YAML、`${VAR:default}`、环境变量覆盖、无效配置报错 | P0 | [IN_PROGRESS] 环境覆盖已补测试 |
| TM-006 | config update/copy | 验证 `Update` 后不丢失 InitDB/Executor/Storage/CORS 等字段 | P0 | [CONFIRMED] TASK-P1-001 已覆盖 |
| TM-007 | migration boundary | 验证 server/initdb/reload 对 demo schema 的触发策略 | P1 | [CONFIRMED] TASK-P1-005 已覆盖 |
| TM-008 | pkg API | 为无测试的公共包补最小行为测试 | P1 | [CONFIRMED] TASK-P1-007 已完成分类；测试补齐仍延后 |

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
- [CONFIRMED] 测试矩阵草案已生成。
- [CONFIRMED] P1 优化候选项已生成。
- [CONFIRMED] 已按 P1 切片进行受控 Go 测试代码修改或文档分类：TASK-P1-001、TASK-P1-002、TASK-P1-003、TASK-P1-004、TASK-P1-005、TASK-P1-006、TASK-P1-007、TASK-P1-008 均已完成并验证通过。
- [CONFIRMED] 下一阶段已选择提升 `types/*` 契约边界，当前合法下一步为 TASK-P1-009 / TS-P1-009。
