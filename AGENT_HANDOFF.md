# AGENT_HANDOFF.md

## Last Updated

- Date: 2026-05-26
- Agent: Codex
- Tool: Codex Desktop

## Project Snapshot

- Project: go-scaffold
- Phase: P2 Linux Docker production 部署制品与远程 Linux 脚本待验证
- Module: 项目优化路线
- Current Task: TASK-P2-004
- Current Time Slice: TS-P2-004
- Overall Status: PENDING_VERIFICATION

## What Was Done Last

- 用户要求“开始，linux、docker、production -> 部署”，并修正为“环境变量在部署脚本上动态配置”。
- 完成 TASK-P2-004 / TS-P2-004 的制品实现部分，状态保持 `PENDING_VERIFICATION`。
- 新增 `Dockerfile`，用于构建 Linux server 镜像，并在运行阶段使用非 root 用户。
- 新增 `.dockerignore`，排除 Git、真实 env、缓存、日志和非运行制品。
- 新增 `deploy/docker-compose.production.example.yml`，使用 `DEPLOY_IMAGE`、外置配置、数据和日志挂载，并包含 `/health` healthcheck。
- 新增 `deploy/config.production.example.yaml`，绑定 `0.0.0.0:9999`，不包含真实密钥。
- 新增 `deploy/remote-linux-deploy.sh`，用于在远程 Linux 主机按参数/环境变量动态生成 `DEPLOY_PATH/.env.deploy`，再执行 Docker build 或 pull、Compose up 和 health/ready 检查。
- 扩展 `.github/workflows/deploy-remote.yml`，支持 `staging` / `production` 手动环境选择，并要求 `deploy-staging` 或 `deploy-production` 确认词。
- 更新 `.env.deploy.example`、`docs/deployment.md` 和 README，记录 Linux Docker、Windows 到远程 Linux 直接部署、production Compose、GitHub Environment、Secrets、目录权限和回滚边界。
- 未修改 Go 代码、测试文件、导出业务 API、配置 schema、HTTP 路由、数据库 schema、依赖文件、真实配置、部署凭据或密钥。
- 未执行真实部署、未触发 workflow、未连接远程服务器、未推送镜像。

## Files Changed Last

| File | Change | Reason |
|---|---|---|
| `Dockerfile` | Added | Linux Docker server 镜像制品 |
| `.dockerignore` | Added | 控制 Docker build context，避免真实 env 和非运行制品进入镜像构建 |
| `deploy/docker-compose.production.example.yml` | Added | production Docker Compose 示例 |
| `deploy/config.production.example.yaml` | Added | production 配置样例，无真实密钥 |
| `deploy/remote-linux-deploy.sh` | Added | 远程 Linux 动态生成 `.env.deploy` 并执行受控 Docker Compose 部署路径 |
| `.github/workflows/deploy-remote.yml` | Updated | 支持手动 staging/production 环境选择和环境绑定确认词 |
| `.env.deploy.example` | Updated | 补充 `APP_PORT`、`DEPLOY_CONTAINER_NAME` 和 production 环境说明 |
| `README.md`、`docs/deployment.md` | Updated | 记录 Docker、Compose、GitHub Environment、Secrets、权限和回滚边界 |
| Project status docs | Updated | 标记 TASK-P2-004 / TS-P2-004 为 `PENDING_VERIFICATION` |

## Commands Run Last

| Command | Result |
|---|---|
| Required file reads | PASS |
| User correction review | ACCEPTED_WITH_RISK |
| `docker version` | FAIL_ENV，当前环境未安装 Docker CLI |
| `Get-Command podman` / `Get-Command nerdctl` / `Get-Command docker.exe` | NOT_AVAILABLE |
| `Get-Command bash` | FAIL_ENV，本机无可用 bash |
| `wsl bash -lc ... bash -n deploy/remote-linux-deploy.sh` | FAIL_ENV，WSL 未安装 Linux 发行版 |
| `go run mvdan.cc/sh/v3/cmd/shfmt@latest -ln bash -tojson` | PASS，脚本 Bash 语法解析通过 |
| Temporary Go YAML parse | PASS |
| `go run github.com/rhysd/actionlint/cmd/actionlint@latest .github/workflows/ci.yml .github/workflows/deploy-remote.yml` | PASS |
| `go test ./... -count=1` | PASS |
| `go build -o <temp> ./cmd/server` | PASS |
| `git diff --check` | PASS, only Windows LF/CRLF conversion warnings |

