# Project Start Template

## 1. Project Goal

- [NEEDS_CONFIRMATION] Problem to solve: Systematically audit and optimize `go-scaffold` by collecting project facts, design boundaries, risks, and a phased optimization route.
- [INFERRED] Target users: Maintainers, future AI agents, and Go developers who use this scaffold.
- [NEEDS_CONFIRMATION] Successful outcome: Clear project boundaries, recoverable documentation, prioritized optimization tasks, and verification evidence for each completed step.

## 2. Users And Scenarios

| User Type | Scenario | Pain Point | Priority | Status |
|---|---|---|---|---|
| Maintainer | Decide what to optimize first | Current design boundaries need consolidation | P0 | [INFERRED] |
| Future AI Agent | Continue after context loss | Missing state docs make next legal task unclear | P0 | [INFERRED] |
| Go Developer | Use scaffold as a baseline | Needs stable extension rules and examples | P1 | [INFERRED] |

## 3. Feature Priority

### P0

- [CONFIRMED] Generate startup, requirements clarification, technical options, architecture constraints, acceptance, and risk confirmation templates.
- [CONFIRMED] Capture current strengths, weaknesses, known risks, and pending confirmation questions.
- [CONFIRMED] Keep code unchanged during intake.

### P1

- [INFERRED] Create focused optimization plans for config, reload, migration, HTTP, demo, testing, and package API stability.
- [INFERRED] Convert confirmed intake into formal requirements and architecture documents.

### P2

- [DEFERRED] Auth/rbac, deployment automation, performance tests, multi-tenancy, plugin module system, and scaffold generator.

### Explicitly Out Of Scope

- [CONFIRMED] Go code changes.
- [CONFIRMED] Refactors.
- [CONFIRMED] New product features.
- [CONFIRMED] Auth/rbac implementation.
- [CONFIRMED] Deployment or irreversible migration commands.

## 4. Technical Preferences

| Item | Preference | Reason | Status |
|---|---|---|---|
| Backend | Go module `github.com/rei0721/go-scaffold` | Existing repository and tests use this module | [CONFIRMED] |
| HTTP | Gin through `pkg/httpserver` and `internal/transport/http` | Existing router/server stack | [CONFIRMED] |
| Database | SQLite local default, MySQL/Postgres supported by package | Existing config and database package | [CONFIRMED] |
| Documentation | Root project docs plus `docs/templates/*` | Required by `docs/ai/prompt.md` | [CONFIRMED] |
| Testing | `go test ./... -count=1` | Existing verification command | [CONFIRMED] |
| Auth | Out of scope for current stage | README says auth/rbac is not implemented | [CONFIRMED] |

## 5. Constraints

| Type | Content | Hard Constraint | Status |
|---|---|---|---|
| Scope | Do not write code before confirmation | Yes | [CONFIRMED] |
| Architecture | Preserve one-way dependency direction | Yes | [INFERRED] |
| Testing | Keep `go test ./... -count=1` passing | Yes | [CONFIRMED] |
| Documentation | Use fact labels for all startup facts | Yes | [CONFIRMED] |
| Safety | Do not deploy, migrate production data, or expose secrets | Yes | [CONFIRMED] |

## 6. Acceptance Draft

| Item | Method | Required | Status |
|---|---|---|---|
| Required startup files exist | File check | Yes | [CONFIRMED] |
| Facts use required labels | Manual review | Yes | [CONFIRMED] |
| Next step is confirmation | `STATUS.md` review | Yes | [CONFIRMED] |
| Go behavior unchanged | `git diff` and optional `go test ./... -count=1` | Yes | [INFERRED] |

## 7. Risk Draft

| Risk | Impact | Probability | Mitigation | Blocking |
|---|---|---|---|---|
| Scope expands into full rewrite | High | Medium | Keep P0/P1/P2 boundaries and require confirmation | Yes |
| Documentation drift | Medium | High | Record facts in root docs and templates | Yes |
| Test gaps hide regressions | High | Medium | Add test strategy in requirements phase | No |
| Migration boundary unclear | High | Medium | Confirm migration policy before code changes | Yes |
| Public package compatibility unclear | Medium | Medium | Confirm `pkg/*` API policy | Yes |

## 8. AI Agent Driving Rules

- Developer deep involvement stage: Project intake, requirements confirmation, architecture confirmation.
- Trigger after developer exits: Only continue when `STATUS.md` has a unique legal task.
- Meaning of `next step`: Execute the current legal time slice only.
- Allow automatic dependency install: [NEEDS_CONFIRMATION]
- Allow automatic architecture doc updates: Yes, after user confirms scope.
- Allow automatic backlog additions: Yes, as `[DEFERRED]`.
- Allow automatic git commit: [NEEDS_CONFIRMATION]
- Allow automatic tests: Yes, for non-destructive checks.
- Allow automatic deployment: No.

## 9. Pending Questions

### Must Confirm

1. [NEEDS_CONFIRMATION] Which technical option should be adopted as the default optimization route?
2. [NEEDS_CONFIRMATION] Are `pkg/*` packages public APIs or internal support packages?
3. [NEEDS_CONFIRMATION] What migration strategy should be authoritative?

### Can Default But Should Confirm

1. [INFERRED] Keep conservative governance-first optimization as default.
2. [INFERRED] Keep current high-level module layout.
3. [INFERRED] Move auth/rbac and JWT details to backlog unless explicitly promoted.

### Can Defer

1. [DEFERRED] Deployment pipeline.
2. [DEFERRED] Performance testing.
3. [DEFERRED] Multi-tenancy.
4. [DEFERRED] Plugin-style module system.

