# Go-Scaffold

[![CI](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml/badge.svg)](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/rin721/go-scaffold)

`go-scaffold` 是一个可直接运行的 Go 后端服务脚手架，适合用作服务端项目的工程基线。当前仓库包含 HTTP 服务、配置加载、结构化日志、数据库访问、Demo CRUD、本地用户认证、RBAC、存储工具、SQL 生成、Docker 构建、CI 检查、部署示例和 AI 运行态文档。

<p align="center">
  <img src="./configs/logo.png" alt="go-scaffold logo" width="180">
</p>

## 项目亮点

- 提供可运行的服务入口，支持平滑启动和优雅关闭。
- 使用清晰的模块分层：`handler -> service -> repository -> model`。
- 本地默认配置开箱即用：SQLite 数据库、Redis 关闭、Demo 模块启用。
- 提供 Docker、Compose、环境变量、远程部署 workflow 和发布脚本示例。
- 内置认证、RBAC、数据库、缓存、国际化、日志、存储、SQL 生成和 HTTP 服务等基础设施包。
- `docs/ai` 保存 AI 协作运行态，不依赖原始提示词或聊天记录作为恢复路径。

## 技术栈

| 领域 | 技术 |
| --- | --- |
| 语言与运行时 | Go 1.24.6 |
| HTTP | Gin, gin-contrib/cors |
| CLI | 本地 `pkg/cli` 命令框架 |
| 配置 | Viper, godotenv, YAML, 环境变量覆盖 |
| 日志 | Zap, lumberjack |
| 数据库 | GORM，支持 SQLite、MySQL、PostgreSQL |
| 缓存 | go-redis，可在本地禁用 |
| 认证 | JWT v5, bcrypt |
| 权限 | Casbin RBAC |
| 国际化 | go-i18n，内置 `zh-CN` 和 `en-US` |
| 存储 | afero, mimetype, imaging |
| SQL 与代码生成 | 本地 `pkg/sqlgen`, Jennifer |
| 后台任务 | ants goroutine pool manager |
| 测试 | Go test, miniredis |
| CI 与交付 | GitHub Actions, Docker, Docker Compose 示例 |

## 快速启动

使用默认本地配置启动服务：

```bash
go run ./cmd/main server
```

默认配置使用本地 SQLite `./data/app.db`，关闭 Redis，启用 demo 模块，并将 HTTP 服务绑定到 `127.0.0.1:9999`。

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

运行完整测试：

```bash
go test ./... -count=1
```

构建服务二进制：

```bash
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
```

构建 Docker 镜像：

```bash
docker build -t go-scaffold:local .
```

## 主要入口

| 范围 | 路径 |
| --- | --- |
| CLI 入口 | `cmd/main` |
| 应用装配 | `internal/app` |
| 配置 | `internal/config`, `configs` |
| HTTP 传输层 | `internal/transport/http` |
| Demo 模块 | `internal/modules/demo` |
| 用户、认证、RBAC | `internal/modules/user`, `pkg/auth`, `pkg/rbac` |
| 基础设施包 | `pkg/database`, `pkg/cache`, `pkg/logger`, `pkg/httpserver`, `pkg/storage`, `pkg/sqlgen` |
| 共享响应和错误类型 | `types` |
| Docker 与部署 | `Dockerfile`, `deploy`, `deploy.sh`, `script/install.sh` |
| 人类工程文档 | `docs/index.md` |
| AI 运行态 | `AGENTS.md`, `docs/ai` |

## API 范围

| 路由 | 用途 |
| --- | --- |
| `GET /health` | 进程存活检查 |
| `GET /ready` | 包含数据库 ping 的就绪检查 |
| `POST /api/v1/demo/todos` | 创建 demo Todo |
| `GET /api/v1/demo/todos` | 查询 demo Todo 列表 |
| `GET /api/v1/demo/todos/:id` | 读取单个 demo Todo |
| `PUT /api/v1/demo/todos/:id` | 更新 demo Todo |
| `DELETE /api/v1/demo/todos/:id` | 删除 demo Todo |
| `POST /api/v1/auth/register` | 在允许公开注册时创建本地用户 |
| `POST /api/v1/auth/login` | 登录并签发 bearer token |
| `GET /api/v1/auth/me` | 读取当前登录主体 |
| `/api/v1/users`, `/api/v1/roles`, `/api/v1/permissions` | 需要认证和权限的用户/RBAC 管理 |

## 配置说明

本地配置从 `configs/config.yaml` 或 `configs/config.example.yaml` 开始。运行时也可以通过环境变量和 `.env` 文件覆盖配置。生产环境最重要的覆盖项是：

```bash
RIN_APP_AUTH_TOKEN_SECRET=<至少-32-byte-的密钥>
```

推荐阅读：

- `docs/environment/configuration.md`
- `.env.example`
- `deploy/config.production.example.yaml`

## 数据库 CLI

`db` 命令可以预览或应用生成的 SQL，也可以通过应用服务层执行 demo Todo 操作。

```bash
go run ./cmd/main db --operation=schema
go run ./cmd/main db --operation=schema --apply
go run ./cmd/main db --operation=todo-list
```

已移除的 `initdb` 命令和 InitDB 配置段不得在没有新确认任务的情况下恢复。

## 工程化工作流

```bash
gofmt -w ./cmd ./internal ./pkg ./types
go test ./... -count=1 -mod=readonly
go build -mod=readonly -o ./tmp/go-scaffold-server ./cmd/main
docker build -t go-scaffold:ci .
```

CI 会在推送到 `main`、`master` 或创建 Pull Request 时运行格式漂移检查、Go 测试、服务构建、Docker 镜像构建和空白检查。

## 部署入口

| 目标 | 入口 |
| --- | --- |
| 本地二进制 | `go build ... ./cmd/main` |
| 本地容器 | `Dockerfile` |
| 生产 Compose 示例 | `deploy/docker-compose.production.example.yml` |
| 生产配置示例 | `deploy/config.production.example.yaml` |
| Shell 部署助手 | `deploy.sh` |
| 安装助手 | `script/install.sh` |
| 远程部署 workflow | `.github/workflows/deploy-remote.yml` |

## 文档

从 `docs/index.md` 开始阅读。文档按当前代码形态组织，覆盖项目概览、目录地图、配置、架构、运行流程、模块、工作流、测试、构建、发布、维护、AI 协作和已知缺口。

也可以通过 README 顶部的 DeepWiki 徽章访问项目知识库：

```text
https://deepwiki.com/rin721/go-scaffold
```

## 生产注意事项

生产配置必须向 `RIN_APP_AUTH_TOKEN_SECRET` 注入至少 32 bytes 的密钥。本地开发在未显式配置 token secret 时可以生成进程内随机密钥，但服务重启后旧 token 会失效。

生产示例默认关闭 demo 模块。除非有明确确认，不要在生产环境暴露 demo 路由或隐式创建 demo schema。

## 开源协议

本项目基于 [MIT License](LICENSE) 开源。
