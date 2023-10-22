package meta

type AppOp func(metadata *App)

func WithAppNameFn(name func() string) AppOp {
	return func(metadata *App) {
		metadata.Name = name()
	}
}

func WithAppName(name string) AppOp {
	return func(metadata *App) {
		metadata.Name = name
	}
}

func WithAppVersionFn(version func() string) AppOp {
	return func(metadata *App) {
		metadata.Version = version()
	}
}

func WithAppVersion(version string) AppOp {
	return func(metadata *App) {
		metadata.Version = version
	}
}

type App struct {
	Name    string
	Version string
}

func NewApp(
	name string,
	version string,
) App {
	return App{
		Name:    name,
		Version: version,
	}
}

func NewAppWithOps(ops ...AppOp) App {

	app := App{}

	for _, op := range ops {
		op(&app)
	}

	return app

}
