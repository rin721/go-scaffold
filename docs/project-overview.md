# 项目概述

## 这个项目是什么

`vibecoding2labs` 当前是一个 Vibe Coding 协作基建仓库。它还没有业务应用代码，已经具备一套可恢复、可审计、可继续演进的 AI 协作运行层。

这套运行层解决的问题是：后续开发不只靠聊天上下文，而是靠仓库里的实体文件恢复状态、需求、任务、证据和下一步。

## 当前已经具备什么

| 能力 | 当前产物 |
| --- | --- |
| Agent 短入口 | `AGENTS.md` |
| 运行期规则索引 | `docs/ai/runtime-rule-index.md` |
| 当前状态 | `docs/ai/status/current.yaml` |
| 初始化任务树 | `docs/ai/tasks/bootstrap-tree.yaml` |
| 任务树林 | `docs/ai/tasks/forest.yaml` |
| 主线想法任务树 | `docs/ai/tasks/main-tree.yaml` |
| Vibe Coding 基建支线 | `docs/ai/tasks/branches/vibe-coding-infra/tree.yaml` |
| 当前执行切片 | `docs/ai/tasks/current-slice.yaml` |
| 需求入口 | `docs/ai/requirements/ledger.yaml` |
| 决策记录 | `docs/ai/decisions/records.md` |
| 证据入口 | `docs/ai/evidence/index.md` |
| 知识库入口 | `docs/ai/knowledge/index.md` |
| 项目技能入口 | `docs/ai/skills/index.md` |
| 交接说明 | `docs/ai/handoff/current.md` |

## 三层边界

| 层级 | 含义 | 当前仓库如何处理 |
| --- | --- | --- |
| 上游指令 | 产生编译器文本的来源 | 不进入仓库运行期权威链。 |
| 编译器提示词 | 用于初始化、迁移或重大升级 | 已被编译成运行期文件，后续不作为日常依赖。 |
| 运行期产物 | Agent 和开发者日常使用的实体文件 | 已落在 `AGENTS.md` 和 `docs/ai/*`。 |

## 工作流概览

```text
想法输入
-> 需求确认
-> 研究验证
-> 任务分析
-> 架构与技术栈确认
-> 基建模式确认
-> 任务树与执行切片
-> 实现或文档变更
-> 验证
-> 证据、知识库、skills、状态和交接同步
-> 轮次关闭
```

当前仓库已经不再是纯文档基建仓库：真实代码包含 Go 后端脚手架、HTTP 服务、demo Todo、用户/认证/RBAC、插件、存储、SQL 生成器和 Docker/部署样例。`docs/ai/*` 仍作为 Agent 运行态与审计状态入口；后续代码修复、功能补齐和架构复位必须同时维护代码、验证结果和运行态文档。

## 当前不是做什么

当前项目还不是：

- 一个已经有业务功能的应用。
- 一个可以直接部署的服务。
- 一个依赖原始提示词才能运行的项目。
- 一个允许 Agent 自由改任何文件的自动化系统。

## 可靠性的来源

这个项目的可靠性来自四件事：

1. 状态落盘：当前事实写在 `docs/ai/status/current.yaml`。
2. 切片约束：每次执行都有允许文件、验证命令和回滚方式。
3. 证据记录：命令、验证和关键判断写入 `docs/ai/evidence/index.md`。
4. 人类确认：需求、架构、模式、风险和验收仍由开发者最终确认。
