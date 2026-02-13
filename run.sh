#!/usr/bin/env bash
set -e

ENV_FILE=".env"

# --- Cleanup Logic ---
cleanup() {
  echo "Cleaning up..."
  rm -rf ./cmd/api/tmp
}
# Execute cleanup function on script exit or interrupt
trap cleanup EXIT
# ---------------------

if [ ! -f "$ENV_FILE" ]; then
  echo ".env file not found at $ENV_FILE"
  exit 1
fi

export $(grep -vE '^(#|$)' "$ENV_FILE" | xargs)

case "$1" in
  seed)
    echo "Running database seed..."
    go run ./cmd/seed/main.go
    ;;
  seed-and-run|with-seed)
    echo "Running database seed..."
    go run ./cmd/seed/main.go
    echo ""
    echo "Starting API server..."
    go run ./cmd/api/main.go
    ;;
  air)
    echo "Starting API server with hot reload..."
    # Run air from the root but point to the config
    (cd cmd/api && air)
    ;;
  *)
    echo "Starting API server..."
    go run ./cmd/api/main.go
    ;;
esac