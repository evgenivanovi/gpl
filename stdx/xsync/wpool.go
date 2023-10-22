package xsync

import (
	"context"
	"fmt"
	"sync"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

type WorkerPool[T any] struct {
	workers int
	queue   *Queue[T]
	action  func(ctx context.Context, task T) error

	mutex *sync.WaitGroup
}

func NewWorkerPool[T any](
	workers int,
	queue *Queue[T],
	action func(ctx context.Context, task T) error,
) *WorkerPool[T] {
	return &WorkerPool[T]{
		workers: workers,
		queue:   queue,
		action:  action,
		mutex:   &sync.WaitGroup{},
	}
}

func (wp *WorkerPool[T]) Run(ctx context.Context) {
	for worker := 0; worker < wp.workers; worker++ {
		go wp.execute(ctx)
	}
}

func (wp *WorkerPool[T]) Close() {
	wp.mutex.Wait()
}

func (wp *WorkerPool[T]) execute(ctx context.Context) {
	for {
		task, err := wp.queue.PopWait()
		if err != nil {
			break
		}
		wp.doExecute(ctx, task)
	}
}

func (wp *WorkerPool[T]) doExecute(ctx context.Context, task T) {
	wp.mutex.Add(1)
	defer wp.mutex.Done()

	err := wp.action(ctx, task)
	if err != nil {
		slogx.FromCtx(ctx).Debug(
			"error has occurred in worker",
			slogx.ErrAttr(fmt.Errorf("worker action: %w", err)),
		)
	}
}
