# 项目概览

`go-scaffold` 是一个围绕真实可运行服务构建的 Go 后端脚手架。它不只是文档仓库或 AI 协作空壳。当前仓库已经包含应用代码、基础设施包、测试、Docker
构建文件、部署示例和 AI 运行态产物。

## 当前能力

| 能力 | 当前状态 |
| --- | --- |
| HTTP 服务 | 通过 `internal/transport/http` 和 `pkg/httpserver` 实现 |
| 配置 | 支持 YAML、`.env`、环境变量覆盖、校验和 watch/reload |
| 数据库 | 通过 `pkg/database` 支持 SQLite/MySQL/PostgreSQL |
| Demo 模块 | Todo CRUD，展示 handler/service/repository/model 分层 |
| 用户模块 | 本地用户、密码哈希、bearer token、角色、权限和 RBAC 检查 |
| RBAC | 基于 Casbin 的授权器，支持配置化种子数据 |
| 插件系统 | 被动插件管理器、本地/HTTP 插件、hook registry 和远程 blog 示例 |
| 存储 | 本地文件系统抽象和可选 watcher 工具 |
| SQL 生成 | Go struct 到 SQL 的辅助工具，用于 DB CLI 和启动 schema 应用 |
| CI/构建 | 根模块测试、远程插件测试、服务构建、Docker 构建和空白检查 |
| 部署 | 生产配置示例、Docker Compose 示例、本地部署脚本、远程 workflow |

## 当前非目标

当前脚手架尚不承诺这些生产级能力：

- refresh token、session revoke、token rotation、审计流或密码重置；
- 外部 IAM 集成；
- 完整生产迁移体系；
- 远程插件的真实持久化，当前 blog 插件只是内存示例；
- v1 发布契约。

## 运行默认值

本地默认配置是 `configs/config.yaml`。它监听 `127.0.0.1:9999`，使用 SQLite
`./data/app.db`，关闭 Redis，启用 demo Todo，并且在未显式配置 auth secret 时可以使用进程内随机 token secret。

生产示例位于 `deploy`。生产路径要求显式提供 `RIN_APP_AUTH_TOKEN_SECRET`，并默认关闭 demo 模块。

## 事实来源

代码和当前配置是首要事实来源。包级 README 可以作为局部说明，但工程定位、目录边界和维护流程应优先以本结构化文档为准。
