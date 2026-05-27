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

- 用户反馈远端 Docker build 在 `RUN go mod download` 阶段很慢并曾因访问 `proxy.golang.org` 超时失败。
- 按协议读取必读项目文档并确认当前唯一合法任务仍为 TASK-P2-004 / TS-P2-004。
- 检查 Dockerfile 后确认旧文件未声明 `GOPROXY` build arg，用户传入的 `--build-arg GOPROXY=...` 不会被 `go mod download` 使用。
- 已更新 Dockerfile：新增 `GOPROXY` / `GOSUMDB` build arg，并为 `go mod download`、`go build` 增加 BuildKit cache mount。
- TASK-P2-004 / TS-P2-004 继续保持 `BLOCKED`，`ISSUE-P2-005` 保持打开，待 Docker 环境用更新后的 Dockerfile 重跑。
- TASK-P2-005 至 TASK-P2-010 插件钩子运行时、HTTP 远程插件传输、独立 IAM 公共接口、配置接入、app 装配、reload 和 lifecycle 保持 `COMPLETED`。
- 本轮未触发 workflow、未连接远程服务器、未推送镜像、未执行真实 production、未写入真实密钥。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `Dockerfile` | Updated | 支持 `GOPROXY` / `GOSUMDB` build arg，并缓存 Go module/build 目录 |
| `docs/deployment.md` | Updated | 增加带 `GOPROXY` 的 Docker 构建示例 |
| `STATUS.md` | Updated | 记录远端 Go 代理超时与 Dockerfile 修补 |
| `TASKS.md` | Updated | 记录 TASK-P2-004 仍阻塞，待代理参数修补后重跑 |
| `TIME_SLICES.md` | Updated | 同步 TS-P2-004 阻塞原因和重跑命令 |
| `ACCEPTANCE.md` | Updated | 同步 ACC-P2-026 的阻塞说明 |
| `TEST_REPORT.md` | Updated | 记录当前最新验证为 Docker build 代理诊断 |
| `CHANGELOG.md` | Updated | 新增本轮 blocked recheck 变更记录 |
| `ISSUES.md` | Updated | 追加 `ISSUE-P2-005` 的远端构建失败与重跑命令 |
| `AGENT_HANDOFF.md` | Updated | 交接说明指向当前合法阻塞任务 |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| Dockerfile / 文档检查 | PASS |
| 用户远端 Docker build 输出审查 | FAIL_REMOTE_NETWORK：`go mod download` 访问 Go 代理超时 |
| `docker version` | FAIL_ENV |
| `Get-Command docker,podman,nerdctl,docker.exe -ErrorAction SilentlyContinue` | NOT_AVAILABLE |
| `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` | NOT_RUN_LOCAL，本机无 Docker CLI |
| Go tests | NOT_RUN，本轮未修改 Go 代码 |
| `git diff --check` | PASS，仅有 Windows LF/CRLF 提示 |

## Test Status

- Docker image build: BLOCKED for TASK-P2-004. 本机无 Docker CLI；远端补跑曾因 Go 代理网络超时失败，Dockerfile 已修补后待重跑。
- `pkg/plugin` and `pkg/plugin/hooks`: PASS from the previous completion audit.
- `pkg/iam` and `pkg/iam/memory`: PASS from the previous completion audit.
- `internal/config` and `internal/app/...`: PASS from the previous completion audit.
- Full regression and server build: PASS from the previous completion audit.
- Diff whitespace check: PASS for this documentation-only update; only Windows LF/CRLF warnings were printed.

## Current Blockers

- `ISSUE-P2-005` remains open for TASK-P2-004: run `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` with the updated Dockerfile in a Docker-enabled environment before closing that deployment task.

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
- Why: TASK-P2-004 Docker image build cannot run in the current local environment because no Docker-compatible CLI is available; the user's remote build reached `go mod download` but failed on Go proxy network timeout before this Dockerfile proxy-arg fix was available.
- Entry condition: switch to a Docker-enabled Linux or Docker Desktop environment.
- Required command: `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .`.
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
4. Run `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` only in a Docker-enabled environment and update project status based on the result.
