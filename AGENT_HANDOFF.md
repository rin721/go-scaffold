# AGENT_HANDOFF.md

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

- 接受用户纠正：当前项目还未开发完整，不应该发布第一版。
- Docker build 通过证据保持有效：用户在 Linux Docker 环境补跑 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` 并通过，BuildKit 输出 `23/23 FINISHED`，镜像标记为 `docker.io/library/go-scaffold:local`。
- TASK-P2-004 至 TASK-P2-010 的切片完成判断保持；它们不代表项目整体完成、v1 可发布或真实 production 已上线。
- 项目整体状态已改为 `IN_DEVELOPMENT_NOT_RELEASE_READY`，当前等待用户确认新的开发范围或第一版发布验收清单。
- 本轮未修改 Go 代码，未触发 workflow、未连接远程服务器、未推送镜像、未执行真实 production、未写入真实密钥。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `STATUS.md` | Updated | 将项目整体状态改为开发中且未达发布条件 |
| `TASKS.md` | Updated | 当前合法任务改为等待用户确认后续范围 |
| `TIME_SLICES.md` | Updated | 当前合法切片改为等待用户确认后续范围 |
| `ACCEPTANCE.md` | Updated | 新增发布验收状态，明确不发布第一版 |
| `TEST_REPORT.md` | Updated | 记录用户纠正审查和文档验证 |
| `CHANGELOG.md` | Updated | 新增 not release-ready 纠偏记录 |
| `ISSUES.md` | Updated | 记录并关闭状态表述过度完成的问题 |
| `PROJECT_BRIEF.md` | Updated | 同步项目仍未达发布条件 |
| `REQUIREMENTS.md` | Updated | 明确 v1 发布前仍需确认完整验收范围 |
| `ARCHITECTURE.md` | Updated | 明确 Docker 制品不等于 v1 架构完成 |
| `MODULES.md` | Updated | 同步模块文档完成不等于发布验收 |
| `ROADMAP.md` | Updated | 新增第一版发布条件确认阶段 |
| `BACKLOG.md` | Updated | 新增 `BL-027` 发布验收清单与剩余路线 |
| `RISK_REGISTER.md` | Updated | 新增 `RISK-022` 防止误判为第一版发布 |
| `DECISIONS.md` | Updated | 新增 `DEC-027` 当前项目未达第一版发布条件 |
| `TEST_MATRIX.md` | Updated | 新增第一版发布验收清单矩阵项 |
| `README.md` | Updated | 入口说明当前不是 v1 发布候选 |
| `docs/deployment.md` | Updated | 部署说明明确制品不代表第一版发布 |
| `AGENT_HANDOFF.md` | Updated | 交接说明指向 not release-ready 状态 |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| 用户纠正审查 | ACCEPT：当前项目未达第一版发布条件 |
| 用户 Linux Docker build 输出审查 | PASS_REMOTE：`23/23 FINISHED`，镜像标记为 `docker.io/library/go-scaffold:local` |
| Go tests | NOT_RUN，本轮仅更新文档，未修改 Go 代码 |
| `git diff --check` | PASS，仅有 Windows LF/CRLF 提示 |

## Test Status

- Docker image build: PASS_REMOTE for TASK-P2-004. 用户 Linux Docker build 已通过。
- `pkg/plugin` and `pkg/plugin/hooks`: PASS from the previous completion audit.
- `pkg/iam` and `pkg/iam/memory`: PASS from the previous completion audit.
- `internal/config` and `internal/app/...`: PASS from the previous completion audit.
- Full regression and server build: PASS from the previous completion audit.
- Diff whitespace check: PASS for this documentation-only update; only Windows LF/CRLF warnings were printed.
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

## Risks

- Some existing workspace changes may predate this slice; do not revert unrelated user or prior-Agent changes.
- Plugin hooks can become a hidden control plane if future work registers broad handlers without tests; keep hook points explicit and covered.
- IAM memory service is infrastructure only, not business login/RBAC; do not market it as complete authentication.
- Remote hook calls use the plugin invoke path; keep `hooks.execute` isolated from manager hook emission to avoid recursion.
- Do not confuse the completed Docker image build with a real production deployment; this session did not deploy, push images, or run production migrations.
- Do not publish or label v1/release-ready until the user confirms a release acceptance checklist and the required tasks pass.

## Legal Next Step

- Task ID: NONE
- Time Slice ID: NONE
- Status: PENDING_USER_CONFIRMATION
- Why: TASK-P2-004 through TASK-P2-010 are complete and verified, but the project is explicitly not release-ready; current confirmed scope has no automatic next task.
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
4. Do not start new implementation or publish v1 until the user confirms a new task or release acceptance checklist and the task/time-slice documents are updated.
