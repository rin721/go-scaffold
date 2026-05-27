# 配置文档说明

本文说明当前项目的配置加载顺序、动态环境变量前缀、`.env` 自动加载规则，以及新增配置字段时应如何维护 `envname`。

## 配置入口

应用默认读取 `configs/config.yaml`。`server` 和 `initdb` 命令都支持显式传入配置文件：

```bash
go run ./cmd/server server --config=configs/config.yaml
go run ./cmd/server initdb --config=configs/config.yaml
```

也可以通过配置路径环境变量指定：

```bash
export RIN_CONFIG_PATH=configs/config.yaml
go run ./cmd/server server
```

`RIN_CONFIG_PATH` 由 `internal/config.EnvConfigPathName()` 基于 `types/constants.AppPrefix` 动态生成。当前 `AppPrefix = "Rin"`，因此配置路径变量是 `RIN_CONFIG_PATH`。

注意：`.env` 文件是在 `Manager.Load` 内加载的，加载时命令行配置路径已经确定。因此 `.env` 适合配置字段覆盖，不适合作为 `RIN_CONFIG_PATH` 的唯一来源。配置路径优先使用命令行 `--config`，其次使用进程环境变量 `RIN_CONFIG_PATH`，最后使用默认路径。

## 加载顺序

当前配置加载流程如下：

1. 命令行层决定配置文件路径。
2. `Manager.Load` 自动加载当前工作目录下的 `.env`，文件不存在时静默跳过。
3. 读取 YAML 配置文件。
4. 处理配置文件中的 `${VAR}` 或 `${VAR:default}` 字符串替换。
5. 反序列化到 `internal/config.Config`。
6. 通过 `envname` 标签读取环境变量并覆盖配置字段。
7. 校验配置并存入配置管理器。

`.env` 使用 `godotenv.Load` 加载，不覆盖已经存在的系统环境变量。也就是说，真实进程环境变量优先于 `.env` 文件中的同名变量。

## 动态环境变量前缀

配置字段环境变量前缀不再写死。`internal/config.EnvPrefix()` 会从 `types/constants.AppPrefix` 派生前缀：

| `AppPrefix` | 配置字段前缀 | 配置路径变量 |
|---|---|---|
| `Rin` | `RIN_APP` | `RIN_CONFIG_PATH` |

配置字段完整环境变量名格式为：

```text
<APP_PREFIX>_APP_<MODULE>_<FIELD>
```

例如当前前缀下：

```bash
export RIN_APP_SERVER_PORT=9090
export RIN_APP_DB_HOST=127.0.0.1
export RIN_APP_LOG_LEVEL=debug
```

未加前缀的变量仍作为兼容 fallback 保留，例如 `DB_HOST`。但新文档、新部署配置和新测试都应使用 `RIN_APP_*` 主格式。

## envname 单一事实源

配置字段的环境变量名只由结构体字段上的 `envname` 标签声明：

```go
type DatabaseConfig struct {
    Host string `mapstructure:"host" envname:"DB_HOST"`
}
```

运行时会先查找动态前缀变量，再查找未加前缀 fallback：

```text
RIN_APP_DB_HOST
DB_HOST
```

不要再为字段环境变量名新增 `EnvDBHost`、`EnvRedisHost`、`EnvServerPort` 这类镜像常量。它们会把 `envname` 标签变成第二事实源，后续改名容易漂移。

## 支持的数据类型

环境变量覆盖目前支持常见标量和切片：

| Go 字段类型 | 说明 |
|---|---|
| `string` | 原样使用环境变量值 |
| `bool` | 使用 `strconv.ParseBool` 解析 |
| `int` / `int64` 等整数 | 使用十进制解析 |
| `uint` / `uint64` 等无符号整数 | 使用十进制解析 |
| `float32` / `float64` | 使用浮点解析 |
| `[]string` | 默认使用英文逗号分隔 |
| 其他 slice | 使用 JSON 解析 |
| 指针字段 | 解析成功后自动创建并赋值 |

空字符串不会覆盖配置文件值。解析失败时该字段保持原配置值。

## 常用变量

完整清单以 `.env.example` 和 `internal/config/*` 中的 `envname` 标签为准。常用变量如下：

