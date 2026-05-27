# DECISIONS.md

## 决策记录

### DEC-001：将插件系统实现为独立 `pkg/plugin`

- Date：2026-05-25
- Status：ACCEPTED
- Context：历史任务要求在 `pkg` 下实现独立插件系统，并支持多协议适配。
- Decision：实现不依赖 `internal/*` 的 `pkg/plugin` 包。
- Alternatives：接入 `internal/app`；使用 Go dynamic plugin；推迟到完整架构阶段。
- Consequences：宿主应用需要显式注册 local plugin factory；库保持跨平台和独立。
- Related Tasks：TASK-HIST-PLUGIN-001

### DEC-002：local 插件使用 factory 注册，不使用 Go dynamic plugin

- Date：2026-05-25
- Status：ACCEPTED
- Context：历史任务希望支持本地插件，但库需要跨平台。
- Decision：提供 `LocalFactory` 注册机制。
- Alternatives：使用 Go `plugin` 包；扫描并编译插件目录；让 app 层感知插件目录。
- Consequences：本地插件加载保持显式、可测试、跨平台。
- Related Tasks：TASK-HIST-PLUGIN-001

### DEC-003：HTTP 插件使用统一 JSON 协议

- Date：2026-05-25
- Status：ACCEPTED
- Context：历史任务需要远程 HTTP 插件能力。
- Decision：HTTP 插件使用 `POST`，请求/响应使用统一 `Request`/`Response` 形状。
- Alternatives：每个插件自定义 payload；GET 调用；协议专用接口。
- Consequences：HTTP 插件服务需要实现文档化 JSON 契约。
- Related Tasks：TASK-HIST-PLUGIN-001

### DEC-004：接受 `pkg/plugin` v1 local/http 边界

- Date：2026-05-25
- Status：ACCEPTED
- Context：`pkg/plugin` v1 已支持 local/http，测试通过，历史 API review 已收尾。
- Decision：当前接受 v1 local/http API 边界。
- Alternatives：继续修改导出 API；立即增加 rpc/ws/discovery。
- Consequences：rpc、ws、discovery、examples 和 app 集成都必须单独提升为任务。
- Related Tasks：TASK-HIST-PLUGIN-002

### DEC-005：当前主线切换为项目治理与优化路线

- Date：2026-05-25
- Status：ACCEPTED
- Context：用户要求按 `docs/ai/prompt.md` 重新启动项目，生成中文启动模板、需求澄清模板、技术方案、架构约束、验收标准和风险确认项。
- Decision：当前合法任务从插件系统后续方向切换为项目优化启动确认；插件系统内容保留为历史和 Backlog。
- Alternatives：继续插件系统 rpc/ws/discovery；直接进入代码重构；全量重写。
- Reason：用户随后发送“下一步”，按当前文档默认值视为确认治理优先路线。
- Consequences：后续“下一步”应先处理用户确认，不直接写代码。
- Related Tasks：TASK-OPT-001

### DEC-006：采用治理优先优化路线

- Date：2026-05-25
- Status：ACCEPTED
- Context：当前项目已有可运行链路和测试基线，但文档主线、包 API、迁移和中文化边界需要先统一。
- Decision：先完成文档治理、模块边界清单、测试矩阵和任务切片，再进入代码优化。
- Alternatives：模块化重构优先；框架化抽取；重写。
- Reason：治理优先风险最低，符合 `docs/ai/prompt.md` 的先确认、先拆分、先验证原则。
- Consequences：代码实现继续阻塞，直到模块边界清单和测试矩阵完成。
- Related Tasks：TASK-OPT-002、TASK-OPT-003

### DEC-007：`pkg/*` 采用混合 API 策略

- Date：2026-05-25
- Status：ACCEPTED
- Context：`pkg/*` 中既有明显可复用基础设施包，也有可能只服务当前脚手架的支撑能力。
- Decision：后续逐包标注为公共 API、内部支撑或待确认；公共 API 需要兼容策略，内部支撑允许按项目需要调整。
- Alternatives：全部视为公共 API；全部视为内部实现。
- Reason：混合策略能避免过早承诺全部包稳定，也避免任意破坏可复用包。
- Consequences：未完成逐包分类前，不做 `pkg/*` 破坏性重构。
- Related Tasks：TASK-OPT-003

