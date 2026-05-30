# 启动流程

主进程从 `cmd/main` 开始。

```text
main.go
  -> cli.NewApp
  -> server command
  -> runApp
  -> app.New
  -> application.Run
```

## server 命令

`cmd/main/app.go` 定义 `server` 命令。它支持 `--config` 并读取 `RIN_CONFIG_PATH`。命令层只处理入口参数，实际启动交给 `runApp`。

## 应用构建

`internal/app.New` 通过应用子包构建应用：

| 包 | 作用 |
| --- | --- |
| `initapp` | 创建核心服务、基础设施、模块和传输层 |
| `modeapp` | 构建选定运行模式；当前真实模式是 `server` |
| `lifecycleapp` | 启动并关闭 HTTP、存储、执行器、缓存、数据库和日志 |
| `reloadapp` | 将配置变化应用到可重载子系统 |

## 表结构应用

启动期间，只有在 Demo 模块启用且 `demo.apply_schema_on_start` 为 true 时才会应用 Demo 表结构。

这是脚手架基线能力，不是完整生产迁移框架。

## HTTP 启动

HTTP 服务由 `pkg/httpserver` 包装。端口绑定错误会同步返回。

## 关闭流程

`cmd/main/run.go` 监听 `SIGINT` 和 `SIGTERM`。关闭过程使用配置的超时时间，并按以下顺序释放资源：

1. 主 HTTP 服务；
2. 存储；
3. 执行器；
4. 缓存；
5. 数据库；
6. 日志 sync。
