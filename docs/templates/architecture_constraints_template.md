# Architecture Constraints Template

## Constraint Status

- Current Status: PENDING_USER_CONFIRMATION
- Source: Repository inspection and `README.md`
- Applies To: Future requirements, architecture, task decomposition, and code changes

## Dependency Direction

- [INFERRED] `cmd/server` should remain the process entry and CLI boundary only.
- [INFERRED] `internal/app` should remain the composition root and lifecycle owner.
- [INFERRED] `internal/transport/http` should own HTTP routing and health endpoints.
- [INFERRED] `internal/modules/*` should own business modules.
- [INFERRED] `pkg/*` should not depend on `internal/*`.

## Layering Rules

- [CONFIRMED] Handler should handle binding, HTTP status, and response conversion.
- [CONFIRMED] Service should handle business validation, transaction orchestration, and repository calls.
- [CONFIRMED] Repository should handle GORM data access and avoid business decisions.
- [NEEDS_CONFIRMATION] Decide whether these demo layering rules are mandatory for all future modules.

## Configuration Boundary

- [INFERRED] Config structures and validation should stay in `internal/config`.
- [INFERRED] Mapping from app config to package config should stay in `internal/app/initapp`.
- [RISK] Multiple defaulting and override locations can cause drift if not consolidated.

## Migration Boundary

- [RISK] Demo `AutoMigrate`, `initdb`, and SQL scripts coexist.
- [NEEDS_CONFIRMATION] Confirm whether:
  - `AutoMigrate` is dev/demo only.
  - SQL scripts are production/bootstrap only.
  - A layered strategy is required.

## Demo Boundary

- [CONFIRMED] Current demo Todo API exists as an example.
- [NEEDS_CONFIRMATION] Decide whether the demo module is canonical, temporary, or removable.
- [RISK] Demo assumptions should not become hidden production requirements.

## Package API Boundary

- [NEEDS_CONFIRMATION] Decide whether `pkg/*` packages are public reusable APIs.
- [RISK] If public, changes require compatibility rules and clearer documentation.
- [RISK] `pkg/sqlgen` has unimplemented capabilities that should be marked deferred or explicitly unsupported.

## Scope Control

- [CONFIRMED] No code changes are allowed during startup intake.
- [INFERRED] Every future optimization must map to a confirmed requirement, task, and time slice.
- [INFERRED] New ideas should enter backlog unless they block current acceptance.

## Security Constraints

- [CONFIRMED] Do not expose `.env` real values.
- [CONFIRMED] Do not deploy automatically.
- [CONFIRMED] Do not run irreversible migrations automatically.
- [RISK] JWT examples in `.env.example` should be reconciled with current auth/rbac non-goal.

