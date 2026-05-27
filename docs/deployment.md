# 部署说明

## 范围

本文记录当前项目的最小发布前检查、Linux Docker 制品和手动远程部署边界。当前仓库提供 CI 质量门禁、`Dockerfile`、production Compose 示例、统一 `deploy.sh` 部署入口和手动远程部署 workflow；本地会话不执行真实部署、不推送镜像、不连接服务器、不处理生产密钥。

## 发布前检查

本地发布前至少执行：

```bash
go test ./... -count=1
go build -o ./bin/go-scaffold-server ./cmd/server
git diff --check
```

CI 中会报告 Go 格式漂移，并强制执行全量测试、server 构建和空白检查。当前仓库存在历史 gofmt 漂移，硬门禁需要单独任务收敛。

## Docker 镜像

仓库根目录提供 `Dockerfile`。镜像用于 Linux 运行环境，构建 server 二进制并以非 root 用户运行。

```bash
docker build -t go-scaffold:local .
```

当前会话环境未安装 Docker CLI，因此上述镜像构建命令尚未在本机完成验证；需在安装 Docker 的 Linux 或 Docker Desktop 环境补跑。

镜像内置 `deploy/config.production.example.yaml` 作为默认 `/app/configs/config.yaml`，其中 server 绑定 `0.0.0.0:9999`。真实 production 应在远程主机用只读挂载覆盖 `/app/configs/config.yaml`，不要把真实数据库密码、Redis 密码或 token 写入镜像。

## 统一脚本部署

clone 后在仓库根目录执行：

```bash
git clone <repo-address>
cd <repo-address>
bash deploy.sh \
  --docker y \
  --env production \
  --image go-scaffold:local \
  --port 9999 \
  --db-driver sqlite \
  --db-name ./data/app.db \
  --confirm
```

直接下载脚本执行：

```bash
curl -fsSL -o deploy.sh \
  https://raw-githubusercontent-com-gh.helloworlds.eu.org/rin721/go-scaffold/main/script/install.sh
bash deploy.sh \
  --docker y \
  --repo https://github.com/rin721/go-scaffold.git \
  --ref main \
  --env production \
  --image go-scaffold:local \
  --confirm
```

`--docker y` 表示使用 Docker Compose 部署。脚本会准备运行目录、复制 production Compose 示例和默认配置、按显式参数导出容器环境变量，然后执行 Docker build 或 pull、Compose up、health/ready 检查。当前不实现非 Docker 部署。

敏感参数示例：

```bash
bash deploy.sh \
  --docker y \
  --image ghcr.io/OWNER/go-scaffold:TAG \
  --pull y \
  --db-driver postgres \
  --db-host db.internal \
  --db-port 5432 \
  --db-user app \
  --db-password "***" \
  --db-name appdb \
  --redis-enabled true \
  --redis-host redis.internal \
  --redis-password "***" \
  --confirm
```

密码、token 和 secret 类参数可能进入 shell history 或进程列表。推荐在受控 shell、CI Secret masking 或主机密钥管理器中使用；脚本不会打印这些值。

## Docker Compose production 示例

production Compose 示例位于 `deploy/docker-compose.production.example.yml`。脚本会复制它为运行目录下的 `docker-compose.yml`，并通过当前 shell 环境变量给 Compose 和容器传值，不依赖单独的部署 env 文件。

如果手工准备运行目录，可执行：

```bash
mkdir -p configs data logs
cp deploy/config.production.example.yaml configs/config.yaml
cp deploy/docker-compose.production.example.yml docker-compose.yml
sudo chown -R 10001:10001 data logs
```

Compose 示例读取 `DEPLOY_IMAGE`、`APP_PORT`、`DEPLOY_CONTAINER_NAME` 以及 DB、Redis、Server、Logger、I18n、Storage、CORS 显式环境变量，并对 `/health` 配置容器内 healthcheck。

## 配置边界

- 默认配置文件：`configs/config.yaml`。
- 示例配置文件：`configs/config.example.yaml`。
- 命令行配置参数：`--config=<path>`。
- 环境变量配置路径：`REI_CONFIG_PATH`。
- 本地开发环境变量示例见 `.env.example`。
- 部署脚本参数直接映射到当前已实现的环境变量覆盖策略。

生产环境不要提交真实 `.env`。数据库密码、Redis 密码和其他敏感值应由运行环境或密钥管理服务注入。

## 远程部署 workflow

`.github/workflows/deploy-remote.yml` 提供手动触发的远程部署 workflow。它不会在 push 或 pull request 时自动运行，也不会在仓库中保存真实 SSH 私钥、token、密码或服务器值。

workflow 支持 `staging` 和 `production` 两个环境。production 必须在 GitHub Environments 中配置 `production` 环境，并建议启用 required reviewers、分支限制和最小权限 Secrets。

### GitHub Variables 与 Secrets

在 GitHub 仓库、`staging` Environment 或 `production` Environment 中配置显式变量。

常用 Variables：

