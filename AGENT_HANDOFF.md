# AGENT_HANDOFF.md

## Latest Current Handoff

- TASK-P2-015 / TS-P2-015: COMPLETED.
- `cmd/server db` now has maintainer comments, and `docs/db-cli.md` documents DB CLI overview, usage, operations, flags, layering, extension rules, forbidden regressions, and verification guidance.
- DB behavior remains unchanged and continues to be driven by `pkg/sqlgen`; old init/migration/script/AutoMigrate paths remain out of current code/config.
- Next legal state: `NONE / NONE / PENDING_USER_CONFIRMATION`.

## Latest Handoff Addendum

- Date: 2026-05-27
- Task: TASK-P2-015 / TS-P2-015
- Summary: Added comments to `cmd/server/db.go`, created `docs/db-cli.md`, and linked it from `docs/configuration.md` and `docs/deployment.md`.
- Files changed in this addendum: `cmd/server/db.go`, `docs/db-cli.md`, `docs/configuration.md`, `docs/deployment.md`, and project status documents.
- Verification: `go test ./cmd/server -count=1` PASS, `go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1` PASS, DB docs `rg` scan PASS, `git diff --check` PASS with only Git LF/CRLF notices.
- Legal next step remains `NONE / NONE / PENDING_USER_CONFIRMATION`; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`.

## Latest Addendum

- Date: 2026-05-27
- Task: TASK-P2-013 / TS-P2-013
- Summary: Added `docs/configuration.md` with configuration loading, `.env` auto-load, dynamic `RIN_APP_*` prefix, `RIN_CONFIG_PATH`, `envname` single source, common variables, and new config field workflow.
- Files changed in this addendum: `docs/configuration.md`, `README.md`, `docs/deployment.md`, and project status documents.
- Verification: config-doc `rg` check PASS, `go test ./internal/config -count=1` PASS, `go test ./... -count=1` PASS, `git diff --check` PASS with only Windows LF/CRLF warnings.
- Legal next step remains `NONE / NONE / PENDING_USER_CONFIRMATION`; project remains `IN_DEVELOPMENT_NOT_RELEASE_READY`.

## Last Updated

- Date: 2026-05-27
- Agent: Codex
- Tool: Codex Desktop

## Project Snapshot

- Project: go-scaffold
- Phase: 项目开发中，未达第一版发布条件
- Module: 项目优化路线
- Current Task: NONE
- Current Time Slice: NONE
- Overall Status: IN_DEVELOPMENT_NOT_RELEASE_READY

## What Was Done Last

- 接受用户修正：`internal/config` 中 `EnvDB*`、`EnvRedis*`、`EnvServer*` 等重复 env-name 常量在 `envname` 标签机制落地后已无存在必要。
- 删除 `internal/config/constants.go` 中按模块镜像字段环境变量名的常量，只保留动态前缀 helper、`.env` 文件名、分隔符和配置段名。
- 更新 `internal/config/manager_test.go`，测试通过 `taggedEnvName` 从配置结构体 `envname` 标签读取环境变量名，不再依赖第二套常量表。
- 验证动态前缀变量优先、未加前缀 fallback、`.env` 自动加载、cmd/app 相关集成和全量回归均不回归。
- 项目整体仍为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，当前等待用户确认新的开发范围或第一版发布验收清单。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `internal/config/constants.go` | Updated | 删除重复 env-name 常量，只保留动态前缀和通用配置常量 |
| `internal/config/manager_test.go` | Updated | 测试从 `envname` 标签读取环境变量名 |
| `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`REQUIREMENTS.md`、`ARCHITECTURE.md`、`MODULES.md`、`DECISIONS.md`、`RISK_REGISTER.md`、`ISSUES.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`AGENT_HANDOFF.md` | Updated | 记录 TASK-P2-012 / TS-P2-012 完成 |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| 用户纠正审查 | ACCEPT：重复 env-name 常量会造成第二事实源，删除后以 `envname` 标签为准 |
| `gofmt -w internal/config/constants.go internal/config/manager_test.go` | PASS |
| env 常量引用扫描 | PASS：`Env(DB|Redis|Server|Log|I18n|CORS|InitDB|Executor|Storage|Plugin|IAM)` 无匹配 |
| `go test ./internal/config -count=1` | PASS |
| `go test ./cmd/server ./internal/app/... -count=1` | PASS |
| `go test ./... -count=1` | PASS |
| `git diff --check` | PASS，仅有 Windows LF/CRLF 提示 |

## Test Status

- `internal/config` env-name single source: PASS. 重复 env-name 常量已删除，测试从 `envname` 标签读取变量名。
- `internal/config` dynamic env prefix: PASS. `RIN_APP_*` 动态前缀、`envname` 覆盖、未加前缀 fallback 和 `.env` 自动加载均有测试。
- `cmd/server` config path env var: PASS. 配置路径 flag 使用 `RIN_CONFIG_PATH`。
- `internal/app/...`: PASS. app 初始化、reload 和环境清理相关测试通过。
- `types/constants` application constants: PASS from the previous correction. `AppPrefix` 固定为 `Rin`，`AppTestsCommandName` 已删除。
- `types` package boundary: PASS from the previous correction. `types/*` 导入边界测试、`types/result`、`types/errors`、`types/constants` 均通过。
- Full regression: PASS.
- Docker image build: PASS_REMOTE for TASK-P2-004. 用户 Linux Docker build 已通过。
- `pkg/plugin` and `pkg/plugin/hooks`: PASS from the previous completion audit.
- `pkg/iam` and `pkg/iam/memory`: PASS from the previous completion audit.
- `internal/config` and `internal/app/...`: PASS from the previous completion audit.
- Full regression and server build: PASS from the previous completion audit.
- Diff whitespace check: PASS for this code-and-documentation update; only Windows LF/CRLF warnings were printed.
- Release readiness: NOT_READY. 当前不得发布第一版。

## Current Blockers

- 当前任务无未关闭阻塞项，但第一版发布被 `RISK-022` 阻塞。真实 production 运行、镜像发布流水线、生产迁移、完整 auth/rbac、插件扩展和 v1 发布验收清单仍需用户重新确认并拆分新任务。

## Important Decisions

- [CONFIRMED] `dev.tmp/new-pllugin.md` is a typo; the design source is `dev.tmp/new-plugin.md`.
- [ACCEPT_WITH_RISK] Mainline switched to hook-aware plugin runtime, HTTP remote plugin transport and independent IAM public API.
- [CONFIRMED] `pkg/plugin` and `pkg/iam` remain decoupled public infrastructure packages.
- [CONFIRMED] `internal/app` is the composition root for binding IAM authorization hooks, config-created HTTP plugin adapters and lifecycle.
- [CONFIRMED] Config-created plugins are HTTP adapters only; local plugins remain explicitly registered by code.
- [CONFIRMED] TASK-P2-004 Docker verification is complete; no automatic next implementation task is active.
- [ACCEPTED] 当前项目未达第一版发布条件；Docker build 和部署制品完成不等于 v1 release-ready。
- [ACCEPTED] `types/*` 不再聚合 `pkg/*` 基础设施接口；缓存、加密、executor 等能力由应用层显式依赖对应 `pkg/*` 包，或在应用层以上另行定义契约。
- [ACCEPTED] `types/constants` 不再提供 tests 命令名；`cmd/server` 自行维护 `tests` 命令名，`AppPrefix` 为 `Rin`。
- [ACCEPTED] `internal/config` 环境变量覆盖主前缀从 `AppPrefix` 动态派生，当前为 `RIN_APP`；配置字段通过 `envname` 标签声明环境变量名。
- [ACCEPTED] 配置字段环境变量名以 `envname` 标签为唯一事实源；不要恢复 `EnvDB*` / `EnvRedis*` 等镜像常量。

## Risks

- Some existing workspace changes may predate this slice; do not revert unrelated user or prior-Agent changes.
- Plugin hooks can become a hidden control plane if future work registers broad handlers without tests; keep hook points explicit and covered.
- IAM memory service is infrastructure only, not business login/RBAC; do not market it as complete authentication.
- Remote hook calls use the plugin invoke path; keep `hooks.execute` isolated from manager hook emission to avoid recursion.
- Do not confuse the completed Docker image build with a real production deployment; this session did not deploy, push images, or run production migrations.
- Do not publish or label v1/release-ready until the user confirms a release acceptance checklist and the required tasks pass.
- Do not reintroduce `types.Crypto`, `types.CacheInjectable`, or direct `types/* -> pkg/*` imports.
- Do not reintroduce `types/constants.AppTestsCommandName`; keep the tests command name local to `cmd/server`.
- Do not reintroduce fixed `REI_APP` config override logic; use `internal/config.EnvPrefixJoin` and `envname` tags.
- Do not reintroduce `EnvDB*` / `EnvRedis*` / `EnvServer*` style env-name constants; read field names from `envname` tags when tests need them.

## Legal Next Step

- Task ID: NONE
- Time Slice ID: NONE
- Status: PENDING_USER_CONFIRMATION
- Why: TASK-P2-015 DB CLI comments/documentation is complete, TASK-P2-014 sqlgen DB CLI remains complete, and the project is explicitly not release-ready; current confirmed scope has no automatic next task.
- Entry condition for future work: user must confirm a new scope or first-version release acceptance checklist, and it must be written into `TASKS.md` and `TIME_SLICES.md`.
- Likely next choices: define v1 acceptance checklist, complete product scope, image publishing pipeline, real staging/production run, production migration framework, complete auth/rbac, plugin discovery, or RPC/WS transport.

## Do Not Do

- Do not trigger GitHub workflow from this session.
- Do not connect to remote servers.
- Do not push images.
- Do not execute staging or production deployment.
- Do not run production migrations or irreversible database changes.
- Do not write, print or invent real `.env`, SSH key, token, password or production host values.
- Do not implement JWT middleware, database-backed IAM, OPA/Casbin, plugin discovery, RPC/WS transport or Go `.so` plugin loading without a new confirmed task.
- Do not commit, push or revert unrelated dirty workspace changes.

## Recovery Instructions

1. Read `AGENTS.md`.
2. Read `STATUS.md`, `TASKS.md`, and `TIME_SLICES.md`.
3. Confirm current state is `NONE / NONE / PENDING_USER_CONFIRMATION` and overall status is `IN_DEVELOPMENT_NOT_RELEASE_READY`.
4. Remember `types/*` no longer exposes `Crypto`, `CacheInjectable`, direct imports of `pkg/*`, or `AppTestsCommandName`; do not restore lower-layer aliases, typed constants, or the tests command constant.
5. Remember config env overrides now use `RIN_APP_*` from `AppPrefix=Rin` and `envname` tags; keep unprefixed variables only as fallback compatibility, and do not reintroduce duplicate env-name constants.
6. Do not start new implementation or publish v1 until the user confirms a new task or release acceptance checklist and the task/time-slice documents are updated.
## Latest Handoff: TASK-P2-015

- Date: 2026-05-27
- Status: COMPLETED
- Summary: `cmd/server db` now has maintainer comments, and `docs/db-cli.md` documents DB CLI overview, usage, operations, flags, layering, extension rules, forbidden regressions, and verification guidance.
- Verification completed: `go test ./cmd/server -count=1`, `go test ./pkg/sqlgen ./cmd/server ./internal/app/dbapp -count=1`, DB docs `rg` scan, and `git diff --check` passed. `git diff --check` emitted only Git LF/CRLF notices.
- Next legal state: `NONE / NONE / PENDING_USER_CONFIRMATION`.
- Important residual risk: production migration framework is not implemented and must be confirmed as a separate task before production schema changes.
