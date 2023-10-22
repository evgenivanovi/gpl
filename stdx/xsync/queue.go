package xsync

import (
	"sync"

	"github.com/pkg/errors"
)

var ErrQueueChannelClosed = errors.New("channel is closed")

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

func (q *Queue[T]) Empty() bool {
	return len(q.channel) == 0
}

func (q *Queue[T]) WaitEmptiness() {
	for q.Empty() == false {
		// waiting
	}
}

func (q *Queue[T]) Push(task T) {
	q.channel <- task
}

func (q *Queue[T]) PushAll(tasks []T) {
	for _, task := range tasks {
		q.Push(task)
	}
}

func (q *Queue[T]) PopWait() (T, error) {
	task, opened := <-q.channel
	if opened == false {
		return task, ErrQueueChannelClosed
	}
	return task, nil
}

func (q *Queue[T]) Close() {
	action := func() { close(q.channel) }
	q.once.Do(action)
}
