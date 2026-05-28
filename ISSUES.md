# ISSUES.md

## Latest Issue Review: TASK-P2-023

- Date: 2026-05-28
- Result: No new blocking issue after verification.
- Note: Auth token issue/verify now uses the public `pkg/auth` JWT-backed token API; business code no longer owns hand-written JWT signing/parsing.
- Verification: `pkg/auth`, user module, initapp, app, router, full root regression, and diff-check passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: `pkg/auth` does not replace production IAM hardening; refresh/session/audit/password-reset work, secret rotation, production migrations, deployment, real secrets/users, and plugin transport changes remain out of scope.

## Latest Issue Review: TASK-P2-022

- Date: 2026-05-28
- Result: No new blocking issue after verification.
- Note: RBAC authorization now uses the public `pkg/rbac` Casbin-backed authorizer API; business code no longer imports `internal/modules/user/rbac`.
- Verification: `pkg/rbac`, config, user module, initapp, app, router, full root regression, and diff-check passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: `pkg/rbac` does not replace production IAM hardening; refresh/session/audit/password-reset work, production migrations, deployment, real secrets/users, and plugin transport changes remain out of scope.

## Latest Issue Review: TASK-P2-021

- Date: 2026-05-28
- Result: No new blocking issue after verification.
- Note: RBAC authorization now uses a Casbin-backed local wrapper and DB-backed role-permission policies; `configs/rbac_model.conf` stores the recoverable model.
- Verification: config, user module, initapp, app, router, full root regression, and diff-check passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: Casbin integration does not replace production IAM hardening; refresh/session/audit/password-reset work, production migrations, deployment, real secrets/users, and plugin transport changes remain out of scope.

## Latest Issue Review: TASK-P2-020

- Date: 2026-05-28
- Result: No new blocking issue after verification.
- Note: RBAC configuration now seeds roles, permissions, and role-permission grants only; no real users/passwords or production policy engine were added.
- Verification: config, user service, initapp, app, full root regression, and diff-check passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: production IAM hardening, OPA/Casbin, external IAM, refresh/session/audit/password-reset work, production migrations, deployment, and plugin transport changes remain out of scope.

## Latest Issue Review: TASK-P2-019

- Date: 2026-05-28
- Result: No new blocking issue after verification.
- Note: User selected both auth hardening and plugin WS/RPC/discovery. The current legal slice is limited to auth token config hardening; plugin WS/RPC/discovery remains deferred as `BL-028`.
- Verification: config, initapp, app, full root regression, and diff-check passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: real secret management, refresh-token/session revocation, audit logging, password reset, external IAM, OPA/Casbin, production migration, deployment, and plugin transport changes remain out of scope.

## Latest Issue Review: TASK-P2-018

- Date: 2026-05-28
- Result: No new blocking issue after verification.
- Note: Full root regression passed after one in-scope repair to update existing `internal/app/initapp/transport_test.go` call sites for the new `NewHTTPServer` user handler parameter.
- Verification: user module, dbapp, router, app/initapp target tests, full `go test ./... -count=1`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: production secret/session management, refresh-token/session revocation, audit logging, password reset, OPA/Casbin, external IAM, production migration, deployment, and plugin transport changes remain out of scope.

## Latest Issue Review: TASK-INFRA-004

- Date: 2026-05-28
- Result: CI failure fixed locally; no new blocking issue after verification.
- Failure: GitHub Actions run `26531295923` failed in `CI / Go verify` at the `Build server` step with `stat .../cmd/server: directory not found`.
- Root cause: CI still referenced the removed `./cmd/server` entrypoint after the repository moved to `./cmd/main`.
- Fix: `.github/workflows/ci.yml` now builds `./cmd/main`.
- Verification: GitHub job logs inspected; `go test ./... -count=1 -mod=readonly`, `go build -mod=readonly -o <temp> ./cmd/main`, `actionlint`, and `git diff --check` passed.
- Residual note: broader historical documentation may still mention `cmd/server`; this slice only repairs the failing CI build target and records the drift for future confirmed cleanup if needed.

## Latest Issue Review: TASK-P2-017

