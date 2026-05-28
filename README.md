# go-scaffold

配置加载、动态环境变量前缀和 `envname` 字段约定见 [`docs/configuration.md`](docs/configuration.md)。

当前主服务已包含本地用户、认证与 RBAC 能力：`/api/v1/auth`、`/api/v1/users`、`/api/v1/roles` 和 `/api/v1/permissions`。该能力仍不代表生产级密钥管理、refresh token/session revoke、审计、密码重置或外部 IAM 已完成。

注意：当前项目仍在开发中，未达第一版发布条件；Docker 构建、部署示例和 CI 门禁只是阶段性制品，不代表 v1 可发布。

## 启动

```bash
go run ./cmd/server server
```

默认配置使用本地 SQLite：`./data/app.db`，Redis 关闭。启动后访问：

- `GET http://127.0.0.1:9999/health`
- `GET http://127.0.0.1:9999/ready`

## 目录边界

- `cmd/server`: 进程入口，只负责 CLI 参数和信号处理。
- `internal/app`: 应用装配层，按顺序初始化配置、日志、数据库、demo/user schema、缓存、executor、HTTP server。
- `internal/transport/http`: HTTP router、基础 health/ready 路由、demo API、用户/认证/RBAC API 注册。
- `internal/modules/demo`: 示例业务模块，展示 `model -> repository -> service -> handler` 写法。
- `internal/modules/user`: 主服务用户、认证与 RBAC 模块，包含用户/角色/权限模型、密码哈希、bearer token、权限门禁和 HTTP handler。
- `internal/config`: 配置结构、加载、环境变量覆盖、校验。
- `pkg/database`: 数据库连接、ping、关闭、事务。
- `pkg/executor`: 独立 goroutine pool manager。
- `pkg/logger`, `pkg/httpserver`, `pkg/cache`, `pkg/i18n`, `pkg/storage`, `pkg/sqlgen`, `pkg/plugin`, `pkg/utils`: 可复用基础设施库。

## Demo Todo API

- `POST /api/v1/demo/todos`
- `GET /api/v1/demo/todos`
- `GET /api/v1/demo/todos/:id`
- `PUT /api/v1/demo/todos/:id`
- `DELETE /api/v1/demo/todos/:id`

分层规则：

- handler 只做参数绑定、HTTP 状态码和响应转换。
- service 只做业务校验、事务编排和调用 repository。
- repository 只做 GORM 数据访问，不写业务判断。

## 测试

```bash
go test ./... -count=1
```

## CI 与部署

- CI 质量门禁见 `.github/workflows/ci.yml`。
- Linux Docker 镜像构建见 `Dockerfile`。
- production Compose 示例见 `deploy/docker-compose.production.example.yml`，production 配置样例见 `deploy/config.production.example.yaml`。
- 统一部署入口见 `deploy.sh`，支持 clone 后执行 `bash deploy.sh --docker y --confirm ...`。
- 直接下载安装入口见 `script/install.sh`，可通过 `curl -fsSL -o deploy.sh <raw-url>` 后执行。
- 手动远程部署 workflow 见 `.github/workflows/deploy-remote.yml`，支持 `staging` / `production` 手动环境选择。
- 部署边界、Secrets 配置和发布前检查见 `docs/deployment.md`。
- 当前不发布第一版；发布验收清单、真实 production 运行、镜像发布、生产迁移和密钥管理仍需单独确认。
