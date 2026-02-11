#!/usr/bin/env bash
set -e

ENV_FILE=".env"

if [ ! -f "$ENV_FILE" ]; then
  echo ".env file not found at $ENV_FILE"
  exit 1
fi

export $(grep -vE '^(#|$)' "$ENV_FILE" | xargs)

echo "Environment variables loaded from $ENV_FILE"

if [ "$1" = "seed" ]; then
  echo "Running DB seed..."
  go run ./cmd/seed/main.go
else
  echo "Starting API server..."
  go run ./cmd/api/main.go
fi