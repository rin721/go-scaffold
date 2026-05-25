# REQUIREMENTS.md

## Requirements Status

- Project: go-scaffold
- Phase: Requirements Confirmation
- Status: PENDING_USER_CONFIRMATION
- Last Updated: 2026-05-25
- Input: `PROJECT_BRIEF.md`, `STATUS.md`, user command `下一步`

## Confirmation Assumption

- [INFERRED] The user's `下一步` means the startup package is accepted well enough to generate formal requirements using default assumptions.
- [INFERRED] Default route is Option A: conservative governance-first optimization.
- [NEEDS_CONFIRMATION] The user still needs to confirm final requirements before architecture design or code implementation.

## Project Goal

- [CONFIRMED] Establish a recoverable, document-driven optimization process for the existing `go-scaffold` repository.
- [INFERRED] Optimize by first consolidating facts, requirements, boundaries, risks, and verification expectations before touching code.
- [NEEDS_CONFIRMATION] Later phases will decide whether to keep a conservative route or move into modular refactor/framework extraction.

## P0 Requirements

| ID | Requirement | Acceptance | Status |
|---|---|---|---|
| REQ-P0-001 | Preserve current working behavior while documenting optimization goals | `go test ./... -count=1` passes; no Go code changed in requirements phase | [CONFIRMED] |
| REQ-P0-002 | Define current project facts and scope boundaries | `PROJECT_BRIEF.md` and `REQUIREMENTS.md` contain fact-labeled scope | [CONFIRMED] |
| REQ-P0-003 | Keep auth/rbac out of implementation scope until explicitly confirmed | Requirements and backlog mark auth/rbac as deferred unless promoted | [CONFIRMED] |
| REQ-P0-004 | Capture and track risks that affect architecture or execution | `RISK_REGISTER.md` contains blocking and non-blocking risks | [CONFIRMED] |
| REQ-P0-005 | Define acceptance criteria before implementation | `ACCEPTANCE.md` contains project, requirements, and future implementation gates | [CONFIRMED] |
| REQ-P0-006 | Prevent unbounded optimization and refactor drift | `BACKLOG.md` captures deferred ideas; `STATUS.md` identifies only the next legal phase | [CONFIRMED] |

## P1 Requirements

| ID | Requirement | Acceptance | Status |
|---|---|---|---|
| REQ-P1-001 | Confirm architecture boundaries for `cmd`, `internal/app`, `internal/modules`, `internal/transport`, and `pkg` | Future `ARCHITECTURE.md` defines responsibilities and non-goals | [NEEDS_CONFIRMATION] |
| REQ-P1-002 | Decide whether `pkg/*` is public API, internal support, or mixed policy | Future `DECISIONS.md` records compatibility policy | [NEEDS_CONFIRMATION] |
| REQ-P1-003 | Confirm demo module role | Future `ARCHITECTURE.md` states whether demo is canonical, temporary, or removable | [NEEDS_CONFIRMATION] |
| REQ-P1-004 | Confirm migration strategy | Future `ARCHITECTURE.md` defines `AutoMigrate`, `initdb`, and SQL script boundaries | [NEEDS_CONFIRMATION] |
| REQ-P1-005 | Define minimum test matrix before code optimization | Future `TEST_REPORT.md` and task docs list unit/integration/smoke commands | [NEEDS_CONFIRMATION] |

## P2 Requirements

| ID | Requirement | Acceptance | Status |
|---|---|---|---|
| REQ-P2-001 | Add auth/rbac after scope confirmation | Backlog item promoted by user decision | [DEFERRED] |
| REQ-P2-002 | Add deployment or CI quality gates | Backlog item promoted by user decision | [DEFERRED] |
| REQ-P2-003 | Add scaffold generator or plugin-style module system | Backlog item promoted by architecture decision | [DEFERRED] |
| REQ-P2-004 | Add performance benchmark suite | Backlog item promoted after functional boundaries stabilize | [DEFERRED] |

## Explicit Non-Requirements

- [CONFIRMED] Do not implement Go code during requirements confirmation.
- [CONFIRMED] Do not refactor existing packages during requirements confirmation.
- [CONFIRMED] Do not introduce auth/rbac as implementation scope without user confirmation.
- [CONFIRMED] Do not run production deployment, irreversible migrations, or secret-dependent commands.
- [INFERRED] Do not treat package README examples as confirmed architecture until reconciled with code and root docs.

## Functional Requirements For The Optimization Process

| ID | Requirement | Priority | Status |
|---|---|---|---|
| F-001 | Maintain a single current legal task in `STATUS.md` | P0 | [CONFIRMED] |
| F-002 | Use fact labels in project documents | P0 | [CONFIRMED] |
| F-003 | Convert new ideas into backlog unless they block current acceptance | P0 | [CONFIRMED] |
| F-004 | Require user confirmation before architecture or code-level changes | P0 | [CONFIRMED] |
| F-005 | Record verification commands and results after each phase | P0 | [CONFIRMED] |

## Non-Functional Requirements

| Area | Requirement | Status |
|---|---|---|
| Maintainability | Each future task must name the module or boundary it affects | [CONFIRMED] |
| Testability | Each future implementation task must include a test or verification command | [CONFIRMED] |
| Recoverability | A new agent must be able to recover status from repository documents | [CONFIRMED] |
| Safety | Secrets, production deployments, and irreversible migrations require explicit confirmation | [CONFIRMED] |
| Scope Control | Unconfirmed optimizations must stay in `BACKLOG.md` | [CONFIRMED] |

## Open Requirement Questions

| ID | Question | Impact | Default | Status |
|---|---|---|---|---|
| Q-REQ-001 | Should Option A remain the official route? | Drives architecture and tasks | Yes, conservative governance first | [NEEDS_CONFIRMATION] |
| Q-REQ-002 | Are `pkg/*` packages public APIs? | Drives compatibility policy | Mixed until confirmed | [NEEDS_CONFIRMATION] |
| Q-REQ-003 | What is the demo module's long-term role? | Drives tests and examples | Canonical demo until confirmed | [NEEDS_CONFIRMATION] |
| Q-REQ-004 | What migration strategy is authoritative? | Drives DB init architecture | Layered dev/prod policy until confirmed | [NEEDS_CONFIRMATION] |
| Q-REQ-005 | Should JWT examples be removed, deferred, or promoted? | Resolves doc drift | Defer auth/JWT to backlog | [NEEDS_CONFIRMATION] |

## Completion Decision

- Requirements document generated: COMPLETED
- Requirements phase overall: PENDING_USER_CONFIRMATION
- Next legal phase after user confirmation: Architecture Confirmation

