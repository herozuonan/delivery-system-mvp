# Delivery System MVP

目标：先做一个**可跑通的端到端流程 MVP**（本地/服务器均可），验证：
- 代码提交 → 构建 → 打包镜像 → 部署到单VPS（Docker Compose）→ 健康检查
- 可回滚（上一版本）
- 全过程证据可落盘（evidence/）
-
## 快速开始（本地）
```bash
make dev
curl -s http://localhost:8080/healthz
```

## 部署（预留）
见 docs/05-DEPLOYMENT.md 与 scripts/
