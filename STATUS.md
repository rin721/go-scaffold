# STATUS.md

## Current Auth Public API Rework Scope

- 2026-05-28: TASK-P2-023 / TS-P2-023 COMPLETED.
- User correction: after RBAC was moved to `pkg/rbac`, auth token capability also needs a public `pkg` infrastructure API before business code uses it.
- Review Result: ACCEPT_WITH_RISK.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Scope completed: introduced `pkg/auth` as a public JWT token API backed by `github.com/golang-jwt/jwt/v5`, replaced hand-written internal token signing/parsing with that public API, and kept main-service user business code responsible only for mapping users to token claims.
- Allowed scope: `go.mod`, `go.sum`, `pkg/auth/**/*`, `internal/modules/user/service/**/*`, `internal/app/initapp/**/*`, focused router/app tests, and project status documents.
- Strict non-goals: no refresh-token/session revocation, audit logging, password reset, external IAM replacement, production secret manager, production migration framework, real users/passwords/secrets, deployment, plugin transport work, or unrelated router/user CRUD rewrites.
- Verification: `go test ./pkg/auth -count=1`, `go test ./internal/modules/user/... -count=1`, `go test ./internal/app/initapp -count=1`, `go test ./internal/app/... -count=1`, `go test ./internal/transport/http -count=1`, `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.

## Current RBAC Public API Rework Scope

- 2026-05-28: TASK-P2-022 / TS-P2-022 COMPLETED.
- User correction: RBAC must be packaged as a public `pkg` infrastructure API first, then business code must call that public API.
- Review Result: ACCEPT_WITH_RISK.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Scope completed: moved the Casbin-backed RBAC authorizer from `internal/modules/user/rbac` to `pkg/rbac`, added public package docs and tests, kept the Casbin model under `configs/rbac_model.conf`, and updated main-service user authorization plus app composition to depend on `pkg/rbac`.
- Allowed scope: `pkg/rbac/**/*`, `internal/modules/user/**/*`, `internal/app/initapp/**/*`, `configs/rbac_model.conf`, config examples and config tests if needed, focused app/router tests, and project status documents.
- Strict non-goals: no external IAM replacement, no refresh/session/audit/password-reset work, no production migration framework, no real users/passwords/secrets, no deployment, no plugin transport work, and no unrelated router/user CRUD rewrites.
- Verification: `go test ./pkg/rbac -count=1`, `go test ./internal/config -count=1`, `go test ./internal/modules/user/... -count=1`, `go test ./internal/app/initapp -count=1`, `go test ./internal/app/... -count=1`, `go test ./internal/transport/http -count=1`, `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.

## Current RBAC Library Rework Scope

- 2026-05-28: TASK-P2-021 / TS-P2-021 COMPLETED.
- User correction: RBAC must be encapsulated through a mainstream library before business authorization logic is implemented.
- Review Result: ACCEPT_WITH_RISK.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Scope completed: replaced the hand-written permission matcher in the main-service user/RBAC module with a Casbin-backed authorizer wrapper, put the Casbin RBAC model under `configs/rbac_model.conf`, kept existing DB-backed role/permission CRUD and config seed behavior, and added focused tests.
- Allowed scope: `go.mod`, `go.sum`, `configs/rbac_model.conf`, `configs/config.example.yaml`, `configs/config.yaml`, `internal/config/**/*`, `internal/modules/user/**/*`, `internal/app/initapp/**/*`, focused app/router tests, and project status documents.
- Strict non-goals: no external IAM replacement, no OPA server, no refresh/session/audit/password-reset work, no production migration framework, no real users/passwords/secrets, no deployment, no plugin transport work, and no unrelated router/user CRUD rewrites.
- Verification: `go test ./internal/config -count=1`, `go test ./internal/modules/user/... -count=1`, `go test ./internal/app/initapp -count=1`, `go test ./internal/app/... -count=1`, `go test ./internal/transport/http -count=1`, `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.

## Current RBAC Config Scope

- 2026-05-28: TASK-P2-020 / TS-P2-020 COMPLETED.
- User intent: put RBAC configuration under `configs` instead of leaving default roles/permissions only implicit in code.
- Review Result: ACCEPT_WITH_RISK.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Scope completed: added a recoverable RBAC seed configuration shape under `internal/config`, documented safe role/permission seed entries in `configs/config.example.yaml` and `configs/config.yaml`, and applied configured roles, permissions, and role-permission grants idempotently during main service startup when enabled.
- Verification: `go test ./internal/config -count=1`, `go test ./internal/modules/user/... -count=1`, `go test ./internal/app/initapp -count=1`, `go test ./internal/app/... -count=1`, `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Strict non-goals: no OPA/Casbin, no external IAM replacement, no refresh/session/audit/password-reset work, no production migration framework, no real users/passwords/secrets, no deployment, no plugin transport work, and no unrelated user/router rewrites.

