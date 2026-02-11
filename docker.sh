#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

CMD="$1"

case "$CMD" in
  up)
    echo "Starting Docker services..."
    docker compose up -d
    ;;
  down)
    echo "Stopping Docker services..."
    docker compose down
    ;;
  reset)
    echo "Resetting Docker services (containers + volumes)..."
    docker compose down -v
    docker compose up -d
    ;;
  logs)
    echo "Showing DB logs (Ctrl+C to exit)..."
    docker compose logs -f
    ;;
  *)
    echo "Usage: $0 {up|down|reset|logs}"
    exit 1
    ;;
esac