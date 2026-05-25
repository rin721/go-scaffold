# Plugin Package

`pkg/plugin` provides an independent plugin runtime for local and remote plugins.

## API 分类

- 定位：[CONFIRMED] 公共基础设施 API。
- 稳定边界：v1 local/http runtime、被动 `Manager.Register` registry、`Plugin`、`Request`、`Response`、`Definition`。
- 当前风险：[DEFERRED] rpc/ws/discovery 仍为预留协议，不属于当前稳定能力。
- 非目标：[CONFIRMED] 本包不依赖 `internal/*`，不感知应用组合层，不主动发现、加载或注册插件服务。

The package does not import `internal/*` and does not know about the application
composition layer. It is a passive registry/runtime: plugin services or the
host composition layer create plugin instances, then register them with the
manager.

## Features

- Unified `Plugin` interface.
- `Manager` for passive register, invoke, list, and close operations.
- Local in-process plugins created by plugin services.
- HTTP remote plugins through a JSON request/response protocol.
- Reserved protocol constants for future `rpc` and `ws` adapters.
- Context-aware invocation.

## Registration Boundary

`pkg/plugin` does not actively load plugin services from configuration. A plugin
service or host composition layer owns service lifecycle and registration:

1. Build a local or HTTP plugin instance.
2. Call `mgr.Register(pluginInstance)`.
3. Use `mgr.Invoke`, `mgr.List`, and `mgr.Close` as the passive runtime.

## Local Plugin

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

Local plugin implementations can live in `plugins/*`, but this package does not
load Go dynamic plugins. That keeps the library cross-platform and independent.

## HTTP Plugin

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

The HTTP adapter sends:

```json
{
  "plugin": "remote",
  "operation": "status",
  "payload": {}
}
```

The endpoint should return:

```json
{
  "payload": {},
  "metadata": {},
  "error": ""
}
```

HTTP status codes outside `2xx` return `ErrHTTPStatus`.
