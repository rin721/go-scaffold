# CHANGELOG.md

## Latest Addendum

### 2026-05-28 - TASK-P2-023 - pkg/auth public token API

- Change: Added `pkg/auth` as the public authentication token API with JWT issue/verify support backed by `github.com/golang-jwt/jwt/v5`.
- Change: Removed hand-written token signing/parsing from the main-service user module; business code now maps users into `auth.Claims`.
- Change: Updated app composition and focused tests so configured `auth.token_secret` and `auth.token_ttl` create the public auth token service.
- Scope: No refresh/session revocation, audit, password reset, external IAM replacement, production secret manager, production migration, deployment, real secrets/users, or plugin transport implementation.
- Verification: `go test ./pkg/auth -count=1`; `go test ./internal/modules/user/... -count=1`; `go test ./internal/app/initapp -count=1`; `go test ./internal/app/... -count=1`; `go test ./internal/transport/http -count=1`; `go test ./... -count=1`; `git diff --check` passed.
- Status: TASK-P2-023 / TS-P2-023 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-P2-022 - pkg/rbac public RBAC API

- Change: Promoted the Casbin-backed RBAC authorizer contracts and implementation from `internal/modules/user/rbac` to the public infrastructure package `pkg/rbac`.
- Change: Added `pkg/rbac` docs, README, and focused Casbin authorization tests.
- Change: Updated main-service user authorization and app composition to use `github.com/rei0721/go-scaffold/pkg/rbac`.
- Scope: No external IAM replacement, refresh/session/audit/password-reset work, production migration, deployment, real secrets/users, or plugin transport implementation.
- Verification: `go test ./pkg/rbac -count=1`; `go test ./internal/config -count=1`; `go test ./internal/modules/user/... -count=1`; `go test ./internal/app/initapp -count=1`; `go test ./internal/app/... -count=1`; `go test ./internal/transport/http -count=1`; `go test ./... -count=1`; `git diff --check` passed.
- Status: TASK-P2-022 / TS-P2-022 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-P2-021 - Casbin-backed RBAC authorization wrapper

- Change: Added `github.com/casbin/casbin/v2` and a local Casbin authorizer wrapper under `internal/modules/user/rbac`.
- Change: Added `configs/rbac_model.conf` and `rbac.model_path` config/env support.
- Change: Replaced business authorization's hand-written permission matcher with Casbin enforcement over DB-backed role-permission policy data.
- Scope: No external IAM replacement, refresh/session/audit/password-reset work, production migration, deployment, real secrets/users, or plugin transport implementation.
- Verification: `go test ./internal/config -count=1`; `go test ./internal/modules/user/... -count=1`; `go test ./internal/app/initapp -count=1`; `go test ./internal/app/... -count=1`; `go test ./internal/transport/http -count=1`; `go test ./... -count=1`; `git diff --check` passed.
- Status: TASK-P2-021 / TS-P2-021 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-P2-020 - RBAC seed config under configs

- Change: Added `rbac` config support for roles, permissions, and role-permission grants, including validation and config copy coverage.
- Change: Added safe RBAC seed entries to `configs/config.example.yaml` and `configs/config.yaml`.
- Change: Added idempotent startup RBAC seed application through the main-service user module when `rbac.enabled` and `rbac.apply_on_start` are true.
- Scope: No seeded real users/passwords, OPA/Casbin, external IAM replacement, refresh/session/audit/password-reset work, production migration, deployment, or plugin transport implementation.
- Verification: `go test ./internal/config -count=1`; `go test ./internal/modules/user/... -count=1`; `go test ./internal/app/initapp -count=1`; `go test ./internal/app/... -count=1`; `go test ./... -count=1`; `git diff --check` passed.
- Status: TASK-P2-020 / TS-P2-020 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-P2-019 - Auth token config hardening

- Change: Added `auth.token_secret` and `auth.token_ttl` config/env support and wired them into the existing main-service user token service.
- Change: Added placeholder examples and tests for validation, env override, and token wiring.
- Scope: Placeholder examples and focused tests only; no real secrets, refresh-token/session revocation, audit logging, password reset, external IAM, OPA/Casbin, production migrations, deployment, or plugin WS/RPC/discovery implementation.
- Verification: `go test ./internal/config -count=1`; `go test ./internal/app/initapp -count=1`; `go test ./internal/app/... -count=1`; `go test ./... -count=1`; `git diff --check` passed.
- Status: TASK-P2-019 / TS-P2-019 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Latest Addendum

### 2026-05-28 - TASK-P2-018 - Main-service user/auth/RBAC

- Change: Added `internal/modules/user` with users, roles, permissions, user-role and role-permission assignments, password hashing, HMAC bearer tokens, and permission checks.
- Change: Added `/api/v1/auth/register`, `/api/v1/auth/login`, `/api/v1/auth/me`, `/api/v1/users`, `/api/v1/roles`, and `/api/v1/permissions` routes.
- Change: Added sqlgen bootstrap for user/RBAC tables and app/router/service/schema regression tests.
- Scope: No production secret/session management, refresh-token/session revocation, OPA/Casbin, audit logging, password reset, deployment, plugin transport changes, or production migration framework.
- Verification: `go test ./internal/modules/user/... -count=1`; `go test ./internal/app/dbapp -count=1`; `go test ./internal/transport/http -count=1`; `go test ./internal/app/... -count=1`; `go test ./... -count=1`; `git diff --check` passed.
- Status: TASK-P2-018 / TS-P2-018 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-INFRA-004 - CI build target repair

- Change: Updated `.github/workflows/ci.yml` so the `Build server` step builds the current `./cmd/main` entrypoint instead of removed `./cmd/server`.
- Root cause: GitHub Actions run `26531295923` passed tests but failed with `stat .../cmd/server: directory not found`.
- Scope: CI workflow and project status records only; no Go behavior, deployment, secrets, workflow trigger, or image publishing.
- Verification: GitHub job log inspection; `go test ./... -count=1 -mod=readonly`; `go build -mod=readonly -o <temp> ./cmd/main`; `actionlint`; `git diff --check` passed.
- Status: TASK-INFRA-004 / TS-INFRA-004 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-P2-017 - Plugin control-plane interface configuration

