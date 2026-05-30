# 用户、认证与 RBAC 模块

用户模块提供本地用户管理、密码哈希、bearer token 登录、角色分配、权限分配和路由级授权。

## 位置

```text
internal/modules/user
pkg/auth
pkg/crypto
pkg/rbac
```

## 领域对象

| 对象 | 用途 |
| --- | --- |
| User | 本地账号，包含用户名、邮箱、密码哈希和状态 |
| Role | 命名角色，例如 admin 或 user |
| Permission | 授权检查使用的权限码 |
| UserRole | 用户和角色关系 |
| RolePermission | 角色和权限关系 |

## 认证流程

```text
register/login request
  -> user handler
  -> user service validation
  -> password hashing or password verification
  -> token service
  -> JSON response with bearer token
```

`GET /api/v1/auth/me` 读取当前 bearer token 并返回 principal。

## 授权流程

需要认证的管理路由使用 handler 层的认证和权限检查：

```text
Bearer token
  -> user service parses token
  -> principal in Gin context
  -> permission check
  -> Casbin authorizer
  -> handler proceeds or returns forbidden
```

principal 同时会注入 request context，便于应用内模块复用身份信息。

## 种子数据

配置区 RBAC seed 可以在启动时创建默认角色和权限。服务规则可让首个注册用户获得 admin 行为，后续用户按配置获得默认 user 角色。

## Token Secret 规则

如果配置了 `auth.token_secret`，长度必须至少 32 bytes。本地开发未配置时，应用可以生成进程内随机 secret。这个 fallback 不会跨重启稳定保存，不能作为生产密钥策略。

## 当前非目标

该模块尚不提供 refresh token、session revoke、密码重置、审计流、token rotation 或外部 IAM 集成。
