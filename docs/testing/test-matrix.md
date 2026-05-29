# 测试矩阵

测试分布在命令、应用、模块、package、type 和远程插件边界。测试应尽量靠近它保护的行为。

## 标准命令

```bash
go test ./... -count=1
cd remote_plugins/blog && go test ./... -count=1
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
```

CI 还会构建 Docker 镜像。

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
| `pkg/plugin` | registry、hook、HTTP 插件适配器 |
| `pkg/sqlgen` | DDL 生成和解析/生成行为 |
| `types` | result 和 error envelope 行为 |
| `remote_plugins/blog` | 独立远程插件 API 和 hook |

## 何时扩大测试范围

当变更触及以下内容时，应从单元测试扩大到集成式测试：

- `internal/app` 应用装配；
- 配置 reload 行为；
- 数据库 schema 生成/应用；
- auth/RBAC 路由授权；
- 插件注册或 hook dispatch；
- 共享响应/错误 helper。

## 已知测试缺口

middleware 与 `types/result` 之间的 trace ID 键名不一致，需要一个真实 Gin 请求穿过 middleware 和普通 result helper 的集成测试。
