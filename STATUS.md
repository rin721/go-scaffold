# STATUS.md

## 项目状态

- 项目：go-scaffold
- 当前阶段：P2 Linux Docker production 部署制品验证阻塞
- 总体状态：BLOCKED
- 最后更新：2026-05-27
- 最近 Agent：Codex
- 最近工具：Codex Desktop

## 当前合法工作

- 当前模块：项目优化路线
- 当前任务 ID：TASK-P2-004
- 当前时间切片 ID：TS-P2-004
- 当前状态：BLOCKED
- 为什么这是当前唯一合法状态：[CONFIRMED] 用户发送“下一步”后按协议处理剩余待验证项。TASK-P2-005 至 TASK-P2-010 已完成并验证；当前唯一未关闭事项是 TASK-P2-004 的 Docker image build。2026-05-27 当前本机仍无可用 Docker 兼容 CLI；用户在 Docker 环境补跑时 `go mod download` 因访问 `proxy.golang.org` 超时失败，随后发现旧 Dockerfile 未声明 `GOPROXY` build arg，导致 `--build-arg GOPROXY=...` 未生效。本轮已补 Dockerfile 的 `GOPROXY` / `GOSUMDB` build arg 和 BuildKit 缓存，`ISSUE-P2-005` 保持打开，待 Docker 环境用更新后的 Dockerfile 重跑。

## 阶段状态