## Current Auth Hardening Scope

- 2026-05-28: TASK-P2-019 / TS-P2-019 COMPLETED.
- User selection: `2+4`, interpreted as promoting both `BL-026` auth hardening and `BL-028` plugin WS/RPC/discovery direction.
- Review Result: ACCEPT_WITH_RISK.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Scope completed: added stable auth token secret and token TTL configuration for the existing main-service user/auth/RBAC module, wired it into `internal/app/initapp`, documented example config/env placeholders, and covered config + wiring tests.
- Verification: `go test ./internal/config -count=1`, `go test ./internal/app/initapp -count=1`, `go test ./internal/app/... -count=1`, `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Deferred from this slice: `BL-028` real plugin WS/RPC transport, heartbeat, reconnect, and persistent discovery stay queued until auth/session boundaries are stable.
- Strict non-goals: refresh tokens, session revoke, audit logging, password reset, external IAM, OPA/Casbin, production migrations, real secrets, deployment, and plugin transport implementation.

## Current User Scope Review

- 2026-05-28: TASK-P2-018 / TS-P2-018 started.
- User intent: confirm whether the main service already implements user management, then implement complete user + auth + RBAC in the main service.
- Current finding: the main service had no dedicated `internal/modules/user` module, no `/api/v1/users` routes, and no login/token/RBAC HTTP flow; existing `pkg/iam` remains a public infrastructure interface and memory implementation, not database-backed business user management.
- Review Result: ACCEPT_WITH_RISK.
- Current legal work: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Scope: implement a main-service user/auth/RBAC module with database-backed users, roles, permissions, password hashing, bearer token login, authentication middleware, permission-gated routes, server-start sqlgen schema bootstrap, and focused tests.
- Strict non-goals: production-grade secret management, refresh-token/session revocation, OPA/Casbin, production migration framework, real secrets, deployment, plugin transport changes, and broad cmd/entrypoint rewrites.

## Latest Current Slice

- 2026-05-28: TASK-P2-018 / TS-P2-018 COMPLETED.
- User correction accepted with risk: implement complete main-service user + auth + RBAC rather than only user CRUD.
- Implemented `internal/modules/user` with users, roles, permissions, user-role and role-permission assignments, password hashing through `pkg/crypto`, HMAC bearer tokens, authentication middleware, and route-level permission checks.
- Wired `/api/v1/auth/register`, `/api/v1/auth/login`, `/api/v1/auth/me`, `/api/v1/users`, `/api/v1/roles`, and `/api/v1/permissions` into the main router, and added sqlgen bootstrap for user/RBAC tables at server start.
- Verification: `go test ./internal/modules/user/... -count=1`, `go test ./internal/transport/http -count=1`, `go test ./internal/app/... -count=1`, `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Residual non-goals: production secret/session management, refresh-token/session revocation, audit logging, password reset, OPA/Casbin, external IAM, production migrations, deployment, and plugin transport changes.

- 2026-05-28: TASK-INFRA-004 / TS-INFRA-004 COMPLETED.
- GitHub Actions push CI run `26531295923` failed in `CI / Go verify` at the `Build server` step.
- Root cause: `.github/workflows/ci.yml` still built the removed `./cmd/server` package, while the current repository entrypoint is `./cmd/main`.
- Fix: aligned the CI build step to `go build -mod=readonly -o /tmp/go-scaffold-server ./cmd/main`.
- Verification: GitHub job logs confirmed tests passed and build failed on `stat .../cmd/server: directory not found`; local `go test ./... -count=1 -mod=readonly`, `go build -mod=readonly -o <temp> ./cmd/main`, `actionlint` for `.github/workflows/ci.yml`, and `git diff --check` passed.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.