## Test Status

- Full regression: PASS.
- Remote Linux deployment script syntax: PASS by Bash parser; `bash -n` itself is unavailable in this Windows session.
- Server build: PASS.
- Workflow static validation: PASS.
- Diff whitespace check: PASS, with LF/CRLF warnings only.
- Docker image build: PENDING_VERIFICATION because Docker CLI/daemon is unavailable in the current environment.

## Current Blockers

- `docker build -t go-scaffold:local .` has not been run because no Docker-compatible CLI is available in this session.
- This is an environment verification gap, not a known implementation failure.
- Image publishing, real production runtime execution and production migration framework remain out of scope and require separate confirmation.

## Important Decisions

- [CONFIRMED] User wants remote deployment configured through an `.env` style file.
- [CONFIRMED] Commit only `.env.deploy.example`, never real `.env.deploy`.
- [CONFIRMED] User explicitly confirmed implementing the remote deployment workflow.
- [CONFIRMED] User explicitly requested Linux/Docker/production deployment artifacts.
- [CONFIRMED] Production workflow access must remain manual and require GitHub Environment `production` plus `deploy-production` confirmation.
- [CONFIRMED] Remote Linux direct deployment should generate `.env.deploy` dynamically from non-secret runtime variables in `deploy/remote-linux-deploy.sh`.
- [CONFIRMED] Current slice does not trigger workflow, connect to any remote host, push images, execute production deployment or run migrations.

## Risks

- Some existing workspace changes may predate this slice; do not revert unrelated user or prior-Agent changes.
- Production-named artifacts can be mistaken for a completed production rollout. They are examples and controlled workflow gates only.
- The dynamic env script must not be treated as a secret manager; database/Redis credentials remain in remote config or a separate secret system.
- Docker build remains unverified until run in a Docker-enabled environment.
- Workflow is present but was not executed in this session; it will connect to the remote host only when manually run in GitHub Actions with real Secrets.
- Image publishing is still out of scope; the remote host pulls an existing `DEPLOY_IMAGE`.
- Production migration framework, auth/rbac, and plugin rpc/ws/discovery remain out of scope.

## Legal Next Step

- Task ID: TASK-P2-004
- Time Slice ID: TS-P2-004
- Status: PENDING_VERIFICATION
- Why this is next: the only unmet acceptance item is Docker image build verification.
- Required action: in an environment with Docker CLI/daemon, run `docker build -t go-scaffold:local .`.
- Completion rule: if Docker build passes, update `STATUS.md`, `TASKS.md`, `TIME_SLICES.md`, `ACCEPTANCE.md`, `TEST_REPORT.md`, `CHANGELOG.md`, `ISSUES.md` and this handoff to close TASK-P2-004. If it fails, record the output in `ISSUES.md` and enter failure repair within the TASK-P2-004 scope.

## Do Not Do

- Do not trigger GitHub workflow from this session.
- Do not connect to remote servers.
- Do not push images.
- Do not execute staging or production deployment.
- Do not run production migrations or irreversible database changes.
- Do not write, print or invent real `.env`, SSH key, token, password or production host values.
- Do not modify Go source, tests, `go.mod`, `go.sum`, schema, business API or HTTP routes unless a new legal task explicitly authorizes it.
- Do not commit, push or revert unrelated dirty workspace changes.

## Recovery Instructions

1. Read `AGENTS.md`.
2. Read `STATUS.md`, `TASKS.md`, and `TIME_SLICES.md`.
3. Confirm current state is `TASK-P2-004 / TS-P2-004 / PENDING_VERIFICATION`.
4. Run `docker build -t go-scaffold:local .` only in a Docker-enabled environment.
5. Update status and reports according to the build result.