| 阶段 | 状态 | 证据 |
|---|---|---|
| 项目启动 | COMPLETED | `PROJECT_BRIEF.md` 和 `docs/templates/*` 已中文化并切回项目优化主线 |
| 需求 | COMPLETED | `REQUIREMENTS.md` 已记录确认结果 |
| 高层架构 | COMPLETED | `ARCHITECTURE.md` 已记录确认边界 |
| 路线图 | COMPLETED | `ROADMAP.md` 已生成 |
| 模块边界清单 | COMPLETED | `MODULES.md` 已生成 |
| 测试矩阵与任务拆分 | COMPLETED | `TEST_MATRIX.md` 已生成，`TASKS.md` 和 `TIME_SLICES.md` 已写入 P1 草案 |
| P1 执行顺序确认 | COMPLETED | 用户再次发送“下一步”，按推荐默认顺序确认 |
| 配置 copy/update 测试与修复 | COMPLETED | `internal/config/manager_test.go` 已新增，`copyConfig` 已修复 |
| 配置环境变量策略收拢 | COMPLETED | `DB_*` 成为数据库主环境变量，`REI_APP_DB_*` 保留兼容 fallback；`.env.example` 与实现一致 |
| HTTP health/ready smoke test | COMPLETED | `internal/transport/http/router_test.go` 已覆盖 `/health`、`/ready` missing/failure/ready 路径 |
| demo CRUD 测试基线 | COMPLETED | `internal/modules/demo/service/todo_test.go` 已用临时 SQLite 覆盖 Todo Create/List/Get/Update/Delete |
| demo 迁移边界收拢 | COMPLETED | `DemoMigrationPolicyFor` 固定 server-start/initdb/reload 策略，reload 不再隐式执行 demo `AutoMigrate` |
| CLI tests 命令语义收拢 | COMPLETED | `cmd/server tests` 现在执行 `go test`，并由 `cmd/server/tests_test.go` 固定命令语义 |
| pkg/* API 分类 | COMPLETED | 13 个 `pkg/*` 包已在各自 README、`ARCHITECTURE.md`、`MODULES.md` 中标注 API 定位 |
| pkg/sqlgen unsupported 边界标注 | COMPLETED | unsupported 链式查询、批量删除和 DB reverse 已显式返回 `ErrCodeUnsupportedOperation`，README 已标注部分能力边界 |
| Agent 基础设施补齐 | COMPLETED | `AGENTS.md`、`AGENT_RULES.md`、`SKILLS.md`、项目 skills、reports/specs 和跨工具目录已补齐 |
| Agent 基础设施一致性修复 | COMPLETED | TASK-INFRA-002 已补齐实际缺失的 `AGENTS.md`，规范化 skills、模板和 `.agents` 适配器 |
| Agent 状态一致性修复 | COMPLETED | TASK-INFRA-003 已生成状态诊断报告，并修复 TASK-P1-016/017 后背景文档中的旧待办表述 |
| types/* 契约边界 | COMPLETED | TASK-P1-009 已补契约说明和最小测试，`go test ./types/... -count=1` 与全量回归通过 |
| pkg/plugin 被动注册边界 | COMPLETED | TASK-P1-010 已移除 manager 主动配置加载/local factory 公共面，local/http 插件改为服务侧显式 `Register` |
| pkg/* 行为测试首批 | COMPLETED | TASK-P1-011 已补 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 最小行为测试，并修复新增测试暴露的 `pkg/yaml2go` 生成 tag/import 顺序缺陷 |
| pkg/* 行为测试第二批 | COMPLETED | TASK-P1-012 已补 `pkg/executor`、`pkg/httpserver`、`pkg/storage` 最小行为测试，并修复新增测试暴露的 `pkg/executor` 错误包装与 panic handler 缺陷 |
| pkg/cache 行为测试第三批 | COMPLETED | TASK-P1-013 已补 `pkg/cache` 隔离行为测试，使用进程内 Redis 测试服务覆盖配置、读写、批量、计数器、过期和 reload 语义 |
| pkg/utils 内部支撑测试 | COMPLETED | TASK-P1-014 已新增 `pkg/utils/utils_test.go`，覆盖 Snowflake、地址校验、端口查找、设备 ID 和 i18n helper |
| app/router/middleware 集成测试 | COMPLETED | TASK-P1-015 已新增 `internal/transport/http/router_integration_test.go`，覆盖 demo Todo HTTP CRUD、TraceID、CORS 和 Recovery 链路 |
| app 装配与 reload/config 集成测试 | COMPLETED | TASK-P1-016 已新增 `internal/app/app_integration_test.go` 和 `internal/app/reloadapp/reload_test.go`，覆盖真实 app server/initdb 装配、配置变更 hook 和 reload 分发 |
| pkg README 中文化 | COMPLETED | TASK-P1-017 已完成第一阶段 `pkg/*/README.md` 中文化，不修改 Go 代码或依赖 |
| CI 质量门禁与部署说明 | COMPLETED | TASK-P2-001 已新增 GitHub Actions CI workflow、手动部署说明和 README 入口，不执行真实部署或使用密钥 |
| 真实 CD 范围确认 | COMPLETED | 用户选择 C、确认使用远程部署，并进一步确认用 `.env` 风格文件配置；TASK-P2-002 已新增远程部署变量模板 |
| 显式参数部署入口 | COMPLETED | `deploy.sh` 和 `script/install.sh` 已新增，旧本地部署 env 文件已删除，部署说明已同步 |
| 远程部署 workflow | COMPLETED | TASK-P2-003 已新增手动 staging workflow、Secrets 配置说明和远程主机前置条件；未执行真实部署 |
| Linux Docker production 部署制品 | BLOCKED | TASK-P2-004 / TS-P2-004 已实现 Dockerfile、production Compose 示例、手动 production workflow 闸门和统一 `deploy.sh` 部署入口；Dockerfile 已补 Go 代理 build arg 和缓存，但 Docker 构建仍待 Docker 环境重跑通过 |
| 插件钩子运行时与 IAM 公共接口 | COMPLETED | TASK-P2-005 至 TASK-P2-010 已完成；`go test ./pkg/plugin/... -count=1`、`go test ./pkg/iam/... -count=1`、`go test ./internal/config ./internal/app/... -count=1`、`go test ./... -count=1`、server build 和 `git diff --check` 均通过 |
| 部署实现 | COMPLETED | TASK-P2-004 已补齐 Dockerfile、production Compose 示例、production 配置样例、统一 `deploy.sh` 部署入口和手动 production workflow 闸门 |
| 部署验证 | BLOCKED | 脚本 Bash 语法解析、YAML 解析、actionlint、`go test ./... -count=1`、server build 和 `git diff --check` 均通过；Docker build 在远端曾因 Go 代理网络超时失败，Dockerfile 已补代理参数，仍待重跑 |
| 交接 | COMPLETED | `AGENT_HANDOFF.md` 已更新到 TASK-P2-004 / TS-P2-004 阻塞状态，并保留 TASK-P2-005 至 TASK-P2-010 完成验证记录 |
| Phase 6 收尾 | COMPLETED | 用户选择 A 后已完成 TASK-PHASE6-001；最终回归和交接文档已更新 |

## 当前关键发现

| ID | 发现 | 来源 | 状态 |
|---|---|---|---|
| FIND-001 | P1 关键测试缺口已持续收敛 | `go test ./... -count=1`、TASK-P1-003 至 TASK-P1-016 | [CONFIRMED] app/router/demo/config/reload 与主要 `pkg/*` 路径已补最小测试 |
| FIND-002 | `.env.example` 与数据库环境变量前缀不一致 | `MODULES.md` BC-001；TASK-P1-002 已修复 | [CONFIRMED] 已处理 |
| FIND-003 | `manager.copyConfig` 未完整复制配置字段 | `MODULES.md` BC-002；TASK-P1-001 已修复 | [CONFIRMED] 已处理 |
| FIND-004 | demo schema 自动迁移触发点需收拢 | `MODULES.md` BC-003；TASK-P1-005 已固定 server-start/initdb/reload 策略 | [CONFIRMED] 已处理 |
| FIND-005 | `cmd/server tests` 命令语义与行为不一致 | `MODULES.md` BC-004；TASK-P1-006 已改为真实 Go test 入口 | [CONFIRMED] 已处理 |
| FIND-010 | `pkg/*` 公共/内部定位未逐包标记 | `ARCHITECTURE.md`、`MODULES.md`；TASK-P1-007 已完成分类 | [CONFIRMED] 已处理 |
| FIND-011 | `pkg/sqlgen` TODO/unsupported 边界不清 | `pkg/sqlgen` README 和源码；TASK-P1-008 已显式返回 unsupported 或文档化 partial 能力 | [CONFIRMED] 已处理 |
| FIND-012 | `types/result`、错误码和跨层类型边界待明确 | TASK-P1-009 已补 `docs/specs/types_contract_boundary.md`、package doc 和最小测试 | [CONFIRMED] 已处理 |
| FIND-013 | `pkg/plugin` 主动注册服务边界需收拢 | 用户修正；TASK-P1-010 已改为被动 registry/runtime | [CONFIRMED] 已处理 |
| FIND-014 | 背景文档保留 TASK-P1-016 前旧状态 | `ARCHITECTURE.md`、`MODULES.md`、`PROJECT_BRIEF.md`、`ROADMAP.md`；TASK-INFRA-003 已修复 | [CONFIRMED] 已处理 |
| FIND-015 | CI/CD 与部署缺少首个安全边界 | `REQ-OPT-P2-003`、`BL-007`、`BL-008`；用户选择 D | [CONFIRMED] TASK-P2-001 已处理非生产 CI 门禁和部署说明 |
| FIND-016 | 真实 CD 自动化缺少环境与密钥决策 | `BL-024`；用户选择 C、确认远程部署，并确认使用 `.env` 风格配置和实现 workflow | [CONFIRMED] 手动 staging 远程部署 workflow 已补；镜像发布、production 和真实运行仍需单独确认 |
| FIND-017 | production Docker 部署缺少可提交制品 | 用户要求“linux、docker、production -> 部署”；用户修正“环境变量在部署脚本上动态配置”；`BL-024` 剩余范围 | [BLOCKED] 制品和统一 `deploy.sh` 入口已补；远端 Docker build 曾在 `go mod download` 阶段因 Go 代理网络超时失败，Dockerfile 已补 `GOPROXY` build arg 和缓存，真实 production 运行仍不在当前会话执行 |
| FIND-006 | P1 执行顺序尚未确认 | `TEST_MATRIX.md`、`RISK_REGISTER.md` RISK-009；用户再次发送“下一步” | [CONFIRMED] 已确认 |
| FIND-007 | `AGENTS.md` 被状态文件声明已补齐但实际缺失 | `Test-Path AGENTS.md`、`docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md` | [CONFIRMED] 已修复 |
| FIND-008 | `/health`、`/ready` 路由缺少 smoke test | `TEST_MATRIX.md` TM-P0-003；TASK-P1-003 已补测试 | [CONFIRMED] 已处理 |
| FIND-009 | demo Todo CRUD 缺少测试基线 | `TEST_MATRIX.md` TM-P0-005；TASK-P1-004 已补 service/repository 隔离测试 | [CONFIRMED] 已处理 |

## 待用户确认

| ID | 问题 | 影响 | 选项 | Required By |
|---|---|---|---|
| CONFIRM-NEXT-001 | 选择 P1 后续范围或进入收尾 | 已确认：用户选择 A | A: 提升 `BL-021` / `TM-P1-005` 做 `types/*` 契约边界 | COMPLETED |
| CONFIRM-NEXT-002 | 选择 `types/*` 契约边界完成后的后续范围 | 已确认：用户修正并选择收拢 `pkg/plugin` 被动注册边界 | 提升 `BL-022` / `TM-P1-006` | COMPLETED |
| CONFIRM-NEXT-003 | 选择 `pkg/plugin` 被动注册边界完成后的后续范围 | 已确认：用户选择 A，提升 `BL-020` 补 `pkg/*` 行为测试 | A: 提升 `BL-020` 补 `pkg/*` 行为测试 | COMPLETED |
| CONFIRM-NEXT-004 | 选择首批 `pkg/*` 行为测试完成后的后续范围 | 已确认：用户发送“下一步”，按选项 A 继续下一批 `pkg/*` 行为测试 | A: 继续下一批 `pkg/*` 行为测试；B: 进入 Phase 6 收尾；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-005 | 选择第二批 `pkg/*` 行为测试完成后的后续范围 | 已确认：用户选择 A，继续 `BL-020` 剩余包，第三批限定 `pkg/cache` | A: 继续剩余 `pkg/*` 行为测试；B: 进入 Phase 6 收尾；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-006 | 选择 `pkg/cache` 行为测试完成后的后续范围 | 已确认：用户选择 B，提升 `pkg/utils` 内部支撑测试 | A: 进入 Phase 6 收尾；B: 提升内部支撑测试；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-007 | 选择 `pkg/utils` 内部支撑测试完成后的后续范围 | 已确认：用户回复 `b`，选择 B | A: 进入 Phase 6 收尾；B: 提升 app/router/middleware 等集成测试；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-008 | 选择 app/router/middleware 集成测试完成后的后续范围 | 已确认：用户选择 A，进入 Phase 6 收尾 | A: 进入 Phase 6 收尾；B: 继续 app 装配/reload/config 等剩余集成测试；C: 结束本轮 | COMPLETED |
| CONFIRM-NEXT-009 | 选择 TASK-INFRA-003 后的后续方向 | 已确认：用户选择 D，进入 CI/CD 与部署方向首切片 | D: CI/CD 与部署；首切片限定 CI 质量门禁与部署说明 | COMPLETED |
| CONFIRM-NEXT-010 | 确认真实 CD / 镜像发布 / 远程部署自动化边界 | 已确认：用户选择 C、使用远程部署、通过 `.env` 风格模板配置，并明确确认实现远程部署 workflow | TASK-P2-003 已完成 | COMPLETED |
| CONFIRM-NEXT-011 | 确认 Linux/Docker/production 部署制品 | 已确认：用户要求“开始，linux、docker、production -> 部署” | TASK-P2-004 已进入环境阻塞，待 Docker 环境补跑构建 | COMPLETED |

## 待验证

| ID | 任务 | 需要验证内容 | 命令/方法 |
|---|---|---|---|
| VERIFY-P2-004 | TASK-P2-004 | Dockerfile 镜像构建 | [BLOCKED] 当前本机无 `docker`、`podman`、`nerdctl` 或 `docker.exe`；远端曾因 Go 代理网络超时失败。需在 Docker 环境用更新后的 Dockerfile 运行 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .` |
| VERIFY-P2-005 | TASK-P2-005 至 TASK-P2-010 | 插件钩子运行时、远程插件传输、IAM 公共接口和 app 装配 | [CONFIRMED] `go test ./pkg/plugin/... -count=1`；`go test ./pkg/iam/... -count=1`；`go test ./internal/config ./internal/app/... -count=1`；`go test ./... -count=1`；`go build -o <temp> ./cmd/server`；`git diff --check` |

## 需要返工

| ID | 任务 | 原因 | 下一步 |
|---|---|---|---|
|  |  |  |  |

## 最近执行

- 摘要：用户反馈远端 Docker build 在 `go mod download` 阶段很慢/超时；本轮确认旧 Dockerfile 未声明 `GOPROXY` build arg，导致用户传入的代理参数未生效。
- 变更文件：`Dockerfile`、`docs/deployment.md`、项目状态文档与交接文档。
- 执行命令：必读文件读取；Dockerfile/文档检查；`docker version`；`Get-Command docker,podman,nerdctl,docker.exe -ErrorAction SilentlyContinue`；`git diff --check`。
- 测试结果：本机仍无 Docker CLI，未在本机执行 Docker build；远端构建失败来源是依赖下载网络/代理问题；本轮未修改 Go 代码，未运行 Go 测试；`git diff --check` PASS，仅有 Windows LF/CRLF 提示。
- 完成判断：TASK-P2-004 / TS-P2-004 仍为 BLOCKED，`ISSUE-P2-005` 保持 OPEN；TASK-P2-005 至 TASK-P2-010 仍为 COMPLETED。

## 下一步

- 合法下一步：TASK-P2-004 / TS-P2-004 处于 BLOCKED；需在具备 Docker CLI/daemon 的环境后补跑 `docker build --build-arg GOPROXY=https://goproxy.cn,direct -t go-scaffold:local .`。
- 解除阻塞条件：Docker build 通过后，更新 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`ACCEPTANCE.md`、`TEST_REPORT.md`、`CHANGELOG.md`、`ISSUES.md` 和 `AGENT_HANDOFF.md`，再关闭 TASK-P2-004。
- 非目标保持：JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS 传输、生产部署、镜像发布和密钥管理仍不属于本轮完成范围。
