# Vibe Infrastructure Action Workflow Skill

## Trigger Conditions

- User input starts exactly with `新增 Vibe Coding 仓库基建设施建立任务线：`.
- The substring after the first full-width colon is the required
  `{{ 优化动作 }}` value.

## Canonical Template

- Load and follow `docs/ai/templates/vibe-infra-action-workflow-sop.md`.
- Treat the template as the canonical SOP for this skill.
- Do not require the original compiler prompt or any raw chat prompt.

## Runtime Binding

1. Extract `{{ 优化动作 }}` from the triggering user message.
2. If `{{ 优化动作 }}` is empty, stop and ask the developer to provide the
   missing concrete action.
3. Execute only phase one from the template by default.
4. Phase one may analyze, plan, and produce a red-light confirmation request;
   it must not generate code, create directories, edit files, install
   dependencies, or write physical artifacts.
5. Enter phase two only after explicit human authorization that names or clearly
   confirms the selected strategy.

## Boundaries

- Applies only to Vibe Coding repository infrastructure work:
  `AGENTS.md`, `docs/ai/*`, schemas, scripts, runtime state, evidence,
  knowledge, skills, handoff, quality gates, and workflow documentation.
- Does not authorize product, business, application, dependency installation,
  production, money, destructive, or supply-chain work.
- Human confirmation in phase one is a hard gate. Ambiguous phrases such as
  "继续", "可以", or "看着办" are insufficient when the selected strategy,
  write scope, or risk boundary is unclear.

## Verification

- Phase one output ends at the red-light confirmation request.
- Phase two records the confirmed strategy, write scope, validation result,
  evidence, status, and handoff before claiming completion.
