package xsync

import (
	"sync"
	"sync/atomic"
)

type SingleInFlight struct {
	mutex *atomic.Value
}

func NewSingleInFlight() SingleInFlight {
	var mutex atomic.Value
	mutex.Store(new(sync.Once))

	return SingleInFlight{
		mutex: &mutex,
	}
}

func (o *SingleInFlight) Do(task func()) {
	o.getOnce().Do(
		func() {
			task()
			o.setOnce()
		},
	)
}

func (o *SingleInFlight) getOnce() *sync.Once {
	return o.mutex.Load().(*sync.Once)
}

func (o *SingleInFlight) setOnce() {
	o.mutex.Store(new(sync.Once))
}
