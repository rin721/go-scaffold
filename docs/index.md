# 工程文档入口

这里是 `go-scaffold` 的人类工程文档入口，描述当前磁盘上的 Go 服务脚手架。
`docs/ai` 子树是独立的 AI 运行态，保存任务状态、证据、决策、知识和交接信息。

## 推荐阅读顺序

1. [项目概览](overview/project.md)
2. [目录地图](structure/directory-map.md)
3. [配置说明](environment/configuration.md)
4. [分层架构](architecture/layers.md)
5. [启动流程](runtime/startup-flow.md)
6. [HTTP 流程](runtime/http-flow.md)
7. [配置流](runtime/config-flow.md)
8. [错误流](runtime/error-flow.md)
9. [模块文档](modules/demo.md)
10. [测试矩阵](testing/test-matrix.md)
11. [Docker 与 CI](build/docker-and-ci.md)
12. [部署发布](release/deployment.md)
13. [维护指南](maintenance/maintenance-guide.md)
14. [AI 运行态](ai-agent/runtime-state.md)
15. [已知缺口](backlog/known-gaps.md)

## 文档地图

| 章节 | 用途 |
| --- | --- |
| [overview](overview/project.md) | 项目状态、能力范围、非目标和运行假设 |
| [structure](structure/directory-map.md) | 目录职责、边界和依赖方向 |
| [environment](environment/configuration.md) | 配置文件、环境变量、`.env` 和生产密钥规则 |
| [architecture](architecture/layers.md) | 应用分层和模块装配方式 |
| [runtime](runtime/startup-flow.md) | 启动、HTTP 请求、配置热加载、数据、状态和错误流 |
| [modules](modules/demo.md) | Demo、用户、认证/RBAC 文档 |
| [workflows](workflows/db-cli.md) | DB CLI 等开发和运维工作流 |
| [testing](testing/test-matrix.md) | 测试归属和验证命令 |
| [build](build/docker-and-ci.md) | CI、本地构建、Docker 构建和质量门禁 |
| [release](release/deployment.md) | 生产配置、部署脚本、远程 workflow 和发布检查 |
| [extension](extension/adding-modules.md) | 如何添加模块、配置和 API |
| [maintenance](maintenance/maintenance-guide.md) | 长期维护流程和文档卫生 |
| [ai-agent](ai-agent/runtime-state.md) | `AGENTS.md` 与 `docs/ai` 运行态说明 |
| [backlog](backlog/known-gaps.md) | 已知实现漂移、文档债和验证缺口 |

## 兼容入口

这些旧顶层文档保留为兼容入口：

- [project-overview.md](project-overview.md)
- [usage-guide.md](usage-guide.md)
- [configuration.md](configuration.md)
- [db-cli.md](db-cli.md)
- [deployment.md](deployment.md)
- [maintenance-guide.md](maintenance-guide.md)
- [agent-workflow-guide.md](agent-workflow-guide.md)

新增文档应优先放入上面的结构化目录。
