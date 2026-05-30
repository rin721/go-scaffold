# Evidence Index

## evidence_001: Readonly Preflight

- Date: `2026-05-29T17:18:17+08:00`
- Type: `preflight_readonly`
- Status: `completed`
- Commands:
  - `Get-Location; Get-ChildItem -Force | Select-Object Mode,Length,Name`
  - `git status --short --branch`
  - `rg --files ...`
  - `rg -n "TODO|FIXME|Vibe Coding|prompt\.md|AGENTS|test|lint|build" ...`
- Summary: Repository initially contained only `.git`; no existing runtime
  rules, docs, dependencies, tests, CI, or business files were found.
- Related artifacts:
  - `docs/ai/reports/preflight-readonly.md`

## evidence_002: Bootstrap Write Scope

- Date: `2026-05-29T17:18:17+08:00`
- Type: `authorization`
- Status: `completed`
- Summary: Bootstrap writes are constrained to `AGENTS.md`, `docs/.cursorrules`,
  and `docs/ai/**`.

## evidence_010: Bootstrap Runtime Verification

- Date: `2026-05-29T17:29:21+08:00`
- Type: `runtime_validation`
- Status: `completed`
- Command: legacy runtime validation command, superseded by `python docs/ai/scripts/validate_runtime.py`
- Result: `Runtime validation passed.`
- Summary: Required runtime artifacts exist, status points to the bootstrap
  slice, and no runtime entry requires reading the original compiler prompt.

## evidence_011: Validation Command Portability Fix

- Date: `2026-05-29T17:29:21+08:00`
- Type: `physical_readability_issue`
- Status: `resolved`
- Finding: `pwsh` was not available in the local Windows environment.
- Adjustment: Runtime validation now uses the Python validator.
- Related invariant: `INV-16`

## evidence_012: Human Documentation Slice

- Date: `2026-05-29T17:37:15+08:00`
- Type: `documentation_update`
- Status: `completed`
- Requirement: `req_round_001_docs`
- Slice: `slice_001_docs_guides`
- Files:
  - `README.md`
  - `docs/index.md`
  - `docs/project-overview.md`
  - `docs/usage-guide.md`
  - `docs/maintenance-guide.md`
  - `docs/agent-workflow-guide.md`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - File existence checks -> all six human documentation files returned `OK`.
  - Markdown local link check -> `Markdown local links OK`.
- Summary: Human-facing tutorial and overview documentation exists and runtime validation passes.

## evidence_013: Runtime Validator Generalized

- Date: `2026-05-29T17:37:15+08:00`
- Type: `quality_gate_update`
- Status: `completed`
- Finding: The validator still assumed the bootstrap slice was current.
- Adjustment: The validator now checks that `status/current.yaml` `current_slice`
  matches `tasks/current-slice.yaml` `id`.
- Validation: `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`

## evidence_014: Python Runtime Script Replacement

