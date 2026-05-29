# AI 运行态

仓库在 `docs/ai` 下包含 Vibe Coding 运行态。该子树不是普通产品文档，而是面向 Agent 的可读、可审计状态系统。

## 入口

| 路径 | 用途 |
| --- | --- |
| `AGENTS.md` | Agent 短索引 |
| `docs/ai/runtime-rule-index.md` | 运行规则和优先级 |
| `docs/ai/status/current.yaml` | 当前状态和当前 slice 指针 |
| `docs/ai/tasks/current-slice.yaml` | 当前执行 slice 定义 |
| `docs/ai/tasks/forest.yaml` | 主任务森林 |
| `docs/ai/requirements/ledger.yaml` | 需求台账 |
| `docs/ai/decisions/records.md` | 决策记录 |
| `docs/ai/evidence/index.md` | 证据索引 |
| `docs/ai/knowledge/index.md` | 知识库索引 |
| `docs/ai/handoff/current.md` | 当前交接说明 |

## 与人类文档的关系

`docs` 下的人类文档说明 Go 脚手架如何工作。`docs/ai` 下的运行态文档说明 Agent 认为当前任务状态是什么。两者需要在当前事实上一致，但服务对象不同。

不要让未来 Agent 依赖原始 prompt 或聊天记录恢复正常状态。如果状态缺失，应新增或修复物理运行态 artifact。

## 当前需关注风险

当前 status 引用了 `slice_011_go_scaffold_engineering_repair`，但部分 task/evidence/requirement 索引可能没有完整列出该 slice。修复前应把它视为运行态文档债。
