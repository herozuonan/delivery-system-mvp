# MVP 测试操作文档（Delivery System MVP）

适用对象：你现在要验证 MVP 的“可用性 + CI/CD + 回滚”链路。

> 现状：
> - Repo: https://github.com/herozuonan/delivery-system-mvp
> - VPS: 47.85.50.81
> - 运行端口：8080
> - CI/CD：GitHub Actions build-and-deploy（push 触发）
> - 回滚：GitHub Actions rollback（手动触发）

---

## 0. 你需要准备的东西
- 能 SSH 登录 VPS（建议用 SSH key）
- 能访问 GitHub Actions 页面

---

## 1) 基础可用性测试（服务是否活着）
### 1.1 在 VPS 上测试（最可靠）
SSH 登录后执行：
```bash
curl -fsS http://127.0.0.1:8080/healthz && echo
curl -fsS http://127.0.0.1:8080/ && echo
```
期望输出：
- `/healthz` → `ok`
- `/` → `delivery-system-mvp`

### 1.2 从你本机测试（取决于安全组是否放行 8080）
如果你安全组已放行 8080：
```bash
curl -fsS http://47.85.50.81:8080/healthz && echo
```
如果访问失败：优先按 1.1 走 VPS 本机回环测试。

---

## 2) 运行态确认（确认镜像版本、端口映射）
在 VPS 上执行（你的机器是 podman）：
```bash
podman ps --format '{{.Names}} {{.Image}} {{.Ports}} {{.Status}}'
```
期望看到容器：
- name：`delivery-system-mvp`
- image：`ghcr.io/herozuonan/delivery-system-mvp:sha-xxxxxxx`（或 latest）
- ports：`0.0.0.0:8080->8080/tcp`（如果你保持对外暴露）

查看容器日志：
```bash
podman logs delivery-system-mvp --tail 100
```

---

## 3) CI/CD 部署测试（验证“提交→自动部署→健康”）
目标：你 push 一次代码后，Actions 自动 build+push，并 SSH 到 VPS 拉新镜像启动，然后 `/healthz` 仍为 ok。

### 3.1 触发方式
任选其一：
- 方式A：改一行 README 然后 push 到 `master`
- 方式B：GitHub → Actions → build-and-deploy → Run workflow（手动触发）

### 3.2 验证点
1) Actions run 结果：build ✅ deploy ✅
2) VPS 上镜像 tag 变化（sha-xxxxxxx 变化）：
```bash
podman ps --format '{{.Image}}'
```
3) 健康检查：
```bash
curl -fsS http://127.0.0.1:8080/healthz && echo
```

---

## 4) 回滚演练测试（验证“可回滚 + 健康”）
目标：手动触发 rollback workflow，把容器切回上一个镜像，并且健康检查 ok。

### 4.1 触发方式
GitHub → Actions → rollback → Run workflow（手动触发）

### 4.2 验证点
1) Actions run：rollback ✅
2) VPS 镜像变化：
```bash
podman ps --format '{{.Image}}'
```
3) 健康检查：
```bash
curl -fsS http://127.0.0.1:8080/healthz && echo
```

---

## 5) 常见问题排查
### 5.1 Actions deploy 失败（SSH）
- 看 Actions 失败日志：通常是 key/host/user 配置问题。
- VPS 侧看 sshd 日志：
```bash
journalctl -u sshd -n 50 --no-pager
```

### 5.2 VPS 拉 ghcr 镜像失败
- 需要配置 GitHub Secrets：`GHCR_USER` / `GHCR_TOKEN`（PAT 至少 read:packages）
- 或在 VPS 手动登录一次：
```bash
podman login ghcr.io
```

### 5.3 端口访问不到
- 优先在 VPS 内部用 127.0.0.1 测；外部访问涉及安全组/防火墙。

---

## 6) 测试结果记录（建议你回传我这几条输出）
把以下命令输出复制给我，我就能把 QA 证据补齐：
```bash
podman ps --format '{{.Names}} {{.Image}} {{.Ports}} {{.Status}}'
curl -fsS http://127.0.0.1:8080/healthz && echo
```
再加上：你触发的 Actions run 链接（build-and-deploy + rollback）。
