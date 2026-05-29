# kb_001: Vibe Runtime Infrastructure

- ID: `kb_001`
- Source: developer compiler run request and generated runtime artifacts
- Trust level: `high`
- Applies to: `AGENTS.md`, `docs/ai/*`
- Version: `0.1.0`
- Updated at: `2026-05-29T17:18:17+08:00`
- Deprecated: `false`

## Fact

This repository should recover agent work from physical runtime artifacts:
`AGENTS.md`, `docs/ai/status/current.yaml`, `docs/ai/tasks/current-slice.yaml`,
`docs/ai/requirements/ledger.yaml`, `docs/ai/evidence/index.md`,
`docs/ai/knowledge/index.md`, `docs/ai/skills/index.md`, and
`docs/ai/handoff/current.md`.

## Evidence

- `docs/ai/runtime-rule-index.md`
- `docs/ai/tasks/bootstrap-tree.yaml`
- `docs/ai/evidence/index.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- Related requirements: `req_bootstrap_001`
- Related tasks: `task_bootstrap_001`
- Related skills: `skill_context_recovery`, `skill_slice_execution`

