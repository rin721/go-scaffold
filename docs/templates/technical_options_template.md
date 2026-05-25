# Technical Options Template

## Decision Context

- [CONFIRMED] The project already runs and tests pass.
- [CONFIRMED] The requested direction is optimization planning and boundary consolidation, not immediate code implementation.
- [RISK] A large refactor before documentation and requirements confirmation can create scope expansion and regressions.

## Option A: Conservative Governance First

- Status: Recommended default
- Summary: Establish documents, boundaries, risks, and time slices first; fix confirmed inconsistencies afterward.
- Pros:
  - Low risk.
  - Preserves working code.
  - Matches `docs/ai/prompt.md` requirement to confirm before implementation.
  - Creates a recoverable path for future agents.
- Cons:
  - Slower visible code progress.
  - Requires user review before deeper optimization.
- Best For:
  - Current startup phase.
  - Boundary consolidation.
  - Existing project with passing tests.

## Option B: Modular Refactor First

- Status: Needs confirmation
- Summary: Redesign module registration, migration, routing, and config wiring before broader planning.
- Pros:
  - Can produce a cleaner framework shape.
  - May reduce long-term duplication.
- Cons:
  - Higher regression risk.
  - Requires stronger tests first.
  - Can expand beyond current intake scope.
- Best For:
  - After requirements and architecture are confirmed.

## Option C: Framework Extraction

- Status: Needs confirmation
- Summary: Treat `pkg/*` as stable reusable libraries and turn demo into a template/reference module.
- Pros:
  - Supports long-term reuse.
  - Creates clearer public package contracts.
- Cons:
  - Highest documentation and compatibility cost.
  - Requires versioning policy and API governance.
- Best For:
  - Productizing the scaffold beyond one project.

## Option D: Rewrite

- Status: Not recommended
- Summary: Replace existing structure with a new scaffold design.
- Pros:
  - Maximum freedom.
- Cons:
  - Discards existing working startup chain.
  - Loses current test evidence.
  - High cost and high regression risk.
- Best For:
  - Only if confirmed architecture goals are impossible in the current codebase.

## Recommended Decision

- [INFERRED] Choose Option A for the next phase.
- [NEEDS_CONFIRMATION] User must confirm whether Option A is accepted or whether the project should move toward Option B or C.

## Decision Record Placeholder

- Decision ID: DEC-INTAKE-001
- Decision: [NEEDS_CONFIRMATION]
- Alternatives: Option A, Option B, Option C, Option D
- Consequences: [NEEDS_CONFIRMATION]

