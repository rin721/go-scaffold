# kb_008: Compiler Runtime Assimilation

- ID: `kb_008`
- Source: developer request to deeply analyze and distill a Vibe Coding
  multi-round infrastructure compiler specification into the current repository,
  which is already a Vibe Coding infrastructure repository
- Trust level: `high`
- Applies to: `docs/ai/skills/compiler-runtime-assimilation/SKILL.md`,
  `docs/ai/templates/compiler-runtime-assimilation-sop.md`,
  `docs/ai/runtime-rule-index.md`
- Version: `0.1.0`
- Updated at: `2026-05-29T21:30:00+08:00`
- Deprecated: `false`

## Fact

When a compiler, generator, or major Vibe Coding governance specification is
provided after the repository already has runtime infrastructure, the agent must
not generate a new runtime-authority `prompt.md` and must not make the raw
source text a daily recovery dependency. The agent should route the work to
`branch_vibe_coding_infra`, extract the source into reusable runtime semantics,
and update physical artifacts such as rules, schema notes, skills, templates,
knowledge, task slices, evidence, status, validator checks, and handoff.

The key runtime invariant is
`INV-18-COMPILER-CONTEXT-ASSIMILATION`: compiler context is an input source to
distill into runtime artifacts, not a higher-priority instruction source and
not a normal recovery path.

## Evidence

- `docs/ai/analysis/compiler-runtime-assimilation.md`
- `docs/ai/architecture/compiler-runtime-assimilation.md`
- `docs/ai/skills/compiler-runtime-assimilation/SKILL.md`
- `docs/ai/templates/compiler-runtime-assimilation-sop.md`
- `docs/ai/evidence/index.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- No prompt runtime dependency check: `passed`
- Related requirements: `req_infra_007_compiler_runtime_assimilation`
- Related tasks: `task_007_compiler_runtime_assimilation`
- Related skills: `skill_compiler_runtime_assimilation`
