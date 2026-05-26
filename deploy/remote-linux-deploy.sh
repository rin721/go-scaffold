#!/usr/bin/env bash
set -euo pipefail

log() {
  printf '[deploy] %s\n' "$*"
}

die() {
  printf '[deploy] ERROR: %s\n' "$*" >&2
  exit 1
}

usage() {
  cat <<'USAGE'
Usage:
  DEPLOY_ENV=production DEPLOY_IMAGE=go-scaffold:local APP_PORT=9999 bash deploy/remote-linux-deploy.sh

Options:
  --env <staging|production>   Deployment environment.
  --image <image>              Docker image tag to build or run.
  --path <path>                Runtime directory, default /opt/go-scaffold.
  --port <port>                Host port, default 9999.
  --source-dir <path>          Repository directory, default current directory.
  --pull                       Pull DEPLOY_IMAGE instead of building it locally.
  --no-build                   Skip docker build.

Environment variables with the same names can be used instead of flags.
This script writes DEPLOY_PATH/.env.deploy dynamically from the final values.
USAGE
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    --env)
      DEPLOY_ENV="$2"
      shift 2
      ;;
    --image)
      DEPLOY_IMAGE="$2"
      shift 2
      ;;
    --path)
      DEPLOY_PATH="$2"
      shift 2
      ;;
    --port)
      APP_PORT="$2"
      shift 2
      ;;
    --source-dir)
      SOURCE_DIR="$2"
      shift 2
      ;;
    --pull)
      DEPLOY_PULL_IMAGE=true
      DEPLOY_BUILD_IMAGE=false
      shift
      ;;
    --no-build)
      DEPLOY_BUILD_IMAGE=false
      shift
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      die "unknown argument: $1"
      ;;
  esac
done

require_cmd() {
  command -v "$1" >/dev/null 2>&1 || die "$1 is required"
}

require_cmd docker

if docker compose version >/dev/null 2>&1; then
  compose=(docker compose)
elif command -v docker-compose >/dev/null 2>&1; then
  compose=(docker-compose)
else
  die "docker compose or docker-compose is required"
fi

SOURCE_DIR="${SOURCE_DIR:-$(pwd)}"
DEPLOY_ENV="${DEPLOY_ENV:-production}"
DEPLOY_PATH="${DEPLOY_PATH:-/opt/go-scaffold}"
DEPLOY_COMPOSE_FILE="${DEPLOY_COMPOSE_FILE:-docker-compose.yml}"
DEPLOY_SERVICE="${DEPLOY_SERVICE:-go-scaffold}"
DEPLOY_CONTAINER_NAME="${DEPLOY_CONTAINER_NAME:-go-scaffold}"
APP_PORT="${APP_PORT:-9999}"
DEPLOY_IMAGE="${DEPLOY_IMAGE:-go-scaffold:local}"
DEPLOY_BUILD_IMAGE="${DEPLOY_BUILD_IMAGE:-true}"
DEPLOY_PULL_IMAGE="${DEPLOY_PULL_IMAGE:-false}"
DEPLOY_CONFIG_OVERWRITE="${DEPLOY_CONFIG_OVERWRITE:-false}"
DEPLOY_HEALTH_URL="${DEPLOY_HEALTH_URL:-http://127.0.0.1:${APP_PORT}/health}"
DEPLOY_READY_URL="${DEPLOY_READY_URL:-http://127.0.0.1:${APP_PORT}/ready}"
CONFIG_SOURCE="${CONFIG_SOURCE:-${SOURCE_DIR}/deploy/config.production.example.yaml}"
COMPOSE_SOURCE="${COMPOSE_SOURCE:-${SOURCE_DIR}/deploy/docker-compose.production.example.yml}"

case "$DEPLOY_ENV" in
  staging|production) ;;
  *) die "DEPLOY_ENV must be staging or production" ;;
esac

case "$DEPLOY_BUILD_IMAGE" in
  true|false) ;;
  *) die "DEPLOY_BUILD_IMAGE must be true or false" ;;
esac

case "$DEPLOY_PULL_IMAGE" in
  true|false) ;;
  *) die "DEPLOY_PULL_IMAGE must be true or false" ;;
esac

case "$DEPLOY_CONFIG_OVERWRITE" in
  true|false) ;;
  *) die "DEPLOY_CONFIG_OVERWRITE must be true or false" ;;
esac

