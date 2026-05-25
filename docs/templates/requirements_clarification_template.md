# Requirements Clarification Template

## Purpose

Clarify the optimization requirements for `go-scaffold` before creating formal `REQUIREMENTS.md`.

## Current Facts

- [CONFIRMED] The project is an existing Go scaffold with a working demo CRUD module.
- [CONFIRMED] The current stage does not implement auth/rbac.
- [CONFIRMED] `go test ./... -count=1` is the known full test command and has passed during intake analysis.
- [RISK] Documentation and examples may drift from current scope.

## Must Confirm

| ID | Question | Why It Matters | Options | Status |
|---|---|---|---|---|
| REQ-Q001 | Optimization direction | Determines roadmap size and risk | Conservative governance; modular refactor; framework extraction | [NEEDS_CONFIRMATION] |
| REQ-Q002 | `pkg/*` API policy | Determines compatibility rules | Public reusable libraries; internal scaffold support; mixed | [NEEDS_CONFIRMATION] |
| REQ-Q003 | Demo module role | Determines module and test strategy | Canonical example; temporary placeholder; remove later | [NEEDS_CONFIRMATION] |
| REQ-Q004 | Migration policy | Determines database initialization boundary | AutoMigrate; SQL scripts; dev/prod layered strategy | [NEEDS_CONFIRMATION] |
| REQ-Q005 | Auth/JWT scope | Resolves README and `.env.example` drift | Defer; keep placeholder; promote into roadmap | [NEEDS_CONFIRMATION] |

## Defaultable Requirements

- [INFERRED] Prefer documentation and boundary consolidation before code-level changes.
- [INFERRED] Preserve the current `cmd -> internal/app -> internal/transport + internal/modules -> pkg` structure unless a later architecture decision changes it.
- [INFERRED] Treat unconfirmed improvements as backlog items.

## Deferred Requirements

- [DEFERRED] Auth/rbac implementation.
- [DEFERRED] Deployment automation.
- [DEFERRED] Performance benchmark suite.
- [DEFERRED] Multi-tenancy.
- [DEFERRED] Plugin/module marketplace or generator.

## Draft P0 Requirements

| Requirement ID | Requirement | Acceptance Signal | Status |
|---|---|---|---|
| REQ-P0-001 | Create project startup documents | Required files exist and use fact labels | [CONFIRMED] |
| REQ-P0-002 | Capture current advantages and weaknesses | `PROJECT_BRIEF.md` includes strengths, weaknesses, and risks | [CONFIRMED] |
| REQ-P0-003 | Prevent code changes before confirmation | `STATUS.md` marks implementation as blocked | [CONFIRMED] |
| REQ-P0-004 | Identify design boundary decisions | Pending confirmation questions are listed | [CONFIRMED] |

## Draft Non-Functional Requirements

- [INFERRED] Maintainability: Every optimization task must name the affected boundary.
- [INFERRED] Testability: Every P0 code-level follow-up must include a verification command.
- [INFERRED] Recoverability: Project status must be restorable from repository documents.
- [INFERRED] Safety: No deployment, secret exposure, or irreversible migration is allowed without explicit confirmation.

## Confirmation Output Expected

After user confirmation, generate:

- `REQUIREMENTS.md`
- `ACCEPTANCE.md`
- `BACKLOG.md`
- `RISK_REGISTER.md`
- updated `STATUS.md`