- Change: Added host plugin control-plane config for `plugin.interface.http.enabled/host/port/public_url` and reserved `plugin.interface.ws.public_url`.
- Change: Added optional dedicated plugin HTTP server startup for `/plugin/v1/register`, while main HTTP registration exposure now requires explicit `plugin.registration.expose_on_main_http`.
- Change: Updated config/env examples and validation so registration must be explicitly exposed either on the plugin interface or on main HTTP.
- Change: Updated `remote_plugins/blog` to use `BLOG_PLUGIN_HOST_HTTP_URL` / `BLOG_PLUGIN_HOST_WS_URL` with fallback to the older `BLOG_PLUGIN_MAIN_*` names, while keeping the standard `/plugin/v1/invoke` service.
- Scope: No real WS/RPC adapter, heartbeat/discovery, production deployment, real secrets, JWT/login, or database-backed IAM.
- Verification: `go test ./internal/config`; `go test ./internal/app/...`; `go test ./internal/transport/http`; `go test ./pkg/plugin/...`; `go test ./...`; `go test ./...` inside `remote_plugins/blog`; `git diff --check` passed with only Git LF/CRLF notices.
- Status: TASK-P2-017 / TS-P2-017 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Previous Addendum

### 2026-05-28 - TASK-P2-016 - Remote plugin registration and Blog sample

- Change: Added host-side remote plugin registration protocol and `POST /plugin/v1/register` HTTP handler with optional shared registration token validation.
- Change: Hook JSON events now support safe `identity.principal` data, injected by the host through a plugin manager event enricher without making `pkg/plugin` import `pkg/iam`.
- Change: Wired plugin registration config into `internal/config`, `internal/app/initapp`, and `internal/transport/http`; updated `.env.example` and `configs/config.example.yaml`.
- Change: Added `remote_plugins/blog` as an independent Go module with config loading, standard `/plugin/v1/invoke` service, startup host registration client, Blog operations, hook handling, README, and tests.
- Scope: No real WS/RPC adapter, automatic discovery daemon, JWT/login flow, database-backed IAM, production deployment, real secrets, or irreversible migrations.
- Verification: `go test ./pkg/plugin/... -count=1`; `go test ./internal/config ./internal/app/... ./internal/transport/http -count=1`; `go test ./pkg/iam/... -count=1`; `go test ./... -count=1`; `go test ./... -count=1` inside `remote_plugins/blog`; `git diff --check` passed with only Git LF/CRLF notices.
- Status: TASK-P2-016 / TS-P2-016 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Latest Addendum

### 2026-05-27 - TASK-P2-015 - DB CLI comments and documentation

