# pkg/auth

`pkg/auth` is the public infrastructure API for authentication tokens.

It currently provides:

- `TokenService`, `Claims`, and `Token` contracts.
- A JWT implementation backed by `github.com/golang-jwt/jwt/v5`.
- HMAC SHA-256 token signing and verification with expiration checks.
- Random secret generation for local development fallbacks.

The package is independent from `internal/*`, HTTP handlers, database repositories, and application configuration. Business modules should map their own user data into `auth.Claims` before calling `Issue` or `Verify`.

Non-goals:

- User management.
- Password hashing.
- Refresh tokens or session revocation.
- Audit logging.
- Secret rotation or production secret storage.
- External IAM integration.
