#!/usr/bin/env bash
set -euo pipefail

# Rollback script intended to run on the VPS.
# Requires /opt/delivery-system-mvp/previous_image or previous_tag.

APP_NAME=${APP_NAME:-delivery-system-mvp}
IMAGE_REPO=${IMAGE_REPO:-ghcr.io/herozuonan/delivery-system-mvp}
PORT=${PORT:-8080}
STATE_DIR=${STATE_DIR:-/opt/${APP_NAME}}

runtime=""
if command -v podman >/dev/null 2>&1; then runtime="podman"; fi
if [[ -z "$runtime" ]] && command -v docker >/dev/null 2>&1; then runtime="docker"; fi
if [[ -z "$runtime" ]]; then
  echo "ERROR: neither podman nor docker found" >&2
  exit 1
fi

prev_image=""
if [[ -f "$STATE_DIR/previous_image" ]]; then
  prev_image=$(cat "$STATE_DIR/previous_image" | tr -d '\r\n')
fi

if [[ -z "$prev_image" ]]; then
  echo "ERROR: no previous_image recorded" >&2
  exit 2
fi

echo "Rolling back to: $prev_image"
$runtime pull "$prev_image" || true
$runtime rm -f "$APP_NAME" >/dev/null 2>&1 || true
$runtime run -d --name "$APP_NAME" -p "${PORT}:8080" --restart=unless-stopped "$prev_image"

for i in {1..20}; do
  if curl -fsS "http://127.0.0.1:${PORT}/healthz" >/dev/null 2>&1; then
    echo "ok"
    exit 0
  fi
  sleep 1
done

echo "ERROR: rollback health check failed" >&2
$runtime logs "$APP_NAME" --tail 200 || true
exit 3