- Change: Added concise maintainer comments to `cmd/server/db.go` describing command ownership, parsed options, side-effect-free DDL preview, and printable SQL routing.
- Change: Added `docs/db-cli.md` with DB CLI overview, quick usage, operation table, flags, layering, extension workflow, forbidden regressions, and verification guidance.
- Change: Linked the new DB CLI guide from `docs/configuration.md` and `docs/deployment.md`.
- Scope: Documentation and comments only; no DB behavior, schema, config, production migration, or old `initdb` path changes.
- Verification: `go test ./cmd/server -count=1`, `go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1`, DB docs `rg` scan, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Status: TASK-P2-015 / TS-P2-015 completed; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`, legal next state is `NONE / NONE / PENDING_USER_CONFIRMATION`.

## 最新补充

### 2026-05-27 - TASK-P2-013 - Config documentation

- 变更：新增 `docs/configuration.md`，说明配置入口、加载顺序、`.env` 自动加载、`RIN_APP_*` 动态前缀、`RIN_CONFIG_PATH`、`envname` 单一事实源、常用变量和新增配置字段流程。
- 变更：`README.md` 和 `docs/deployment.md` 增加配置文档入口。
- 范围：仅文档和状态记录；未修改 Go 实现、配置 schema、数据库 schema、真实 `.env`、密钥、部署凭据或生产配置。
- 验证：配置文档关键文本 `rg` 检索通过；`go test ./internal/config -count=1`、`go test ./... -count=1` 和 `git diff --check` 通过。
- 状态：TASK-P2-013 / TS-P2-013 完成；项目整体仍为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，当前合法任务回到 `NONE / NONE / PENDING_USER_CONFIRMATION`。

## 最新变更

### 2026-05-27 - TASK-P2-012 - Config envname constants cleanup

- 变更：接受用户修正，删除 `internal/config/constants.go` 中 `EnvDB*`、`EnvRedis*`、`EnvServer*`、`EnvLog*`、`EnvI18n*`、`EnvCORS*`、`EnvInitDB*`、`EnvExecutor*`、`EnvStorage*`、`EnvPlugin*`、`EnvIAM*` 等重复 env-name 常量。
- 变更：`internal/config/constants.go` 仅保留动态前缀 helper、`.env` 文件名、分隔符和配置段名常量；字段环境变量名由配置结构体 `envname` 标签单一维护。
- 变更：`internal/config/manager_test.go` 新增标签读取 helper，测试从 `envname` 标签生成环境变量名，避免测试继续依赖第二套常量表。
- 验证：env 常量 `rg` 扫描无匹配；`go test ./internal/config -count=1`、`go test ./cmd/server ./internal/app/... -count=1`、`go test ./... -count=1`、`git diff --check` 通过。
- 状态：TASK-P2-012 / TS-P2-012 完成；项目整体仍为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，当前合法任务回到 `NONE / NONE / PENDING_USER_CONFIRMATION`。

### 2026-05-27 - TASK-P2-011 - Dynamic config environment prefix

- 变更：接受用户修正，`internal/config` 环境变量前缀不再固定，改为从 `types/constants.AppPrefix` 动态派生；当前 `AppPrefix=Rin`，配置覆盖主前缀为 `RIN_APP`，配置路径变量为 `RIN_CONFIG_PATH`。
- 变更：新增 `envname` 标签驱动的统一反射覆盖逻辑，覆盖 Database、Redis、Server、Logger、I18n、InitDB、Executor、Storage、Plugin、IAM 和 CORS 可配置字段；动态前缀变量优先，未加前缀变量作为兼容 fallback。
- 变更：`Manager.Load` 和配置热重载路径均执行 `.env` 加载与环境变量覆盖；`cmd/server` 配置路径 flag 改用动态环境变量名。
- 变更：同步 `.env.example`、Dockerfile、production Compose 示例、`deploy.sh` 和部署说明到 `RIN_APP_*` / `RIN_CONFIG_PATH`。
- 验证：`go test ./internal/config -count=1`、`go test ./cmd/server ./internal/app/... -count=1`、`go test ./... -count=1`、`git diff --check` 通过。
- 状态：TASK-P2-011 / TS-P2-011 完成；项目整体仍为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，当前合法任务回到 `NONE / NONE / PENDING_USER_CONFIRMATION`。

### 2026-05-27 - User correction - types app constants

- 变更：接受用户修正，将 `types/constants.AppPrefix` 从 `Rei` 改为 `Rin`。
- 变更：删除 `types/constants.AppTestsCommandName`，`cmd/server` 的 `tests` 命令名改为命令包内私有常量，不再由 `types` 提供。
- 变更：新增 `types/constants/app_test.go` 固定应用常量，更新 `cmd/server/tests_test.go` 保持 CLI tests 命令语义。
- 验证：`go test ./types/... ./cmd/server -count=1` 与 `go test ./... -count=1` 通过。
- 状态：本次 `types` 常量修正完成；当前合法任务仍为 `NONE / NONE / PENDING_USER_CONFIRMATION`。

### 2026-05-27 - User correction - root types layering

- 变更：接受用户纠正，`types/*` 不得直接聚合 `pkg/*` 基础设施接口；根 `types` 不得为 `pkg/crypto.Crypto` 提供别名，也不得定义依赖 `pkg/cache.Cache` 的 `CacheInjectable`。
- 变更：删除 `types/interfaces.go`，更新 `types/doc.go`，将 `types/constants` executor pool 名称改为字符串常量，新增 `types/import_boundary_test.go` 固定 `types/*` 不导入 `pkg/*`。
- 变更：更新 `pkg/executor` 文档示例，不再建议下层 `pkg` 文档反向指定 `types/constants` 作为 typed constants 位置。
- 变更：同步 `ARCHITECTURE.md`、`MODULES.md`、`REQUIREMENTS.md`、`ACCEPTANCE.md`、`TEST_MATRIX.md`、`docs/specs/types_contract_boundary.md` 和项目状态文档，明确 `types` 只能承载应用层以上确认过的跨层契约。
- 验证：`go test ./types/... -count=1` 与 `go test ./... -count=1` 通过；`git diff --check` 仅有 Windows LF/CRLF 提示。
- 状态：`types` 分层修正完成；当前合法任务仍为 `NONE / NONE / PENDING_USER_CONFIRMATION`。

### 2026-05-27 - User correction - project not release-ready

- 变更：接受用户纠正，当前项目仍未开发完整，不应发布第一版。
- 变更：将项目整体状态从易误解的 `COMPLETED` 调整为 `IN_DEVELOPMENT_NOT_RELEASE_READY`；TASK-P2-004 至 TASK-P2-010 的切片完成证据保留。
- 变更：新增 `DEC-027`、`RISK-022`、`BL-027`、`TM-P2-012` 和 `ISSUE-STATUS-004`，明确 Docker build 通过、部署制品和基础设施切片完成不等于 v1 release-ready。
- 范围：仅更新文档和状态；未修改 Go 代码、配置 schema、部署脚本或 workflow；未触发真实部署、未推送镜像、未执行生产迁移。
- 状态：当前合法任务为 `NONE / NONE / PENDING_USER_CONFIRMATION`；后续需用户确认新的开发范围或第一版发布验收清单。

### 2026-05-27 - TASK-P2-004 - Docker build verification completed

- 变更：记录用户在 Linux Docker 环境补跑 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` 成功，解除 TASK-P2-004 / TS-P2-004 的唯一阻塞项。
- 验证：Docker BuildKit 输出 `23/23 FINISHED`，镜像写入 `sha256:4df5520bcf1c45a922be8db2e6c5e58ae8fc025f34bea5f1d4bf33f0b2301785`，并标记为 `docker.io/library/go-scaffold:local`。
- 范围：本轮仅更新状态、验收、测试报告、问题记录、变更记录和交接说明；未修改 Go 代码，未触发 workflow，未连接远程服务器，未推送镜像，未执行真实 production。
- 状态：TASK-P2-004 / TS-P2-004 转为 `COMPLETED`，`ISSUE-P2-005` 关闭；当前无自动下一实现任务。

### 2026-05-27 - TASK-P2-004 - Docker build proxy args

- 变更：诊断用户远端 Docker build 慢/超时问题，确认旧 `Dockerfile` 未声明 `GOPROXY` build arg，导致 `--build-arg GOPROXY=...` 未生效。
- 变更：`Dockerfile` 新增 `GOPROXY` / `GOSUMDB` build arg，并为 `go mod download`、`go build` 增加 BuildKit cache mount。
- 变更：`docs/deployment.md` 增加带 `GOPROXY` 的 Docker 构建示例，项目状态文档记录远端 Go 代理超时阻塞。
- 验证：本机仍无 Docker CLI，未在本机执行 Docker build；本轮未修改 Go 代码，未运行 Go 测试；`git diff --check` PASS，仅有 Windows LF/CRLF 提示。
- 状态：TASK-P2-004 / TS-P2-004 保持 `BLOCKED`，`ISSUE-P2-005` 保持 OPEN；待 Docker 环境用更新后的 Dockerfile 重跑。

### 2026-05-27 - TASK-P2-004 - TS-P2-004 blocked recheck

- 变更：用户发送“下一步”后，按当前唯一合法任务复验 Docker build 前置环境，不推进新功能。
- 验证：`docker version` 失败；`docker`、`podman`、`nerdctl`、`docker.exe` 均不可用；`docker build -t go-scaffold:local .` 因前置 Docker CLI/daemon 缺失未执行。
- 范围：本轮仅更新项目状态与交接文档，未修改 Go 代码、部署脚本、workflow、真实配置或密钥。
- 状态：TASK-P2-004 / TS-P2-004 保持 `BLOCKED`，`ISSUE-P2-005` 保持 OPEN；TASK-P2-005 至 TASK-P2-010 插件/IAM 主线仍保持 `COMPLETED`。

### 2026-05-27 - dev.tmp/new-plugin completion audit

- 变更：按 `dev.tmp/new-plugin.md` 重新审计插件钩子运行时、HTTP 远程插件传输、IAM 公共接口、配置/app/reload/lifecycle 接入。
- 变更：`pkg/plugin/hooks` 现在会在注册时拒绝 nil `HandlerFunc`，直接调用 nil `HandlerFunc` 也返回 `ErrNilHandler`。
- 变更：HTTP 远程插件响应超过 `maxResponseBytes` 时返回包装 `ErrInvalidResponse` 的明确错误，不再依赖截断后的 JSON 解码失败。
- 变更：新增测试覆盖 nil hook handler 拒绝、HTTP 响应大小限制、`after_invoke` hook 失败时返回插件响应和包装后的 hook 错误。
- 验证：`go test ./pkg/plugin/... -count=1`、`go test ./pkg/iam/... -count=1`、`go test ./internal/config ./internal/app/... -count=1`、`go test ./... -count=1`、server build、`git diff --check` 均通过。
- 状态：`dev.tmp/new-plugin.md` 设计完成；TASK-P2-004 Docker build 阻塞仍独立保持打开。

### 2026-05-27 - TASK-P2-004 - TS-P2-004 blocked verification

- 变更：用户发送“下一步”后，按协议处理剩余 Docker build 验证项，不推进新功能。
- 验证：`docker version` 失败；`docker`、`podman`、`nerdctl`、`docker.exe` 均不可用；`docker build -t go-scaffold:local .` 因前置 Docker CLI/daemon 缺失未执行。
- 状态：TASK-P2-004 / TS-P2-004 记录为 `BLOCKED`，`ISSUE-P2-005` 保持 OPEN；TASK-P2-005 至 TASK-P2-010 插件/IAM 主线仍保持 `COMPLETED`。

### 2026-05-27 - TASK-P2-005 至 TASK-P2-010 - TS-P2-005 至 TS-P2-010

- 变更：按 `dev.tmp/new-plugin.md` 设计完成插件钩子运行时、HTTP 远程插件传输、独立 IAM 公共接口和 app 组合层接入；原 `dev.tmp/new-pllugin.md` 视为笔误。
- 变更：新增 `pkg/plugin/hooks`，提供钩子点、事件、结果、处理器、注册表和服务查找能力。
- 变更：扩展 `pkg/plugin.Manager`，新增 `Hooks()`、`RegisterHook`、`WithHooks` 和注册/调用/错误/关闭/配置/日志/IAM 等标准钩子点，保持被动注册模型。
- 变更：新增 `NewHTTPServer`、`hooks.execute` 标准操作和 `RemoteHook`，沿用 JSON `Request` / `Response` 协议。
- 变更：新增 `pkg/iam` 公共类型和 `pkg/iam/memory` 实现，覆盖 token 凭证、策略授权、拒绝优先、通配、过期和默认拒绝。
- 变更：新增 `plugin` 与 `iam` 配置字段，默认 disabled；`internal/app/initapp.Infrastructure` 增加 IAM 和 Plugins，reload/lifecycle 已接入。
- 范围：未实现 JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS 传输、生产部署、镜像发布或密钥管理；TASK-P2-004 Docker build 阻塞项保持打开。
- 验证：`go test ./pkg/plugin/... -count=1`、`go test ./pkg/iam/... -count=1`、`go test ./internal/config ./internal/app/... -count=1`、`go test ./... -count=1`、server build、`git diff --check` 均通过。

### 2026-05-27 - TASK-P2-004 - TS-P2-004 rework

- 变更：删除已跟踪的旧部署 env 示例和旧远程 Linux 动态 env 脚本；本地旧部署 env 文件已删除且未读取内容。
- 变更：新增根 `deploy.sh` 和 `script/install.sh`，支持 clone 后执行和 direct curl 执行两种流程，部署与应用配置均通过显式参数传入。
- 变更：重构 `.github/workflows/deploy-remote.yml`，改为 GitHub Variables/Secrets 组装显式参数，通过 SSH 在远端执行 `script/install.sh` / `deploy.sh`。
- 变更：Compose production 示例改为读取显式导出的 DB、Redis、Server、Logger、I18n、Storage、CORS 环境变量，不再依赖旧部署 env 文件。
- 验证：shfmt Bash parser、workflow YAML 解析、actionlint、旧引用 `rg` 检查、`go test ./... -count=1`、server build、`git diff --check` 均通过；Docker CLI 不存在，`docker build -t go-scaffold:local .` 保持 `PENDING_VERIFICATION`。

### 2026-05-26 - TASK-P2-004 - TS-P2-004

- 变更：用户要求“开始，linux、docker、production -> 部署”，并修正“环境变量在部署脚本上动态配置”，接受为 production Docker 远程部署制品、远程 Linux 统一 `deploy.sh` 入口和手动闸门切片。
- 变更：新增 `Dockerfile`，构建 Linux server 镜像并以非 root 用户运行。
- 变更：新增 `.dockerignore`，避免 Git、真实 env、缓存、日志和非运行制品进入构建上下文。
- 变更：新增 `deploy/docker-compose.production.example.yml` 和 `deploy/config.production.example.yaml`，提供 production Compose 示例和无密钥配置样例。
- 变更：新增 `deploy.sh`，用于在远程 Linux 主机按参数/环境变量动态生成 `运行期显式部署参数` 并执行 Docker Compose 部署路径。
- 变更：扩展 `.github/workflows/deploy-remote.yml`，支持 `staging` / `production` 手动选择，确认词改为 `deploy-staging` 或 `deploy-production`。
- 变更：更新 `deploy.sh` / `script/install.sh` 显式参数契约、`docs/deployment.md` 和 README，补充 `APP_PORT`、`DEPLOY_CONTAINER_NAME`、Linux Docker、Windows 到远程 Linux 直接部署、GitHub Environment、production Secrets、目录权限和回滚边界说明。
- 范围：未修改 Go 代码、测试文件、导出业务 API、配置 schema、HTTP 路由、数据库 schema、`go.mod`、`go.sum`、真实 `.env`、真实服务器地址、部署凭据或密钥；未执行真实部署、未连接远程服务器、未推送镜像、未触发 workflow。
- 验证：
  - `docker version`：FAIL_ENV，当前环境未安装 Docker CLI
  - `podman` / `nerdctl` / `docker.exe`：NOT_AVAILABLE
  - `bash -n deploy.sh`：FAIL_ENV，本机无可用 bash，WSL 未安装 Linux 发行版
  - `go run mvdan.cc/sh/v3/cmd/shfmt@latest -ln bash -tojson`：PASS
  - 临时 Go YAML 解析：PASS
  - `go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml`：PASS
  - `go test ./... -count=1`：PASS
  - `go build -o <temp> ./cmd/server`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P2-004 PENDING_VERIFICATION；Docker build 待具备 Docker 的环境补跑。

### 2026-05-26 - TASK-P2-003 - TS-P2-003

- 变更：用户明确确认实现远程部署 workflow。
- 变更：新增 `.github/workflows/deploy-remote.yml`，提供手动触发的 staging 远程部署 workflow。
- 变更：workflow 使用 显式部署参数、`DEPLOY_SSH_KEY`、可选 `DEPLOY_SSH_KNOWN_HOSTS`、可选 `GHCR_USERNAME` / `GHCR_TOKEN` 等 GitHub Secrets。
- 变更：workflow 校验远程 SSH 输入，要求环境绑定确认词，通过 SSH 执行 `script/install.sh` 并把 GitHub Variables/Secrets 映射为 `deploy.sh` 显式参数。
- 变更：`deploy.sh` / `script/install.sh` 显式参数契约、`docs/deployment.md` 和 README 已补 workflow、Secrets、远程主机前置条件和手动触发说明。
- 范围：未修改 Go 代码、依赖、配置 schema、HTTP 路由、数据库 schema、真实 `.env`、真实服务器地址、部署凭据或密钥；未执行真实部署、未连接远程服务器、未推送镜像。
- 验证：
  - 临时 Go YAML 解析：PASS
  - `go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P2-003 COMPLETED；当前无自动下一实现任务。

### 2026-05-26 - TASK-P2-002 - TS-P2-002

- 变更：用户要求远程部署使用 `.env` 风格文件配置。
- 变更：新增 `deploy.sh` / `script/install.sh` 显式参数契约，提供远程部署目标、Docker Compose、镜像和健康检查变量占位。
- 变更：删除旧本地部署 env 文件依赖，部署配置改由显式参数传入。
- 变更：`docs/deployment.md` 增加远程部署变量说明，README 增加模板入口。
- 变更：同步需求、架构、测试矩阵、验收、Backlog、风险、决策、状态、测试报告和交接文档。
- 范围：未修改 `.github/workflows/*`，未实现真实部署、未连接服务器、未推送镜像、未写入真实密钥。
- 验证：
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P2-002 COMPLETED；当前无自动下一实现任务。

### 2026-05-26 - TASK-NEXT-SCOPE-010 - TS-NEXT-SCOPE-010

- 变更：用户选择 C，意图进入真实 CD / 镜像发布 / 远程部署自动化。
- 变更：用户补充使用远程部署；默认建议收敛为 GHCR + 手动触发 + staging + SSH 到 Linux 服务器 + Docker Compose。
- 变更：执行用户修正审查，结论为 `NEEDS_USER_DECISION`，确认前不得实现真实 CD workflow、推送镜像、连接远程环境或读取 secrets。
- 变更：新增 TASK-NEXT-SCOPE-010 / TS-NEXT-SCOPE-010 待确认状态，要求补充镜像仓库、SSH/Docker 等远程方式、发布环境、触发策略和 secrets 命名。
- 范围：仅更新项目状态文档；未修改 `.github/workflows/*`、Go 代码、依赖、配置 schema、数据库 schema、真实配置或密钥。
- 验证：
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：COMPLETED；后续由 TASK-P2-002 完成显式参数部署入口。

### 2026-05-26 - TASK-P2-001 - TS-P2-001

- 变更：用户选择 D，确认进入 CI/CD 与部署方向首切片。
- 变更：新增 `.github/workflows/ci.yml`，建立 GitHub Actions CI 质量门禁，报告 gofmt 漂移并执行全量测试、server 构建和空白检查。
- 变更：新增 `docs/deployment.md`，记录发布前检查、配置入口、手动运行、`initdb` 边界、手动发布步骤和未实现项。
- 变更：README 新增 CI 与部署说明入口。
- 变更：同步 `REQUIREMENTS.md`、`ARCHITECTURE.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`DECISIONS.md`、`ISSUES.md`、`TEST_REPORT.md` 和 `AGENT_HANDOFF.md`。
- 范围：未修改 Go 代码、测试文件、导出业务 API、配置 schema、HTTP 路由、数据库 schema、`go.mod`、`go.sum`、真实配置、部署凭据或密钥。
- 验证：
  - gofmt 漂移审计：KNOWN_DRIFT（历史 Go 文件格式漂移，已记录 `BL-025`）
  - `go test ./... -count=1`：PASS
  - `go build -o <temp> ./cmd/server`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P2-001 COMPLETED；当前无自动下一实现任务。

### 2026-05-26 - TASK-INFRA-003 - TS-INFRA-003

- 变更：用户发送“下一步”后执行状态恢复检查，发现背景文档仍保留 TASK-P1-016 前的 app 装配、reload/config 待补表述。
- 变更：新增 `docs/reports/status_diagnostics/2026-05-26-task-p1-017-post-completion-doc-drift.md`。
- 变更：同步 `ARCHITECTURE.md`、`MODULES.md`、`PROJECT_BRIEF.md` 和 `ROADMAP.md`，确认 TASK-P1-016 已覆盖 app 装配、配置变更 hook 与 reload/config 分发路径。
- 变更：同步 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`、`TEST_REPORT.md`、`ISSUES.md`、`RISK_REGISTER.md` 和 `AGENT_HANDOFF.md`。
- 范围：未修改 Go 代码、测试文件、导出业务 API、配置 schema、HTTP 路由、数据库 schema、`go.mod`、`go.sum`、部署配置或密钥。
- 验证：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-INFRA-003 COMPLETED；当前无自动下一实现任务。

### 2026-05-26 - TASK-P1-017 - TS-P1-017

- 变更：用户选择 A，确认进入 `BL-006` 第一阶段包 README 中文化。
- 变更：统一 `pkg/*/README.md` 包标题，将 `pkg/plugin/README.md` 英文主体转为中文，并把 License/FAQ/Unsupported 等读者文本中文化。
- 变更：同步 `pkg/cache`、`pkg/cli`、`pkg/executor`、`pkg/httpserver`、`pkg/i18n`、`pkg/storage`、`pkg/yaml2go` README 中过期的“缺少测试”风险描述。
- 变更：同步 `REQUIREMENTS.md`、`ARCHITECTURE.md`、`MODULES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`DECISIONS.md`、`ISSUES.md`、`TEST_REPORT.md` 和 `AGENT_HANDOFF.md`。
- 范围：未修改 Go 代码、导出业务 API、配置 schema、HTTP 路由、数据库 schema、`go.mod`、`go.sum`、部署配置或密钥。
- 验证：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-017 COMPLETED；当前无自动下一实现任务。

