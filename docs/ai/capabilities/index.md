# Capability Groups Index

- Catalog id: `capability_catalog_001`
- Branch: `branch_vibe_coding_infra`
- Created at: `2026-05-29T18:50:41+08:00`
- Last official source refresh: `2026-05-29T19:07:53+08:00`
- Status: `candidate_catalog_ready`
- Strategy: `confirmed_A_catalog_only`
- Current-year policy: derive the current year from the source refresh timestamp
  or runtime date; do not hard-code a calendar year into future research prompts.

This catalog lists optional capability groups that the future business mainline
can choose from. It does not install dependencies, create a frontend project, or
create a Python project skeleton.

## Red Light Boundary

| Action | Status |
| --- | --- |
| Create candidate capability catalog | allowed |
| Create skills and knowledge entries for selection | allowed |
| Create `package.json`, `pyproject.toml`, `src/`, `tools/`, or lockfiles | not allowed in strategy A |
| Run `npm install`, `pnpm add`, `pip install`, or similar dependency installation | not allowed in strategy A |
| Add secrets, API keys, production config, or external service credentials | not allowed |

## Groups

| ID | Group | Default recommendation | When to choose | Path |
| --- | --- | --- | --- | --- |
| `cap_ui_001` | UI library and design system | `shadcn/ui + Radix UI + Tailwind CSS v4` | Product UI, dashboards, internal tools, design-system-friendly apps | `docs/ai/capabilities/groups/ui-library.yaml` |
| `cap_arch_001` | Frontend engineering architecture | React framework-first decision; `Vite` for SPA/local tools | Future browser app or local web tool | `docs/ai/capabilities/groups/frontend-architecture.yaml` |
| `cap_agent_scripts_001` | AI Agent script tools | Python CLI/runbook tools with typed IO | Repo automation, validation, doc checks, status repair | `docs/ai/capabilities/groups/agent-script-tools.yaml` |
| `cap_py_tools_001` | Python engineering tools | `uv + pytest + ruff + typer + pydantic` | Python utilities, CLI tools, validators, AI support scripts | `docs/ai/capabilities/groups/python-tooling.yaml` |
| `cap_agent_frameworks_001` | Agent framework candidates | `OpenAI Agents SDK` or `Pydantic AI`; `LangGraph` for stateful workflows | Tool-calling agents, structured AI workflows, HITL loops | `docs/ai/capabilities/groups/agent-frameworks.yaml` |
| `cap_quality_001` | Quality and verification | Runtime validator + future test/build/lint gates | Every capability adoption | `docs/ai/capabilities/groups/quality-verification.yaml` |

## Use Protocol

1. Start from a confirmed `branch_mainline_idea` requirement.
2. Select only the capability groups needed for that requirement.
3. Re-check official sources at execution time, deriving "current year" from the
   runtime date or the source refresh timestamp.
4. Record the chosen groups in a decision record.
5. Create a new execution slice with allowed files, dependencies, commands, and rollback plan.
6. Ask for hard confirmation before installing dependencies or creating lockfiles.
