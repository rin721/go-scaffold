# Requirement Workflow Engine SOP

This SOP is the canonical procedure for strategy C: a complete declarative
workflow engine for requirement discovery. It extends project requirement
discovery without becoming an executable framework.

## Core Principle

During `REQUIREMENTS_GATHERING`, the agent is a product-manager guide and
runtime recorder. It may ask questions and update requirement artifacts. It may
not generate business code, package manifests, dependency installs, source
directories, or application skeletons.

## 1. Engine Entry

Entry conditions:

- A developer provides a compact project idea.
- A future lifecycle gate reaches idea intake or requirement analysis.
- A developer asks for product-manager style requirement elicitation.

Required action:

1. Read `docs/ai/requirements/workflow_engine.yaml`.
2. Read `docs/ai/requirements/template_discovery.yaml`.
3. Read `docs/ai/requirements/state_machine.yaml`.
4. Confirm that the active state for raw idea intake is
   `REQUIREMENTS_GATHERING` and `code_generation_allowed: false`.

## 2. Intent Capture And Record Creation

Required output:

- Raw developer idea.
- Normalized intent.
- Selected task line.
- Selected probe template.
- Initial requirement discovery record path.

Default record path:

```text
docs/ai/requirements/intake/round_<nnn>_<feature_slug>.yaml
```

The record is created from
`docs/ai/templates/project-requirement-discovery-record.yaml` and then extended
with `probe_template_id`, `question_rounds`, `coverage_matrix`, and
`followup_triggers`.

## 3. Probe Template Selection

Select from `template_discovery.yaml`:

- `commerce` for shop, store, order, cart, SKU, marketplace, or mall ideas.
- `saas` for tenants, subscriptions, workspaces, dashboards, or B2B products.
- `blog` for publishing, CMS, article, post, or documentation sites.
- `generic_app` when no stronger project class is found.

Do not invent a new project class inside an active round unless the developer
confirms that the existing templates are insufficient.

## 4. Question Batch Loop

Each batch must:

- Ask 5-7 numbered questions.
- Tie each question to a requirement domain.
- Include options when options clarify the decision.
- Preserve a free-form answer path.
- Avoid implementation details before product shape is clear.

After each developer answer:

1. Overwrite the intake YAML with the answer.
2. Move answered items to confirmed, deferred, out of scope, rejected,
   research-needed, or decision-needed.
3. Apply template follow-up rules.
4. Add finer questions when the answer opens a domain.
5. Sync stable confirmed facts to the requirement ledger.
6. Update evidence and handoff when durable state changes.

## 5. Follow-Up Rule Examples

Commerce SKU answer:

- Ask SKU dimensions.
- Ask SKU combination rules.
- Ask stock deduction timing.
- Ask whether order lines store product/SKU snapshots.

Real payment answer:

- Ask provider and sandbox expectations.
- Ask callback, secret, refund, and audit needs.
- Mark provider selection as current-source research-needed.

SaaS multi-tenant answer:

- Ask tenant isolation model.
- Ask invitations and roles.
- Ask support access and audit rules.
- Ask whether billing is simulated, deferred, or real.

## 6. Coverage And Baseline Readiness

The engine may request requirement-baseline confirmation only when:

- Critical domains are confirmed, deferred, out of scope, or blocked by a named
  decision.
- AI-inferred candidates remain clearly marked as candidates.
- Open questions are non-blocking or explicitly deferred.
- Proposed first-round scope and non-goals are visible.
- Acceptance criteria draft exists.
- Research-needed items are named and routed downstream.

If the gate fails, ask the next smallest useful batch instead of requesting
approval.

## 7. Exit To Downstream Lifecycle

After hard baseline confirmation:

1. Initialize `docs/ai/lifecycle/<round_id>-downstream.yaml`.
2. Run the downstream lifecycle detailing workflow.
3. Keep implementation locked until research, task analysis, architecture, and
   mode confirmations authorize an executable slice.
