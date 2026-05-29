# Task Forest Branching Analysis

- Analysis id: `analysis_001_task_forest_branching`
- Date: `2026-05-29T17:54:03+08:00`
- Source: developer request
- Status: `implemented`

## Original Sentence

> 将当前任务树林，实现新能力，像以上的操作实则是在迭代Vibe Coding仓库，而不是提出想法树建立主线任务树，所以应该为Vibe Coding建立一条单独的任务树支线，用于Vibe Coding仓库的基建设施建立

## Problem In The Existing Runtime

The previous runtime had only one visible round task tree:
`docs/ai/tasks/tree.yaml`. Infrastructure tasks such as documentation updates,
runtime validator replacement, schema edits, and rule upgrades were placed into
that tree. That makes the repository look as if these infrastructure changes
were product or business idea work.

This weakens the workflow because future agents cannot cleanly distinguish:

- a developer's product or business idea mainline;
- Vibe Coding runtime infrastructure work for this repository;
- bootstrap work that already happened;
- maintenance or repair slices for the agent-driving system itself.

## Optimized Rule Statement

Use a `TaskForest` instead of a single ambiguous task tree.

The `mainline_idea` branch is reserved for developer product, feature, business,
or application ideas after requirement confirmation. The `vibe_coding_infra`
branch is reserved for Vibe Coding repository infrastructure work, including
runtime docs, `AGENTS.md`, `docs/ai/*`, schemas, quality gates, scripts,
knowledge entries, project skills, handoff, and state recovery.

If a task changes how this repository is governed, recovered, validated, or
driven by agents, route it to `branch_vibe_coding_infra`. Do not start or pollute
the mainline idea task tree with those infrastructure tasks.

## Operational Consequence

- `docs/ai/tasks/forest.yaml` becomes the task forest entry.
- `docs/ai/tasks/main-tree.yaml` remains available for future product or
  business idea work.
- `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml` owns current Vibe Coding
  infrastructure tasks.
- `docs/ai/tasks/tree.yaml` is retained as a compatibility pointer to the task
  forest, not as the only source of task truth.

