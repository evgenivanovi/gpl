package fw

import (
	"sync"

	"github.com/evgenivanovi/gpl/stdx/xsync"
)

type Application struct {
	Settings ServerSettings
	Context  map[string]any

	onStartExec           xsync.Executor
	onStartBackgroundExec xsync.Executor

	onCloseWG             *sync.WaitGroup
	onCloseBackgroundExec xsync.Executor
	onCloseExec           xsync.Executor
}

func NewApplication() *Application {
	return &Application{
		Context: make(map[string]any),

		onStartExec:           xsync.NewParallelExecutor(),
		onStartBackgroundExec: xsync.NewParallelExecutor(),

		onCloseWG:             &sync.WaitGroup{},
		onCloseBackgroundExec: xsync.NewParallelExecutor(),
		onCloseExec:           xsync.NewParallelExecutor(),
	}
}

func (a *Application) Put(key string, value any) {
	a.Context[key] = value
}

func (a *Application) Get(key string) any {
	return a.Context[key]
}

func (a *Application) RegisterOnStart(task func()) {
	a.onStartExec.Add(task)
}

func (a *Application) RegisterOnStartBackground(task func()) {
	a.onStartBackgroundExec.Add(task)
}

func (a *Application) Start() {
	// sync
	a.onStartExec.Execute()
	// parallel
	xsync.GO(a.onStartBackgroundExec.Execute)
}

func (a *Application) RegisterOnCloseBackground(task func()) {
	a.onCloseBackgroundExec.Add(task)
}

func (a *Application) RegisterOnClose(task func()) {
	a.onCloseExec.Add(task)
}

func (a *Application) Close() {
	// parallel
	xsync.GOWG(a.onCloseBackgroundExec.Execute, a.onCloseWG)
	// wait
	a.onCloseWG.Wait()
	// sync
	a.onCloseExec.Execute()
}