### 2026-05-26 - TASK-P1-016 - TS-P1-016

- 变更：用户明确要求实施 TASK-P1-016，确认提升 `BL-002` 剩余 app 装配、reload/config 集成测试范围。
- 变更：新增 `internal/app/app_integration_test.go`，使用临时 YAML、临时 SQLite 和真实 `app.New` 覆盖 server/initdb 装配、demo schema 创建、资源 shutdown 和 app 配置变更 hook。
- 变更：新增 `internal/app/reloadapp/reload_test.go`，使用 fake cache/database/logger/executor/httpserver/storage 覆盖 reload 未变化跳过、单组件变化分发、Redis/executor/storage 关闭置空和 database reload 不隐式迁移。
- 范围：未修改导出业务 API、配置 schema、HTTP 路由、数据库 schema、`go.mod`、`go.sum`、部署配置或密钥。
- 验证：
  - `gofmt -w internal/app/app_integration_test.go internal/app/reloadapp/reload_test.go`：PASS
  - `go test ./internal/app/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-016 COMPLETED；当前无自动下一实现任务。

### 2026-05-26 - TASK-PHASE6-001 - TS-PHASE6-001

- 变更：用户选择 A，确认进入 Phase 6 收尾与交接。
- 变更：关闭 TASK-NEXT-SCOPE-008，记录 TASK-PHASE6-001 / TS-PHASE6-001 完成。
- 变更：更新状态、任务、时间切片、验收、测试矩阵、路线图、项目简介、风险、Backlog、决策、问题记录、测试报告和交接说明。
- 范围：未新增或修改 Go 源码、测试文件、依赖、数据库 schema、部署配置或密钥。
- 验证：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：Phase 6 收尾完成；当前无自动下一实现任务，后续工作需要用户重新确认。

### 2026-05-26 - TASK-P1-015 - TS-P1-015

- 变更：用户选择 B，确认提升 `BL-002` 的 router/middleware/demo HTTP 集成测试部分。
- 变更：新增 `internal/transport/http/router_integration_test.go`，使用临时 SQLite 和真实 demo repository/service/handler 注入 `NewRouter`。
- 变更：覆盖 demo Todo HTTP Create/List/Get/Update/Delete、删除后 404、CORS preflight/actual origin header、TraceID header round-trip，以及 Recovery 500 响应 traceId 和 logger 调用。
- 修复：前两次相关包测试失败来自测试构造问题：`httptest.NewRequest` 默认 Host 与 Origin 同源，导致 CORS 中间件跳过；固定测试 Host 为 `api.local` 后通过。
- 验证：
  - `gofmt -w internal/transport/http/router_integration_test.go`：PASS
  - `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-015 COMPLETED；TASK-NEXT-SCOPE-008 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-014 - TS-P1-014

