package meta

/* __________________________________________________ */

type AppMetadataOp func(metadata *AppMetadata)

func WithAppNameFn(name func() string) AppMetadataOp {
	return func(metadata *AppMetadata) {
		metadata.Name = name()
	}
}

func WithAppName(name string) AppMetadataOp {
	return func(metadata *AppMetadata) {
		metadata.Name = name
	}
}

func WithAppVersionFn(version func() string) AppMetadataOp {
	return func(metadata *AppMetadata) {
		metadata.Version = version()
	}
}

func WithAppVersion(version string) AppMetadataOp {
	return func(metadata *AppMetadata) {
		metadata.Version = version
	}
}

/* __________________________________________________ */

type AppMetadata struct {
	Name    string
	Version string
}

func NewAppMetadata(
	name string,
	version string,
) AppMetadata {
	return AppMetadata{
		Name:    name,
		Version: version,
	}
}

func NewAppMetadataWithOps(ops ...AppMetadataOp) AppMetadata {

	app := AppMetadata{}

	for _, op := range ops {
		op(&app)
	}

	return app

}

/* __________________________________________________ */
