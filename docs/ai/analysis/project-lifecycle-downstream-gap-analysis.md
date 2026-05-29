# Project Lifecycle Downstream Gap Analysis

- ID: `analysis_005_project_lifecycle_downstream_gap`
- Date: `2026-05-29T22:00:59+08:00`
- Branch: `branch_vibe_coding_infra`
- Related requirement: `req_infra_009_project_lifecycle_downstream_detail`
- Related task: `task_009_project_lifecycle_downstream_detail`

## Diagnosis

The repository now has a detailed requirement discovery workflow for the early
idea-to-baseline stage. The downstream stages after requirement-baseline
confirmation still remain thinner than the requested engineering process.

The full-project lifecycle SOP names research, task analysis, architecture,
mode selection, Agent driving infrastructure, task tree and slices,
implementation, testing, documentation, state sync, acceptance, closure, and
round `n+1`. It also lists required outputs for each stage. That is useful, but
it can still be executed too broadly if a future agent treats each stage as a
short paragraph rather than a physical gate with records, evidence, readiness
checks, and confirmation boundaries.

## Why The Downstream Flow Was Still Too General

| Stage | Existing shape | Remaining gap |
| --- | --- | --- |
| Research | Requires sources, findings, stack candidates, risks, and recommendation. | It does not define a source matrix, freshness policy, decision mapping, rejected option structure, or research record. |
| Task analysis | Requires task candidates, slices, dependencies, tests, docs, and closure criteria. | It does not force a dependency graph, critical path, risk-to-slice mapping, acceptance coverage, or estimation of verification cost. |
| Architecture | Requires overview, stack, data, API, UI, security, rollback, and quality gates. | It does not force an architecture dossier that maps every confirmed requirement to a design decision and unresolved tradeoff. |
| Mode recommendation | Requires a recommended mode and rationale. | It does not force a risk matrix, escalation triggers, artifact depth, or confirmation packet. |
| Agent driving infrastructure | Lists possible outputs. | It does not force a project-specific plan for skills, knowledge, schemas, prompts, evidence, handoff, and quality gates. |
| Task tree and slices | Requires slice fields. | It does not force slice contracts to prove allowed files, forbidden files, verification, rollback, state updates, and next condition before execution. |
| Implementation loop | Lists implementation, tests, docs, evidence, and state sync. | It does not force an iteration ledger, failure triage, retry limits, evidence packet, or invariant that each slice closes before the next expands scope. |
| Acceptance and closure | Lists closure conditions. | It does not force an acceptance packet, residual-risk register, diff summary, handoff update, and explicit next-round trigger. |

## Root Cause

The lifecycle was repaired at the gate-name level and then at the requirement
discovery level. The downstream gates still need their own engineering
operating procedure and durable record shape.

Without that layer, a future agent could correctly name the stages but still
produce short summaries such as:

- "Research confirms the stack."
- "Tasks are split into backend and frontend."
- "Architecture is REST API plus Nuxt."
- "Proceed in standard mode."

Those statements are not enough to safely drive a full project. Each downstream
stage must produce traceable records and a confirmation packet before the next
stage is allowed to consume it.

## Selected Repair Strategy

Add a `project_lifecycle_downstream_detail` workflow under Vibe Coding
infrastructure. It becomes the canonical downstream companion to
`project_requirement_discovery`.

The repair adds:

- A project skill:
  `docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md`.
- A canonical SOP:
  `docs/ai/templates/project-lifecycle-downstream-gates-sop.md`.
- A durable record template:
  `docs/ai/templates/project-lifecycle-downstream-record.yaml`.
- A default record storage guide:
  `docs/ai/lifecycle/README.md`.
- An architecture note:
  `docs/ai/architecture/project-lifecycle-downstream-workflow.md`.
- A knowledge entry:
  `docs/ai/knowledge/entries/kb_010_project_lifecycle_downstream.md`.
- Runtime index, full-project lifecycle, task tree, current slice, evidence,
  status, validator, and handoff updates.

The repair stays inside repository infrastructure. It does not start a future
business project, install dependencies, create package manifests, select a
final project stack, or write application code.
