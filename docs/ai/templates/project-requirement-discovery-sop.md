# Project Requirement Discovery SOP

This SOP is the canonical operating procedure for turning a raw developer idea
into a durable requirement baseline candidate. It is a sub-procedure of the
full-project lifecycle and must run before requirement-baseline confirmation.

## Core Principle

A developer idea is not a requirement baseline.

An agent may summarize the idea to show understanding, but it must not collapse
the flow into a generic MVP proposal. The agent must discover the project shape,
persist what is known, mark what is inferred, and guide the developer through
the missing decisions.

When the declarative requirement workflow engine is present, it is the
controlling surface for this SOP:

- Engine: `docs/ai/requirements/workflow_engine.yaml`
- Probe templates: `docs/ai/requirements/template_discovery.yaml`
- State lock: `docs/ai/requirements/state_machine.yaml`

`REQUIREMENTS_GATHERING` forbids business code generation, package manifests,
dependency installation, and application skeleton creation.

## 1. `idea_seed_intake_gate`

Capture the raw seed.

Required output:

- Raw developer text.
- Normalized one-sentence intent.
- Project type guess.
- Known technology or operational constraints.
- Stated exclusions.
- Unknowns that affect project scope.

Rule:

- The raw idea may mention a stack or shape, such as Go + Nuxt and separated
  startup, but that does not confirm the whole stack, architecture, data model,
  feature scope, or delivery plan.

## 2. `idea_interpretation_gate`

Translate the idea into an interpretation that remains reversible.

Required output:

- Product intent: what the developer appears to want to build.
- System category: the likely system class.
- Implied requirements: AI-inferred candidates, marked as unconfirmed.
- Missing decisions: what cannot be safely assumed.
- Risk flags: money, auth, privacy, production, destructive actions, external
  services, compliance, supply chain, or high user impact.

Hard stop:

- Do not present inferred details as final requirements.

## 3. `domain_surface_mapping_gate`

Map the requirement domains that must be explored for this project class.

For an application project, scan at least:

| Domain | Example questions |
| --- | --- |
| Goal and users | Who uses it, what job should it help them complete, and what outcome counts as success? |
| Roles and permissions | Are there buyers, admins, operators, merchants, guests, or staff? |
| Core workflows | What are the end-to-end flows from entry to completion? |
| Data model | What entities exist, who owns them, and what lifecycle do they follow? |
| Integrations | Are payment, messaging, maps, AI APIs, storage, analytics, or third-party systems needed? |
| Admin and operations | Who configures, moderates, imports, exports, audits, or recovers data? |
| Security and privacy | What auth, authorization, sensitive data, and audit rules are required? |
| Non-functional quality | Performance, accessibility, reliability, observability, backups, and error handling. |
| Development topology | Local startup, environment variables, ports, scripts, test data, and developer workflow. |
| Testing and acceptance | What must be tested, demoed, or inspected before acceptance? |
| Documentation | What docs are needed for users, developers, operators, and future agents? |
| Non-goals | What is explicitly out of scope for this round? |

## 4. `requirement_inventory_gate`

Create a requirement inventory with statuses.

Allowed statuses:

- `developer_stated`
- `ai_inferred_candidate`
- `pending_confirmation`
- `research_needed`
- `decision_needed`
- `confirmed`
- `deferred`
- `out_of_scope`
- `rejected`

Required categories:

- Functional requirements.
- Non-functional requirements.
- Data requirements.
- User or operator workflows.
- Constraints.
- Non-goals.
- Risks.
- Open questions.

## 5. `question_backlog_gate`

Build a prioritized question backlog.

Question rules:

- Ask 3 to 7 high-leverage questions in the first batch.
- When a matching project template exists, ask 5 to 7 numbered questions from
  that template unless the developer requests a smaller batch.
- Tie each question to a requirement domain and a blocking decision.
- Provide 2 to 4 concrete options when options help the developer decide.
- Include a free-form path when the developer may have a different intent.
- Avoid asking about implementation details before product shape is clear.

The backlog must preserve:

- Question ID.
- Domain.
- Why the answer matters.
- Suggested options, if useful.
- Status.
- Developer answer, when provided.
- Ledger sync target, when confirmed.

## 6. `requirement_collection_round_gate`

After each developer response:

1. Update the discovery record.
2. Move answered items from `pending_confirmation` to `confirmed`, `deferred`,
   `out_of_scope`, `rejected`, or `decision_needed`.
3. Detect contradictions with prior answers.
4. Apply follow-up trigger rules from `template_discovery.yaml`.
5. Add new questions only when the answer opens a new requirement domain or
   increases the required detail inside an already-open domain.
6. Synchronize stable facts to the requirement ledger.
7. Ask the next smallest useful batch of questions.

This gate may repeat across turns.

## 7. `requirement_plan_gate`

Plan the remaining requirement work before research and architecture.

Required output:

- Requirement domains that are ready for baseline confirmation.
- Requirement domains that still need developer answers.
- Items requiring current-source research.
- Decisions deferred to later rounds.
- Proposed first-round scope and explicit non-goals.
- Acceptance criteria draft.
- Evidence plan for proving requirements later.

## 8. `requirement_persistence_sync_gate`

Persist requirement discovery state.

Default storage:

```text
docs/ai/requirements/intake/<round_id>.yaml
```

Use `docs/ai/templates/project-requirement-discovery-record.yaml` as the shape.

Sync rules:

- Confirmed requirements go to `docs/ai/requirements/ledger.yaml`.
- Unconfirmed AI inferences stay in the discovery record.
- Open questions stay in the discovery record and handoff.
- Decisions go to `docs/ai/decisions/records.md` only after confirmation.
- Evidence links must point to the artifact or command that proves the update.

## 9. `baseline_readiness_gate`

Requirement baseline confirmation is allowed only when:

- Critical domains are confirmed, explicitly deferred, out of scope, or marked
  as blocked by a named decision.
- AI-inferred candidates are not mixed with confirmed requirements.
- Acceptance criteria exist for the proposed first round.
- Non-goals are explicit.
- Research-needed items are named and routed to the research gate.
- The developer can see what will and will not be built in this round.

Hard stop:

- If the readiness gate fails, keep collecting requirements instead of asking
  for broad project approval.

## Commerce Seed Response Pattern

For a future seed like "I want a Go + Nuxt, frontend/backend separated,
separately started commerce system", the first response should be shaped like:

1. Route the work: this is a future mainline project idea, while current Vibe
   Coding infrastructure state may still control whether mainline work can
   start immediately.
2. Interpret the idea: separated frontend/backend commerce system, likely
   local development first, but business model, roles, checkout, admin, data,
   and risk boundaries are not yet confirmed.
3. State the discovery domains: roles, catalog, cart, checkout, orders, admin,
   inventory, auth, payment, fulfillment, data, local startup, quality,
   security, deployment, docs, and non-goals.
4. Persist candidate requirements as unconfirmed.
5. Ask the first requirement batch, for example:
   - Is the first round for learning/demo, internal tool, or real commercial
     use?
   - Who are the roles in round one: guest buyer, registered buyer, admin,
     merchant, operator?
   - Should checkout create unpaid demo orders, or should any payment provider
     be explored later?
   - Should inventory be real stock control, simple product availability, or
     out of scope?
   - What admin operations are required in round one?
   - What local database and data persistence expectations are acceptable for
     development?
6. Continue collecting answers before requirement-baseline confirmation.

The agent may offer a recommended first-round scope, but it must label it as a
candidate and ask the developer to confirm, edit, or reject it.
