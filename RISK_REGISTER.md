# RISK_REGISTER.md

## 风险登记状态

- Project：go-scaffold
- Phase：P2 Linux Docker production 部署制品与远程 Linux 动态 env 脚本
- Status：PENDING_VERIFICATION
- Last Updated：2026-05-26

## 风险列表

### RISK-001：范围膨胀成全量重写

- Type：Scope
- Severity：High
- Probability：Medium
- Impact：优化工作不可收敛，测试成本上升，当前可运行链路可能被破坏。
- Trigger：未确认需求和架构边界就开始重构。
- Mitigation：默认采用治理优先；未确认内容进入 Backlog。
- Owner：User/Agent
- Status：[RISK]
- Blocking：Yes，阻塞代码实现。

### RISK-002：文档状态漂移

- Type：Documentation
- Severity：High
- Probability：High
- Impact：后续 Agent 可能继续执行插件系统扩展，而不是项目优化主线。
- Trigger：`STATUS.md`、`TASKS.md`、`TIME_SLICES.md` 未更新。
- Mitigation：当前合法任务已随每个切片更新；TASK-INFRA-003 已修复 TASK-P1-016/017 后背景文档旧状态漂移；后续仍需持续同步 `STATUS.md`、`TASKS.md`、`TIME_SLICES.md`、`TEST_REPORT.md` 和 `AGENT_HANDOFF.md`。
- Owner：Agent
- Status：[RISK]
- Blocking：Yes，阻塞“下一步”稳定执行。

### RISK-003：迁移策略冲突

- Type：Architecture
- Severity：High
- Probability：Medium
- Impact：开发态、测试态和生产态数据库初始化职责混乱。
- Trigger：demo `AutoMigrate`、`initdb` 命令、SQL 脚本并存但无统一策略。
- Mitigation：TASK-P1-005 已将 demo 迁移触发策略显式化：server-start/initdb 允许 demo `AutoMigrate`，reload 跳过隐式 schema 变更；生产迁移框架仍延后。
- Owner：User/Agent
- Status：[CONFIRMED] 触发边界已收拢；生产迁移框架仍延后
- Blocking：No。

### RISK-004：`pkg/*` API 兼容不明

- Type：Architecture
- Severity：High
- Probability：Medium
- Impact：包级重构可能破坏未来或外部使用者。
- Trigger：未确认 `pkg/*` 是公共复用库、内部支撑或混合策略。
- Mitigation：TASK-P1-007 已逐包标注公共基础设施 API、公共工具 API 或内部支撑工具包；任何 `pkg/*` 破坏性重构仍需单独任务和确认。
- Owner：User/Agent
- Status：[CONFIRMED] API 分类已完成；破坏性重构仍受门禁约束
- Blocking：No。

### RISK-005：包 README 中英混杂

- Type：Documentation
- Severity：Medium
- Probability：High
- Impact：中文项目体验不一致，交接成本上升。
- Trigger：根文档中文化后，包 README 仍保留英文或混合风格。
- Mitigation：根文档和模板已中文化；第一阶段 `pkg/*/README.md` 已由 TASK-P1-017 完成，历史文档更大范围中文化仍需单独确认。
- Owner：Agent
- Status：[CONFIRMED] 第一阶段包 README 中文化已完成
- Blocking：No。

### RISK-006：插件系统历史主线干扰当前目标

- Type：Process
- Severity：Medium
- Probability：Medium
- Impact：后续工作可能误入 rpc/ws/discovery 扩展。
- Trigger：插件系统 v1 closeout 继续作为当前合法任务。
- Mitigation：插件系统历史保留，扩展项进入 Backlog。
- Owner：Agent
- Status：[RISK]
- Blocking：Yes，阻塞当前主线统一。

### RISK-007：JWT/auth 示例暗示未实现能力

