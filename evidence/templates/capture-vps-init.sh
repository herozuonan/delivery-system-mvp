#!/bin/bash
# VPS 初始化记录采集脚本（可选）
# 用途：把关键系统信息落盘，作为初始化证据的一部分

set -euo pipefail

OUT=${1:-"./vps-init-$(date +%Y%m%d-%H%M%S).log"}
{
  echo "== TIME =="; date -Is
  echo "== HOSTNAME =="; hostname
  echo "== OS =="; cat /etc/os-release || true
  echo "== KERNEL =="; uname -a
  echo "== USERS =="; whoami
  echo "== UPTIME =="; uptime || true
  echo "== IP =="; ip a || true
  echo "== ROUTE =="; ip r || true
  echo "== DNS =="; cat /etc/resolv.conf || true
  echo "== FIREWALL (ufw) =="; (ufw status verbose || true)
  echo "== FIREWALL (iptables) =="; (iptables -S || true)
  echo "== TIME SYNC =="; (timedatectl status || true)
  echo "== DISK =="; df -h
  echo "== MEM =="; free -h || true
} | tee "$OUT"

echo "Wrote: $OUT"