- 变更：用户选择 B，确认提升 `BL-023`，为 `pkg/utils` 内部支撑工具补最小确定性测试。
- 变更：新增 `pkg/utils/utils_test.go`，覆盖 Snowflake 生成/非法 node、监听地址校验、端口范围与 exclude、设备 ID 稳定/盐值和 i18n helper 默认语言转发。
- 修复：前两次包测试失败来自测试代码对端口占用语义的环境假设；改为确定性无效地址和 exclude/range 断言后通过。
- 验证：
  - `gofmt -w pkg/utils/utils_test.go`：PASS
  - `go test ./pkg/utils -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-014 COMPLETED；TASK-NEXT-SCOPE-007 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-013 - TS-P1-013

- 变更：用户选择 A，确认继续 `BL-020` 剩余范围，提升第三批 `pkg/cache` 隔离行为测试。
- 变更：新增 `pkg/cache/cache_test.go`，使用进程内 Redis 测试服务覆盖配置默认值、配置校验、连接、读写、缺失键、过期、批量操作、计数器和 reload 语义。
- 变更：新增纯测试依赖 `github.com/alicebob/miniredis/v2`，同步更新 `go.mod` 和 `go.sum`。
- 修复：首次包测试为测试代码编译失败，原因是误读 `miniredis.Get` 返回值；修正断言后通过。
- 验证：
  - `go get github.com/alicebob/miniredis/v2@latest`：PASS
  - `gofmt -w pkg/cache/cache_test.go`：PASS
  - `go test ./pkg/cache -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-013 COMPLETED；TASK-NEXT-SCOPE-006 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-012 - TS-P1-012

