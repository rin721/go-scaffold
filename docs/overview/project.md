# 项目概述

`go-scaffold` 是一个可运行的 Go 后端脚手架，包含服务代码、基础设施包、测试、Docker 构建文件、部署示例和 AI 运行时制品。

## 当前能力

| 能力 | 状态 |
| --- | --- |
| HTTP 服务 | 由 `internal/transport/http` 和 `pkg/httpserver` 实现 |
| 配置 | 支持 YAML、`.env`、环境变量覆盖、校验和监听重载 |
| 数据库 | `pkg/database` 支持 SQLite、MySQL、PostgreSQL |
| Demo 模块 | Todo CRUD，采用 handler/service/repository/model 分层 |
| 存储 | 本地文件系统抽象和可选 watcher 辅助能力 |
| SQL 生成 | Go 结构体到 SQL 的辅助工具，用于 DB CLI 和表结构应用 |
| CI/构建 | Go 测试、服务构建、Docker 构建和空白字符检查 |
| 部署 | 生产配置示例、Docker Compose 示例、本地部署脚本和远程工作流 |

## 已移除范围

脚手架不再内置用户管理栈。以下内容有意不存在：

- 本地用户模块；
- IAM 服务；
- JWT 认证令牌包；
- Casbin RBAC 包；
- 用户、角色、权限表结构；
- `/api/v1/auth`、`/api/v1/users`、`/api/v1/roles`、`/api/v1/permissions` 路由。

后续身份认证或访问控制能力应从新的需求基线开始，不应默认复活已删除实现。

## 当前非目标

- 完整生产迁移框架；
- 外部扩展运行时；
- v1 发布保证；
- 内置账号生命周期管理。

## 运行时默认值

本地默认配置是 `configs/config.yaml`。服务监听 `127.0.0.1:9999`，使用 SQLite `./data/app.db`，关闭 Redis，并启用 Demo Todo 模块。

生产示例位于 `deploy` 目录，默认关闭 Demo 模块。
