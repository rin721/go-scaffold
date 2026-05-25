# CHANGELOG.md

## 最新变更

### 2026-05-26 - TASK-PHASE6-001 - TS-PHASE6-001

- 变更：用户选择 A，确认进入 Phase 6 收尾与交接。
- 变更：关闭 TASK-NEXT-SCOPE-008，记录 TASK-PHASE6-001 / TS-PHASE6-001 完成。
- 变更：更新状态、任务、时间切片、验收、测试矩阵、路线图、项目简介、风险、Backlog、决策、问题记录、测试报告和交接说明。
- 范围：未新增或修改 Go 源码、测试文件、依赖、数据库 schema、部署配置或密钥。
- 验证：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：Phase 6 收尾完成；当前无自动下一实现任务，后续工作需要用户重新确认。

### 2026-05-26 - TASK-P1-015 - TS-P1-015

- 变更：用户选择 B，确认提升 `BL-002` 的 router/middleware/demo HTTP 集成测试部分。
- 变更：新增 `internal/transport/http/router_integration_test.go`，使用临时 SQLite 和真实 demo repository/service/handler 注入 `NewRouter`。
- 变更：覆盖 demo Todo HTTP Create/List/Get/Update/Delete、删除后 404、CORS preflight/actual origin header、TraceID header round-trip，以及 Recovery 500 响应 traceId 和 logger 调用。
- 修复：前两次相关包测试失败来自测试构造问题：`httptest.NewRequest` 默认 Host 与 Origin 同源，导致 CORS 中间件跳过；固定测试 Host 为 `api.local` 后通过。
- 验证：
  - `gofmt -w internal/transport/http/router_integration_test.go`：PASS
  - `go test ./internal/transport/http ./internal/middleware ./internal/modules/demo/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-015 COMPLETED；TASK-NEXT-SCOPE-008 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-014 - TS-P1-014

- 变更：用户选择 B，确认提升 `BL-023`，为 `pkg/utils` 内部支撑工具补最小确定性测试。
- 变更：新增 `pkg/utils/utils_test.go`，覆盖 Snowflake 生成/非法 node、监听地址校验、端口范围与 exclude、设备 ID 稳定/盐值和 i18n helper 默认语言转发。
- 修复：前两次包测试失败来自测试代码对端口占用语义的环境假设；改为确定性无效地址和 exclude/range 断言后通过。
- 验证：
  - `gofmt -w pkg/utils/utils_test.go`：PASS
  - `go test ./pkg/utils -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-014 COMPLETED；TASK-NEXT-SCOPE-007 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-013 - TS-P1-013

