package handler

import (
	"context"
	"log/slog"
)

// LogEventType represents the type of log event
type LogEventType string

const (
	LogEventTypeInfo  LogEventType = "info"
	LogEventTypeError LogEventType = "error"
	LogEventTypeWarn  LogEventType = "warn"
	LogEventTypeDebug LogEventType = "debug"
)

// Preserved tag keys
const (
	TagHandler         = "handler"
	TagError           = "err"
	TagUserID          = "user_id"
	TagRequestID       = "request_id"
	TagDuration        = "duration_ms"
	TagStatus          = "status"
	TagMethod          = "method"
	TagPath            = "path"
	TagStatusCode      = "status_code"
	TagProductID       = "product_id"
	TagTransactionID   = "transaction_id"
	TagCategoriesCount = "categories_count"
)

// LogTags holds both preserved and custom tags
type LogTags struct {
	Handler         string
	Error           error
	UserID          string
	RequestID       string
	Duration        int64
	Status          string
	Method          string
	Path            string
	StatusCode      int
	ProductID       uint32
	TransactionID   string
	CategoriesCount int
	Custom          map[string]interface{}
}

// NewLogTags creates a new LogTags with custom tags
func NewLogTags(custom map[string]interface{}) *LogTags {
	return &LogTags{
		Custom: custom,
	}
}

// WithHandler sets the handler tag
func (t *LogTags) WithHandler(handler string) *LogTags {
	t.Handler = handler
	return t
}

// WithError sets the error tag
func (t *LogTags) WithError(err error) *LogTags {
	t.Error = err
	return t
}

// WithUserID sets the user_id tag
func (t *LogTags) WithUserID(userID string) *LogTags {
	t.UserID = userID
	return t
}

// WithRequestID sets the request_id tag
func (t *LogTags) WithRequestID(requestID string) *LogTags {
	t.RequestID = requestID
	return t
}

// WithDuration sets the duration_ms tag
func (t *LogTags) WithDuration(duration int64) *LogTags {
	t.Duration = duration
	return t
}

// WithStatus sets the status tag
func (t *LogTags) WithStatus(status string) *LogTags {
	t.Status = status
	return t
}

// WithMethod sets the method tag
func (t *LogTags) WithMethod(method string) *LogTags {
	t.Method = method
	return t
}

// WithPath sets the path tag
func (t *LogTags) WithPath(path string) *LogTags {
	t.Path = path
	return t
}

// WithStatusCode sets the status_code tag
func (t *LogTags) WithStatusCode(code int) *LogTags {
	t.StatusCode = code
	return t
}

// WithProductID sets the product_id tag
func (t *LogTags) WithProductID(productID uint32) *LogTags {
	t.ProductID = productID
	return t
}

// WithTransactionID sets the transaction_id tag
func (t *LogTags) WithTransactionID(transactionID string) *LogTags {
	t.TransactionID = transactionID
	return t
}

// WithCategoriesCount sets the categories_count tag
func (t *LogTags) WithCategoriesCount(count int) *LogTags {
	t.CategoriesCount = count
	return t
}

// ToMap converts LogTags to map[string]interface{} for logging
func (t *LogTags) ToMap() map[string]interface{} {
	tags := make(map[string]interface{})

	// Add preserved tags if set
	if t.Handler != "" {
		tags[TagHandler] = t.Handler
	}
	if t.Error != nil {
		tags[TagError] = t.Error
	}
	if t.UserID != "" {
		tags[TagUserID] = t.UserID
	}
	if t.RequestID != "" {
		tags[TagRequestID] = t.RequestID
	}
	if t.Duration > 0 {
		tags[TagDuration] = t.Duration
	}
	if t.Status != "" {
		tags[TagStatus] = t.Status
	}
	if t.Method != "" {
		tags[TagMethod] = t.Method
	}
	if t.Path != "" {
		tags[TagPath] = t.Path
	}
	if t.StatusCode > 0 {
		tags[TagStatusCode] = t.StatusCode
	}
	if t.ProductID > 0 {
		tags[TagProductID] = t.ProductID
	}
	if t.TransactionID != "" {
		tags[TagTransactionID] = t.TransactionID
	}
	if t.CategoriesCount > 0 {
		tags[TagCategoriesCount] = t.CategoriesCount
	}

	// Add custom tags (custom tags override preserved tags if same key)
	if t.Custom != nil {
		for k, v := range t.Custom {
			tags[k] = v
		}
	}

	return tags
}

// Log wraps logger and context for convenient logging using slog's standard API
type Log struct {
	logger *slog.Logger
	ctx    context.Context
}

// NewLog creates a new Log instance
func NewLog(logger *slog.Logger, ctx context.Context) *Log {
	return &Log{
		logger: logger,
		ctx:    ctx,
	}
}

// WithContext returns a new Log instance with updated context
func (l *Log) WithContext(ctx context.Context) *Log {
	return NewLog(l.logger, ctx)
}

// Info logs an info message using slog's standard API
func (l *Log) Info(event string, msg string, tags *LogTags) {
	logger := l.logger.With(LogCtx(l.ctx)...)
	if tags != nil {
		logger = logger.With(tagsToAttrs(tags)...)
	}
	logger.Info(msg, "event", event, "event_type", string(LogEventTypeInfo))
}

// Error logs an error message using slog's standard API
func (l *Log) Error(event string, msg string, tags *LogTags) {
	logger := l.logger.With(LogCtx(l.ctx)...)
	if tags != nil {
		logger = logger.With(tagsToAttrs(tags)...)
	}
	logger.Error(msg, "event", event, "event_type", string(LogEventTypeError))
}

// Warn logs a warning message using slog's standard API
func (l *Log) Warn(event string, msg string, tags *LogTags) {
	logger := l.logger.With(LogCtx(l.ctx)...)
	if tags != nil {
		logger = logger.With(tagsToAttrs(tags)...)
	}
	logger.Warn(msg, "event", event, "event_type", string(LogEventTypeWarn))
}

// Debug logs a debug message using slog's standard API
func (l *Log) Debug(event string, msg string, tags *LogTags) {
	logger := l.logger.With(LogCtx(l.ctx)...)
	if tags != nil {
		logger = logger.With(tagsToAttrs(tags)...)
	}
	logger.Debug(msg, "event", event, "event_type", string(LogEventTypeDebug))
}

// tagsToAttrs converts LogTags to slog attributes
func tagsToAttrs(tags *LogTags) []interface{} {
	tagMap := tags.ToMap()
	attrs := make([]interface{}, 0, len(tagMap)*2)
	for k, v := range tagMap {
		attrs = append(attrs, k, v)
	}
	return attrs
}

// LogCtx extracts useful context info for logging
func LogCtx(ctx context.Context) []interface{} {
	return []interface{}{}
}
