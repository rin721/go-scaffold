# DECISIONS.md

## Decision Records

### DEC-001: Implement Plugin System As Independent `pkg/plugin`

- Date: 2026-05-25
- Status: ACCEPTED
- Context: User requested a plugin system under `pkg` that is independent from the application layer and supports multiple protocol adapters.
- Decision: Implement a standalone `pkg/plugin` package with no `internal/*` imports.
- Alternatives: Wire plugins through `internal/app`; use Go dynamic plugins; defer implementation until full architecture phase.
- Reason: The requested system is a reusable library and must not touch application composition.
- Consequences: Host applications must register local plugin factories explicitly; the library stays cross-platform.
- Related Requirements: REQ-PLUGIN-001, REQ-PLUGIN-002
- Related Tasks: TASK-PLUGIN-001

### DEC-002: Use Local Factory Registration Instead Of Dynamic Plugin Loading

- Date: 2026-05-25
- Status: ACCEPTED
- Context: User wants local plugins under `plugins/*`, but the library must stay independent and cross-platform.
- Decision: Provide `LocalFactory` registration and let host code import/register local implementations.
- Alternatives: Use Go `plugin` package; scan and compile plugin folders; make app layer aware of plugins.
- Reason: Go dynamic plugins are not portable across all supported environments, and auto-loading would couple the library to host layout.
- Consequences: Local plugin loading is explicit, testable, and independent.
- Related Requirements: REQ-PLUGIN-003
- Related Tasks: TASK-PLUGIN-001

### DEC-003: Define A JSON Wire Contract For HTTP Plugins

- Date: 2026-05-25
- Status: ACCEPTED
- Context: v1 needs remote plugin support over HTTP.
- Decision: HTTP plugins use `POST` with JSON `Request` and JSON `Response`.
- Alternatives: Custom per-plugin payloads; GET-based calls; protocol-specific interfaces.
- Reason: Shared wire types keep the manager protocol-neutral.
- Consequences: HTTP plugin servers must implement the documented JSON contract.
- Related Requirements: REQ-PLUGIN-004
- Related Tasks: TASK-PLUGIN-001

### DEC-004: Accept `pkg/plugin` v1 Boundary

- Date: 2026-05-25
- Status: ACCEPTED
- Context: `pkg/plugin` v1 has local and HTTP support, tests pass, and the user advanced with `下一步` after API review.
- Decision: Treat v1 local/http API as accepted for the current phase.
- Alternatives: Revise exported API before closeout; immediately add rpc/ws/discovery.
- Reason: v1 satisfies the confirmed request while keeping future protocols controlled by backlog promotion.
- Consequences: RPC, WebSocket, discovery, examples, and app integration require separate task/time-slice promotion.
- Related Requirements: REQ-PLUGIN-001, REQ-PLUGIN-002, REQ-PLUGIN-003, REQ-PLUGIN-004, REQ-PLUGIN-005, REQ-PLUGIN-006
- Related Tasks: TASK-PLUGIN-002, TASK-PLUGIN-003