- 变更：用户选择 A，确认继续 `BL-020` 剩余范围，提升第三批 `pkg/cache` 隔离行为测试。
- 变更：新增 `pkg/cache/cache_test.go`，使用进程内 Redis 测试服务覆盖配置默认值、配置校验、连接、读写、缺失键、过期、批量操作、计数器和 reload 语义。
- 变更：新增纯测试依赖 `github.com/alicebob/miniredis/v2`，同步更新 `go.mod` 和 `go.sum`。
- 修复：首次包测试为测试代码编译失败，原因是误读 `miniredis.Get` 返回值；修正断言后通过。
- 验证：
  - `go get github.com/alicebob/miniredis/v2@latest`：PASS
  - `gofmt -w pkg/cache/cache_test.go`：PASS
  - `go test ./pkg/cache -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-013 COMPLETED；TASK-NEXT-SCOPE-006 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-012 - TS-P1-012

- 变更：用户发送“下一步”，确认继续 `BL-020` 第二批 `pkg/*` 行为测试。
- 变更：新增 `pkg/executor/executor_test.go`，覆盖配置校验、任务执行、缺失池、过载、关闭、失败 reload 和 panic handler。
- 变更：新增 `pkg/httpserver/httpserver_test.go`，覆盖构造、默认配置、配置错误、停止态 reload/shutdown 和已运行 start 拒绝路径。
- 变更：新增 `pkg/storage/storage_test.go`，覆盖内存文件系统读写、复制、MIME、Excel、图片和配置错误路径。
- 修复：`pkg/executor` 缺失池、过载和重复配置错误改为包装公开 sentinel，支持 `errors.Is` 判断。
- 修复：`pkg/executor` panic 恢复路径现在会调用通过 `SetPanicHandler` 注册的 handler。
- 验证：
  - `gofmt -w pkg/executor/executor_test.go pkg/httpserver/httpserver_test.go pkg/storage/storage_test.go`：PASS
  - `go test ./pkg/executor ./pkg/httpserver ./pkg/storage -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-012 COMPLETED；TASK-NEXT-SCOPE-005 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-011 - TS-P1-011

- 变更：新增 `pkg/cli/app_test.go`，覆盖命令注册、flag/env/args 解析、help/version 输出、usage error 和 command error 包装。
- 变更：新增 `pkg/i18n/i18n_test.go`，覆盖 JSON/YAML 消息加载、模板渲染、默认语言回退、缺失消息 fallback、`MustT` panic 和加载错误路径。
- 变更：新增 `pkg/yaml2go/converter_test.go`，覆盖多文件生成、空输入、非法 YAML、配置校验，并用 Go parser 校验生成代码合法性。
- 修复：`pkg/yaml2go` 使用 Jennifer tag map 生成合法 struct tag，避免输出 `:"..."` 形式的非法 tag。
- 修复：`pkg/yaml2go` 将子配置 struct 与方法追加到同一个 Jennifer 文件，避免 import 块被拼接到声明之后。
- 验证：
  - `gofmt -w pkg/cli/app_test.go pkg/i18n/i18n_test.go pkg/yaml2go/converter_test.go`：PASS
  - `go test ./pkg/cli ./pkg/i18n ./pkg/yaml2go -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-011 COMPLETED；TASK-NEXT-SCOPE-004 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-NEXT-SCOPE-003 - TS-NEXT-SCOPE-003

- 变更：用户选择 A，确认提升 `BL-020` 补 `pkg/*` 行为测试。
- 变更：将 `BL-020` 首批拆为 TASK-P1-011 / TS-P1-011，覆盖 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 最小行为测试。
- 变更：新增状态诊断报告 `docs/reports/status_diagnostics/2026-05-25-task-p1-011-handoff-stale.md`，修复交接和测试报告滞后。
- 变更：更新 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`BACKLOG.md`、`RISK_REGISTER.md`、`DECISIONS.md`、`TEST_REPORT.md`、`ISSUES.md` 和 `AGENT_HANDOFF.md`。
- 测试：
  - 状态一致性文本检查：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
  - Go 测试未运行：本切片仅修改项目文档，未改 Go 代码
- 状态：TASK-NEXT-SCOPE-003 COMPLETED；TASK-P1-011 NOT_STARTED。

### 2026-05-25 - TASK-P1-010 - TS-P1-010

- 变更：接受用户修正，`pkg/plugin` 注册责任改为被动 registry/runtime。
- 变更：新增状态诊断报告 `docs/reports/status_diagnostics/2026-05-25-agent-handoff-stale-after-types-boundary.md`，记录并修复交接文件滞后。
- 变更：`Manager` 接口移除 `Load`、`RegisterLocalFactory` 和 manager option 主动装配公共面。
- 变更：新增 `NewHTTP` 和 HTTP option，使 HTTP 插件可由插件服务或宿主装配层显式构造后 `Register`。
- 变更：移除 local factory API 和相关错误，local 插件改为服务侧构造后注册。
- 变更：更新 `pkg/plugin` README、package doc、架构、模块清单、测试矩阵、验收、风险、Backlog、状态和交接文档。
- 测试：
  - `gofmt -w pkg/plugin/manager.go pkg/plugin/http.go pkg/plugin/constants.go pkg/plugin/errors.go pkg/plugin/doc.go pkg/plugin/config.go pkg/plugin/local.go pkg/plugin/plugin_test.go`：PASS
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-010 COMPLETED；TASK-NEXT-SCOPE-003 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-NEXT-SCOPE - TS-NEXT-SCOPE

- 变更：用户选择选项 A，确认提升 `BL-021` / `TM-P1-005`。
- 变更：新增 TASK-P1-009 / TS-P1-009，作为当前唯一合法下一步，用于明确 `types/*` 契约边界。
- 变更：更新 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_MATRIX.md`、`ACCEPTANCE.md`、`BACKLOG.md`、`RISK_REGISTER.md` 和交接相关文档，关闭待确认状态。
- 测试：
  - 状态一致性文本检查：PASS
  - `go test ./types/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-NEXT-SCOPE COMPLETED；TASK-P1-009 NOT_STARTED。

### 2026-05-25 - TASK-P1-008 - TS-P1-008

- 变更：新增 `ErrCodeUnsupportedOperation` 和 `NewUnsupportedError`，统一表示 `pkg/sqlgen` 未支持能力。
- 变更：`Or`、`Not`、`Group`、`Having`、`Distinct`、`Joins` 不再静默 no-op，后续 SQL 生成会返回 unsupported 错误。
- 变更：`DeleteInBatches` 不再退化为普通删除，直接返回 unsupported 错误。
- 变更：`ReverseDB(...).Generate`、`GenerateAll`、`GenerateToDir` 返回 unsupported 错误。
- 变更：`pkg/sqlgen/README.md` 标注 unsupported / partial 能力边界，`doc.go` 不再声称完整 GORM 兼容。
- 变更：新增 `pkg/sqlgen` unsupported 行为测试。
- 测试：
  - `gofmt -w pkg/sqlgen/errors.go pkg/sqlgen/types.go pkg/sqlgen/generator.go pkg/sqlgen/query.go pkg/sqlgen/update.go pkg/sqlgen/delete.go pkg/sqlgen/reverse.go pkg/sqlgen/doc.go pkg/sqlgen/sqlgen_test.go`：PASS
  - `go test ./pkg/sqlgen -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-008 COMPLETED；后续范围 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-P1-007 - TS-P1-007

- 变更：为 13 个 `pkg/*/README.md` 新增 API 分类段。
- 变更：将 `pkg/cache`、`pkg/crypto`、`pkg/database`、`pkg/executor`、`pkg/httpserver`、`pkg/i18n`、`pkg/logger`、`pkg/plugin`、`pkg/storage` 标注为公共基础设施 API。
- 变更：将 `pkg/cli`、`pkg/sqlgen`、`pkg/yaml2go` 标注为公共工具 API。
- 变更：将 `pkg/utils` 标注为内部支撑工具包。
- 变更：同步 `ARCHITECTURE.md` 和 `MODULES.md`，记录稳定边界、测试缺口和后续约束。
- 变更：将当前合法下一步推进为 TASK-P1-008。
- 测试：
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows LF/CRLF 转换警告
- 状态：TASK-P1-007 COMPLETED；TASK-P1-008 NOT_STARTED。

### 2026-05-25 - TASK-P1-006 - TS-P1-006

- 变更：`cmd/server tests` 从 yaml2go 示例转换改为真实 Go test 入口。
- 变更：新增 `--package/-p`，默认执行 `go test ./...`，可指定 package pattern。
- 变更：移除 `TestsCommand.Execute` 中的 `log.Fatal` 行为，runner 失败时返回可包装错误。
- 变更：新增 `cmd/server/tests_test.go`，覆盖命令元信息、默认包范围、指定包范围和失败返回。
- 变更：新增 `docs/specs/cli_tests_command_boundary.md`，记录 CLI tests 命令语义。
- 变更：将当前合法下一步推进为 TASK-P1-007。
- 测试：
  - `go test ./cmd/server -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows CRLF 转换警告
