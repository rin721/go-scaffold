# Compiler Runtime Assimilation SOP

This SOP is the canonical procedure for absorbing a Vibe Coding compiler,
generator, or major infrastructure-governance specification into an already
initialized Vibe Coding repository.

## 0. Intake And Route

Classify the source:

| Source shape | Route |
| --- | --- |
| Compiler, generator, prompt-generation, runtime-governance, schema, skill, knowledge, task-tree, or status-machine specification | `branch_vibe_coding_infra` |
| Product, application, business feature, bugfix, refactor, or user-facing implementation | `branch_mainline_idea` after the full-project lifecycle gates |
| Mixed compiler and product idea | Split them; assimilate compiler semantics first, then ask for product requirement confirmation |

Required output:

- Source summary.
- Selected branch.
- Whether physical assignment is explicitly requested.
- Write boundary.

## 1. Four-Layer Distillation

Extract the source into four layers:

1. Core invariants, human authority, safety, and agent agency boundaries.
2. State machine, initialization task tree, round lifecycle, task-tree closure,
   escape hatches, and retry limits.
3. Physical artifacts: directories, schema fields, evidence, knowledge, skills,
   roles, templates, handoff, and validation.
4. Mode matrices, task-type acceptance matrices, forbidden matrices, version
   strategy, and example output requirements.

Do not copy the source wholesale. Capture only the runtime semantics that the
current repository needs.

## 2. Repository Fit

Compare the source semantics with current artifacts:

- Already covered: leave as-is unless too thin.
- Covered but too implicit: add skill, SOP, schema note, validator term, or
  knowledge entry.
- Missing but relevant: add a scoped infrastructure task and slice.
- Irrelevant to current repo: record as deferred or out of scope.
- Unsafe or ambiguous: stop for developer confirmation.

## 3. Runtime Translation

Translate the confirmed scope into physical files:

- `docs/ai/requirements/ledger.yaml`
- `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/tasks/slices/*`
- `docs/ai/runtime-rule-index.md`
- `docs/ai/skills/*`
- `docs/ai/templates/*`
- `docs/ai/knowledge/*`
- `docs/ai/architecture/*`
- `docs/ai/evidence/index.md`
- `docs/ai/status/current.yaml`
- `docs/ai/handoff/current.md`

Keep `AGENTS.md` short. Do not move the source compiler text into `AGENTS.md`.

## 4. Validation

Run:

```powershell
python docs/ai/scripts/validate_runtime.py
rg -n "compiler_runtime_assimilation|INV-18|no_prompt_runtime_dependency_recheck_gate" docs/ai
Test-Path package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools
```

Record the result in `docs/ai/evidence/index.md`.

## 5. Closure

Before claiming completion:

- Status points to the new slice.
- Task tree includes the new task.
- Knowledge and skills indexes include the new entries.
- Handoff explains the next condition.
- No runtime file asks future agents to read the original compiler prompt as a
  normal recovery path.

The next allowed action is developer acceptance, requested edits, or explicit
closure/migration of the infrastructure branch.