### DEC-008：demo 模块暂定为长期标准示例

- Date：2026-05-25
- Status：ACCEPTED
- Context：`internal/modules/demo` 当前展示 Todo CRUD 和 handler/service/repository/model 分层。
- Decision：demo 暂定为长期标准示例，用于后续模块边界说明和测试样板。
- Alternatives：临时占位；后续移除。
- Reason：保留 demo 有助于验证脚手架的默认使用路径。
- Consequences：demo 的示例假设需要和生产约束明确分离。
- Related Tasks：TASK-OPT-003

### DEC-009：迁移采用 dev-prod 分层策略

- Date：2026-05-25
- Status：ACCEPTED
- Context：项目同时存在 demo `AutoMigrate`、`initdb` 命令和 SQL 脚本。
- Decision：开发/demo 可使用 `AutoMigrate`，生产/bootstrap 倾向显式 SQL 或迁移流程，后续需要写清职责边界。
- Alternatives：只用 `AutoMigrate`；只用 SQL 脚本。
- Reason：分层策略兼顾开发便利和生产可控。
- Consequences：数据库相关优化前必须先补充迁移边界和验证策略。
- Related Tasks：TASK-OPT-003

### DEC-010：中文化先覆盖根文档和模板

- Date：2026-05-25
- Status：ACCEPTED
- Context：项目是中文项目，但历史内容和部分包 README 存在中英文混杂。
- Decision：当前阶段优先保证根文档和模板为中文；包 README 与历史内容分阶段处理。
- Alternatives：只中文化根文档；立即全量中文化全部历史文档和包 README。
- Reason：分阶段处理可以避免文档任务过大而影响优化主线。
- Consequences：包 README 中文化进入 Backlog 或后续文档时间切片。
- Related Tasks：TASK-OPT-003

### DEC-011：提升 `types/*` 契约边界为下一任务

- Date：2026-05-25
- Status：ACCEPTED
- Context：TASK-P1-008 完成后，当前已确认 P1 任务列表全部完成，状态进入 `PENDING_USER_CONFIRMATION`，要求在 `BL-021`、`BL-020`、Phase 6 收尾或结束本轮之间选择。
- Decision：用户选择 A，提升 `BL-021` / `TM-P1-005`，下一任务为 TASK-P1-009 / TS-P1-009：明确 `types/*` 契约边界。
- Alternatives：提升 `BL-020` 补 `pkg/*` 行为测试；进入 Phase 6 收尾；结束本轮。
- Reason：`types/result` 依赖 Gin，属于 HTTP 响应契约而非纯类型包；`types/errors` 也包含 auth/rbac 预留错误码，需要先标注边界以免后续复用和重构误判。
- Consequences：下一次“下一步”只能执行 TS-P1-009，不得插队补 `pkg/*` 测试或进入收尾。
- Related Tasks：TASK-NEXT-SCOPE、TASK-P1-009

### DEC-012：`pkg/plugin` 采用被动注册边界

- Date：2026-05-25
- Status：ACCEPTED_WITH_RISK
- Context：用户修正 `pkg/plugin` 的注册责任：`pkg/plugin` 不应主动注册插件服务，而应由插件服务或宿主装配层主动调用注册接口。
- Decision：`pkg/plugin` 保持为被动 registry/runtime。它提供插件接口、local/http 插件实现、`Register`/`Invoke`/`Close` 等运行时能力，但不再通过 manager 公共 API 从配置主动加载并注册插件服务；插件服务负责构造插件实例并注册到 manager。
- Alternatives：保留 `Manager.Load(config)` 主动创建并注册插件；继续使用 local factory 装配；推迟到 rpc/ws/discovery 阶段。
- Reason：被动注册能避免基础设施包反向感知插件服务生命周期，也更符合 `cmd -> internal/app -> pkg` 的依赖方向。
- Consequences：这是对历史 v1 API 的边界收窄，可能影响依赖 `Load` 或 local factory 的调用方；因此必须通过 TASK-P1-010 记录风险、更新 README 并跑 `pkg/plugin` 包测试和全量回归。
- Related Tasks：TASK-NEXT-SCOPE-002、TASK-P1-010

