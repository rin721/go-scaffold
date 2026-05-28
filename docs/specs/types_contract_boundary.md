# types 契约边界说明

## 状态

- 任务：TASK-P1-009
- 时间切片：TS-P1-009
- 状态：COMPLETED
- 最后更新：2026-05-27

## 结论

- [CONFIRMED] `types/result` 是 HTTP API 响应契约包，其中 `helpers.go` 直接依赖 Gin，不能被视为纯领域类型包。
- [CONFIRMED] `types/errors` 是错误码和 `BizError` 契约包。3000-3999 是 auth/rbac 错误码范围，当前由 `internal/modules/user` 的认证与权限 HTTP 层使用。
- [CONFIRMED] `types/constants` 是应用层以上运行常量包，当前包含应用命令、配置路径、关闭超时、缓存键和 executor pool 名称；常量保持为字符串，不直接导入 `pkg/executor`。
- [ACCEPT] 用户修正：根 `types` 包不得直接向 `pkg/crypto.Crypto` 提供别名，也不得定义依赖 `pkg/cache.Cache` 的 `CacheInjectable` 接口；`types` 只能承载应用层以上确认过的跨层契约。
- [CONFIRMED] 根 `types` 包当前不再导入 `pkg/cache` 或 `pkg/crypto`，也不再作为 `pkg/*` 基础设施聚合入口。

## 包边界

| 包 | 边界定位 | 可依赖内容 | 不负责 |
|---|---|---|---|
| `types/result` | HTTP API 响应契约 | JSON 响应结构、分页结构、Gin 响应 helper、trace id 提取 | 领域模型、业务逻辑、认证实现 |
| `types/errors` | 错误码和业务错误契约 | 错误码分段、`BizError`、错误链 | HTTP 写响应、token 校验、权限判断 |
| `types/constants` | 应用层以上运行常量 | 应用命令名、默认配置路径、关闭超时、executor pool 名称、缓存键常量；不直接依赖 `pkg/*` | 可变配置、业务枚举、HTTP 响应 |
| `types` | 应用层以上契约说明入口 | 已确认的应用级跨层契约说明；当前无 `pkg/*` 别名或注入接口 | `pkg/*` 基础设施聚合、新业务类型集合、HTTP helper 聚合 |

## 依赖说明

- [CONFIRMED] `types/result` 依赖 `github.com/gin-gonic/gin`，因为它提供 `OK`、`BadRequest`、`Unauthorized`、`Forbidden`、`NotFound`、`InternalError`、`Fail` 和 `Page` helper。
- [CONFIRMED] `types/result` 依赖 `types/errors` 来映射响应错误码。
- [CONFIRMED] `types/constants` 不再依赖 `pkg/executor` 的 `PoolName` 类型；executor pool 名称以字符串常量形式保留，应用层调用 executor 时由使用点承担类型适配。
- [CONFIRMED] `types/*` 不依赖 `pkg/*`；需要缓存、加密、executor 等基础设施能力的应用层代码应显式依赖对应 `pkg/*` 包，或在应用层以上另行定义本层契约。

## auth/rbac 范围

- [CONFIRMED] `ErrUnauthorized`、`ErrInvalidToken`、`ErrTokenExpired`、`ErrPermissionDenied` 是错误码契约。
- [CONFIRMED] 上述错误码现在覆盖 `internal/modules/user` 的 bearer token 登录、token 校验和权限拒绝响应。
- [DEFERRED] refresh token、session revoke、生产级密钥管理、外部 IAM/OPA/Casbin、审计和账号恢复仍需单独确认需求、架构、任务和验收标准。

## 验收证据

- `types/result/result_test.go` 固定基础响应结构、分页总页数、Gin helper 的 HTTP 状态码和错误码映射、trace id 提取。
- `types/errors/error_test.go` 固定 `BizError` 错误链和错误码分段。
- `types/constants/executor_test.go` 已固定 executor pool 名称字符串常量。
- `types/import_boundary_test.go` 固定 `types/*` 包不得直接导入 `pkg/*` 基础设施包。

## 非目标

- 不移动 `types/*` 包。
- 不修改 HTTP router、middleware 或 demo handler 行为。
- 不实现 production-grade IAM、refresh token/session revoke、外部策略引擎或账号恢复。
- 不把 `pkg/*` 行为测试并入本切片。
