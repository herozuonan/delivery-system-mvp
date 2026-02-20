# Evidence Vault（证据库）

本目录用于“证据化交付”归档，确保每个关键结论、采购、环境验证、演练均可追溯到时间点与版本号。

## 目录结构
- 01_review/            三方评审证据（记录/反馈/决议/签字）
- 02_procurement/       采购与交付时间证据（确认单/订单/截图/工单）
- 03_vps_init/          VPS初始化与安全加固记录
- 04_script_v2/         初始化脚本v2执行日志/问题单/修复记录/验证报告
- 05_observability/     监控告警配置与演练证据（Prometheus/Grafana/Alertmanager）
- 06_cicd_rollback/     CI/CD流水线与回滚演练证据
- 07_env_validation/    环境验证清单（逐项附证据）
- 08_baseline_freeze/   基线冻结证据（tag/commit/版本号/变更流程生效说明）
- templates/            证据模板

## 归档原则
1. 每份证据都要包含：日期时间、Owner、关联版本（tag/commit）、外部链接（如有）与本地备份（截图/导出/日志）。
2. 关键路径门槛：
   - 硬证据#1：可签字评审决议（MVP/需求冻结、架构与技术栈批准、变更规则）
   - 硬证据#2：VPS采购确认单 + 明确交付时间
3. 文件命名建议：`YYYYMMDD-主题-Owner`，例如：`20260216-review-record-engineer.md`
