# kb_004: Task Forest Branching

- ID: `kb_004`
- Source: developer request and implemented task forest runtime artifacts
- Trust level: `high`
- Applies to: `docs/ai/tasks/forest.yaml`, `docs/ai/tasks/main-tree.yaml`, `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- Version: `0.1.0`
- Updated at: `2026-05-29T17:54:03+08:00`
- Deprecated: `false`

## Fact

The repository uses a task forest. Vibe Coding repository infrastructure work is
routed to `branch_vibe_coding_infra`; future product, business, application, or
feature ideas are routed to `branch_mainline_idea` only after confirmation.

## Evidence

- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/main-tree.yaml`
- `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- `docs/ai/analysis/task-forest-branching.md`
- `docs/ai/sandbox/sandbox_001_task_forest_branching.yaml`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- Related requirements: `req_infra_003_task_forest_branching`
- Related tasks: `task_003_task_forest_branching`
- Related skills: `skill_context_recovery`, `skill_slice_execution`

