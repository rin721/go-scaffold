# 目录地图

项目采用务实的分层布局。核心规则是：`cmd/main` 启动进程，`internal/app` 装配应用，`internal`
保存应用特有实现，`pkg` 保存可复用支撑包，`types` 保存共享响应、错误和常量辅助。

## 根目录

| 路径 | 职责 | 被谁使用 |
| --- | --- | --- |
| `README.md` | 仓库快速入口 | 人类和 AI Agent |
| `go.mod`, `go.sum` | 根 Go module 依赖声明 | 根模块下所有包 |
| `Dockerfile` | 多阶段 Linux 镜像构建 | CI 和发布流程 |
| `.github/workflows/ci.yml` | 分支/PR 质量检查 | 维护者 |
| `.github/workflows/deploy-remote.yml` | 手动远程部署 workflow | 维护者 |
| `deploy.sh` | Bash 部署包装器 | 远程安装和手动部署 |

## 应用代码

| 路径 | 职责 | 说明 |
| --- | --- | --- |
| `cmd/main` | CLI 命令、进程入口、信号处理 | 保持轻量，不放业务逻辑 |
| `internal/app` | 组合根、生命周期、reload 编排 | 负责装配顺序和模块组装 |
| `internal/config` | 配置结构、加载、环境覆盖、校验、watch | `envname` 映射的单一事实来源 |
| `internal/transport/http` | Gin 路由和 HTTP handler 注册 | 通过 handler 接入模块 |
| `internal/middleware` | trace、logger、recovery、i18n、CORS 等 HTTP 中间件 | 只处理传输层关注点 |
| `internal/modules/demo` | Demo Todo 业务模块 | 模块分层示例 |
| `internal/modules/user` | 用户、认证、RBAC 业务模块 | 当前主要身份模块 |
| `internal/iam` | 插件/IAM hook 所需的访问上下文抽象 | 应用特有上下文 |

## 基础设施包

| 路径 | 职责 | 边界 |
| --- | --- | --- |
| `pkg/database` | GORM 数据库管理、ping、reload、事务 | 不知道应用模块 |
| `pkg/cache` | Redis 管理器 | 可选基础设施 |
| `pkg/logger` | Zap 日志封装 | 被 app 和 package 复用 |
| `pkg/httpserver` | HTTP server 启停和 reload 包装 | 不定义路由 |
| `pkg/executor` | goroutine pool 管理器 | 可选支撑设施 |
| `pkg/auth` | JWT token 服务 | 通用 token 基础设施 |
| `pkg/rbac` | Casbin 授权适配器 | 通用 RBAC 基础设施 |
| `pkg/crypto` | 密码哈希 | 通用加密辅助 |
| `pkg/plugin` | 插件管理、hook、HTTP 插件支持 | 不直接依赖业务模块 |
| `pkg/storage` | 文件存储和 watcher 工具 | 可选存储基础设施 |
| `pkg/sqlgen` | SQL DDL/代码生成辅助 | 被 DB CLI 和 schema 应用使用 |
| `pkg/i18n` | 消息包和语言辅助 | 可选本地化支撑 |
| `pkg/utils` | 零散工具 | 需避免膨胀成隐藏业务层 |

## 共享类型

`types` 保存常量、类型化错误和 HTTP result 辅助。注意：`types/result` 依赖 Gin，所以当前
`types` 并不是完全纯净的领域类型层。修改 result helper 时要把它视为明确的现状边界。

## 远程插件

`remote_plugins/blog` 是独立 Go module，用作远程 HTTP 插件示例。它有自己的 `go.mod`、测试、配置和 README。根目录的
`go test ./...` 不会覆盖它，CI 需要单独测试。

## 文档和运行态

| 路径 | 职责 |
| --- | --- |
| `docs` | 人类工程文档 |
| `docs/ai` | AI 运行态、任务树、决策、证据、技能和交接 |
| `AGENTS.md` | Agent 短索引 |

不要把“回读原始 prompt”作为未来恢复路径。如果运行态太薄，应修复 `docs/ai` 下的物理 artifact。
