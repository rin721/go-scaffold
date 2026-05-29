# go-scaffold

`go-scaffold` 是一个 Go 后端服务脚手架。当前仓库已经包含可运行的 HTTP
服务、配置加载、日志、数据库访问、Demo Todo API、本地用户认证、RBAC、插件注册、存储工具、SQL
生成、Docker 构建、CI 检查和部署示例。

项目尚未达到 v1 发布条件。Docker、CI、部署脚本和运行文档都是当前阶段的工程基线，不等同于生产发布承诺。

## 快速启动

```bash
go run ./cmd/main server
```

默认配置使用本地 SQLite `./data/app.db`，关闭 Redis，启用 demo 模块，并把 HTTP 服务绑定到
`127.0.0.1:9999`。

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

运行测试：

```bash
go test ./... -count=1
cd remote_plugins/blog && go test ./... -count=1
```

构建服务：

```bash
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
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
| 基础设施包 | `pkg/database`, `pkg/cache`, `pkg/logger`, `pkg/httpserver`, `pkg/storage`, `pkg/plugin`, `pkg/sqlgen` |
| 共享响应和错误类型 | `types` |
| 远程插件示例 | `remote_plugins/blog` |
| Docker 与部署 | `Dockerfile`, `deploy`, `deploy.sh`, `script/install.sh` |
| 人类工程文档 | `docs/index.md` |
| AI 运行态 | `AGENTS.md`, `docs/ai` |

## API 范围

主服务当前注册的核心路由：

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

## 数据库 CLI

`db` 命令当前聚焦于生成 SQL 和 demo Todo CRUD。

```bash
go run ./cmd/main db --operation=schema
go run ./cmd/main db --operation=schema --apply
go run ./cmd/main db --operation=todo-list
```

已移除的 `initdb` 命令和 InitDB 配置段不得在没有新确认任务的情况下恢复。

## 文档

从 [`docs/index.md`](docs/index.md) 开始阅读。文档按真实工程边界组织：项目概览、目录地图、配置、架构、运行流程、模块、工作流、测试、构建、发布、扩展、维护、AI 协作和已知缺口。

## 生产注意事项

生产配置必须注入至少 32 bytes 的 `RIN_APP_AUTH_TOKEN_SECRET`。本地配置在没有显式 auth secret 时可以回退到进程内随机 token
secret，这只适合开发调试；服务重启后旧 token 会失效。

生产示例默认关闭 demo 模块。除非有明确确认，不要在生产环境暴露 demo 路由或隐式创建 demo schema。
