# Capability Selection Skill

## Trigger Conditions

- A future `branch_mainline_idea` requirement needs UI, frontend architecture,
  Python tools, Agent scripts, or agent framework capabilities.
- The developer asks what libraries, tools, or skills should be enabled.
- A task proposes installing dependencies or creating engineering skeletons.
- The developer asks to research current-year/latest technology choices.

## Inputs

- `docs/ai/capabilities/index.md`
- `docs/ai/capabilities/dependency-candidates.yaml`
- `docs/ai/capabilities/groups/*.yaml`
- `docs/ai/research/research_001_capability_groups.md`
- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`

## Outputs

- A proposed capability group selection.
- A confirmation request for any install, lockfile, skeleton, or package-manager
  choice.
- A future execution slice if the developer confirms adoption.

## Boundaries

- Do not install dependencies from the catalog alone.
- Do not create `package.json`, `pyproject.toml`, `src/`, `tools/`, or lockfiles
  unless a future confirmed slice allows it.
- Do not treat candidate tools as already adopted project dependencies.
- Do not mix Vibe Coding infrastructure catalog work into the business mainline
  task tree.
- Do not hard-code a calendar year into reusable research prompts. Derive the
  current year from the runtime date or source refresh timestamp.

## Steps

1. Read the confirmed business mainline requirement.
2. Read the capability catalog.
3. Refresh only the relevant official sources, using the runtime date/current
   checked_at for current-year wording.
4. Select the smallest capability groups that satisfy the requirement.
5. Explain why lighter and heavier options were not selected.
6. Ask for hard confirmation before dependency installation or skeleton creation.
7. After confirmation, create a bounded execution slice with validation and
   rollback.

## Verification

- Proposed groups map to a confirmed requirement.
- Current-source refresh is recorded with a timestamp, not a hard-coded year.
- Install/scaffold actions are gated by explicit confirmation.
- Evidence and status updates are planned before execution.
