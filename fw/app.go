package fw

import (
	"github.com/evgenivanovi/gpl/stdx/xsync"
)

type Application struct {
	Settings ServerSettings
	Context  map[string]any

	onStartTasks           []func()
	onStartBackgroundTasks []func()

	onCloseTasks           []func()
	onCloseBackgroundTasks []func()
}

func NewApplication() *Application {
	return &Application{
		Context: make(map[string]any),

		onStartTasks:           make([]func(), 0),
		onStartBackgroundTasks: make([]func(), 0),

		onCloseTasks:           make([]func(), 0),
		onCloseBackgroundTasks: make([]func(), 0),
	}
}

func (a *Application) Put(key string, value any) {
	a.Context[key] = value
}

func (a *Application) Get(key string) any {
	return a.Context[key]
}

func (a *Application) RegisterOnStart(task func()) {
	a.onStartTasks = append(a.onStartTasks, task)
}

func (a *Application) RegisterOnStartBackground(task func()) {
	a.onStartBackgroundTasks = append(a.onStartBackgroundTasks, task)
}

func (a *Application) Start() {
	xsync.ExecuteSequential(a.onStartTasks...)
	xsync.ExecuteParallel(a.onStartBackgroundTasks...)
}

func (a *Application) RegisterOnClose(task func()) {
	a.onCloseTasks = append(a.onCloseTasks, task)
}

func (a *Application) RegisterOnCloseBackground(task func()) {
	a.onCloseBackgroundTasks = append(a.onCloseBackgroundTasks, task)
}

func (a *Application) Close() {
	xsync.ExecuteParallel(a.onCloseBackgroundTasks...)
	xsync.ExecuteSequential(a.onCloseTasks...)
}
