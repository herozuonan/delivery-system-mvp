# Evidence Index

**T3-Lock（强约束来源）：** evidence/LOCK.md（无索引不算完成）

| 类别 | 状态 | 证据文件/链接 | 备注 |
|---|---|---|---|
| T3-Lock 规范 | IN_PROGRESS | evidence/LOCK.md | 命名规范/目录结构/索引字段/更新流程 |
| 计划恢复（PM最小输入包） | DONE | evidence/10_plan_restore/pm-min-input-v0.1-20260217-1640.md | v0.1 @2026-02-17 16:40 Owner=Engineer；关联：Checkpoint A GO→T3硬证据链；验证：一跳定位/版本时间戳在文件头一致 |
| 计划恢复（验收&证据规格） | DONE | evidence/10_plan_restore/acceptance-spec-v0.1-20260217-1640.md | v0.1(LOCKED) @2026-02-17 16:40 Owner=Engineer；关联：作为后续Gate判定依据；验证：DONE标准与证据最小集合已固化 |
| 计划恢复（WP清单） | DONE | evidence/10_plan_restore/wp-list-v0.1-20260217-1640.md | v0.1 @2026-02-17 16:40 Owner=Engineer；关联：覆盖T3→技术交付→上线；验证：依赖总览D0-D5可形成执行顺序 |
| 计划恢复（Go/No-Go） | DONE | evidence/10_plan_restore/go-nogo-checkpoint-a-v0.1-20260217-1640.md | v0.1 @2026-02-17 16:40 Owner=Engineer；关联：引用#1-#3版本号；验证：Decision=GO + 冻结策略符合性核验记录已写入 |
| 计划恢复（恢复推进指令） | DONE | evidence/10_plan_restore/resume-directive-v0.1-20260217-1640.md | v0.1 @2026-02-17 16:40 Owner=Engineer；关联：解冻准入与Gate顺序；验证：明确解冻条件/允许边界/违例处理 |
| 三方评审记录 | DONE | evidence/01_review/20260220-review-record-async-engineer-v1.md | v1 @2026-02-20T11:30:00+08:00 Owner=Engineer；验证：书面确认闭环 |
| 反馈清单 | DONE | evidence/01_review/20260220-feedback-list-async-engineer-v1.md | v1 @2026-02-20T11:30:00+08:00 Owner=Engineer |
| 决议项（可签字） | DONE | evidence/01_review/20260220-decision-items-async-techlead-v1.md | v1 @2026-02-20T11:30:00+08:00 Owner=Engineer/TechLead |
| 签字页/书面确认 | DONE | evidence/01_review/20260220-written-confirmation-user-v1.md; evidence/01_review/20260220-written-confirmation-techlead-v1.md | v1 @2026-02-20T11:32:00+08:00 Owner=User/TechLead；验证：两方书面确认齐备 |
| 版本号/时间戳一致性校验 | DONE | evidence/01_review/20260220-consistency-check-01_review-v1.md | v1 @2026-02-20T11:33:00+08:00 Owner=Engineer；校验：一跳定位/版本/时间戳 |
| VPS采购确认单 | IN_PROGRESS | evidence/02_procurement/20260220-vps-equivalent-evidence-package-v1.md | v1 @2026-02-20T11:41:00+08:00 Owner=Engineer；等效证据包（非新采购） |
| 交付时间证据 | IN_PROGRESS | evidence/02_procurement/20260220-vps-equivalent-evidence-package-v1.md | v1 @2026-02-20T11:41:00+08:00 Owner=Engineer；待补齐控制台截图时间戳 |
| 规格/地域说明 | IN_PROGRESS | evidence/02_procurement/20260220-vps-capture-instructions-v1.md | v1 @2026-02-20T11:41:00+08:00 Owner=Engineer；待补齐控制台截图/采集输出 |
| 付款凭证 | N/A | evidence/02_procurement/20260220-vps-equivalent-evidence-package-v1.md | v1 @2026-02-20T11:41:00+08:00 Owner=Engineer；等效证据包不要求 |
| 端到端归档演练报告 | IN_PROGRESS | evidence/09_archiving_drill/20260220-archiving-drill-report-v1.md | v1 @2026-02-20T11:42:00+08:00 Owner=Engineer；待闭环#2材料与索引DONE |
| VPS初始化记录 | TODO | evidence/03_vps_init/ | 安全加固/防火墙 |
| 脚本v2执行日志 | TODO | evidence/04_script_v2/ | 成功/失败路径 |
| 脚本v2验证报告 | TODO | evidence/04_script_v2/ | 幂等/可重入/回滚 |
| 监控部署与配置 | TODO | evidence/05_observability/ | compose/config/export |
| 告警演练记录 | TODO | evidence/05_observability/ | 宕机/资源耗尽 |
| CI/CD流水线证据 | TODO | evidence/06_cicd_rollback/ | run link/export |
| 回滚演练证据 | TODO | evidence/06_cicd_rollback/ | 版本切换/健康 |
| 环境验证清单 | TODO | evidence/07_env_validation/ | 逐项附证据 |
| 基线冻结（tag/版本） | TODO | evidence/08_baseline_freeze/ | 变更流程生效 |
