# 工程文档索引

本目录概述当前磁盘上的 Go 脚手架。`docs/ai` 是独立的 AI 运行时状态树，用于保存任务状态、证据、决策、知识和交接信息。

## 推荐阅读顺序

1. [项目概述](overview/project.md)
2. [目录地图](structure/directory-map.md)
3. [配置说明](environment/configuration.md)
4. [分层架构](architecture/layers.md)
5. [启动流程](runtime/startup-flow.md)
6. [HTTP 流程](runtime/http-flow.md)
7. [配置流程](runtime/config-flow.md)
8. [错误流程](runtime/error-flow.md)
9. [Demo 模块](modules/demo.md)
10. [测试矩阵](testing/test-matrix.md)
11. [Docker 和 CI](build/docker-and-ci.md)
12. [部署说明](release/deployment.md)
13. [维护指南](maintenance/maintenance-guide.md)
14. [AI 运行时状态](ai-agent/runtime-state.md)
15. [已知缺口](backlog/known-gaps.md)

## 文档地图

| 分区 | 内容概述 |
| --- | --- |
| [overview](overview/project.md) | 当前能力、非目标和运行时默认假设 |
| [structure](structure/directory-map.md) | 目录职责和依赖方向 |
| [environment](environment/configuration.md) | 配置文件、环境变量、`.env` 和生产示例 |
| [architecture](architecture/layers.md) | 应用分层和装配方式 |
| [runtime](runtime/startup-flow.md) | 启动、HTTP、配置重载、状态和错误流程 |
| [modules](modules/demo.md) | Demo 模块说明 |
| [workflows](workflows/db-cli.md) | DB CLI 和运维型命令 |
| [testing](testing/test-matrix.md) | 测试归属和验证命令 |
| [build](build/docker-and-ci.md) | CI、本地构建、Docker 构建和质量门禁 |
| [release](release/deployment.md) | 生产配置、部署脚本和发布检查 |
| [extension](extension/adding-modules.md) | 新增模块、配置和 API 的方式 |
| [maintenance](maintenance/maintenance-guide.md) | 长期维护工作流 |
| [ai-agent](ai-agent/runtime-state.md) | `AGENTS.md` 和 `docs/ai` 运行时说明 |
| [backlog](backlog/known-gaps.md) | 当前已知实现和文档缺口 |

## 兼容入口

- [project-overview.md](project-overview.md)
- [usage-guide.md](usage-guide.md)
- [configuration.md](configuration.md)
- [db-cli.md](db-cli.md)
- [deployment.md](deployment.md)
- [maintenance-guide.md](maintenance-guide.md)
- [agent-workflow-guide.md](agent-workflow-guide.md)
