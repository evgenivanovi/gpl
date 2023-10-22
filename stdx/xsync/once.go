package xsync

import (
	"sync"
	"sync/atomic"
)

/* __________________________________________________ */

type Once struct {
	mutex atomic.Value
}

func NewOnce() Once {
	var mutex atomic.Value
	mutex.Store(new(sync.Once))
	return Once{
		mutex: mutex,
	}
}

func (o *Once) Do(task func()) {
	o.getOnce().Do(
		func() {
			task()
			o.setOnce()
		},
	)
}

func (o *Once) getOnce() *sync.Once {
	return o.mutex.Load().(*sync.Once)
}

func (o *Once) setOnce() {
	o.mutex.Store(new(sync.Once))
}

/* __________________________________________________ */
