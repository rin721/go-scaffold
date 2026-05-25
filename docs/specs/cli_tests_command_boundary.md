# CLI Tests Command Boundary

## 状态

- 任务：TASK-P1-006
- 时间切片：TS-P1-006
- 最后更新：2026-05-25
- 状态：COMPLETED

## 目标

- [CONFIRMED] `cmd/server tests` 不应再执行 yaml2go 示例转换。
- [CONFIRMED] `tests` 命令应与名称和描述一致，作为 Go 测试入口。
- [CONFIRMED] 本切片不调整其他 CLI 命令、不重命名应用主命令、不修改 `pkg/cli`。

## 当前语义

| 项 | 结果 | 状态 |
|---|---|---|
| 命令名 | `tests` | [CONFIRMED] |
| 描述 | `Run Go tests` | [CONFIRMED] |
| 默认行为 | 执行 `go test ./...` | [CONFIRMED] |
| 可选范围 | `--package=<pattern>` 或 `-p <pattern>` | [CONFIRMED] |
| 失败处理 | 返回包装后的 runner error，不调用 `log.Fatal` | [CONFIRMED] |

## 非目标

- [CONFIRMED] 不新增 CI/CD。
- [CONFIRMED] 不改变 `server` 或 `initdb` 行为。
- [CONFIRMED] 不移动 yaml2go 包，也不新增 yaml2go demo 命令。

## 验证

- `cmd/server/tests_test.go` 固定命令元信息、默认包范围、指定包范围和 runner 失败返回。
- `go test ./cmd/server -count=1` 必须通过。
- `go test ./... -count=1` 必须通过。
