# 启动流程

主进程从 `cmd/main` 启动。

```text
main.go
  -> cli.NewApp
  -> server command
  -> runApp
  -> app.New
  -> application.Run
```

## server 命令

`cmd/main/app.go` 定义 `server` 命令。它支持 `--config`，也读取 `RIN_CONFIG_PATH`。命令本身只负责入口参数，实际启动交给
`runApp`。

## 应用构建

`internal/app.New` 负责构建应用，但细节拆分到多个 app 包：

| 包 | 角色 |
| --- | --- |
| `initapp` | 创建核心服务、基础设施、模块和传输层 |
| `modeapp` | 构建指定运行模式；当前真实模式是 server |
| `lifecycleapp` | 启动和关闭 HTTP、插件、存储、executor、缓存、数据库、日志 |
| `reloadapp` | 将配置变更应用到可 reload 的子系统 |

## Schema 应用

启动期间：

- demo 模块启用且 `demo.apply_schema_on_start` 为 true 时应用 demo schema；
- user 模块初始化时应用 user schema；
- 配置开启时应用 RBAC seed 数据。

这是脚手架基线，不是完整生产迁移体系。

## HTTP 启动

主 HTTP server 先启动，可选 plugin HTTP server 后启动。如果 plugin server 启动失败，主 server 会被关闭。端口绑定错误由 HTTP server
包装器同步返回。

## 关闭流程

`cmd/main/run.go` 监听 `SIGINT` 和 `SIGTERM`。关闭流程使用配置的 app shutdown timeout，并按以下顺序关闭：

1. plugin HTTP server；
2. main HTTP server；
3. storage；
4. plugin manager；
5. executor；
6. cache；
7. database；
8. logger sync。