- Date: 2026-05-28
- Result: No new blocking issue for TASK-P2-017.
- Note: Plugin registration exposure is now explicit: dedicated plugin HTTP interface must be enabled, or main HTTP exposure must be opted in with `plugin.registration.expose_on_main_http`. WS remains a reserved URL placeholder only.
- Verification: config/app/router/plugin target tests, root `go test ./...`, Blog module `go test ./...`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: broad network exposure of the plugin control plane, real WS/RPC transport, plugin heartbeat/persistent discovery, JWT/login, database-backed IAM, production deployment, and real secret management remain out of scope.

## Previous Issue Review: TASK-P2-016

- Date: 2026-05-28
- Result: No new blocking issue.
- Note: Existing workspace had pre-existing `cmd/server` deletion and `cmd/main` addition before this slice; this slice did not revert or modify that move. Full root regression passed against the current workspace layout.
- Verification: plugin/app/router/IAM target tests, root `go test ./...`, Blog module `go test ./...`, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Residual risks: real WS/RPC transport, plugin heartbeat/persistent discovery, JWT/login, database-backed IAM, production deployment, and real secret management remain out of scope.

## Latest Issue Review: TASK-P2-015

- Date: 2026-05-27
- Result: No new blocking issue.
- Note: This slice only added DB CLI comments and documentation. It did not change DB behavior, execute schema changes, restore old init paths, or touch production migration flow.
- Verification: `go test ./cmd/server -count=1`, `go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1`, DB docs `rg` scan, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.

## 最新补充

- ISSUE-CONFIG-DOCS-001：无新增失败项。TASK-P2-013 仅新增配置文档说明和入口，未修改 Go 实现、真实 `.env`、密钥、数据库 schema、HTTP 路由、业务模块或生产部署配置；关键文档检索、`go test ./internal/config -count=1`、`go test ./... -count=1` 和 `git diff --check` 均通过。

## Issue 状态

- 项目：go-scaffold
- 最后更新：2026-05-28
- 规则：失败、返工和阻塞问题记录在本文；风险项仍记录在 `RISK_REGISTER.md`。

## Open Issues

| ID | Linked Task | Severity | Status | Summary | Next Action |
|---|---|---|---|---|---|
| 无 |  |  |  | 当前无 open issue |  |

## Issue Details

