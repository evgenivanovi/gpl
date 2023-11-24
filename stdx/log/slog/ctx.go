package slog

import (
	"context"
	"log/slog"
)

/* __________________________________________________ */

type contextKey string

const ctxKey contextKey = "ctx.log.slog"

/* __________________________________________________ */

// FromCtx
// Takes a context.Context and returns the zap.Logger associated with it (if any).
func FromCtx(ctx context.Context) *slog.Logger {

	value, ok := ctx.Value(ctxKey).(*slog.Logger)

	if ok && value != nil {
		return value
	}

	return logger

}

// WithCtx
// Associates a slog.Logger instance with a context.Context and returns it.
func WithCtx(ctx context.Context, log *slog.Logger) context.Context {

	if value, ok := ctx.Value(ctxKey).(*slog.Logger); ok {
		if value == log {
			return ctx
		}
	}

	return context.WithValue(ctx, ctxKey, log)

}

func WithKV(ctx context.Context, key string, value any) context.Context {
	log := FromCtx(ctx).With(key, value)
	return WithCtx(ctx, log)
}
