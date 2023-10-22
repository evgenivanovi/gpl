package xsync

import "sync"

type Semaphore struct {
	channel chan struct{}
	once    sync.Once
}

func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		channel: make(chan struct{}, max),
	}
}

func (s *Semaphore) Acquire() {
	s.channel <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.channel
}

func (s *Semaphore) Close() {
	action := func() { close(s.channel) }
	s.once.Do(action)
}
