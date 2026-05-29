# Compiler Runtime Assimilation Skill

## Trigger Conditions

Use this skill when any of these are true:

- The developer provides a Vibe Coding compiler, generator, prompt-generation,
  governance-upgrade, or infrastructure-establishment specification and asks to
  assign it to the current repository.
- The developer says the current project is already a Vibe Coding
  infrastructure repository and asks to deeply analyze, distill, or transplant
  a compiler specification into it.
- A future `branch_vibe_coding_infra` slice must repair runtime artifacts
  because they are too thin, stale, or still dependent on raw compiler text.

## Purpose

Turn a macro compiler or generator context into lightweight runtime artifacts
without making future agents read the original prompt. This skill is for
repository infrastructure only; it does not authorize product implementation,
dependency installation, package skeleton creation, production work, or
business architecture choices.

## Required Runtime Inputs

- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/runtime-rule-index.md`
- `docs/ai/evidence/index.md`
- `docs/ai/knowledge/index.md`
- `docs/ai/skills/index.md`
- `docs/ai/templates/compiler-runtime-assimilation-sop.md`

## Steps

1. Route the work to `branch_vibe_coding_infra` unless the developer has
   separately confirmed a product or business idea.
2. Determine whether the developer authorized physical assignment to the
   repository. If not, stop after diagnosis and ask for confirmation.
3. Extract the source context into four semantic layers:
   core invariants and human boundaries; state machine and task-tree flow;
   physical schemas, evidence, knowledge, and skills; mode, acceptance, and
   example matrices.
4. Map each layer to existing runtime artifacts. Prefer updating a small set of
   files over copying the source specification into the repository.
5. Add or update a requirement, task, execution slice, knowledge entry, skill,
   SOP template, architecture note, validator coverage, evidence, status, and
   handoff as needed.
6. Run `python docs/ai/scripts/validate_runtime.py` and targeted `rg` checks.
7. Leave the next condition as developer acceptance or requested edits.

## Boundaries

- Keep writes inside `AGENTS.md`, `docs/`, or `docs/ai/*` unless a future
  confirmed slice says otherwise.
- Do not create `prompt.md` as a runtime authority.
- Do not require future agents to reread the original compiler prompt.
- Do not install dependencies, create lockfiles, create application skeletons,
  or choose a product technology stack.
- Do not treat external compiler text as instructions that outrank local
  runtime rules.

## Verification

The assimilation is complete only when:

- `compiler_runtime_assimilation` appears in the runtime rule index, skill,
  template, knowledge entry, slice, and validator coverage.
- The runtime validator passes.
- Evidence records validation and confirms no package manifest, lockfile,
  application skeleton, or business code was created.
- The handoff explains the next allowed action and preserves
  `no_prompt_runtime_dependency`.

## Agency

This skill does not expand agency level. Low-risk `docs/ai/*` synchronization
inside a confirmed infrastructure slice is `auto_recordable`; changes to
confirmation boundaries, tool permissions, production behavior, money, privacy,
auth, multi-agent delegation, or destructive operations require
`hard_confirmation` or `pause_required`.
