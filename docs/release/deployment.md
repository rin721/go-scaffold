# 部署发布

当前部署能力是一条示例 production 路径，不是 v1 发布保证。真实环境使用前必须
复核配置、secret、数据库选择和回滚策略。

## 文件

| 路径 | 用途 |
| --- | --- |
| `Dockerfile` | 构建服务镜像 |
| `deploy/config.production.example.yaml` | production 风格应用配置 |
| `deploy/docker-compose.production.example.yml` | Compose 服务定义 |
| `deploy.sh` | Bash 部署包装器 |
| `script/install.sh` | 调用 `deploy.sh` 的远程安装入口 |
| `.github/workflows/deploy-remote.yml` | 手动 GitHub Actions 远程部署 |

## 必需密钥

生产必须设置：

```bash
RIN_APP_AUTH_TOKEN_SECRET=<at-least-32-bytes>
```

Compose 示例会强制要求它。生产不要依赖本地随机 token secret fallback。

## 手动 Docker Compose 路径

```bash
export DEPLOY_IMAGE=go-scaffold:local
export RIN_APP_AUTH_TOKEN_SECRET=replace-with-at-least-32-bytes
docker compose -f deploy/docker-compose.production.example.yml up -d
```

然后检查：

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

## deploy.sh 路径

`deploy.sh` 可以 clone/update repo、准备配置、build 或 pull image、运行 Compose，
并检查 health/readiness。破坏性或 production-like 操作需要显式确认 flag。

部署脚本应在 Linux Bash 环境执行。

## GitHub 远程 Workflow

远程 workflow 是手动触发的，需要选择环境并输入确认字符串。它通过 SSH 进入目标
主机并调用 `script/install.sh`。

命名注意：应用本身优先使用 `RIN_APP_*` 变量，而 workflow 当前读取若干未加前缀
的 GitHub Variables，例如 `DB_DRIVER`，再转换成 deploy script 参数。修改部署
配置时要保持文档和 workflow 变量一致。

## 发布检查清单

真实发布前：

1. 选择并验证生产数据库 driver；
2. 安全注入 `RIN_APP_AUTH_TOKEN_SECRET`；
3. 除非明确需要，否则关闭 demo 路由；
4. 复核 CORS origins；
5. 验证 `/health` 和 `/ready`；
6. 运行根模块测试；
7. 在干净环境构建 Docker 镜像；
8. 记录回滚和数据备份步骤；
9. 在合适的运行态 artifact 中记录部署证据。
