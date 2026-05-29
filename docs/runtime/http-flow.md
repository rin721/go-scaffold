# HTTP 流程

HTTP 路由位于 `internal/transport/http`。router 在应用启动期间创建，此时模块和基础设施已经可用。

## 中间件顺序

主 router 安装这些 middleware：

1. i18n，可用时启用；
2. request trace ID；
3. CORS；
4. request logging；
5. panic recovery。

传输层 middleware 应只处理 HTTP 关注点。业务判断属于模块 service。

## 路由组

| 路由 | handler 来源 |
| --- | --- |
| `GET /health` | transport router |
| `GET /ready` | transport router，并执行 database ping |
| `/api/v1/demo/todos` | demo handler |
| `/api/v1/auth/*` | user auth handler |
| `/api/v1/users` | user handler，带认证和权限检查 |
| `/api/v1/roles` | role handler，带认证和权限检查 |
| `/api/v1/permissions` | permission handler，带认证和权限检查 |
| plugin register path | plugin 启用时注册 |

## 请求处理形态

```text
HTTP request
  -> middleware
  -> handler bind/parse
  -> service validation/business rules
  -> repository/database or infrastructure package
  -> service result
  -> handler result helper
  -> JSON response
```

handler 不应直接隐藏事务或业务规则。service 负责业务校验和事务边界。

## 认证与授权

用户相关路由通过 user 模块 token service 解析 bearer token。principal 会写入 Gin context，也会注入 request context，供 IAM/plugin
hook 增强读取。权限检查调用 user service，再委托给配置的 authorizer。
