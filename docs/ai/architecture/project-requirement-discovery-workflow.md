# Project Requirement Discovery Workflow Architecture

- ID: `architecture_004_project_requirement_discovery_workflow`
- Date: `2026-05-29T21:48:21+08:00`
- Branch: `branch_vibe_coding_infra`
- Requirement: `req_infra_008_project_requirement_discovery`
- Task: `task_008_project_requirement_discovery`

## Purpose

This architecture adds a concrete requirement-discovery layer to the
full-project lifecycle. A future product idea must be decomposed into durable
requirement artifacts before the agent asks for a requirement-baseline
confirmation or proceeds to technology research, task analysis, architecture,
or implementation.

## Components

| Component | Path | Responsibility |
| --- | --- | --- |
| Runtime rule index | `docs/ai/runtime-rule-index.md` | Declares requirement discovery as a required project-lifecycle sub-procedure. |
| Project skill | `docs/ai/skills/project-requirement-discovery/SKILL.md` | Tells agents when and how to run project requirement discovery. |
| SOP template | `docs/ai/templates/project-requirement-discovery-sop.md` | Defines the gates, outputs, question batches, persistence rules, and readiness checks. |
| Intake record template | `docs/ai/templates/project-requirement-discovery-record.yaml` | Provides the durable per-round structure for idea interpretation, requirement domains, questions, answers, and baseline readiness. |
| Requirement workflow engine | `docs/ai/requirements/workflow_engine.yaml` | Declares the full requirement-discovery node graph, loop controls, outputs, and permission invariants. |
| Probe template catalog | `docs/ai/requirements/template_discovery.yaml` | Provides reusable commerce, SaaS, blog, and generic app question templates and follow-up triggers. |
| Requirement state machine | `docs/ai/requirements/state_machine.yaml` | Locks `REQUIREMENTS_GATHERING` so code generation, manifests, dependencies, and app skeletons are forbidden while requirements are open. |
| Intake directory guide | `docs/ai/requirements/intake/README.md` | Documents the default storage location for future per-round discovery records. |
| Requirement ledger | `docs/ai/requirements/ledger.yaml` | Stores confirmed, pending, deferred, rejected, or research-needed requirements after discovery sync. |
| Full-project lifecycle skill | `docs/ai/skills/full-project-lifecycle/SKILL.md` | Calls requirement discovery before baseline confirmation. |
| Evidence and handoff | `docs/ai/evidence/index.md`, `docs/ai/handoff/current.md` | Preserve validation and recovery state. |

## Gate Model

1. `idea_seed_intake_gate`
2. `idea_interpretation_gate`
3. `domain_surface_mapping_gate`
4. `requirement_inventory_gate`
5. `question_backlog_gate`
6. `requirement_collection_round_gate`
7. `requirement_plan_gate`
8. `requirement_persistence_sync_gate`
9. `baseline_readiness_gate`

These gates sit inside the full-project lifecycle between `idea_intake_gate`
and `requirement_baseline_confirmation_gate`.

## Persistence Model

Future mainline rounds should create a discovery record from
`docs/ai/templates/project-requirement-discovery-record.yaml`, normally at:

```text
docs/ai/requirements/intake/<round_id>.yaml
```

The record is the working surface for:

- Raw developer idea.
- Agent interpretation and assumptions.
- Requirement domain scan.
- Candidate requirements.
- Open questions.
- Developer answers.
- Deferred or out-of-scope items.
- Research-needed items.
- Baseline readiness.

Confirmed or explicitly rejected requirement facts are then synchronized into
`docs/ai/requirements/ledger.yaml` with evidence references. AI-inferred
details stay marked as candidates until the developer confirms them.

## Go + Nuxt Commerce Seed Implication

For a future "Go + Nuxt separated-start commerce system" idea, the agent must
not stop at a generic MVP list. It must scan at least:

- Business goal and target users.
- Buyer, admin, operator, and optional merchant roles.
- Product catalog and inventory expectations.
- Cart, checkout, order, payment, fulfillment, and refund boundaries.
- User accounts, auth, permissions, and privacy assumptions.
- Admin workflows and audit needs.
- Frontend/backend startup, API boundary, and local development topology.
- Data storage, seed data, migrations, and backup expectations.
- Non-functional requirements: performance, accessibility, security,
  observability, testability, deployment, and recovery.
- Explicit non-goals for the first round.

Only after enough answers are persisted can the lifecycle move to current-source
research and stack confirmation.