- Date: `2026-05-29T17:44:56+08:00`
- Type: `runtime_script_update`
- Status: `completed`
- Requirement: `req_round_001_python_scripts`
- Slice: `slice_002_python_scripts`
- Files:
  - `docs/ai/scripts/validate_runtime.py`
  - `README.md`
  - `docs/usage-guide.md`
  - `docs/maintenance-guide.md`
  - `docs/ai/manifest.yaml`
  - `docs/ai/quality/gates.md`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/slices/slice_002_python_scripts.yaml`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `python -m py_compile docs/ai/scripts/validate_runtime.py` -> passed.
  - script inventory check under `docs/ai/scripts` -> only `validate_runtime.py` found.
  - Markdown local link check -> `Markdown local links OK`.
- Summary: Runtime validator was replaced with Python and command references were updated.

## evidence_015: Task Forest Branching

- Date: `2026-05-29T17:58:59+08:00`
- Type: `runtime_task_governance_update`
- Status: `completed`
- Requirement: `req_infra_003_task_forest_branching`
- Slice: `slice_003_task_forest_branching`
- Analysis:
  - `docs/ai/analysis/task-forest-branching.md`
- Sandbox:
  - `docs/ai/sandbox/sandbox_001_task_forest_branching.yaml`
- Files:
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/tasks/main-tree.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/tree.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/slices/slice_003_task_forest_branching.yaml`
  - `docs/ai/status/current.yaml`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/scripts/validate_runtime.py`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - task forest routing search for `branch_vibe_coding_infra`, `branch_mainline_idea`, and `TaskForest` -> matched status, task forest, branch tree, main tree, compatibility pointer, slices, and runtime rule index.
- Summary: Task forest and Vibe Coding infrastructure branch are implemented and validated.

## evidence_016: Catalog-Only Capability Groups

- Date: `2026-05-29T18:53:11+08:00`
- Type: `runtime_capability_catalog_update`
- Status: `completed`
- Requirement: `req_infra_004_capability_groups`
- Slice: `slice_004_capability_groups`
- Research:
  - `docs/ai/research/research_001_capability_groups.md`
- Architecture:
  - `docs/ai/architecture/capability-group-strategy.md`
- Files:
  - `docs/ai/capabilities/index.md`
  - `docs/ai/capabilities/dependency-candidates.yaml`
  - `docs/ai/capabilities/groups/ui-library.yaml`
  - `docs/ai/capabilities/groups/frontend-architecture.yaml`
  - `docs/ai/capabilities/groups/agent-script-tools.yaml`
  - `docs/ai/capabilities/groups/python-tooling.yaml`
  - `docs/ai/capabilities/groups/agent-frameworks.yaml`
  - `docs/ai/capabilities/groups/quality-verification.yaml`
  - `docs/ai/skills/capability-selection/SKILL.md`
  - `docs/ai/knowledge/entries/kb_005_capability_groups.md`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `Test-Path package.json, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
  - `rg -n "capability_catalog_001|req_infra_004_capability_groups|slice_004_capability_groups|skill_capability_selection|kb_005" docs/ai` -> matched catalog, requirement, slice, skill, and knowledge entries.
  - Markdown local link check -> `Markdown local links OK`.
- Summary: Catalog-only capability groups were added without dependency
  installation, package manifests, lockfiles, or engineering skeleton creation.

## evidence_017: Dynamic Current-Year Capability Refresh

- Date: `2026-05-29T19:07:53+08:00`
- Type: `runtime_capability_catalog_refresh`
- Status: `completed`
- Requirement: `req_infra_004_capability_groups`
- Slice: `slice_004_capability_groups`
- Trigger: Developer green-lighted phase two strategy A and required official
  source verification with current-year wording derived dynamically rather than
  hard-coded.
- Official source refresh:
  - React, Vite, shadcn/ui, Tailwind CSS, Radix UI, React Router, TanStack Query, Storybook
  - pnpm, Zod, Biome
  - uv, pytest, Ruff, Typer, Pydantic, Playwright Python
  - OpenAI Agents SDK, Pydantic AI, LangGraph
- Files:
  - `docs/ai/capabilities/index.md`
  - `docs/ai/capabilities/dependency-candidates.yaml`
  - `docs/ai/capabilities/groups/ui-library.yaml`
  - `docs/ai/capabilities/groups/frontend-architecture.yaml`
  - `docs/ai/capabilities/groups/agent-script-tools.yaml`
  - `docs/ai/capabilities/groups/python-tooling.yaml`
  - `docs/ai/capabilities/groups/agent-frameworks.yaml`
  - `docs/ai/capabilities/groups/quality-verification.yaml`
  - `docs/ai/research/research_001_capability_groups.md`
  - `docs/ai/architecture/capability-group-strategy.md`
  - `docs/ai/skills/capability-selection/SKILL.md`
  - `docs/ai/knowledge/entries/kb_005_capability_groups.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/slices/slice_004_capability_groups.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/decisions/records.md`
  - `docs/ai/handoff/current.md`
  - `docs/ai/manifest.yaml`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `Test-Path package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools, tests` -> all returned `False`.
  - `rg -n "current_year_policy|Current-Year Policy|official_source_refresh|last_official_source_refresh|uv|zod|ui.shadcn.com" docs/ai` -> matched refreshed catalog, group, research, architecture, knowledge, status, and slice entries.
  - `rg -n "v3.shadcn.com|2026 年当前官方资料|2026年当前官方资料" docs/ai/capabilities docs/ai/research docs/ai/architecture docs/ai/skills docs/ai/knowledge` -> no matches.
