# pkg/iam - 身份与权限公共接口

`pkg/iam` 提供独立的身份认证与授权公共接口，以及最小内存实现。它面向基础设施组合层，不绑定 HTTP、数据库、插件系统或业务登录流程。

## API 分类

- 定位：[CONFIRMED] 公共基础设施 API。
- 稳定边界：`Principal`、`Credential`、`Action`、`Resource`、`Decision`、`Authenticator`、`Authorizer`、`Service`、`Policy`、context helper 和 `memory` 实现。
- 当前风险：[DEFERRED] JWT 中间件、数据库版权限、OPA/Casbin、业务登录流程和密钥管理不属于当前稳定能力。
- 非目标：[CONFIRMED] 本包不导入 `pkg/plugin` 或 `internal/*`，不主动接入 HTTP 中间件，不读取真实密钥。

## 功能特性

- token / bearer 凭证认证契约。
- 主体、动作、资源和授权决策类型。
- context 中写入和读取主体的 helper。
- memory service 支持精确匹配和 `*` 通配策略。
- deny 优先、策略过期检查和默认拒绝。

## 内存实现

```go
svc, err := memory.NewService(memory.Config{
    DefaultDeny: true,
    Principals: map[string]iam.Principal{
        "token": {ID: "user-1"},
    },
    Policies: []iam.Policy{
        {
            Subject:  "user-1",
            Action:   "plugin.invoke",
            Resource: "echo",
            Effect:   iam.EffectAllow,
        },
    },
})
if err != nil {
    return err
}
```

内存实现适合测试、开发和基础设施组合验证。生产登录、令牌签发、数据库权限存储和密钥轮换需要单独设计。
