# Project Requirement Discovery Gap Analysis

- ID: `analysis_004_project_requirement_discovery_gap`
- Date: `2026-05-29T21:48:21+08:00`
- Branch: `branch_vibe_coding_infra`
- Related requirement: `req_infra_008_project_requirement_discovery`
- Related task: `task_008_project_requirement_discovery`

## Diagnosis

The repository already has a full-project lifecycle workflow, but the earliest
part of that lifecycle can still be executed too shallowly. A compact developer
idea such as "build a Go + Nuxt separated-start commerce system" can be turned
too quickly into a suggested MVP, a few assumed features, and a single
confirmation sentence.

That behavior is helpful as a first orientation, but it does not fully guide
the developer through the engineering work of defining a complete project. It
can skip the discovery work that must happen before a requirement baseline is
ready: understanding the idea, identifying the system domain, inventorying
requirements, planning requirement collection, persisting answers, and asking
targeted follow-up questions across multiple rounds.

## Example Failure Mode

For the Go + Nuxt commerce seed, the thin response pattern is:

- Accept the technology framing too early.
- Infer an MVP feature list before the commerce domain is explored.
- Mention requirement confirmation, research, task analysis, architecture, and
  mode confirmation, but not run a requirement-discovery process.
- Ask for one acceptance sentence instead of collecting the missing product,
  user, data, checkout, admin, security, integration, quality, and operation
  requirements.
- Leave inferred requirements in chat instead of a persistent intake record
  linked to the requirement ledger.

## Root Cause

The current lifecycle SOP names `idea_intake_gate` and
`requirement_analysis_gate`, but those gates do not yet contain a detailed
requirement discovery sub-workflow. The missing layer is the one between a raw
idea and a confirmable requirement baseline.

That layer needs to define:

- How to parse the idea without over-committing to assumptions.
- Which requirement domains must be scanned for a complete project.
- How to separate developer-confirmed requirements from AI-inferred candidates.
- How to produce a prioritized question backlog instead of a generic prompt.
- How to persist answers and open questions between turns.
- How to decide when the requirement baseline is ready for hard confirmation.

## Selected Repair Strategy

Add a dedicated `project_requirement_discovery` workflow under Vibe Coding
infrastructure. It becomes a required sub-procedure inside the full-project
lifecycle before requirement-baseline confirmation.

The repair adds:

- A project skill:
  `docs/ai/skills/project-requirement-discovery/SKILL.md`.
- A canonical SOP:
  `docs/ai/templates/project-requirement-discovery-sop.md`.
- A durable intake record template:
  `docs/ai/templates/project-requirement-discovery-record.yaml`.
- An architecture note:
  `docs/ai/architecture/project-requirement-discovery-workflow.md`.
- A knowledge entry:
  `docs/ai/knowledge/entries/kb_009_project_requirement_discovery.md`.
- Runtime index, task tree, current slice, evidence, status, validator, and
  handoff updates.

The repair does not start the Go + Nuxt commerce project, choose its final
stack, create package manifests, install dependencies, or write business code.
It strengthens how future mainline ideas are elicited and persisted.
