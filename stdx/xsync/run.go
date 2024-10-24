package xsync

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

var ErrRecovered = errors.New("recovered from panic")

// Run executes a function.
// Panics are suppressed.
func Run(task func()) {
	_ = execFunc(task)
}

// RunFunc executes function.
// Panics are captured and returned as errors.
func RunFunc(task func()) error {
	return execFunc(task)
}

// RunErrorFunc executes function.
// Panics are captured and returned as errors.
func RunErrorFunc(task func() error) error {
	return execErrorFunc(task)
}

// Go runs a function in a new goroutine.
// Panics are suppressed.
func Go(task func()) {
	go func() {
		defer suppressed()
		task()
	}()
}

// GoFunc executes a function in a new goroutine.
// Panics are captured and returned as errors.
func GoFunc(task func()) <-chan error {
	out := make(chan error, 1)
	go func() {
		defer close(out)
		out <- execFunc(task)
	}()
	return out
}

// GoFunc executes a function in a new goroutine.
// Panics are captured and returned as errors.
func GoErrorFunc(task func() error) <-chan error {
	out := make(chan error, 1)
	go func() {
		defer close(out)
		out <- execErrorFunc(task)
	}()
	return out
}

func GOWG(task func(), wgs ...*sync.WaitGroup) {
	addWG(wgs...)
	go func() {
		defer doneWG(wgs...)
		_ = execFunc(task)
	}()
}

func GOWGFunc(task func(), wgs ...*sync.WaitGroup) <-chan error {
	out := make(chan error, 1)
	addWG(wgs...)
	go func() {
		defer doneWG(wgs...)
		out <- execFunc(task)
	}()
	return out
}

func GOWGErrorFunc(task func() error, wgs ...*sync.WaitGroup) <-chan error {
	out := make(chan error, 1)
	addWG(wgs...)
	go func() {
		defer doneWG(wgs...)
		out <- execErrorFunc(task)
	}()
	return out
}

func addWG(wgs ...*sync.WaitGroup) {
	for _, wg := range wgs {
		wg.Add(1)
	}
}

func doneWG(wgs ...*sync.WaitGroup) {
	for _, wg := range wgs {
		wg.Done()
	}
}

func execFunc(task func()) (result error) {
	defer func() {
		if res := recover(); res != nil {
			// wrap error if result is error otherwise no
			if err, ok := res.(error); ok {
				err = fmt.Errorf("xsync: %w: %w", ErrRecovered, err)
				slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
				result = err
				return
			} else {
				err = fmt.Errorf("xsync: %w: %v", ErrRecovered, res)
				slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
				result = err
				return
			}
		}
	}()

	task()
	return result
}

func execErrorFunc(task func() error) (result error) {
	defer func() {
		if res := recover(); res != nil {
			// wrap error if result is error otherwise no
			if err, ok := res.(error); ok {
				err = fmt.Errorf("xsync: %w: %w", ErrRecovered, err)
				slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
				result = err
				return
			} else {
				err = fmt.Errorf("xsync: %w: %v", ErrRecovered, res)
				slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
				result = err
				return
			}
		}
	}()

	return task()
}

func execContextualFunc(ctx context.Context, tasks <-chan func(), errs chan<- error) {
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		case task, open := <-tasks:
			{
				if !open {
					return
				}

				errs <- execFunc(task)
				return
			}
		}
	}
}

func execContextualErrorFunc(ctx context.Context, tasks <-chan func() error, errs chan<- error) {
	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		case task, open := <-tasks:
			{
				if !open {
					return
				}

				errs <- execErrorFunc(task)
				return
			}
		}
	}
}

// Just for reference
func suppressed() {
	if res := recover(); res != nil {
		// wrap error if result is error otherwise no
		if err, ok := res.(error); ok {
			err = fmt.Errorf("xsync: %w: %w", ErrRecovered, err)
			slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
		} else {
			err = fmt.Errorf("xsync: %w: %v", ErrRecovered, res)
			slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
		}
	}
}

// Just for reference
func returned() error {
	if res := recover(); res != nil {
		// wrap error if result is error otherwise no
		if err, ok := res.(error); ok {
			err = fmt.Errorf("xsync: %w: %w", ErrRecovered, err)
			slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
			return err
		} else {
			err = fmt.Errorf("xsync: %w: %v", ErrRecovered, res)
			slogx.Log().With("stack", string(debug.Stack())).Error(err.Error())
			return err
		}
	}
	return nil
}
