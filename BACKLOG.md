# BACKLOG.md

## Current Backlog Promotion

- 2026-05-28: User requested moving RBAC config under `configs`.
- The config-seed portion of `BL-026` was promoted to and completed by `TASK-P2-020 / TS-P2-020`.
- `BL-026` remaining hardening stays deferred: production IAM, OPA/Casbin, refresh-token/session revoke, audit logging, password reset, external IAM, production migrations, and real secret management.
- 2026-05-28: User selected `2+4`.
- `BL-026` was partially promoted to and completed by `TASK-P2-019 / TS-P2-019` for auth token secret/TTL configuration only.
- `BL-026` remaining hardening stays deferred: refresh-token/session revoke, audit logging, password reset, external IAM, OPA/Casbin, production migrations, and real secret management.
- `BL-028` remains deferred/queued for a later plugin WS/RPC/heartbeat/persistent-discovery slice.

## Backlog 状态

- Project：go-scaffold
- Last Updated：2026-05-28
- Rule：未被用户确认并拆成任务/时间切片的事项，不属于当前实现范围。

## Backlog 项

| ID | 标题 | 来源 | 优先级 | 延后原因 | 状态 |
|---|---|---|---|---|---|
| BL-001 | 处理 JWT/auth 示例与 README 范围不一致 | `.env.example`、README | P1 | README 已由 TASK-P2-018 对齐本地 user/auth/RBAC 能力；真实生产密钥、refresh token/session revoke 和外部 IAM 仍由 BL-026 跟踪 | [COMPLETED] |
| BL-002 | 增加 app/router/demo 集成测试 | 测试风险分析 | P1 | router/middleware/demo HTTP 集成已在 TASK-P1-015 覆盖；app 装配、配置变更 hook 与 reload/config 剩余范围已在 TASK-P1-016 覆盖 | [COMPLETED] |
| BL-003 | 定义 `pkg/*` API 兼容策略 | 包边界风险 | P1 | 需要架构决策 | [DEFERRED] |
| BL-004 | 统一 AutoMigrate、initdb、SQL 脚本迁移策略 | 迁移边界风险 | P1 | 已由 TASK-P2-014 统一到 sqlgen DB CLI；生产迁移框架另行确认 | [COMPLETED] |
| BL-005 | 明确 `pkg/sqlgen` 未实现能力的边界 | TODO/unsupported 风险 | P1 | 已提升为 TASK-P1-008 并完成 | [COMPLETED] |
| BL-006 | 分阶段中文化包 README | 中文项目要求 | P1 | 第一阶段 `pkg/*/README.md` 已由 TASK-P1-017 完成；历史文档更大范围中文化仍需单独确认 | [COMPLETED] |
| BL-007 | 增加 CI 质量门禁 | 质量工程 | P2 | 已提升为 TASK-P2-001 并完成 | [COMPLETED] |
| BL-008 | 增加部署说明 | 发布工程 | P2 | 已提升为 TASK-P2-001 并完成 | [COMPLETED] |
| BL-009 | 实现 auth/rbac 模块 | 未来功能 | P2 | 已由 TASK-P2-018 完成主服务本地 user/auth/RBAC；生产级 IAM hardening 仍由 BL-026 跟踪 | [COMPLETED] |
| BL-010 | 增加脚手架生成器 | 产品化方向 | P2 | 需要确认框架化抽取路线 | [DEFERRED] |
| BL-011 | 增加性能基准测试 | 性能质量 | P2 | 需要先稳定功能边界 | [DEFERRED] |
| BL-012 | 增加多租户支持 | 产品架构 | P2 | 未确认且范围较大 | [DEFERRED] |
| BL-013 | 增加插件系统 rpc adapter | 插件扩展 | P2 | hook-aware runtime 与 HTTP 远程插件已完成；RPC adapter 仍需单独提升 | [DEFERRED] |
| BL-014 | 增加插件系统 ws adapter | 插件扩展 | P2 | hook-aware runtime 与 HTTP 远程插件已完成；WS adapter 仍需单独提升 | [DEFERRED] |
| BL-015 | 增加插件发现机制 | 插件运行时增强 | P2 | hooks 和 RemoteHook 已完成；插件发现仍需单独提升为任务 | [DEFERRED] |
| BL-016 | 增加 local/http 插件示例 | 插件 v1 后续文档 | P2 | 当前主线已切回项目优化 | [DEFERRED] |
| BL-017 | 统一配置环境变量命名策略 | `MODULES.md` BC-001 | P1 | 需要先生成测试矩阵和任务切片 | [DEFERRED] |
| BL-018 | 修复 `manager.copyConfig` 字段覆盖 | `MODULES.md` BC-002 | P1 | 已提升为 TASK-P1-001 并完成 | [COMPLETED] |
| BL-019 | 处理 `cmd/server tests` 命令语义 | `MODULES.md` BC-004 | P1 | 需要确认命令保留、重命名或改造 | [DEFERRED] |
| BL-020 | 为无测试的公共 `pkg/*` 包补最小行为测试 | `MODULES.md` pkg 分类草案 | P1 | 首批 TASK-P1-011、第二批 TASK-P1-012 和第三批 TASK-P1-013 已完成，公共 `pkg/*` 行为测试已覆盖 | [COMPLETED] |
| BL-021 | 明确 `types/result` 是否属于 HTTP 契约 | `MODULES.md` types 边界 | P1 | 已提升为 TASK-P1-009 并完成 | [COMPLETED] |
| BL-022 | 收拢 `pkg/plugin` 被动注册边界 | 用户架构修正 | P1 | 已提升为 TASK-P1-010 并完成 | [COMPLETED] |
| BL-023 | 为 `pkg/utils` 内部支撑工具补最小测试 | `MODULES.md` pkg 分类草案 | P2 | 已提升为 TASK-P1-014 / TS-P1-014 并完成 | [COMPLETED] |
| BL-024 | 增加真实 CD、镜像发布或远程部署自动化 | CI/CD 首切片后续 | P2 | 显式参数部署入口已由 TASK-P2-002/TASK-P2-004 收敛；手动 staging/production 远程部署 workflow 闸门已补齐；Docker build 已由用户在 Linux 环境验证通过；镜像发布流水线和真实运行仍需单独确认 | [DEFERRED] |
| BL-025 | 统一 Go 文件 gofmt 格式 | CI 格式审计 | P2 | 当前仓库存在历史 gofmt 漂移，批量格式化会触碰大量 Go 文件，需单独确认后执行 | [DEFERRED] |
| BL-026 | 完整 IAM / auth hardening | TASK-P2-008/TASK-P2-018 后续 | P2 | 主服务本地 user/auth/RBAC 已完成；生产级密钥管理、refresh token/session revoke、审计、密码重置、外部 IAM/OPA/Casbin 和生产迁移需单独确认 | [DEFERRED] |
| BL-027 | 定义第一版发布验收清单与剩余开发路线 | 用户纠正当前项目未达第一版发布条件 | P0 | 需要用户确认 v1 的功能范围、质量门禁、真实环境验证、迁移、密钥管理和发布流程后，才能拆分任务/时间切片 | [DEFERRED] |
| BL-028 | 远程插件 WS/RPC 常连接、心跳与持久发现 | TASK-P2-016 非目标 | P2 | 本切片只实现 HTTP 显式注册和 Blog 示例；`ws://.../ws` 仅作为示例配置预留，真实 WS/RPC transport、重连、心跳和持久注册表需单独确认 | [DEFERRED] |

## 提升规则

- [CONFIRMED] Backlog 项只有在用户确认后才能提升。
- [CONFIRMED] 提升后必须映射到需求、架构、任务和时间切片。
- [CONFIRMED] 未提升事项不得顺手实现。
