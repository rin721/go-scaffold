# kb_005: Business Mainline Capability Groups

- ID: `kb_005`
- Source: official documentation research, strategy A confirmation, and
  dynamic current-year source refresh
- Trust level: `high`
- Applies to: `docs/ai/capabilities/*`, future `branch_mainline_idea`
- Version: `0.2.0`
- Updated at: `2026-05-29T19:07:53+08:00`
- Deprecated: `false`

## Fact

The repository has a candidate capability catalog for future business mainline
work. It includes UI library candidates, frontend architecture candidates,
AI Agent script tools, Python tooling, agent framework candidates, and quality
verification requirements. These are not installed dependencies.

The catalog must treat "current-year" research as dynamic: future agents should
derive the current year from the runtime date or source refresh timestamp rather
than writing a fixed year into reusable prompts. The refreshed candidate set
adds `uv` for Python tooling and `Zod 4` for future TypeScript runtime schema
validation, while preserving the no-install/no-scaffold boundary.

## Evidence

- `docs/ai/capabilities/index.md`
- `docs/ai/capabilities/dependency-candidates.yaml`
- `docs/ai/research/research_001_capability_groups.md`
- `docs/ai/architecture/capability-group-strategy.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- Dynamic current-year wording check: `passed`
- Related requirements: `req_infra_004_capability_groups`
- Related tasks: `task_004_capability_groups`
- Related skills: `skill_capability_selection`
