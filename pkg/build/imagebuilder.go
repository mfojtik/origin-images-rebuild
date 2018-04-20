package build

import (
	"bytes"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/openshift/imagebuilder"
	"github.com/openshift/imagebuilder/dockerclient"
)

type imageBuilder struct {
	binarySourceDir string
}

func NewImageBuilder(binarySourceDir string) Interface {
	return &imageBuilder{
		binarySourceDir: binarySourceDir,
	}
}

func (b *imageBuilder) Build(imageName, imageTag string, config *BuildConfig) error {
	client, err := dockerclient.NewClientFromEnv()
	if err != nil {
		return err
	}
	options := dockerclient.NewClientExecutor(nil)
	options.LogFn = func(format string, args ...interface{}) {
		log.Infof("%s", fmt.Sprintf(format, args...))
	}
	options.AuthFn = dockerclient.NoAuthFn
	options.Client = client
	options.AllowPull = true
	options.Tag = fmt.Sprintf("%s:%s", imageName, imageTag)

	binaryPaths := []string{}
	for _, binary := range config.Binaries {
		binaryPaths = append(binaryPaths, binary.ToDockerAdd(b.binarySourceDir))
	}

	dockerfile := bytes.NewBufferString(strings.Join(binaryPaths, "\n"))
	node, err := imagebuilder.ParseDockerfile(dockerfile)

	builder := imagebuilder.NewBuilder(map[string]string{})
	stages := imagebuilder.NewStages(node, builder)

	var stageExecutor *dockerclient.ClientExecutor

	log.Infof("Building %q ...", imageName+":"+imageTag)
	for _, stage := range stages {
		stageExecutor = options.WithName(stage.Name)
		if err := stageExecutor.Prepare(stage.Builder, stage.Node, imageName+":"+imageTag); err != nil {
			return err
		}
		if err := stageExecutor.Execute(stage.Builder, stage.Node); err != nil {
			return err
		}
	}
	return stageExecutor.Commit(stages[len(stages)-1].Builder)
}
