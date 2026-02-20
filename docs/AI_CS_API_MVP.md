# AI CS API (MVP) - 公网测试说明

当前环境：直接用公网 IP 测试，后续再挂域名。

## Base URL
- `http://47.85.50.81:18882`

## Endpoints
### 1) Health
```bash
curl -fsS http://47.85.50.81:18882/healthz && echo
```
期望：`ok`

### 2) Chat (同步)
```bash
curl -fsS http://47.85.50.81:18882/v1/chat \
  -H 'Content-Type: application/json' \
  -d '{"user_id":"u1","session_id":"","message":"你好","stage":"pre_sales","goal":"collect_requirements"}' | jq .
```
期望：返回 JSON，包含 `reply/session_id/trace_id`。

### 3) Chat (流式 SSE)
```bash
curl -N "http://47.85.50.81:18882/v1/chat/stream?user_id=u1&session_id=&stage=pre_sales&goal=collect_requirements&message=你好" 
```
期望：持续输出 `event: start/delta/done`。

### 4) Metrics
```bash
curl -fsS http://47.85.50.81:18882/metrics | head
```

## 当前模式
- `AI_MODE=mock`：MVP mock 回复（先把链路、并发、SSE、观测跑通）
- 后续会接入真实模型（走 Model Gateway），对测试工具与接口不破坏。
