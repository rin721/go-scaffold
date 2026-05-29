# 已知缺口

本文记录工程文档整理期间发现的已知缺口。它用于让未来工作可见，同时避免把尚未完成的能力写成已经完成。

## 实现漂移

| 缺口 | 影响 | 建议验证 |
| --- | --- | --- |
| middleware 使用 `traceId`，result helper 读取 `trace_id` | 普通错误响应可能缺少 trace ID | 增加真实 middleware + result helper 集成测试，再统一键名 |
| HTTP server 启动后的异步 serve 错误没有通过 app interface 暴露 | 后续 server 错误可能只被日志记录，不能传回 `cmd/main` | 复核 `pkg/httpserver` interface 和 lifecycle 行为 |
| DB CLI 聚焦 demo，而 user schema 在 server 启动时应用 | 运维者可能误以为一个 CLI 覆盖所有 schema | 确认这是设计选择，或扩展 DB CLI |

## 文档债

| 缺口 | 影响 | 建议动作 |
| --- | --- | --- |
| 部分 package README 含有过时 API 示例 | 读者可能复制不存在的方法或错误签名 | 中心文档稳定后，统一修正 package README |
| 部署 workflow 变量混用了 `RIN_APP_*` 应用变量和未加前缀 GitHub Variables | 运维配置容易混淆 | 在部署专项任务中对齐 workflow 和文档变量 |
| `docs/ai` 对 slice 011 的索引可能不完整 | 未来 Agent 可能缺少 task/evidence 上下文 | 通过已确认的运行态修复任务补齐 artifact |

## 生产就绪缺口

| 缺口 | 影响 |
| --- | --- |
| 没有 refresh token、session revoke、密码重置流程 | 当前 auth 是本地脚手架级能力，不是完整账户生命周期 |
| 没有完整迁移框架 | 生成 schema 是有用基线，不等于生产迁移体系 |
| 远程 blog 插件使用内存存储 | 只适合作为演示 |
| 发布/回滚证据尚未形成稳定 v1 流程 | v1 发布仍需验收工作 |
