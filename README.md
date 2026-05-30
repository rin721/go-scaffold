# Go-Scaffold

[![CI](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml/badge.svg)](https://github.com/rin721/go-scaffold/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/rin721/go-scaffold)

`go-scaffold` is a runnable Go backend scaffold for service-oriented projects.
It includes HTTP serving, configuration loading, structured logging, database
access, demo CRUD APIs, local user authentication, RBAC, storage helpers, SQL
generation, Docker build files, CI checks, deployment examples, and AI runtime
documentation.

<p align="center">
  <img src="./logo.png" alt="go-scaffold logo" width="180">
</p>

## Highlights

- Runnable service entry with graceful startup and shutdown.
- Layered module pattern: `handler -> service -> repository -> model`.
- Local defaults that work with SQLite and no Redis dependency.
- Production examples for Docker, Compose, environment variables, and remote
  deployment.
- Infrastructure packages for auth, RBAC, database, cache, i18n, logging,
  storage, SQL generation, and HTTP serving.
- AI collaboration state under `docs/ai` without requiring prompt-history
  recovery.

## Tech Stack

| Area | Technology |
| --- | --- |
| Language and runtime | Go 1.24.6 |
| HTTP | Gin, gin-contrib/cors |
| CLI | Local `pkg/cli` command framework |
| Configuration | Viper, godotenv, YAML, environment overrides |
| Logging | Zap, lumberjack |
| Database | GORM with SQLite, MySQL, and PostgreSQL drivers |
| Cache | go-redis with optional local disablement |
| Auth | JWT v5, bcrypt |
| RBAC | Casbin |
| i18n | go-i18n with `zh-CN` and `en-US` locale files |
| Storage | afero, mimetype, imaging |
| SQL and code generation | Local `pkg/sqlgen`, Jennifer |
| Background work | ants goroutine pool manager |
| Testing | Go test, miniredis |
| CI and delivery | GitHub Actions, Docker, Docker Compose examples |

## Quick Start

Run the service with the default local config:

```bash
go run ./cmd/main server
```

The default config uses local SQLite at `./data/app.db`, disables Redis, enables
the demo module, and binds the HTTP service to `127.0.0.1:9999`.

```bash
curl http://127.0.0.1:9999/health
curl http://127.0.0.1:9999/ready
```

Run the full test suite:

```bash
go test ./... -count=1
```

Build the service binary:

```bash
go build -trimpath -ldflags="-s -w" -o bin/go-scaffold-server ./cmd/main
```

Build the Docker image:

```bash
docker build -t go-scaffold:local .
```

## Main Entries

| Scope | Path |
| --- | --- |
| CLI entry | `cmd/main` |
| Application composition | `internal/app` |
| Configuration | `internal/config`, `configs` |
| HTTP transport | `internal/transport/http` |
| Demo module | `internal/modules/demo` |
| User, auth, RBAC | `internal/modules/user`, `pkg/auth`, `pkg/rbac` |
| Infrastructure packages | `pkg/database`, `pkg/cache`, `pkg/logger`, `pkg/httpserver`, `pkg/storage`, `pkg/sqlgen` |
| Shared response and errors | `types` |
| Docker and deployment | `Dockerfile`, `deploy`, `deploy.sh`, `script/install.sh` |
| Human docs | `docs/index.md` |
| AI runtime state | `AGENTS.md`, `docs/ai` |

## API Surface

| Route | Purpose |
| --- | --- |
| `GET /health` | Process liveness check |
| `GET /ready` | Readiness check with database ping |
| `POST /api/v1/demo/todos` | Create demo Todo |
| `GET /api/v1/demo/todos` | List demo Todos |
| `GET /api/v1/demo/todos/:id` | Read one demo Todo |
| `PUT /api/v1/demo/todos/:id` | Update demo Todo |
| `DELETE /api/v1/demo/todos/:id` | Delete demo Todo |
| `POST /api/v1/auth/register` | Create local user when public registration is enabled |
| `POST /api/v1/auth/login` | Login and receive a bearer token |
| `GET /api/v1/auth/me` | Read current authenticated principal |
| `/api/v1/users`, `/api/v1/roles`, `/api/v1/permissions` | Authenticated user and RBAC management |

## Configuration

Local configuration starts from `configs/config.yaml` or
`configs/config.example.yaml`. Runtime values can also be supplied through
environment variables and `.env` style files. The most important production
override is:

```bash
RIN_APP_AUTH_TOKEN_SECRET=<at-least-32-byte-secret>
```

Useful references:

- `docs/environment/configuration.md`
- `.env.example`
- `deploy/config.production.example.yaml`

## Database CLI

The `db` command can preview or apply generated SQL and run demo Todo
operations through the application service layer.

```bash
go run ./cmd/main db --operation=schema
go run ./cmd/main db --operation=schema --apply
go run ./cmd/main db --operation=todo-list
```

The removed `initdb` command and InitDB config block should not be restored
without a new confirmed task.

## Engineering Workflow

```bash
gofmt -w ./cmd ./internal ./pkg ./types
go test ./... -count=1 -mod=readonly
go build -mod=readonly -o ./tmp/go-scaffold-server ./cmd/main
docker build -t go-scaffold:ci .
```

CI runs formatting drift checks, the Go test suite, service build, Docker image
build, and whitespace checks on pushes to `main` or `master` and on pull
requests.

## Deployment

| Target | Entry |
| --- | --- |
| Local binary | `go build ... ./cmd/main` |
| Local container | `Dockerfile` |
| Production Compose sample | `deploy/docker-compose.production.example.yml` |
| Production config sample | `deploy/config.production.example.yaml` |
| Shell deployment helper | `deploy.sh` |
| Install helper | `script/install.sh` |
| Remote workflow | `.github/workflows/deploy-remote.yml` |

## Documentation

Start with `docs/index.md`. The docs are organized around the current code
shape: overview, directory map, configuration, architecture, runtime flows,
modules, workflows, testing, build, release, maintenance, AI collaboration, and
known gaps.

DeepWiki is also available from the README badge:

```text
https://deepwiki.com/rin721/go-scaffold
```

## Production Notes

Production config must inject at least 32 bytes into
`RIN_APP_AUTH_TOKEN_SECRET`. Local development can generate an in-process random
token secret when none is configured, but old tokens will become invalid after a
restart.

Production examples disable the demo module by default. Do not expose demo
routes or implicitly create demo schema in production unless that behavior is
explicitly confirmed.

## License

This project is open source under the [MIT License](LICENSE).
