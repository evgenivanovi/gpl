package xsync

import (
	"sync"
	"time"
)

type Executor interface {
	Add(task func())
	AddAll(tasks []func())
	Execute()
}

type SequentialExecutor struct {
	lock  sync.Mutex
	tasks []func()
}

func NewSequentialExecutor() *SequentialExecutor {
	tasks := make([]func(), 0)
	return &SequentialExecutor{
		tasks: tasks,
	}
}

func (ex *SequentialExecutor) Add(task func()) {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.doAdd(task)
}

func (ex *SequentialExecutor) doAdd(task func()) {
	ex.tasks = append(ex.tasks, task)
}

func (ex *SequentialExecutor) AddAll(tasks []func()) {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.doAddAll(tasks)
}

func (ex *SequentialExecutor) doAddAll(tasks []func()) {
	for _, task := range tasks {
		ex.doAdd(task)
	}
}

func (ex *SequentialExecutor) Execute() {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.doExecute()
}

func (ex *SequentialExecutor) doExecute() {
	for _, task := range ex.tasks {
		RUN(task)
	}
}

type ParallelExecutor struct {
	lock  sync.Mutex
	tasks []func()
}

func NewParallelExecutor() *ParallelExecutor {
	tasks := make([]func(), 0)
	return &ParallelExecutor{
		tasks: tasks,
	}
}

func (ex *ParallelExecutor) Add(task func()) {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.doAdd(task)
}

func (ex *ParallelExecutor) doAdd(task func()) {
	ex.tasks = append(ex.tasks, task)
}

func (ex *ParallelExecutor) AddAll(tasks []func()) {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.doAddAll(tasks)
}

func (ex *ParallelExecutor) doAddAll(tasks []func()) {
	for _, task := range tasks {
		ex.doAdd(task)
	}
}

func (ex *ParallelExecutor) Execute() {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.doExecute()
}

func (ex *ParallelExecutor) doExecute() {
	var wg sync.WaitGroup
	for _, task := range ex.tasks {
		GOWG(task, &wg)
	}
	wg.Wait()
}

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
