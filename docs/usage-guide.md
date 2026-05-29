# 使用教程

这份教程讲的是怎么“食用”当前仓库：你可以把它当成一个已经装好 Vibe Coding 协作底座的项目空壳，用来继续生成业务需求、文档、代码和验证流程。

## 1. 第一次打开仓库

先看三个文件：

```text
README.md
docs/index.md
AGENTS.md
```

它们分别回答：

- 这个仓库现在是什么状态。
- 文档应该从哪里读。
- Agent 应该从哪里恢复上下文。

## 2. 检查当前状态

让 Agent 或自己查看：

```text
docs/ai/status/current.yaml
docs/ai/tasks/current-slice.yaml
docs/ai/handoff/current.md
```

这三个文件回答：

- 当前在哪个阶段。
- 当前唯一合法执行切片是什么。
- 下一步条件是什么。

## 3. 运行基础验证

在终端里运行：

```powershell
python docs/ai/scripts/validate_runtime.py
```

看到下面输出就说明运行期入口完整：

```text
Runtime validation passed.
```

## 4. 给出下一轮目标

你可以这样告诉 Agent：

```text
下一步：我想把这个仓库做成一个本地任务管理工具。
```

或者：

```text
下一步：先帮我搭一个前端原型，能展示项目列表和任务状态。
```

Agent 应该先把你的想法记录进需求入口，然后用少量问题确认：

1. 谁会用它？
2. 第一版必须完成什么？
3. 哪些明确不做？
4. 有没有账号、隐私、资金、生产或外部服务风险？
5. 验收时你希望看到什么？

## 5. 不要直接跳到写代码

这个仓库的协作约定是先确认需求，再做研究、任务分析、架构、模式和切片。低风险任务可以很轻，但不能没有状态和证据。

正确顺序：

```text
目标
-> 需求基线
-> 研究或说明无需研究
-> 任务分析
-> 架构与技术栈
-> 模式确认
-> 任务树与切片
-> 实现
-> 验证和证据
```

## 6. 常见操作

### 查看当前入口

```powershell
Get-Content -Raw AGENTS.md
Get-Content -Raw docs/ai/status/current.yaml
Get-Content -Raw docs/ai/tasks/current-slice.yaml
```

### 查看所有文档

```powershell
Get-ChildItem -Recurse -File docs | Select-Object FullName
```

### 搜索规则或状态

```powershell
rg -n "round_001|current_slice|no_prompt_runtime_dependency" AGENTS.md docs
```

如果没有 `rg`，可以用 PowerShell：

```powershell
Get-ChildItem -Recurse -File docs | Select-String -Pattern "round_001|current_slice|no_prompt_runtime_dependency"
```

## 7. 文档应该怎么改

- 面向人的说明放在 `docs/`。
- Agent 状态和证据放在 `docs/ai/`。
- 改了流程、运行方式或重要入口，就更新 `README.md` 和 `docs/index.md`。
- 改了 Agent 执行边界，就更新 `AGENTS.md`、`docs/ai/runtime-rule-index.md` 和相关 skill。
- 做了验证，就把结果写到 `docs/ai/evidence/index.md`。

## 8. 什么时候算完成

一个切片完成至少要满足：

- 文件已经修改。
- 对应验证已经运行，或有明确替代验证。
- 证据已经记录。
- 状态和交接已经同步。
- 没有把未确认的推测写成开发者确认。

当前文档切片的完成标准是：

- `README.md` 存在并能快速说明项目。
- `docs/index.md` 存在并能导航所有人类文档。
- `docs/project-overview.md` 存在并说明项目结构和运行期基建。
- `docs/usage-guide.md` 存在并给出实际使用步骤。
- `docs/maintenance-guide.md` 存在并说明后续维护方式。
- `docs/agent-workflow-guide.md` 存在并说明人机协作流程。
- 运行期验证通过。
