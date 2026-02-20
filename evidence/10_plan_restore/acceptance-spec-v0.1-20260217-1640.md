# 验收清单 & 证据规格（锁定版）— Checkpoint A

- Project: 一人软件服务公司自动化交付系统（Delivery System）
- Owner: Engineer
- Version: v0.1 (LOCKED)
- Timestamp (Asia/Shanghai): 2026-02-17 16:40

> 说明：本文件为 `docs/09-ACCEPTANCE_AND_EVIDENCE_SPEC.md` 的“锁定执行版”。
> 锁定含义：后续如需变更验收口径或证据最小集合，必须以“变更记录+新版本”形式落盘，并在 00-index 记录变更指针。
> 命名/目录/索引字段：遵循 `evidence/LOCK.md`。

## A. 通用DONE标准（适用于所有条目）
每条验收/证据达到 DONE 时，必须满足：
1) 一跳定位：从 `evidence/00-index.md` 可直接定位到证据文件
2) 可复核：证据内版本号/时间点与索引登记一致
3) 可追溯：至少具备关联指针（决议 → 行动项 → 风险），或明确说明不适用

## B. Checkpoint A（Plan Restore）验收项
### B0. 五件套落盘 + 索引DONE（本轮P0）
- 验收：
  - 五件套文件存在且版本化
  - 00-index对应条目为DONE
  - 00-index备注中包含：路径/版本号/时间点/责任人/关联指针/验证记录
- 证据最小集合：
  - evidence/10_plan_restore/pm-min-input-*.md
  - evidence/10_plan_restore/acceptance-spec-*.md
  - evidence/10_plan_restore/wp-list-*.md
  - evidence/10_plan_restore/go-nogo-checkpoint-a-*.md
  - evidence/10_plan_restore/resume-directive-*.md

## C. T3-Gate（下一阶段，解冻后首批）
### C1. 硬证据#1：可签字评审决议包（evidence/01_review/）
- 必备文件：
  - review-record-*.md/pdf/png
  - feedback-list-*.md
  - decision-items-*.md（可签字）
  - signoff-*.pdf/png/邮件截图
  - consistency-check-*.md（版本号/时间点一致）
- 验收：文件齐备、可执行、可签字；00-index对应条目DONE且附验证记录。

### C2. 硬证据#2：VPS采购确认证据包（evidence/02_procurement/）
- 必备文件：
  - order-confirmation-*
  - delivery-time-proof-*
  - spec-region-*
  - payment-proof-*
- 验收：信息完整可复核；交付时间可追溯；00-index条目DONE且附验证记录。

### C3. 硬证据#3：端到端归档演练报告（evidence/09_archiving_drill/）
- 必备内容：
  - 覆盖硬证据#1/#2全链路：采集→命名→关联→索引更新→TODO→DONE→可追溯性验证
  - 问题清单与修复闭环（至少闭环到可追溯达标）
- 验收：报告可复现；索引更新过程可证明；00-index条目DONE且附验证记录。
