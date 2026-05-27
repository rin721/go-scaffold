# REQUIREMENTS.md

## 需求状态

- 项目：go-scaffold
- 阶段：需求确认
- 状态：COMPLETED
- 发布状态：NOT_RELEASE_READY
- 最后更新：2026-05-28
- 输入：用户发送“下一步”，按当前文档默认值确认 `TASK-OPT-002`

## 确认结果

- [CONFIRMED] 优化路线采用“治理优先”。
- [CONFIRMED] `pkg/*` 采用混合策略：可复用基础设施包按公共 API 管理，明确项目内部支撑包不承诺外部兼容。
- [CONFIRMED] demo 模块暂定为长期标准示例，用于展示模块分层和后续测试样板。
- [CONFIRMED] 数据库 schema 与 demo CRUD/bootstrap 统一通过 `pkg/sqlgen` 工具链生成 SQL；`cmd/server db` 是显式 DB CLI 入口；`initdb`、SQL 脚本目录和运行期 `AutoMigrate` 已移除；生产迁移框架仍需单独确认。
- [CONFIRMED] 中文化范围先覆盖根文档和模板；第一阶段包 README 已纳入 TASK-P1-017，历史文档仍分阶段处理。
- [CONFIRMED] auth/JWT 先延后处理，不在当前 P0/P1 代码实现范围内。

## 项目目标

- [CONFIRMED] 分析当前项目优缺点，收拢不一致设计边界，形成可拆分、可验证、可交接的优化路线。
- [CONFIRMED] 优先建立稳定项目事实、状态、风险、验收和路线图，再进入任何代码优化。
- [CONFIRMED] 后续每个代码优化必须映射到已确认需求、架构边界、任务和时间切片。
- [ACCEPT] 当前项目未达第一版发布条件；发布 v1 前必须另行确认完整功能范围、质量门禁、真实环境验证、迁移和安全边界。
- [ACCEPT] `types/*` 不得直接为 `pkg/*` 基础设施提供别名、注入接口或 typed constants；`types` 只能承载应用层以上确认过的跨层契约。
- [ACCEPT_WITH_RISK] `internal/config` 环境变量覆盖必须从 `AppPrefix` 动态派生前缀，不再固定 `REI_APP`；配置字段使用 `envname` 标签声明环境变量名，并自动加载 `.env`。`EnvDB*` / `EnvRedis*` 等重复 env-name 常量已删除，不再作为字段命名来源。

## P0 需求

| ID | 需求 | 验收 | 状态 |
|---|---|---|---|
| REQ-OPT-P0-001 | 中文启动材料生成 | `docs/templates/*` 六个模板已中文化 | [CONFIRMED] |
| REQ-OPT-P0-002 | 当前主线切换到项目治理与优化路线 | `STATUS.md`、`TASKS.md`、`TIME_SLICES.md` 不再指向插件扩展 | [CONFIRMED] |
| REQ-OPT-P0-003 | 明确默认优化路线 | `DECISIONS.md` 记录治理优先和关键边界默认值 | [CONFIRMED] |
| REQ-OPT-P0-004 | 生成路线图 | `ROADMAP.md` 存在并包含阶段目标、退出条件和验证要求 | [CONFIRMED] |
| REQ-OPT-P0-005 | 保持代码行为不变 | 不修改 Go 代码；`go test ./... -count=1` 通过 | [CONFIRMED] |

## P1 需求

| ID | 需求 | 验收 | 状态 |
|---|---|---|---|
| REQ-OPT-P1-001 | 生成模块边界清单 | `MODULES.md` 梳理 `cmd`、`internal/*`、`pkg/*`、`types/*` 的职责、风险和优化方向 | [CONFIRMED] |
| REQ-OPT-P1-002 | 生成测试矩阵 | `MODULES.md` 包含 app 启动、health/ready、demo CRUD、配置加载/重载、迁移策略测试矩阵草案 | [CONFIRMED] |
| REQ-OPT-P1-003 | 生成设计边界收拢清单 | `MODULES.md` 明确迁移、配置、插件、包 API、demo、auth/JWT 的统一策略风险 | [CONFIRMED] |
| REQ-OPT-P1-004 | 拆分优化任务和时间切片 | 每个 P1 优化项有允许文件范围、验证命令和退出条件 | [CONFIRMED] |
| REQ-OPT-P1-005 | 分阶段中文化历史内容 | 根文档和模板已完成；第一阶段 `pkg/*/README.md` 中文化已由 TASK-P1-017 覆盖 | [CONFIRMED] |

