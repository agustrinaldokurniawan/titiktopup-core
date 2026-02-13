package handler

import (
	"log/slog"
)

// HandlerDeps holds common dependencies for all handlers
type HandlerDeps struct {
	Logger *slog.Logger
}

// NewHandlerDeps creates a new HandlerDeps with the given logger
func NewHandlerDeps(logger *slog.Logger) *HandlerDeps {
	return &HandlerDeps{
		Logger: logger,
	}
}