- 2026-05-28: TASK-P2-017 / TS-P2-017 COMPLETED.
- User correction accepted with risk: remote plugin services expose standard `/plugin/v1/invoke` and register to host `/plugin/v1/register`, while the host explicitly configures the plugin control-plane HTTP interface address/port and reserved WS address.
- Implemented host plugin interface config (`interface.http` and `interface.ws`), optional dedicated plugin HTTP server for `/plugin/v1/register`, explicit `registration.expose_on_main_http`, config/env examples, and Blog sample host URL alignment.
- Verification: `go test ./internal/config`, `go test ./internal/app/...`, `go test ./internal/transport/http`, `go test ./pkg/plugin/...`, `go test ./...`, `go test ./...` in `remote_plugins/blog`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Strict non-goals remain: real WS/RPC transport, persistent discovery/heartbeat, production deployment, real secrets, JWT/login, database-backed IAM, and unrelated `cmd/*` rewrites.

## Current Legal Work Override

- Current Module: NONE
- Current Task ID: NONE
- Current Time Slice ID: NONE
- Current Status: PENDING_USER_CONFIRMATION
- Why this is the only legal work now: TASK-P2-023 / TS-P2-023 completed and returned the project to `NONE / NONE / PENDING_USER_CONFIRMATION`. The completed slice promoted reusable auth token signing and verification into `pkg/auth` using `github.com/golang-jwt/jwt/v5`, updated business usage to the public API, and verified targeted plus full regressions; production IAM/session/audit/migration/deployment work remains out of scope.

- 2026-05-28: TASK-P2-016 / TS-P2-016 COMPLETED.
- User goal accepted with risk: implement a host-side remote plugin registration loop, inject safe IAM principal data into the plugin hook JSON event protocol, and add `remote_plugins/blog` as an independently deployable remote Blog plugin sample.
- Implemented `POST /plugin/v1/register`, host config for registration token gating, IAM principal injection into `hooks.Event.identity`, and a standalone `remote_plugins/blog` module with registration client, invoke endpoint, README, and tests.
- Verification: `go test ./pkg/plugin/... -count=1`, `go test ./internal/config ./internal/app/... ./internal/transport/http -count=1`, `go test ./pkg/iam/... -count=1`, `go test ./... -count=1`, `go test ./... -count=1` in `remote_plugins/blog`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Non-goals: real WebSocket/RPC transport, plugin discovery service beyond explicit registration, JWT/login flow, database-backed IAM, production deployment, real secrets, and irreversible migrations.

- 2026-05-27: TASK-P2-015 / TS-P2-015 COMPLETED.
- Added maintainer comments for `cmd/server db` and created `docs/db-cli.md` with DB CLI overview, usage, extension rules, and verification guidance.
- Existing DB CLI behavior remains unchanged: database DDL, demo schema, and Todo CRUD SQL stay backed by `pkg/sqlgen`.
- Verification: `go test ./cmd/server -count=1`, `go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1`, DB docs `rg` scan, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Current legal work after completion: `NONE / NONE / PENDING_USER_CONFIRMATION`.

## 最新补充

- 2026-05-27 已完成 TASK-P2-013 / TS-P2-013：新增 `docs/configuration.md` 配置文档说明，补充配置入口、`.env` 自动加载、`RIN_APP_*` 动态前缀、`RIN_CONFIG_PATH`、`envname` 单一事实源、常用变量和新增配置字段流程。
- 已在 `README.md` 和 `docs/deployment.md` 增加配置文档入口；本次未修改 Go 实现、配置 schema、数据库 schema、真实 `.env`、部署凭据或生产配置。
- 验证结果：配置文档关键约定 `rg` 检查通过，`go test ./internal/config -count=1` 通过，`go test ./... -count=1` 通过，`git diff --check` 通过（仅 Windows LF/CRLF 提示）。
- 当前合法工作仍为 `NONE / NONE / PENDING_USER_CONFIRMATION`；项目整体仍为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，不代表 v1 release-ready。

## 项目状态

- 项目：go-scaffold
- 当前阶段：项目开发中，未达第一版发布条件
- 总体状态：IN_DEVELOPMENT_NOT_RELEASE_READY
- 最后更新：2026-05-28
- 最近 Agent：Codex
- 最近工具：Codex Desktop

## 当前合法工作

- 当前模块：NONE
- 当前任务 ID：NONE
- 当前时间切片 ID：NONE
- 当前状态：PENDING_USER_CONFIRMATION
- 为什么这是当前唯一合法状态：TASK-P2-018 / TS-P2-018 已完成，当前合法状态回到 `NONE / NONE / PENDING_USER_CONFIRMATION`。本切片已实现数据库用户、角色、权限、密码哈希、bearer token 登录、认证中间件、权限门禁 HTTP 路由、sqlgen schema bootstrap 和测试；生产级密钥/会话管理、OPA/Casbin、生产迁移、真实密钥、部署和无关重构仍不属于当前范围。

