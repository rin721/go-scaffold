# Lifecycle Records

This directory is the default storage location for future per-round downstream
project lifecycle records.

Use `docs/ai/templates/project-lifecycle-downstream-record.yaml` as the
starting shape and store records as:

```text
docs/ai/lifecycle/<round_id>-downstream.yaml
```

Rules:

- Create a downstream lifecycle record after requirement-baseline confirmation.
- Keep research, task analysis, architecture, mode, Agent infrastructure,
  slices, implementation ledger, verification, closure, and next-round state in
  the record.
- Link evidence to `docs/ai/evidence/index.md`.
- Link confirmed decisions to `docs/ai/decisions/records.md`.
- Do not use this directory for secrets, credentials, private production data,
  payment details, or generated application artifacts.
