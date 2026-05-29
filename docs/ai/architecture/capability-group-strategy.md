# Capability Group Architecture Strategy

- Architecture id: `arch_001_capability_groups`
- Branch: `branch_vibe_coding_infra`
- Status: `candidate_catalog_ready`
- Updated at: `2026-05-29T19:07:53+08:00`

## Goal

Prepare selectable capability groups for the future `branch_mainline_idea`
without prematurely installing dependencies or creating a business application
skeleton.

## Dynamic Current-Year Rule

Future research prompts must say "current official sources" or derive the
current year from the runtime date/source refresh timestamp. Do not hard-code a
calendar year in reusable skills, catalog protocols, or selection prompts.

## Architecture

```text
branch_vibe_coding_infra
  -> docs/ai/capabilities/index.md
  -> docs/ai/capabilities/dependency-candidates.yaml
  -> docs/ai/capabilities/groups/*.yaml
  -> docs/ai/research/research_001_capability_groups.md
  -> docs/ai/knowledge/entries/kb_005_capability_groups.md
  -> docs/ai/skills/capability-selection/SKILL.md

branch_mainline_idea
  -> future confirmed requirement
  -> refresh current official sources using runtime date
  -> choose capability groups
  -> create execution slice
  -> request hard confirmation for install/scaffold
```

## Selection Flow

1. Developer confirms a business mainline goal.
2. Agent reads the capability catalog.
3. Agent refreshes relevant official sources using the runtime date/current
   checked_at rather than a hard-coded year.
4. Agent proposes only the groups needed by the goal.
5. Developer confirms selected groups and package manager.
6. Agent creates a new execution slice with install commands, allowed files,
   verification, rollback, and evidence targets.
7. Only then may dependencies or skeleton files be created.

## Refreshed Strategy Notes

- UI: keep `shadcn/ui + Radix UI + Tailwind CSS v4` as the default candidate
  for copy-owned, accessible, design-system-friendly React UI.
- Frontend architecture: decide React framework first; use Vite when the future
  requirement is a SPA, static/local web tool, or scratch React boundary.
- Python tooling: add `uv` as a candidate, but treat `uv init`, `uv add`,
  `uv lock`, and virtual environment creation as future confirmed actions.
- TypeScript validation: add `Zod 4` as a candidate for runtime schemas after a
  TypeScript boundary exists.
- Agent runtime: keep direct scripts first; select OpenAI Agents SDK,
  Pydantic AI, or LangGraph only when the future product requirement needs that
  level of orchestration, typed agent IO, persistence, or HITL.

## Non-Goals

- No `package.json`.
- No `pyproject.toml`.
- No `src/`, `app/`, `components/`, `tools/`, or `tests/` skeleton.
- No dependency installation.
- No lockfile.

## Quality Gates

- `python docs/ai/scripts/validate_runtime.py`
- Verify no forbidden skeleton or package files were created.
- Verify capability groups are referenced from knowledge, skills, and handoff.
- Verify current-year research wording is dynamic and source-refresh based.
