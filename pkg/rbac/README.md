# pkg/rbac

`pkg/rbac` is the public infrastructure API for role-based authorization.

It currently provides:

- `Authorizer`, `Principal`, and `Policy` contracts.
- A Casbin-backed authorizer implementation.
- A default RBAC model compatible with role-to-permission grants and glob-style permission patterns such as `users:*` and `*:*`.

The package is intentionally independent from `internal/*`, HTTP handlers, database repositories, and application configuration. Business modules should adapt their own user, role, and permission data into this package's contracts before calling `Authorize`.

Non-goals:

- User management.
- Token/session management.
- Audit logging.
- Secret management.
- Production migration or deployment policy.