- Summary: Catalog-only capability groups were refreshed against current
  official sources, dynamic current-year wording was encoded in the runtime
  artifacts, and no dependency installation, package manifest, lockfile, or
  engineering skeleton was created.

## evidence_018: Trigger-Bound Vibe Infrastructure Action Workflow SOP

- Date: `2026-05-29T19:32:05+08:00`
- Type: `runtime_workflow_sop_update`
- Status: `completed`
- Requirement: `req_infra_005_vibe_infra_action_workflow`
- Slice: `slice_005_vibe_infra_action_workflow`
- Files:
  - `docs/ai/skills/vibe-infra-action-workflow/SKILL.md`
  - `docs/ai/templates/vibe-infra-action-workflow-sop.md`
  - `docs/ai/knowledge/entries/kb_006_vibe_infra_action_workflow.md`
  - `docs/ai/skills/index.md`
  - `docs/ai/knowledge/index.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/tasks/slices/slice_005_vibe_infra_action_workflow.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/decisions/records.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/handoff/current.md`
  - `docs/ai/changelog.md`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `rg -n "新增 Vibe Coding 仓库基建设施建立任务线|\\{\\{ 优化动作 \\}\\}|强制红绿灯|阶段二：物理产出" docs/ai/skills docs/ai/templates docs/ai/runtime-rule-index.md` -> matched the runtime rule index, project skill, skill index, and SOP template.
  - `Test-Path package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
- Summary: A trigger-bound project skill and canonical Markdown SOP template now
  enforce phase-one four-dimensional diagnosis, a red-light hard stop, and
  phase-two physical output only after explicit human authorization.

## evidence_019: Full-Project Lifecycle Workflow

- Date: `2026-05-29T21:12:37+08:00`
- Type: `runtime_lifecycle_workflow_update`
- Status: `completed`
- Requirement: `req_infra_006_full_project_lifecycle`
- Slice: `slice_006_full_project_lifecycle`
- Analysis:
  - `docs/ai/analysis/mainline-full-project-lifecycle-gap-analysis.md`
- Architecture:
  - `docs/ai/architecture/full-project-lifecycle-workflow.md`
- Files:
  - `docs/ai/skills/full-project-lifecycle/SKILL.md`
  - `docs/ai/templates/full-project-lifecycle-workflow-sop.md`
  - `docs/ai/knowledge/entries/kb_007_full_project_lifecycle.md`
  - `docs/ai/skills/index.md`
  - `docs/ai/knowledge/index.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/tasks/slices/slice_006_full_project_lifecycle.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/main-tree.yaml`
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/decisions/records.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/handoff/current.md`
  - `docs/ai/changelog.md`
  - `docs/ai/scripts/validate_runtime.py`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `rg -n "full_project_lifecycle|task_006_full_project_lifecycle|kb_007|skill_full_project_lifecycle|standard_light_risk_escalated|enterprise_high_assurance" docs/ai` -> matched the status, manifest, requirements, runtime rule index, schema, skill index, knowledge index, slice, task tree, analysis, architecture, SOP, knowledge entry, and validator.
  - Validator coverage now checks all 18 lifecycle gate identifiers, the two
    mode identifiers, and the round `n+1` loop identifier.
  - `Test-Path package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
- Summary: A full-project lifecycle workflow now requires future mainline or
  project task lines to pass through task-line routing, requirement analysis and
  confirmation, current-source research and technology confirmation, task
  analysis and confirmation, architecture and stack confirmation, infrastructure
  mode recommendation and confirmation, Agent driving infrastructure, task tree
  and execution slices, implementation, tests, documentation, state/evidence
  sync, next-step continuation, closure, and round `n+1`.

## evidence_020: Compiler Runtime Assimilation Workflow

- Date: `2026-05-29T21:30:00+08:00`
- Type: `runtime_compiler_assimilation_update`
- Status: `completed`
- Requirement: `req_infra_007_compiler_runtime_assimilation`
- Slice: `slice_007_compiler_runtime_assimilation`
- Analysis:
  - `docs/ai/analysis/compiler-runtime-assimilation.md`
- Architecture:
  - `docs/ai/architecture/compiler-runtime-assimilation.md`
- Files:
  - `docs/ai/skills/compiler-runtime-assimilation/SKILL.md`
  - `docs/ai/templates/compiler-runtime-assimilation-sop.md`
  - `docs/ai/knowledge/entries/kb_008_compiler_runtime_assimilation.md`
  - `docs/ai/skills/index.md`
  - `docs/ai/knowledge/index.md`
  - `docs/ai/architecture/index.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/tasks/slices/slice_007_compiler_runtime_assimilation.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/authorizations/records.md`
  - `docs/ai/decisions/records.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/handoff/current.md`
  - `docs/ai/changelog.md`
  - `docs/ai/scripts/validate_runtime.py`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `rg -n "compiler_runtime_assimilation|INV-18|compiler_context_intake_gate|no_prompt_runtime_dependency_recheck_gate|skill_compiler_runtime_assimilation|kb_008" docs/ai` -> matched the runtime rule index, manifest, requirement, status, task tree, current slice, slice file, architecture index and note, analysis, skill index, project skill, template, knowledge index, knowledge entry, evidence, handoff, rounds file, and validator.
  - `Test-Path prompt.md, package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
