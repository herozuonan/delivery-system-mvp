#!/usr/bin/env bash
set -euo pipefail

# Deploy script intended to run on the VPS.
# Usage:
#   TAG=sha-xxxx IMAGE_REPO=ghcr.io/herozuonan/delivery-system-mvp ./deploy.sh

APP_NAME=${APP_NAME:-delivery-system-mvp}
IMAGE_REPO=${IMAGE_REPO:-ghcr.io/herozuonan/delivery-system-mvp}
TAG=${TAG:-latest}
PORT=${PORT:-8080}

STATE_DIR=${STATE_DIR:-/opt/${APP_NAME}}
PREV_FILE="$STATE_DIR/previous_tag"
mkdir -p "$STATE_DIR"

runtime=""
if command -v podman >/dev/null 2>&1; then runtime="podman"; fi
if [[ -z "$runtime" ]] && command -v docker >/dev/null 2>&1; then runtime="docker"; fi
if [[ -z "$runtime" ]]; then
  echo "ERROR: neither podman nor docker found" >&2
  exit 1
fi

echo "Runtime=$runtime"

echo "Saving previous tag (best-effort)..."
if $runtime ps --format '{{.Image}} {{.Names}}' 2>/dev/null | grep -q " ${APP_NAME}$"; then
  prev_image=$($runtime ps --format '{{.Image}} {{.Names}}' | awk -v n="$APP_NAME" '$2==n{print $1}')
  echo "$prev_image" | tee "$STATE_DIR/previous_image"
fi

# Optional GHCR auth (recommended if pulls fail)
if [[ -n "${GHCR_USER:-}" && -n "${GHCR_TOKEN:-}" ]]; then
  echo "Logging into GHCR as $GHCR_USER"
  echo "$GHCR_TOKEN" | $runtime login ghcr.io -u "$GHCR_USER" --password-stdin
fi

image="${IMAGE_REPO}:${TAG}"
echo "Pulling $image"
$runtime pull "$image"

echo "Stopping old container (best-effort)"
$runtime rm -f "$APP_NAME" >/dev/null 2>&1 || true

echo "Starting new container"
$runtime run -d --name "$APP_NAME" -p "${PORT}:8080" --restart=unless-stopped "$image"

echo "Health check"
for i in {1..20}; do
  if curl -fsS "http://127.0.0.1:${PORT}/healthz" >/dev/null 2>&1; then
    echo "ok"
    echo "$TAG" > "$PREV_FILE"  # used as "current" for now
    exit 0
  fi
  sleep 1
done

echo "ERROR: health check failed" >&2
$runtime logs "$APP_NAME" --tail 200 || true
exit 2
