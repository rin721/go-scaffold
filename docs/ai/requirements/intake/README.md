# Requirement Intake Records

This directory is the default storage location for future per-round project
requirement discovery records.

Use `docs/ai/templates/project-requirement-discovery-record.yaml` as the
starting shape and store records as:

```text
docs/ai/requirements/intake/round_<nnn>_<feature_slug>.yaml
```

Future agents should select a probe template from
`docs/ai/requirements/template_discovery.yaml` and honor the
`REQUIREMENTS_GATHERING` lock in `docs/ai/requirements/state_machine.yaml`
before creating or updating a record.

Rules:

- Raw developer ideas are stored as discovery input, not as confirmed
  requirements.
- AI-inferred details remain candidates until the developer confirms them.
- Confirmed, rejected, deferred, out-of-scope, or research-needed items are
  synchronized to `docs/ai/requirements/ledger.yaml`.
- Open questions remain in the intake record and are summarized in handoff.
- Do not store secrets, credentials, private production data, or payment
  details in this directory.
