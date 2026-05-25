# types 契约边界说明

## 状态

- 任务：TASK-P1-009
- 时间切片：TS-P1-009
- 状态：COMPLETED
- 最后更新：2026-05-25

## 结论

- [CONFIRMED] `types/result` 是 HTTP API 响应契约包，其中 `helpers.go` 直接依赖 Gin，不能被视为纯领域类型包。
- [CONFIRMED] `types/errors` 是错误码和 `BizError` 契约包。3000-3999 是 auth/rbac 预留错误码范围，但当前项目未实现 auth/rbac。
- [CONFIRMED] `types/constants` 是跨层运行常量包，当前包含应用命令、配置路径、关闭超时、缓存键和 executor pool 名称。
- [CONFIRMED] 根 `types` 包是有限聚合入口，当前只提供 `pkg/crypto.Crypto` 别名和 `CacheInjectable` 接口。

## 包边界

| 包 | 边界定位 | 可依赖内容 | 不负责 |
|---|---|---|---|
| `types/result` | HTTP API 响应契约 | JSON 响应结构、分页结构、Gin 响应 helper、trace id 提取 | 领域模型、业务逻辑、认证实现 |
| `types/errors` | 错误码和业务错误契约 | 错误码分段、`BizError`、错误链 | HTTP 写响应、token 校验、权限判断 |
| `types/constants` | 跨层运行常量 | 应用命令名、默认配置路径、关闭超时、executor pool 名称、缓存键常量 | 可变配置、业务枚举、HTTP 响应 |
| `types` | 有限聚合入口 | `Crypto` 别名、`CacheInjectable` 接口 | 新业务类型集合、HTTP helper 聚合 |

## 依赖说明

- [CONFIRMED] `types/result` 依赖 `github.com/gin-gonic/gin`，因为它提供 `OK`、`BadRequest`、`Unauthorized`、`Forbidden`、`NotFound`、`InternalError`、`Fail` 和 `Page` helper。
- [CONFIRMED] `types/result` 依赖 `types/errors` 来映射响应错误码。
- [CONFIRMED] `types/constants` 依赖 `pkg/executor` 的 `PoolName` 类型，属于跨层运行契约。
- [CONFIRMED] 根 `types` 依赖 `pkg/cache` 和 `pkg/crypto`，用于保留跨层注入和加密接口的稳定入口。

## auth/rbac 范围

- [CONFIRMED] `ErrUnauthorized`、`ErrInvalidToken`、`ErrTokenExpired`、`ErrPermissionDenied` 是错误码契约。
- [CONFIRMED] 上述错误码不代表当前项目已实现登录、JWT、token 刷新、角色或权限系统。
- [DEFERRED] auth/rbac 实现仍保留在 Backlog，必须单独确认需求、架构、任务和验收标准。

## 验收证据

- `types/result/result_test.go` 固定基础响应结构、分页总页数、Gin helper 的 HTTP 状态码和错误码映射、trace id 提取。
- `types/errors/error_test.go` 固定 `BizError` 错误链和错误码分段。
- `types/constants/executor_test.go` 已固定 executor pool 名称常量。

## 非目标

- 不移动 `types/*` 包。
- 不修改 HTTP router、middleware 或 demo handler 行为。
- 不实现 auth/rbac。
- 不把 `pkg/*` 行为测试并入本切片。
