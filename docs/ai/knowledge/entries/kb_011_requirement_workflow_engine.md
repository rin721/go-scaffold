# kb_011: Requirement Workflow Engine

- ID: `kb_011`
- Source: `req_infra_010_requirement_workflow_engine`
- Trust level: `high`
- Applies to: `requirement workflow engine`, `project requirement discovery`
- Updated at: `2026-05-29T22:24:00+08:00`
- Sensitive info check: `passed`
- Prompt injection check: `not_applicable`
- Deprecated: `false`

## Fact

Strategy C is implemented as a declarative requirement workflow engine under
`docs/ai/requirements/`. It includes a workflow node graph, a
`REQUIREMENTS_GATHERING` state lock, and reusable probe templates for commerce,
SaaS, blog, and generic app ideas.

The engine locks business code generation, package manifests, dependency
installation, and application skeleton creation while requirements are being
gathered. During that state, agents may only update requirement intake records,
ledger entries, tasks, status, evidence, decisions, and handoff artifacts.

## Evidence

- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`
- `docs/ai/skills/requirement-workflow-engine/SKILL.md`
- `docs/ai/templates/requirement-workflow-engine-sop.md`
- `evidence_023`

## Related

- Requirement: `req_infra_010_requirement_workflow_engine`
- Task: `task_010_requirement_workflow_engine`
- Skill: `skill_requirement_workflow_engine`
