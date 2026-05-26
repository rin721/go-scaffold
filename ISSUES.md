# ISSUES.md

## Issue 状态

- 项目：go-scaffold
- 最后更新：2026-05-26
- 规则：失败、返工和阻塞问题记录在本文；风险项仍记录在 `RISK_REGISTER.md`。

## Open Issues

| ID | Linked Task | Severity | Status | Summary | Next Action |
|---|---|---|---|---|---|
| ISSUE-P2-005 | TASK-P2-004 | Medium | OPEN | 当前环境未安装 Docker CLI，无法验证 `Dockerfile` 镜像构建 | 在具备 Docker CLI/daemon 的 Linux 或 Docker Desktop 环境运行 `docker build -t go-scaffold:local .` |

## Issue Details

- ISSUE-P2-005：TASK-P2-004 已补齐 Dockerfile、production Compose 示例、production 配置样例、远程 Linux 动态 env 部署脚本和手动 production workflow 闸门；临时 Go YAML 解析、actionlint、全量 Go 回归、server build 和 `git diff --check` 均通过。但当前本机缺少 `docker`，且未发现 `podman`、`nerdctl` 或 `docker.exe`，因此 `docker build -t go-scaffold:local .` 无法执行。任务保持 `PENDING_VERIFICATION`，待 Docker 环境补跑构建。
- ISSUE-INFRA-002：`AGENTS.md` 缺失但状态文件声称已补齐。已在 TASK-INFRA-002 中修复，诊断报告见 `docs/reports/status_diagnostics/2026-05-25-task-infra-002-agents-md-missing.md`。
- ISSUE-P1-002：`.env.example` 与数据库环境变量前缀不一致，且 JWT 示例暗示未实现能力。已在 TASK-P1-002 中修复，相关测试通过。
- ISSUE-P1-003：无新增失败项。TASK-P1-003 新增 router smoke test 后，包测试和全量回归均通过。
- ISSUE-P1-004：无新增失败项。TASK-P1-004 新增 demo CRUD 测试基线后，demo 模块测试和全量回归均通过。
- ISSUE-P1-005：无新增失败项。TASK-P1-005 收拢 demo 迁移边界后，app 包测试和全量回归均通过。
- ISSUE-P1-006：无新增失败项。TASK-P1-006 收拢 CLI tests 命令语义后，cmd/server 包测试和全量回归均通过。
- ISSUE-P1-007：无新增失败项。TASK-P1-007 完成 `pkg/*` API 分类后，全量回归通过。
- ISSUE-P1-008：无新增失败项。TASK-P1-008 标注 `pkg/sqlgen` unsupported 边界后，包测试和全量回归均通过。
- ISSUE-NEXT-001：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE 已将 `BL-021` / `TM-P1-005` 提升为 TASK-P1-009。
- ISSUE-P1-010：无新增失败项。TASK-P1-010 收拢 `pkg/plugin` 被动注册边界后，包测试和全量回归均通过。
- ISSUE-NEXT-003：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE-003 已将 `BL-020` 首批提升为 TASK-P1-011 / TS-P1-011。
- ISSUE-P1-011：无未解决失败项。TASK-P1-011 中 `pkg/yaml2go` 新增测试暴露生成 tag 与方法 import 顺序缺陷，已在当前允许范围内修复并通过回归。
- ISSUE-NEXT-004：无新增失败项。用户发送“下一步”后，TASK-NEXT-SCOPE-004 已将 `BL-020` 第二批提升为 TASK-P1-012 / TS-P1-012。
- ISSUE-P1-012：无未解决失败项。TASK-P1-012 中 `pkg/executor` 新增测试暴露 sentinel 错误包装和 panic handler 未调用缺陷，已在当前允许范围内修复并通过回归；测试自身等待竞态已在第二轮修复。
- ISSUE-NEXT-005：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE-005 已将 `BL-020` 第三批 `pkg/cache` 隔离行为测试提升为 TASK-P1-013 / TS-P1-013。
- ISSUE-P1-013：无未解决失败项。TASK-P1-013 首次包测试为测试代码编译失败，原因是误读 `miniredis.Get` 返回值；修正测试断言后，`pkg/cache` 包测试和全量回归均通过。
- ISSUE-NEXT-006：无新增失败项。用户选择 B 后，TASK-NEXT-SCOPE-006 已将 `BL-023` `pkg/utils` 内部支撑测试提升为 TASK-P1-014 / TS-P1-014。
- ISSUE-P1-014：无未解决失败项。TASK-P1-014 前两次包测试失败来自测试代码对端口占用语义的环境假设；改为确定性无效地址、端口范围和 exclude 断言后，`pkg/utils` 包测试和全量回归均通过。
- ISSUE-NEXT-007：无新增失败项。用户选择 B 后，TASK-NEXT-SCOPE-007 已将 `BL-002` router/middleware/demo HTTP 集成测试提升为 TASK-P1-015 / TS-P1-015。
- ISSUE-P1-015：无未解决失败项。TASK-P1-015 前两次相关包测试失败来自测试构造问题：`httptest.NewRequest` 默认 Host 与 Origin 同源导致 CORS 中间件跳过；固定测试 Host 为 `api.local` 后，相关包测试和全量回归均通过。
- ISSUE-NEXT-008：无新增失败项。用户选择 A 后，TASK-NEXT-SCOPE-008 已关闭并进入 TASK-PHASE6-001 / TS-PHASE6-001。
- ISSUE-PHASE6-001：无未解决失败项。Phase 6 收尾仅更新项目状态文档，最终 `go test ./... -count=1` 与 `git diff --check` 均通过。
- ISSUE-P1-016：无未解决失败项。TASK-P1-016 新增 app 装配与 reload/config 集成测试后，`go test ./internal/app/... -count=1`、`go test ./... -count=1` 与 `git diff --check` 均通过。
- ISSUE-P1-017：无未解决失败项。TASK-P1-017 第一阶段包 README 中文化后，`go test ./... -count=1` 与 `git diff --check` 均通过。
- ISSUE-INFRA-003：TASK-P1-016/017 完成后部分背景文档仍保留旧待办表述。已在 TASK-INFRA-003 中修复，诊断报告见 `docs/reports/status_diagnostics/2026-05-26-task-p1-017-post-completion-doc-drift.md`。
- ISSUE-P2-001：无未解决阻塞失败项。TASK-P2-001 新增 CI 质量门禁与部署说明后，全量测试、server 构建和 `git diff --check` 均通过；gofmt 漂移审计发现 82 个历史格式漂移文件，已记录 `BL-025`。
- ISSUE-P2-002：无实现失败项。用户选择 C 并确认使用远程部署后，当时真实 CD 自动化仍因缺少镜像仓库、SSH/Docker 等远程方式、环境、触发策略和 secrets 命名而处于 `PENDING_USER_CONFIRMATION`；后续已由 TASK-P2-003 完成手动 staging workflow。
- ISSUE-P2-003：无新增失败项。TASK-P2-002 仅新增 `.env.deploy.example` 和忽略真实 `.env.deploy`，当时未实现 workflow、未连接远程环境、未写入密钥；后续已由 TASK-P2-003 完成手动 staging workflow。
- ISSUE-P2-004：无新增失败项。TASK-P2-003 新增手动 staging 远程部署 workflow 后，临时 Go YAML 解析、actionlint 和 `git diff --check` 均通过；未执行真实部署、未连接远程服务器、未写入真实密钥。

