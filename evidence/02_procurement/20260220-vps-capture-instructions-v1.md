# VPS 规格/地域 采集指令（v1）

目标：补齐“VPS 等效证据包”的可复核材料。**不要把密码/私钥发出来。**

## A) 控制台证据（你做）
请在 VPS 提供商控制台截 1-2 张图（可打码）：
- 实例列表页：显示实例名/ID、地域、配置（CPU/RAM/磁盘/带宽）
- 账单/资源详情页：能证明该实例在你名下并可用（金额/订单号可打码）

把图片发我（或把关键信息文字贴出来也行）。

## B) 系统侧采集（你在 VPS 上执行）
SSH 登录到 VPS 后执行：

```bash
mkdir -p /root/mvp-evidence && cd /root/mvp-evidence

# 基本信息
{
  echo "# CAPTURE_TIME=$(date -Is)";
  echo "# HOSTNAME=$(hostname)";
  echo "# UPTIME=$(uptime -p)";
  echo;
  echo "## OS"; cat /etc/os-release 2>/dev/null || true; echo;
  echo "## KERNEL"; uname -a; echo;
  echo "## CPU"; lscpu 2>/dev/null || cat /proc/cpuinfo | head -n 50; echo;
  echo "## MEM"; free -h; echo;
  echo "## DISK"; df -h; echo;
  echo "## IP"; ip -br a 2>/dev/null || ifconfig 2>/dev/null || true; echo;
} | tee 20260220-vps-spec-capture-v1.txt
```

然后把这个文件内容复制给我（或用 scp 下载后发我）。

## C) 最小安全建议（不改密码也请做这个）
- 确认 22 端口仅对你的办公 IP 开放（安全组/防火墙）。
- 部署完成后立刻轮换 root 密码，并切到 SSH key + 非 root 用户。