- Summary: The repository now has a compiler-runtime assimilation task line for
  distilling compiler, generator, or macro governance specifications into local
  runtime artifacts without creating `prompt.md` or restoring raw compiler
  prompt dependency.

## evidence_021: Project Requirement Discovery Workflow

- Date: `2026-05-29T21:48:21+08:00`
- Type: `runtime_requirement_discovery_workflow_update`
- Status: `completed`
- Requirement: `req_infra_008_project_requirement_discovery`
- Slice: `slice_008_project_requirement_discovery`
- Analysis:
  - `docs/ai/analysis/project-requirement-discovery-gap-analysis.md`
- Architecture:
  - `docs/ai/architecture/project-requirement-discovery-workflow.md`
- Files:
  - `docs/ai/skills/project-requirement-discovery/SKILL.md`
  - `docs/ai/templates/project-requirement-discovery-sop.md`
  - `docs/ai/templates/project-requirement-discovery-record.yaml`
  - `docs/ai/requirements/intake/README.md`
  - `docs/ai/knowledge/entries/kb_009_project_requirement_discovery.md`
  - `docs/ai/skills/full-project-lifecycle/SKILL.md`
  - `docs/ai/templates/full-project-lifecycle-workflow-sop.md`
  - `docs/ai/knowledge/entries/kb_007_full_project_lifecycle.md`
  - `docs/ai/skills/index.md`
  - `docs/ai/knowledge/index.md`
  - `docs/ai/architecture/index.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/tasks/slices/slice_008_project_requirement_discovery.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/main-tree.yaml`
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/decisions/records.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/handoff/current.md`
  - `docs/ai/changelog.md`
  - `docs/ai/scripts/validate_runtime.py`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `python -m py_compile docs/ai/scripts/validate_runtime.py` -> passed.
  - `rg -n "project_requirement_discovery|idea_seed_intake_gate|domain_surface_mapping_gate|question_backlog_gate|requirement_discovery_record|skill_project_requirement_discovery|kb_009" docs/ai` -> matched the runtime rule index, manifest, status, requirements, task tree, current slice, slice file, analysis, architecture, skill index, project skill, SOP, record template, knowledge index, knowledge entry, evidence, handoff, and validator.
  - `Test-Path prompt.md, package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
- Summary: The repository now has a project requirement discovery task line for
  turning raw ideas into persistent requirement discovery records, domain maps,
  question backlogs, requirement plans, and baseline readiness checks before
  requirement confirmation.

## evidence_022: Project Lifecycle Downstream Detailing Workflow

