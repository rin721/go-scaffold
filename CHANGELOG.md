# CHANGELOG.md

## Latest Changes

### 2026-05-25 - TASK-PLUGIN-002 - TS-PLUGIN-002

- Changed: Accepted and closed `pkg/plugin` v1 local/http API boundary.
- Files:
  - `STATUS.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `ARCHITECTURE.md`
  - `DECISIONS.md`
  - `ACCEPTANCE.md`
  - `BACKLOG.md`
  - `RISK_REGISTER.md`
  - `REQUIREMENTS.md`
  - `TEST_REPORT.md`
  - `CHANGELOG.md`
  - `AGENT_HANDOFF.md`
- Reason: User advanced with `下一步` after plugin v1 implementation.
- Tests:
  - `go test ./pkg/plugin -count=1`: PASS
- Status: COMPLETED.
- Notes: RPC, WebSocket, discovery, examples, and app integration remain separate promotions.

### 2026-05-25 - TASK-PLUGIN-001 - TS-PLUGIN-001

- Changed: Implemented independent plugin system package with local and HTTP protocol support.
- Files:
  - `pkg/plugin/*`
  - `README.md`
  - `ARCHITECTURE.md`
  - `DECISIONS.md`
  - `TASKS.md`
  - `TIME_SLICES.md`
  - `REQUIREMENTS.md`
  - `ACCEPTANCE.md`
  - `BACKLOG.md`
  - `RISK_REGISTER.md`
  - `STATUS.md`
  - `TEST_REPORT.md`
  - `AGENT_HANDOFF.md`
- Reason: User requested a standalone plugin library under `pkg` with local/http support.
- Tests:
  - `go test ./pkg/plugin -count=1`: PASS
  - `go test ./... -count=1`: PASS
- Status: PENDING_USER_CONFIRMATION.
- Notes: No application-layer integration was added.

## Historical Changes

### 2026-05-25 - REQ-001 - TS-REQ-001

- Changed: Generated formal requirements phase documents from the project startup package.
- Files:
  - `REQUIREMENTS.md`
  - `ACCEPTANCE.md`
  - `BACKLOG.md`
  - `RISK_REGISTER.md`
  - `STATUS.md`
  - `TEST_REPORT.md`
  - `AGENT_HANDOFF.md`
- Reason: User requested `下一步`; default startup route was used to move into requirements confirmation.
- Tests: `go test ./... -count=1` passed.
- Status: COMPLETED.
- Notes: No Go code was changed.

### 2026-05-25 - INTAKE-001 - TS-INTAKE-001

- Changed: Generated project startup package.
- Files:
  - `PROJECT_BRIEF.md`
  - `STATUS.md`
  - `docs/templates/project_start_template.md`
  - `docs/templates/requirements_clarification_template.md`
  - `docs/templates/technical_options_template.md`
  - `docs/templates/architecture_constraints_template.md`
  - `docs/templates/acceptance_template.md`
  - `docs/templates/risk_confirmation_template.md`
- Reason: Establish project facts before implementation.
- Tests: `go test ./... -count=1` passed.
- Status: PENDING_USER_CONFIRMATION.