- ISSUE-CONFIG-002：无新增失败项。用户指出 `EnvDBDriver`、`EnvDBHost` 等重复 env-name 常量已无存在必要；本轮已删除 `internal/config/constants.go` 中按模块镜像 `envname` 标签值的常量，并将配置测试改为从结构体标签读取环境变量名。目标包测试、cmd/app 相关测试、全量回归和 diff 检查均通过。
- ISSUE-CONFIG-001：无新增失败项。用户要求 `internal/config` 基于 `AppPrefix` 动态生成环境变量前缀、使用 `envname` 字段标签并自动加载 `.env`；本轮已实现 `RIN_APP_*` 主前缀、未加前缀 fallback、`.env` 自动加载测试和 app/cmd 相关回归验证。
- ISSUE-TYPES-002：无新增失败项。用户要求将 `types/constants.AppPrefix` 改为 `Rin` 并删除 `AppTestsCommandName`；本轮已将 tests 命令名收回 `cmd/server` 私有常量，新增应用常量测试，并通过目标包和全量回归验证。
- ISSUE-TYPES-001：CLOSED。用户纠正 `types` 分层：不能直接向 `pkg/crypto.Crypto` 提供别名，也不能定义依赖 `pkg/cache.Cache` 的 `CacheInjectable`；`types` 只能承载应用层以上确认过的跨层契约。本轮已删除 `types/interfaces.go`，将 `types/constants` executor pool 名称改为字符串常量，更新 `types/doc.go` 和 `docs/specs/types_contract_boundary.md`，并新增 `types/import_boundary_test.go` 固定 `types/*` 不得导入 `pkg/*`。
- ISSUE-STATUS-004：CLOSED。用户纠正当前项目还未开发完整，不应发布第一版；此前项目总体状态写为 `COMPLETED`，容易把 TASK-P2-004 Docker build 通过误判为 v1 release-ready。本轮已将项目整体状态改为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，并在状态、验收、风险、决策、路线图、Backlog、README、部署说明和交接文档中明确 Docker build 只作为切片证据，不代表第一版发布。
- ISSUE-P2-005：CLOSED。TASK-P2-004 已补齐 Dockerfile、production Compose 示例、production 配置样例、统一 `deploy.sh` 部署入口和手动 production workflow 闸门；shfmt Bash parser、临时 Go YAML 解析、actionlint、旧引用 `rg` 检查、全量 Go 回归、server build 和 `git diff --check` 均通过。2026-05-27 用户在 Linux Docker 环境补跑 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` 成功，BuildKit 输出 `23/23 FINISHED`，镜像标记为 `docker.io/library/go-scaffold:local`。
- ISSUE-P2-006：无新增失败项。TASK-P2-005 至 TASK-P2-010 已完成插件钩子运行时、HTTP 远程插件传输、IAM 公共接口、配置/app/reload/lifecycle 接入；`go test ./pkg/plugin/... -count=1`、`go test ./pkg/iam/... -count=1`、`go test ./internal/config ./internal/app/... -count=1`、`go test ./... -count=1`、server build 和 `git diff --check` 均通过。
- ISSUE-INFRA-002：`AGENTS.md` 缺失但状态文件声称已补齐。已在 TASK-INFRA-002 中修复，诊断报告见 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md`。
- ISSUE-P1-002：`.env.example` 与数据库环境变量前缀不一致，且 JWT 示例暗示未实现能力。已在 TASK-P1-002 中修复，相关测试通过。
- ISSUE-P1-003：无新增失败项。TASK-P1-003 新增 router smoke test 后，包测试和全量回归均通过。
- ISSUE-P1-004：无新增失败项。TASK-P1-004 新增 demo CRUD 测试基线后，demo 模块测试和全量回归均通过。
- ISSUE-P1-005：无新增失败项。TASK-P1-005 收拢 demo 迁移边界后，app 包测试和全量回归均通过。
- ISSUE-P1-006：无新增失败项。TASK-P1-006 收拢 CLI tests 命令语义后，cmd/server 包测试和全量回归均通过。
- ISSUE-P1-007：无新增失败项。TASK-P1-007 完成 `pkg/*` API 分类后，全量回归通过。
- ISSUE-P1-008：无新增失败项。TASK-P1-008 标注 `pkg/sqlgen` unsupported 边界后，包测试和全量回归均通过。
- ISSUE-NEXT-001：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE 已将 `BL-021` / `TM-P1-005` 提升为 TASK-P1-009。
- ISSUE-P1-010：无新增失败项。TASK-P1-010 收拢 `pkg/plugin` 被动注册边界后，包测试和全量回归均通过。
- ISSUE-NEXT-003：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE-003 已将 `BL-020` 首批提升为 TASK-P1-011 / TS-P1-011。
- ISSUE-P1-011：无未解决失败项。TASK-P1-011 中 `pkg/yaml2go` 新增测试暴露生成 tag 与方法 import 顺序缺陷，已在当前允许范围内修复并通过回归。
- ISSUE-NEXT-004：无新增失败项。用户发送“下一步”后，TASK-NEXT-SCOPE-004 已将 `BL-020` 第二批提升为 TASK-P1-012 / TS-P1-012。
- ISSUE-P1-012：无未解决失败项。TASK-P1-012 中 `pkg/executor` 新增测试暴露 sentinel 错误包装和 panic handler 未调用缺陷，已在当前允许范围内修复并通过回归；测试自身等待竞态已在第二轮修复。
- ISSUE-NEXT-005：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE-005 已将 `BL-020` 第三批 `pkg/cache` 隔离行为测试提升为 TASK-P1-013 / TS-P1-013。
- ISSUE-P1-013：无未解决失败项。TASK-P1-013 首次包测试为测试代码编译失败，原因是误读 `miniredis.Get` 返回值；修正测试断言后，`pkg/cache` 包测试和全量回归均通过。
- ISSUE-NEXT-006：无新增失败项。用户选择 B 后，TASK-NEXT-SCOPE-006 已将 `BL-023` `pkg/utils` 内部支撑测试提升为 TASK-P1-014 / TS-P1-014。
- ISSUE-P1-014：无未解决失败项。TASK-P1-014 前两次包测试失败来自测试代码对端口占用语义的环境假设；改为确定性无效地址、端口范围和 exclude 断言后，`pkg/utils` 包测试和全量回归均通过。
- ISSUE-NEXT-007：无新增失败项。用户选择 B 后，TASK-NEXT-SCOPE-007 已将 `BL-002` router/middleware/demo HTTP 集成测试提升为 TASK-P1-015 / TS-P1-015。
- ISSUE-P1-015：无未解决失败项。TASK-P1-015 前两次相关包测试失败来自测试构造问题：`httptest.NewRequest` 默认 Host 与 Origin 同源导致 CORS 中间件跳过；固定测试 Host 为 `api.local` 后，相关包测试和全量回归均通过。
- ISSUE-NEXT-008：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE-008 已关闭并进入 TASK-PHASE6-001 / TS-PHASE6-001。
- ISSUE-PHASE6-001：无未解决失败项。Phase 6 收尾仅更新项目状态文档，最终 `go test ./... -count=1` 与 `git diff --check` 均通过。
- ISSUE-P1-016：无未解决失败项。TASK-P1-016 新增 app 装配与 reload/config 集成测试后，`go test ./internal/app/... -count=1`、`go test ./... -count=1` 与 `git diff --check` 均通过。
- ISSUE-P1-017：无未解决失败项。TASK-P1-017 第一阶段包 README 中文化后，`go test ./... -count=1` 与 `git diff --check` 均通过。
- ISSUE-INFRA-003：TASK-P1-016/017 完成后部分背景文档仍保留旧待办表述。已在 TASK-INFRA-003 中修复，诊断报告见 `docs/reports/status_diagnostics/2026-05-26-task-p1-017-post-completion-doc-drift.md`。
- ISSUE-P2-001：无未解决阻塞失败项。TASK-P2-001 新增 CI 质量门禁与部署说明后，全量测试、server 构建和 `git diff --check` 均通过；gofmt 漂移审计发现 82 个历史格式漂移文件，已记录 `BL-025`。
- ISSUE-P2-002：无实现失败项。用户选择 C 并确认使用远程部署后，当时真实 CD 自动化仍因缺少镜像仓库、SSH/Docker 等远程方式、环境、触发策略和 secrets 命名而处于 `PENDING_USER_CONFIRMATION`；后续已由 TASK-P2-003 完成手动 staging workflow。
- ISSUE-P2-003：无新增失败项。TASK-P2-002 后续改为 `deploy.sh` / `script/install.sh` 显式参数契约，并删除旧本地部署 env 文件依赖；后续已由 TASK-P2-003 完成手动 staging workflow。
- ISSUE-P2-004：无新增失败项。TASK-P2-003 新增手动 staging 远程部署 workflow 后，临时 Go YAML 解析、actionlint 和 `git diff --check` 均通过；未执行真实部署、未连接远程服务器、未写入真实密钥。

