# 插件系统

插件系统允许 host 注册插件元数据、调用插件操作并执行 hook 集成。当前设计是被动式：应用装配阶段注册配置中的插件，远程 blog 插件是示例实现。

## Host 侧包

| 路径 | 角色 |
| --- | --- |
| `pkg/plugin` | 插件管理器、registry、hook dispatcher、本地/HTTP 插件适配器 |
| `internal/app/initapp` | 创建 plugin manager 并接入 host hook |
| `internal/transport/http` | 启用插件时暴露注册路由 |
| `internal/iam` | 为 IAM/plugin hook 增强提供 principal context |

## 远程 Blog 插件

`remote_plugins/blog` 是独立 Go module。它提供：

- plugin manifest；
- health check；
- blog post create/list 操作；
- hook execute；
- 可选 shared secret 保护；
- 向 host 注册的能力。

它把数据保存在内存中，是演示插件，不是生产 blog 服务。

## Hook 流程

```text
host action
  -> plugin manager hook dispatch
  -> local or HTTP plugin adapter
  -> plugin operation/hook endpoint
  -> response normalized by plugin package
```

host 应用代码应在边界处注入应用特有上下文。可复用的 plugin package 不应导入业务模块。

## 扩展清单

新增插件集成时：

1. 决定是 local 插件还是 remote HTTP 插件；
2. 定义 manifest 和 operations；
3. 配置注册、timeout 和 max response size；
4. 需要时增加 hook binding；
5. 增加 package 和集成边界测试；
6. 记录失败行为和安全 header。
