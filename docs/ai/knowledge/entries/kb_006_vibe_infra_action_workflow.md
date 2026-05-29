# kb_006: Vibe Infrastructure Action Workflow SOP

- ID: `kb_006`
- Source: developer request for a trigger-bound, two-stage SOP template for
  Vibe Coding repository infrastructure establishment tasks
- Trust level: `high`
- Applies to: `docs/ai/skills/vibe-infra-action-workflow/SKILL.md`,
  `docs/ai/templates/vibe-infra-action-workflow-sop.md`
- Version: `0.1.0`
- Updated at: `2026-05-29T19:32:05+08:00`
- Deprecated: `false`

## Fact

The repository has a project skill for inputs that start with
`新增 Vibe Coding 仓库基建设施建立任务线：`. The text after the full-width colon is
bound to `{{ 优化动作 }}`.

The skill enforces an isolated two-stage workflow. Phase one performs only
deep scanning, diagnosis, option selection, and a red-light confirmation stop.
It forbids code generation, directory creation, file writes, dependency
installation, and physical output. Phase two can begin only after explicit
human authorization and must self-check the confirmed strategy before writing
runtime artifacts.

## Evidence

- `docs/ai/skills/vibe-infra-action-workflow/SKILL.md`
- `docs/ai/templates/vibe-infra-action-workflow-sop.md`
- `docs/ai/evidence/index.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- No prompt runtime dependency check: `passed`
- Related requirements: `req_infra_005_vibe_infra_action_workflow`
- Related tasks: `task_005_vibe_infra_action_workflow`
- Related skills: `skill_vibe_infra_action_workflow`
