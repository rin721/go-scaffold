# Project Lifecycle Downstream Detailing Skill

## Trigger Conditions

Use this skill when any of these are true:

- A future mainline or project task line has a confirmed requirement baseline
  and needs to proceed to research, task analysis, architecture, mode
  selection, Agent driving infrastructure, task tree/slices, implementation,
  verification, acceptance, closure, or round `n+1`.
- The full-project lifecycle reaches any downstream gate after
  `requirement_baseline_confirmation_gate`.
- A developer asks to make the later Vibe Coding project flow as detailed as
  the requirement discovery flow.
- A repository infrastructure task asks to refine downstream project lifecycle
  gates, execution loops, evidence packets, or closure behavior.

## Purpose

Prevent the post-requirement lifecycle from collapsing into short summaries.
The skill requires every downstream stage to produce a durable record,
decision-ready packets, evidence links, confirmation status, and a next
condition before the following stage can consume it.

## Required Runtime Inputs

- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/main-tree.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/decisions/records.md`
- `docs/ai/evidence/index.md`
- `docs/ai/handoff/current.md`
- `docs/ai/templates/project-lifecycle-downstream-gates-sop.md`
- `docs/ai/templates/project-lifecycle-downstream-record.yaml`
- `docs/ai/lifecycle/README.md`

## Gate Order

1. Run `research_plan_gate`.
   Convert the confirmed requirement baseline into research questions, source
   priorities, unstable facts, official-source targets, and research exit
   criteria.
2. Run `current_source_research_execution_gate`.
   Gather current-source evidence where needed, preserve sources, summarize
   findings, identify stale or uncertain facts, and reject unsuitable options.
3. Run `research_synthesis_confirmation_gate`.
   Produce a research recommendation packet and stop for developer
   confirmation when stack, API, library, model, hosting, pricing, regulation,
   or other unstable choices affect the project.
4. Run `task_graph_analysis_gate`.
   Build a task graph with deliverables, dependencies, critical path, slice
   candidates, test plan, documentation plan, evidence plan, and closure
   criteria.
5. Run `task_analysis_confirmation_packet_gate`.
   Produce a confirmation packet that lets the developer accept, edit, defer,
   or reject the task plan.
6. Run `architecture_dossier_gate`.
   Map confirmed requirements and research to architecture decisions, data
   model, API boundaries, UI model, security/privacy rules, rollback strategy,
   quality gates, and unresolved tradeoffs.
7. Run `architecture_confirmation_packet_gate`.
   Stop for architecture and stack confirmation before implementation mode or
   slice creation.
8. Run `infra_mode_risk_matrix_gate`.
   Score risk drivers and recommend `standard_light_risk_escalated` or
   `enterprise_high_assurance` with artifact depth, escalation triggers, and
   cost of assurance.
9. Run `infra_mode_confirmation_packet_gate`.
   Stop for developer confirmation of the mode.
10. Run `agent_driving_infra_plan_gate`.
    Define the project-specific skills, knowledge entries, schemas, prompt
    surfaces, evidence requirements, quality gates, and handoff anchors needed
    for the confirmed mode.
11. Run `task_tree_slice_contract_gate`.
    Create or update task trees and slices only when each slice has allowed
    files, forbidden files, verification commands, acceptance criteria,
    rollback plan, evidence plan, and next condition.
12. Run `implementation_iteration_ledger_gate`.
    During each slice, record implementation iterations, failures, fixes,
    tests, documentation updates, evidence, state sync, and retry limits.
13. Run `verification_evidence_packet_gate`.
    Before claiming completion, produce a validation packet that maps commands,
    tests, manual checks, and evidence to acceptance criteria.
14. Run `acceptance_closure_packet_gate`.
    Produce a closure packet with requirement coverage, residual risks, docs,
    handoff, diff summary, developer acceptance status, and rollback notes.
15. Run `next_round_reentry_gate`.
    Decide whether the next allowed state is acceptance, a new slice, closure,
    or `next_round_intake[n+1]`.

## Hard Stops

Do not proceed from one downstream stage to the next using only a narrative
summary. Each stage needs record state, evidence, and the required confirmation
status.

Do not install dependencies, create package manifests, write application code,
or create a business task tree before the upstream requirement, research,
architecture, and mode confirmation gates are satisfied.

Do not expand an active slice beyond its contract. Create a new slice or return
to task analysis when scope changes.

## Verification

Downstream detailing is complete only when:

- A downstream record exists or the active slice explicitly records why it is
  not yet created.
- Research, task analysis, architecture, mode, Agent infra, slices,
  implementation, verification, closure, and next-round sections have statuses.
- Evidence links prove validation where the gate requires it.
- Confirmation packets are recorded for hard-confirmation gates.
- Handoff points to the next allowed action.

Chat-only intent is not enough.
