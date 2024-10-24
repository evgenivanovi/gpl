package xsync

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// given
	done := make(chan bool, 1)

	// when
	Run(func() {
		done <- true
	})

	// when
	if !<-done {
		t.Error("Run did not execute task")
	}
}

func TestRunFunc_NoPanic(t *testing.T) {
	// given
	task := func() {
		// No panic
	}

	// when
	actual := RunFunc(task)

	// then
	if actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestRunFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() {
		panic(expected)
	}

	// when
	actual := RunFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestRunFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() {
		panic("test panic")
	}

	// when
	actual := RunFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestRunErrorFunc_NoError(t *testing.T) {
	// given
	task := func() error {
		return nil
	}

	// when
	actual := RunErrorFunc(task)

	// then
	if actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestRunErrorFunc_OnError(t *testing.T) {
	// given
	expected := errors.New("test error")

	task := func() error {
		return expected
	}

	// when
	actual := RunErrorFunc(task)

	// then
	if !errors.Is(actual, expected) {
		t.Errorf("Expected error %v, got %v", expected, actual)
	}
}

func TestRunErrorFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() error {
		panic(expected)
	}

	// when
	actual := RunErrorFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestRunErrorFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() error {
		panic("test panic")
	}

	// when
	actual := RunErrorFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestGo(t *testing.T) {
	// given
	done := make(chan bool, 1)

	// when
	Go(func() {
		done <- true
	})

	// then
	if !<-done {
		t.Error("Go did not execute task")
	}
}

func TestGoFunc_NoPanic(t *testing.T) {
	// given
	task := func() {
		// No panic
	}

	// when
	channel := GoFunc(task)

	// then
	if actual := <-channel; actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestGoFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() {
		panic(expected)
	}

	// when
	channel := GoFunc(task)

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestGoFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() {
		panic("test panic")
	}

	// when
	channel := GoFunc(task)

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestGoErrorFunc_NoError(t *testing.T) {
	// given
	task := func() error {
		return nil
	}

	// when
	channel := GoErrorFunc(task)

	// then
	if actual := <-channel; actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestGoErrorFunc_OnError(t *testing.T) {
	// given
	expected := errors.New("test error")

	task := func() error {
		return expected
	}

	// when
	channel := GoErrorFunc(task)

	// then
	actual := <-channel

	if !errors.Is(actual, expected) {
		t.Errorf("Expected error %v, got %v", expected, actual)
	}
}

func TestGoErrorFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() error {
		panic(expected)
	}

	// when
	channel := GoErrorFunc(task)

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestGoErrorFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() error {
		panic("test panic")
	}

	// when
	channel := GoErrorFunc(task)

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestGOWG(t *testing.T) {
	// given
	done := make(chan bool, 1)

	task := func() {
		done <- true
	}

	// when
	var wg sync.WaitGroup
	GOWG(task, &wg)
	wg.Wait()

	// then
	if !<-done {
		t.Error("GOWG did not execute task")
	}
}

func TestGOWGFunc_NoPanic(t *testing.T) {
	// given
	task := func() {
		// No panic
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGFunc(task, &wg)
	wg.Wait()

	// then
	if actual := <-channel; actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestGOWGFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() {
		panic(expected)
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGFunc(task, &wg)
	wg.Wait()

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestGOWGFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() {
		panic("test panic")
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGFunc(task, &wg)
	wg.Wait()

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestGOWGErrorFunc_NoError(t *testing.T) {
	// given
	task := func() error {
		return nil
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGErrorFunc(task, &wg)
	wg.Wait()

	// then
	if actual := <-channel; actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestGOWGErrorFunc_OnError(t *testing.T) {
	// given
	expected := errors.New("test error")

	task := func() error {
		return expected
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGErrorFunc(task, &wg)
	wg.Wait()

	// then
	actual := <-channel

	if !errors.Is(actual, expected) {
		t.Errorf("Expected error %v, got %v", expected, actual)
	}
}

func TestGOWGErrorFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() error {
		panic(expected)
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGErrorFunc(task, &wg)
	wg.Wait()

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestGOWGErrorFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() error {
		panic("test panic")
	}

	// when
	var wg sync.WaitGroup
	channel := GOWGErrorFunc(task, &wg)
	wg.Wait()

	// then
	actual := <-channel

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestExecFunc_NoPanic(t *testing.T) {
	// given
	task := func() {
		// No panic
	}

	// when
	actual := execFunc(task)

	// then
	if actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestExecFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() {
		panic(expected)
	}

	// when
	actual := execFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestExecFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() {
		panic("test panic")
	}

	// when
	actual := execFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestExecErrorFunc_NoError(t *testing.T) {
	// given
	task := func() error {
		return nil
	}

	// when
	actual := execErrorFunc(task)

	// then
	if actual != nil {
		t.Errorf("Expected no error, got %v", actual)
	}
}

func TestExecErrorFunc_OnError(t *testing.T) {
	// given
	expected := errors.New("test error")

	task := func() error {
		return expected
	}

	// when
	actual := execErrorFunc(task)

	// then
	if !errors.Is(actual, expected) {
		t.Errorf("Expected error %v, got %v", expected, actual)
	}
}

func TestExecErrorFunc_OnPanic(t *testing.T) {
	// given
	expected := errors.New("test panic")

	task := func() error {
		panic(expected)
	}

	// when
	actual := execErrorFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestExecErrorFunc_OnPanicWithString(t *testing.T) {
	// given
	task := func() error {
		panic("test panic")
	}

	// when
	actual := execErrorFunc(task)

	// then
	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestExecContextualFunc(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func(), 1)

	done := make(chan bool, 1)

	tasks <- func() {
		done <- true
	}

	// when
	go execContextualFunc(context.Background(), tasks, errs)

	// then
	if !<-done {
		t.Error("Task was not executed")
	}
}

func TestExecContextualFunc_OnPanic(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func(), 1)

	expected := errors.New("test panic")

	tasks <- func() {
		panic(expected)
	}

	// when
	go execContextualFunc(context.Background(), tasks, errs)

	// then
	err := <-errs

	if !errors.Is(err, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, err)
	}

	if !errors.Is(err, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, err)
	}
}

func TestExecContextualFunc_ContextCancelled(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func(), 1)

	// when
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// then
	execContextualFunc(ctx, tasks, errs)
}

func TestExecContextualFunc_ContextTimeout(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func(), 1)

	// when
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	go execContextualFunc(ctx, tasks, errs)

	// then
	<-ctx.Done()

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("Expected error %v, got %v", context.DeadlineExceeded, ctx.Err())
	}
}

func TestExecContextualErrorFunc(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func() error, 1)

	done := make(chan bool, 1)

	tasks <- func() error {
		done <- true
		return nil
	}

	// when
	go execContextualErrorFunc(context.Background(), tasks, errs)

	// then
	if !<-done {
		t.Error("Task was not executed")
	}
}

func TestExecContextualFunc_OnError(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func() error, 1)

	expected := errors.New("test error")

	tasks <- func() error {
		return expected
	}

	// when
	go execContextualErrorFunc(context.Background(), tasks, errs)

	// then
	actual := <-errs

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestExecContextualErrorFunc_OnPanic(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func() error, 1)

	expected := errors.New("test panic")

	tasks <- func() error {
		panic(expected)
	}

	// when
	go execContextualErrorFunc(context.Background(), tasks, errs)

	// then
	actual := <-errs

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if !errors.Is(actual, expected) {
		t.Errorf("Expected panic error %v, got %v", expected, actual)
	}
}

func TestExecContextualErrorFunc_OnPanicWithString(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func() error, 1)

	tasks <- func() error {
		panic("test panic")
	}

	// when
	go execContextualErrorFunc(context.Background(), tasks, errs)

	// then
	actual := <-errs

	if !errors.Is(actual, ErrRecovered) {
		t.Errorf("Expected error %v, got %v", ErrRecovered, actual)
	}

	if actual == nil || actual.Error() != "xsync: recovered from panic: test panic" {
		t.Errorf("Expected panic error message 'xsync: recovered from panic: test panic', got %v", actual)
	}
}

func TestExecContextualErrorFunc_ContextCancelled(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func() error, 1)

	// when
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// then
	execContextualErrorFunc(ctx, tasks, errs)
}

func TestExecContextualErrorFunc_ContextTimeout(t *testing.T) {
	// given
	errs := make(chan error, 1)
	tasks := make(chan func() error, 1)

	// when
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	go execContextualErrorFunc(ctx, tasks, errs)

	// then
	<-ctx.Done()

	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("Expected error %v, got %v", context.DeadlineExceeded, ctx.Err())
	}
}