## 阶段状态

| 阶段 | 状态 | 证据 |
|---|---|---|
| 项目整体 / 第一版发布 | NOT_RELEASE_READY | 用户 2026-05-27 明确纠正：项目仍未开发完整，不应发布第一版；Docker build 通过仅为 TASK-P2-004 切片证据 |
| 项目启动 | COMPLETED | `PROJECT_BRIEF.md` 和 `docs/templates/*` 已中文化并切回项目优化主线 |
| 需求 | COMPLETED | `REQUIREMENTS.md` 已记录确认结果 |
| 高层架构 | COMPLETED | `ARCHITECTURE.md` 已记录确认边界 |
| 路线图 | COMPLETED | `ROADMAP.md` 已生成 |
| 模块边界清单 | COMPLETED | `MODULES.md` 已生成 |
| 测试矩阵与任务拆分 | COMPLETED | `TEST_MATRIX.md` 已生成，`TASKS.md` 和 `TIME_SLICES.md` 已写入 P1 草案 |
| P1 执行顺序确认 | COMPLETED | 用户再次发送“下一步”，按推荐默认顺序确认 |
| 配置 copy/update 测试与修复 | COMPLETED | `internal/config/manager_test.go` 已新增，`copyConfig` 已修复 |
| 配置环境变量策略收拢 | COMPLETED | 历史 TASK-P1-002 曾收拢数据库 `DB_*` 策略；TASK-P2-011 已进一步升级为 `RIN_APP_*` 动态前缀主策略 |
| 配置环境变量动态前缀 | COMPLETED | `internal/config` 已从 `types/constants.AppPrefix` 推导 `RIN_APP` 前缀，字段通过 `envname` 自动覆盖；TASK-P2-012 已删除 `EnvDB*` / `EnvRedis*` 等重复 env-name 常量，未加前缀变量保留兼容 fallback |
| HTTP health/ready smoke test | COMPLETED | `internal/transport/http/router_test.go` 已覆盖 `/health`、`/ready` missing/failure/ready 路径 |
| demo CRUD 测试基线 | COMPLETED | `internal/modules/demo/service/todo_test.go` 已用临时 SQLite 覆盖 Todo Create/List/Get/Update/Delete |
| demo 迁移边界收拢 | COMPLETED | `DemoMigrationPolicyFor` 固定 server-start/initdb/reload 策略，reload 不再隐式执行 demo `AutoMigrate` |
| CLI tests 命令语义收拢 | COMPLETED | `cmd/server tests` 现在执行 `go test`，并由 `cmd/server/tests_test.go` 固定命令语义 |
| pkg/* API 分类 | COMPLETED | 13 个 `pkg/*` 包已在各自 README、`ARCHITECTURE.md`、`MODULES.md` 中标注 API 定位 |
| pkg/sqlgen unsupported 边界标注 | COMPLETED | unsupported 链式查询、批量删除和 DB reverse 已显式返回 `ErrCodeUnsupportedOperation`，README 已标注部分能力边界 |
| Agent 基础设施补齐 | COMPLETED | `AGENTS.md`、`AGENT_RULES.md`、`SKILLS.md`、项目 skills、reports/specs 和跨工具目录已补齐 |
| Agent 基础设施一致性修复 | COMPLETED | TASK-INFRA-002 已补齐实际缺失的 `AGENTS.md`，规范化 skills、模板和 `.agents` 适配器 |
| Agent 状态一致性修复 | COMPLETED | TASK-INFRA-003 已生成状态诊断报告，并修复 TASK-P1-016/017 后背景文档中的旧待办表述 |
| types/* 契约边界 | COMPLETED | TASK-P1-009 已补契约说明和最小测试；2026-05-27 用户修正 `types/*` 不得聚合 `pkg/*`，已移除 `Crypto` 别名、`CacheInjectable` 和 `types/constants` 对 `pkg/executor` 的直接依赖，并补导入边界测试；同日 `AppPrefix` 已改为 `Rin`，`AppTestsCommandName` 已删除 |
| pkg/plugin 被动注册边界 | COMPLETED | TASK-P1-010 已移除 manager 主动配置加载/local factory 公共面，local/http 插件改为服务侧显式 `Register` |
| pkg/* 行为测试首批 | COMPLETED | TASK-P1-011 已补 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 最小行为测试，并修复新增测试暴露的 `pkg/yaml2go` 生成 tag/import 顺序缺陷 |
| pkg/* 行为测试第二批 | COMPLETED | TASK-P1-012 已补 `pkg/executor`、`pkg/httpserver`、`pkg/storage` 最小行为测试，并修复新增测试暴露的 `pkg/executor` 错误包装与 panic handler 缺陷 |
| pkg/cache 行为测试第三批 | COMPLETED | TASK-P1-013 已补 `pkg/cache` 隔离行为测试，使用进程内 Redis 测试服务覆盖配置、读写、批量、计数器、过期和 reload 语义 |
| pkg/utils 内部支撑测试 | COMPLETED | TASK-P1-014 已新增 `pkg/utils/utils_test.go`，覆盖 Snowflake、地址校验、端口查找、设备 ID 和 i18n helper |
| app/router/middleware 集成测试 | COMPLETED | TASK-P1-015 已新增 `internal/transport/http/router_integration_test.go`，覆盖 demo Todo HTTP CRUD、TraceID、CORS 和 Recovery 链路 |
| app 装配与 reload/config 集成测试 | COMPLETED | TASK-P1-016 已新增 `internal/app/app_integration_test.go` 和 `internal/app/reloadapp/reload_test.go`，覆盖真实 app server/initdb 装配、配置变更 hook 和 reload 分发 |
| pkg README 中文化 | COMPLETED | TASK-P1-017 已完成第一阶段 `pkg/*/README.md` 中文化，不修改 Go 代码或依赖 |
| CI 质量门禁与部署说明 | COMPLETED | TASK-P2-001 已新增 GitHub Actions CI workflow、手动部署说明和 README 入口，不执行真实部署或使用密钥 |
| 真实 CD 范围确认 | COMPLETED | 用户选择 C、确认使用远程部署，并进一步确认用 `.env` 风格文件配置；TASK-P2-002 已新增远程部署变量模板 |
| 显式参数部署入口 | COMPLETED | `deploy.sh` 和 `script/install.sh` 已新增，旧本地部署 env 文件已删除，部署说明已同步 |
| 远程部署 workflow | COMPLETED | TASK-P2-003 已新增手动 staging workflow、Secrets 配置说明和远程主机前置条件；未执行真实部署 |
| Linux Docker production 部署制品 | COMPLETED | TASK-P2-004 / TS-P2-004 已实现 Dockerfile、production Compose 示例、手动 production workflow 闸门和统一 `deploy.sh` 部署入口；用户已在 Linux Docker 环境用 `GOPROXY=https://goproxy.cn,direct` 构建通过 |
| 插件钩子运行时与 IAM 公共接口 | COMPLETED | TASK-P2-005 至 TASK-P2-010 已完成；`go test ./pkg/plugin/... -count=1`、`go test ./pkg/iam/... -count=1`、`go test ./internal/config ./internal/app/... -count=1`、`go test ./... -count=1`、server build 和 `git diff --check` 均通过 |
| 部署实现 | COMPLETED | TASK-P2-004 已补齐 Dockerfile、production Compose 示例、production 配置样例、统一 `deploy.sh` 部署入口和手动 production workflow 闸门 |
| 部署验证 | COMPLETED | 脚本 Bash 语法解析、YAML 解析、actionlint、`go test ./... -count=1`、server build、`git diff --check` 和用户 Linux Docker build 均已通过 |
| 交接 | COMPLETED | `AGENT_HANDOFF.md` 已更新到当前无自动下一实现任务，并保留 TASK-P2-004 至 TASK-P2-010 完成验证记录 |
| Phase 6 收尾 | COMPLETED | 用户选择 A 后已完成 TASK-PHASE6-001；最终回归和交接文档已更新 |

