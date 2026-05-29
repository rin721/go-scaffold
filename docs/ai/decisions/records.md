# Decision Records

## decision_001: Bootstrap Runtime Infrastructure

- Date: `2026-05-29T17:18:17+08:00`
- Status: `accepted_for_bootstrap`
- Source: developer compiler run request
- Decision: Create lightweight Vibe Coding runtime infrastructure in `AGENTS.md`
  and `docs/ai/*`.
- Rationale: The repository had no existing runtime entries, rules,
  dependency files, tests, CI, or business files. The bootstrap can safely stay
  within documentation and AI runtime infrastructure.
- Alternatives:
  - Do nothing: rejected because runtime state would remain only in chat.
  - Build business code immediately: rejected because core governance comes
    before business execution.
- Confirmation level: `auto_recordable` for bootstrap files only.
- Evidence: `evidence_001`, `evidence_010`
- No prompt runtime dependency: true

## decision_002: Human Documentation Slice

- Date: `2026-05-29T17:34:47+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer goal, "当前项目的食用（使用）教程等相关详细概述文件"
- Decision: Create root and `docs/` human-facing documentation for project
  overview, usage, maintenance, and Agent workflow.
- Rationale: The repository had runtime infrastructure but no human tutorial or
  overview. Documentation is low risk and improves recoverability.
- Alternatives:
  - Ask for more detail before writing: not needed because the requested scope
    is clear and low risk.
  - Create only README: rejected because the user asked for detailed overview
    files.
- Confirmation level: `auto_recordable` within documentation and docs/ai
  status/evidence files.
- Evidence: `evidence_012`
- No prompt runtime dependency: true

## decision_003: Python Runtime Validator

- Date: `2026-05-29T17:43:05+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request, "将项目中脚本，替换成py"
- Decision: Replace the runtime validation script with
  `docs/ai/scripts/validate_runtime.py` and update command references.
- Rationale: Python is available in the local environment and is more portable
  across Windows, macOS, and Linux for this repository's lightweight validator.
- Alternatives:
  - Keep the PowerShell validator: rejected by developer request.
  - Add both PowerShell and Python validators: rejected because the request was
    to replace scripts with Python.
- Confirmation level: `auto_recordable` for low-risk runtime script
  maintenance.
- Evidence: `evidence_014`
- No prompt runtime dependency: true

## decision_004: Task Forest With Vibe Coding Infrastructure Branch

