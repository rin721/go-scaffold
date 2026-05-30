# 配置流

配置是运行时状态。它在基础设施和模块创建前加载，并在 server 模式下被 watch。

## 首次加载

```text
command flag/env/default path
  -> Manager.Load
  -> .env load
  -> YAML read
  -> placeholder substitution
  -> unmarshal
  -> env override
  -> validation
  -> atomic store
```

config manager 向应用装配代码暴露当前配置。各 package 通常只接收自己需要的
配置片段，而不是整份 app config。

## Watch 与 Reload

server 模式会注册配置 watcher。配置变化时，manager 会 shadow-load 新文件，
应用环境变量覆盖，完成校验，然后存储新配置并调用变更处理器。

`reloadapp` 比较新旧配置，并 reload 受影响子系统：

- logger；
- database；
- cache；
- executor；
- HTTP server；
- storage；
- IAM。

demo schema 应用属于启动关注点，不会在普通 reload 中重新执行。

## 维护清单

修改配置结构时：

1. 更新 `internal/config` 中的结构体、tag 和校验；
2. 更新 `configs` 和 `deploy` 下的示例；
3. 字段面向环境变量时更新 `.env.example`；
4. 更新运行文档；
5. 增加 env override 和校验测试；
6. 判断是否需要 reload 行为。
