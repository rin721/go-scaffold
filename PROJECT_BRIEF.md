# PROJECT_BRIEF.md

## Project Identity

- Project: go-scaffold
- Current Phase: Project Intake
- Overall Status: PENDING_USER_CONFIRMATION
- Last Updated: 2026-05-25
- Source Rule: docs/ai/prompt.md

## Summary

- [CONFIRMED] The repository is an existing Go scaffold project, not a greenfield project.
- [CONFIRMED] The current project contains `cmd/server`, `internal/app`, `internal/modules/demo`, reusable `pkg/*` infrastructure packages, and a demo Todo CRUD API.
- [CONFIRMED] `README.md` states that the current stage keeps only the infrastructure startup chain and one demo CRUD example, and does not implement auth/rbac.
- [CONFIRMED] `go test ./... -count=1` passed during startup analysis.
- [RISK] `.env.example` contains JWT-related sample configuration while README says auth/rbac is out of scope, creating scope/documentation drift.
- [RISK] The `docs/` directory is currently untracked in git, so project facts are not yet stable until these documents are added to version control.
- [INFERRED] The requested work is documentation and project-governance startup work, not Go code implementation.

## One-Sentence Goal

- [NEEDS_CONFIRMATION] Systematically audit `go-scaffold` design, implementation, documentation, tests, and boundaries, then create a phased optimization roadmap that can be executed by humans or AI agents without relying on chat history.

## Target Users

| User Type | Scenario | Pain Point | Priority | Status |
|---|---|---|---|---|
| Maintainer | Maintain and evolve the scaffold safely | Module and design boundaries can drift over time | P0 | [INFERRED] |
| Future AI Agent | Continue work after context loss | Missing state files make the next legal task ambiguous | P0 | [INFERRED] |
| Go Developer | Use scaffold as project baseline | Needs clear extension rules and reliable startup/test path | P1 | [INFERRED] |

## Success Criteria

- [NEEDS_CONFIRMATION] Project boundaries are unified and documented.
- [NEEDS_CONFIRMATION] Optimization work is split into phases and time slices.
- [NEEDS_CONFIRMATION] Code-level changes do not start until startup documents and user confirmations are complete.
- [CONFIRMED] The known verification command is `go test ./... -count=1`.

## Priority Scope

### P0

- [CONFIRMED] Establish project facts in repository documents.
- [CONFIRMED] Capture current strengths, weaknesses, and risks.
- [CONFIRMED] Create startup, clarification, technical option, architecture constraint, acceptance, and risk confirmation templates.
- [CONFIRMED] Set project status to `PENDING_USER_CONFIRMATION`.

### P1

- [INFERRED] Create focused follow-up plans for config, reload, migration, HTTP, demo module, testing, and package API stability.
- [INFERRED] Convert confirmed P0 findings into `REQUIREMENTS.md`, `ARCHITECTURE.md`, `ROADMAP.md`, `TASKS.md`, and `TIME_SLICES.md`.

### P2

- [DEFERRED] Add quality gates, scaffold generators, auth/rbac, deployment pipelines, performance tests, multi-tenancy, and plugin-style module systems.

## Explicit Non-Goals

- [CONFIRMED] Do not write Go code in this startup phase.
- [CONFIRMED] Do not refactor existing modules in this startup phase.
- [CONFIRMED] Do not add new product features in this startup phase.
- [CONFIRMED] Do not implement auth/rbac before explicit confirmation.
- [CONFIRMED] Do not run deployment or irreversible migration commands.

## Current Strengths

- [CONFIRMED] The project already has a working startup command documented in README.
- [CONFIRMED] The project has a simple demo module with handler, service, repository, and model layers.
- [CONFIRMED] Infrastructure packages exist for database, logger, HTTP server, cache, i18n, storage, executor, sqlgen, utils, and yaml2go.
- [CONFIRMED] The full Go test suite passed during startup analysis.
- [INFERRED] The current composition root in `internal/app` provides a useful place to centralize wiring and enforce boundaries.

## Current Weaknesses

- [RISK] Core project-management documents required by `docs/ai/prompt.md` were missing before this startup work.
- [RISK] README and `.env.example` have a possible scope mismatch around auth/JWT.
- [RISK] Current tests cover several packages but do not yet cover the full app/router/demo/config reload path.
- [RISK] Migration responsibilities need confirmation because demo `AutoMigrate`, `initdb`, and SQL scripts coexist.
- [RISK] Some `pkg/sqlgen` APIs advertise or contain unimplemented capabilities that need backlog or explicit unsupported boundaries.

## Recommended Default

- [INFERRED] Use Option A: conservative governance first.
- [INFERRED] Keep the current high-level structure: `cmd -> internal/app -> internal/transport + internal/modules -> pkg`.
- [INFERRED] Put all new optimization ideas into backlog until a confirmed time slice authorizes them.

## Pending User Confirmations

1. [NEEDS_CONFIRMATION] Should optimization remain conservative, or prepare for a larger framework-style refactor?
2. [NEEDS_CONFIRMATION] Are `pkg/*` packages public reusable libraries or internal scaffold support packages?
3. [NEEDS_CONFIRMATION] Is `internal/modules/demo` a long-term canonical example or a temporary placeholder?
4. [NEEDS_CONFIRMATION] Should migration policy prefer `AutoMigrate`, SQL scripts, or a layered strategy?
5. [NEEDS_CONFIRMATION] Should JWT/auth examples be removed from startup docs until auth/rbac is explicitly in scope?

