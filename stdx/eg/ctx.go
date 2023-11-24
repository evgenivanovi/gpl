package eg

import (
	"context"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	"golang.org/x/sync/errgroup"
)

/* __________________________________________________ */

func WithContext(ctx context.Context) (*Group, context.Context) {
	eg, egCtx := errgroup.WithContext(ctx)
	log := slogx.FromCtx(ctx)
	return &Group{eg: eg, log: log}, egCtx
}

/* __________________________________________________ */
