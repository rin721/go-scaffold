# Risk Confirmation Template

## Risk Summary

| Risk ID | Title | Type | Severity | Probability | Blocking | Status |
|---|---|---|---|---|---|---|
| RISK-001 | Scope expansion into full rewrite | Scope | High | Medium | Yes | [RISK] |
| RISK-002 | Documentation drift | Documentation | Medium | High | Yes | [RISK] |
| RISK-003 | Test coverage gaps | Testing | High | Medium | No | [RISK] |
| RISK-004 | Migration boundary ambiguity | Architecture | High | Medium | Yes | [RISK] |
| RISK-005 | Public package compatibility ambiguity | Architecture | Medium | Medium | Yes | [RISK] |
| RISK-006 | Agent execution ambiguity | Process | High | High | Yes | [RISK] |

## Risk Details

### RISK-001: Scope Expansion Into Full Rewrite

- Impact: Optimization work may become unbounded and difficult to verify.
- Trigger: Starting code refactors before requirements and architecture confirmation.
- Mitigation: Use Option A by default and require user confirmation for Option B/C/D.
- Owner: User and Agent.
- Required Confirmation: [NEEDS_CONFIRMATION] Confirm optimization route.

### RISK-002: Documentation Drift

- Impact: Future agents and developers may rely on stale or conflicting facts.
- Trigger: README, `.env.example`, package README files, and code disagree.
- Mitigation: Promote confirmed facts into root project documents and keep docs updated per time slice.
- Owner: Agent.
- Required Confirmation: [NEEDS_CONFIRMATION] Confirm whether JWT/auth examples should be deferred or kept as placeholders.

### RISK-003: Test Coverage Gaps

- Impact: Code optimization could regress app startup, routing, demo, or reload behavior without detection.
- Trigger: Making code changes before integration tests exist.
- Mitigation: Add test strategy during requirements and architecture phases.
- Owner: Agent.
- Required Confirmation: [NEEDS_CONFIRMATION] Confirm minimum test matrix for P0 optimizations.

### RISK-004: Migration Boundary Ambiguity

- Impact: Dev and production database initialization paths may conflict.
- Trigger: `AutoMigrate`, `initdb`, and SQL scripts being used without a single policy.
- Mitigation: Confirm migration policy before database-related changes.
- Owner: User.
- Required Confirmation: [NEEDS_CONFIRMATION] Choose migration strategy.

### RISK-005: Public Package Compatibility Ambiguity

- Impact: Changes to `pkg/*` may accidentally break external or future consumers.
- Trigger: Treating packages as both internal implementation details and reusable libraries.
- Mitigation: Confirm API policy and document compatibility expectations.
- Owner: User.
- Required Confirmation: [NEEDS_CONFIRMATION] Decide public/internal/mixed policy.

### RISK-006: Agent Execution Ambiguity

- Impact: Future "next step" requests may be handled from chat memory rather than project state.
- Trigger: Missing `TASKS.md`, `TIME_SLICES.md`, `AGENT_HANDOFF.md`, and related state files.
- Mitigation: Build full agent infrastructure after requirements and architecture confirmation.
- Owner: Agent.
- Required Confirmation: [NEEDS_CONFIRMATION] Confirm startup documents, then generate formal agent state files.

## Confirmation Checklist

- [ ] Confirm optimization option.
- [ ] Confirm `pkg/*` API policy.
- [ ] Confirm demo module role.
- [ ] Confirm migration policy.
- [ ] Confirm auth/JWT scope.
- [ ] Confirm whether to proceed to requirements generation.

