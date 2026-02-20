# CI/CD 部署方式（GitHub Actions → SSH 到 VPS）

## 你需要在 GitHub Repo Secrets 配置
在仓库 Settings → Secrets and variables → Actions → New repository secret：

必需：
- `VPS_HOST`：47.85.50.81
- `VPS_USER`：root（或你后续创建的 deploy 用户）
- `VPS_SSH_KEY`：用于 Actions 登录 VPS 的 **私钥**（推荐单独生成一对 key）

可选：
- `GHCR_USER`：ghcr 用户名（例如 herozuonan）
- `GHCR_TOKEN`：PAT（至少 `read:packages`），当 VPS 拉 ghcr 需要认证时使用

## 如何生成一对专用 SSH Key（推荐）
在你本机执行：
```bash
ssh-keygen -t ed25519 -C "gha-delivery-system-mvp" -f ./gha_delivery_system_mvp
```

- 把 `gha_delivery_system_mvp`（私钥）内容复制到 GitHub Secret：`VPS_SSH_KEY`
- 把 `gha_delivery_system_mvp.pub`（公钥）追加到 VPS：`~/.ssh/authorized_keys`

VPS 上执行（示例）：
```bash
mkdir -p ~/.ssh && chmod 700 ~/.ssh
cat >> ~/.ssh/authorized_keys <<'EOF'
<把公钥整行粘贴到这里>
EOF
chmod 600 ~/.ssh/authorized_keys
```

## 工作流说明
- `build-and-deploy.yml`：push 到 master → build 镜像 → push 到 GHCR（sha-xxxx + latest）→ SSH 到 VPS 执行部署脚本。
- 部署脚本位置（VPS）：`/opt/delivery-system-mvp/deploy.sh`（由 workflow 通过 raw URL 拉取）
- 回滚 workflow：`rollback.yml`（手动触发，执行 `/opt/delivery-system-mvp/rollback.sh`）

## 安全建议（强烈）
- 建议尽快改成：创建 `deploy` 用户 + sudo 白名单 + 禁用 root 密码登录。
