# Quality Gates

## Minimum Runtime Gate

Run:

```powershell
python docs/ai/scripts/validate_runtime.py
```

The gate checks:

- Required runtime files exist.
- Task forest, mainline idea tree, and Vibe Coding infrastructure branch exist.
- Required minimum governance paths exist.
- `docs/ai/status/current.yaml` branch and slice match `docs/ai/tasks/current-slice.yaml`.
- `AGENTS.md` and runtime docs do not instruct agents to read the original
  compiler prompt as a normal recovery path.
- The status file points to a current slice and next allowed phase.

## Round Implementation Gate

Before business implementation:

1. Confirm requirement baseline.
2. Record research validation or state why no external research is needed.
3. Confirm task analysis.
4. Confirm architecture and stack.
5. Confirm infrastructure mode and agency level.
6. Create task tree and execution slices.
7. Run available tests, build, lint, smoke checks, or documented alternatives.
8. Sync docs, knowledge, skills, status, and evidence.
