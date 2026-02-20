# 端到端归档演练报告（v1，草稿）

- Timestamp: 2026-02-20T11:42:00+08:00
- Owner: Engineer
- Scope: 硬证据#1（01_review）+ 等效硬证据#2（02_procurement）→ 索引 TODO→DONE → 可追溯性验证

## 1) 演练目标
验证：证据采集→命名→关联→索引更新→状态闭环→可复核/可追溯。

## 2) 演练输入
- 硬证据#1：evidence/01_review/（v1）
- 硬证据#2：evidence/02_procurement/（v1，待补齐控制台证据与规格采集输出）

## 3) 执行步骤（已完成）
1. 生成并落盘硬证据#1：review-record/decision-items/feedback-list/written-confirmation/consistency-check
2. 更新 evidence/00-index.md：将对应条目从 TODO→DONE/IN_PROGRESS
3. 生成硬证据#2 等效证据包模板与采集指令

## 4) 待完成步骤（阻塞点）
- 从 Owner 获取控制台截图（地域/规格/可用性）
- 获取 VPS 系统侧规格采集输出（txt）
- 将 evidence/00-index.md 中 #2 的条目置为 DONE，并补齐版本/时间点/Owner/验证记录

## 5) 可追溯性验证
- 决议（01_review decision-items）→ 行动项（补齐#2证据）→ 风险（密码暴露风险：要求轮换）

## 6) 结论
- 当前：IN_PROGRESS（#2材料未齐，演练未闭环）
- 完成标准：#2齐备 + 00-index 对应条目 DONE + 本报告补齐“验证记录”段落
