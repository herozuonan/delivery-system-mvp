# T3-Lock：证据归档强约束（唯一执行标准）

> 强约束：**无索引不算完成**。任何证据文件若未在 `evidence/00-index.md` 登记（含路径/版本/时间点/责任人/状态），一律视为未交付。

## 1. 目录结构（固定）
不允许随意新增顶层目录；新增内容必须落在既有分类下。

- `evidence/01_review/`：评审决议包（硬证据#1）
- `evidence/02_procurement/`：采购确认证据包（硬证据#2）
- `evidence/03_vps_init/`：VPS初始化记录
- `evidence/04_script_v2/`：脚本v2执行/问题单/验证报告
- `evidence/05_observability/`：监控告警配置/演练记录
- `evidence/06_cicd_rollback/`：CI/CD与回滚演练证据
- `evidence/07_env_validation/`：环境验证清单（逐项附证据）
- `evidence/08_baseline_freeze/`：基线冻结证据（tag/commit）
- `evidence/09_archiving_drill/`：端到端归档演练报告
- `evidence/templates/`：模板

## 2. 命名规范（强制）
文件名必须包含：`YYYYMMDD` 日期 + 主题 + owner（或团队角色），必要时加版本号。

推荐格式：
- `YYYYMMDD-<topic>-<owner>[-vX].<ext>`

示例：
- `20260216-review-record-engineer-v1.md`
- `20260216-decision-items-techlead-v1.md`
- `20260216-procurement-confirmation-pm.png`
- `20260216-archiving-drill-report-engineer-v1.md`

时间戳格式：
- 文档内使用 ISO-8601：`2026-02-16T12:34:56+08:00`

## 3. 版本号策略（证据集一致）
- 同一里程碑/同一证据包，版本号统一（如 v1）。
- 评审决议包、归档演练报告、索引更新应能对齐到同一版本与时间点。

## 4. 索引字段（必须）
`evidence/00-index.md` 的每条证据必须包含：
- 类别
- 状态：`TODO | IN_PROGRESS | DONE`
- 证据文件路径或链接
- 版本号（如 v1）
- 时间点（ISO-8601）
- 责任人（owner）
- 关联指针（可选但推荐：指向决议项/行动项/风险条目）

## 5. 索引更新流程（强制）
- 产出证据文件 **当日** 必须同步更新索引。
- 索引状态变更规则：
  - `TODO`：未产出文件
  - `IN_PROGRESS`：已产出部分文件/待签字/待补齐
  - `DONE`：证据包完整且可追溯性校验通过

## 6. 可追溯性最低要求
至少满足：
- 决议项（decision-items）→ 反馈清单（feedback-list）→ 风险条目（RISKS编号或风险说明）
并能从索引一跳定位到对应文件。
