# 添加模块

新增应用功能模块时使用本指南。

## 推荐形态

```text
internal/modules/<module>/
  model.go
  repository.go
  service.go
  handler.go
  *_test.go
```

不是每个模块都必须严格拥有这些文件，但职责应保持分离。

## 步骤

1. 定义持久化 model 和表名。
2. 实现 repository，只处理数据库访问。
3. 实现 service，处理校验、业务规则和事务。
4. 实现 HTTP handler，处理请求绑定和响应转换。
5. 在 `internal/app/initapp` 装配模块。
6. 在 `internal/transport/http` 注册路由。
7. 只有运行行为需要配置时才新增配置字段。
8. 在 service 和 transport 边界增加测试。
9. 在 `docs/modules` 下增加模块文档。

## 边界规则

- 不要把业务逻辑放进 `cmd/main`。
- 不要让 `pkg` 包导入 `internal/modules`。
- 不要把数据库事务藏在 handler 里。
- 不要把非通用能力塞进 `pkg/utils`。
- 不要在没有配置和文档的情况下新增生产可见默认行为。

## 添加路由

新路由应在模块由 `internal/app` 构建完成后，通过 transport 层注册。受保护路由应沿用 user 模块中的认证和权限模式。

## 添加 Schema

如果需要脚手架级生成 schema，沿用 `pkg/sqlgen` 模式，并明确记录 schema 是通过启动、DB CLI，还是未来独立迁移工作流应用。
