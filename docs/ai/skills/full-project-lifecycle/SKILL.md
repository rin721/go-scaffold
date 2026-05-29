# Full Project Lifecycle Skill

## Trigger Conditions

Use this skill when any of these are true:

- A developer provides a new product, business, application, feature, bugfix,
  upgrade, expansion, or refactor idea for the mainline task tree.
- A developer asks why the Vibe Coding flow did not fully drive an idea into a
  complete project.
- A developer asks to continue with `next step` and the physical state points
  to `branch_mainline_idea`, a future product branch, or a lifecycle gate.
- A repository infrastructure task asks to establish or repair the mainline
  full-project workflow.

## Purpose

Prevent the mainline flow from collapsing into a brief requirement summary.
Every mainline or future project task line must pass through detailed,
auditable gates before implementation and must keep looping through slices and
round `n+1` until the task tree is closed or the developer redirects it.

## Required Runtime Inputs

- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/main-tree.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/decisions/records.md`
- `docs/ai/evidence/index.md`
- `docs/ai/handoff/current.md`
- `docs/ai/templates/full-project-lifecycle-workflow-sop.md`
- `docs/ai/skills/requirement-workflow-engine/SKILL.md`
- `docs/ai/templates/requirement-workflow-engine-sop.md`
- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`
- `docs/ai/skills/project-requirement-discovery/SKILL.md`
- `docs/ai/templates/project-requirement-discovery-sop.md`
- `docs/ai/templates/project-requirement-discovery-record.yaml`
- `docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md`
- `docs/ai/templates/project-lifecycle-downstream-gates-sop.md`
- `docs/ai/templates/project-lifecycle-downstream-record.yaml`

## Gate Order

1. Route the work to the correct task line before analysis:
   `branch_vibe_coding_infra`, `branch_mainline_idea`, or another confirmed
   branch.
2. Capture the developer's target idea for the current round.
3. Run the requirement workflow engine and project requirement discovery before
   treating the idea as a requirement baseline candidate: enter
   `REQUIREMENTS_GATHERING`, select a probe template, interpret the idea, map
   requirement domains, inventory developer-stated and AI-inferred
   requirements, ask 5-7 numbered questions, build a question backlog, collect
   answers across rounds, persist the discovery record, generate follow-ups,
   plan remaining requirement work, and check baseline readiness.
4. Analyze requirements into problem, users, scope, non-goals, constraints,
   risks, acceptance criteria, and open questions.
5. Stop for hard confirmation of the requirement baseline.
6. Initialize the downstream lifecycle record from
   `docs/ai/templates/project-lifecycle-downstream-record.yaml`.
7. Run `research_plan_gate` and `current_source_research_execution_gate` for
   technology, API, library, model, hosting, pricing, regulation, or other
   unstable choices.
8. Run `research_synthesis_confirmation_gate` and stop for research and stack
   confirmation where required.
9. Run `task_graph_analysis_gate` to create deliverables, dependencies,
   critical path, slice candidates, test strategy, documentation strategy,
   evidence strategy, and closure criteria.
10. Run `task_analysis_confirmation_packet_gate`.
11. Run `architecture_dossier_gate` to map requirements and research to data
    model, API boundaries, UI model, security/privacy rules, rollback, quality
    gates, alternatives, and unresolved decisions.
12. Run `architecture_confirmation_packet_gate`.
13. Run `infra_mode_risk_matrix_gate` and recommend
    `standard_light_risk_escalated` or `enterprise_high_assurance`.
14. Run `infra_mode_confirmation_packet_gate`.
15. Run `agent_driving_infra_plan_gate` for project-specific skills,
    knowledge, schemas, templates, validation, evidence, and handoff anchors.
16. Run `task_tree_slice_contract_gate` before any slice becomes executable.
17. Run `implementation_iteration_ledger_gate` for every implementation slice.
18. Run `verification_evidence_packet_gate` before claiming slice or round
    completion.
19. Run `acceptance_closure_packet_gate` when all required work is complete.
20. Run `next_round_reentry_gate` to decide whether to wait for acceptance,
    continue another slice, close the current tree, or enter
    `next_round_intake[n+1]`.

## Hard Stops

Do not implement product code, install dependencies, create package manifests,
choose a technology stack, change production settings, or expand agent agency
before the relevant confirmation gate is satisfied and recorded.

Do not treat `next step` as permission to skip requirement, research, task
analysis, architecture, or mode confirmation. `next step` only advances to the
next permitted gate or active slice in physical state.

## Verification

A lifecycle step is complete only when the corresponding physical artifact and
evidence exist. Chat-only intent is not enough.