| 模块 | 常用变量 |
|---|---|
| Server | `RIN_APP_SERVER_HOST`, `RIN_APP_SERVER_PORT`, `RIN_APP_SERVER_MODE`, `RIN_APP_SERVER_READ_TIMEOUT`, `RIN_APP_SERVER_WRITE_TIMEOUT`, `RIN_APP_SERVER_IDLE_TIMEOUT` |
| Database | `RIN_APP_DB_DRIVER`, `RIN_APP_DB_HOST`, `RIN_APP_DB_PORT`, `RIN_APP_DB_USER`, `RIN_APP_DB_PASSWORD`, `RIN_APP_DB_NAME`, `RIN_APP_DB_MAX_OPEN_CONNS`, `RIN_APP_DB_MAX_IDLE_CONNS` |
| Redis | `RIN_APP_REDIS_ENABLED`, `RIN_APP_REDIS_HOST`, `RIN_APP_REDIS_PORT`, `RIN_APP_REDIS_PASSWORD`, `RIN_APP_REDIS_DB`, `RIN_APP_REDIS_POOL_SIZE` |
| Logger | `RIN_APP_LOG_LEVEL`, `RIN_APP_LOG_FORMAT`, `RIN_APP_LOG_OUTPUT`, `RIN_APP_LOG_FILE_PATH`, `RIN_APP_LOG_MAX_SIZE`, `RIN_APP_LOG_MAX_BACKUPS`, `RIN_APP_LOG_MAX_AGE` |
| I18n | `RIN_APP_I18N_DEFAULT`, `RIN_APP_I18N_SUPPORTED`, `RIN_APP_I18N_MESSAGES_DIR` |
| InitDB | `RIN_APP_INITDB_SCRIPT_DIR`, `RIN_APP_INITDB_LOCK_FILE`, `RIN_APP_INITDB_SCRIPT_FILE_PREFIX` |
| Executor | `RIN_APP_EXECUTOR_ENABLED` |
| Storage | `RIN_APP_STORAGE_ENABLED`, `RIN_APP_STORAGE_FS_TYPE`, `RIN_APP_STORAGE_BASE_PATH`, `RIN_APP_STORAGE_ENABLE_WATCH`, `RIN_APP_STORAGE_WATCH_BUFFER_SIZE` |
| Plugin | `RIN_APP_PLUGIN_ENABLED`, `RIN_APP_PLUGIN_DEFAULT_TIMEOUT`, `RIN_APP_PLUGIN_MAX_RESPONSE_BYTES` |
| IAM | `RIN_APP_IAM_ENABLED`, `RIN_APP_IAM_MODE`, `RIN_APP_IAM_DEFAULT_DENY` |
| CORS | `RIN_APP_CORS_ENABLED`, `RIN_APP_CORS_ALLOW_ORIGINS`, `RIN_APP_CORS_ALLOW_METHODS`, `RIN_APP_CORS_ALLOW_HEADERS`, `RIN_APP_CORS_EXPOSE_HEADERS`, `RIN_APP_CORS_ALLOW_CREDENTIALS`, `RIN_APP_CORS_MAX_AGE` |

## .env 示例

本地开发可以在仓库根目录创建 `.env`。不要提交真实 `.env`。

```dotenv
RIN_APP_SERVER_PORT=9090
RIN_APP_DB_DRIVER=sqlite
RIN_APP_DB_NAME=./data/app.db
RIN_APP_REDIS_ENABLED=false
RIN_APP_LOG_LEVEL=debug
```

生产环境应优先使用系统环境变量、CI/CD Secrets、容器编排环境变量或密钥管理服务注入敏感值，不要把真实数据库密码、Redis 密码、token 或 SSH 凭据写进仓库。

## 新增配置字段流程

新增配置字段时按以下顺序维护：

1. 在对应 `internal/config/*Config` 结构体字段上添加 `mapstructure` 和 `envname` 标签。
2. 如字段需要 YAML 示例，同步更新 `configs/config.example.yaml` 或 production 示例。
3. 如字段面向本地开发或部署使用，同步更新 `.env.example` 和相关部署文档。
4. 在配置测试中通过结构体 `envname` 标签读取变量名，不要新增字段环境变量常量。
5. 运行相关验证，至少覆盖 `go test ./internal/config -count=1` 和 `git diff --check`。

新增字段示例：

```go
type ServerConfig struct {
    PublicURL string `mapstructure:"public_url" envname:"SERVER_PUBLIC_URL"`
}
```

当前完整变量名会自动派生为：

```text
RIN_APP_SERVER_PUBLIC_URL
```

如果未来 `AppPrefix` 变化，完整变量名前缀会随之变化，不需要修改 `internal/config` 的字段常量表。
