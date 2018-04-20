package build

import (
	"fmt"
	"path/filepath"
)

type BinaryTransport struct {
	Source      string
	Destination string
}

func (t BinaryTransport) ToDockerAdd(sourceDir string) string {
	return fmt.Sprintf("ADD %s %s", filepath.Join(sourceDir, t.Source), t.Destination)
}

type BuildConfig struct {
	Binaries []BinaryTransport
}

type Interface interface {
	Build(imageName, imageTag string, config *BuildConfig) error
}
