# research_001: Business Mainline Capability Groups

- Round: `infra_001`
- Branch: `branch_vibe_coding_infra`
- Checked at: `2026-05-29T19:07:53+08:00`
- Trigger: developer confirmed green-light A for phase two, catalog-only, with
  a requirement to verify current official sources using the current year
  dynamically instead of hard-coding a year.
- Status: `candidate_refreshed`

## Current-Year Policy

When a future prompt asks for current-year research, derive the current year
from the runtime date or the `Checked at` timestamp. Do not write a fixed
calendar year into reusable prompts, skills, or selection protocols.

## Question

Which current UI, frontend architecture, AI Agent script, Python tooling, and
agent framework capabilities should be prepared as selectable capability groups
for the future business mainline?

## Sources

| Source | Trust | Used for |
| --- | --- | --- |
| https://react.dev/learn/start-a-new-react-project | official | React project direction and framework/build-tool framing |
| https://vite.dev/guide/ | official | Vite React TypeScript candidate baseline |
| https://ui.shadcn.com/docs/installation | official | shadcn/ui setup modes and supported templates |
| https://ui.shadcn.com/docs/installation/vite | official | shadcn/ui Vite and Tailwind v4 setup path |
| https://www.radix-ui.com/primitives/docs | official | accessible primitive UI architecture |
| https://tailwindcss.com/docs/installation/using-vite | official | Tailwind CSS Vite plugin path |
| https://tanstack.com/query/docs/docs | official | server-state management candidate |
| https://reactrouter.com/start/data/routing | official | React routing candidate |
| https://storybook.js.org/docs/10/get-started | official | component workshop and UI state docs |
| https://pnpm.io/ | official | JavaScript package-manager candidate |
| https://zod.dev/ | official | TypeScript runtime schema validation candidate |
| https://biomejs.dev/linter | official | JavaScript/TypeScript linting candidate |
| https://docs.astral.sh/uv/ | official | Python package, project, script, and tool runner candidate |
| https://docs.astral.sh/uv/concepts/projects/dependencies/ | official | Python dependency-management candidate |
| https://docs.pytest.org/en/stable/getting-started.html | official | Python test framework |
| https://docs.astral.sh/ruff/linter/ | official | Python linting |
| https://docs.astral.sh/ruff/formatter/ | official | Python formatting |
| https://typer.tiangolo.com/ | official | Python CLI tooling |
| https://docs.pydantic.dev/latest/concepts/models/ | official | structured data validation |
| https://playwright.dev/python/docs/intro | official | browser automation candidate |
| https://platform.openai.com/docs/guides/agents-sdk/ | official | OpenAI Agents SDK purpose and official SDK docs entry |
| https://openai.github.io/openai-agents-python/agents/ | official | OpenAI agent orchestration |
| https://openai.github.io/openai-agents-python/sessions/ | official | OpenAI Agents SDK session memory candidate |
| https://openai.github.io/openai-agents-python/tracing/ | official | OpenAI Agents SDK tracing behavior |
| https://pydantic.dev/docs/ai/overview/ | official | Python-first agent framework |
| https://docs.langchain.com/oss/python/langgraph/workflows-agents | official | graph/stateful agent workflows |
| https://docs.langchain.com/oss/python/langgraph/persistence | official | LangGraph persistence, memory, HITL, and recovery |

## Refresh Findings

- React official docs recommend starting new apps with a framework. For this
  repository, Vite remains a candidate when the future confirmed requirement is
  a SPA, local browser tool, or scratch frontend rather than a full-stack app.
- shadcn/ui current official docs use `ui.shadcn.com`, support shadcn/create,
  CLI templates, and existing-project paths, and the Vite path uses
  Tailwind CSS with `@tailwindcss/vite`.
- Radix Primitives remain a candidate accessibility primitive layer; shadcn/ui
  remains a candidate copy-owned component workflow rather than a dependency
  adoption decision.
- Add `uv` as the default Python tooling candidate, but do not let it create
  `pyproject.toml`, `uv.lock`, virtual environments, or project files until a
  future confirmed slice allows that.
- Add `Zod 4` as a TypeScript runtime schema candidate for forms, configs,
  API payloads, and Agent IO after a frontend/TypeScript boundary exists.
- Keep OpenAI Agents SDK, Pydantic AI, and LangGraph as distinct candidates:
  OpenAI-native orchestration; Python type-safe/provider-flexible agents; and
  persistent graph/HITL workflows respectively.

## Conclusion

Create candidate capability groups, not installed dependencies. The current
repository has no frontend or Python package boundary. Installing packages now
would create premature lockfiles and technical commitment before a business
mainline requirement exists.

## Recommended Catalog

1. UI library and design system: `shadcn/ui + Radix UI + Tailwind CSS v4`.
2. Frontend architecture: React framework-first decision; `React + TypeScript +
   Vite` for SPA/local-tool boundaries with explicit module boundaries.
3. AI Agent script tools: Python CLI and validation scripts with typed IO,
   keeping `uv` as a candidate runner only after Python boundaries are confirmed.
4. Python tooling: `uv + pytest + ruff + typer + pydantic`, with Playwright only
   when browser automation is needed.
5. Agent frameworks: `OpenAI Agents SDK` or `Pydantic AI` for agent runtime;
   `LangGraph` only when explicit stateful graph workflows are needed.
6. TypeScript schema validation: `Zod 4` only after a TypeScript boundary exists.
7. Quality verification: every capability adoption must add validation and
   evidence.

## Confirmation Boundary

This research supports strategy A. Installing dependencies, creating lockfiles,
or generating app/tool skeletons remains blocked until a future explicit
confirmation chooses strategy B or C.
