# kb_002: Human Documentation Map

- ID: `kb_002`
- Source: generated project documentation
- Trust level: `high`
- Applies to: `docs/README.md`, `docs/**`
- Version: `0.1.0`
- Updated at: `2026-05-29T17:34:47+08:00`
- Deprecated: `false`

## Fact

Human-facing documentation for this repository starts at `docs/README.md`. The
old root `README.md`, `docs/index.md`, and top-level compatibility entry files
have been folded into the docs tree or removed. Detailed topic docs now live
under structured subdirectories such as `docs/overview`, `docs/structure`,
`docs/environment`, `docs/runtime`, `docs/workflows`, `docs/maintenance`, and
`docs/ai-agent`.

## Evidence

- `docs/README.md`
- `docs/overview/project.md`
- `docs/structure/directory-map.md`
- `docs/environment/configuration.md`
- `docs/runtime/startup-flow.md`
- `docs/workflows/db-cli.md`
- `docs/maintenance/maintenance-guide.md`
- `docs/ai-agent/runtime-state.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- Related requirements: `req_round_001_docs`
- Related tasks: `task_001_human_docs`
- Related skills: `skill_context_recovery`, `skill_slice_execution`
