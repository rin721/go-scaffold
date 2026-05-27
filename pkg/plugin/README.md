# pkg/plugin - 插件注册与运行时

`pkg/plugin` 提供独立的插件运行时，用于本地插件和远程 HTTP 插件。

## API 分类

- 定位：[CONFIRMED] 公共基础设施 API。
- 稳定边界：local/http runtime、被动 `Manager.Register` registry、`Plugin`、`Request`、`Response`、`Definition`、`hooks`、HTTP server helper、`RemoteHook`。
- 当前风险：[DEFERRED] rpc/ws/discovery、插件发现和 Go `.so` 插件仍不属于当前稳定能力。
- 非目标：[CONFIRMED] 本包不依赖 `internal/*`、`pkg/iam`、日志或配置，不感知应用组合层，不主动发现、加载或注册插件服务。

本包不导入 `internal/*`，也不感知应用组合层。它是被动 registry/runtime：插件服务或宿主装配层负责创建插件实例，再把实例注册到 manager。

## 功能特性

- 统一的 `Plugin` 接口。
- `Manager` 提供被动注册、调用、列出和关闭能力。
- `Manager` 暴露 `Hooks()` 和 `RegisterHook`，支持注册、调用、错误、关闭、配置变化、日志就绪和权限检查等标准钩子点。
- 本地进程内插件由插件服务显式创建。
- HTTP 远程插件使用统一 JSON 请求/响应协议。
- `NewHTTPServer` 可把任意 `Plugin` 暴露为标准 HTTP 插件端点。
- `RemoteHook` 可把远程 HTTP 插件适配为本地 hook handler。
- 为后续 `rpc` 和 `ws` 适配器保留协议常量。
- 调用链路支持 `context.Context`。

## 注册边界

`pkg/plugin` 不会主动从配置加载插件服务。插件服务或宿主装配层负责服务生命周期和注册：

1. 构造本地或 HTTP 插件实例。
2. 调用 `mgr.Register(pluginInstance)`。
3. 使用 `mgr.Invoke`、`mgr.List` 和 `mgr.Close` 作为被动运行时能力。

## 本地插件

```go
mgr := plugin.NewManager()
echo, err := plugin.NewLocal(plugin.Metadata{
    Name:     "echo",
    Protocol: plugin.ProtocolLocal,
}, func(ctx context.Context, req plugin.Request) (*plugin.Response, error) {
    return plugin.NewResponse(map[string]string{"ok": "true"})
})
if err != nil {
    return err
}

err = mgr.Register(echo)
```

本地插件实现可以放在 `plugins/*` 中，但本包不会加载 Go dynamic plugin。这样可以保持库跨平台且独立于宿主应用。

## HTTP 插件

```go
mgr := plugin.NewManager()
remote, err := plugin.NewHTTP(plugin.Definition{
    Name:     "remote",
    Protocol: plugin.ProtocolHTTP,
    Endpoint: "https://example.com/plugin",
})
if err != nil {
    return err
}
err = mgr.Register(remote)

resp, err := mgr.Invoke(ctx, "remote", plugin.MustNewRequest("status", nil))
```

HTTP 适配器发送如下请求：

```json
{
  "plugin": "remote",
  "operation": "status",
  "payload": {}
}
```

插件端点应返回如下响应：

```json
{
  "payload": {},
  "metadata": {},
  "error": ""
}
```

HTTP 状态码不在 `2xx` 范围时返回 `ErrHTTPStatus`。

## Hook 运行时

```go
mgr := plugin.NewManager()
err := mgr.RegisterHook(plugin.HookBeforeInvoke, hooks.HandlerFunc(
    func(ctx context.Context, event hooks.Event) (hooks.Result, error) {
        return hooks.Result{}, nil
    },
), hooks.WithPriority(100))
```

Hook handler 按优先级从高到低执行；执行前会复制处理器列表；每个处理器执行前检查 `context.Context`；handler 可返回 `hooks.Stop(...)` 表示停止语义。

## HTTP 插件服务端

```go
handler := plugin.NewHTTPServer(localPlugin)
```

服务端只接受 `POST /plugin/v1/invoke`，请求和响应沿用 `Request` / `Response` JSON 结构。远程 hook 使用标准操作 `hooks.execute`，返回 payload 应可解码为 `hooks.Result`。
