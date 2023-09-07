package splitslog

import (
	"context"
	"fmt"
	"log/slog"
)

// Splitter is a map of log levels to handlers.
// The default log levels (slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError) must be present, otherwise the SplitHandler panics.
type Splitter map[slog.Level]slog.Handler

// SplitHandler is a handler that splits log records to different handlers based on their level.
type SplitHandler struct {
	Splitter Splitter

	attrs []slog.Attr
	group string
}

// NewSplitHandler returns a new SplitHandler.
func NewSplitHandler(splitter Splitter) *SplitHandler {
	switch {
	case splitter == nil:
		panic("splitter of SplitHandler must not be nil")
	case splitter[slog.LevelDebug] == nil:
		panic("splitter of SplitHandler must have a handler for debug level")
	case splitter[slog.LevelInfo] == nil:
		panic("splitter of SplitHandler must have a handler for info level")
	case splitter[slog.LevelWarn] == nil:
		panic("splitter of SplitHandler must have a handler for warn level")
	case splitter[slog.LevelError] == nil:
		panic("splitter of SplitHandler must have a handler for error level")
	}

	return &SplitHandler{Splitter: splitter}
}

// Enabled implements Handler.Enabled.
func (h *SplitHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.getHandler(level).Enabled(ctx, level)
}

// Handle implements Handler.Handle.
func (h *SplitHandler) Handle(ctx context.Context, r slog.Record) error {
	handler := h.getHandler(r.Level)
	handler = handler.WithGroup(h.group)
	handler = handler.WithAttrs(h.attrs)
	return handler.Handle(ctx, r)
}

// WithAttrs implements Handler.WithAttrs.
func (h *SplitHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &SplitHandler{
		Splitter: h.Splitter,
		attrs:    attrs,
	}
}

// WithGroup implements Handler.WithGroup.
func (h *SplitHandler) WithGroup(name string) slog.Handler {
	return &SplitHandler{
		Splitter: h.Splitter,
		group:    name,
	}
}

func (h *SplitHandler) getHandler(level slog.Level) slog.Handler {
	handler, ok := h.Splitter[level]
	if !ok {
		panic(fmt.Sprintf("no handler registered for level %s", level))
	}

	return handler
}
