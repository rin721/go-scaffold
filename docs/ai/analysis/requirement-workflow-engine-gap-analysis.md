# Requirement Workflow Engine Gap Analysis

- ID: `analysis_006_requirement_workflow_engine_gap`
- Date: `2026-05-29T22:24:00+08:00`
- Branch: `branch_vibe_coding_infra`
- Related requirement: `req_infra_010_requirement_workflow_engine`
- Related task: `task_010_requirement_workflow_engine`

## Diagnosis

The repository already has project requirement discovery and downstream
lifecycle detailing. That repair prevents a raw idea from becoming a generic
MVP summary, but it still relies on agents reading a prose SOP and choosing
reasonable questions.

The developer explicitly selected strategy C: a complete workflow engine. In
this repository that should not mean an executable framework, dependency
installation, package manifests, or application scaffolding. It should mean a
declarative runtime control plane under `docs/ai/requirements/` that future
agents can inspect and follow deterministically.

## Remaining Gaps Before This Repair

| Area | Existing behavior | Gap |
| --- | --- | --- |
| Project-type probes | The requirement discovery SOP lists general domains and a commerce example. | There is no reusable `template_discovery.yaml` catalog with commerce, SaaS, blog, and generic project probes. |
| Write lock | SOPs say not to generate code before confirmation. | There is no machine-readable `REQUIREMENTS_GATHERING` state that explicitly sets code generation and package creation to false. |
| Engine topology | Full-project lifecycle and requirement discovery list gates. | There is no single declarative engine record tying routing, template selection, intake record creation, question loops, follow-up triggers, coverage checks, and baseline readiness together. |
| Follow-up mechanics | The SOP says to ask follow-ups. | There are no explicit trigger rules such as SKU answers producing SKU dimension, stock deduction, and order snapshot questions. |
| Validation | The validator checks discovery and downstream terms. | It does not yet require the new engine artifacts, state lock terms, or template catalog. |

## Selected Repair Strategy

Implement strategy C as a declarative requirement workflow engine:

- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`
- `docs/ai/skills/requirement-workflow-engine/SKILL.md`
- `docs/ai/templates/requirement-workflow-engine-sop.md`
- schema, validator, runtime index, manifest, task tree, evidence, knowledge,
  and handoff synchronization.

The repair remains repository infrastructure. It does not start the Go + Nuxt
commerce project, select a final business stack, create package manifests,
install dependencies, or write application code.
