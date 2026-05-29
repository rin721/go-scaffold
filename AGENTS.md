# Agent Runtime Entry

This repository uses lightweight Vibe Coding runtime artifacts compiled from
`vibe-coding-infra-compiler-v1.8.0`.

## Role

Agents must treat this file as a short runtime index. It is not a replacement
for the structured state under `docs/ai/`.

## Rule Priority

1. Safety, privacy, permissions, supply chain, production, money, destructive
   operations, and developer confirmations.
2. Repository runtime state, current slice, task tree, requirements, decisions,
   evidence, knowledge entries, and project skills.
3. Existing project conventions and local directory rules.
4. Best practices and templates.

## Runtime Entries

- Runtime rule index: `docs/ai/runtime-rule-index.md`
- Current state: `docs/ai/status/current.yaml`
- Current slice: `docs/ai/tasks/current-slice.yaml`
- Task forest: `docs/ai/tasks/forest.yaml`
- Mainline idea tree: `docs/ai/tasks/main-tree.yaml`
- Vibe Coding infrastructure tree: `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- Bootstrap tree: `docs/ai/tasks/bootstrap-tree.yaml`
- Requirements: `docs/ai/requirements/ledger.yaml`
- Decisions: `docs/ai/decisions/records.md`
- Evidence: `docs/ai/evidence/index.md`
- Knowledge base: `docs/ai/knowledge/index.md`
- Project skills: `docs/ai/skills/index.md`
- Handoff: `docs/ai/handoff/current.md`

## No Prompt Runtime Dependency

Do not require future agents to read the original compiler prompt or any raw
chat prompt as a normal recovery path. If a runtime artifact is missing or too
thin, create or run a repair slice that updates the physical artifact.
