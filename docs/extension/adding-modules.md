# 新增模块

新增应用能力时，优先放在 `internal/modules/<name>`，再通过 `internal/app/initapp` 装配。

## 推荐形态

```text
model -> repository -> service -> handler
```

只使用实际需要的层。很小的模块起步时不一定需要每个文件都齐全。

## 接线步骤

1. 在 `internal/modules/<name>` 下新增模块代码。
2. 如需配置，在 `internal/config` 中新增配置结构。
3. 在 `internal/app/initapp` 中装配 repository、service、handler。
4. 在 `internal/transport/http` 中注册路由。
5. 为 service 和路由行为补充聚焦测试。
6. 如果属于托管任务，同步更新工程文档和运行时证据。

## 身份和访问控制

之前内置的 user/IAM/auth/RBAC 服务已经移除。新的身份认证或访问控制工作应从需求、架构和验收标准重新开始。
