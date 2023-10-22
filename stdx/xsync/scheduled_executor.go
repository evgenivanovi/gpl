package xsync

import (
	"sync"
	"time"
)

/* __________________________________________________ */

type ScheduledExecutor struct {
	interval time.Duration
	task     func()

	stopChan      chan struct{}
	stopChanActor *sync.Once
}

func NewScheduledExecutor(
	interval time.Duration,
	task func(),
) *ScheduledExecutor {
	return &ScheduledExecutor{
		interval: interval,
		task:     task,

		stopChan:      make(chan struct{}),
		stopChanActor: new(sync.Once),
	}
}

func (ex *ScheduledExecutor) Execute() {

	ticker := time.NewTicker(ex.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			{
				RUN(ex.task)
			}
		case <-ex.stopChan:
			{
				return
			}
		}
	}

}

func (ex *ScheduledExecutor) Close() {
	action := func() { close(ex.stopChan) }
	ex.stopChanActor.Do(action)
}

/* __________________________________________________ */
