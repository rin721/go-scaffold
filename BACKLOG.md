# BACKLOG.md

## Backlog Status

- Project: go-scaffold
- Last Updated: 2026-05-25
- Rule: Items here are not current scope unless promoted by user confirmation and task decomposition.

## Backlog Items

| ID | Title | Source | Priority | Reason Deferred | Status |
|---|---|---|---|---|---|
| BL-001 | Reconcile JWT/auth examples with README scope | `.env.example`, README | P1 | Requires user decision on auth/JWT scope | [DEFERRED] |
| BL-002 | Add app/router/demo integration tests | Project risk analysis | P1 | Requires architecture/test strategy confirmation | [DEFERRED] |
| BL-003 | Define package API compatibility policy | `pkg/*` public/internal ambiguity | P1 | Requires architecture decision | [DEFERRED] |
| BL-004 | Define migration policy across AutoMigrate, initdb, and SQL scripts | Migration boundary risk | P1 | Requires architecture decision | [DEFERRED] |
| BL-005 | Mark unsupported sqlgen capabilities explicitly | `pkg/sqlgen` TODO/unimplemented items | P1 | Requires package API policy | [DEFERRED] |
| BL-006 | Add CI quality gate | Future quality work | P2 | Out of current requirements phase | [DEFERRED] |
| BL-007 | Add deployment guidance | Future release work | P2 | Out of current requirements phase | [DEFERRED] |
| BL-008 | Add auth/rbac module | User-facing feature | P2 | README currently says auth/rbac is not implemented | [DEFERRED] |
| BL-009 | Add scaffold generator | Productization idea | P2 | Requires architecture and module policy | [DEFERRED] |
| BL-010 | Add performance benchmark suite | Quality/performance idea | P2 | Requires stable functional boundaries | [DEFERRED] |
| BL-011 | Add multi-tenancy support | Product architecture idea | P2 | Not confirmed and likely large scope | [DEFERRED] |
| BL-012 | Add rpc plugin adapter | User plugin request | P2 | v1 only implements local and HTTP | [DEFERRED] |
| BL-013 | Add ws plugin adapter | User plugin request | P2 | v1 only implements local and HTTP | [DEFERRED] |
| BL-014 | Add plugin discovery from manifests or registries | Future plugin runtime enhancement | P2 | Requires explicit promotion after v1 API acceptance | [DEFERRED] |
| BL-015 | Add local/http plugin examples | Plugin v1 closeout | P1 | Useful but not required for core v1 acceptance | [DEFERRED] |

## Promotion Rules

- [CONFIRMED] A backlog item can be promoted only after user confirmation.
- [CONFIRMED] Promoted items must be mapped to requirements, architecture, tasks, and time slices.
- [CONFIRMED] Backlog items must not be implemented opportunistically.
