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

- 按 `dev.tmp/new-plugin.md` 重新审计 TASK-P2-005 至 TASK-P2-010 的当前实现。
- 补强 `pkg/plugin/hooks`：nil `HandlerFunc` 作为空处理器在注册时拒绝，直接调用也返回 `ErrNilHandler`。
- 补强 HTTP 远程插件客户端：响应超过 `maxResponseBytes` 时返回包装 `ErrInvalidResponse` 的明确错误。
- 新增测试覆盖 nil hook handler 拒绝、HTTP 响应大小限制、`after_invoke` hook 失败时返回插件响应和包装后的 hook 错误。
- TASK-P2-005 至 TASK-P2-010 插件钩子运行时、HTTP 远程插件传输、独立 IAM 公共接口、配置接入、app 装配、reload 和 lifecycle 保持 `COMPLETED`。
- TASK-P2-004 / TS-P2-004 Docker build 仍因当前环境缺少 Docker 兼容 CLI 保持 `BLOCKED`，`ISSUE-P2-005` 保持打开。
- 本轮未触发 workflow、未连接远程服务器、未推送镜像、未执行真实 production、未写入真实密钥。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `pkg/plugin/hooks/*` | Added | 独立钩子引擎、注册表、服务查找、停止语义和测试 |
| `pkg/plugin/hooks/registry.go`、`pkg/plugin/hooks/types.go`、`pkg/plugin/hooks/registry_test.go` | Updated | 拒绝 nil `HandlerFunc` 并补测试 |
| `pkg/plugin/manager.go`、`pkg/plugin/constants.go` | Updated | manager 钩子 API、标准钩子点、调用语义 |
| `pkg/plugin/http_server.go`、`pkg/plugin/remote_hook.go` | Added | HTTP 插件服务端 helper 和远程钩子适配器 |
| `pkg/plugin/http.go`、`pkg/plugin/plugin_test.go` | Updated | 明确 HTTP 响应大小超限错误，并覆盖 after hook 错误返回响应 |
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
| `gofmt -w pkg/plugin/hooks/types.go pkg/plugin/hooks/registry.go pkg/plugin/hooks/registry_test.go pkg/plugin/http.go pkg/plugin/plugin_test.go` | PASS |
| `go test ./pkg/plugin/... -count=1` | PASS |
| `go test ./pkg/iam/... -count=1` | PASS |
| `go test ./internal/config ./internal/app/... -count=1` | PASS |
| `go test ./... -count=1` | PASS |
| `go build -o <temp> ./cmd/server` | PASS |
| `git diff --check` | PASS |

## Test Status

- Docker image build: BLOCKED for TASK-P2-004 because Docker CLI/daemon is unavailable in the current environment.
- `pkg/plugin` and `pkg/plugin/hooks`: PASS after completion audit.
- `pkg/iam` and `pkg/iam/memory`: PASS after completion audit.
- `internal/config` and `internal/app/...`: PASS after completion audit.
- Full regression and server build: PASS after completion audit.
- Diff whitespace check: PASS; only Windows LF/CRLF warnings were printed.

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
