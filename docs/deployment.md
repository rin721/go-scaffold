# 部署说明

## 范围

本文记录当前项目的最小发布前检查和手动部署边界。当前切片只补 CI 质量门禁和部署说明，不执行真实部署、不推送镜像、不连接服务器、不处理生产密钥。

## 发布前检查

本地发布前至少执行：

```bash
go test ./... -count=1
go build -o ./bin/go-scaffold-server ./cmd/server
git diff --check
```

CI 中会报告 Go 格式漂移，并强制执行全量测试、server 构建和空白检查。当前仓库存在历史 gofmt 漂移，硬门禁需要单独任务收敛。

## 配置边界

- 默认配置文件：`configs/config.yaml`。
- 示例配置文件：`configs/config.example.yaml`。
- 命令行配置参数：`--config=<path>`。
- 环境变量配置路径：`REI_CONFIG_PATH`。
- 业务环境变量示例见 `.env.example`。
- 远程部署变量示例见 `.env.deploy.example`。

生产环境不要提交真实 `.env`。数据库密码、Redis 密码和其他敏感值应由运行环境或密钥管理服务注入。

## 远程部署变量

远程部署使用 `.env.deploy.example` 作为模板。真实文件建议命名为 `.env.deploy`，并只保存在本机、CI secret 或受控部署环境中。

`.env.deploy` 当前只用于记录部署目标和运行参数，不应存放 SSH 私钥、镜像仓库 token、数据库密码或 Redis 密码。SSH 私钥和镜像仓库 token 后续应通过 GitHub Secrets 注入。

推荐默认远程部署形态：

- 镜像仓库：GHCR。
- 触发方式：手动触发。
- 发布环境：`staging`。
- 远程方式：SSH 到 Linux 服务器。
- 运行方式：Docker Compose 拉取镜像并重启服务。

后续若新增自动部署 workflow，只能读取变量名和 GitHub Secrets，不能把真实值写入仓库。

## 远程部署 workflow

`.github/workflows/deploy-remote.yml` 提供手动触发的 staging 远程部署 workflow。它不会在 push 或 pull request 时自动运行，也不会在仓库中保存真实 `.env.deploy`、SSH 私钥或 token。

### GitHub Secrets

在 GitHub 仓库或 `staging` Environment 中配置：

| Secret | 必须 | 说明 |
|---|---|---|
| `DEPLOY_ENV_FILE` | 是 | `.env.deploy` 的完整内容，按 `.env.deploy.example` 填写真实值 |
| `DEPLOY_SSH_KEY` | 是 | 能以 `DEPLOY_USER` 登录远程主机的 SSH 私钥 |
| `DEPLOY_SSH_KNOWN_HOSTS` | 否 | 推荐提供远程主机 known_hosts；缺省时 workflow 会用 `ssh-keyscan` 生成 |
| `GHCR_USERNAME` | 否 | 私有 GHCR 镜像需要；公开镜像可不填 |
| `GHCR_TOKEN` | 否 | 私有 GHCR 镜像需要；只授予拉取镜像所需权限 |

`DEPLOY_ENV_FILE` 中的 `DEPLOY_ENV` 当前必须为 `staging`。生产部署仍需单独确认和新的时间切片。

### 远程主机前置条件

- Linux 主机允许 `DEPLOY_USER` 通过 SSH 登录。
- 远程主机已安装 Docker，并安装 `docker compose` 插件或 `docker-compose`。
- `DEPLOY_PATH` 可由 `DEPLOY_USER` 创建或写入。
- `DEPLOY_PATH` 下存在 `DEPLOY_COMPOSE_FILE` 指向的 Compose 文件。
- Compose 文件中的服务名与 `DEPLOY_SERVICE` 一致，并能使用 `.env.deploy` 中的 `DEPLOY_IMAGE`。
- `DEPLOY_HEALTH_URL` 和 `DEPLOY_READY_URL` 应能从远程主机访问，常见写法是 `http://127.0.0.1:<port>/health` 和 `/ready`。

workflow 不构建或推送镜像；远程主机会按 `DEPLOY_IMAGE` 拉取已存在的镜像。镜像发布需要单独任务确认。

### 手动触发

1. 打开 GitHub Actions。
2. 选择 `Remote Deploy`。
3. 点击 `Run workflow`。
4. `deploy_env` 选择 `staging`。
5. `confirm` 输入 `deploy`。
6. 运行后查看 SSH、Docker Compose 和 health/ready 检查结果。

失败时不要在日志中粘贴真实密钥或 `.env.deploy` 内容。优先检查 GitHub Secrets、远程 Compose 文件、镜像是否存在、远程 Docker 权限和 health/ready 地址。

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

## 尚未实现

- 自动生产 CD。
- Dockerfile 或镜像发布。
- Kubernetes、systemd、云平台部署模板。
- 生产数据库迁移框架。
- 密钥管理集成。

以上内容必须分别确认并拆成新的任务/时间切片后才能实现。
