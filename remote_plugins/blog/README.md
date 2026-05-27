# Remote Blog Plugin

`remote_plugins/blog` is an independently deployable Blog plugin sample. It exposes the standard plugin JSON endpoint and can register itself with the host service.

## Startup Flow

1. Start the host service with plugin registration enabled.
2. Start this Blog plugin service.
3. The Blog plugin posts its HTTP invoke endpoint to the host registration endpoint.
4. The host creates an HTTP plugin adapter and uses the existing `hooks.execute` JSON protocol for remote hooks.

## Environment

| Variable | Purpose |
|---|---|
| `BLOG_PLUGIN_LISTEN_ADDR` | Blog service listen address, for example `127.0.0.1:18081`. |
| `BLOG_PLUGIN_PUBLIC_HTTP_URL` | Public URL the host uses, for example `http://blog.example.com:18081`. |
| `BLOG_PLUGIN_MAIN_HTTP_URL` | Host HTTP base URL, for example `http://host.example.com:9999`. |
| `BLOG_PLUGIN_MAIN_WS_URL` | Reserved host WS URL, for example `ws://host.example.com:9999/ws`; this sample does not implement WS transport. |
| `BLOG_PLUGIN_REGISTRATION_TOKEN` | Placeholder shared token for host registration. Do not commit real values. |
| `BLOG_PLUGIN_SHARED_SECRET` | Optional placeholder secret sent by the host when invoking the Blog plugin. |
| `BLOG_PLUGIN_HOOK_POINTS` | Comma-separated hook points, default `plugin.after_invoke`. |

## Operations

- `manifest`
- `health`
- `blog.create`
- `blog.list`
- `hooks.execute`

The hook event may include `identity.principal` when the host request context has an IAM principal. The host sends only safe principal fields; it must not send IAM tokens, policies, credentials, or service internals.

## Local Check

```sh
go test ./...
```