- Type：Scope/Security
- Severity：Medium
- Probability：Medium
- Impact：用户可能误认为 auth/rbac 已支持，带来安全误解。
- Trigger：`.env.example` 包含 JWT 示例，而 README 说暂不实现 auth/rbac。
- Mitigation：TASK-P1-002 已从 `.env.example` 移除 JWT 示例；auth/rbac 继续作为延后需求处理。
- Owner：User/Agent
- Status：[CONFIRMED] 示例风险已修复；auth/rbac 仍延后
- Blocking：No。

### RISK-008：测试覆盖不足

- Type：Testing
- Severity：High
- Probability：Medium
- Impact：后续代码优化可能回归 app 启动、路由、demo CRUD、配置热更新或迁移。
- Trigger：多个关键路径当前没有测试文件。
- Mitigation：`TEST_MATRIX.md` 已生成；TASK-P1-003 已补齐 `internal/transport/http` health/ready smoke test；TASK-P1-004 已补齐 demo Todo service/repository CRUD 测试基线；TASK-P1-005 已补齐 demo 迁移策略测试；TASK-P1-006 已补齐 `cmd/server` tests 命令语义测试；TASK-P1-007 已完成 `pkg/*` API 分类；TASK-P1-008 已补齐 `pkg/sqlgen` unsupported 行为测试；TASK-P1-011 已补齐首批 `pkg/cli`、`pkg/i18n`、`pkg/yaml2go` 行为测试；TASK-P1-012 已补齐第二批 `pkg/executor`、`pkg/httpserver`、`pkg/storage` 行为测试；TASK-P1-013 已补齐 `pkg/cache` 隔离行为测试；TASK-P1-014 已补齐 `pkg/utils` 内部支撑测试；TASK-P1-015 已补齐 router/middleware/demo HTTP 集成测试；TASK-P1-016 已补齐 app 装配、配置变更 hook 与 reload/config 剩余集成测试。
- Owner：Agent
- Status：[CONFIRMED] 当前 P1 已覆盖 app/router/demo/config/reload 与主要 `pkg/*` 最小测试路径；更大范围端到端、生产迁移、CI/CD 与 auth/rbac 仍按独立 Backlog 处理
- Blocking：No，但必须在代码优化前处理。

### RISK-009：P1 执行顺序未确认

- Type：Process
- Severity：Medium
- Probability：Medium
- Impact：如果直接进入代码优化，可能出现先修复后补测试、或任务顺序与用户预期不一致。
- Trigger：生成测试矩阵和任务切片后未确认执行顺序。
- Mitigation：用户再次发送“下一步”后已按推荐默认顺序确认，并已完成 TASK-P1-001。
- Owner：User/Agent
- Status：[CONFIRMED]
- Blocking：No。

### RISK-010：Agent 入口状态与文件系统事实不一致

- Type：Process/Documentation
- Severity：High
- Probability：Medium
- Impact：新 Agent 读取 `CLAUDE.md`、Cursor、Kiro 或 Codex 配置时会引用缺失的 `AGENTS.md`，导致上下文恢复失败。
- Trigger：`TASK-INFRA-001` 记录 `AGENTS.md` 已补齐，但实际文件缺失。
- Mitigation：TASK-INFRA-002 已新增 `AGENTS.md`，统一跨工具引用，补齐 skills frontmatter 和 `.agents` adapters，并生成状态诊断报告。
- Owner：Agent
- Status：[CONFIRMED] 已修复
- Blocking：No。

### RISK-011：配置环境变量命名策略不一致

- Type：Configuration/Documentation
- Severity：Medium
- Probability：High
- Impact：使用 `.env.example` 的开发者设置 `DB_*` 后数据库配置不生效，造成本地和部署配置漂移。
- Trigger：数据库 override 曾读取 `REI_APP_DB_*`，而 `.env.example` 和其他模块使用未加前缀变量。
- Mitigation：TASK-P1-002 已统一为 `DB_*` 优先，旧 `REI_APP_DB_*` 兼容 fallback；`.env.example` 已同步；配置测试已覆盖。
- Owner：Agent
- Status：[CONFIRMED] 已修复
- Blocking：No。

