package eg

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func WithContext(parent context.Context) (*Group, context.Context) {
	eg, egCtx := errgroup.WithContext(parent)
	log := slogx.FromCtx(parent)
	return &Group{eg: eg, log: log}, egCtx
}

var ErrRecovered = errors.New("recovered from panic")

type Group struct {
	eg  *errgroup.Group
	log *slog.Logger
}

func (g *Group) Go(action func() error) {
	g.eg.Go(g.invoke(action))
}

func (g *Group) TryGo(action func() error) bool {
	return g.eg.TryGo(g.invoke(action))
}

func (g *Group) invoke(action func() error) func() (err error) {
	return func() (err error) {
		defer func() {
			if result := recover(); result != nil {
				err = fmt.Errorf(
					"errgroup: %w: %v",
					ErrRecovered, result,
				)
				g.
					logger().
					With("stack", string(debug.Stack())).
					Error(err.Error())
			}
		}()
		err = action()
		return
	}
}

func (g *Group) Wait() error {
	return g.eg.Wait()
}

func (g *Group) SetLimit(n int) {
	g.eg.SetLimit(n)
}

func (g *Group) logger() *slog.Logger {
	if g.log != nil {
		return g.log
	}
	return slogx.Log()
}
