# AGENT_HANDOFF.md

## Last Updated

- Date: 2026-05-26
- Agent: Codex
- Tool: Codex Desktop

## Project Snapshot

- Project: go-scaffold
- Phase: P2 远程部署 workflow 完成
- Module: 项目优化路线
- Current Task: NONE
- Current Time Slice: NONE
- Overall Status: COMPLETED

## What Was Done Last

- 用户明确确认实现远程部署 workflow。
- 完成 TASK-P2-003 / TS-P2-003。
- 新增 `.github/workflows/deploy-remote.yml`，提供手动触发的 staging 远程部署 workflow。
- workflow 使用 `DEPLOY_ENV_FILE`、`DEPLOY_SSH_KEY`、可选 `DEPLOY_SSH_KNOWN_HOSTS`、可选 `GHCR_USERNAME` / `GHCR_TOKEN` 等 GitHub Secrets。
- workflow 校验 `.env.deploy` 必需变量，要求 `confirm=deploy`，通过 SSH/SCP 上传 `.env.deploy`，并在远程执行 Docker Compose pull/up 和 health/ready 检查。
- 更新 `.env.deploy.example`、`docs/deployment.md` 和 README 的 workflow、Secrets 和远程主机前置条件说明。
- 未修改 Go 代码、测试文件、导出业务 API、配置 schema、HTTP 路由、数据库 schema、依赖文件、真实配置、部署凭据或密钥。
- 未执行真实部署、未连接远程服务器、未推送镜像。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `.github/workflows/deploy-remote.yml` | Added | 手动 staging 远程部署 workflow |
| `.env.deploy.example` | Updated | 补充 workflow Secret 名称 |
| `docs/deployment.md`、`README.md` | Updated | 记录 workflow、Secrets、远程主机前置条件和手动触发步骤 |
| Project status docs | Updated | Mark TASK-P2-003 / TS-P2-003 completed |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| User correction review | ACCEPTED_WITH_RISK |
| Temporary Go YAML parse | PASS |
| `go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml` | PASS |
| `git diff --check` | PASS, only Windows LF/CRLF conversion warnings |

## Test Status

- Full regression: not run for this workflow/documentation slice; no Go code changed.
- Workflow static validation: PASS.
- Diff whitespace check: PASS, with LF/CRLF warnings only.
- Pending verification: none.

## Current Blockers

- None for TASK-P2-003.
- Image publishing, Dockerfile, production deployment and production migration framework remain unimplemented and require separate confirmation.

## Important Decisions

- [CONFIRMED] User wants remote deployment configured through an `.env` style file.
- [CONFIRMED] Commit only `.env.deploy.example`, never real `.env.deploy`.
- [CONFIRMED] User explicitly confirmed implementing the remote deployment workflow.
- [CONFIRMED] Workflow is staging-only, manually triggered, and requires `confirm=deploy`.
- [CONFIRMED] Workflow reads deployment configuration from GitHub Secrets, not committed real values.

## Risks

- Some existing workspace changes predate this slice; do not revert unrelated user or prior-Agent changes.
- The workflow is present but was not executed in this session; it will connect to the remote host only when manually run in GitHub Actions with real Secrets.
- Image publishing is still out of scope; the remote host pulls an existing `DEPLOY_IMAGE`.
- Production deployment is still out of scope.
- Production migration framework, auth/rbac, and plugin rpc/ws/discovery remain out of scope.
- Further README or historical document localization must be confirmed as a new task/time slice.

## Legal Next Step

- Task ID: NONE
- Time Slice ID: NONE
- Status: COMPLETED
- Why this is next: TASK-P2-003 completed the hand-triggered staging remote deployment workflow; there is no automatic next implementation task.
- If the user asks to continue, confirm whether they want image publishing/Dockerfile, production deployment, or production migration framework.

## Do Not Do

- Do not start new Go code, tests, production deployment, image publishing, Dockerfile, or migration work without a new confirmed task/time slice.
- Do not modify `cmd/**/*`, `internal/**/*`, `pkg/**/*`, `types/**/*`, `go.mod`, `go.sum`, schema, deployment config, or secrets without a new legal task.
- Do not commit, push, deploy, run irreversible migrations, or expose real `.env` values without explicit user confirmation.
- Do not revert unrelated dirty workspace changes.

## Recovery Instructions

1. Read `AGENTS.md`.
2. Read `STATUS.md`, `TASKS.md`, and `TIME_SLICES.md`.
3. Confirm current state is `NONE / COMPLETED` after TASK-P2-003.
4. If the user confirms image publishing, production deployment, Dockerfile, or production migration, create a new task/time slice before editing related files.
