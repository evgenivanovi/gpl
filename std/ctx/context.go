package ctx

import (
	"context"
	"time"
)

// DetachedContext
// Контекст, отвязывающий сигнал отмены от родительского контекста
type DetachedContext struct {
	ctx context.Context
}

func (d *DetachedContext) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d *DetachedContext) Done() <-chan struct{} {
	return nil
}

func (d *DetachedContext) Err() error {
	return nil
}

func (d *DetachedContext) Value(key any) any {
	return d.ctx.Value(key)
}

func NewDetachedContext(ctx context.Context) *DetachedContext {
	return &DetachedContext{
		ctx: ctx,
	}
}

func WithValueDetached(parent context.Context, key, val any) *DetachedContext {
	return NewDetachedContext(
		context.WithValue(parent, key, val),
	)
}