### RISK-012：CLI tests 命令语义误导

- Type：CLI/Testing
- Severity：Medium
- Probability：High
- Impact：开发者可能误以为 `cmd/server tests` 会执行 Go 测试，但旧实现实际运行 yaml2go 示例转换。
- Trigger：`TestsCommand.Description()` 曾返回 `Run tests`，但 `Execute` 调用 yaml2go converter。
- Mitigation：TASK-P1-006 已将 `tests` 改为真实 Go test 入口，默认 `go test ./...`，并新增命令语义测试。
- Owner：Agent
- Status：[CONFIRMED] 已修复
- Blocking：No。

### RISK-013：`pkg/sqlgen` 未实现能力边界不清

- Type：Documentation/API
- Severity：Medium
- Probability：High
- Impact：使用者可能把 TODO 或未实现能力误认为可用能力，后续 API 兼容和测试验收会变得含糊。
- Trigger：`pkg/sqlgen` 文档、TODO 或导出能力没有显式标注 unsupported 边界。
- Mitigation：TASK-P1-008 已在 `pkg/sqlgen` 允许范围内标注 unsupported 边界；高级查询、批量删除和 DB reverse 均有 `ErrCodeUnsupportedOperation` 测试覆盖。
- Owner：Agent
- Status：[CONFIRMED] 已修复
- Blocking：No。

### RISK-014：`types/*` 契约边界不清

- Type：Architecture/API
- Severity：Medium
- Probability：Medium
- Impact：`types/result` 依赖 Gin 且承担 HTTP 响应契约，若被误认为纯类型包，后续重构或复用会破坏跨层边界；auth/rbac 错误码预留也可能被误解为已实现能力。
- Trigger：`types/*` 当前承载常量、错误码、响应结构和聚合类型，但公共契约与实现范围尚未完整标注。
- Mitigation：TASK-P1-009 已明确 `types/*` 契约边界，补充 `types/result` 和 `types/errors` 最小测试，并运行 `go test ./types/... -count=1` 与全量回归。
- Owner：Agent
- Status：[CONFIRMED] 已处理
- Blocking：No。

### RISK-015：`pkg/plugin` 注册责任反向

- Type：Architecture/API
- Severity：Medium
- Probability：Medium
- Impact：如果 `pkg/plugin` 主动从配置加载并注册插件服务，基础设施包会承担服务装配职责，后续接入真实插件服务时容易形成生命周期和依赖方向混乱。
- Trigger：`Manager.Load(config)` 和 local factory 公共 API 让 manager 主动创建并注册插件实例。
- Mitigation：TASK-P1-010 已收拢为被动注册边界：插件服务或宿主装配层显式创建插件并调用 `Register`；`pkg/plugin` 仅保留 registry/runtime 能力。
- Owner：Agent
- Status：[CONFIRMED] 已处理
- Blocking：No。

### RISK-016：CI/CD 误触发生产动作

- Type：Release/Safety
- Severity：High
- Probability：Medium
- Impact：如果 CI/CD workflow 自动连接生产环境、推送镜像或执行部署，可能造成未授权发布、密钥暴露或生产变更。
- Trigger：在未确认密钥、环境、权限和回滚策略前实现自动 CD。
- Mitigation：TASK-P2-001 已新增只读 CI 质量门禁和手动部署说明；TASK-P2-003 已新增手动 staging 远程部署 workflow；TASK-P2-004 已补 production Docker 制品、手动 production 闸门和远程 Linux 动态 env 部署脚本，Docker build 待具备 Docker 的环境补跑。workflow 和脚本均不在本会话触发远程连接、不推送镜像、不写真实 secrets。
- Owner：User/Agent
- Status：[CONFIRMED] 非生产 CI、手动 staging workflow 和 production 手动闸门均按受控方向推进
- Blocking：No；会阻塞 production 自动化，直到用户单独确认。