- Date: `2026-05-29T17:54:03+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request to analyze, sandbox, and implement a separate Vibe
  Coding task tree branch.
- Decision: Introduce `docs/ai/tasks/forest.yaml`, reserve
  `docs/ai/tasks/main-tree.yaml` for future product or business ideas, and route
  Vibe Coding repository infrastructure work to
  `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`.
- Rationale: Documentation, runtime validation scripts, schemas, state, evidence,
  knowledge, skills, and handoff changes iterate this repository's Vibe Coding
  infrastructure. They should not pollute the future mainline idea task tree.
- Alternatives:
  - Keep one task tree: rejected because it mixes infrastructure maintenance
    with product idea work.
  - Create a new mainline round for every infrastructure edit: rejected because
    these edits are not product idea rounds.
- Confirmation level: `auto_recordable` because the developer explicitly asked
  for this governance capability and the change is documentation/runtime-state
  scoped.
- Evidence: `evidence_015`
- No prompt runtime dependency: true

## decision_005: Catalog-Only Capability Groups

- Date: `2026-05-29T18:53:11+08:00`
- Status: `accepted_for_confirmed_A_execution`
- Source: developer confirmation, "确认 A，进入阶段二，只建立能力目录、skills、知识库、依赖候选清单，不创建工程骨架，不安装依赖。"
- Decision: Create `docs/ai/capabilities/`, grouped capability candidate
  files, a capability-selection skill, a knowledge entry, research notes, and
  architecture strategy notes as catalog-only infrastructure.
- Rationale: The developer confirmed the conservative A route: preserve
  future choice surfaces without committing the repository to a UI library,
  framework, Python toolchain, package manifest, lockfile, or application
  skeleton.
- Alternatives:
  - Create an engineering skeleton now: rejected by explicit developer
    constraint.
  - Install dependencies now: rejected by explicit developer constraint.
  - Keep only a chat recommendation: rejected because runtime recovery needs
    physical docs, skills, knowledge, and evidence.
- Confirmation level: `hard_confirmation` for the A strategy boundary,
  `auto_recordable` for low-risk docs/ai runtime artifacts inside that boundary.
- Evidence: `evidence_016`
- No prompt runtime dependency: true

## decision_006: Dynamic Current-Year Capability Refresh

- Date: `2026-05-29T19:07:53+08:00`
- Status: `accepted_for_confirmed_A_refresh`
- Source: developer confirmation, "绿灯 A：进入阶段二，只更新 catalog-only 基建产物，不安装依赖、不建工程骨架，应当“请先联网核验 {当前年份，不能写死} 年当前官方资料”"
- Decision: Refresh the catalog-only capability groups against current official
  sources and add a dynamic current-year rule: future research wording must
  derive the current year from runtime/source-refresh time instead of hard-coding
  a calendar year.
- Rationale: Capability choices change over time, but this repository should not
  bake stale date wording into reusable skills or prompts. The catalog remains a
  developer choice surface and must stay separate from installation or
  scaffolding authorization.
- Alternatives:
  - Hard-code the current calendar year: rejected because the developer
    explicitly requested a non-hard-coded current-year policy.
  - Install refreshed candidate dependencies: rejected because strategy A only
    allows catalog-only infrastructure.
  - Defer all source refresh to a future mainline slice: rejected because the
    developer explicitly green-lit phase-two catalog updates now.
- Confirmation level: `hard_confirmation` for the strategy A refresh boundary,
  `auto_recordable` for docs/ai updates inside that boundary.
- Evidence: `evidence_017`
- No prompt runtime dependency: true

## decision_007: Trigger-Bound Vibe Infrastructure Action Workflow SOP

- Date: `2026-05-29T19:32:05+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request to add an action workflow for
  `新增 Vibe Coding 仓库基建设施建立任务线：[具体动作]`.
- Decision: Add a project skill and canonical Markdown SOP template that bind
  the trigger payload to `{{ 优化动作 }}`, enforce phase-one four-dimensional
  diagnosis and red-light confirmation, and allow phase-two physical output only
  after explicit human authorization.
- Rationale: The workflow itself governs repository infrastructure changes and
  must be recoverable from physical runtime artifacts rather than chat memory.
  A project skill plus template gives future agents a precise trigger surface
  and a hard confirmation gate.
- Alternatives:
  - Keep the SOP only in chat: rejected because runtime recovery requires
    physical artifacts.
  - Put the whole rule only in `AGENTS.md`: rejected because project skills are
    the existing catalog for repeatable or fragile workflows.
  - Let the trigger authorize immediate writes: rejected because the requested
    SOP requires an absolute phase-one/phase-two separation.
- Confirmation level: `auto_recordable` for docs/ai runtime artifacts;
  `hard_confirmation` remains required by the SOP before any future phase-two
  physical output generated through that workflow.
- Evidence: `evidence_018`
- No prompt runtime dependency: true

## decision_008: Full-Project Lifecycle Workflow