## P2 延后需求

| ID | 需求 | 延后原因 | 状态 |
|---|---|---|---|
| REQ-OPT-P2-001 | auth/rbac 实现 | 当前 README 明确暂不实现，需要单独需求确认 | [DEFERRED] |
| REQ-OPT-P2-002 | 插件系统 rpc/ws/discovery 扩展 | 插件 v1 已完成，扩展需独立提升 | [DEFERRED] |
| REQ-OPT-P2-003 | CI/CD 与部署 | 首切片已完成 CI 质量门禁和部署说明；显式参数部署入口、手动 staging/production 远程部署 workflow 闸门、Dockerfile、production Compose 示例、统一 `deploy.sh` 部署入口和 Docker build 验证已完成；镜像发布和真实 production 运行仍需单独确认 | [CONFIRMED] |
| REQ-OPT-P2-004 | 性能基准测试 | 需先完成测试矩阵和功能边界收拢 | [DEFERRED] |
| REQ-OPT-P2-005 | 脚手架生成器 | 需先确认框架化抽取路线 | [DEFERRED] |
| REQ-OPT-P2-006 | 插件钩子运行时、HTTP 远程插件传输和 IAM 公共接口 | 用户确认 `dev.tmp/new-plugin.md` 设计；本轮仅实现公共基础设施接口、内存 IAM 和 app 组合层接入，不实现完整业务登录/RBAC | [CONFIRMED] |
| REQ-OPT-P2-007 | 配置环境变量动态前缀与字段标签覆盖 | 用户要求 `internal/config` 结合 `AppPrefix` 自动注入环境变量，并使用 `envname` 配置字段环境变量名；已实现 `RIN_APP_*` 动态前缀、未加前缀 fallback 和 `.env` 自动加载测试；重复 env-name 常量已删除，字段名以 `envname` 标签为唯一事实源 | [CONFIRMED] |
| REQ-OPT-P2-008 | 远程插件显式注册、IAM hook JSON 上下文和 Blog 示例 | 用户通过 `/goal` 确认新范围；TASK-P2-016 已完成 HTTP 显式注册、IAM 安全主体上下文注入和 `remote_plugins/blog` 示例；真实 WS/RPC、JWT/login、数据库版权限或生产部署仍不实现 | [CONFIRMED] |
| REQ-OPT-P2-009 | Plugin control-plane interface address configuration | User correction requires remote plugin services to expose standard `/plugin/v1/invoke`, register to host `/plugin/v1/register`, and align HTTP/WS protocol addresses through host plugin-system configuration. TASK-P2-017 implements configurable host HTTP interface exposure and reserved WS address placeholders only. | [CONFIRMED] |

## 非需求

- [CONFIRMED] 当前不修改 Go API、配置结构、数据库结构或 HTTP 路由。
- [CONFIRMED] 当前不执行部署、生产命令或不可逆迁移。
- [CONFIRMED] 当前不发布第一版、不创建 v1 release、不宣称 release-ready。
- [CONFIRMED] 当前不删除插件系统历史记录。
- [CONFIRMED] 当前不实现 JWT 中间件、数据库版权限、OPA/Casbin、Go `.so` 插件、插件发现、RPC/WS 传输、镜像发布、真实生产部署或密钥管理。
- [CONFIRMED] `types/*` 不作为 `pkg/*` 基础设施聚合入口；缓存、加密和 executor 等接口不通过 `types` 重新导出或直接依赖。

## 完成判断

- TASK-OPT-004：COMPLETED
- 下一合法任务：TASK-OPT-005，确认正式测试矩阵和 P1 执行顺序。
