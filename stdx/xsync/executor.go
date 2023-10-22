package xsync

import (
	"sync"
)

/* __________________________________________________ */

type Executor interface {
	Add(task func())
	AddAll(tasks []func())
	Execute()
}

/* __________________________________________________ */

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
	ex.doAdd(task)
	ex.lock.Unlock()
}

func (ex *SequentialExecutor) doAdd(task func()) {
	ex.tasks = append(ex.tasks, task)
}

func (ex *SequentialExecutor) AddAll(tasks []func()) {
	ex.lock.Lock()
	ex.doAddAll(tasks)
	ex.lock.Unlock()
}

func (ex *SequentialExecutor) doAddAll(tasks []func()) {
	for _, task := range tasks {
		ex.doAdd(task)
	}
}

func (ex *SequentialExecutor) Execute() {
	ex.lock.Lock()
	ex.doExecute()
	ex.lock.Unlock()
}

func (ex *SequentialExecutor) doExecute() {
	for _, task := range ex.tasks {
		RUN(task)
	}
}

/* __________________________________________________ */

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
	ex.doAdd(task)
	ex.lock.Unlock()
}

func (ex *ParallelExecutor) doAdd(task func()) {
	ex.tasks = append(ex.tasks, task)
}

func (ex *ParallelExecutor) AddAll(tasks []func()) {
	ex.lock.Lock()
	ex.doAddAll(tasks)
	ex.lock.Unlock()
}

func (ex *ParallelExecutor) doAddAll(tasks []func()) {
	for _, task := range tasks {
		ex.doAdd(task)
	}
}

func (ex *ParallelExecutor) Execute() {
	ex.lock.Lock()
	ex.doExecute()
	ex.lock.Unlock()
}

func (ex *ParallelExecutor) doExecute() {
	var wg sync.WaitGroup
	for _, task := range ex.tasks {
		GOWG(task, &wg)
	}
	wg.Wait()
}

/* __________________________________________________ */
