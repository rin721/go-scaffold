# 需求澄清模板

## 1. 澄清目标

- [CONFIRMED] 本模板用于确认 `go-scaffold` 的项目优化需求。
- [CONFIRMED] 当前阶段只澄清需求、边界、风险和验收标准，不写 Go 代码。
- [INFERRED] 需求确认完成后，才能生成正式优化路线、架构约束、任务拆分和时间切片。

## 2. 已确认事实

- [CONFIRMED] `go-scaffold` 是已有 Go 脚手架项目。
- [CONFIRMED] 当前保留基础设施启动链路和 demo CRUD 示例。
- [CONFIRMED] README 明确当前暂不实现 auth/rbac。
- [CONFIRMED] `go test ./... -count=1` 当前通过。
- [CONFIRMED] 插件系统 v1 local/http 已作为历史工作完成，但不是本轮当前主线。

## 3. 必须确认的问题

| ID | 问题 | 选项 | 推荐默认 | 影响 | 状态 |
|---|---|---|---|---|---|
| REQ-Q001 | 优化路线 | 保守治理优先 / 模块化重构 / 框架化抽取 | 保守治理优先 | 决定是否先改文档还是进入大规模架构调整 | [NEEDS_CONFIRMATION] |
| REQ-Q002 | `pkg/*` 定位 | 公共复用库 / 项目内部支撑 / 混合策略 | 混合策略 | 决定 API 兼容、README、测试和重构边界 | [NEEDS_CONFIRMATION] |
| REQ-Q003 | demo 模块定位 | 长期标准示例 / 临时占位 / 后续移除 | 长期标准示例 | 决定 demo 是否成为新模块模板 | [NEEDS_CONFIRMATION] |
| REQ-Q004 | 迁移策略 | 开发态 `AutoMigrate` / SQL 脚本 / dev-prod 分层策略 | dev-prod 分层策略 | 决定数据库初始化、测试和生产约束 | [NEEDS_CONFIRMATION] |
| REQ-Q005 | 中文化范围 | 只中文化根文档 / 根文档加模板 / 全量历史文档与包 README 分阶段中文化 | 根文档加模板优先 | 决定文档工作量和验收口径 | [NEEDS_CONFIRMATION] |

## 4. 可默认但应记录的要求

- [INFERRED] 先统一文档主线、状态文件和设计边界，再进入代码优化。
- [INFERRED] 保持当前依赖方向：`cmd -> internal/app -> internal/transport + internal/modules -> pkg`。
- [INFERRED] 所有新想法先进入 Backlog，只有被确认并拆成时间切片后才能实现。
- [INFERRED] 后续代码优化必须至少保留 `go test ./... -count=1` 通过。

## 5. 暂缓需求

- [DEFERRED] auth/rbac 实现。
- [DEFERRED] 插件系统 rpc/ws/discovery 扩展。
- [DEFERRED] 部署流水线和生产发布说明。
- [DEFERRED] 性能压测和基准测试体系。
- [DEFERRED] 多租户、框架市场、脚手架生成器。
- [DEFERRED] 包 README 的全量中文化，除非用户确认提升优先级。

## 6. P0 需求草案

| ID | 需求 | 验收信号 | 状态 |
|---|---|---|---|
| REQ-P0-001 | 重新启动全项目优化主线 | `STATUS.md` 当前合法任务指向项目优化确认 | [CONFIRMED] |
| REQ-P0-002 | 生成中文启动模板 | `docs/templates/*` 六个模板均为中文 | [CONFIRMED] |
| REQ-P0-003 | 记录当前优缺点和风险 | 启动模板、风险模板和状态文件列出事实标签 | [CONFIRMED] |
| REQ-P0-004 | 阻止未确认代码实现 | 当前状态为 `PENDING_USER_CONFIRMATION` | [CONFIRMED] |
| REQ-P0-005 | 保留插件系统历史但不继续扩展 | 插件相关任务进入历史或 Backlog | [CONFIRMED] |

## 7. 后续输出

用户确认后，应生成或更新：

- `REQUIREMENTS.md`
- `ARCHITECTURE.md`
- `ACCEPTANCE.md`
- `RISK_REGISTER.md`
- `BACKLOG.md`
- `TASKS.md`
- `TIME_SLICES.md`
- `STATUS.md`
- `AGENT_HANDOFF.md`
