# Work Packages（可执行工作包）— Plan-to-Execute Gate

> Owner 默认：Engineer
> 证据命名/目录/索引：遵循 `evidence/LOCK.md`

## WP-01：补齐PM最小输入包并固化边界（P0）
- Owner: Engineer
- Inputs: `docs/08-PM_MIN_INPUT.md`（待PM填写）
- Steps:
  1) 校验关键字段是否齐全（目标/范围/优先级/验收口径）
  2) 生成/更新 Go/No-Go 结论记录
- Outputs:
  - PM最小输入包（完成版）
  - Go/No-Go 结论记录（证据化）
- DoD:
  - 关键字段齐全或明确No-Go；结论可复核
- Evidence Links:
  - evidence/10_plan_restore/pm-min-input-*.md
  - evidence/00-index.md（登记）
- Rollback Plan:
  - No-Go则继续冻结，不产生工程增量

## WP-02：固化验收清单与证据规格（P0）
- Owner: Engineer
- Inputs: `docs/09-ACCEPTANCE_AND_EVIDENCE_SPEC.md`
- Steps:
  1) 将验收条目与证据路径/命名/验证方法绑定
  2) 在索引中引用该规范（如需要）
- Outputs: 固化后的验收与证据规格（版本化）
- DoD: 每条验收均有对应证据位置/格式/命名与验证方法
- Evidence Links: evidence/10_plan_restore/acceptance-spec-*.md
- Rollback Plan: 不适用（文档性变更）

## WP-03：拆解P0工作项并建立依赖（P0）
- Owner: Engineer
- Inputs: PM最小输入包 + 验收/证据规格
- Steps:
  1) 输出P0/P1工作包映射与依赖
  2) 明确阻断条件
- Outputs: WP清单（含依赖）
- DoD: 所有P0映射到至少一个WP；依赖清晰
- Evidence Links: evidence/10_plan_restore/wp-list-*.md

## WP-04：建立恢复推进准入检查点（Go/No-Go）（P0）
- Owner: Engineer
- Inputs: PM最小输入包（完成版）+ WP清单
- Steps:
  1) 执行Checkpoint A（输入齐全）检查
  2) 产出Go/No-Go清单与结论
- Outputs: Go/No-Go清单与结论
- DoD: 可被第三方复核
- Evidence Links: evidence/10_plan_restore/go-nogo-*.md

## WP-05：发布“恢复推进指令”（仅Go后）（P0）
- Owner: Engineer
- Inputs: Go结论
- Steps:
  1) 明确本迭代允许的技术增量范围、禁区、回滚要求
  2) 同步到执行清单，作为后续审计依据
- Outputs: 恢复推进指令
- DoD: 指令与执行清单一致；后续动作必须引用该指令
- Evidence Links: evidence/10_plan_restore/resume-directive-*.md
