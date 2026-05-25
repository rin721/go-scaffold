# 状态诊断报告

## 1. 诊断结论

已发现并修复 Agent 基础设施状态冲突：`STATUS.md`、`TASKS.md`、`TEST_REPORT.md` 和 `AGENT_HANDOFF.md` 声称 `AGENTS.md` 已补齐，但仓库根目录实际缺失该文件。

## 2. 原因

- `TASK-INFRA-001` 记录了跨 Agent 入口补齐，但 `AGENTS.md` 未实际落盘。
- `CLAUDE.md`、`.cursor/rules/agent_project_driver.mdc`、`.kiro/steering/agent_project_driver.md`、`.codex/config.toml` 均引用 `AGENTS.md`，导致新 Agent 无法按入口恢复项目规则。
- `skills/*/SKILL.md` 缺少 YAML frontmatter 和完整项目 skill 结构，无法通过标准 skill 验证。

## 3. 读取到的状态

- 当前阶段：配置环境变量策略收拢。
- 当前模块：internal/config。
- 当前任务：TASK-P1-002。
- 当前时间切片：TS-P1-002。
- 阻塞项：Agent 入口文件缺失，属于流程基础设施冲突。
- 待验证项：基础设施一致性、skill frontmatter、Go 全量回归。

## 4. 冲突或缺失

| 文件 | 问题 | 影响 |
|---|---|---|
| `AGENTS.md` | 实际缺失 | 跨工具主入口断裂 |
| `STATUS.md` | 记录 `AGENTS.md` 已补齐 | 状态与文件系统事实冲突 |
| `CLAUDE.md` / Cursor / Kiro / Codex 配置 | 引用缺失入口 | 新 Agent 恢复失败 |
| `skills/*/SKILL.md` | 缺少 frontmatter 和完整 skill 模板结构 | skill 验证失败，触发和交接不稳定 |
| `docs/templates/*` | 部分模板混入当前项目实例事实 | 模板复用性不足 |

## 5. 修复方案

| 方案 | 做法 | 风险 | 推荐度 |
|---|---|---|---|
| A | 只新增 `AGENTS.md` | skills 和模板仍不完整 | 中 |
| B | 完整补齐入口、skills、模板、适配器和状态记录 | 修改文档较多，但不触碰 Go 代码 | 高 |

## 6. 推荐下一步

- 已采用方案 B。
- 完成 `TASK-INFRA-002 / TS-INFRA-002` 后，恢复当前合法下一步为 `TASK-P1-002 / TS-P1-002`。