- 变更：用户发送“下一步”，确认继续 `BL-020` 第二批 `pkg/*` 行为测试。
- 变更：新增 `pkg/executor/executor_test.go`，覆盖配置校验、任务执行、缺失池、过载、关闭、失败 reload 和 panic handler。
- 变更：新增 `pkg/httpserver/httpserver_test.go`，覆盖构造、默认配置、配置错误、停止态 reload/shutdown 和已运行 start 拒绝路径。
- 变更：新增 `pkg/storage/storage_test.go`，覆盖内存文件系统读写、复制、MIME、Excel、图片和配置错误路径。
- 修复：`pkg/executor` 缺失池、过载和重复配置错误改为包装公开 sentinel，支持 `errors.Is` 判断。
- 修复：`pkg/executor` panic 恢复路径现在会调用通过 `SetPanicHandler` 注册的 handler。
- 验证：
  - `gofmt -w pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`：PASS
  - `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-012 COMPLETED；TASK-NEXT-SCOPE-005 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-011 - TS-P1-011

- 变更：新增 `pkg/cli/app_test.go`，覆盖命令注册、flag/env/args 解析、help/version 输出、usage error 和 command error 包装。
- 变更：新增 `pkg/i18n/i18n_test.go`，覆盖 JSON/YAML 消息加载、模板渲染、默认语言回退、缺失消息 fallback、`MustT` panic 和加载错误路径。
- 变更：新增 `pkg/yaml2go/converter_test.go`，覆盖多文件生成、空输入、非法 YAML、配置校验，并用 Go parser 校验生成代码合法性。
- 修复：`pkg/yaml2go` 使用 Jennifer tag map 生成合法 struct tag，避免输出 `:"..."` 形式的非法 tag。
- 修复：`pkg/yaml2go` 将子配置 struct 与方法追加到同一个 Jennifer 文件，避免 import 块被拼接到声明之后。
- 验证：
  - `gofmt -w pkg/cli/app_test.go pkg/i18n/i18n_test.go pkg/yaml2go/converter_test.go`：PASS
  - `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-011 COMPLETED；TASK-NEXT-SCOPE-004 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-NEXT-SCOPE-003 - TS-NEXT-SCOPE-003

- 变更：用户选择 A，确认提升 `BL-020` 补 `pkg/*` 行为测试。
- 变更：将 `BL-020` 首批拆为 TASK-P1-011 / TS-P1-011，覆盖 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 最小行为测试。
- 变更：新增状态诊断报告 `docs/reports/status_diagnostics/2026-05-25-task-p1-011-handoff-stale.md`，修复交接和测试报告滞后。
- 变更：更新 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`DECISIONS.md`、`TEST_REPORT.md`、`ISSUES.md` 和 `AGENT_HANDOFF.md`。
- 测试：
  - 状态一致性文本检查：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
  - Go 测试未运行：本切片仅修改项目文档，未改 Go 代码
- 状态：TASK-NEXT-SCOPE-003 COMPLETED；TASK-P1-011 NOT_STARTED。

### 2026-05-25 - TASK-P1-010 - TS-P1-010

