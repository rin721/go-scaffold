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
