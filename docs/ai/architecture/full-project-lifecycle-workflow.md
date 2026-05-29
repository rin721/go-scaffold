# Full-Project Lifecycle Workflow Architecture

- ID: `architecture_002_full_project_lifecycle_workflow`
- Date: `2026-05-29T21:04:30+08:00`
- Branch: `branch_vibe_coding_infra`
- Requirement: `req_infra_006_full_project_lifecycle`
- Task: `task_006_full_project_lifecycle`

## Purpose

This architecture turns the Vibe Coding mainline from a short requirement
prompt into a repeatable project lifecycle. A developer idea must be routed,
analyzed, researched, confirmed, decomposed, implemented, tested, documented,
closed, and then re-entered as round `n+1` when the developer wants to iterate,
upgrade, expand, or refactor.

## Components

| Component | Path | Responsibility |
| --- | --- | --- |
| Runtime rule index | `docs/ai/runtime-rule-index.md` | Declares the mandatory full-project lifecycle gates. |
| Project skill | `docs/ai/skills/full-project-lifecycle/SKILL.md` | Tells agents when and how to run the lifecycle. |
| Requirement discovery skill | `docs/ai/skills/project-requirement-discovery/SKILL.md` | Expands raw ideas into domain scans, question backlogs, persistent intake records, and baseline readiness checks before requirement confirmation. |
| SOP template | `docs/ai/templates/full-project-lifecycle-workflow-sop.md` | Defines gate-by-gate required outputs, confirmations, and stop conditions. |
| Requirement discovery SOP | `docs/ai/templates/project-requirement-discovery-sop.md` | Defines the detailed idea-to-requirement sub-procedure used inside the lifecycle. |
| Downstream detailing skill | `docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md` | Expands post-baseline stages into research, task analysis, architecture, mode, Agent infrastructure, slices, implementation, verification, closure, and n+1 record gates. |
| Downstream detailing SOP | `docs/ai/templates/project-lifecycle-downstream-gates-sop.md` | Defines the detailed post-baseline lifecycle sub-procedure used inside the lifecycle. |
| Task forest | `docs/ai/tasks/forest.yaml` | Routes work to the right task line before mainline execution. |
| Mainline tree | `docs/ai/tasks/main-tree.yaml` | Holds future product or business task trees after confirmation. |
| Infrastructure tree | `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml` | Holds workflow and governance changes like this one. |
| Requirements, decisions, evidence, knowledge, handoff | `docs/ai/*` | Preserve traceability and recovery without relying on chat memory. |

## Lifecycle Gates

1. `task_line_routing_gate`
2. `idea_intake_gate`
3. `requirement_analysis_gate`
4. `requirement_baseline_confirmation_gate`
5. `current_source_research_gate`
6. `research_confirmation_gate`
7. `task_analysis_gate`
8. `task_analysis_confirmation_gate`
9. `architecture_and_stack_design_gate`
10. `architecture_confirmation_gate`
11. `infra_mode_recommendation_gate`
12. `infra_mode_confirmation_gate`
13. `agent_driving_infra_gate`
14. `task_tree_and_slice_gate`
15. `implementation_test_docs_state_loop`
16. `next_step_protocol_gate`
17. `round_closure_gate`
18. `next_round_intake[n+1]`

## Mode Choice Surface

Future project rounds must recommend one of two infrastructure modes and wait
for developer confirmation:

| Mode | Use when | Default controls |
| --- | --- | --- |
| `standard_light_risk_escalated` | Local prototypes, documentation, small tools, low/medium risk features, or early exploration. | Lightweight artifacts, focused evidence, required tests for touched behavior, escalation when privacy, money, production, destructive operations, or high uncertainty appears. |
| `enterprise_high_assurance` | Production, money, auth, private data, compliance, irreversible migrations, multi-agent autonomy, or unclear high-risk architecture. | Full audit trail, stricter approvals, broader tests, architecture/security review, rollback plans, richer evidence and acceptance gates. |

The recommendation must explain why the selected mode fits the project risk.
It must also explain why the other mode is not the best default for the current
round, what would cause escalation, and the exact developer confirmation needed
before proceeding. The developer has final authority.

## Next-Step Protocol

`next step` is not a vague continuation command. It means:

1. Read status, task forest, current tree, current slice, evidence, and handoff.
2. Determine the next lifecycle gate or slice from physical state.
3. Refuse to skip a hard confirmation gate.
4. If a slice is active, implement only within its allowed scope.
5. Run declared verification.
6. Record evidence and update status/handoff.
7. If the current task tree is complete, enter `round_closure_gate`.
8. After closure, begin `next_round_intake[n+1]` for a new idea, iteration,
   upgrade, expansion, or refactor.
