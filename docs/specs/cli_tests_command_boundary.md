# CLI Tests Command Boundary

## Status

- Task: TASK-P1-006
- Time Slice: TS-P1-006
- Last Updated: 2026-05-27
- Status: COMPLETED

## Goal

- `cmd/server tests` is the Go test entry point.
- The command name, description, and behavior stay aligned.
- This spec does not define the database CLI; current DB behavior is documented in `docs/specs/demo_migration_boundary.md`.

## Current Semantics

| Item | Result | Status |
|---|---|---|
| Command name | `tests` | CONFIRMED |
| Description | `Run Go tests` | CONFIRMED |
| Default behavior | Run `go test ./...` | CONFIRMED |
| Optional scope | `--package=<pattern>` or `-p <pattern>` | CONFIRMED |
| Failure handling | Return wrapped runner errors; do not call `log.Fatal` | CONFIRMED |

## Non-Goals

- Do not add CI/CD here.
- Do not change `server` behavior here.
- Do not move `yaml2go` or add a yaml2go demo command here.

## Verification

- `cmd/server/tests_test.go` fixes command metadata, default package scope, explicit package scope, and runner error return.
- `go test ./cmd/server -count=1` must pass.
- `go test ./... -count=1` must pass before completion.
