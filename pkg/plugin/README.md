# pkg/plugin - 插件注册与运行时

`pkg/plugin` 提供独立的插件运行时，用于本地插件和远程 HTTP 插件。

## API 分类

- 定位：[CONFIRMED] 公共基础设施 API。
- 稳定边界：v1 local/http runtime、被动 `Manager.Register` registry、`Plugin`、`Request`、`Response`、`Definition`。
- 当前风险：[DEFERRED] rpc/ws/discovery 仍为预留协议，不属于当前稳定能力。
- 非目标：[CONFIRMED] 本包不依赖 `internal/*`，不感知应用组合层，不主动发现、加载或注册插件服务。

本包不导入 `internal/*`，也不感知应用组合层。它是被动 registry/runtime：插件服务或宿主装配层负责创建插件实例，再把实例注册到 manager。

## 功能特性

- 统一的 `Plugin` 接口。
- `Manager` 提供被动注册、调用、列出和关闭能力。
- 本地进程内插件由插件服务显式创建。
- HTTP 远程插件使用统一 JSON 请求/响应协议。
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