- 状态：TASK-P1-006 COMPLETED；TASK-P1-007 NOT_STARTED。

### 2026-05-25 - TASK-P1-005 - TS-P1-005

- 变更：新增 demo 迁移触发策略，显式区分 `server-start`、`initdb` 和 `reload`。
- 变更：`NewModules` 继续在 server 启动路径执行 demo `AutoMigrate`，`BuildInitDB` 继续作为显式 demo bootstrap 执行迁移。
- 变更：`reloadDatabase` 改为使用 reload 策略，数据库 reload 不再隐式执行 demo schema 迁移。
- 变更：新增 `internal/app/initapp/demo_migration_test.go`，用隔离 SQLite 验证触发策略。
- 变更：新增 `docs/specs/demo_migration_boundary.md`，记录 dev/demo 与生产/bootstrap 迁移职责。
- 变更：将当前合法下一步推进为 TASK-P1-006。
- 测试：
  - `go test ./internal/app/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows CRLF 转换警告
- 状态：TASK-P1-005 COMPLETED；TASK-P1-006 NOT_STARTED。

### 2026-05-25 - TASK-P1-004 - TS-P1-004

- 变更：新增 `internal/modules/demo/service/todo_test.go`，为 demo Todo 建立 service/repository CRUD 测试基线。
- 变更：使用临时 SQLite 执行真实 repository/service 路径，不依赖外部数据库或 HTTP server。
- 变更：覆盖 Create/List/Get/Update/Delete 成功路径、空标题校验、缺失资源 not found 和软删除后不可见语义。
- 变更：将当前合法下一步推进为 TASK-P1-005。
- 测试：
  - `go test ./internal/modules/demo/... -count=1`：PASS
  - `go test ./... -count=1`：PASS
  - `git diff --check`：PASS，仅有 Windows CRLF 转换警告
- 状态：TASK-P1-004 COMPLETED；TASK-P1-005 NOT_STARTED。

### 2026-05-25 - TASK-P1-003 - TS-P1-003

- 变更：新增 `internal/transport/http/router_test.go`，用 `httptest` 固定 `/health` 和 `/ready` 行为。
- 变更：`/health` 覆盖 HTTP 200、`code=0`、`message=success`、`data.status=ok`。
- 变更：`/ready` 覆盖数据库缺失、ping 失败、ping 成功三条路径，断言 HTTP 状态码、`data.status` 和 `data.checks.database`。
- 变更：将当前合法下一步推进为 TASK-P1-004。
- 测试：
  - `go test ./internal/transport/http -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-P1-003 COMPLETED；TASK-P1-004 NOT_STARTED。