[[ "$DEPLOY_PATH" = /* ]] || die "DEPLOY_PATH must be absolute"
[[ "$APP_PORT" =~ ^[0-9]+$ ]] || die "APP_PORT must be numeric"
[ -d "$SOURCE_DIR" ] || die "SOURCE_DIR does not exist: $SOURCE_DIR"
[ -f "$COMPOSE_SOURCE" ] || die "compose template not found: $COMPOSE_SOURCE"
[ -f "$CONFIG_SOURCE" ] || die "config template not found: $CONFIG_SOURCE"

validate_env_value() {
  local key="$1"
  local value="$2"
  [ -n "$value" ] || die "$key cannot be empty"
  if [[ "$value" == *$'\n'* || "$value" == *$'\r'* ]]; then
    die "$key cannot contain newlines"
  fi
}

for key in \
  DEPLOY_ENV \
  DEPLOY_COMPOSE_FILE \
  DEPLOY_SERVICE \
  DEPLOY_CONTAINER_NAME \
  APP_PORT \
  DEPLOY_IMAGE \
  DEPLOY_HEALTH_URL \
  DEPLOY_READY_URL; do
  validate_env_value "$key" "${!key}"
done

log "source: $SOURCE_DIR"
log "target: $DEPLOY_PATH"
log "environment: $DEPLOY_ENV"
log "image: $DEPLOY_IMAGE"

mkdir -p "$DEPLOY_PATH/configs" "$DEPLOY_PATH/data" "$DEPLOY_PATH/logs"

if [ ! -f "$DEPLOY_PATH/configs/config.yaml" ] || [ "$DEPLOY_CONFIG_OVERWRITE" = "true" ]; then
  cp "$CONFIG_SOURCE" "$DEPLOY_PATH/configs/config.yaml"
  log "wrote $DEPLOY_PATH/configs/config.yaml from template"
else
  log "kept existing $DEPLOY_PATH/configs/config.yaml"
fi

cp "$COMPOSE_SOURCE" "$DEPLOY_PATH/$DEPLOY_COMPOSE_FILE"

env_file="$DEPLOY_PATH/.env.deploy"
tmp_env_file="$DEPLOY_PATH/.env.deploy.tmp"
{
  printf 'DEPLOY_ENV=%s\n' "$DEPLOY_ENV"
  printf 'DEPLOY_COMPOSE_FILE=%s\n' "$DEPLOY_COMPOSE_FILE"
  printf 'DEPLOY_SERVICE=%s\n' "$DEPLOY_SERVICE"
  printf 'DEPLOY_CONTAINER_NAME=%s\n' "$DEPLOY_CONTAINER_NAME"
  printf 'APP_PORT=%s\n' "$APP_PORT"
  printf 'DEPLOY_IMAGE=%s\n' "$DEPLOY_IMAGE"
  printf 'DEPLOY_HEALTH_URL=%s\n' "$DEPLOY_HEALTH_URL"
  printf 'DEPLOY_READY_URL=%s\n' "$DEPLOY_READY_URL"
} > "$tmp_env_file"
chmod 600 "$tmp_env_file"
mv "$tmp_env_file" "$env_file"
log "generated $env_file"

if [ "$(id -u)" = "0" ]; then
  chown -R 10001:10001 "$DEPLOY_PATH/data" "$DEPLOY_PATH/logs"
elif command -v sudo >/dev/null 2>&1 && sudo -n true >/dev/null 2>&1; then
  sudo chown -R 10001:10001 "$DEPLOY_PATH/data" "$DEPLOY_PATH/logs"
else
  log "warning: cannot chown data/logs to 10001:10001 without passwordless sudo"
fi

if [ "$DEPLOY_BUILD_IMAGE" = "true" ]; then
  log "building image"
  docker build -t "$DEPLOY_IMAGE" "$SOURCE_DIR"
fi

cd "$DEPLOY_PATH"

if [ "$DEPLOY_PULL_IMAGE" = "true" ]; then
  log "pulling image"
  docker pull "$DEPLOY_IMAGE"
  "${compose[@]}" --env-file "$env_file" -f "$DEPLOY_COMPOSE_FILE" pull "$DEPLOY_SERVICE"
fi

log "starting service"
"${compose[@]}" --env-file "$env_file" -f "$DEPLOY_COMPOSE_FILE" up -d "$DEPLOY_SERVICE"

if command -v curl >/dev/null 2>&1; then
  check_url() {
    local name="$1"
    local url="$2"

    for attempt in $(seq 1 30); do
      if curl -fsS --max-time 5 "$url" >/dev/null; then
        log "$name check passed"
        return 0
      fi
      sleep 2
    done

    die "$name check failed: $url"
  }

  check_url health "$DEPLOY_HEALTH_URL"
  check_url ready "$DEPLOY_READY_URL"
else
  log "warning: curl not found on host; skipped health/ready checks"
fi

log "deployment finished"
