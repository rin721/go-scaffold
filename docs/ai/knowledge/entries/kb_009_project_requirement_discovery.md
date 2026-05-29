# kb_009: Project Requirement Discovery

- ID: `kb_009`
- Source: developer request to refine the idea-to-requirement flow so agents
  persistently collect and guide complete project requirements instead of
  summarizing an idea into a shallow MVP
- Trust level: `high`
- Applies to: `docs/ai/skills/project-requirement-discovery/SKILL.md`,
  `docs/ai/templates/project-requirement-discovery-sop.md`,
  `docs/ai/templates/project-requirement-discovery-record.yaml`,
  `docs/ai/requirements/intake/README.md`,
  `docs/ai/skills/full-project-lifecycle/SKILL.md`
- Version: `0.1.0`
- Updated at: `2026-05-29T21:48:21+08:00`
- Deprecated: `false`

## Fact

The repository now has a dedicated `project_requirement_discovery` workflow.
Future mainline project ideas must pass through idea seed intake, reversible
idea interpretation, domain surface mapping, requirement inventory, question
backlog creation, repeated requirement collection rounds, requirement planning,
persistence sync, and baseline readiness checks before requirement-baseline
confirmation.

The workflow distinguishes developer-stated facts from AI-inferred candidates.
It persists discovery state in a per-round intake record and synchronizes only
confirmed, rejected, deferred, out-of-scope, or research-needed items into the
requirement ledger. It is designed to guide the developer toward more complete
requirements through small, targeted question batches.

The discovery workflow is now backed by the requirement workflow engine:
`docs/ai/requirements/workflow_engine.yaml`, the probe template catalog
`docs/ai/requirements/template_discovery.yaml`, and the
`REQUIREMENTS_GATHERING` state lock in
`docs/ai/requirements/state_machine.yaml`.

## Evidence

- `docs/ai/analysis/project-requirement-discovery-gap-analysis.md`
- `docs/ai/architecture/project-requirement-discovery-workflow.md`
- `docs/ai/skills/project-requirement-discovery/SKILL.md`
- `docs/ai/templates/project-requirement-discovery-sop.md`
- `docs/ai/templates/project-requirement-discovery-record.yaml`
- `docs/ai/requirements/intake/README.md`
- `docs/ai/evidence/index.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- No prompt runtime dependency check: `passed`
- Related requirements: `req_infra_008_project_requirement_discovery`
- Related tasks: `task_008_project_requirement_discovery`
- Related skills: `skill_project_requirement_discovery`
