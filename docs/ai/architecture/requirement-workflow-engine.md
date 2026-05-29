# Requirement Workflow Engine Architecture

- ID: `architecture_006_requirement_workflow_engine`
- Date: `2026-05-29T22:24:00+08:00`
- Branch: `branch_vibe_coding_infra`
- Requirement: `req_infra_010_requirement_workflow_engine`
- Task: `task_010_requirement_workflow_engine`

## Purpose

The requirement workflow engine turns project idea intake into a recoverable
state machine. It gives future agents a deterministic path for product-manager
style elicitation: capture intent, create a durable intake record, ask 5-7
numbered questions, persist answers, generate finer follow-ups, and keep
business code generation locked until the requirement baseline and downstream
gates are confirmed.

## Components

| Component | Path | Responsibility |
| --- | --- | --- |
| Engine definition | `docs/ai/requirements/workflow_engine.yaml` | Declares the node graph, loop controls, outputs, and permission invariants. |
| Probe catalog | `docs/ai/requirements/template_discovery.yaml` | Provides reusable question templates for commerce, SaaS, blog, and generic applications. |
| State machine | `docs/ai/requirements/state_machine.yaml` | Defines `REQUIREMENTS_GATHERING` and locks code generation, package manifests, dependency installation, and app skeletons. |
| Project skill | `docs/ai/skills/requirement-workflow-engine/SKILL.md` | Tells future agents when and how to run the engine. |
| SOP | `docs/ai/templates/requirement-workflow-engine-sop.md` | Provides the operating procedure for engine execution and recovery. |
| Intake records | `docs/ai/requirements/intake/` | Stores per-round `round_<nnn>_<feature_slug>.yaml` discovery records. |
| Ledger | `docs/ai/requirements/ledger.yaml` | Receives confirmed, rejected, deferred, out-of-scope, or research-needed requirements. |
| Validator | `docs/ai/scripts/validate_runtime.py` | Verifies that the engine artifacts and key invariants exist. |

## Engine Topology

```text
raw_project_idea
-> route_task_line
-> capture_intent
-> select_probe_template
-> create_requirement_record
-> enter_requirements_gathering_lock
-> ask_question_batch
-> absorb_answers
-> generate_followups
-> coverage_check
   -> ask_question_batch [while critical domains remain open]
-> requirement_plan
-> baseline_readiness
-> requirement_baseline_confirmation
-> downstream_lifecycle_entry
```

The loop is finite. Each iteration must add a durable answer, a new question, a
deferral, an out-of-scope decision, or a named blocker. Three repeated blockers
trigger human decision instead of silent repetition.

## Permission Model

`REQUIREMENTS_GATHERING` is a hard lock:

- Business code generation: forbidden.
- Package manifests and lockfiles: forbidden.
- Dependency installation: forbidden.
- Application skeletons: forbidden.
- Requirement intake records, ledger sync, status, evidence, decisions, tasks,
  and handoff updates: allowed.

Implementation remains impossible until requirement baseline, research, task
analysis, architecture, and infrastructure mode confirmations authorize an
executable slice.
