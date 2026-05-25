# ARCHITECTURE.md

## 架构状态

- 项目：go-scaffold
- 当前焦点：项目治理与优化路线 v1
- 状态：COMPLETED
- 最后更新：2026-05-25

## 已确认架构原则

- [CONFIRMED] 采用治理优先路线：先稳定文档、边界、任务和测试矩阵，再改代码。
- [CONFIRMED] 保持当前依赖方向：`cmd -> internal/app -> internal/transport + internal/modules -> pkg`。
- [CONFIRMED] `internal/app` 是组合根和生命周期边界。
- [CONFIRMED] `internal/modules/demo` 暂定为长期标准示例。
- [CONFIRMED] `pkg/*` 采用混合策略，后续必须逐包标注公共 API 或内部支撑定位。
- [CONFIRMED] 迁移采用 dev-prod 分层策略，后续需要把 `AutoMigrate`、`initdb` 和 SQL 脚本职责写清。

## 目标依赖方向

```text
cmd/server
  -> internal/app
      -> internal/config
      -> internal/modules/*
      -> internal/transport/http
      -> pkg/*
types/*
```

## 已确认边界

| 边界 | 结论 | 后续动作 | 状态 |
|---|---|---|---|
| `cmd/server` | 进程入口、CLI 命令、信号处理 | 检查是否有业务逻辑外溢 | [CONFIRMED] |
| `internal/app` | 组合根、生命周期、跨层装配 | 生成装配链路清单 | [CONFIRMED] |
| `internal/transport/http` | HTTP router、middleware、health/ready、API 注册 | 补充 smoke/integration 测试计划 | [CONFIRMED] |
| `internal/modules/demo` | 长期标准示例 | 生成 demo 分层验收和测试路线 | [CONFIRMED] |
| `pkg/*` | 混合策略 | 逐包分类：公共 API / 内部支撑 / 待确认 | [CONFIRMED] |
| 数据库迁移 | dev-prod 分层 | 明确 `AutoMigrate`、`initdb`、SQL 脚本职责 | [CONFIRMED] |
| 插件系统 | v1 local/http 保留为历史已完成能力 | rpc/ws/discovery 留在 Backlog | [CONFIRMED] |
| auth/JWT | 当前不实现，示例存在范围漂移 | 后续决定删除、保留占位或提升需求 | [CONFIRMED] |

## 需要详细分析的模块

| 模块 | 主要问题 | 下一步 |
|---|---|---|
| `internal/app` | 装配、reload、mode、lifecycle 边界需形成清单 | TASK-OPT-003 |
| `internal/config` | 环境覆盖、热更新和默认值职责需统一说明 | TASK-OPT-003 |
| `internal/transport/http` | health/ready 和 demo 路由缺少集成测试 | TASK-OPT-003 |
| `internal/modules/demo` | 示例职责与生产约束需分离 | TASK-OPT-003 |
| `pkg/*` | 公共/内部定位未逐包标记 | TASK-OPT-003 |
| `types/*` | 错误码和响应类型是否属于公共契约需确认 | TASK-OPT-003 |

## 代码变更门禁

- [CONFIRMED] 未完成模块边界清单前，不进入代码优化。
- [CONFIRMED] 未完成测试矩阵前，不修改 app/router/demo/config/migration 核心路径。
- [CONFIRMED] 未确认包 API 分类前，不做 `pkg/*` 破坏性重构。

## 下一架构任务

- [CONFIRMED] TASK-OPT-003：生成模块边界清单和优化路线明细。
