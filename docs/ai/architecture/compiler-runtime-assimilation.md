# Compiler Runtime Assimilation Architecture

- ID: `architecture_003_compiler_runtime_assimilation`
- Date: `2026-05-29T21:30:00+08:00`
- Branch: `branch_vibe_coding_infra`
- Related requirement: `req_infra_007_compiler_runtime_assimilation`
- Related analysis: `docs/ai/analysis/compiler-runtime-assimilation.md`

## Purpose

This architecture defines how the repository absorbs future Vibe Coding
compiler, generator, or major governance specifications after the core runtime
already exists. The goal is to preserve the useful semantics while preventing
the raw source text from becoming a hidden runtime dependency.

## Boundary Model

| Layer | Meaning | Repository handling |
| --- | --- | --- |
| Source compiler context | A pasted or referenced generation specification, upgrade spec, or macro governance prompt. | Treat as external input and requirement source. Preserve only summarized, traceable facts in runtime artifacts. |
| Current runtime artifacts | `AGENTS.md`, `docs/ai/*`, task forest, status, requirements, evidence, knowledge, skills, schemas, and handoff. | This is the daily authority for future agents. Repair it directly when thin or stale. |
| Future operating behavior | How agents recover context, route work, confirm requirements, execute slices, validate, and close rounds. | Express through skills, SOPs, schemas, task trees, validator checks, and knowledge entries. |

## Assimilation Flow

```text
compiler_context_intake_gate
-> repository_fit_gate
-> semantic_layer_extraction_gate
-> runtime_artifact_translation_gate
-> no_prompt_runtime_dependency_recheck_gate
-> validation_and_evidence_gate
-> handoff_and_next_condition_gate
```

## Translation Rules

1. `compiler_context_intake_gate`: capture the source goal and whether the
   developer asked for analysis only or physical repository assignment.
2. `repository_fit_gate`: decide whether the work belongs to
   `branch_vibe_coding_infra`; compiler and generator governance always does
   unless it contains a separate confirmed product idea.
3. `semantic_layer_extraction_gate`: extract invariants, state-machine gates,
   physical artifact requirements, mode/acceptance matrices, and forbidden
   behaviors.
4. `runtime_artifact_translation_gate`: update only the smallest runtime
   artifacts needed to make the semantics recoverable: rules, schemas, skills,
   templates, knowledge, tasks, evidence, and handoff.
5. `no_prompt_runtime_dependency_recheck_gate`: confirm no file tells future
   agents to reread the raw compiler prompt as a normal recovery path.
6. `validation_and_evidence_gate`: run `python docs/ai/scripts/validate_runtime.py`
   and targeted text checks; record evidence.
7. `handoff_and_next_condition_gate`: synchronize `current.yaml`,
   `current-slice.yaml`, task tree, manifest, changelog, and handoff.

## Mode Recommendation

Use `standard_light` for the current assimilation because it updates only
repository governance files and does not touch production, data, dependencies,
or business implementation. Escalate to `risk_scaled_strict` if a future
compiler update changes agency boundaries, multi-agent delegation, production
actions, external data sources, knowledge trust policy, or skill permissions.
Escalate to `enterprise_high_assurance` if a compiler update affects money,
privacy, auth, compliance, irreversible operations, or team-wide release
governance.

## Artifact Contract

Every compiler-runtime assimilation slice must produce or update:

- A requirement ledger entry with the source request and normalized runtime
  meaning.
- A task tree entry and execution slice in `branch_vibe_coding_infra`.
- A skill or template when the behavior is repeatable.
- A knowledge entry when the distilled fact should survive context loss.
- A validator or manual check if the behavior can regress.
- Evidence and handoff updates before claiming completion.

The raw compiler context may be summarized in requirements and evidence, but it
must not become the daily recovery path.
