# kb_003: Python Runtime Validator

- ID: `kb_003`
- Source: runtime script replacement slice
- Trust level: `high`
- Applies to: `docs/ai/scripts/validate_runtime.py`
- Version: `0.1.0`
- Updated at: `2026-05-29T17:43:05+08:00`
- Deprecated: `false`

## Fact

The runtime validator for this repository is now the Python script
`docs/ai/scripts/validate_runtime.py`. Run it from the repository root with:

```powershell
python docs/ai/scripts/validate_runtime.py
```

## Evidence

- `docs/ai/scripts/validate_runtime.py`
- `docs/ai/evidence/index.md`

## Checks

- Sensitive information check: `passed`
- Prompt injection check: `passed`
- Related requirements: `req_round_001_python_scripts`
- Related tasks: `task_002_python_scripts`
- Related skills: `skill_context_recovery`, `skill_slice_execution`