### DEC-013：提升 `BL-020` 首批 `pkg/*` 行为测试

- Date：2026-05-25
- Status：ACCEPTED
- Context：TASK-P1-010 完成后，当前状态进入 `PENDING_USER_CONFIRMATION`，要求在提升 `BL-020`、进入 Phase 6 收尾或结束本轮之间选择。
- Decision：用户选择 A，提升 `BL-020` 的首批测试工作为 TASK-P1-011 / TS-P1-011，优先覆盖无外部服务依赖且当前无包级测试的 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go`。
- Alternatives：一次性补齐所有无测试 `pkg/*` 包；进入 Phase 6 收尾；结束本轮。
- Reason：`BL-020` 范围较大，首批选择轻依赖包可以降低测试实现和修复风险，避免一次性扩大到 Redis、HTTP server、storage 或 executor 生命周期测试。
- Consequences：下一次“下一步”只能执行 TS-P1-011；`pkg/cache`、`pkg/executor`、`pkg/httpserver`、`pkg/storage` 等剩余测试缺口需要后续批次或收尾前再次确认。
- Related Tasks：TASK-NEXT-SCOPE-003、TASK-P1-011

### DEC-014：提升 `BL-020` 第二批 `pkg/*` 行为测试

- Date：2026-05-25
- Status：ACCEPTED
- Context：TASK-P1-011 完成后，当前状态进入 `PENDING_USER_CONFIRMATION`，要求在继续 `BL-020` 后续包、进入 Phase 6 收尾或结束本轮之间选择。
- Decision：用户发送“下一步”，按确认选项 A 提升 `BL-020` 第二批测试工作为 TASK-P1-012 / TS-P1-012，覆盖 `pkg/executor`、`pkg/httpserver`、`pkg/storage`。
- Alternatives：直接进入 Phase 6 收尾；结束本轮；将 `pkg/cache` Redis 路径纳入同一批。
- Reason：第二批三包可通过内存文件系统、停止态 HTTP server 和本地协程池测试完成，不需要 Redis、数据库或外部服务；`pkg/cache` 需要单独隔离策略。
- Consequences：TASK-P1-012 完成后，`pkg/cache` 等剩余 `BL-020` 范围仍需用户再次确认才能继续。
- Related Tasks：TASK-NEXT-SCOPE-004、TASK-P1-012

### DEC-015：提升 `BL-020` 第三批 `pkg/cache` 隔离行为测试

- Date：2026-05-25
- Status：ACCEPTED
- Context：TASK-P1-012 完成后，当前状态进入 `PENDING_USER_CONFIRMATION`，要求在继续 `BL-020` 剩余范围、进入 Phase 6 收尾或结束本轮之间选择。
- Decision：用户回复 `A`，提升 `BL-020` 第三批测试工作为 TASK-P1-013 / TS-P1-013，限定覆盖 `pkg/cache`。
- Alternatives：直接进入 Phase 6 收尾；结束本轮；把 `pkg/utils` 内部支撑测试并入同一批。
- Reason：`pkg/cache` 是公共基础设施 API，Redis 依赖路径需要隔离测试；`pkg/utils` 已分类为内部支撑工具包，不属于公共包补测范围，若补测应单独确认。
- Consequences：允许引入纯测试用进程内 Redis 依赖以避免真实 Redis 服务；TASK-P1-013 完成后，公共 `pkg/*` 行为测试覆盖告一段落。
- Related Tasks：TASK-NEXT-SCOPE-005、TASK-P1-013

### DEC-016：提升 `BL-023` `pkg/utils` 内部支撑测试

- Date：2026-05-25
- Status：ACCEPTED
- Context：TASK-P1-013 完成后，当前状态进入 `PENDING_USER_CONFIRMATION`，要求在进入 Phase 6 收尾、提升内部支撑测试或结束本轮之间选择。
- Decision：用户回复 `b`，提升 `BL-023` 为 TASK-P1-014 / TS-P1-014，限定覆盖 `pkg/utils` 内部支撑工具最小行为测试。
- Alternatives：直接进入 Phase 6 收尾；结束本轮；扩大到 `internal/app`、middleware 或集成测试。
- Reason：`pkg/utils` 已分类为内部支撑工具包，虽然不属于公共 `pkg/*` 补测范围，但其 ID、地址、端口、设备 ID 和 i18n helper 被多处支撑路径使用，补最小测试可以降低后续维护风险。
- Consequences：本切片只允许修改 `pkg/utils` 测试和必要状态文档，不改变 `pkg/utils` 公共 API 或默认 Snowflake panic 策略。
- Related Tasks：TASK-NEXT-SCOPE-006、TASK-P1-014

### DEC-017：提升 `BL-002` router/middleware/demo HTTP 集成测试

- Date：2026-05-26
- Status：ACCEPTED
- Context：TASK-P1-014 完成后，当前状态进入 `PENDING_USER_CONFIRMATION`，要求在进入 Phase 6 收尾、提升 app/router/middleware 等集成测试或结束本轮之间选择。
- Decision：用户回复 `b`，提升 `BL-002` 的 router/middleware/demo HTTP 集成部分为 TASK-P1-015 / TS-P1-015，限定用 `httptest` 覆盖 demo Todo HTTP 路由、handler/service/repository 集成，以及 TraceID、CORS、Recovery 中间件链路。
- Alternatives：直接进入 Phase 6 收尾；结束本轮；一次性扩大到完整 app 装配、reload/config 集成测试。
- Reason：router/middleware/demo HTTP 链路是现有可隔离验证的高价值缺口，能在不启动真实 HTTP server、不引入生产配置和外部服务的前提下降低后续改动风险。
- Consequences：`BL-002` 进入部分完成状态；app 装配、reload/config 等剩余集成路径仍需后续用户确认后才能提升为新任务。
- Related Tasks：TASK-NEXT-SCOPE-007、TASK-P1-015

### DEC-018：进入 Phase 6 收尾

- Date：2026-05-26
- Status：ACCEPTED
- Context：TASK-P1-015 完成后，当前状态进入 `PENDING_USER_CONFIRMATION`，要求在进入 Phase 6 收尾、继续 app 装配/reload/config 等剩余集成测试或结束本轮之间选择；用户最新回复 `a`。
- Decision：接受用户选择 A，进入 Phase 6 收尾与交接，限定为项目状态文档、验证记录、变更记录、风险/Backlog 和交接说明更新。
- Alternatives：继续 app 装配、reload/config 等剩余集成测试；结束本轮但不做 Phase 6 收尾。
- Reason：当前 P1 已完成多轮受控补测与边界收拢，选择 A 符合当前确认项，能先冻结本轮成果并留下可恢复交接。
- Consequences：本收尾切片不新增 Go 代码或测试；app 装配、reload/config 等剩余集成路径继续保留为 Backlog/风险，后续必须重新确认后才能提升。
- Related Tasks：TASK-NEXT-SCOPE-008、TASK-PHASE6-001

### DEC-019：提升 `BL-002` 剩余 app 装配/reload/config 集成测试

- Date：2026-05-26
- Status：ACCEPTED
- Context：Phase 6 收尾后，用户明确要求实施 TASK-P1-016 计划，覆盖 `BL-002` 中尚未完成的 app 装配、配置变更 hook 与 reload/config 剩余路径。
- Decision：将该剩余范围提升为 TASK-P1-016 / TS-P1-016，限定新增测试文件 `internal/app/app_integration_test.go` 与 `internal/app/reloadapp/reload_test.go`，不新增功能、不启动真实 HTTP server、不依赖外部服务。
- Alternatives：保持 Phase 6 收尾完成状态并继续 defer；扩大到端到端 server 启动或生产迁移验证；将 reload 与 app 装配拆成两个切片。
- Reason：用户已提供具体计划和测试边界，范围清晰且能通过临时 SQLite、临时配置和 fake 组件隔离验证。
- Consequences：允许在 `internal/app/**` 测试范围内补齐集成测试；若测试暴露超出当前范围的缺陷，只能记录为 issue 或单独提升。
- Related Tasks：TASK-P1-016

### DEC-020：提升 `BL-006` 第一阶段包 README 中文化

- Date：2026-05-26
- Status：ACCEPTED
- Context：TASK-P1-017 前当前状态为 `COMPLETED / NONE`，用户在后续方向列表中选择 `a`，对应“包 README 分阶段中文化”；该方向与 `REQ-OPT-P1-005`、`BL-006`、`RISK-005` 对齐。
- Decision：将 `BL-006` 第一阶段提升为 TASK-P1-017 / TS-P1-017，限定为 `pkg/*/README.md` 主要读者文本中文化，并同步需求、架构、模块和状态文档。
- Alternatives：继续保持 `BL-006` defer；一次性中文化全部历史文档；转向 auth/rbac、生产迁移、CI/CD 或部署。
- Reason：包 README 中仍存在英文标题、插件 README 英文主体和过期测试风险描述；第一阶段限定在 `pkg/*/README.md` 可降低范围扩张风险。
- Consequences：允许修改包 README 和必要状态/架构文档；不修改 Go 代码、依赖、配置 schema、HTTP 路由或数据库 schema。历史文档更大范围中文化仍需单独确认。
- Related Tasks：TASK-P1-017

### DEC-021：提升 CI 质量门禁与部署说明首切片

- Date：2026-05-26
- Status：ACCEPTED_WITH_RISK
- Context：当前状态为 `NONE / COMPLETED`，用户在后续方向中选择 D，对应 CI/CD 与部署方向；该方向与 `REQ-OPT-P2-003`、`BL-007` 和 `BL-008` 对齐。
- Decision：将首切片限定为 TASK-P2-001 / TS-P2-001：新增 GitHub Actions CI 质量门禁和手动部署说明，不实现真实 CD、不推送镜像、不连接服务器、不使用 secrets。
- Alternatives：继续保持 defer；直接实现 Docker/镜像发布/云部署；先只写部署文档不加 CI。
- Reason：CI 质量门禁和部署说明能降低发布前回归风险，且不触碰生产环境或密钥；真实 CD 需要更多环境、权限和回滚策略确认。
- Consequences：允许新增 `.github/workflows/ci.yml`、`docs/deployment.md` 和 README 入口；真实 CD、镜像发布和远程部署自动化进入 Backlog。
- Related Tasks：TASK-P2-001

### DEC-022：进入真实 CD / 镜像发布 / 远程部署自动化范围确认

- Date：2026-05-26
- Status：ACCEPTED_WITH_RISK
- Context：TASK-P2-001 完成后，用户在后续方向中选择 C，对应 `BL-024` 真实 CD、镜像发布或远程部署自动化；随后用户补充使用远程部署。
- Decision：不直接实现真实 CD；先完成安全的远程部署 `.env` 风格模板 TASK-P2-002 / TS-P2-002。真实 workflow、镜像发布和远程连接仍需后续单独确认。
- Alternatives：继续保持 `BL-024` defer；只实现 staging/manual dry-run；直接实现生产部署 workflow。
- Reason：真实 CD 涉及远程环境、密钥、镜像仓库和发布权限，缺少这些输入时实现 workflow 会造成错误发布或凭据暴露风险。
- Consequences：提交 `deploy.sh` / `script/install.sh` 显式参数契约，并删除旧本地部署 env 文件依赖；在 TASK-P2-003 确认前不得提交真实服务器值、密钥或自动部署 workflow。后续用户已确认并完成手动 staging workflow。
- Related Tasks：TASK-NEXT-SCOPE-010、TASK-P2-002

### DEC-023：实现手动 staging 远程部署 workflow

- Date：2026-05-26
- Status：ACCEPTED_WITH_RISK
- Context：TASK-P2-002 完成后，用户明确回复“确认实现远程部署 workflow”。
- Decision：实现 TASK-P2-003 / TS-P2-003，新增 `.github/workflows/deploy-remote.yml`，采用 `workflow_dispatch` 手动触发、`confirm=deploy` 确认词、staging-only 环境、GitHub Secrets 注入 显式部署参数 和 SSH key，并在远程主机执行 Docker Compose pull/up 与 health/ready 检查。
- Alternatives：继续保持 workflow defer；直接实现 production workflow；同时实现 Dockerfile 和镜像发布。
- Reason：手动 staging workflow 能落地远程部署路径，同时避免自动生产发布、真实密钥入库和镜像发布范围膨胀。
- Consequences：允许提交 workflow 和部署说明；不在本会话触发 workflow、不连接远程服务器、不推送镜像、不写真实密钥。Dockerfile、镜像发布、production 和生产迁移框架仍需单独确认。
- Related Tasks：TASK-P2-003

### DEC-024：提升 Linux Docker production 部署制品

- Date：2026-05-26
- Status：ACCEPTED_WITH_RISK
- Context：TASK-P2-003 完成后，用户明确要求“开始，linux、docker、production -> 部署”。
- Decision：将 `BL-024` 中 production Docker 远程部署剩余范围提升为 TASK-P2-004 / TS-P2-004；新增 Dockerfile、production Compose 示例、production 配置样例，并把远程部署 workflow 扩展为手动 staging/production 环境选择。
- Alternatives：保持 production defer；只补 Dockerfile 不开放 production workflow；直接执行真实 production 部署。
- Reason：用户已确认 production 部署方向，但真实生产运行涉及远程服务器、密钥、镜像标签和回滚风险，因此本切片只实现可提交制品和受控手动闸门。
- Consequences：production workflow 必须仍为 `workflow_dispatch`，并要求 `deploy-production` 确认词和 GitHub Environment `production`；本会话不触发 workflow、不连接服务器、不推送镜像、不执行生产迁移。
- Related Tasks：TASK-P2-004

### DEC-025：统一 `deploy.sh` 显式参数部署入口

- Date：2026-05-26
- Status：ACCEPTED_WITH_RISK
- Context：用户修正“环境变量在部署脚本上动态配置”，并强调 Windows 本机不应直接执行 Linux Docker 部署，而应通过远程 Linux 主机部署。
- Decision：在 TASK-P2-004 范围内新增根 `deploy.sh` 和 `script/install.sh`，由远程 Linux 主机按显式参数注入运行环境，再执行 Docker build 或 pull、Compose up 和 health/ready 检查。
- Alternatives：继续要求用户手动维护本地部署 env 文件；只依赖 GitHub Actions workflow；在 Windows 本机运行 Docker 部署命令。
- Reason：显式参数入口可以统一 clone 后执行和直接下载执行两种流程，删除旧本地部署 env 文件依赖；通过 SSH 到 Linux 执行也更符合目标运行环境。
- Consequences：脚本只写入非密钥部署变量，真实数据库密码、Redis 密码、SSH key、registry token 和生产迁移仍不由脚本生成或打印；本会话不执行脚本、不连接远程服务器。
- Related Tasks：TASK-P2-004

### DEC-026：提升支持钩子的插件运行时与 IAM 公共接口

- Date：2026-05-27
- Status：ACCEPTED_WITH_RISK
- Context：用户要求实现 `dev.tmp/new-pllugin.md` 设计，并明确说明该路径为笔误，实际设计文件是 `dev.tmp/new-plugin.md`；同时要求 TASK-P2-004 不能因 Docker 环境缺失而标记完成。
- Decision：将新主线限定为 hook-aware `pkg/plugin` runtime、HTTP 远程插件传输和独立 `pkg/iam` 公共接口，并由 `internal/app` 负责配置装配、IAM hook 绑定、reload 和 lifecycle。
- Alternatives：先关闭 TASK-P2-004；直接实现完整业务登录/RBAC；引入 OPA/Casbin；实现 Go `.so` 插件、插件发现或 RPC/WS 传输。
- Reason：该设计延续 `pkg/plugin` 被动注册边界，又能为后续插件与权限能力建立公共基础设施；风险在于 IAM/插件容易被误解为完整生产权限系统，因此必须明确非目标并保留 Docker 验证阻塞。
- Consequences：`pkg/plugin` 不导入 `pkg/iam`、日志、配置或 `internal/*`；`pkg/iam` 不导入 `pkg/plugin`；配置创建的插件仅限 HTTP adapter；本地插件继续由代码显式注册；JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS、生产部署、镜像发布和密钥管理仍需单独确认。
- Related Tasks：TASK-P2-005、TASK-P2-006、TASK-P2-007、TASK-P2-008、TASK-P2-009、TASK-P2-010
