# Work Packages 清单（完成版）+ 依赖清单 — Checkpoint A

- Project: 一人软件服务公司自动化交付系统（Delivery System）
- Owner: Engineer
- Version: v0.1
- Timestamp (Asia/Shanghai): 2026-02-17 16:40
- Evidence rules: 遵循 `evidence/LOCK.md`

> 目标：把“连续推进直到上线”拆成可执行WP，并明确依赖顺序、验收与证据路径。

## 0) 依赖总览（可形成执行顺序）
- D0：Checkpoint A闭环（五件套 + 00-index DONE）
- D1：T3硬证据#1（评审决议包）
- D2：T3硬证据#2（VPS采购确认证据包）
- D3：T3硬证据#3（端到端归档演练报告，覆盖D1+D2链路）
- D4：T3-Gate通过（审计链成立）
- D5：技术交付执行（VPS init → script v2 → observability → cicd/rollback → env validation → baseline freeze → launch）

## WP-A0：Checkpoint A（五件套+索引DONE）（P0, Freeze-Only）
- Inputs: `docs/08..10` 模板 + `evidence/00-index.md`
- Outputs:
  - pm-min-input（完成版）
  - acceptance-spec（锁定版）
  - wp-list（本文件）
  - go-nogo-checkpoint-a
  - resume-directive
- DoD:
  - 五件套均版本化落盘
  - 00-index条目DONE且备注字段齐全（路径/版本/时间/责任人/关联/验证）
- Evidence Path: evidence/10_plan_restore/

## WP-T3-01：硬证据#1 评审决议包（P0）
- Inputs: acceptance-spec（锁定版）
- Outputs: evidence/01_review/ 下5件套（review-record/feedback-list/decision-items/signoff/consistency-check）
- DoD: 按验收规格C1；00-index对应条目DONE
- Dependencies: D0

## WP-T3-02：硬证据#2 VPS采购确认证据包（P0）
- Inputs: acceptance-spec（锁定版）
- Outputs: evidence/02_procurement/ 下4件套（订单/交付时间证据/规格地域/付款凭证）
- DoD: 按验收规格C2；00-index对应条目DONE
- Dependencies: D0

## WP-T3-03：端到端归档演练报告（P0）
- Inputs: WP-T3-01, WP-T3-02产物
- Outputs: evidence/09_archiving_drill/ 报告（含问题清单与修复闭环）
- DoD: 按验收规格C3；00-index对应条目DONE
- Dependencies: D0, D1, D2

## WP-TECH-01：VPS初始化记录（P0）
- Outputs: evidence/03_vps_init/（安全加固/防火墙/基础配置记录+日志）
- DoD: 00-index条目DONE；可复核/可追溯
- Dependencies: D4

## WP-TECH-02：脚本v2执行与验证（P0）
- Outputs: evidence/04_script_v2/（执行日志 + 验证报告：幂等/可重入/失败路径/回滚）
- Dependencies: D4, WP-TECH-01

## WP-TECH-03：监控与告警（P0）
- Outputs: evidence/05_observability/（部署与配置 + 告警演练记录）
- Dependencies: D4, WP-TECH-02

## WP-TECH-04：CI/CD与回滚演练（P0）
- Outputs: evidence/06_cicd_rollback/（流水线证据 + 回滚演练证据）
- Dependencies: D4, WP-TECH-02

## WP-TECH-05：环境验证清单（P0）
- Outputs: evidence/07_env_validation/（逐项验证与证据）
- Dependencies: D4, WP-TECH-03, WP-TECH-04

## WP-TECH-06：基线冻结（tag/版本）（P0）
- Outputs: evidence/08_baseline_freeze/（基线tag/版本/变更流程生效记录）
- Dependencies: D4, WP-TECH-05

## WP-TECH-07：上线推进（P0）
- Outputs: 上线证明（需要在后续补充上线证据目录/条目）
- Dependencies: D4, WP-TECH-06
