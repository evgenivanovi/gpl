package xsync

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

/* __________________________________________________ */

func RUN(task func()) {
	doRUN(task)
}

func doRUN(task func()) {
	if task != nil {
		var action = func(task func()) {
			defer rec()
			task()
		}
		action(task)
	}
}

/* __________________________________________________ */

func GO(task func()) {
	go doGO(task)
}

func doGO(task func()) {
	if task != nil {
		var action = func(task func()) {
			defer rec()
			task()
		}
		action(task)
	}
}

/* __________________________________________________ */

func GOWG(task func(), wgs ...*sync.WaitGroup) {
	for _, wg := range wgs {
		wg.Add(1)
	}
	go doGOWG(task, wgs...)
}

func doGOWG(task func(), wgs ...*sync.WaitGroup) {

	done := func() {
		for _, wg := range wgs {
			wg.Done()
		}
	}

	defer done()

	if task == nil {
		return
	}

	var action = func(task func()) {
		defer rec()
		task()
	}
	action(task)

}

func GOCLWG(
	task func(ctx context.Context) error,
	ctx context.Context,
	cancel context.CancelFunc,
	wgs ...*sync.WaitGroup,
) {
	for _, wg := range wgs {
		wg.Add(1)
	}
	go doGOCLWG(task, ctx, cancel, wgs...)
}

func doGOCLWG(
	task func(ctx context.Context) error,
	ctx context.Context,
	cancel context.CancelFunc,
	wgs ...*sync.WaitGroup,
) {

	done := func() {
		for _, wg := range wgs {
			wg.Done()
		}
	}

	defer done()

	if task == nil {
		return
	}

	var action = func(task func(ctx context.Context) error) {
		defer ccl(cancel)
		select {
		case <-ctx.Done():
			_ = ctx.Err()
		default:
			if err := task(ctx); err != nil {
				cancel()
			}
		}
	}
	action(task)

}

func GOCHWG(task func() error, ch chan error, wgs ...*sync.WaitGroup) {
	for _, wg := range wgs {
		wg.Add(1)
	}
	go doGOCHWG(task, ch, wgs...)
}

func doGOCHWG(task func() error, ch chan error, wgs ...*sync.WaitGroup) {

	done := func() {
		for _, wg := range wgs {
			wg.Done()
		}
	}

	defer done()

	if task == nil {
		return
	}

	var action = func(task func() error) {
		defer rec()
		if err := task(); err != nil {
			ch <- err
			return
		}
	}
	action(task)

}

/* __________________________________________________ */

func rec() {
	if result := recover(); result != nil {
		err := fmt.Errorf("xsync: %v", result)
		slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
	}
}

func ccl(cancel context.CancelFunc) {
	if result := recover(); result != nil {
		err := fmt.Errorf("xsync: %v", result)
		slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
		cancel()
	}
}

/* __________________________________________________ */