### 2026-05-25 - TASK-P1-002 - TS-P1-002

- 变更：数据库环境变量覆盖改为优先读取 `DB_*`，旧 `REI_APP_DB_*` 保留为兼容 fallback。
- 变更：`.env.example` 对齐实际环境变量策略，补齐 Storage/CORS 示例，移除未实现的 JWT 示例。
- 变更：新增配置测试，覆盖 `DB_*` 主策略、旧前缀 fallback、Redis/Server/Logger/I18n 环境变量覆盖。
- 变更：将当前合法下一步推进为 TASK-P1-003。
- 测试：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-P1-002 COMPLETED；TASK-P1-003 NOT_STARTED。

### 2026-05-25 - TASK-INFRA-002 - TS-INFRA-002

- 变更：新增实际缺失的 `AGENTS.md`，修复状态文档与文件系统事实冲突。
- 变更：统一 `CLAUDE.md`、`AGENT_RULES.md`、Cursor、Kiro、Codex 配置对 `AGENTS.md` 和 `docs/ai/prompt.md` 的引用。
- 变更：扩充 14 个 canonical `skills/*/SKILL.md`，补齐 YAML frontmatter 和完整执行结构。
- 变更：新增 14 个 `.agents/skills/*/SKILL.md` 轻量适配器。
- 变更：将 `docs/templates/*` 标准化为可复用模板，项目事实继续保留在根目录项目文档。
- 变更：新增状态诊断报告 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md`。
- 测试：
  - Agent 基础设施文件存在性核对：PASS
  - `quick_validate.py` 验证 28 个 skill 目录：PASS
  - 跨工具入口引用一致性检查：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-INFRA-002 COMPLETED；TASK-P1-002 NOT_STARTED。

### 2026-05-25 - TASK-INFRA-001 - TS-INFRA-001

- 变更：补齐 `docs/ai/prompt.md` 要求的跨 Agent 入口、规则和 skills 索引。
- 变更：新增任务拆分模板、时间切片模板、reports/specs 目录入口和跨工具目录入口。
- 变更：新增 14 个项目专用 `skills/*/SKILL.md`。
- 变更：记录 TASK-INFRA-001 完成，并恢复当前合法下一步为 TASK-P1-002。
- 测试：
  - Prompt 全量产物存在性核对：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-INFRA-001 COMPLETED；TASK-P1-002 NOT_STARTED。

## 历史变更

### 2026-05-25 - TASK-P1-001 - TS-P1-001

- 变更：用户再次发送“下一步”，按推荐默认顺序确认 P1 执行顺序。
- 变更：修复 `internal/config/manager.go` 的 `copyConfig`，改为完整结构体复制并深拷贝 slice。
- 变更：新增 `internal/config/manager_test.go`，覆盖字段完整复制、slice 深拷贝和 `Update` 保留未修改字段。
- 变更：将当前合法下一步推进为 TASK-P1-002。
- 测试：
  - `go test ./internal/config -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-005 COMPLETED；TASK-P1-001 COMPLETED；TASK-P1-002 NOT_STARTED。

### 2026-05-25 - TASK-OPT-004 - TS-OPT-004

- 变更：新增 `TEST_MATRIX.md`，正式记录 P0/P1 测试矩阵、验证命令、退出条件和推荐执行顺序。
- 变更：新增 `ISSUES.md`，补齐项目问题记录入口。
- 变更：在 `TASKS.md` 和 `TIME_SLICES.md` 中写入 TASK-P1-001 至 TASK-P1-008、TS-P1-001 至 TS-P1-008 草案。
- 变更：将当前合法下一步推进为 TASK-OPT-005，等待确认 P1 执行顺序。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-004 COMPLETED；TASK-OPT-005 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-OPT-003 - TS-OPT-003

- 变更：新增 `MODULES.md`，记录模块职责、依赖方向、边界冲突、测试矩阵草案和 P1 优化候选项。
- 变更：确认 `.env.example` 与数据库环境变量前缀不一致、`copyConfig` 字段复制不完整、demo 自动迁移触发点分散、`cmd/server tests` 语义不一致等为优先收口风险。
- 变更：更新需求、验收、路线图、Backlog、任务、时间切片、状态、测试报告和交接文档。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-003 COMPLETED；TASK-OPT-004 NOT_STARTED。

### 2026-05-25 - TASK-OPT-002 - TS-OPT-002

- 变更：按用户“下一步”确认推荐默认值。
- 变更：确认治理优先、`pkg/*` 混合策略、demo 长期标准示例、迁移 dev-prod 分层、中文化根文档和模板优先。
- 变更：新增 `ROADMAP.md`。
- 变更：更新需求、架构、决策、任务、时间切片、状态、验收、风险和交接文档。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-002 COMPLETED；TASK-OPT-003 NOT_STARTED。

### 2026-05-25 - TASK-OPT-001 - TS-OPT-001

- 变更：重新启动全项目治理与优化路线主线。
- 变更：将六个启动模板重写为中文：
  - `docs/templates/project_start_template.md`
  - `docs/templates/requirements_clarification_template.md`
  - `docs/templates/technical_options_template.md`
  - `docs/templates/architecture_constraints_template.md`
  - `docs/templates/acceptance_template.md`
  - `docs/templates/risk_confirmation_template.md`
- 变更：更新根目录项目文档和状态文件，使当前合法任务从插件系统扩展切换为项目优化启动确认。
- 变更：将插件系统 v1 内容保留为历史记录和 Backlog，不作为当前主线继续扩展。
- 测试：
  - `go test ./... -count=1`：PASS
- 状态：TASK-OPT-001 COMPLETED；TASK-OPT-002 PENDING_USER_CONFIRMATION。

### 2026-05-25 - TASK-HIST-PLUGIN-002 - TS-HIST-PLUGIN-002

- 历史：接受并关闭 `pkg/plugin` v1 local/http API 边界。
- 测试：
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：COMPLETED。

### 2026-05-25 - TASK-HIST-PLUGIN-001 - TS-HIST-PLUGIN-001

- 历史：实现独立 `pkg/plugin` 包，支持 local 和 HTTP 协议。
- 测试：
  - `go test ./pkg/plugin -count=1`：PASS
  - `go test ./... -count=1`：PASS
- 状态：COMPLETED。
