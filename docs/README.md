# Go-Scaffold

[![CI](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml/badge.svg)](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](../LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/rin721/go-scaffold)

`go-scaffold` 是一个可运行的 Go 后端服务脚手架。当前保留 HTTP 服务、配置加载、结构化日志、数据库访问、Demo Todo CRUD、存储辅助能力、SQL 生成、Docker 构建、CI 检查、部署示例和 AI 运行时文档。

原本内置的本地用户管理栈已经移除：不再提供 IAM 服务、认证令牌服务、RBAC 适配器、用户模块、用户表结构或用户管理 HTTP API。后续身份认证或访问控制能力需要重新从需求、架构和验收标准开始设计。

<p align="center">
  <img src="../configs/logo.png" alt="go-scaffold logo" width="180">
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
| 工程文档 | `docs/README.md`, `docs` |
| AI 运行时 | `AGENTS.md`, `docs/ai` |

## 工程文档索引

`docs/README.md` 是当前工程文档入口。`docs/ai` 是独立的 AI 运行时状态树，用于保存任务状态、证据、决策、知识和交接信息；普通工程阅读优先从下面的结构化文档开始。

### 推荐阅读顺序

1. [项目概述](overview/project.md)
2. [目录地图](structure/directory-map.md)
3. [配置说明](environment/configuration.md)
4. [分层架构](architecture/layers.md)
5. [启动流程](runtime/startup-flow.md)
6. [HTTP 流程](runtime/http-flow.md)
7. [配置流程](runtime/config-flow.md)
8. [错误流程](runtime/error-flow.md)
9. [Demo 模块](modules/demo.md)
10. [测试矩阵](testing/test-matrix.md)
11. [Docker 和 CI](build/docker-and-ci.md)
12. [部署说明](release/deployment.md)
13. [维护指南](maintenance/maintenance-guide.md)
14. [AI 运行时状态](ai-agent/runtime-state.md)
15. [已知缺口](backlog/known-gaps.md)

### 文档地图

| 分区 | 内容概述 |
| --- | --- |
| [overview](overview/project.md) | 当前能力、非目标和运行时默认假设 |
| [structure](structure/directory-map.md) | 目录职责和依赖方向 |
| [environment](environment/configuration.md) | 配置文件、环境变量、`.env` 和生产示例 |
| [architecture](architecture/layers.md) | 应用分层和装配方式 |
| [runtime](runtime/startup-flow.md) | 启动、HTTP、配置重载、状态和错误流程 |
| [modules](modules/demo.md) | Demo 模块说明 |
| [workflows](workflows/db-cli.md) | DB CLI 和运维型命令 |
| [testing](testing/test-matrix.md) | 测试归属和验证命令 |
| [build](build/docker-and-ci.md) | CI、本地构建、Docker 构建和质量门禁 |
| [release](release/deployment.md) | 生产配置、部署脚本和发布检查 |
| [extension](extension/adding-modules.md) | 新增模块、配置和 API 的方式 |
| [maintenance](maintenance/maintenance-guide.md) | 长期维护工作流 |
| [ai-agent](ai-agent/runtime-state.md) | `AGENTS.md` 和 `docs/ai` 运行时说明 |
| [backlog](backlog/known-gaps.md) | 当前已知实现和文档缺口 |

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

- `environment/configuration.md`
- `../.env.example`
- `../deploy/config.production.example.yaml`

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

本项目使用 [MIT License](../LICENSE)。
