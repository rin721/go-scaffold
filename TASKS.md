# TASKS.md

## Current Legal Task

- Task ID: TASK-PLUGIN-003
- Status: PENDING_USER_CONFIRMATION
- Time Slice: TS-PLUGIN-003
- Summary: Select next plugin/system direction after v1 local/http acceptance.

## Tasks

### TASK-PLUGIN-001: Implement Independent Plugin Package

- Status: COMPLETED
- Goal: Provide a standalone plugin system library under `pkg` with local and HTTP protocol support.
- Requirements Covered:
  - REQ-PLUGIN-001
  - REQ-PLUGIN-002
  - REQ-PLUGIN-003
  - REQ-PLUGIN-004
- Allowed Files:
  - `pkg/plugin/*`
  - Project documentation files
- Non-Goals:
  - No app-layer integration.
  - No `internal/*` imports.
  - No rpc/ws implementation in v1.
  - No Go dynamic plugin loader.
- Verification:
  - `go test ./pkg/plugin -count=1`
  - `go test ./... -count=1`
- Completion Evidence:
  - Both verification commands passed.

### TASK-PLUGIN-002: Review And Confirm Plugin Package API

- Status: COMPLETED
- Goal: User reviews exported API and confirms whether local/http v1 is sufficient.
- Requirements Covered:
  - REQ-PLUGIN-005
- Allowed Files:
  - Project documentation files
  - `pkg/plugin/*` only if API corrections are requested
- Verification:
  - `go test ./pkg/plugin -count=1`
- Completion Evidence:
  - User advanced with `下一步` after v1 implementation.
  - No API corrections requested.

### TASK-PLUGIN-003: Select Next Plugin Or Governance Work

- Status: PENDING_USER_CONFIRMATION
- Goal: Decide the next promoted work item now that `pkg/plugin` v1 local/http is accepted.
- Candidate Promotions:
  - RPC adapter from `BL-012`.
  - WebSocket adapter from `BL-013`.
  - Plugin discovery from `BL-014`.
  - Additional examples or documentation for local/http usage.
  - Return to broader project optimization governance from the original roadmap.
- Allowed Files:
  - Project documentation files until a candidate is confirmed.
  - Implementation files only after a candidate gets a task and time slice.
- Non-Goals:
  - No opportunistic rpc/ws/discovery implementation.
  - No application-layer integration unless explicitly requested.
  - No unrelated package refactors.
  - No auth/rbac promotion.
  - No migrations or deployment changes.
  - No staged/committed changes unless explicitly requested.
  - No broader `pkg/*` API policy changes without a separate decision.
- Verification:
  - Depends on selected candidate.
