# PM 最小输入包（完成版）— Checkpoint A / Plan-to-Execute Gate

- Project: 一人软件服务公司自动化交付系统（Delivery System）
- Owner: Engineer
- Version: v0.1
- Timestamp (Asia/Shanghai): 2026-02-17 16:40
- Status intent: 用于Checkpoint A输入；缺字段则默认No-Go。

> 目的：用最小但完整的输入，解除“计划层面”阻塞，使项目从冻结态进入可控执行（先过T3-Gate硬证据链，再进入技术交付）。
> 硬规则：任一关键字段缺失或不可执行 => No-Go（继续冻结）。

## 1) 迭代目标（必须）
- Iteration Goal：完成T3阶段的“可审计交付链”前置（硬证据#1/#2/#3），并建立可持续的Evidence→Index→Gate闭环，使后续技术交付可被验收与追溯。
- 目标Gate/交付物：
  - Gate-Checkpoint A：五件套（本文件 + 验收&证据规格锁定版 + WP清单完成版 + Go/No-Go结论 + 恢复推进指令）落盘并在 `evidence/00-index.md` 登记为 DONE
  - Gate-T3：
    - 硬证据#1：可签字评审决议包（evidence/01_review/）
    - 硬证据#2：VPS采购确认证据包（evidence/02_procurement/）
    - 硬证据#3：端到端归档演练报告（evidence/09_archiving_drill/）
- 成功定义（可度量）：
  - `evidence/00-index.md` 相关条目从 TODO/IN_PROGRESS → DONE
  - 每个DONE条目满足：一跳定位 / 可复核（版本号+时间点一致）/ 可追溯（决议→行动→风险）

## 2) 范围边界（必须）
- In Scope：
  - Freeze-Only（证据/计划）工作：文档、证据、索引、验证记录、审计链条完善
  - 产出并固化：验收口径、证据最小集合、命名与索引规则（引用 `evidence/LOCK.md`）
  - 输出可执行WP清单与依赖，形成后续“连续推进直到上线”的执行序列
- Out of Scope（明确不做清单）：
  - 任何实现代码增量
  - 任何环境变更（VPS初始化、部署、配置修改等）
  - 任何Gate动作执行（除Checkpoint A文档性决策）
  - baseline/tag 及任何需要回滚的技术性变更

## 3) 优先级与关键路径（必须）
- P0（阻断项）：
  1) Checkpoint A五件套落盘 + 00-index条目DONE（本轮）
  2) 准备并推进T3硬证据#1/#2/#3的WP（作为解冻后首批执行）
- P1（可延后）：
  - 技术交付相关WP的细化（VPS初始化、脚本v2、监控、CI/CD、回滚演练、环境验证、基线冻结、上线）在T3硬证据链完成后再进入执行
- 关键路径顺序：
  - Checkpoint A闭环 → （Go）发布恢复推进指令 → 启动T3硬证据#1 → 硬证据#2 → 端到端归档演练（覆盖#1/#2） → T3-Gate验收通过 → 进入技术交付WP序列
- 阻断条件（Blocking Conditions）：
  - 任一证据条目无法满足“一跳定位/可复核/可追溯”
  - 00-index字段不全仍标DONE（审计失败）
  - 未Go即发生实现/环境变更/Gate动作/baseline-tag（冻结策略违例，默认No-Go并要求回滚证据）

## 4) 验收口径与证据标准（必须）
- DoD（Definition of Done）总则：
  - 所有可验收产出必须落盘在 evidence 体系内，并在 `evidence/00-index.md` 登记
  - DONE 必须包含：路径、版本号、时间点、责任人、关联指针、验证记录
- 必须产出证据清单（本轮Checkpoint A）：
  - PM最小输入完成版（本文件）
  - 验收&证据规格锁定版
  - WP清单完成版 + 依赖清单
  - Checkpoint A Go/No-Go结论
  - 恢复推进指令
- 证据存放与命名：遵循 `evidence/LOCK.md`

## 5) 风险与回滚要求（推荐）
- 风险列表：
  - 计划层面字段虽齐，但后续执行范围膨胀（Scope creep）
  - 证据口径不一致导致后续返工
  - 冻结策略被破坏引入不可审计增量
- 回滚要求：
  - 本轮为Freeze-Only，无技术回滚；若出现违规技术增量，必须提供回滚证据并在索引记录（默认No-Go）。