- 变更：接受用户修正，`pkg/plugin` 注册责任改为被动 registry/runtime。
- 变更：新增状态诊断报告 `docs/reports/status_diagnostics/2026-05-25-agent-handoff-stale-after-types-boundary.md`，记录并修复交接文件滞后。
- 变更：`Manager` 接口移除 `Load`、`RegisterLocalFactory` 和 manager option 主动装配公共面。
- 变更：新增 `NewHTTP` 和 HTTP option，使 HTTP 插件可由插件服务或宿主装配层显式构造后 `Register`。
- 变更：移除 local factory API 和相关错误，local 插件改为服务侧构造后注册。
- 变更：更新 `pkg/plugin` README、package doc、架构、模块清单、测试矩阵、验收、风险、Backlog、状态和交接文档。
- 测试：
  - `gofmt -w pkg/plugin/manager.go pkg/plugin/http.go pkg/plugin/constants.go pkg/plugin/errors.go pkg/plugin/doc.go pkg/plugin/config.go pkg/plugin/local.go pkg/plugin/plugin_test.go`：PASS
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-010 COMPLETED；TASK-NEXT-SCOPE-003 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-NEXT-SCOPE - TS-NEXT-SCOPE

- 变更：用户选择选项 A，确认提升 `BL-021` / `TM-P1-005`。
- 变更：新增 TASK-P1-009 / TS-P1-009，作为当前唯一合法下一步，用于明确 `types/*` 契约边界。
- 变更：更新 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`BACKLOG.md`、`RISK_REGISTER.md` 和交接相关文档，关闭待确认状态。
- 测试：
  - 状态一致性文本检查：PASS
  - `go test ./types/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-NEXT-SCOPE COMPLETED；TASK-P1-009 NOT_STARTED。

### 2026-05-25 - TASK-P1-008 - TS-P1-008

- 变更：新增 `ErrCodeUnsupportedOperation` 和 `NewUnsupportedError`，统一表示 `pkg/sqlgen` 未支持能力。
- 变更：`Or`、`Not`、`Group`、`Having`、`Distinct`、`Joins` 不再静默 no-op，后续 SQL 生成会返回 unsupported 错误。
- 变更：`DeleteInBatches` 不再退化为普通删除，直接返回 unsupported 错误。
- 变更：`ReverseDB(...).Generate`、`GenerateAll`、`GenerateToDir` 返回 unsupported 错误。
- 变更：`pkg/sqlgen/README.md` 标注 unsupported / partial 能力边界，`doc.go` 不再声称完整 GORM 兼容。
- 变更：新增 `pkg/sqlgen` unsupported 行为测试。
- 测试：
  - `gofmt -w pkg/sqlgen/errors.go pkg/sqlgen/types.go pkg/sqlgen/generator.go pkg/sqlgen/query.go pkg/sqlgen/update.go pkg/sqlgen/delete.go pkg/sqlgen/reverse.go pkg/sqlgen/doc.go pkg/sqlgen/sqlgen_test.go`：PASS
  - `go test ./pkg/sqlgen -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-008 COMPLETED；后续范围 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-007 - TS-P1-007

- 变更：为 13 个 `pkg/*/README.md` 新增 API 分类段。
- 变更：将 `pkg/cache`、`pkg/crypto`、`pkg/database`、`pkg/executor`、`pkg/httpserver`、`pkg/i18n`、`pkg/logger`、`pkg/plugin`、`pkg/storage` 标注为公共基础设施 API。
- 变更：将 `pkg/cli`、`pkg/sqlgen`、`pkg/yaml2go` 标注为公共工具 API。
- 变更：将 `pkg/utils` 标注为内部支撑工具包。
- 变更：同步 `ARCHITECTURE.md` 和 `MODULES.md`，记录稳定边界、测试缺口和后续约束。
- 变更：将当前合法下一步推进为 TASK-P1-008。
- 测试：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-007 COMPLETED；TASK-P1-008 NOT_STARTED。

### 2026-05-25 - TASK-P1-006 - TS-P1-006

- 变更：`cmd/server tests` 从 yaml2go 示例转换改为真实 Go test 入口。
- 变更：新增 `--package/-p`，默认执行 `go test ./...`，可指定 package pattern。
- 变更：移除 `TestsCommand.Execute` 中的 `log.Fatal` 行为，runner 失败时返回可包装错误。
- 变更：新增 `cmd/server/tests_test.go`，覆盖命令元信息、默认包范围、指定包范围和失败返回。
- 变更：新增 `docs/specs/cli_tests_command_boundary.md`，记录 CLI tests 命令语义。
- 变更：将当前合法下一步推进为 TASK-P1-007。
- 测试：
  - `go test ./cmd/server -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows CRLF 转换警告
- 状态：TASK-P1-006 COMPLETED；TASK-P1-007 NOT_STARTED。

### 2026-05-25 - TASK-P1-005 - TS-P1-005

- 变更：新增 demo 迁移触发策略，显式区分 `server-start`、`initdb` 和 `reload`。
- 变更：`NewModules` 继续在 server 启动路径执行 demo `AutoMigrate`，`BuildInitDB` 继续作为显式 demo bootstrap 执行迁移。
- 变更：`reloadDatabase` 改为使用 reload 策略，数据库 reload 不再隐式执行 demo schema 迁移。
- 变更：新增 `internal/app/initapp/demo_migration_test.go`，用隔离 SQLite 验证触发策略。
- 变更：新增 `docs/specs/demo_migration_boundary.md`，记录 dev/demo 与生产/bootstrap 迁移职责。
- 变更：将当前合法下一步推进为 TASK-P1-006。
- 测试：
  - `go test ./internal/app/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows CRLF 转换警告
- 状态：TASK-P1-005 COMPLETED；TASK-P1-006 NOT_STARTED。

### 2026-05-25 - TASK-P1-004 - TS-P1-004

- 变更：新增 `internal/modules/demo/service/todo_test.go`，为 demo Todo 建立 service/repository CRUD 测试基线。
- 变更：使用临时 SQLite 执行真实 repository/service 路径，不依赖外部数据库或 HTTP server。
- 变更：覆盖 Create/List/Get/Update/Delete 成功路径、空标题校验、缺失资源 not found 和软删除后不可见语义。
- 变更：将当前合法下一步推进为 TASK-P1-005。
- 测试：
  - `go test ./internal/modules/demo/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows CRLF 转换警告
- 状态：TASK-P1-004 COMPLETED；TASK-P1-005 NOT_STARTED。

### 2026-05-25 - TASK-P1-003 - TS-P1-003

- 变更：新增 `internal/transport/http/router_test.go`，用 `httptest` 固定 `/health` 和 `/ready` 行为。
- 变更：`/health` 覆盖 HTTP 200、`code=0`、`message=success`、`data.status=ok`。
- 变更：`/ready` 覆盖数据库缺失、ping 失败、ping 成功三条路径，断言 HTTP 状态码、`data.status` 和 `data.checks.database`。
- 变更：将当前合法下一步推进为 TASK-P1-004。
- 测试：
  - `go test ./internal/transport/http -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-P1-003 COMPLETED；TASK-P1-004 NOT_STARTED。