## 当前关键发现

| ID | 发现 | 来源 | 状态 |
|---|---|---|---|
| FIND-001 | P1 关键测试缺口已持续收敛 | `go test ./... -count=1`、TASK-P1-003 至 TASK-P1-016 | [CONFIRMED] app/router/demo/config/reload 与主要 `pkg/*` 路径已补最小测试 |
| FIND-002 | `.env.example` 与数据库环境变量前缀不一致 | `MODULES.md` BC-001；TASK-P1-002 已修复 | [CONFIRMED] 已处理 |
| FIND-003 | `manager.copyConfig` 未完整复制配置字段 | `MODULES.md` BC-002；TASK-P1-001 已修复 | [CONFIRMED] 已处理 |
| FIND-004 | demo schema 自动迁移触发点需收拢 | `MODULES.md` BC-003；TASK-P1-005 已固定 server-start/initdb/reload 策略 | [CONFIRMED] 已处理 |
| FIND-005 | `cmd/server tests` 命令语义与行为不一致 | `MODULES.md` BC-004；TASK-P1-006 已改为真实 Go test 入口 | [CONFIRMED] 已处理 |
| FIND-010 | `pkg/*` 公共/内部定位未逐包标记 | `ARCHITECTURE.md`、`MODULES.md`；TASK-P1-007 已完成分类 | [CONFIRMED] 已处理 |
| FIND-011 | `pkg/sqlgen` TODO/unsupported 边界不清 | `pkg/sqlgen` README 和源码；TASK-P1-008 已显式返回 unsupported 或文档化 partial 能力 | [CONFIRMED] 已处理 |
| FIND-012 | `types/result`、错误码和跨层类型边界待明确 | TASK-P1-009 已补 `docs/specs/types_contract_boundary.md`、package doc 和最小测试；用户修正 `types` 只能存在应用层以上的层级再定义，不得直接聚合 `pkg/*` 基础设施接口 | [CONFIRMED] 已处理 |
| FIND-013 | `pkg/plugin` 主动注册服务边界需收拢 | 用户修正；TASK-P1-010 已改为被动 registry/runtime | [CONFIRMED] 已处理 |
| FIND-014 | 背景文档保留 TASK-P1-016 前旧状态 | `ARCHITECTURE.md`、`MODULES.md`、`PROJECT_BRIEF.md`、`ROADMAP.md`；TASK-INFRA-003 已修复 | [CONFIRMED] 已处理 |
| FIND-015 | CI/CD 与部署缺少首个安全边界 | `REQ-OPT-P2-003`、`BL-007`、`BL-008`；用户选择 D | [CONFIRMED] TASK-P2-001 已处理非生产 CI 门禁和部署说明 |
| FIND-016 | 真实 CD 自动化缺少环境与密钥决策 | `BL-024`；用户选择 C、确认远程部署，并确认使用 `.env` 风格配置和实现 workflow | [CONFIRMED] 手动 staging 远程部署 workflow 已补；镜像发布、production 和真实运行仍需单独确认 |
| FIND-017 | production Docker 部署缺少可提交制品 | 用户要求“linux、docker、production -> 部署”；用户修正“环境变量在部署脚本上动态配置”；`BL-024` 剩余范围 | [CONFIRMED] 制品和统一 `deploy.sh` 入口已补；用户已在 Linux Docker 环境完成镜像构建验证，真实 production 运行仍不在当前会话执行 |
| FIND-018 | 当前仓库未达第一版发布条件 | 用户纠正；`README.md` 当前仍说明仅保留基础设施启动链路和 demo CRUD，auth/rbac、生产迁移、镜像发布、真实 production 运行和完整产品验收均未完成 | [CONFIRMED] 不发布第一版；后续需先确认发布验收清单和剩余开发范围 |
| FIND-019 | 配置环境变量前缀曾固定或分散 | 用户修正；`internal/config` 旧实现硬编码 `REI_APP` 且按模块手写覆盖 | [CONFIRMED] 已改为 `AppPrefix` 动态前缀 + `envname` 反射覆盖，`.env` 自动加载路径已测试 |
| FIND-020 | 配置 env-name 常量与 `envname` 标签重复 | 用户修正；`EnvDB*`、`EnvRedis*`、`EnvServer*` 等常量在 `envname` 标签落地后成为第二事实源 | [CONFIRMED] 已删除重复常量；测试 helper 从配置结构体标签读取环境变量名 |
| FIND-006 | P1 执行顺序尚未确认 | `TEST_MATRIX.md`、`RISK_REGISTER.md` RISK-009；用户再次发送“下一步” | [CONFIRMED] 已确认 |
| FIND-007 | `AGENTS.md` 被状态文件声明已补齐但实际缺失 | `Test-Path AGENTS.md`、`docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` | [CONFIRMED] 已修复 |
| FIND-008 | `/health`、`/ready` 路由缺少 smoke test | `TEST_MATRIX.md` TM-P0-003；TASK-P1-003 已补测试 | [CONFIRMED] 已处理 |
| FIND-009 | demo Todo CRUD 缺少测试基线 | `TEST_MATRIX.md` TM-P0-005；TASK-P1-004 已补 service/repository 隔离测试 | [CONFIRMED] 已处理 |

