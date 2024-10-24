package xsync

import (
	"sync"

	"github.com/evgenivanovi/gpl/stdx"
)

// Async converts the given function into an asynchronous one.
// The asynchronous function returns a channel where the result will appear.
func Async[T any](action func() T) func() <-chan T {
	return func() <-chan T {
		done := make(chan T, 1)
		go func() {
			defer close(done)
			done <- action()
		}()
		return done
	}
}

// AsyncOnce converts the given function into an asynchronous one that is guaranteed to execute only once.
// The asynchronous function returns a channel where the result will appear.
// Subsequent calls will return the result from the first execution.
//
// The function is useful for scenarios where you need to perform an initialization or a resource-intensive
// operation only once, even if the function is called multiple times concurrently.
func AsyncOnce[T any](action func() T) func() <-chan T {
	var once sync.Once
	var result T

	return func() <-chan T {
		done := make(chan T, 1)
		go func() {
			defer close(done)
			once.Do(
				func() {
					result = action()
				},
			)
			done <- result
		}()
		return done
	}
}

// Await waits for the result of the asynchronous function on the given channel.
func Await[T any](in <-chan T) T {
	return <-in
}

func AsyncResult[T any](action func() (T, error)) func() <-chan stdx.Result[T] {
	return Async(func() stdx.Result[T] {
		return stdx.NewResult(action)
	})
}

func AsyncResultOnce[T any](action func() (T, error)) func() <-chan stdx.Result[T] {
	return AsyncOnce(func() stdx.Result[T] {
		return stdx.NewResult(action)
	})
}

func AwaitResult[T any](in <-chan stdx.Result[T]) stdx.Result[T] {
	return <-in
}
