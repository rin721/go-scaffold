# ARCHITECTURE.md

## Architecture Status

- Project: go-scaffold
- Current Focus: Plugin system v1 closeout
- Status: COMPLETED
- Last Updated: 2026-05-25

## Plugin System Architecture

- [CONFIRMED] The plugin system is implemented as an independent package under `pkg/plugin`.
- [CONFIRMED] The package does not import `internal/*` and has no application-layer dependency.
- [CONFIRMED] v1 supports `local` and `http` protocols.
- [CONFIRMED] `rpc` and `ws` protocol constants are reserved for future adapters but are not implemented.
- [CONFIRMED] v1 API boundary is accepted as of 2026-05-25 after user advanced with `下一步`.

## Module Boundary

| Module | Responsibility | Does Not Do | Status |
|---|---|---|---|
| `pkg/plugin` | Define plugin interfaces, config, manager, local adapter, HTTP adapter, request/response wire types | Does not import app/internal code; does not auto-load application plugins | [CONFIRMED] |
| `plugins/*` | Optional host-side local plugin implementations | Not part of the independent library | [INFERRED] |
| `internal/app` | Application composition | Not touched by plugin system v1 | [CONFIRMED] |

## Public Interfaces

- [CONFIRMED] `Plugin` exposes `Metadata`, `Invoke`, and `Close`.
- [CONFIRMED] `Manager` exposes `Load`, `Register`, `RegisterLocalFactory`, `Invoke`, `Get`, `List`, and `Close`.
- [CONFIRMED] `Definition` describes a plugin instance by name, protocol, endpoint, headers, timeout, and metadata.
- [CONFIRMED] `Request` and `Response` provide the common JSON wire format for all protocols.

## Local Protocol

- [CONFIRMED] Local plugins are in-process Go implementations.
- [CONFIRMED] Local plugin implementations are provided through `LocalFactory`.
- [CONFIRMED] The library does not use Go dynamic plugins, keeping it cross-platform and independent.
- [INFERRED] Projects can place local implementations under `plugins/*` and register their factories from host code.

## HTTP Protocol

- [CONFIRMED] HTTP plugins are invoked with `POST` and JSON request bodies.
- [CONFIRMED] HTTP responses use the same `Response` shape.
- [CONFIRMED] Non-2xx responses return `ErrHTTPStatus`.
- [CONFIRMED] Configured headers and per-request headers are both supported.

## Future Protocols

- [DEFERRED] RPC adapter.
- [DEFERRED] WebSocket adapter.
- [DEFERRED] Plugin discovery from files or network registries.

## Next Architecture Decision

- [NEEDS_CONFIRMATION] Select one next promoted direction: RPC adapter, WebSocket adapter, discovery, examples/docs, or return to broader project optimization governance.
- [CONFIRMED] Unselected directions remain in `BACKLOG.md` and must not be implemented opportunistically.
