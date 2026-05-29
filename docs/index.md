# 文档入口

这里是 `vibecoding2labs` 的人类文档入口。`docs/` 面向开发者、维护者和未来接手者；`docs/ai/` 面向 AI Agent 的运行期状态和证据。

## 推荐阅读顺序

1. [项目概述](project-overview.md)
2. [使用教程](usage-guide.md)
3. [维护与演进指南](maintenance-guide.md)
4. [Agent 协作流程](agent-workflow-guide.md)

## 文档地图

| 文件 | 适合谁看 | 解决什么问题 |
| --- | --- | --- |
| [../README.md](../README.md) | 第一次打开仓库的人 | 快速了解项目状态和入口。 |
| [project-overview.md](project-overview.md) | 开发者、维护者 | 解释项目定位、目录结构、运行期基建和边界。 |
| [usage-guide.md](usage-guide.md) | 想直接使用仓库的人 | 给出从打开仓库到推进下一轮工作的步骤。 |
| [maintenance-guide.md](maintenance-guide.md) | 长期维护者 | 说明怎么更新状态、证据、知识库、skills 和文档。 |
| [agent-workflow-guide.md](agent-workflow-guide.md) | 人和 Agent | 说明协作顺序、确认点、下一步协议和禁止事项。 |
| [ai/tasks/forest.yaml](ai/tasks/forest.yaml) | Agent、维护者 | 区分 Vibe Coding 基建支线和未来业务主线任务树。 |

## AI 运行期入口

| 文件 | 用途 |
| --- | --- |
| [../AGENTS.md](../AGENTS.md) | Agent 短索引。 |
| [ai/runtime-rule-index.md](ai/runtime-rule-index.md) | 运行期规则索引。 |
| [ai/status/current.yaml](ai/status/current.yaml) | 当前状态第一入口。 |
| [ai/tasks/current-slice.yaml](ai/tasks/current-slice.yaml) | 当前执行切片。 |
| [ai/requirements/ledger.yaml](ai/requirements/ledger.yaml) | 需求台账。 |
| [ai/evidence/index.md](ai/evidence/index.md) | 证据入口。 |
| [ai/handoff/current.md](ai/handoff/current.md) | 当前交接说明。 |

## 文档维护原则

- 稳定说明放在 `docs/`。
- 过程状态、证据和 Agent 工作文件放在 `docs/ai/`。
- 外部资料只作为候选事实，验证后才能进入知识库。
- 不要把“回读原始提示词”写成恢复路径。
