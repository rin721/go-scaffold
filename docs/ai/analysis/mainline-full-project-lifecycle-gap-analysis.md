# Mainline Full-Project Lifecycle Gap Analysis

- ID: `analysis_002_mainline_full_project_lifecycle_gap`
- Date: `2026-05-29T21:04:30+08:00`
- Branch: `branch_vibe_coding_infra`
- Related requirement: `req_infra_006_full_project_lifecycle`
- Related task: `task_006_full_project_lifecycle`

## Diagnosis

The current runtime already names a confirmation chain in
`docs/ai/runtime-rule-index.md`, but the physical artifacts only encode the
chain as phase labels. That is enough to remind an agent that requirements,
research, task analysis, architecture, mode selection, execution, validation,
documentation, state sync, and next-step handoff exist. It is not enough to
force a developer idea through a complete project lifecycle.

The mainline idea tree currently says that future product or business ideas
must complete requirement confirmation, research, task analysis, architecture,
and mode confirmation before task creation. It does not define the detailed
gates, required outputs, confirmation questions, evidence shape, or the
`next step` behavior that carries the work through every slice and then into
round `n+1`.

## Why The Previous Flow Was Too Thin

| Area | Existing behavior | Gap |
| --- | --- | --- |
| Task-line routing | The task forest separates mainline idea work from Vibe Coding infrastructure work. | There is no mandatory routing gate before every new idea to decide whether the work belongs to `branch_mainline_idea`, `branch_vibe_coding_infra`, or another future branch. |
| Requirement analysis | The runtime has requirement baseline fields and a requirements ledger. | A mainline idea can be summarized too quickly without a required problem statement, user outcomes, scope, non-goals, acceptance criteria, constraints, risks, and hard confirmation. |
| Research and stack validation | The state machine mentions research validation. | It does not require internet-backed, current-source validation when selecting technology stacks, libraries, APIs, or architecture-sensitive dependencies. |
| Task analysis | Task trees and slices exist. | There is no required work breakdown that proves the whole project can be completed, including dependencies, test strategy, documentation work, and closure conditions. |
| Architecture and stack design | The state machine names architecture and stack design. | There is no architecture dossier template tying requirements, research, stack choices, risks, data model, integration points, and acceptance gates together. |
| Infrastructure mode | The state machine mentions mode recommendation and confirmation. | The choice surface is not explicit: standard lightweight mode with risk-based escalation versus enterprise high-assurance mode with full audit priority. |
| Agent driving infrastructure | Skills, knowledge, status, evidence, and handoff exist. | There is no rule that each confirmed project round must establish or update the project-specific agent facilities needed to execute that round. |
| Slice loop | Slice execution skill exists. | The loop does not spell out the mandatory order: implement code, write/run tests, document, record evidence, update state, then expose the `next step` protocol. |
| Round closure | The state machine includes closure and `next_round_intake[n+1]`. | The physical docs do not define how a completed task tree re-enters the same idea-intake lifecycle for iteration, expansion, upgrade, or refactor work. |

## Root Cause

The runtime captured the lifecycle as a conceptual state machine, but not as a
physical operating procedure with:

- Required inputs and outputs per gate.
- Hard confirmation points that cannot be collapsed into a summary.
- A task-line routing gate before mainline work.
- A current-source research requirement for technology stack decisions.
- A mode-selection rubric with developer confirmation.
- A durable `next step` protocol for slice continuation and round `n+1`.

## Selected Repair Strategy

Add a dedicated full-project lifecycle workflow as Vibe Coding infrastructure:

- A project skill: `docs/ai/skills/full-project-lifecycle/SKILL.md`.
- A canonical SOP template:
  `docs/ai/templates/full-project-lifecycle-workflow-sop.md`.
- An architecture note:
  `docs/ai/architecture/full-project-lifecycle-workflow.md`.
- A knowledge entry:
  `docs/ai/knowledge/entries/kb_007_full_project_lifecycle.md`.
- Runtime index, task tree, slice, evidence, status, and handoff updates.

This repair stays inside repository infrastructure. It does not create product
code, install dependencies, choose a business technology stack, or start a
mainline implementation before the future developer idea has passed the full
confirmation chain.
