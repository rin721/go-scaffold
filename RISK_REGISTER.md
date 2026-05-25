# RISK_REGISTER.md

## Risk Register Status

- Project: go-scaffold
- Phase: Requirements Confirmation
- Status: PENDING_USER_CONFIRMATION
- Last Updated: 2026-05-25

## Risk Format

- Risk ID:
- Title:
- Type:
- Severity:
- Probability:
- Impact:
- Trigger:
- Mitigation:
- Owner:
- Status:
- Blocking:

## Risks

### RISK-001: Scope Expansion Into Full Rewrite

- Type: Scope
- Severity: High
- Probability: Medium
- Impact: Optimization work may become unbounded, hard to test, and hard to complete.
- Trigger: Starting refactors before requirements, architecture, tasks, and time slices are confirmed.
- Mitigation: Use conservative governance-first route by default; keep new ideas in `BACKLOG.md`.
- Owner: User/Agent
- Status: [RISK]
- Blocking: Yes, blocks implementation.

### RISK-002: Documentation Drift

- Type: Documentation
- Severity: Medium
- Probability: High
- Impact: Future agents and developers may rely on stale or conflicting information.
- Trigger: README, `.env.example`, package README files, and implementation details disagree.
- Mitigation: Keep root project docs authoritative and update them after each phase.
- Owner: Agent
- Status: [RISK]
- Blocking: Yes, blocks architecture confirmation until major drift points are captured.

### RISK-003: Test Coverage Gaps

- Type: Testing
- Severity: High
- Probability: Medium
- Impact: Future code optimizations may regress startup, router, demo CRUD, config reload, or migration behavior.
- Trigger: Making code changes with only partial package-level tests.
- Mitigation: Define test matrix in architecture phase before implementation.
- Owner: Agent
- Status: [RISK]
- Blocking: No, but must be addressed before code optimization.

### RISK-004: Migration Boundary Ambiguity

- Type: Architecture
- Severity: High
- Probability: Medium
- Impact: Development and production database initialization paths may conflict.
- Trigger: `AutoMigrate`, `initdb`, and SQL scripts existing without a confirmed policy.
- Mitigation: Confirm migration strategy before database-related code changes.
- Owner: User
- Status: [RISK]
- Blocking: Yes, blocks database architecture decisions.

### RISK-005: Public Package Compatibility Ambiguity

- Type: Architecture
- Severity: Medium
- Probability: Medium
- Impact: Changes to `pkg/*` may accidentally break consumers if they are treated as stable public APIs.
- Trigger: Refactoring `pkg/*` without deciding public/internal/mixed policy.
- Mitigation: Record package API policy in future architecture and decisions documents.
- Owner: User
- Status: [RISK]
- Blocking: Yes, blocks package-level refactors.

### RISK-006: Agent Execution Ambiguity

- Type: Process
- Severity: High
- Probability: High
- Impact: Future `下一步` requests may become ambiguous without full task and time-slice infrastructure.
- Trigger: Missing `TASKS.md`, `TIME_SLICES.md`, `AGENT_HANDOFF.md`, `AGENT_RULES.md`, and related files.
- Mitigation: Generate full agent infrastructure after requirements and architecture confirmation.
- Owner: Agent
- Status: [RISK]
- Blocking: No for requirements, yes before implementation.

### RISK-007: Auth/JWT Scope Drift

- Type: Scope/Security
- Severity: Medium
- Probability: Medium
- Impact: Auth-related examples may imply unsupported security behavior.
- Trigger: `.env.example` includes JWT settings while README says auth/rbac is out of scope.
- Mitigation: Keep auth/rbac deferred or explicitly promote it after user confirmation.
- Owner: User/Agent
- Status: [RISK]
- Blocking: No for requirements, yes before auth-related documentation or implementation changes.

### RISK-008: Plugin System Scope Expansion

- Type: Scope/Architecture
- Severity: Medium
- Probability: Medium
- Impact: Plugin runtime could grow into application composition, service discovery, or dynamic code loading before boundaries are confirmed.
- Trigger: Adding rpc/ws/discovery/app integration before v1 API review.
- Mitigation: Keep v1 limited to independent `pkg/plugin` with local and HTTP adapters.
- Owner: Agent/User
- Status: [RISK]
- Blocking: No for v1; yes for further protocol expansion until a candidate is promoted.

## Required User Decisions

| ID | Decision | Blocking What | Status |
|---|---|---|---|
| RD-001 | Confirm Option A or choose Option B/C | Architecture and task decomposition | [NEEDS_CONFIRMATION] |
| RD-002 | Decide `pkg/*` API policy | Package architecture | [NEEDS_CONFIRMATION] |
| RD-003 | Decide demo module role | Module architecture | [NEEDS_CONFIRMATION] |
| RD-004 | Decide migration strategy | Database architecture | [NEEDS_CONFIRMATION] |
| RD-005 | Decide auth/JWT scope | Documentation and backlog promotion | [NEEDS_CONFIRMATION] |
| RD-006 | Select next plugin/system direction after `pkg/plugin` v1 | Further plugin protocols, discovery, examples, or governance route | [NEEDS_CONFIRMATION] |
