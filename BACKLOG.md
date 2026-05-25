# BACKLOG.md

## Backlog 状态

- Project：go-scaffold
- Last Updated：2026-05-25
- Rule：未被用户确认并拆成任务/时间切片的事项，不属于当前实现范围。

## Backlog 项

| ID | 标题 | 来源 | 优先级 | 延后原因 | 状态 |
|---|---|---|---|---|---|
| BL-001 | 处理 JWT/auth 示例与 README 范围不一致 | `.env.example`、README | P1 | 需要确认 auth/JWT 范围 | [DEFERRED] |
| BL-002 | 增加 app/router/demo 集成测试 | 测试风险分析 | P1 | 需要确认 P0 测试矩阵 | [DEFERRED] |
| BL-003 | 定义 `pkg/*` API 兼容策略 | 包边界风险 | P1 | 需要架构决策 | [DEFERRED] |
| BL-004 | 统一 AutoMigrate、initdb、SQL 脚本迁移策略 | 迁移边界风险 | P1 | 需要架构决策 | [DEFERRED] |
| BL-005 | 明确 `pkg/sqlgen` 未实现能力的边界 | TODO/unsupported 风险 | P1 | 需要包 API 策略 | [DEFERRED] |
| BL-006 | 分阶段中文化包 README | 中文项目要求 | P1 | 需要确认中文化范围 | [DEFERRED] |
| BL-007 | 增加 CI 质量门禁 | 质量工程 | P2 | 超出启动确认阶段 | [DEFERRED] |
| BL-008 | 增加部署说明 | 发布工程 | P2 | 超出当前阶段 | [DEFERRED] |
| BL-009 | 实现 auth/rbac 模块 | 未来功能 | P2 | README 当前说明暂不实现 | [DEFERRED] |
| BL-010 | 增加脚手架生成器 | 产品化方向 | P2 | 需要确认框架化抽取路线 | [DEFERRED] |
| BL-011 | 增加性能基准测试 | 性能质量 | P2 | 需要先稳定功能边界 | [DEFERRED] |
| BL-012 | 增加多租户支持 | 产品架构 | P2 | 未确认且范围较大 | [DEFERRED] |
| BL-013 | 增加插件系统 rpc adapter | 插件扩展 | P2 | v1 仅支持 local/http | [DEFERRED] |
| BL-014 | 增加插件系统 ws adapter | 插件扩展 | P2 | v1 仅支持 local/http | [DEFERRED] |
| BL-015 | 增加插件发现机制 | 插件运行时增强 | P2 | 需要单独提升为任务 | [DEFERRED] |
| BL-016 | 增加 local/http 插件示例 | 插件 v1 后续文档 | P2 | 当前主线已切回项目优化 | [DEFERRED] |
| BL-017 | 统一配置环境变量命名策略 | `MODULES.md` BC-001 | P1 | 需要先生成测试矩阵和任务切片 | [DEFERRED] |
| BL-018 | 修复 `manager.copyConfig` 字段覆盖 | `MODULES.md` BC-002 | P1 | 已提升为 TASK-P1-001 并完成 | [COMPLETED] |
| BL-019 | 处理 `cmd/server tests` 命令语义 | `MODULES.md` BC-004 | P1 | 需要确认命令保留、重命名或改造 | [DEFERRED] |
| BL-020 | 为无测试的公共 `pkg/*` 包补最小行为测试 | `MODULES.md` pkg 分类草案 | P1 | 需要测试矩阵确认 | [DEFERRED] |
| BL-021 | 明确 `types/result` 是否属于 HTTP 契约 | `MODULES.md` types 边界 | P1 | 需要公共契约整理 | [DEFERRED] |

## 提升规则

- [CONFIRMED] Backlog 项只有在用户确认后才能提升。
- [CONFIRMED] 提升后必须映射到需求、架构、任务和时间切片。
- [CONFIRMED] 未提升事项不得顺手实现。