- Date: `2026-05-29T21:04:30+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request to add a Vibe Coding repository infrastructure
  task line that explains and repairs why mainline ideas were only summarized
  into requirements instead of being driven through a complete project.
- Decision: Add a full-project lifecycle project skill, canonical SOP, gap
  analysis, architecture note, knowledge entry, slice, validator coverage, and
  runtime state updates. Future mainline or confirmed project task lines must
  route the task line, analyze and confirm requirements, perform current-source
  research and technology validation, confirm research, analyze and confirm
  tasks, design and confirm architecture and stack, recommend and confirm
  infrastructure mode, establish Agent Vibe Coding driving facilities, create
  task trees and slices, implement, test, document, update state/evidence, use
  the next-step protocol, close the task tree, and re-enter round `n+1`.
- Rationale: The existing state machine listed the phases but did not provide a
  detailed physical SOP that prevents skipped or over-summarized gates. A
  dedicated workflow makes future recovery and execution depend on local
  artifacts rather than chat memory.
- Alternatives:
  - Keep only the existing state machine: rejected because phase labels alone
    did not force detailed outputs, confirmations, or next-step behavior.
  - Start a mainline product implementation now: rejected because this task
    changes Vibe Coding infrastructure and no product idea has passed the
    lifecycle gates yet.
  - Treat the existing trigger-bound infrastructure SOP as enough: rejected
    because it governs two-stage infrastructure actions, not the complete
    mainline project lifecycle and round `n+1` loop.
- Confirmation level: `auto_recordable` for docs/ai runtime artifacts;
  `hard_confirmation` remains required for future requirement baselines,
  research/technology stack choices, task analysis, architecture, and mode
  selection before business implementation.
- Evidence: `evidence_019`
- No prompt runtime dependency: true

## decision_009: Compiler Runtime Assimilation Workflow

- Date: `2026-05-29T21:30:00+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request to deeply analyze, distill, and assign a Vibe
  Coding multi-round infrastructure compiler specification into this repository,
  with the clarification that the current project is already an infrastructure
  repository.
- Decision: Add a compiler-runtime assimilation project skill, canonical SOP,
  analysis, architecture note, knowledge entry, slice, validator coverage, and
  runtime state updates. Future compiler, generator, prompt-generation, or macro
  governance specifications must be routed to `branch_vibe_coding_infra` and
  translated into local runtime artifacts rather than copied into `prompt.md` or
  used as a daily recovery dependency.
- Rationale: The provided source is valuable as semantic input, but the runtime
  repository already has `AGENTS.md` and `docs/ai/*` artifacts. The reliable
  path is to strengthen those artifacts and preserve `INV-02` instead of
  reintroducing a large hidden prompt dependency.
- Alternatives:
  - Generate a new `prompt.md`: rejected because the current repository is
    already initialized and runtime recovery must not depend on raw compiler
    text.
  - Only keep the analysis in chat: rejected because future agents need
    physical skills, knowledge, state, evidence, and handoff entries.
  - Fold this entirely into the full-project lifecycle workflow: rejected
    because compiler/generator assimilation is a distinct infrastructure task
    line that updates repository governance rather than product lifecycle work.
- Confirmation level: `auto_recordable` for docs/ai runtime artifacts inside
  the explicit "assign to current repository" request; `hard_confirmation`
  remains required for any future change that expands agency, changes
  permissions, touches production, installs dependencies, or creates business
  implementation.
- Evidence: `evidence_020`
- No prompt runtime dependency: true

## decision_010: Project Requirement Discovery Workflow

