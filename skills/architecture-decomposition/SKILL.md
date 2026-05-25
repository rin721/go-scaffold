---
name: architecture-decomposition
description: Decompose confirmed requirements into architecture options, module boundaries, interface boundaries, decisions, and risks for go-scaffold. Use after requirements are confirmed and before task decomposition or implementation.
---

# Skill: architecture-decomposition

## Purpose

这个 skill 用来生成或修正架构方案、模块边界、接口边界、测试策略、安全策略和技术取舍。

## When to Use

在以下情况必须使用：

- 需求已确认，需要进入架构设计。
- 模块职责、依赖方向或接口边界不清。
- 需要比较多个技术方案或记录架构决策。

## Inputs

必须读取：

- `AGENTS.md`
- `REQUIREMENTS.md`
- `ACCEPTANCE.md`
- `RISK_REGISTER.md`
- `STATUS.md`

可选读取：

- `README.md`
- `go.mod`
- `MODULES.md`
- `DECISIONS.md`
- 相关源码目录。

## Outputs

必须写入或更新：

- `ARCHITECTURE.md`
- `MODULES.md`
- `ROADMAP.md`
- `DECISIONS.md`
- `RISK_REGISTER.md`
- `STATUS.md`

## Preconditions

执行前必须满足：

- P0 需求都有验收标准。
- 关键需求冲突已解决或被标记为阻塞。

## Procedure

1. 读取确认需求和验收标准。
2. 识别系统边界、非目标和外部依赖。
3. 生成至少 2 个方案，除非需求已指定唯一方案。
4. 比较复杂度、成本、风险、扩展性、测试难度和维护难度。
5. 推荐一个方案并记录被放弃方案。
6. 定义模块、数据流、接口边界、错误处理、测试、安全和部署策略。
7. 将决策写入 `DECISIONS.md`，风险写入 `RISK_REGISTER.md`。

## Acceptance Criteria

只有满足以下条件才算本 skill 执行完成：

- 推荐方案明确。
- P0 需求能映射到模块。
- 模块边界和非目标清晰。
- 测试策略、安全策略和风险登记存在。

## Completion Decision

- `COMPLETED`：架构已确认并记录。
- `PENDING_USER_CONFIRMATION`：关键技术取舍待确认。
- `BLOCKED`：核心需求或约束冲突未解决。
- `REWORK_REQUIRED`：架构与需求、验收或现有代码冲突。

## Failure Handling

如果失败：

1. 标记冲突。
2. 提供可选方案。
3. 等待用户确认关键取舍。

最大修复次数：

- 同一问题最多 3 轮。

## Evidence Required

必须记录：

- 修改文件。
- 方案比较。
- 决策记录。
- 风险更新。
- 下一步。