| Variable | 必须 | 说明 |
|---|---|---|
| `DEPLOY_HOST` | 是 | 远程 Linux 主机 |
| `DEPLOY_PORT` | 否 | SSH 端口，默认 `22` |
| `DEPLOY_USER` | 是 | SSH 用户 |
| `DEPLOY_PATH` | 否 | 运行目录，默认 `/opt/go-scaffold` |
| `DEPLOY_REPO` | 否 | 仓库地址，默认 `https://github.com/rin721/go-scaffold.git` |
| `DEPLOY_REF` | 否 | Git ref，默认当前分支或 `main` |
| `DEPLOY_IMAGE` | 否 | 镜像名，默认 `go-scaffold:local` |
| `DEPLOY_BUILD` | 否 | 是否构建镜像，默认 `y` |
| `DEPLOY_PULL` | 否 | 是否拉取镜像，默认 `n` |
| `APP_PORT` | 否 | 主机端口，默认 `9999` |
| `DEPLOY_CONTAINER_NAME` | 否 | 容器名，默认 `go-scaffold` |
| `DB_DRIVER`、`DB_HOST`、`DB_PORT`、`DB_USER`、`DB_NAME` | 否 | 数据库非密钥参数 |
| `REDIS_ENABLED`、`REDIS_HOST`、`REDIS_PORT`、`REDIS_DB` | 否 | Redis 非密钥参数 |

常用 Secrets：

| Secret | 必须 | 说明 |
|---|---|---|
| `DEPLOY_SSH_KEY` | 是 | 能以 `DEPLOY_USER` 登录远程主机的 SSH 私钥 |
| `DEPLOY_SSH_KNOWN_HOSTS` | 否 | 推荐提供远程主机 known_hosts；缺省时 workflow 会用 `ssh-keyscan` 生成 |
| `DB_PASSWORD` | 否 | 数据库密码 |
| `REDIS_PASSWORD` | 否 | Redis 密码 |
| `GHCR_USERNAME` | 否 | 私有 GHCR 镜像需要；公开镜像可不填 |
| `GHCR_TOKEN` | 否 | 私有 GHCR 镜像需要；只授予拉取镜像所需权限 |

### 远程主机前置条件

- Linux 主机允许 `DEPLOY_USER` 通过 SSH 登录。
- 远程主机已安装 Git、Docker，并安装 `docker compose` 插件或 `docker-compose`。
- `DEPLOY_PATH` 可由 `DEPLOY_USER` 创建或写入。
- `DEPLOY_PATH/data` 和 `DEPLOY_PATH/logs` 对容器用户 `10001:10001` 可写，或远程用户具有 passwordless sudo 来修正权限。

workflow 会通过 SSH 在远程主机执行 `script/install.sh`，由该脚本 clone 仓库并调用根 `deploy.sh`。镜像发布流水线仍需单独任务确认。

### 手动触发

1. 打开 GitHub Actions。
2. 选择 `Remote Deploy`。
3. 点击 `Run workflow`。
4. `deploy_env` 选择 `staging` 或 `production`。
5. `confirm` 输入与环境一致的确认词：`deploy-staging` 或 `deploy-production`。
6. 运行后查看 SSH、Docker Compose 和 health/ready 检查结果。

production 运行前必须确认 GitHub Environment 审批已生效、显式 Variables/Secrets 指向 production 主机、`DEPLOY_IMAGE` 是要发布的不可变镜像标签，并且远程主机已有可回滚的上一版本。

失败时不要在日志中粘贴真实密钥或命令行完整敏感参数。优先检查 GitHub Variables、Secrets、远程 Docker 权限、镜像是否存在和 health/ready 地址。

## 本地运行

```bash
go run ./cmd/server server --config=configs/config.yaml
```

Windows PowerShell 可使用：

```powershell
$env:REI_CONFIG_PATH = "configs/config.yaml"
go run ./cmd/server server
```

服务启动后检查：

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

## 初始化数据库

```bash
go run ./cmd/server initdb --config=configs/config.yaml
```

`initdb` 当前用于 demo schema bootstrap。生产迁移框架尚未实现，生产数据库结构变更必须单独确认，不能依赖运行期隐式 `AutoMigrate`。

## 手动发布步骤

1. 从干净工作区构建二进制。
2. 在目标环境准备配置文件或环境变量。
3. 先在目标环境执行只读健康检查所需依赖验证。
4. 如需要初始化 demo schema，显式执行 `initdb`。
5. 启动 server。
6. 检查 `/health` 和 `/ready`。
7. 保留上一版本二进制和配置以便回滚。

## production 回滚边界

当前 workflow 不自动执行数据库迁移，也不自动回滚。production 回滚建议固定为镜像标签级回滚：

1. 将 GitHub Environment 中的 `DEPLOY_IMAGE` 改回上一版本镜像标签。
2. 手动运行 `Remote Deploy`，选择 `production`，输入 `deploy-production`。
3. 检查 `/health` 和 `/ready`。
4. 如涉及数据库 schema 变更，必须使用单独确认的生产迁移流程，不要依赖当前 workflow 自动处理。

## 尚未实现

- 自动 production CD。
- 镜像发布 workflow。
- Kubernetes、systemd、云平台部署模板。
- 生产数据库迁移框架。
- 密钥管理集成。

以上内容必须分别确认并拆成新的任务/时间切片后才能实现。
