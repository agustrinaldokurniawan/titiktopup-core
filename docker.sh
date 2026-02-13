#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

CMD="$1"

case "$CMD" in
  up)
    echo "Starting Docker services..."
    docker compose up -d --build
    ;;
  down)
    echo "Stopping Docker services..."
    docker compose down
    ;;
  reset)
    echo "Resetting Docker services (containers + volumes)..."
    docker compose down -v
    docker compose up -d --build
    echo "Waiting for API container to be ready..."
    # Wait for API container to be running
    for i in {1..30}; do
      if docker compose ps api | grep -q "Up"; then
        echo "API container is ready!"
        break
      fi
      echo "Waiting... ($i/30)"
      sleep 2
    done
    echo "Running database seed..."
    docker compose exec api ./seed
    echo "✅ Reset complete! Database seeded."
    ;;
  seed)
    echo "Running database seed..."
    docker compose exec api ./seed
    ;;
  logs)
    echo "Showing logs (Ctrl+C to exit)..."
    docker compose logs -f
    ;;
  logs-api)
    echo "Showing API logs (Ctrl+C to exit)..."
    docker compose logs -f api
    ;;
  restart)
    echo "Restarting services..."
    docker compose restart
    ;;
  *)
    echo "Usage: $0 {up|down|reset|seed|logs|logs-api|restart}"
    echo ""
    echo "Commands:"
    echo "  up         - Start services with build"
    echo "  down       - Stop services"
    echo "  reset      - Stop, remove volumes, and restart"
    echo "  seed       - Run database seed in API container"
    echo "  logs       - Show all logs"
    echo "  logs-api   - Show API logs only"
    echo "  restart    - Restart services"
    exit 1
    ;;
esac