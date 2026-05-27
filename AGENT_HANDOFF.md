# AGENT_HANDOFF.md

## Last Updated

- Date: 2026-05-27
- Agent: Codex
- Tool: Codex Desktop

## Project Snapshot

- Project: go-scaffold
- Phase: P2 插件钩子运行时与 IAM 公共接口完成
- Module: 项目优化路线
- Current Task: NONE
- Current Time Slice: NONE
- Overall Status: COMPLETED

## What Was Done Last

- 用户要求实现 `dev.tmp/new-plugin.md` 设计，原路径 `dev.tmp/new-pllugin.md` 视为笔误；审查结论沿用 `ACCEPT_WITH_RISK`。
- TASK-P2-004 / TS-P2-004 未标记完成，Docker build 验证仍因当前环境缺少 Docker CLI 保持 `ISSUE-P2-005` 打开。
- 完成 TASK-P2-005 至 TASK-P2-010：插件钩子运行时、HTTP 远程插件传输、独立 IAM 公共接口、配置接入、app 装配、reload 和 lifecycle。
- `pkg/plugin` 保持独立基础设施包，不导入 `pkg/iam`、日志、配置或 `internal/*`；`pkg/iam` 不导入 `pkg/plugin`；IAM 权限钩子只在 `internal/app` 注册。
- 本轮未实现 JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS 传输、生产部署、镜像发布或密钥管理。

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
| Project status docs | Updated | 记录 TASK-P2-005 至 TASK-P2-010 完成，并保留 TASK-P2-004 Docker 验证阻塞 |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| User correction review | ACCEPTED_WITH_RISK |
| `gofmt -w ...` | PASS |
| `go test ./pkg/plugin/... -count=1` | PASS |
| `go test ./pkg/iam/... -count=1` | PASS |
| `go test ./internal/config ./internal/app/... -count=1` | PASS |
| `go test ./... -count=1` | PASS |
| `go build -o <temp> ./cmd/server` | PASS |
| `git diff --check` | PASS |

## Test Status

- `pkg/plugin` and `pkg/plugin/hooks`: PASS.
- `pkg/iam` and `pkg/iam/memory`: PASS.
- `internal/config` and `internal/app/...`: PASS.
- Full regression: PASS.
- Server build: PASS.
- Diff whitespace check: PASS.
- Docker image build: still PENDING_VERIFICATION for TASK-P2-004 because Docker CLI/daemon is unavailable in the current environment.

## Current Blockers

- No blocker for TASK-P2-005 至 TASK-P2-010.
- `ISSUE-P2-005` remains open for TASK-P2-004: run `docker build -t go-scaffold:local .` in a Docker-enabled environment before closing that older deployment task.

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
- Docker build remains unverified until run in a Docker-enabled environment.

## Legal Next Step

- Task ID: NONE
- Time Slice ID: NONE
- Status: COMPLETED
- Why: TASK-P2-005 至 TASK-P2-010 are implemented and verified.
- Outstanding external verification: TASK-P2-004 / ISSUE-P2-005 still requires `docker build -t go-scaffold:local .` in a Docker-enabled environment.
- Any new feature work must be confirmed by the user and written as a new task/time slice before implementation.

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
3. Confirm current state is `NONE / COMPLETED` for the plugin/IAM mainline.
4. If the user wants to close TASK-P2-004, run `docker build -t go-scaffold:local .` only in a Docker-enabled environment and update project status based on the result.