## 历史说明

- 2026-05-25：记录并关闭 `.env.example` 与数据库环境变量实现不一致问题。
- 2026-05-25：记录 TASK-P1-003 无新增失败项，HTTP router smoke test 和全量回归通过。
- 2026-05-25：记录 TASK-P1-004 无新增失败项，demo CRUD 测试和全量回归通过。
- 2026-05-25：记录 TASK-P1-005 无新增失败项，demo 迁移策略测试和全量回归通过。
- 2026-05-25：记录 TASK-P1-006 无新增失败项，CLI tests 命令语义测试和全量回归通过。
- 2026-05-25：记录 TASK-P1-007 无新增失败项，`pkg/*` API 分类后全量回归通过。
- 2026-05-25：记录 TASK-P1-008 无新增失败项，`pkg/sqlgen` unsupported 行为测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE 无新增失败项，`types/*` 契约边界已提升为下一合法任务。
- 2026-05-25：记录 TASK-P1-010 无新增失败项，`pkg/plugin` 被动注册边界测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE-003 无新增失败项，首批 `pkg/*` 行为测试已排期。
- 2026-05-25：记录 TASK-NEXT-SCOPE-004 无新增失败项，第二批 `pkg/*` 行为测试已排期。
- 2026-05-25：记录 TASK-P1-012 无未解决失败项，`pkg/executor` 暴露缺陷已修复，第二批包测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE-005 无新增失败项，第三批 `pkg/cache` 行为测试已排期。
- 2026-05-25：记录 TASK-P1-013 无未解决失败项，测试代码编译问题已修复，`pkg/cache` 包测试和全量回归通过。
- 2026-05-25：记录 TASK-NEXT-SCOPE-006 无新增失败项，`pkg/utils` 内部支撑测试已排期。
- 2026-05-25：记录 TASK-P1-014 无未解决失败项，测试环境假设已修正，`pkg/utils` 包测试和全量回归通过。
- 2026-05-26：记录 TASK-NEXT-SCOPE-007 无新增失败项，router/middleware/demo HTTP 集成测试已排期。
- 2026-05-26：记录 TASK-P1-015 无未解决失败项，CORS 测试构造问题已修正，相关包测试和全量回归通过。
- 2026-05-26：记录 TASK-NEXT-SCOPE-008 与 TASK-PHASE6-001 无新增失败项，Phase 6 收尾完成。
- 2026-05-26：记录 TASK-P1-016 无未解决失败项，app 装配与 reload/config 集成测试通过。
- 2026-05-26：记录 TASK-P1-017 无未解决失败项，第一阶段包 README 中文化和全量回归通过。
- 2026-05-26：记录并关闭 TASK-P1-016/017 后背景文档状态漂移。
- 2026-05-26：记录 TASK-P2-001 无未解决失败项，CI 质量门禁和部署说明首切片完成。
- 2026-05-26：记录 TASK-P2-003 无新增失败项，手动 staging 远程部署 workflow 通过静态验证。
- 2026-05-26：记录 TASK-P2-004 环境待验证项，当前机器无 Docker CLI，Docker build 待补跑；其他静态验证和 Go 回归通过。
- 2026-05-25：记录并关闭 `AGENTS.md` 缺失导致的 Agent 入口冲突。
- 2026-05-25：创建 `ISSUES.md`，补齐 `docs/ai/prompt.md` 要求的项目问题记录入口。