### RISK-017：真实 CD 输入不足导致错误发布

- Type：Release/Safety
- Severity：High
- Probability：High
- Impact：在未确认镜像仓库、远程环境、触发策略和 secrets 权限前实现真实 CD，可能导致错误环境发布、凭据暴露、镜像覆盖或不可回滚变更。
- Trigger：用户选择 C 并确认使用远程部署；TASK-P2-002 已提供 `.env.deploy.example`，TASK-P2-003 已新增手动 staging 远程部署 workflow；用户随后确认 Linux/Docker/production 部署制品。
- Mitigation：`.env.deploy.example` 只提供占位变量，真实 `.env.deploy` 已被忽略；workflow 仅手动触发，production 需要 GitHub Environment 和 `deploy-production` 确认词；远程 Linux 脚本只动态写入非密钥部署运行变量，真实数据库/Redis 密码仍应位于远程 `configs/config.yaml` 或密钥系统；本会话不触发 workflow、不连接远程环境、不读取真实 secrets、不部署生产。
- Owner：User/Agent
- Status：[CONFIRMED] env 模板、手动 workflow 和 production Docker 制品方向已确认；镜像发布流水线和真实运行仍受阻
- Blocking：Yes，阻塞镜像发布流水线 / 真实 production 运行 / 生产迁移后续范围。

### RISK-018：production Docker 制品误用为已完成真实生产部署

- Type：Release/Safety
- Severity：High
- Probability：Medium
- Impact：使用者可能把 Dockerfile、Compose 示例或 production workflow 闸门误认为已经完成真实生产上线，从而跳过 GitHub Environment 审批、真实镜像标签确认、远程目录权限、回滚和迁移检查。
- Trigger：用户要求进入 production 部署；仓库新增 production 命名的部署制品。
- Mitigation：所有制品使用占位值和 example 文件；workflow 仅手动触发，production 要求 `deploy-production`；远程 Linux 部署脚本会按参数动态生成 `.env.deploy`，且只写入部署环境、镜像、端口和健康检查地址等非密钥变量；文档明确本会话未触发真实部署、未连接服务器、未推送镜像、未执行迁移。
- Owner：User/Agent
- Status：[RISK] 制品已补齐，Docker build 待验证；不得宣称真实 production 已上线
- Blocking：No；但阻塞把本切片结果宣称为真实 production 已上线。

## 决策状态

| ID | 决策 | 阻塞内容 | 状态 |
|---|---|---|---|
| RD-001 | 确认优化路线 | 已确认治理优先 | [CONFIRMED] |
| RD-002 | 确认 `pkg/*` API 策略 | 已确认混合策略 | [CONFIRMED] |
| RD-003 | 确认 demo 模块定位 | 已确认长期标准示例 | [CONFIRMED] |
| RD-004 | 确认迁移策略 | 已确认 dev-prod 分层 | [CONFIRMED] |
| RD-005 | 确认中文化范围 | 已确认根文档和模板优先 | [CONFIRMED] |
| RD-006 | 确认 auth/JWT 范围 | 已确认延后处理 | [CONFIRMED] |
| RD-007 | 确认 `pkg/plugin` 注册责任 | 用户修正为被动注册边界 | [CONFIRMED] |
| RD-008 | 确认包 README 中文化第一阶段 | `pkg/*/README.md` 已完成第一阶段中文化 | [CONFIRMED] |
| RD-009 | 确认 CI/CD 与部署首切片 | 仅新增 CI 质量门禁和部署说明，不做真实 CD | [CONFIRMED] |
| RD-010 | 确认真实 CD / 镜像发布 / 远程部署自动化边界 | 用户选择 C 并确认远程部署；env 模板和手动 staging workflow 已完成，production 与镜像发布仍需确认 | [CONFIRMED] |
| RD-011 | 确认 Linux Docker production 部署制品 | 用户确认 production 部署方向；本切片只补制品和手动闸门，不执行真实 production | [CONFIRMED] |
