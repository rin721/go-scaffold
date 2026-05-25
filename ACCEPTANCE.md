# ACCEPTANCE.md

## Acceptance Status

- Project: go-scaffold
- Phase: Requirements Confirmation
- Status: PENDING_USER_CONFIRMATION
- Last Updated: 2026-05-25

## Project-Level Acceptance

| ID | Acceptance Item | Method | Required | Status |
|---|---|---|---|---|
| ACC-PROJ-001 | Project facts are stored in repository documents | Review `PROJECT_BRIEF.md`, `REQUIREMENTS.md`, `STATUS.md` | Yes | [CONFIRMED] |
| ACC-PROJ-002 | Current optimization route is documented | Review `REQUIREMENTS.md` and `STATUS.md` | Yes | [CONFIRMED] |
| ACC-PROJ-003 | Code implementation is blocked until confirmation | Review `STATUS.md` | Yes | [CONFIRMED] |
| ACC-PROJ-004 | Full Go tests pass after documentation changes | Run `go test ./... -count=1` | Yes | [CONFIRMED] |
| ACC-PROJ-005 | Pending decisions are explicit | Review open questions in `REQUIREMENTS.md` | Yes | [CONFIRMED] |

## Requirements Phase Acceptance

| ID | Acceptance Item | Method | Required | Status |
|---|---|---|---|---|
| ACC-REQ-001 | P0/P1/P2 requirements are separated | Review `REQUIREMENTS.md` | Yes | [CONFIRMED] |
| ACC-REQ-002 | P0 requirements have acceptance criteria | Review P0 table in `REQUIREMENTS.md` | Yes | [CONFIRMED] |
| ACC-REQ-003 | Non-goals are explicit | Review `Explicit Non-Requirements` | Yes | [CONFIRMED] |
| ACC-REQ-004 | Backlog captures deferred ideas | Review `BACKLOG.md` | Yes | [CONFIRMED] |
| ACC-REQ-005 | Risks are registered with mitigation | Review `RISK_REGISTER.md` | Yes | [CONFIRMED] |

## Future Architecture Acceptance

| ID | Acceptance Item | Method | Required | Status |
|---|---|---|---|---|
| ACC-ARCH-001 | Architecture boundaries are confirmed | Future `ARCHITECTURE.md` | Yes | [NEEDS_CONFIRMATION] |
| ACC-ARCH-002 | `pkg/*` API policy is recorded | Future `DECISIONS.md` | Yes | [NEEDS_CONFIRMATION] |
| ACC-ARCH-003 | Demo module role is decided | Future `ARCHITECTURE.md` | Yes | [NEEDS_CONFIRMATION] |
| ACC-ARCH-004 | Migration strategy is decided | Future `ARCHITECTURE.md` | Yes | [NEEDS_CONFIRMATION] |
| ACC-ARCH-005 | Test strategy covers app/router/demo/config paths | Future `ARCHITECTURE.md` and `TEST_REPORT.md` | Yes | [NEEDS_CONFIRMATION] |

## Future Implementation Acceptance

Future code-level tasks cannot be marked `COMPLETED` unless all apply:

1. [CONFIRMED] The task maps to a confirmed requirement.
2. [CONFIRMED] The task maps to a confirmed architecture boundary.
3. [CONFIRMED] The current time slice authorizes the file scope.
4. [CONFIRMED] Relevant tests or verification commands ran.
5. [CONFIRMED] `STATUS.md`, `CHANGELOG.md`, `TEST_REPORT.md`, and `AGENT_HANDOFF.md` are updated.
6. [CONFIRMED] The next legal step is clear.

## Current Completion Decision

- Requirements documents generated: COMPLETED
- Requirements phase confirmation: PENDING_USER_CONFIRMATION
- Code implementation: BLOCKED
