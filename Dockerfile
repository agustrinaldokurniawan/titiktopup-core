# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
COPY go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build both API and seed binaries (statically linked, stripped for smaller size)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/api ./cmd/api && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/seed ./cmd/seed

# Runtime stage - minimal alpine
FROM alpine:latest

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/api .
COPY --from=builder /app/seed .

# Change ownership to non-root user
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose ports
EXPOSE 8080 50051

# Run the application
CMD ["./api"]
