# kb_007: Full Project Lifecycle Workflow

- ID: `kb_007`
- Source: developer request to repair the mainline Vibe Coding flow so ideas
  are driven through a complete project lifecycle instead of a summarized
  requirement-only path
- Trust level: `high`
- Applies to: `docs/ai/skills/full-project-lifecycle/SKILL.md`,
  `docs/ai/templates/full-project-lifecycle-workflow-sop.md`,
  `docs/ai/runtime-rule-index.md`
- Version: `0.1.0`
- Updated at: `2026-05-29T22:24:00+08:00`
- Deprecated: `false`

## Fact

The repository now has a dedicated full-project lifecycle workflow for future
mainline or other confirmed project task lines. The workflow starts with a
task-line routing gate, then requires idea intake, requirement analysis,
requirement confirmation, current-source research and technology validation,
research confirmation, task analysis, task-analysis confirmation, architecture
and stack design, architecture confirmation, infrastructure mode recommendation
and confirmation, Agent Vibe Coding driving infrastructure, task tree and slice
creation, implementation, tests, documentation, state sync, evidence, next-step
continuation, task-tree closure, and round `n+1` intake.

The workflow explicitly distinguishes the two infrastructure modes:
`standard_light_risk_escalated` and `enterprise_high_assurance`. The agent may
recommend a mode and explain the risk rationale, but the developer must confirm
the mode before implementation work proceeds.

The lifecycle now delegates raw idea elaboration to the requirement workflow
engine and `skill_project_requirement_discovery` before requirement-baseline
confirmation. That sub-procedure enters `REQUIREMENTS_GATHERING`, locks
business code generation, selects a probe template, and requires idea
interpretation, domain surface mapping, requirement inventory, numbered
question batches, question backlog creation, collection rounds, follow-up
triggers, requirement planning, durable intake records, and baseline readiness
checks.

After requirement-baseline confirmation, the lifecycle delegates post-baseline
execution to `skill_project_lifecycle_downstream_detailing`. That sub-procedure
requires research planning, current-source research execution, research
confirmation packet, task graph analysis, task-analysis confirmation packet,
architecture dossier, architecture confirmation packet, mode risk matrix, mode
confirmation packet, Agent driving infrastructure plan, slice contracts,
implementation iteration ledger, verification evidence packet, acceptance
closure packet, and next-round reentry state.

## Evidence

- `docs/ai/analysis/mainline-full-project-lifecycle-gap-analysis.md`
- `docs/ai/architecture/full-project-lifecycle-workflow.md`
- `docs/ai/skills/full-project-lifecycle/SKILL.md`
- `docs/ai/templates/full-project-lifecycle-workflow-sop.md`
- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`
- `docs/ai/skills/requirement-workflow-engine/SKILL.md`
- `docs/ai/templates/requirement-workflow-engine-sop.md`
- `docs/ai/templates/project-requirement-discovery-sop.md`
- `docs/ai/templates/project-lifecycle-downstream-gates-sop.md`
- `docs/ai/evidence/index.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- No prompt runtime dependency check: `passed`
- Related requirements: `req_infra_006_full_project_lifecycle`
- Related tasks: `task_006_full_project_lifecycle`
- Related skills: `skill_full_project_lifecycle`, `skill_requirement_workflow_engine`, `skill_project_requirement_discovery`, `skill_project_lifecycle_downstream_detailing`
