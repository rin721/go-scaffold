# REQUIREMENTS.md

## 需求状态

- 项目：go-scaffold
- 阶段：需求确认
- 状态：COMPLETED
- 最后更新：2026-05-26
- 输入：用户发送“下一步”，按当前文档默认值确认 `TASK-OPT-002`

## 确认结果

- [CONFIRMED] 优化路线采用“治理优先”。
- [CONFIRMED] `pkg/*` 采用混合策略：可复用基础设施包按公共 API 管理，明确项目内部支撑包不承诺外部兼容。
- [CONFIRMED] demo 模块暂定为长期标准示例，用于展示模块分层和后续测试样板。
- [CONFIRMED] 迁移策略采用 dev-prod 分层：开发/demo 可使用 `AutoMigrate`，生产/bootstrap 倾向显式 SQL 或迁移流程。
- [CONFIRMED] 中文化范围先覆盖根文档和模板；第一阶段包 README 已纳入 TASK-P1-017，历史文档仍分阶段处理。
- [CONFIRMED] auth/JWT 先延后处理，不在当前 P0/P1 代码实现范围内。

## 项目目标

- [CONFIRMED] 分析当前项目优缺点，收拢不一致设计边界，形成可拆分、可验证、可交接的优化路线。
- [CONFIRMED] 优先建立稳定项目事实、状态、风险、验收和路线图，再进入任何代码优化。
- [CONFIRMED] 后续每个代码优化必须映射到已确认需求、架构边界、任务和时间切片。

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
| REQ-OPT-P2-003 | CI/CD 与部署 | 首切片已完成 CI 质量门禁和部署说明；显式参数部署入口、手动 staging/production 远程部署 workflow 闸门、Dockerfile、production Compose 示例和统一 `deploy.sh` 部署入口已补齐；Docker build 待具备 Docker 的环境验证，镜像发布和真实 production 运行仍需单独确认 | [CONFIRMED] |
| REQ-OPT-P2-004 | 性能基准测试 | 需先完成测试矩阵和功能边界收拢 | [DEFERRED] |
| REQ-OPT-P2-005 | 脚手架生成器 | 需先确认框架化抽取路线 | [DEFERRED] |

## 非需求

- [CONFIRMED] 当前不修改 Go API、配置结构、数据库结构或 HTTP 路由。
- [CONFIRMED] 当前不执行部署、生产命令或不可逆迁移。
- [CONFIRMED] 当前不删除插件系统历史记录。

## 完成判断

- TASK-OPT-004：COMPLETED
- 下一合法任务：TASK-OPT-005，确认正式测试矩阵和 P1 执行顺序。
