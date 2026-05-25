# go-scaffold

当前阶段只保留基础设施启动链路和一个 demo CRUD 示例，暂不实现 auth/rbac。

## 启动

```bash
go run ./cmd/server server
```

默认配置使用本地 SQLite：`./data/app.db`，Redis 关闭。启动后访问：

- `GET http://127.0.0.1:9999/health`
- `GET http://127.0.0.1:9999/ready`

## 目录边界

- `cmd/server`: 进程入口，只负责 CLI 参数和信号处理。
- `internal/app`: 应用装配层，按顺序初始化配置、日志、数据库、demo schema、缓存、executor、HTTP server。
- `internal/transport/http`: HTTP router、基础 health/ready 路由。
- `internal/modules/demo`: 示例业务模块，展示 `model -> repository -> service -> handler` 写法。
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
