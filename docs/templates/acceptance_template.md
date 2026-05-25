# Acceptance Template

## Project Intake Acceptance

| ID | Acceptance Item | Method | Required | Status |
|---|---|---|---|---|
| ACC-INTAKE-001 | `PROJECT_BRIEF.md` exists | File check | Yes | [CONFIRMED] |
| ACC-INTAKE-002 | `STATUS.md` exists and status is `PENDING_USER_CONFIRMATION` | File review | Yes | [CONFIRMED] |
| ACC-INTAKE-003 | Required startup templates exist under `docs/templates/` | File check | Yes | [CONFIRMED] |
| ACC-INTAKE-004 | Facts use labels from `docs/ai/prompt.md` | Manual review | Yes | [CONFIRMED] |
| ACC-INTAKE-005 | Next step is user confirmation, not code implementation | `STATUS.md` review | Yes | [CONFIRMED] |

## Future Optimization Acceptance

| ID | Acceptance Item | Method | Required | Status |
|---|---|---|---|---|
| ACC-FUTURE-001 | Full Go tests pass after code changes | `go test ./... -count=1` | Yes | [NEEDS_CONFIRMATION] |
| ACC-FUTURE-002 | Startup command remains documented and valid | Manual or smoke test | Yes | [NEEDS_CONFIRMATION] |
| ACC-FUTURE-003 | `/health` and `/ready` stay available | HTTP smoke test | Yes | [NEEDS_CONFIRMATION] |
| ACC-FUTURE-004 | README and project status docs stay synchronized | Manual review | Yes | [NEEDS_CONFIRMATION] |
| ACC-FUTURE-005 | Each P0 optimization has evidence | Command/test/status report | Yes | [NEEDS_CONFIRMATION] |

## Non-Functional Acceptance

- [INFERRED] Maintainability: Each accepted change identifies the boundary it affects.
- [INFERRED] Testability: Each implementation task has a verification command or documented reason it cannot be run.
- [INFERRED] Recoverability: A future agent can determine the next task from repository documents.
- [INFERRED] Safety: No deployment, production migration, or secret exposure occurs without explicit confirmation.

## Completion Gate

A future task can be marked `COMPLETED` only when:

1. [CONFIRMED] Its required output exists.
2. [CONFIRMED] Its scope did not exceed the current time slice.
3. [CONFIRMED] Required verification ran or a blocker was recorded.
4. [CONFIRMED] Relevant docs and status were updated.
5. [CONFIRMED] The next legal step is clear.

## Current Completion Decision

- Startup template generation: COMPLETED
- Project optimization work: PENDING_USER_CONFIRMATION
- Code implementation: BLOCKED until confirmation