- Date: `2026-05-29T22:00:59+08:00`
- Type: `runtime_downstream_lifecycle_workflow_update`
- Status: `completed`
- Requirement: `req_infra_009_project_lifecycle_downstream_detail`
- Slice: `slice_009_project_lifecycle_downstream_detail`
- Analysis:
  - `docs/ai/analysis/project-lifecycle-downstream-gap-analysis.md`
- Architecture:
  - `docs/ai/architecture/project-lifecycle-downstream-workflow.md`
- Files:
  - `docs/ai/skills/project-lifecycle-downstream-detailing/SKILL.md`
  - `docs/ai/templates/project-lifecycle-downstream-gates-sop.md`
  - `docs/ai/templates/project-lifecycle-downstream-record.yaml`
  - `docs/ai/lifecycle/README.md`
  - `docs/ai/knowledge/entries/kb_010_project_lifecycle_downstream.md`
  - `docs/ai/skills/full-project-lifecycle/SKILL.md`
  - `docs/ai/templates/full-project-lifecycle-workflow-sop.md`
  - `docs/ai/knowledge/entries/kb_007_full_project_lifecycle.md`
  - `docs/ai/skills/index.md`
  - `docs/ai/knowledge/index.md`
  - `docs/ai/architecture/index.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/tasks/slices/slice_009_project_lifecycle_downstream_detail.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/main-tree.yaml`
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/decisions/records.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/handoff/current.md`
  - `docs/ai/changelog.md`
  - `docs/ai/scripts/validate_runtime.py`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `rg -n "project_lifecycle_downstream|research_plan_gate|task_graph_analysis_gate|architecture_dossier_gate|implementation_iteration_ledger_gate|acceptance_closure_packet_gate|skill_project_lifecycle_downstream_detailing|kb_010" docs/ai` -> matched the runtime rule index, manifest, requirements, status, task tree, current slice, slice file, analysis, architecture, skill index, project skill, SOP, record template, knowledge index, knowledge entry, evidence, handoff, and validator.
  - `Test-Path prompt.md, package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
- Summary: The repository now has a downstream lifecycle detailing task line for
  turning post-baseline stages into durable research, task analysis,
  architecture, mode, Agent infrastructure, slice, implementation, verification,
  closure, and next-round records instead of narrative summaries.

## evidence_023: Strategy C Requirement Workflow Engine

- Date: `2026-05-29T22:24:00+08:00`
- Type: `runtime_requirement_workflow_engine_update`
- Status: `completed`
- Requirement: `req_infra_010_requirement_workflow_engine`
- Slice: `slice_010_requirement_workflow_engine`
- Trigger: Developer confirmed strategy C and phase-two execution.
- Analysis:
  - `docs/ai/analysis/requirement-workflow-engine-gap-analysis.md`
- Architecture:
  - `docs/ai/architecture/requirement-workflow-engine.md`
