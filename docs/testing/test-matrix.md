# 测试矩阵

测试分布在命令、应用、模块、package 和共享类型边界。测试应尽量靠近它保护的行为。

## 标准命令

```bash
go test ./... -count=1
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
```

CI 还会构建 Docker 镜像并检查空白字符。

## 归属矩阵

| 范围 | 典型测试 |
| --- | --- |
| `cmd/main` | CLI 参数解析、DB 命令行为 |
| `internal/app` | 应用装配、生命周期、DB schema helper、reload 行为 |
| `internal/config` | 配置加载、环境覆盖、校验 |
| `internal/transport/http` | router 注册、health/ready 行为 |
| `internal/modules/demo` | Todo service/repository/handler 行为 |
| `internal/modules/user` | 注册、登录、授权、RBAC seed 行为 |
| `pkg/database` | 连接管理、事务辅助、reload |
| `pkg/httpserver` | 启动、关闭、reload 行为 |
| `pkg/sqlgen` | DDL 生成和解析生成行为 |
| `types` | result 和 error envelope 行为 |

## 何时扩大测试范围

当变更触及以下内容时，应从单元测试扩大到集成式测试：

- `internal/app` 应用装配；
- 配置 reload 行为；
- 数据库 schema 生成或应用；
- auth/RBAC 路由授权；
- 共享响应或错误 helper。

## 已知测试缺口

middleware 中 `traceId` 和 `types/result` 的 `trace_id` 键名仍需真实 Gin
请求集成测试来统一确认。
