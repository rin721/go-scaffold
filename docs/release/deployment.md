# 部署说明

当前部署能力是生产风格示例，不是 v1 发布保证。真实环境使用前需要审查配置、密钥、数据库选择和回滚策略。

## 相关文件

| 路径 | 用途 |
| --- | --- |
| `Dockerfile` | 构建服务镜像 |
| `deploy/config.production.example.yaml` | 生产风格应用配置 |
| `deploy/docker-compose.production.example.yml` | Compose 服务定义 |
| `deploy.sh` | Bash 部署包装脚本 |
| `script/install.sh` | 调用 `deploy.sh` 的远程安装入口 |
| `.github/workflows/deploy-remote.yml` | 手动触发的 GitHub Actions 远程部署 |

## 手动 Docker Compose 路径

```bash
export DEPLOY_IMAGE=go-scaffold:local
docker compose -f deploy/docker-compose.production.example.yml up -d
```

然后检查：

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

## deploy.sh 路径

`deploy.sh` 可以克隆仓库或使用本地仓库、准备配置、构建或拉取镜像、运行 Compose，并检查健康/就绪状态。破坏性或类生产操作必须显式传入 `--confirm`。

该脚本应在 Linux Bash 环境运行。

## 发布清单

1. 选择并验证生产数据库驱动。
2. 除非明确需要，否则关闭 Demo 路由。
3. 审查 CORS origins 和 headers。
4. 验证 `/health` 和 `/ready`。
5. 运行根模块测试。
6. 在干净环境构建 Docker 镜像。
7. 记录回滚和备份步骤。
8. 如果属于托管任务，在对应运行时制品中记录部署证据。
