# Go-Scaffold

[![CI](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml/badge.svg)](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/rin721/go-scaffold)

`go-scaffold` 是一个可运行的 Go 后端服务脚手架。当前保留 HTTP 服务、配置加载、结构化日志、数据库访问、Demo Todo CRUD、存储辅助能力、SQL 生成、Docker 构建、CI 检查、部署示例和 AI 运行时文档。

原本内置的本地用户管理栈已经移除：不再提供 IAM 服务、认证令牌服务、RBAC 适配器、用户模块、用户表结构或用户管理 HTTP API。后续身份认证或访问控制能力需要重新从需求、架构和验收标准开始设计。

<p align="center">
  <img src="./configs/logo.png" alt="go-scaffold logo" width="180">
</p>

## 亮点

- 可直接启动的服务入口，支持优雅启动和关闭。
- Demo 模块采用 `handler -> service -> repository -> model` 分层，适合作为新增业务模块参考。
- 本地默认配置使用 SQLite `./data/app.db`，Redis 默认关闭，Demo 模块默认开启，HTTP 监听 `127.0.0.1:9999`。
- 提供 Docker、Compose、环境变量、部署脚本和 CI 示例。
- `docs/ai` 保存可恢复的 AI 协作状态，不依赖原始提示词或聊天历史。

## 技术栈

| 范围 | 技术 |
| --- | --- |
| 运行时 | Go 1.24.6 |
| HTTP | Gin, gin-contrib/cors |
| CLI | 本地 `pkg/cli` 命令框架 |
| 配置 | Viper, godotenv, YAML, 环境变量覆盖 |
| 日志 | Zap, lumberjack |
| 数据库 | GORM，支持 SQLite、MySQL、PostgreSQL |
| 缓存 | go-redis，可选启用 |
| 国际化 | go-i18n，包含 `zh-CN` 和 `en-US` 示例 |
| 存储 | afero, mimetype, imaging |
| SQL/代码生成 | 本地 `pkg/sqlgen`, Jennifer |
| 后台任务 | ants 协程池管理 |
| 测试 | Go test, miniredis |
| CI 和交付 | GitHub Actions, Docker, Docker Compose 示例 |

## 快速启动

```bash
go run ./cmd/main server
```

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

```bash
go test ./... -count=1
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
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
| 基础设施包 | `pkg/database`, `pkg/cache`, `pkg/logger`, `pkg/httpserver`, `pkg/storage`, `pkg/sqlgen` |
| 共享响应和错误类型 | `types` |
| Docker 和部署 | `Dockerfile`, `deploy`, `deploy.sh`, `script/install.sh` |
| 工程文档 | `docs/index.md` |
| AI 运行时 | `AGENTS.md`, `docs/ai` |

## API 范围

| 路由 | 用途 |
| --- | --- |
| `GET /health` | 进程存活检查 |
| `GET /ready` | 包含数据库 ping 的就绪检查 |
| `POST /api/v1/demo/todos` | 创建 Demo Todo |
| `GET /api/v1/demo/todos` | 列出 Demo Todo |
| `GET /api/v1/demo/todos/:id` | 读取单个 Demo Todo |
| `PUT /api/v1/demo/todos/:id` | 更新单个 Demo Todo |
| `DELETE /api/v1/demo/todos/:id` | 删除单个 Demo Todo |

## 配置

本地配置从 `configs/config.yaml` 或 `configs/config.example.yaml` 开始。运行时也可以通过环境变量和 `.env` 覆盖配置值。

推荐参考：

- `docs/environment/configuration.md`
- `.env.example`
- `deploy/config.production.example.yaml`

## 数据库 CLI

```bash
go run ./cmd/main db --operation=schema
go run ./cmd/main db --operation=schema --apply
go run ./cmd/main db --operation=todo-list
```

已移除的 `initdb` 命令和 InitDB 配置块不要在没有新确认任务的情况下恢复。

## 工程工作流

```bash
gofmt -w ./cmd ./internal ./pkg ./types
go test ./... -count=1 -mod=readonly
go build -mod=readonly -o ./tmp/go-scaffold-server ./cmd/main
docker build -t go-scaffold:ci .
```

## 生产提示

生产示例默认关闭 Demo 模块。除非明确需要，不要在生产或类生产环境暴露 Demo 路由，也不要隐式创建 Demo 表结构。

真实部署前需要审查数据库、Redis、存储、日志、CORS、健康/就绪检查、备份和回滚策略。

## 许可证

本项目使用 [MIT License](LICENSE)。
