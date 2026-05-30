# 测试矩阵

测试分布在命令、应用包、模块、可复用包和共享类型边界附近。测试应尽量靠近它保护的行为。

## 标准命令

```bash
go test ./... -count=1
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
```

CI 还会构建 Docker 镜像并检查空白字符。

## 归属矩阵

| 范围 | 常见测试 |
| --- | --- |
| `cmd/main` | CLI 解析和 DB 命令行为 |
| `internal/app` | 应用装配、生命周期、DB 表结构辅助、重载行为 |
| `internal/config` | 配置加载、环境变量覆盖、校验 |
| `internal/transport/http` | 路由注册、health/ready 行为 |
| `internal/modules/demo` | Todo service/repository/handler 行为 |
| `pkg/database` | 连接管理、事务辅助、重载 |
| `pkg/httpserver` | 启动、关闭、重载行为 |
| `pkg/sqlgen` | DDL 生成和 parser/generator 行为 |
| `types` | 响应和错误信封行为 |

## 何时扩大测试范围

当改动触及以下区域时，应从单元测试扩大到集成风格测试：

- `internal/app` 装配；
- 配置重载行为；
- 数据库表结构生成或应用；
- 共享响应或错误辅助函数；
- HTTP 中间件行为。

## 已知测试缺口

中间件使用 `traceId`，部分 result helper 读取 `trace_id`。后续应增加真实 Gin 请求集成测试来固定预期 key 格式。
