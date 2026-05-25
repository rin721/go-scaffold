# Plugin Package

`pkg/plugin` provides an independent plugin runtime for local and remote plugins.

## API 分类

- 定位：[CONFIRMED] 公共基础设施 API。
- 稳定边界：v1 local/http runtime、`Plugin`、`Manager`、`Request`、`Response`、`Definition`。
- 当前风险：[DEFERRED] rpc/ws/discovery 仍为预留协议，不属于当前稳定能力。
- 非目标：[CONFIRMED] 本包不依赖 `internal/*`，不感知应用组合层。

The package does not import `internal/*` and does not know about the application
composition layer. Projects can keep local implementations under `plugins/*`,
then register factories with the manager.

## Features

- Unified `Plugin` interface.
- `Manager` for load, register, invoke, list, and close operations.
- Local in-process plugins through registered factories.
- HTTP remote plugins through a JSON request/response protocol.
- Reserved protocol constants for future `rpc` and `ws` adapters.
- Context-aware invocation.

## Local Plugin

```go
mgr := plugin.NewManager(plugin.WithLocalFactory("echo", func(def plugin.Definition) (plugin.Plugin, error) {
    return plugin.NewLocal(def.Metadata(), func(ctx context.Context, req plugin.Request) (*plugin.Response, error) {
        return plugin.NewResponse(map[string]string{"ok": "true"})
    })
}))

err := mgr.Load(&plugin.Config{Plugins: []plugin.Definition{
    {Name: "echo", Protocol: plugin.ProtocolLocal},
}})
```

Local plugin implementations can live in `plugins/*`, but this package does not
load Go dynamic plugins. That keeps the library cross-platform and independent.

## HTTP Plugin

```go
mgr := plugin.NewManager()
err := mgr.Load(&plugin.Config{Plugins: []plugin.Definition{
    {
        Name:     "remote",
        Protocol: plugin.ProtocolHTTP,
        Endpoint: "https://example.com/plugin",
    },
}})

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