## 待用户确认

| ID | 问题 | 影响 | 选项 | Required By |
|---|---|---|---|
| CONFIRM-NEXT-001 | 选择 P1 后续范围或进入收尾 | 已确认：用户选择 A | A: 提升 `BL-021` / `TM-P1-005` 做 `types/*` 契约边界 | COMPLETED |
| CONFIRM-NEXT-002 | 选择 `types/*` 契约边界完成后的后续范围 | 已确认：用户修正并选择收拢 `pkg/plugin` 被动注册边界 | 提升 `BL-022` / `TM-P1-006` | COMPLETED |
| CONFIRM-NEXT-003 | 选择 `pkg/plugin` 被动注册边界完成后的后续范围 | 已确认：用户选择 A，提升 `BL-020` 补 `pkg/*` 行为测试 | A: 提升 `BL-020` 补 `pkg/*` 行为测试 | COMPLETED |
| CONFIRM-NEXT-004 | 选择首批 `pkg/*` 行为测试完成后的后续范围 | 已确认：用户发送“下一步”，按选项 A 继续下一批 `pkg/*` 行为测试 | A: 继续下一批 `pkg/*` 行为测试；B: 进入 Phase 6 收尾；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-005 | 选择第二批 `pkg/*` 行为测试完成后的后续范围 | 已确认：用户选择 A，继续 `BL-020` 剩余包，第三批限定 `pkg/cache` | A: 继续剩余 `pkg/*` 行为测试；B: 进入 Phase 6 收尾；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-006 | 选择 `pkg/cache` 行为测试完成后的后续范围 | 已确认：用户选择 B，提升 `pkg/utils` 内部支撑测试 | A: 进入 Phase 6 收尾；B: 提升内部支撑测试；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-007 | 选择 `pkg/utils` 内部支撑测试完成后的后续范围 | 已确认：用户回复 `b`，选择 B | A: 进入 Phase 6 收尾；B: 提升 app/router/middleware 等集成测试；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-008 | 选择 app/router/middleware 集成测试完成后的后续范围 | 已确认：用户选择 A，进入 Phase 6 收尾 | A: 进入 Phase 6 收尾；B: 继续 app 装配/reload/config 等剩余集成测试；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-009 | 选择 TASK-INFRA-003 后的后续方向 | 已确认：用户选择 D，进入 CI/CD 与部署方向首切片 | D: CI/CD 与部署；首切片限定 CI 质量门禁与部署说明 | COMPLETED |
| CONFIRM-NEXT-010 | 确认真实 CD / 镜像发布 / 远程部署自动化边界 | 已确认：用户选择 C、使用远程部署、通过 `.env` 风格模板配置，并明确确认实现远程部署 workflow | TASK-P2-003 已完成 | COMPLETED |
| CONFIRM-NEXT-011 | 确认 Linux/Docker/production 部署制品 | 已确认：用户要求“开始，linux、docker、production -> 部署” | TASK-P2-004 已完成 Docker 构建验证 | COMPLETED |
| CONFIRM-NEXT-012 | 确认下一阶段开发范围或第一版发布验收清单 | 当前项目未达发布条件，不能把已完成切片直接当作 v1 | 需由用户明确选择下一阶段：补产品功能、完善 auth/rbac、生产迁移、镜像发布/真实运行、发布验收清单或其他范围 | PENDING_USER_CONFIRMATION |

