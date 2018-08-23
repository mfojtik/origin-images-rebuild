package build

type Interface interface {
	Build(imageName, imageTag string, config *BuildConfig) error
}
