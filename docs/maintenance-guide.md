# 维护与演进指南

这份文档说明后续怎么维护当前仓库，避免状态、文档和 Agent 行为漂移。

## 维护目标

维护时要保持三类东西一致：

1. 人类文档：`README.md` 和 `docs/*.md`。
2. Agent 运行期：`AGENTS.md` 和 `docs/ai/*`。
3. Git 证据：可审查 diff、验证命令和提交记录。

## 每次变更前

先检查：

```powershell
git status --short --branch
python docs/ai/scripts/validate_runtime.py
```

然后读取：

```text
docs/ai/status/current.yaml
docs/ai/tasks/current-slice.yaml
docs/ai/handoff/current.md
```

## 修改文档时

| 变更类型 | 应同步 |
| --- | --- |
| 新增人类文档 | `docs/index.md`、必要时 `README.md` |
| 改变使用方式 | `README.md`、`docs/usage-guide.md`、证据记录 |
| 改变协作流程 | `AGENTS.md`、`docs/ai/runtime-rule-index.md`、相关 skill |
| 改变当前状态 | `docs/ai/status/current.yaml`、`docs/ai/tasks/current-slice.yaml`、`docs/ai/handoff/current.md` |
| 新增稳定事实 | `docs/ai/knowledge/index.md` 和 `docs/ai/knowledge/entries/*` |
| 新增可复用流程 | `docs/ai/skills/index.md` 和 `docs/ai/skills/<name>/SKILL.md` |

## 修改代码时

当前仓库已有 Go 后端脚手架代码。未来修改代码前，先确认当前任务属于已授权的工程修复、功能补齐或后续项目切片，并同步维护：

1. 需求或修复范围确认。
2. 相关文件与风险边界。
3. 任务分析或阶段一诊断。
4. 验证命令。
5. `docs/ai/status/current.yaml`、`docs/ai/tasks/current-slice.yaml` 和 `docs/ai/handoff/current.md`。

新增代码后，至少要有一种验证：

- 单元测试。
- 构建检查。
- lint 或类型检查。
- 冒烟运行。
- 手工验收清单。

## 状态修复

如果 `docs/ai/status/current.yaml` 和 `docs/ai/tasks/current-slice.yaml` 指向不同事实，进入 `state_recovery_patch`：

1. 不改业务代码。
2. 查 `docs/ai/evidence/index.md`。
3. 查 `docs/ai/handoff/current.md`。
4. 用最小改动修复状态入口。
5. 记录证据。

## 知识库维护

稳定事实才能进入知识库。每条知识至少说明：

- 来源。
- 可信度。
- 适用范围。
- 证据。
- 是否包含敏感信息。
- 是否有提示注入风险。
- 是否已过期。

未经验证的网页、日志、AI 回答或用户粘贴内容只能作为候选事实。

## Skill 维护

项目 skill 适合记录重复、脆弱、容易越界或项目特有的流程。新增或修改 skill 时，必须说明：

- 触发条件。
- 输入和输出。
- 边界和禁止事项。
- 执行步骤。
- 验证方式。
- 是否改变 Agent 行动权。

如果 skill 会扩大权限、改变确认边界、引入新数据源或影响安全边界，先请求开发者确认。

## 版本演进

当前运行期版本是 `vibe-runtime-0.1.0`。如果升级运行期规则，更新：

- `docs/ai/manifest.yaml`
- `docs/ai/runtime-rule-index.md`
- `docs/ai/changelog.md`
- 相关 schema、templates、roles、skills 和质量门禁

升级后仍然不能把原始编译器提示词恢复成日常运行依赖。