## 待验证

| ID | 任务 | 需要验证内容 | 命令/方法 |
|---|---|---|---|
| VERIFY-P2-004 | TASK-P2-004 | Dockerfile 镜像构建 | [CONFIRMED] 用户已在 Linux Docker 环境运行 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .`，BuildKit 输出 `23/23 FINISHED`，镜像标记为 `docker.io/library/go-scaffold:local` |
| VERIFY-P2-005 | TASK-P2-005 至 TASK-P2-010 | 插件钩子运行时、远程插件传输、IAM 公共接口和 app 装配 | [CONFIRMED] `go test ./pkg/plugin/... -count=1`；`go test ./pkg/iam/... -count=1`；`go test ./internal/config ./internal/app/... -count=1`；`go test ./... -count=1`；`go build -o <temp> ./cmd/server`；`git diff --check` |
| VERIFY-P2-011 | TASK-P2-011 | 动态配置环境变量前缀、`envname` 覆盖、`.env` 自动加载和相关 app/cmd 集成 | [CONFIRMED] `go test ./internal/config -count=1`；`go test ./cmd/server ./internal/app/... -count=1`；`go test ./... -count=1`；`git diff --check` |
| VERIFY-P2-012 | TASK-P2-012 | 删除重复 env-name 常量，确认 `envname` 标签是字段环境变量名唯一来源 | [CONFIRMED] `rg -n "Env(DB|Redis|Server|Log|I18n|CORS|InitDB|Executor|Storage|Plugin|IAM)" internal cmd types deploy docs .env.example Dockerfile -S` 无匹配；`go test ./internal/config -count=1`；`go test ./cmd/server ./internal/app/... -count=1`；`go test ./... -count=1`；`git diff --check` |

## 需要返工

| ID | 任务 | 原因 | 下一步 |
|---|---|---|---|
|  |  |  |  |

## 最近执行

- 摘要：接受用户修正：`internal/config` 不再保留 `EnvDB*`、`EnvRedis*`、`EnvServer*` 等重复 env-name 常量；字段环境变量名统一由配置结构体 `envname` 标签声明。
- 变更文件：`internal/config/constants.go`、`internal/config/manager_test.go` 和项目状态文档。
- 执行命令：必读文件读取；用户纠正审查；`gofmt -w internal/config/constants.go internal/config/manager_test.go`；`rg -n "Env(DB|Redis|Server|Log|I18n|CORS|InitDB|Executor|Storage|Plugin|IAM)" internal cmd types deploy docs .env.example Dockerfile -S`；`go test ./internal/config -count=1`；`go test ./cmd/server ./internal/app/... -count=1`；`go test ./... -count=1`；`git diff --check`。
- 测试结果：重复常量引用扫描无匹配；目标包测试、cmd/app 相关测试和全量回归均通过。
- 完成判断：TASK-P2-012 / TS-P2-012 完成；TASK-P2-011 的动态前缀实现保持完成；项目整体仍为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，当前等待用户确认下一阶段范围或第一版发布验收清单。

## 下一步

- 合法下一步：当前无自动下一实现任务；项目未达发布条件。如需继续，必须由用户确认新的开发范围或第一版发布验收清单，并拆分任务/时间切片。
- 可选后续方向：镜像发布流水线、真实 staging/production 运行、生产迁移框架、完整 auth/rbac、插件发现或 RPC/WS 传输。
- 非目标保持：JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS 传输、生产部署、镜像发布和密钥管理仍不属于本轮完成范围。
## Current Update: TASK-P2-015 / TS-P2-015

- Date: 2026-05-27
- Status: COMPLETED
- Current legal work after this slice: `NONE / NONE / PENDING_USER_CONFIRMATION`
- User goal: add comments for `cmd/db` and generate overview, usage, and extension documentation.
- Result: `cmd/server/db.go` now includes concise maintainer comments; `docs/db-cli.md` documents DB CLI overview, quick usage, operation semantics, flags, layering, extension workflow, forbidden regressions, and verification guidance; `docs/configuration.md` and `docs/deployment.md` link to it.
- Verification: `go test ./cmd/server -count=1` PASS; `go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1` PASS; DB docs `rg` scan PASS; `git diff --check` PASS with only Git LF/CRLF notices.
- No DB behavior, schema, production migration, old `initdb`, InitDB config, SQL script, or `AutoMigrate` path was changed.

## Current Update: TASK-P2-014 / TS-P2-014

- Date: 2026-05-27
- Status: COMPLETED
- Current legal work after this slice: `NONE / NONE / PENDING_USER_CONFIRMATION`
- User goal: remove current init/migration commands and reimplement DB operations through the sqlgen toolchain; table/database creation and CRUD data operations must use tool-chain style programming, not hand-written scripts.
- Result: `cmd/server db` is now the explicit DB CLI. It can generate database DDL, print/apply sqlgen-generated demo schema, and run Todo CRUD operations through sqlgen-generated SQL. `initdb`, InitDB config, SQL bootstrap scripts, demo migration wrappers, and GORM `AutoMigrate` were removed from current code/config paths.
- Verification: code/config removed-path scan passed; targeted package tests passed; `go test ./... -count=1` passed; `git diff --check` passed with only Git LF/CRLF notices.