- Files:
  - `docs/ai/requirements/workflow_engine.yaml`
  - `docs/ai/requirements/template_discovery.yaml`
  - `docs/ai/requirements/state_machine.yaml`
  - `docs/ai/skills/requirement-workflow-engine/SKILL.md`
  - `docs/ai/templates/requirement-workflow-engine-sop.md`
  - `docs/ai/templates/project-requirement-discovery-record.yaml`
  - `docs/ai/knowledge/entries/kb_011_requirement_workflow_engine.md`
  - `docs/ai/skills/index.md`
  - `docs/ai/knowledge/index.md`
  - `docs/ai/architecture/index.md`
  - `docs/ai/runtime-rule-index.md`
  - `docs/ai/schemas/core.schema.yaml`
  - `docs/ai/tasks/slices/slice_010_requirement_workflow_engine.yaml`
  - `docs/ai/tasks/current-slice.yaml`
  - `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
  - `docs/ai/tasks/main-tree.yaml`
  - `docs/ai/tasks/forest.yaml`
  - `docs/ai/requirements/ledger.yaml`
  - `docs/ai/authorizations/records.md`
  - `docs/ai/decisions/records.md`
  - `docs/ai/status/current.yaml`
  - `docs/ai/manifest.yaml`
  - `docs/ai/handoff/current.md`
  - `docs/ai/changelog.md`
  - `docs/ai/scripts/validate_runtime.py`
- Validation:
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `python -m py_compile docs/ai/scripts/validate_runtime.py` -> passed.
  - `rg -n "requirement_workflow_engine|RequirementDiscoveryWorkflowEngine|REQUIREMENTS_GATHERING|template_discovery|skill_requirement_workflow_engine|kb_011" docs/ai` -> matched the runtime rule index, manifest, status, requirements, task tree, current slice, slice file, analysis, architecture, skill index, project skill, SOP, state machine, template catalog, engine definition, schema, knowledge index, knowledge entry, evidence, handoff, and validator.
  - `Test-Path prompt.md, package.json, package-lock.json, pnpm-lock.yaml, yarn.lock, pyproject.toml, requirements.txt, src, app, components, tools` -> all returned `False`.
- Summary: Strategy C is represented as a complete declarative requirement
  workflow engine with project-type probes, a requirement gathering state lock,
  question loops, follow-up triggers, and validator coverage while avoiding
  dependency installation, package manifests, application skeletons, and
  business code.

## evidence_024: Extension Runtime Removal

- Date: `2026-05-30T12:21:38+08:00`
- Type: `engineering_removal`
- Status: `completed`
- Slice: `slice_011_go_scaffold_engineering_repair`
- Trigger: Developer explicitly requested removing the current extension system
  and related services.
- Files:
  - Deleted source package and hook registry under the old runtime directory.
  - Deleted remote blog sample service directory.
  - Removed config schema, env examples, CI step, HTTP registration route,
    service lifecycle wiring, reload coupling, and stale tests.
  - Updated human docs and runtime handoff/current slice.
- Validation:
  - `rg ... -g '!docs/ai/**'` for old runtime identifiers -> no matches.
  - `go test ./internal/config ./internal/app/initapp ./internal/app/lifecycleapp ./internal/app/reloadapp ./internal/app ./internal/transport/http ./pkg/iam/... -count=1` -> passed.
  - `go test ./... -count=1 -mod=readonly` -> passed.
  - `go build -mod=readonly -o ./tmp/go-scaffold-server.exe ./cmd/main` -> passed.
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
- Summary: The current project no longer contains the old extension runtime,
  remote sample service, registration endpoint, config/env/CI surface, or docs
  page. Future extension work should start from fresh requirements.

## evidence_025: 用户管理栈移除

- Date: `2026-05-30T13:12:40+08:00`
- Type: `engineering_removal`
- Status: `completed`
- Slice: `slice_012_user_management_stack_removal`
- Trigger: 开发者明确要求删除 IAM、auth、RBAC 和相关用户管理服务。
- Files:
  - 删除 `internal/modules/user`、`pkg/auth`、`pkg/iam`、`pkg/rbac` 和 `configs/rbac_model.conf`。
  - 移除 auth/IAM/RBAC 配置结构、校验、环境变量覆盖测试、应用装配、reload 接线、用户 schema helper、路由注册、集成测试、配置块、部署环境变量和部署脚本认证密钥要求。
  - 通过 `go mod tidy` 清理 `go.mod` 和 `go.sum`。
  - 更新面向人的文档，说明当前脚手架不再内置用户管理栈；后续身份或访问控制能力需要从新需求开始。
- Validation:
  - `gofmt -w ./cmd ./internal ./pkg ./types` -> passed.
  - `go test ./internal/config ./internal/app/initapp ./internal/app/reloadapp ./internal/app ./internal/transport/http -count=1` -> passed.
  - `go test ./cmd/main ./internal/app/dbapp ./internal/modules/demo/... -count=1` -> passed.
  - `go test ./... -count=1 -mod=readonly` -> passed.
  - `go build -mod=readonly -o ./tmp/go-scaffold-server.exe ./cmd/main` -> passed.
  - `rg -n "github.com/rei0721/go-scaffold/(internal/modules/user|pkg/auth|pkg/rbac|pkg/iam)|RIN_APP_(AUTH|IAM|RBAC)|rbac_model|\\b(AuthConfig|RBACConfig|IAMConfig|AppAuthName|AppRBACName|AppIAMName|NewIAM|ApplyUserSchema|UserSchemaSQL|NewUserModule|ApplyConfiguredRBAC|IsIAMConfigChanged|reloadIAM|UserModule)\\b" --glob "!docs/ai/**" . -S` -> no matches.
- Summary: Go 脚手架不再包含旧的本地用户管理栈、路由面、配置面、schema helper 或直接 Casbin/JWT 依赖。后续身份认证或访问控制能力应从新的需求基线重新设计。

## evidence_026: 中文交付规则和文档同步

- Date: `2026-05-30T13:33:57+08:00`
- Type: `runtime_rule_and_documentation_sync`
- Status: `completed`
- Requirement: `req_infra_013_chinese_delivery_rule`
- Slice: `slice_013_chinese_delivery_rule`
- Trigger: 开发者要求新增以后中文交付的规则，并将上一轮修改内容调整为代码中文注释、文档中文概述。
- Files:
  - `docs/ai/runtime-rule-index.md` 新增 `INV-22` 中文交付规则。
  - README 和上一轮移除栈相关工程文档改为中文概述。
  - `.env.example`、配置示例、部署帮助文本清理英文说明和旧认证口径。
  - `internal/config/app_cors.go` 的修改注释保持中文，并移除默认 Authorization header 说明。
  - 同步 `docs/ai` 的 requirement、current slice、task tree、status、evidence 和 handoff。
- Validation:
  - `gofmt -w internal/config/app_cors.go internal/config/manager_test.go internal/app/app_integration_test.go` -> passed.
  - `rg -n "\\b(The|This|Current|Future|Removed|Deployment|Configuration|Testing|Known|Project|Directory|Startup|Layered|Adding|Usage:|Security note|Application options|Deployment options|Registry options|Local default|Replace with|user management stack|built-in)\\b" README.md docs .env.example configs deploy deploy.sh --glob "!docs/ai/**" -S` -> no matches.
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `go test ./... -count=1 -mod=readonly` -> passed.
  - `go build -mod=readonly -o .\\tmp\\go-scaffold-server.exe .\\cmd\\main` -> passed.
  - `git diff --check` -> passed.
  - Removed-stack residual `rg` outside `docs/ai` -> no matches.
  - Added/changed Go comment scan -> only Chinese CORS example comment found.
- Summary: 后续面向人的交付默认使用中文；新增或修改代码中的解释性注释使用中文；文档以中文概述当前事实，技术标识符和外部专有名词可保留原文。

## evidence_027: modeapp 重命名为 mainapp

- Date: `2026-05-30T14:15:15+08:00`
- Type: `internal_package_rename`
- Status: `completed`
- Requirement: `req_infra_014_modeapp_to_mainapp`
- Slice: `slice_014_modeapp_to_mainapp`
- Trigger: 开发者要求将 `modeapp` 改成 `mainapp`。
- Files:
  - 将 `internal/app/modeapp` 移动为 `internal/app/mainapp`。
  - 将 Go package 声明从 `modeapp` 改为 `mainapp`。
  - 更新 `internal/app/app.go` 和 `internal/app/constants.go` 的导入与调用点。
  - 更新 `docs/runtime/startup-flow.md` 中的应用子包说明。
  - 同步 `docs/ai` 的 requirement、current slice、task tree、status、evidence 和 handoff。
- Validation:
  - `gofmt -w internal/app/mainapp/mode.go internal/app/app.go internal/app/constants.go` -> passed.
  - `rg -n "modeapp|ModeApp|mode app|mode_app|internal/app/modeapp" . -S --glob "!docs/ai/**"` -> no matches.
  - `python docs/ai/scripts/validate_runtime.py` -> `Runtime validation passed.`
  - `go test ./... -count=1 -mod=readonly` -> passed.
  - `go build -mod=readonly -o .\\tmp\\go-scaffold-server.exe .\\cmd\\main` -> passed.
- Summary: 该改动只重命名内部应用子包，不改变 CLI mode 值或服务运行行为。
