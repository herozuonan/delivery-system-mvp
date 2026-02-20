# 恢复推进指令（执行准入依据）— Resume Directive

- Project: 一人软件服务公司自动化交付系统（Delivery System）
- Owner: Engineer
- Version: v0.1
- Effective Timestamp (Asia/Shanghai): 2026-02-17 16:40
- Prerequisite: Checkpoint A = GO 且 00-index 对应条目 DONE

## 1) 解冻条件（必须同时满足）
1) Checkpoint A Go/No-Go结论为 GO（引用#1-#3版本号）
2) `evidence/00-index.md` 中“计划恢复/Checkpoint A”五件套条目均为 DONE
3) DONE条目备注字段齐全：路径/版本号/时间点/责任人/关联指针/验证记录

## 2) 生效范围
- 生效对象：从本指令生效起的所有工程活动（含文档、证据、代码、环境、配置、发布）
- 生效目标：以“证据可验收”为准入门槛，持续推进直到上线

## 3) 允许的增量边界（按阶段）
### 3.1 阶段一：T3硬证据链（优先）
- 允许：为获得硬证据#1/#2/#3而进行的资料采集、文档落盘、归档演练与索引更新
- 禁止：任何不产生可验收证据的忙碌工作；任何技术实现/环境变更

### 3.2 阶段二：技术交付（在T3-Gate通过后）
- 允许：按WP依赖顺序执行 VPS init / script v2 / observability / cicd&rollback / env validation / baseline freeze / launch
- 禁止：跳过证据与验收；未更新00-index就宣称完成；绕开回滚要求

## 4) Gate顺序（必须遵循）
- Gate-Checkpoint A（当前） → Gate-T3（硬证据#1/#2/#3全部DONE） → 技术交付WP序列 → 基线冻结 → 上线

## 5) 每日检查机制（执行纪律）
- 每日开始：确认当日目标WP、预期证据路径、00-index需要更新的条目
- 每日结束：更新00-index（TODO/IN_PROGRESS/DONE）、写入验证记录；若未达标，必须记录原因与下一步补救

## 6) 回滚与违例处理
- 若出现冻结策略违例（未满足解冻条件即实施技术增量）：
  - 立即停止
  - 产出回滚证据
  - 在00-index新增/更新“违例与回滚”验证记录
  - 决策默认切回 No-Go，直到审计链恢复
