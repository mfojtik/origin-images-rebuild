package build

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

func binaryToDockerAdd(sourceDir string, t BinaryTransport) string {
	return fmt.Sprintf("ADD %s %s", filepath.Join(sourceDir, t.Source), t.Destination)
}

func (b *imageBuilder) Build(imageName, imageTag string, config *BuildConfig) error {
	client, err := dockerclient.NewClientFromEnv()
	if err != nil {
		return err
	}
	options := dockerclient.NewClientExecutor(nil)
	options.LogFn = func(format string, args ...interface{}) {
		log.Printf("%s", fmt.Sprintf(format, args...))
	}
	options.AuthFn = dockerclient.NoAuthFn
	options.Client = client
	options.AllowPull = true
	options.Tag = fmt.Sprintf("%s:%s", imageName, imageTag)
	options.Out, options.ErrOut = os.Stdout, os.Stderr

	defer func() {
		for _, err := range options.Release() {
			log.Printf("unable to clean up the build: %v", err)
		}
	}()

	var binaryPaths []string
	for _, binary := range config.Binaries {
		binaryPaths = append(binaryPaths, binaryToDockerAdd(b.binarySourceDir, binary))
	}

	dockerfile := bytes.NewBufferString(strings.Join(binaryPaths, "\n"))
	node, err := imagebuilder.ParseDockerfile(dockerfile)

	builder := imagebuilder.NewBuilder(map[string]string{})
	stages := imagebuilder.NewStages(node, builder)

	var stageExecutor *dockerclient.ClientExecutor

	log.Printf("Building %q ...", imageName+":"+imageTag)
	for _, stage := range stages {
		stageExecutor = options.WithName(stage.Name)
		if err := stageExecutor.Prepare(stage.Builder, stage.Node, imageName+":"+imageTag); err != nil {
			return fmt.Errorf("%s prepare failed with: %v", stage.Name, err)
		}
		if err := stageExecutor.Execute(stage.Builder, stage.Node); err != nil {
			return fmt.Errorf("%s build execute failed with: %v", stage.Name, err)
		}
	}
	return stageExecutor.Commit(stages[len(stages)-1].Builder)
}
