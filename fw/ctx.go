package fw

import (
	"context"
)

/* __________________________________________________ */

type contextKey string

const key contextKey = "ctx.app"

/* __________________________________________________ */

// FromCtx
// Takes a context.Context and returns the Application associated with it (if any).
func FromCtx(ctx context.Context) *Application {
	if value, ok := ctx.Value(key).(*Application); ok {
		return value
	}
	return nil
}

// WithCtx
// Associates an Application instance with a context.Context and returns it.
func WithCtx(ctx context.Context, app *Application) context.Context {
	value := FromCtx(ctx)
	if value == app {
		return ctx
	}
	return context.WithValue(ctx, key, app)
}
