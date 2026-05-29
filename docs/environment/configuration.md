# 配置说明

配置由 `internal/config` 加载。默认文件是 `configs/config.yaml`，示例位于
`configs/config.example.yaml` 和 `deploy/config.production.example.yaml`。

## 配置路径解析

`server` 和 `db` 命令都支持 `--config`：

```bash
go run ./cmd/main server --config=configs/config.yaml
go run ./cmd/main db --config=configs/config.yaml --operation=schema
```

如果没有传入 `--config`，进程会检查 `RIN_CONFIG_PATH`，最后回退到默认配置路径。

`.env` 在 `Manager.Load` 内部加载，此时命令层已经选定配置路径。因此 `.env` 适合做配置字段覆盖，不适合作为
`RIN_CONFIG_PATH` 的唯一来源。

## 加载顺序

1. 从命令参数、环境变量或默认值选择配置路径。
2. 如果当前工作目录存在 `.env`，加载它。
3. 读取 YAML 配置。
4. 替换 `${VAR}` 和 `${VAR:default}` 占位符。
5. 反序列化到 `internal/config.Config`。
6. 根据 `envname` tag 应用环境变量覆盖。
7. 校验最终配置。
8. 存入配置管理器。

真实系统环境变量优先于 `.env` 中的同名值。

## 环境变量命名

当前 app prefix 是 `Rin`，派生结果如下：

| 用途 | 名称 |
| --- | --- |
| 配置路径 | `RIN_CONFIG_PATH` |
| 配置字段前缀 | `RIN_APP` |

每个配置字段通过 `envname` 声明未加前缀名称。加载器会先查找带前缀名称，再查找未加前缀 fallback：

```text
RIN_APP_DB_HOST
DB_HOST
```

新文档、新部署配置和新测试应优先使用 `RIN_APP_*` 形式。未加前缀名称仅作为兼容路径保留。

## 常用变量

| 范围 | 变量 |
| --- | --- |
| Server | `RIN_APP_SERVER_HOST`, `RIN_APP_SERVER_PORT`, `RIN_APP_SERVER_MODE`, `RIN_APP_SERVER_READ_TIMEOUT`, `RIN_APP_SERVER_WRITE_TIMEOUT`, `RIN_APP_SERVER_IDLE_TIMEOUT` |
| Database | `RIN_APP_DB_DRIVER`, `RIN_APP_DB_HOST`, `RIN_APP_DB_PORT`, `RIN_APP_DB_USER`, `RIN_APP_DB_PASSWORD`, `RIN_APP_DB_NAME`, `RIN_APP_DB_MAX_OPEN_CONNS`, `RIN_APP_DB_MAX_IDLE_CONNS` |
| Redis | `RIN_APP_REDIS_ENABLED`, `RIN_APP_REDIS_HOST`, `RIN_APP_REDIS_PORT`, `RIN_APP_REDIS_PASSWORD`, `RIN_APP_REDIS_DB`, `RIN_APP_REDIS_POOL_SIZE` |
| Logger | `RIN_APP_LOG_LEVEL`, `RIN_APP_LOG_FORMAT`, `RIN_APP_LOG_OUTPUT`, `RIN_APP_LOG_FILE_PATH`, `RIN_APP_LOG_MAX_SIZE`, `RIN_APP_LOG_MAX_BACKUPS`, `RIN_APP_LOG_MAX_AGE` |
| I18n | `RIN_APP_I18N_DEFAULT`, `RIN_APP_I18N_SUPPORTED`, `RIN_APP_I18N_MESSAGES_DIR` |
| Executor | `RIN_APP_EXECUTOR_ENABLED` |
| Storage | `RIN_APP_STORAGE_ENABLED`, `RIN_APP_STORAGE_FS_TYPE`, `RIN_APP_STORAGE_BASE_PATH`, `RIN_APP_STORAGE_ENABLE_WATCH`, `RIN_APP_STORAGE_WATCH_BUFFER_SIZE` |
| Demo | `RIN_APP_DEMO_ENABLED`, `RIN_APP_DEMO_APPLY_SCHEMA_ON_START` |
| Plugin | `RIN_APP_PLUGIN_ENABLED`, `RIN_APP_PLUGIN_DEFAULT_TIMEOUT`, `RIN_APP_PLUGIN_MAX_RESPONSE_BYTES` |
| IAM | `RIN_APP_IAM_ENABLED`, `RIN_APP_IAM_MODE`, `RIN_APP_IAM_DEFAULT_DENY` |
| Auth | `RIN_APP_AUTH_TOKEN_SECRET`, `RIN_APP_AUTH_TOKEN_TTL`, `RIN_APP_AUTH_PUBLIC_REGISTRATION` |
| CORS | `RIN_APP_CORS_ENABLED`, `RIN_APP_CORS_ALLOW_ORIGINS`, `RIN_APP_CORS_ALLOW_METHODS`, `RIN_APP_CORS_ALLOW_HEADERS`, `RIN_APP_CORS_EXPOSE_HEADERS`, `RIN_APP_CORS_ALLOW_CREDENTIALS`, `RIN_APP_CORS_MAX_AGE` |

完整字段以 `.env.example` 和 `internal/config/*` 中的 `envname` tag 为准。

## 本地与生产默认值

本地配置：

- SQLite 路径为 `./data/app.db`；
- Redis 关闭；
- demo 模块启用；
- RBAC seed 启用；
- auth 未显式配置 secret 时，可以生成进程内随机 token secret。

生产示例：

- 绑定到 `0.0.0.0`；
- 要求 `RIN_APP_AUTH_TOKEN_SECRET`；
- 默认关闭 demo；
- 敏感值应通过环境变量、CI/CD secret 或容器编排注入。

## 新增配置字段流程

1. 在对应的 `internal/config/*Config` 结构体增加字段。
2. 增加 `mapstructure` 和 `envname` tag。
3. 如字段有安全或格式约束，补充校验。
4. 必要时更新 `configs/config.example.yaml`、`.env.example` 和 production 配置示例。
5. 在 `internal/config` 添加或调整测试。
6. 如果字段面向使用者或运维者，同步更新本文档。
