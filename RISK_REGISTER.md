# RISK_REGISTER.md

## 风险登记状态

- Project：go-scaffold
- Phase：模块边界清单
- Status：IN_PROGRESS
- Last Updated：2026-05-25

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
- Mitigation：当前合法任务已切换到模块边界清单，后续仍需持续更新状态文件。
- Owner：Agent
- Status：[RISK]
- Blocking：Yes，阻塞“下一步”稳定执行。

### RISK-003：迁移策略冲突

- Type：Architecture
- Severity：High
- Probability：Medium
- Impact：开发态、测试态和生产态数据库初始化职责混乱。
- Trigger：demo `AutoMigrate`、`initdb` 命令、SQL 脚本并存但无统一策略。
- Mitigation：架构确认阶段明确迁移策略。
- Owner：User/Agent
- Status：[RISK]
- Blocking：Yes，阻塞数据库相关优化。

### RISK-004：`pkg/*` API 兼容不明

- Type：Architecture
- Severity：High
- Probability：Medium
- Impact：包级重构可能破坏未来或外部使用者。
- Trigger：未确认 `pkg/*` 是公共复用库、内部支撑或混合策略。
- Mitigation：在 `DECISIONS.md` 中记录 API 策略后再重构。
- Owner：User/Agent
- Status：[RISK]
- Blocking：Yes，阻塞包级重构。

### RISK-005：包 README 中英混杂

- Type：Documentation
- Severity：Medium
- Probability：High
- Impact：中文项目体验不一致，交接成本上升。
- Trigger：根文档中文化后，包 README 仍保留英文或混合风格。
- Mitigation：先中文化根文档和模板；包 README 分阶段处理。
- Owner：Agent
- Status：[RISK]
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
- Mitigation：后续确认删除、保留占位或提升为正式需求。
- Owner：User/Agent
- Status：[RISK]
- Blocking：No，但阻塞 auth 相关文档或实现。

### RISK-008：测试覆盖不足

- Type：Testing
- Severity：High
- Probability：Medium
- Impact：后续代码优化可能回归 app 启动、路由、demo CRUD、配置热更新或迁移。
- Trigger：多个关键路径当前没有测试文件。
- Mitigation：代码优化前确认 P0 测试矩阵。
- Owner：Agent
- Status：[RISK]
- Blocking：No，但必须在代码优化前处理。

## 决策状态

| ID | 决策 | 阻塞内容 | 状态 |
|---|---|---|---|
| RD-001 | 确认优化路线 | 已确认治理优先 | [CONFIRMED] |
| RD-002 | 确认 `pkg/*` API 策略 | 已确认混合策略 | [CONFIRMED] |
| RD-003 | 确认 demo 模块定位 | 已确认长期标准示例 | [CONFIRMED] |
| RD-004 | 确认迁移策略 | 已确认 dev-prod 分层 | [CONFIRMED] |
| RD-005 | 确认中文化范围 | 已确认根文档和模板优先 | [CONFIRMED] |
| RD-006 | 确认 auth/JWT 范围 | 已确认延后处理 | [CONFIRMED] |
