# AGENT_HANDOFF.md

## Last Updated

- Date: 2026-05-27
- Agent: Codex
- Tool: Codex Desktop

## Project Snapshot

- Project: go-scaffold
- Phase: P2 Linux Docker production 部署制品验证阻塞
- Module: 项目优化路线
- Current Task: TASK-P2-004
- Current Time Slice: TS-P2-004
- Overall Status: BLOCKED

## What Was Done Last

- 用户发送“下一步”后，按协议处理当前唯一未关闭事项：TASK-P2-004 / TS-P2-004 Docker image build 验证。
- 已执行 `docker version`，当前环境未安装 Docker CLI；已检查 `docker`、`podman`、`nerdctl`、`docker.exe`，均不可用。
- `docker build -t go-scaffold:local .` 因前置 Docker CLI/daemon 缺失未执行，TASK-P2-004 / TS-P2-004 已记录为 `BLOCKED`，`ISSUE-P2-005` 保持打开。
- TASK-P2-005 至 TASK-P2-010 插件钩子运行时、HTTP 远程插件传输、独立 IAM 公共接口、配置接入、app 装配、reload 和 lifecycle 仍保持 `COMPLETED`。
- 本轮未触发 workflow、未连接远程服务器、未推送镜像、未执行真实 production、未写入真实密钥。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `pkg/plugin/hooks/*` | Added | 独立钩子引擎、注册表、服务查找、停止语义和测试 |
| `pkg/plugin/manager.go`、`pkg/plugin/constants.go` | Updated | manager 钩子 API、标准钩子点、调用语义 |
| `pkg/plugin/http_server.go`、`pkg/plugin/remote_hook.go` | Added | HTTP 插件服务端 helper 和远程钩子适配器 |
| `pkg/plugin/plugin_test.go` | Updated | 覆盖钩子、HTTP server、RemoteHook 行为 |
| `pkg/iam/*`、`pkg/iam/memory/*` | Added | IAM 公共 API 和内存实现 |
| `internal/config/*` | Updated | 新增 `plugin` / `iam` 配置、默认值、校验、env override 和深拷贝 |
| `internal/app/initapp/*` | Updated | Infrastructure 新增 IAM/Plugins，server 模式装配和 app 层 hook 注册 |
| `internal/app/reloadapp/reload.go` | Updated | 配置重载先构建新 IAM/plugin 实例再替换，失败保留旧实例 |
| `internal/app/lifecycleapp/lifecycle.go` | Updated | HTTP server 停止后、cache/database 前关闭插件管理器 |
| Project status docs | Updated | 记录用户“下一步”后的 Docker 复验结果；TASK-P2-004 / TS-P2-004 为 BLOCKED，TASK-P2-005 至 TASK-P2-010 保持完成 |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| `docker version` | FAIL_ENV |
| `Get-Command docker,podman,nerdctl,docker.exe -ErrorAction SilentlyContinue` | NOT_AVAILABLE |
| `docker build -t go-scaffold:local .` | NOT_RUN，Docker CLI/daemon 不可用 |
| `git diff --check` | PASS |

## Test Status

- Docker image build: BLOCKED for TASK-P2-004 because Docker CLI/daemon is unavailable in the current environment.
- `pkg/plugin` and `pkg/plugin/hooks`: PASS in the previous TASK-P2-005 to TASK-P2-010 verification.
- `pkg/iam` and `pkg/iam/memory`: PASS in the previous TASK-P2-005 to TASK-P2-010 verification.
- `internal/config` and `internal/app/...`: PASS in the previous TASK-P2-005 to TASK-P2-010 verification.
- Full regression and server build: PASS in the previous TASK-P2-005 to TASK-P2-010 verification.
- Diff whitespace check: PASS after this status update.

## Current Blockers

- `ISSUE-P2-005` remains open for TASK-P2-004: run `docker build -t go-scaffold:local .` in a Docker-enabled environment before closing that deployment task.

## Important Decisions

- [CONFIRMED] `dev.tmp/new-pllugin.md` is a typo; the design source is `dev.tmp/new-plugin.md`.
- [ACCEPT_WITH_RISK] Mainline switched to hook-aware plugin runtime, HTTP remote plugin transport and independent IAM public API.
- [CONFIRMED] `pkg/plugin` and `pkg/iam` remain decoupled public infrastructure packages.
- [CONFIRMED] `internal/app` is the composition root for binding IAM authorization hooks, config-created HTTP plugin adapters and lifecycle.
- [CONFIRMED] Config-created plugins are HTTP adapters only; local plugins remain explicitly registered by code.
- [CONFIRMED] TASK-P2-004 Docker verification remains separate and open.

## Risks

- Some existing workspace changes may predate this slice; do not revert unrelated user or prior-Agent changes.
- Plugin hooks can become a hidden control plane if future work registers broad handlers without tests; keep hook points explicit and covered.
- IAM memory service is infrastructure only, not business login/RBAC; do not market it as complete authentication.
- Remote hook calls use the plugin invoke path; keep `hooks.execute` isolated from manager hook emission to avoid recursion.
- Docker build remains blocked until run in a Docker-enabled environment.

## Legal Next Step

- Task ID: TASK-P2-004
- Time Slice ID: TS-P2-004
- Status: BLOCKED
- Why: TASK-P2-004 Docker image build cannot run in the current environment because no Docker-compatible CLI is available.
- Entry condition: switch to a Docker-enabled Linux or Docker Desktop environment.
- Required command: `docker build -t go-scaffold:local .`.
- After a passing Docker build, update `STATUS.md`, `TASKS.md`, `TIME_SLICES.md`, `ACCEPTANCE.md`, `TEST_REPORT.md`, `CHANGELOG.md`, `ISSUES.md` and `AGENT_HANDOFF.md` before closing TASK-P2-004.

## Do Not Do

- Do not mark TASK-P2-004 complete until Docker build verification runs and passes.
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
3. Confirm current state is `TASK-P2-004 / TS-P2-004 / BLOCKED`.
4. Run `docker build -t go-scaffold:local .` only in a Docker-enabled environment and update project status based on the result.
