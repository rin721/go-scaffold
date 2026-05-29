# 分层架构

服务采用务实分层架构。关键设计点是：`internal/app` 是唯一知道如何把所有子系统装配到一起的组合根。

## 层级

```text
cmd/main
  -> internal/app
      -> internal/config
      -> pkg infrastructure
      -> internal/modules
      -> internal/transport/http
```

| 层 | 职责 |
| --- | --- |
| CLI | 解析命令参数、选择配置路径、处理进程信号 |
| 应用装配 | 创建核心服务、基础设施、模块、传输层和生命周期 hook |
| 配置 | 加载、校验、watch 并暴露当前配置 |
| 基础设施 | 数据库、缓存、日志、executor、存储、认证、RBAC、插件支撑 |
| 模块 | 业务行为和模块内校验 |
| 传输层 | HTTP 路由、middleware、请求绑定、响应转换 |

## 依赖方向

应用模块可以使用基础设施接口和 package。可复用的 `pkg` 包不得导入应用模块。`cmd/main` 应保持轻量，不承载业务逻辑。

当前 `types/result` 依赖 Gin，因此 `types` 并非完全与传输层无关。修改 result helper 时必须承认这个现状。

## 模块形态

Demo 和 User 模块都遵循同一基本形态：

```text
model -> repository -> service -> handler
```

- model：持久化结构和表信息；
- repository：只处理数据库访问；
- service：处理校验、事务和业务规则；
- handler：处理 HTTP 请求绑定、状态码选择和响应转换。

## 组合根

`internal/app/initapp` 按以下顺序构建应用：

1. 核心服务：config、logger、i18n、ID generator；
2. 基础设施：database、cache、executor、storage、IAM、plugin manager；
3. 模块：demo 和 user；
4. 传输层：主 HTTP server 和可选 plugin HTTP server。

reload 和 shutdown 由 app 专属包统一编排，而不是散落到模块内部。

## 扩展规则

新增功能时，优先在 `internal/modules/<name>` 下添加应用模块，并在 `internal/app/initapp` 中装配。只有当能力确实与应用业务无关时，才放入
`pkg` 作为可复用支撑。
