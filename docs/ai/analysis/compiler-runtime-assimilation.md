# Compiler Runtime Assimilation Analysis

- ID: `analysis_003_compiler_runtime_assimilation`
- Date: `2026-05-29T21:30:00+08:00`
- Branch: `branch_vibe_coding_infra`
- Related requirement: `req_infra_007_compiler_runtime_assimilation`
- Related task: `task_007_compiler_runtime_assimilation`

## Diagnosis

The developer provided a complete `vibe-coding-infra-compiler-v1.8.0`
generation specification and asked to deeply analyze, distill, and assign that
context into this repository. The important repository fact is that this project
is already a Vibe Coding infrastructure repository with physical runtime
artifacts under `docs/ai/*`. Therefore the correct action is not to generate a
new `prompt.md`, not to make the original compiler text a runtime dependency,
and not to restart bootstrap from zero. The correct action is to assimilate the
compiler semantics into recoverable runtime artifacts: task-line rules, skills,
knowledge, templates, schema notes, validation, evidence, status, and handoff.

## Four-Dimensional Scan

| Dimension | Finding | Runtime adjustment |
| --- | --- | --- |
| Context and meta-philosophy | The source specification treats Vibe Coding as a controlled engineering system: human authority, auditable agent action, physical state, evidence, and recoverable task trees are the reliability source. | Preserve the compiler-versus-runtime boundary. Record the source as a requirement and knowledge input, then express the usable behavior through local runtime artifacts. |
| Structure and execution flow | The source has a four-layer architecture: invariants and human boundaries; state machine and task-tree flow; physical schemas, evidence, knowledge, and skills; mode matrices, acceptance matrices, and examples. | Add a dedicated compiler-runtime assimilation workflow that maps these layers into existing `docs/ai` entries without duplicating the raw prompt. |
| Engineering density and executability | The source is dense enough to drive artifacts, but only if it is translated into concrete repository files, IDs, validation commands, and next-step conditions. | Create a project skill, canonical SOP, architecture note, knowledge entry, execution slice, validator coverage, and synchronized indexes. |
| Contradictions and smells | The main risk is turning the compiler text into a hidden mega-rule that future agents must reread. A second risk is treating every compiler upgrade as permission to rewrite the whole repository. | Encode `INV-18-COMPILER-CONTEXT-ASSIMILATION`: compiler text is a semantic source to distill, not a runtime authority. Assimilation must be scoped, validated, evidenced, and reversible. |

## Extracted Runtime Semantics

The compiler context contributes five concrete runtime requirements for this
repository:

1. A future compiler or generator specification must be routed to
   `branch_vibe_coding_infra`, not to `branch_mainline_idea`.
2. The agent must distinguish three layers: source compiler text, current
   repository runtime artifacts, and future daily runtime behavior.
3. The source must be decomposed into reusable runtime units: invariants,
   state gates, schemas, task trees, execution slices, evidence, knowledge,
   skills, handoff, and mode matrices.
4. The repository must avoid prompt-runtime coupling. If a runtime artifact is
   too thin, create a repair or assimilation slice; do not ask future agents to
   read the original compiler prompt.
5. Assimilation is not broad rewrite authorization. It is a controlled
   infrastructure slice inside `docs/ai/*`, with no dependency installation, no
   application skeleton, no business implementation, and no production action.

## Selected Strategy

Use `standard_light` for this repository-infrastructure slice. The work is
documentation and runtime-governance scoped, reversible, and does not touch
business code or external services. The slice still records evidence, updates
the validator, and adds a project skill because future compiler-spec updates
are fragile enough to deserve a repeatable workflow.

This strategy upgrades the current runtime in place:

- Add `docs/ai/skills/compiler-runtime-assimilation/SKILL.md`.
- Add `docs/ai/templates/compiler-runtime-assimilation-sop.md`.
- Add `docs/ai/architecture/compiler-runtime-assimilation.md`.
- Add `docs/ai/knowledge/entries/kb_008_compiler_runtime_assimilation.md`.
- Add `slice_007_compiler_runtime_assimilation` and synchronize status,
  tasks, evidence, manifest, changelog, and handoff.

The repair remains a Vibe Coding infrastructure task line. It does not start a
mainline product tree and does not close the existing infrastructure branch
without developer acceptance.
