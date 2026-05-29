# Authorization Records

## auth_001: Bootstrap Write Scope

- Date: `2026-05-29T17:18:17+08:00`
- Source: developer message containing the compiler run request.
- Scope:
  - `AGENTS.md`
  - `docs/.cursorrules`
  - `docs/ai/**`
- Excluded:
  - business implementation files
  - dependency installation or lockfile updates
  - production configuration
  - destructive file operations
- Confirmation status: `inferred_for_bootstrap_from_explicit_compiler_run`
- Risk level: `low`
- Next decision: developer must provide and confirm the first `round_001`
  business or project goal before implementation work.

## auth_002: Compiler Runtime Assimilation Scope

- Date: `2026-05-29T21:30:00+08:00`
- Source: developer message requesting deep analysis, distillation, and
  assignment of a Vibe Coding infrastructure compiler specification into the
  current repository, with the clarification that the repository is already a
  Vibe Coding infrastructure repository.
- Scope:
  - `docs/ai/analysis/*`
  - `docs/ai/architecture/*`
  - `docs/ai/skills/*`
  - `docs/ai/templates/*`
  - `docs/ai/knowledge/*`
  - `docs/ai/tasks/*`
  - `docs/ai/status/*`
  - `docs/ai/evidence/*`
  - `docs/ai/scripts/validate_runtime.py`
  - `docs/ai/manifest.yaml`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/changelog.md`
  - `docs/ai/handoff/current.md`
- Excluded:
  - `prompt.md` as a runtime authority
  - business implementation files
  - dependency installation or lockfile updates
  - application skeleton creation
  - production configuration
  - destructive file operations
- Confirmation status: `explicit_assignment_request_for_docs_ai_runtime_artifacts`
- Risk level: `low`
- Next decision: developer can accept the compiler-runtime assimilation
  workflow or request edits.

## auth_003: Strategy C Requirement Workflow Engine Scope

- Date: `2026-05-29T22:24:00+08:00`
- Source: developer confirmation, "确认方案 C，进入阶段二"
- Scope:
  - `docs/ai/requirements/**`
  - `docs/ai/analysis/*`
  - `docs/ai/architecture/*`
  - `docs/ai/skills/*`
  - `docs/ai/templates/*`
  - `docs/ai/knowledge/*`
  - `docs/ai/tasks/*`
  - `docs/ai/status/*`
  - `docs/ai/evidence/*`
  - `docs/ai/scripts/validate_runtime.py`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/changelog.md`
  - `docs/ai/handoff/current.md`
  - `docs/ai/decisions/records.md`
- Excluded:
  - `prompt.md`
  - business implementation files
  - dependency installation or lockfile updates
  - package manifests
  - application skeleton creation
  - production configuration
  - destructive file operations
- Confirmation status: `explicit_confirmed_strategy_C_phase_two`
- Risk level: `low`
- Next decision: developer can accept the strategy C requirement workflow
  engine or request edits.