## 历史说明

- 2026-05-25：记录并关闭 `.env.example` 与数据库环境变量实现不一致问题。
- 2026-05-25：记录 TASK-P1-003 无新增失败项，HTTP router smoke test 和全量回归通过。
- 2026-05-25：记录 TASK-P1-004 无新增失败项，demo CRUD 测试和全量回归通过。
- 2026-05-25：记录 TASK-P1-005 无新增失败项，demo 迁移策略测试和全量回归通过。
- 2026-05-25：记录 TASK-P1-006 无新增失败项，CLI tests 命令语义测试和全量回归通过。
- 2026-05-25：记录 TASK-P1-007 无新增失败项，`pkg/*` API 分类后全量回归通过。
- 2026-05-25：记录 TASK-P1-008 无新增失败项，`pkg/sqlgen` unsupported 行为测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE 无新增失败项，`types/*` 契约边界已提升为下一合法任务。
- 2026-05-25：记录 TASK-P1-010 无新增失败项，`pkg/plugin` 被动注册边界测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE-003 无新增失败项，首批 `pkg/*` 行为测试已排期。
- 2026-05-25：记录 TASK-NEXT-SCOPE-004 无新增失败项，第二批 `pkg/*` 行为测试已排期。
- 2026-05-25：记录 TASK-P1-012 无未解决失败项，`pkg/executor` 暴露缺陷已修复，第二批包测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE-005 无新增失败项，第三批 `pkg/cache` 行为测试已排期。
- 2026-05-25：记录 TASK-P1-013 无未解决失败项，测试代码编译问题已修复，`pkg/cache` 包测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE-006 无新增失败项，`pkg/utils` 内部支撑测试已排期。
- 2026-05-25：记录 TASK-P1-014 无未解决失败项，测试环境假设已修正，`pkg/utils` 包测试和全量回归通过。
- 2026-05-26：记录 TASK-NEXT-SCOPE-007 无新增失败项，router/middleware/demo HTTP 集成测试已排期。
- 2026-05-26：记录 TASK-P1-015 无未解决失败项，CORS 测试构造问题已修正，相关包测试和全量回归通过。
- 2026-05-26：记录 TASK-NEXT-SCOPE-008 与 TASK-PHASE6-001 无新增失败项，Phase 6 收尾完成。
- 2026-05-26：记录 TASK-P1-016 无未解决失败项，app 装配与 reload/config 集成测试通过。
- 2026-05-26：记录 TASK-P1-017 无未解决失败项，第一阶段包 README 中文化和全量回归通过。
- 2026-05-26：记录并关闭 TASK-P1-016/017 后背景文档状态漂移。
- 2026-05-26：记录 TASK-P2-001 无未解决失败项，CI 质量门禁和部署说明首切片完成。
- 2026-05-26：记录 TASK-P2-003 无新增失败项，手动 staging 远程部署 workflow 通过静态验证。
- 2026-05-27：记录 TASK-P2-004 部署流程重构后的环境待验证项，当前机器无 Docker CLI，Docker build 待补跑；其他静态验证和 Go 回归通过。
- 2026-05-27：用户发送“下一步”后复验 TASK-P2-004 Docker build 阻塞；`docker version` 失败，`docker`、`podman`、`nerdctl`、`docker.exe` 均不可用，`ISSUE-P2-005` 保持 OPEN，TASK-P2-004 / TS-P2-004 记录为 BLOCKED。
- 2026-05-27：用户再次发送“下一步”后复验同一阻塞；当前环境仍无 Docker 兼容 CLI，`docker build -t go-scaffold:local .` 未执行，`ISSUE-P2-005` 保持 OPEN。
- 2026-05-27：用户远端补跑 Docker build 时 `go mod download` 因 Go 代理网络超时失败；旧 Dockerfile 未声明 `GOPROXY` build arg，导致代理参数未生效。本轮已补 Dockerfile 代理参数和 BuildKit 缓存，`ISSUE-P2-005` 保持 OPEN。
- 2026-05-27：用户在 Linux Docker 环境用更新后的 Dockerfile 补跑 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` 成功，`ISSUE-P2-005` 关闭。
- 2026-05-27：用户纠正当前项目未达第一版发布条件；记录并关闭 `ISSUE-STATUS-004`，后续不得把当前切片完成误判为 v1 发布。
- 2026-05-27：记录 TASK-P2-005 至 TASK-P2-010 无新增失败项，插件钩子运行时、远程插件传输、IAM 公共接口和 app 组合层接入已通过验证。
- 2026-05-27：记录并关闭 `types/*` 分层问题；已移除 `Crypto` 别名、`CacheInjectable` 和 `types/constants` 对 `pkg/executor` 的直接依赖，并补充导入边界测试。
- 2026-05-27：记录 `internal/config` 动态环境变量前缀修正，无新增失败项；目标包、cmd/app 相关测试和全量回归均通过。
- 2026-05-27：记录 `internal/config` 重复 env-name 常量清理，无新增失败项；`envname` 标签成为字段环境变量名唯一事实源。
- 2026-05-25：记录并关闭 `AGENTS.md` 缺失导致的 Agent 入口冲突。
- 2026-05-25：创建 `ISSUES.md`，补齐 `docs/ai/prompt.md` 要求的项目问题记录入口。
## Latest Issue Review: TASK-P2-014

- Date: 2026-05-27
- Result: No new blocking issue.
- Note: Production migration framework remains intentionally out of scope and must be separately confirmed before real production schema changes.
