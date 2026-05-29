# Full Project Lifecycle Workflow SOP

This SOP is the canonical operating procedure for turning a developer idea into
a complete Vibe Coding project round and then continuing into round `n+1`.

## 0. Task-Line Routing Gate

Before analyzing the idea, classify the work:

| Route | Use for |
| --- | --- |
| `branch_vibe_coding_infra` | Changes to `AGENTS.md`, `docs/ai/*`, runtime rules, schemas, scripts, skills, knowledge, evidence, handoff, quality gates, or workflow governance. |
| `branch_mainline_idea` | Product, business, application, feature, bugfix, upgrade, expansion, or refactor work after confirmation. |
| Future branch | A new durable workstream only after the developer confirms why mainline and infrastructure are not sufficient. |

Required output:

- Selected task line.
- Reason for the route.
- Whether the current physical state allows starting that line.

Hard stop:

- Do not create mainline tasks while an active infrastructure branch must be
  accepted, edited, migrated, or explicitly closed first.

## 1. Idea Intake Gate

Capture the developer's idea in plain language.

Required output:

- Goal idea.
- Intended users or operators.
- Desired outcome.
- Known constraints.
- Unknowns that must be resolved before implementation.

Before moving to requirement baseline work, run the requirement workflow engine
and project requirement discovery SOP:

- `docs/ai/skills/requirement-workflow-engine/SKILL.md`
- `docs/ai/templates/requirement-workflow-engine-sop.md`
- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`
- `docs/ai/skills/project-requirement-discovery/SKILL.md`
- `docs/ai/templates/project-requirement-discovery-sop.md`
- `docs/ai/templates/project-requirement-discovery-record.yaml`

The discovery sub-procedure must enter `REQUIREMENTS_GATHERING`, lock business
code generation, select a probe template, produce idea interpretation, domain
surface mapping, requirement inventory, numbered question batches, question
backlog, requirement collection rounds, follow-up triggers, requirement
planning, persistence sync, and baseline readiness evidence.

## 2. Requirement Analysis Gate

Turn the idea into an implementation-ready requirement candidate.

Required output:

- Problem statement.
- User workflows.
- Functional requirements.
- Non-functional requirements.
- Non-goals.
- Data and privacy assumptions.
- Acceptance criteria.
- Open questions.
- Risks and dependencies.

Hard stop:

- Ask for requirement-baseline confirmation before research or architecture is
  treated as committed.
- Do not ask for a single broad confirmation if the requirement discovery
  record still contains unresolved critical domains or unmarked AI-inferred
  candidates.

## 3. Research And Technology Validation Gate

Research is mandatory when current technology stack, APIs, library versions,
pricing, regulations, models, hosting behavior, or other unstable facts affect
the project.

Before running downstream research and subsequent post-baseline stages, load:

- `docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md`
- `docs/ai/templates/project-lifecycle-downstream-gates-sop.md`
- `docs/ai/templates/project-lifecycle-downstream-record.yaml`

The downstream sub-procedure must produce research plan, source matrix,
research confirmation packet, task graph, task-analysis confirmation packet,
architecture dossier, architecture confirmation packet, mode risk matrix, mode
confirmation packet, Agent driving infrastructure plan, slice contracts,
implementation iteration ledger, verification evidence packet, acceptance
closure packet, and next-round reentry state.

Required output:

- Research questions.
- Sources used, preferring official or primary sources.
- Findings.
- Technology stack candidates.
- Risks, tradeoffs, and rejected options.
- Recommendation.

Hard stop:

- Ask for research and technology-stack confirmation before task analysis is
  treated as committed.

## 4. Task Analysis Gate

Break the confirmed requirements and research into a complete project plan.

Required output:

- Task tree candidates.
- Slice boundaries.
- Dependencies and ordering.
- Test strategy.
- Documentation plan.
- Evidence plan.
- Completion and closure criteria.

Hard stop:

- Ask for task-analysis confirmation before architecture and slice creation are
  treated as committed.

## 5. Architecture And Stack Design Gate

Design the system before implementation.

Required output:

- Architecture overview.
- Technology stack.
- Data model and storage choices.
- API or integration boundaries.
- UI or interaction model when applicable.
- Security, privacy, and permission boundaries.
- Error handling and rollback strategy.
- Quality gates and test matrix.
- Mapping from requirements to architecture decisions.

Hard stop:

- Ask for architecture and stack confirmation before mode selection and
  implementation.

## 6. Infrastructure Mode Recommendation Gate

Recommend one mode and explain why.

| Mode | Meaning |
| --- | --- |
| `standard_light_risk_escalated` | Standard lightweight mode: keep the workflow lean by default, and add stricter controls when risk increases. |
| `enterprise_high_assurance` | Enterprise high-assurance mode: prefer full audit and higher assurance first, even when heavier. |

Use `standard_light_risk_escalated` when the project is local, reversible,
low/medium risk, exploratory, or mostly documentation/prototype work.

Use `enterprise_high_assurance` when the project touches production, money,
private data, authentication, compliance, destructive operations, external
deployment, irreversible migrations, high user impact, or broad agent autonomy.

Required output:

- Recommended mode.
- Why that mode fits.
- Why the other mode is not the best default right now.
- What would cause escalation.
- What artifacts the mode requires.
- The exact developer confirmation needed before proceeding.

Hard stop:

- Ask for developer confirmation of the mode.

## 7. Agent Driving Infrastructure Gate

After mode confirmation, establish the project-specific Agent Vibe Coding
facilities for the round.

Possible outputs:

- Skills.
- Knowledge entries.
- Context/rule updates.
- Prompt surfaces or templates.
- Schemas.
- Task tree and execution slices.
- Evidence and handoff anchors.

These outputs must be scoped to the confirmed mode and architecture.

## 8. Task Tree And Slice Gate

Create the first-round task tree and execution slices only after the prior
gates are satisfied.

Each slice must define:

- Goal.
- Inputs and outputs.
- Allowed and forbidden files.
- Allowed tools.
- Approval-required actions.
- Verification commands.
- Acceptance criteria.
- Rollback plan.
- Next condition.

## 9. Implementation, Test, Docs, And State Loop

For each active slice:

1. Implement code or documentation inside the allowed scope.
2. Write or update tests appropriate to risk.
3. Run verification commands.
4. Fix failures within the slice boundary.
5. Update docs and knowledge when stable facts changed.
6. Record evidence.
7. Update status, current slice, task tree, and handoff.
8. Tell the developer the next allowed action.

Hard stop:

- Do not claim completion without validation evidence.

## 10. Next-Step Protocol

When the developer says `next step`:

1. Read physical state first.
2. Determine the next allowed lifecycle gate or slice.
3. If a hard confirmation is pending, ask for that confirmation.
4. If a slice is active and confirmed, execute that slice.
5. If the task tree is complete, enter closure.
6. If the round is closed, begin `next_round_intake[n+1]`.

## 11. Round Closure And n+1 Loop

A task tree is complete only after:

- Requirements are satisfied.
- Tests and validation pass.
- Documentation is updated.
- Evidence is recorded.
- Status and handoff are synchronized.
- Developer acceptance is recorded or the remaining acceptance decision is
  explicitly pending.

After closure, future iteration, upgrade, expansion, or refactor work starts a
new round and repeats this SOP from the task-line routing gate.
