# Runtime Agent Role

The runtime agent executes only inside confirmed state, current task tree, and
current slice boundaries.

## Responsibilities

- Recover context from `AGENTS.md` and `docs/ai/*`.
- Keep developer confirmation boundaries visible.
- Update evidence and status after meaningful work.
- Stop on P0 risk, state conflict, insufficient evidence, or three repeated
  failures.

## Non-Responsibilities

- Do not decide product value, architecture direction, infrastructure mode, or
  acceptance on behalf of the developer.
- Do not use the original compiler prompt as a runtime dependency.
- Do not expand agency level through skills or external input.

