package xsync

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	"golang.org/x/sync/errgroup"
)

func WithContext(parent context.Context) (*ErrorGroup, context.Context) {
	eg, ctx := errgroup.WithContext(parent)
	log := slogx.FromCtx(parent)
	return &ErrorGroup{eg: eg, log: log}, ctx
}

type ErrorGroup struct {
	eg  *errgroup.Group
	log *slog.Logger
}

func (g *ErrorGroup) Go(action func() error) {
	g.eg.Go(g.invoke(action))
}

func (g *ErrorGroup) TryGo(action func() error) bool {
	return g.eg.TryGo(g.invoke(action))
}

func (g *ErrorGroup) RetryGo(ctx context.Context, action func() error) bool {
	for {
		select {
		case <-ctx.Done():
			{
				return false
			}
		default:
			{
				if g.TryGo(action) {
					return true
				}
			}
		}
	}
}

func (g *ErrorGroup) invoke(action func() error) func() (err error) {
	return func() (err error) {
		defer func() {
			if result := recover(); result != nil {
				err = fmt.Errorf("errgroup: %w: %v", ErrRecovered, result)
				g.logger().With("stack", string(debug.Stack())).Error(err.Error())
			}
		}()
		err = action()
		return
	}
}

func (g *ErrorGroup) Wait() error {
	return g.eg.Wait()
}

func (g *ErrorGroup) SetLimit(n int) {
	g.eg.SetLimit(n)
}

func (g *ErrorGroup) logger() *slog.Logger {
	if g.log != nil {
		return g.log
	}
	return slogx.Log()
}

// TryGo continuously tries to execute the given function in a new goroutine,
// repeatedly calling TryGo on the errgroup.Group
// until it is successfully started or the context is done.
func TryGo(ctx context.Context, eg *errgroup.Group, action func() error) bool {
	for {
		select {
		case <-ctx.Done():
			{
				return false
			}
		default:
			{
				if eg.TryGo(action) {
					return true
				}
			}
		}
	}
}
