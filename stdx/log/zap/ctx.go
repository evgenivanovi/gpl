package zap

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const zapCtxKey contextKey = "ctx.log.zap"
const zapSugarCtxKey contextKey = "ctx.log.zap.sugar"

// FromCtx
// Takes a context.Context and returns the zap.Logger associated with it (if any).
func FromCtx(ctx context.Context) *zap.SugaredLogger {

	value, ok := ctx.Value(zapSugarCtxKey).(*zap.SugaredLogger)

	if ok && value != nil {
		return value
	}

	return sugarLogger

}

// WithCtx
// Associates a zap.SugaredLogger instance with a context.Context and returns it.
func WithCtx(ctx context.Context, log *zap.SugaredLogger) context.Context {

	if value, ok := ctx.Value(zapSugarCtxKey).(*zap.SugaredLogger); ok {
		if value == log {
			return ctx
		}
	}

	return context.WithValue(ctx, zapSugarCtxKey, log)

}

func WithKV(ctx context.Context, key string, value any) context.Context {
	log := FromCtx(ctx).With(key, value)
	return WithCtx(ctx, log)
}

func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	log := FromCtx(ctx).Desugar().With(fields...).Sugar()
	return WithCtx(ctx, log)
}
