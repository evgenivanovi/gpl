package xsync

import (
	"sync"

	"github.com/pkg/errors"
)

/* __________________________________________________ */

type Queue[T any] struct {
	channel chan T
	once    *sync.Once
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		once:    new(sync.Once),
		channel: make(chan T, size),
	}
}

func (q *Queue[T]) Push(task T) {
	q.channel <- task
}

func (q *Queue[T]) PopWait() (T, error) {
	task, opened := <-q.channel
	if !opened {
		return task, errors.New("channel is closed")
	}
	return task, nil
}

func (q *Queue[T]) Close() {
	action := func() { close(q.channel) }
	q.once.Do(action)
}

/* __________________________________________________ */
