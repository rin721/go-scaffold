# Time Slice Template

## Time Slice Format

每个时间切片必须包含：

- Time Slice ID:
- Parent Task:
- Module:
- Purpose:
- Status:
- Inputs:
- Outputs:
- Allowed Files:
- Forbidden Files:
- Strict Non-Goals:
- Execution Steps:
- Test Commands:
- Verification Method:
- Acceptance Criteria:
- Completion Criteria:
- Failure Handling:
- Max Repair Attempts:
- Evidence Required:
- Next Slice Entry Conditions:

---

## Current Legal Time Slice

- Time Slice ID:
- Parent Task:
- Status:
- Why this is the only legal next slice:

---

## Time Slices

### TS-001: <title>

- Parent Task:
- Module:
- Purpose:
- Status: NOT_STARTED
- Inputs:
  -
- Outputs:
  -
- Allowed Files:
  -
- Forbidden Files:
  -
- Strict Non-Goals:
  - 不做任何未列入本切片的优化。
  - 不修改未授权文件。
  - 不新增未确认功能。
- Execution Steps:
  1.
  2.
- Test Commands:
  -
- Verification Method:
  -
- Acceptance Criteria:
  -
- Completion Criteria:
  -
- Failure Handling:
  - 同一失败最多修复 3 轮。
  - 仍失败则生成 issue report。
- Max Repair Attempts: 3
- Evidence Required:
  - 修改文件：
  - 命令：
  - 测试结果：
  - 验证结论：
- Next Slice Entry Conditions:
  -
