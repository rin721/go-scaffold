# 当前交接

- Project: `vibecoding2labs`
- Runtime version: `vibe-runtime-0.1.0`
- Current phase: `verification`
- Current round: `infra_001`
- Current branch: `branch_vibe_coding_infra`
- Current tree: `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- Current mode: `chinese_delivery_rule_synced`
- Agency level: `controlled_execution`
- Current slice: `slice_013_chinese_delivery_rule`

## 当前事实

- 仓库当前是一个可运行的 Go 后端脚手架，同时包含 Vibe Coding 的 `docs/ai` 运行时制品。
- 已移除内置 IAM、auth、RBAC 和相关本地用户管理服务。当前服务只保留 Demo Todo、基础设施包、配置、HTTP、Docker/CI/部署示例和 AI 运行时文档。
- 最新开发者规则：以后面向人的交付内容默认使用中文；新增或修改代码中的解释性注释使用中文；文档以中文概述当前事实。技术标识符、命令、配置键、包名、路径和外部专有名词可保留原文。
- 不要把旧记录里“没有业务代码”的说法当作当前事实；该说法已被当前 Go 脚手架实物状态取代。

## 已完成

- 建立 `AGENTS.md` 和 `docs/ai/*` 最小运行时治理入口。
- 建立任务森林，区分主线产品任务和 `branch_vibe_coding_infra` 基建任务。
- 增加完整项目生命周期、需求发现、下游生命周期、编译器上下文吸收、声明式需求工作流等运行时规则和技能。
- 完成 Go 脚手架维护修复：Docker/CI/部署漂移、生产配置示例、Demo 默认值、HTTP 生命周期、SQL DDL、运行时文档漂移等已被整理。
- 移除旧扩展运行时和相关服务，并通过 Go 测试、构建和运行时校验。
- 移除 IAM/auth/RBAC/用户管理栈：删除 `internal/modules/user`、`pkg/auth`、`pkg/iam`、`pkg/rbac`、RBAC model、用户 schema helper、路由注册、配置/env/deploy 入口、过期测试，并清理 `go.mod`/`go.sum`。
- 新增中文交付规则，并将上一轮移除工作涉及的 README、工程文档、配置示例、部署帮助和 CORS 注释同步为中文口径。

## 未完成

- 用户管理栈移除和中文交付规则同步都处于本地验证通过、等待开发者验收状态。
- 未来身份认证、访问控制、扩展运行时或产品能力扩展都应从新的需求基线开始，不应默认恢复已删除实现。

## 关键文件

- `AGENTS.md`
- `docs/ai/runtime-rule-index.md`
- `docs/ai/status/current.yaml`
- `docs/ai/tasks/current-slice.yaml`
- `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml`
- `docs/ai/requirements/ledger.yaml`
- `docs/ai/evidence/index.md`
- `docs/ai/handoff/current.md`
- `README.md`
- `docs/index.md`
- `docs/overview/project.md`
- `docs/environment/configuration.md`
- `docs/architecture/layers.md`
- `docs/runtime/startup-flow.md`
- `docs/runtime/http-flow.md`
- `docs/runtime/config-flow.md`
- `.env.example`
- `deploy.sh`

## 下一条件

开发者验收 `slice_013_chinese_delivery_rule`。验收后，后续产物默认中文交付；如需英文或双语，应由开发者明确提出。

## 禁止事项

- 不要把原始编译器提示词作为普通恢复路径。
- 不要创建 `prompt.md` 作为运行时权威；外部治理规格必须提炼进本地运行时制品。
- 不要在未确认需求、研究、任务分析、架构和模式前启动新的主线业务实现。
- 不要从候选能力目录直接安装依赖、创建工程骨架或引入包管理清单。
- 不要在没有新确认任务时恢复已删除的 IAM/auth/RBAC/用户管理实现。
- 不要忽略中文交付规则；技术标识符可保留原文，但解释性内容应使用中文。