- Date: `2026-05-29T21:48:21+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request to add a Vibe Coding infrastructure task line that
  refines the idea-to-requirement flow because the Go + Nuxt commerce example
  was still too summarized and did not fully guide complete project
  requirements.
- Decision: Add a project requirement discovery skill, canonical SOP, durable
  intake record template, gap analysis, architecture note, knowledge entry,
  slice, validator coverage, and runtime state updates. Future raw project
  ideas must be interpreted, mapped into requirement domains, inventoried into
  developer-stated and AI-inferred candidates, guided through question
  backlogs, persisted, planned, and checked for baseline readiness before
  requirement-baseline confirmation.
- Rationale: The existing full-project lifecycle names the early gates, but
  the repository still needed an engineering-grade sub-procedure for collecting
  requirements across turns. The new workflow prevents a raw idea from being
  prematurely converted into a generic MVP list or a single broad confirmation
  request.
- Alternatives:
  - Keep only the full-project lifecycle SOP: rejected because it did not
    define the detailed requirement discovery record, question backlog, and
    readiness checks.
  - Start the Go + Nuxt commerce project now: rejected because this request is
    an infrastructure workflow repair and no mainline requirement baseline has
    been confirmed.
  - Store requirement discovery only in chat: rejected because future agents
    need physical recovery artifacts and ledger links.
- Confirmation level: `auto_recordable` for docs/ai runtime artifacts inside
  the developer's infrastructure workflow request; `hard_confirmation` remains
  required for future mainline requirement baselines, stack choices,
  architecture, and implementation mode.
- Evidence: `evidence_021`
- No prompt runtime dependency: true

## decision_011: Project Lifecycle Downstream Detailing Workflow

- Date: `2026-05-29T22:00:59+08:00`
- Status: `accepted_for_low_risk_execution`
- Source: developer request to analyze the current Vibe Coding flow and make
  later stages follow the same detailed engineering process as the requirement
  discovery repair.
- Decision: Add a downstream lifecycle detailing skill, canonical SOP, durable
  downstream record template, lifecycle record guide, gap analysis,
  architecture note, knowledge entry, slice, validator coverage, and runtime
  state updates. Future project rounds after requirement-baseline confirmation
  must run detailed gates for research, task analysis, architecture, mode,
  Agent driving infrastructure, slice contracts, implementation ledger,
  verification evidence, closure, and next-round reentry.
- Rationale: The lifecycle and requirement discovery repairs prevent shallow
  intake, but post-baseline stages still needed a physical operating procedure
  and record shape. The new workflow prevents research, task planning,
  architecture, mode selection, execution, and closure from becoming brief
  narrative summaries.
- Alternatives:
  - Keep only the full-project lifecycle SOP: rejected because it named the
    stages but did not define downstream records, packets, and iteration
    ledgers.
  - Start a business project now: rejected because this request changes Vibe
    Coding infrastructure and no mainline requirement baseline has been
    confirmed.
  - Store downstream detail only in chat: rejected because future agents need
    recoverable physical artifacts and evidence links.
- Confirmation level: `auto_recordable` for docs/ai runtime artifacts inside
  the developer's infrastructure workflow request; `hard_confirmation` remains
  required for future mainline research conclusions, architecture, mode, and
  implementation choices.
- Evidence: `evidence_022`
- No prompt runtime dependency: true

## decision_012: Strategy C Requirement Workflow Engine

- Date: `2026-05-29T22:24:00+08:00`
- Status: `accepted_for_confirmed_C_execution`
- Source: developer confirmation, "确认方案 C，进入阶段二"
- Decision: Implement strategy C as a complete declarative requirement
  workflow engine under `docs/ai/requirements/`, with a workflow node graph,
  reusable probe templates, a `REQUIREMENTS_GATHERING` state lock, PM-style
  question loops, follow-up triggers, schema/validator coverage, task state,
  knowledge, evidence, and handoff updates.
- Rationale: The existing requirement discovery SOP is detailed, but strategy C
  needs a machine-readable control plane that future agents can inspect before
  asking questions or writing requirement records. Implementing it as
  declarative runtime infrastructure gives strong control without turning the
  repository into an executable framework or dependency-bearing project.
- Alternatives:
  - Strategy A, SOP-only update: rejected because the developer confirmed C.
  - Strategy B, template plus state lock only: rejected because the developer
    confirmed the stronger complete workflow engine.
  - Build an executable workflow framework: rejected because this repository
    should stay docs/ai runtime infrastructure and must not create package
    manifests, install dependencies, or scaffold application code.
- Confirmation level: `hard_confirmation` for strategy C phase-two execution;
  `auto_recordable` for docs/ai runtime artifacts inside the confirmed scope.
- Evidence: `evidence_023`
- No prompt runtime dependency: true
