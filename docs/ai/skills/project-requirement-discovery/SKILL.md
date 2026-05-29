# Project Requirement Discovery Skill

## Trigger Conditions

Use this skill when any of these are true:

- A developer provides a compact project idea, product idea, application idea,
  feature idea, upgrade, expansion, or refactor request.
- The full-project lifecycle reaches `idea_intake_gate`,
  `requirement_analysis_gate`, or `requirement_baseline_confirmation_gate`.
- A developer asks why the Vibe Coding flow summarized an idea too quickly or
  did not guide them through complete project requirements.
- A repository infrastructure task asks to improve requirement collection,
  requirement planning, requirement persistence, or developer requirement
  elicitation.

## Purpose

Prevent an idea from being prematurely converted into a generic MVP summary.
The skill turns a raw idea into a durable requirement-discovery record, a
domain scan, a prioritized question backlog, a requirement collection plan, and
a readiness check for hard baseline confirmation.

## Required Runtime Inputs

- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/main-tree.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/evidence/index.md`
- `docs/ai/handoff/current.md`
- `docs/ai/templates/project-requirement-discovery-sop.md`
- `docs/ai/templates/project-requirement-discovery-record.yaml`
- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`

## Gate Order

0. When `docs/ai/requirements/workflow_engine.yaml` exists, run the
   requirement workflow engine first. Its template catalog, state machine, and
   `REQUIREMENTS_GATHERING` permission lock control question batches,
   follow-up generation, and write boundaries.
1. Run `idea_seed_intake_gate`.
   Capture the developer's raw idea, target outcome, known constraints, and
   unknowns. Treat the idea as a seed, not as a confirmed requirement set.
2. Run `idea_interpretation_gate`.
   Restate the idea in three layers: product intent, likely system category,
   and assumptions that require confirmation.
3. Run `domain_surface_mapping_gate`.
   Identify the requirement domains that must be scanned for the project type:
   actors, workflows, data, integrations, permissions, operations,
   non-functional requirements, tests, docs, and non-goals.
4. Run `requirement_inventory_gate`.
   Separate requirements into developer-stated, AI-inferred candidate,
   research-needed, decision-needed, deferred, out-of-scope, and rejected.
5. Run `question_backlog_gate`.
   Build a prioritized backlog of high-leverage questions. Ask the first small
   batch rather than overwhelming the developer.
6. Run `requirement_collection_round_gate`.
   Persist developer answers, update statuses, identify contradictions, and
   prepare the next batch of questions when needed.
7. Run `requirement_plan_gate`.
   Produce a requirement plan: which domains are ready, which need decisions,
   which require current-source research, and which are deferred.
8. Run `requirement_persistence_sync_gate`.
   Store durable discovery state in a per-round intake record and synchronize
   confirmed items to `docs/ai/requirements/ledger.yaml`.
9. Run `baseline_readiness_gate`.
   Ask for requirement-baseline confirmation only when each critical domain is
   confirmed, explicitly deferred, out of scope, or blocked by a named decision.

## Developer Guidance Rules

- Ask specific questions tied to missing requirement domains.
- Prefer project-type probe templates from
  `docs/ai/requirements/template_discovery.yaml` before inventing a new
  question set.
- Provide concrete options only as decision aids; do not silently choose for the
  developer.
- Preserve AI-inferred details as candidates until confirmed.
- Record non-goals as carefully as goals.
- Keep question batches small enough for the developer to answer.
- Continue collecting requirements across turns until the baseline is ready or
  the developer redirects the round.

## Hard Stops

Do not implement product code, install dependencies, create package manifests,
select a final technology stack, or create a mainline task tree from a raw idea
alone.

When the state machine is in `REQUIREMENTS_GATHERING`, code generation,
package manifests, dependency installation, and application skeleton creation
are explicitly locked.

Do not ask for a single broad confirmation when critical requirement domains
are still unknown.

Do not treat a suggested MVP as a confirmed requirement baseline unless the
developer explicitly confirms the baseline and the confirmation is recorded.

## Verification

Requirement discovery is complete only when the physical state proves it:

- A discovery record exists or the active slice explicitly records why it is
  not yet created.
- Requirement candidates have statuses.
- Confirmed requirements are synchronized to the requirement ledger.
- Open questions and deferred items are preserved.
- Evidence and handoff point to the relevant artifacts.

Chat-only intent is not enough.
