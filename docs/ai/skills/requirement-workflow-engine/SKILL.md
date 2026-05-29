# Requirement Workflow Engine Skill

## Trigger Conditions

Use this skill when any of these are true:

- A developer confirms strategy C for requirement discovery or asks for a
  complete workflow engine for project idea intake.
- A future mainline idea needs product-manager style requirement elicitation
  before baseline confirmation.
- The active requirement discovery record or state machine is in
  `REQUIREMENTS_GATHERING`.
- A developer asks to ensure code generation is locked while requirements are
  still being gathered.

## Purpose

Run the declarative requirement workflow engine. The engine converts a raw
project idea into durable requirement intake state, template-driven question
batches, answer persistence, follow-up question generation, coverage checks,
and baseline readiness packets.

## Required Runtime Inputs

- `docs/ai/requirements/workflow_engine.yaml`
- `docs/ai/requirements/template_discovery.yaml`
- `docs/ai/requirements/state_machine.yaml`
- `docs/ai/templates/requirement-workflow-engine-sop.md`
- `docs/ai/templates/project-requirement-discovery-record.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/main-tree.yaml`
- `docs/ai/evidence/index.md`
- `docs/ai/handoff/current.md`

## Steps

1. Route the work to the correct task line.
2. Capture the raw idea and normalize it without confirming assumptions.
3. Select a probe template from `template_discovery.yaml`; use `generic_app`
   when no specific template matches.
4. Create or update
   `docs/ai/requirements/intake/round_<nnn>_<feature_slug>.yaml`.
5. Enter the `REQUIREMENTS_GATHERING` lock from `state_machine.yaml`.
6. Ask 5-7 numbered, domain-bound questions.
7. Record developer answers by overwriting the intake YAML.
8. Generate follow-up questions from the selected template and newly confirmed
   answers.
9. Repeat until critical domains are confirmed, deferred, out of scope, or
   blocked by a named decision.
10. Produce a baseline readiness packet and request hard confirmation only when
    the readiness gate passes.

## Hard Stops

- Do not generate business code while `REQUIREMENTS_GATHERING` is active.
- Do not create package manifests, lockfiles, dependency installs, source
  directories, or application skeletons from a raw idea.
- Do not ask for broad baseline confirmation while critical domains remain
  unknown.
- Do not mark AI-inferred requirements as confirmed without a developer answer.
- Do not keep asking the same blocked question indefinitely; after three
  repeated blockers, ask for a human decision.

## Verification

The engine is active only when the physical artifacts exist and validation
passes:

- Engine definition exists.
- Template catalog exists with commerce, SaaS, blog, and generic app probes.
- State machine exists with `REQUIREMENTS_GATHERING` and
  `code_generation_allowed: false`.
- Runtime validator covers the engine artifacts.
- Evidence and handoff point to the next allowed action.
