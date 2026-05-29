# Project Lifecycle Downstream Workflow Architecture

- ID: `architecture_005_project_lifecycle_downstream_workflow`
- Date: `2026-05-29T22:00:59+08:00`
- Branch: `branch_vibe_coding_infra`
- Requirement: `req_infra_009_project_lifecycle_downstream_detail`
- Task: `task_009_project_lifecycle_downstream_detail`

## Purpose

This architecture adds detailed downstream gates after requirement-baseline
confirmation. Its job is to prevent future project work from drifting into
short summaries after requirements are collected.

The downstream workflow starts after a requirement baseline is confirmed and
continues until the project round is accepted, closed, and ready for round
`n+1`.

## Components

| Component | Path | Responsibility |
| --- | --- | --- |
| Runtime rule index | `docs/ai/runtime-rule-index.md` | Declares downstream lifecycle detailing as mandatory after requirement-baseline confirmation. |
| Full-project lifecycle skill | `docs/ai/skills/full-project-lifecycle/SKILL.md` | Calls the downstream detailing workflow for the post-baseline stages. |
| Downstream detail skill | `docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md` | Tells agents how to execute detailed post-baseline gates. |
| Downstream SOP | `docs/ai/templates/project-lifecycle-downstream-gates-sop.md` | Defines gate-by-gate inputs, outputs, confirmations, evidence, and stop conditions. |
| Downstream record template | `docs/ai/templates/project-lifecycle-downstream-record.yaml` | Provides a durable record for research, task analysis, architecture, mode, Agent infra, slices, execution, closure, and next round. |
| Lifecycle record storage guide | `docs/ai/lifecycle/README.md` | Documents where future per-round downstream lifecycle records live. |
| Evidence and handoff | `docs/ai/evidence/index.md`, `docs/ai/handoff/current.md` | Preserve validation and recovery state. |

## Gate Model

1. `research_plan_gate`
2. `current_source_research_execution_gate`
3. `research_synthesis_confirmation_gate`
4. `task_graph_analysis_gate`
5. `task_analysis_confirmation_packet_gate`
6. `architecture_dossier_gate`
7. `architecture_confirmation_packet_gate`
8. `infra_mode_risk_matrix_gate`
9. `infra_mode_confirmation_packet_gate`
10. `agent_driving_infra_plan_gate`
11. `task_tree_slice_contract_gate`
12. `implementation_iteration_ledger_gate`
13. `verification_evidence_packet_gate`
14. `acceptance_closure_packet_gate`
15. `next_round_reentry_gate`

These gates refine the existing lifecycle stages:

- `current_source_research_gate`
- `research_confirmation_gate`
- `task_analysis_gate`
- `task_analysis_confirmation_gate`
- `architecture_and_stack_design_gate`
- `architecture_confirmation_gate`
- `infra_mode_recommendation_gate`
- `infra_mode_confirmation_gate`
- `agent_driving_infra_gate`
- `task_tree_and_slice_gate`
- `implementation_test_docs_state_loop`
- `next_step_protocol_gate`
- `round_closure_gate`
- `next_round_intake[n+1]`

## Persistence Model

Future project rounds should create a downstream lifecycle record from
`docs/ai/templates/project-lifecycle-downstream-record.yaml`, normally at:

```text
docs/ai/lifecycle/<round_id>-downstream.yaml
```

This record is the working surface for:

- Research questions and source matrix.
- Research findings, options, recommendation, and rejected alternatives.
- Task graph, dependency order, critical path, and slice candidates.
- Architecture dossier and requirement-to-design coverage.
- Infrastructure mode risk matrix and confirmation packet.
- Agent driving infrastructure plan.
- Slice contracts and implementation iteration ledger.
- Verification evidence, acceptance packet, closure packet, and next-round
  trigger.

## Completion Rule

Each downstream gate is complete only when its record section, evidence links,
and confirmation status match the gate's requirements. A paragraph summary is
not enough.