### 2026-05-25 - TASK-P1-002 - TS-P1-002

- 变更：数据库环境变量覆盖改为优先读取 `DB_*`，旧 `REI_APP_DB_*` 保留为兼容 fallback。
- 变更：`.env.example` 对齐实际环境变量策略，补齐 Storage/CORS 示例，移除未实现的 JWT 示例。
- 变更：新增配置测试，覆盖 `DB_*` 主策略、旧前缀 fallback、Redis/Server/Logger/I18n 环境变量覆盖。
- 变更：将当前合法下一步推进为 TASK-P1-003。
- 测试：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-P1-002 COMPLETED；TASK-P1-003 NOT_STARTED。

### 2026-05-25 - TASK-INFRA-002 - TS-INFRA-002

- 变更：新增实际缺失的 `AGENTS.md`，修复状态文档与文件系统事实冲突。
- 变更：统一 `CLAUDE.md`、`AGENT_RULES.md`、Cursor、Kiro、Codex 配置对 `AGENTS.md` 和 `docs/ai/prompt.md` 的引用。
- 变更：扩充 14 个 canonical `skills/*/SKILL.md`，补齐 YAML frontmatter 和完整执行结构。
- 变更：新增 14 个 `.agents/skills/*/SKILL.md` 轻量适配器。
- 变更：将 `docs/templates/*` 标准化为可复用模板，项目事实继续保留在根目录项目文档。
- 变更：新增状态诊断报告 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md`。
- 测试：
  - Agent 基础设施文件存在性核对：PASS
  - `quick_validate.py` 验证 28 个 skill 目录：PASS
  - 跨工具入口引用一致性检查：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-INFRA-002 COMPLETED；TASK-P1-002 NOT_STARTED。

### 2026-05-25 - TASK-INFRA-001 - TS-INFRA-001

- 变更：补齐 `docs/ai/prompt.md` 要求的跨 Agent 入口、规则和 skills 索引。
- 变更：新增任务拆分模板、时间切片模板、reports/specs 目录入口和跨工具目录入口。
- 变更：新增 14 个项目专用 `skills/*/SKILL.md`。
- 变更：记录 TASK-INFRA-001 完成，并恢复当前合法下一步为 TASK-P1-002。
- 测试：
  - Prompt 全量产物存在性核对：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-INFRA-001 COMPLETED；TASK-P1-002 NOT_STARTED。

## 历史变更

### 2026-05-25 - TASK-P1-001 - TS-P1-001

- 变更：用户再次发送“下一步”，按推荐默认顺序确认 P1 执行顺序。
- 变更：修复 `internal/config/manager.go` 的 `copyConfig`，改为完整结构体复制并深拷贝 slice。
- 变更：新增 `internal/config/manager_test.go`，覆盖字段完整复制、slice 深拷贝和 `Update` 保留未修改字段。
- 变更：将当前合法下一步推进为 TASK-P1-002。
- 测试：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-005 COMPLETED；TASK-P1-001 COMPLETED；TASK-P1-002 NOT_STARTED。

### 2026-05-25 - TASK-OPT-004 - TS-OPT-004

- 变更：新增 `TEST_MATRIX.md`，正式记录 P0/P1 测试矩阵、验证命令、退出条件和推荐执行顺序。
- 变更：新增 `ISSUES.md`，补齐项目问题记录入口。
- 变更：在 `TASKS.md` 和 `TIME_SLICES.md` 中写入 TASK-P1-001 至 TASK-P1-008、TS-P1-001 至 TS-P1-008 草案。
- 变更：将当前合法下一步推进为 TASK-OPT-005，等待确认 P1 执行顺序。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-004 COMPLETED；TASK-OPT-005 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-OPT-003 - TS-OPT-003

- 变更：新增 `MODULES.md`，记录模块职责、依赖方向、边界冲突、测试矩阵草案和 P1 优化候选项。
- 变更：确认 `.env.example` 与数据库环境变量前缀不一致、`copyConfig` 字段复制不完整、demo 自动迁移触发点分散、`cmd/server tests` 语义不一致等为优先收口风险。
- 变更：更新需求、验收、路线图、Backlog、任务、时间切片、状态、测试报告和交接文档。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-003 COMPLETED；TASK-OPT-004 NOT_STARTED。

### 2026-05-25 - TASK-OPT-002 - TS-OPT-002

- 变更：按用户“下一步”确认推荐默认值。
- 变更：确认治理优先、`pkg/*` 混合策略、demo 长期标准示例、迁移 dev-prod 分层、中文化根文档和模板优先。
- 变更：新增 `ROADMAP.md`。
- 变更：更新需求、架构、决策、任务、时间切片、状态、验收、风险和交接文档。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-002 COMPLETED；TASK-OPT-003 NOT_STARTED。

### 2026-05-25 - TASK-OPT-001 - TS-OPT-001

- 变更：重新启动全项目治理与优化路线主线。
- 变更：将六个启动模板重写为中文：
  - `docs/templates/project_start_template.md`
  - `docs/templates/requirements_clarification_template.md`
  - `docs/templates/technical_options_template.md`
  - `docs/templates/architecture_constraints_template.md`
  - `docs/templates/acceptance_template.md`
  - `docs/templates/risk_confirmation_template.md`
- 变更：更新根目录项目文档和状态文件，使当前合法任务从插件系统扩展切换为项目优化启动确认。
- 变更：将插件系统 v1 内容保留为历史记录和 Backlog，不作为当前主线继续扩展。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-001 COMPLETED；TASK-OPT-002 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-HIST-PLUGIN-002 - TS-HIST-PLUGIN-002

- 历史：接受并关闭 `pkg/plugin` v1 local/http API 边界。
- 测试：
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：COMPLETED。

### 2026-05-25 - TASK-HIST-PLUGIN-001 - TS-HIST-PLUGIN-001

- 历史：实现独立 `pkg/plugin` 包，支持 local 和 HTTP 协议。
- 测试：
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：COMPLETED。
## 2026-05-27

- Replaced the old database init path with `cmd/server db`, backed by `pkg/sqlgen`.
- Added sqlgen-generated database DDL, demo schema print/apply, and Todo CRUD operations.
- Removed `initdb`, InitDB config, SQL bootstrap scripts, demo migration wrappers, and GORM `AutoMigrate` usage from current code/config paths.
- Updated docs, acceptance, matrix, risk, and handoff records for TASK-P2-014.
