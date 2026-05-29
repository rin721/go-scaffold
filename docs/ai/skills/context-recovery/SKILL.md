# Context Recovery Skill

## Trigger Conditions

- The user says "next step".
- A new agent resumes this repository.
- `docs/ai/status/current.yaml` and `docs/ai/tasks/current-slice.yaml` need to
  be reconciled.

## Inputs

- `AGENTS.md`
- `docs/ai/status/current.yaml`
- `docs/ai/tasks/forest.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/tasks/bootstrap-tree.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/evidence/index.md`
- `docs/ai/handoff/current.md`
- `git status --short --branch`

## Boundaries

- Do not read the original compiler prompt as a normal recovery path.
- Do not start a new round task tree until the current task tree is closed,
  explicitly discarded, migrated, or partially closed with confirmation.
- Do not modify business code during state recovery.

## Steps

1. Read `AGENTS.md`.
2. Read `docs/ai/status/current.yaml`.
3. Read `docs/ai/tasks/forest.yaml`.
4. Read `docs/ai/tasks/current-slice.yaml`.
5. Confirm the status `current_branch` matches the slice `branch_id`.
6. Check `docs/ai/evidence/index.md` and `docs/ai/handoff/current.md`.
7. Run `git status --short --branch`.
8. If entries conflict, enter `state_recovery_patch` and update only runtime
   state artifacts.
9. If entries are consistent, choose the next action from current slice status.

## Verification

- State and current slice point to the same phase or a documented transition.
- Evidence entry exists for any state mutation.
