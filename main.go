package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mfojtik/origin-images-rebuild/pkg/build"
)

var (
	defaultImageTag  = "latest"
	defaultImageRepo = "openshift"
)

func getOriginPath() (string, error) {
	if len(os.Getenv("GOPATH")) == 0 {
		return "", fmt.Errorf("GOPATH cannot be empty")
	}
	originDir := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "openshift", "origin")
	if _, err := os.Stat(originDir); err != nil {
		return "", fmt.Errorf("unable to stat %q: %v", originDir, err)
	}
	return originDir, nil
}

func main() {
	originDir, err := getOriginPath()
	if err != nil {
		log.Fatalf("Unable to determine origin directory: %v", err)
	}
	configPath := filepath.Join(originDir, "images.yaml")

	configPathFlag := flag.String("config", configPath, "A path to config file to use for rebuilding images")
	tagFlag := flag.String("tag", defaultImageTag, "Image tag to use")
	repoFlag := flag.String("repo", defaultImageRepo, "Repository prefix to use")

	flag.Parse()

	if configPathFlag != nil {
		configPath = *configPathFlag
	}

	config, err := build.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	if tagFlag != nil {
		defaultImageTag = *tagFlag
	}

	if repoFlag != nil {
		defaultImageRepo = *repoFlag
	}

	builder := build.NewImageBuilder(filepath.Join(originDir, "_output", "local", "bin", "linux", "amd64"))

	for _, image := range config.Images {
		err := builder.Build(defaultImageRepo+"/origin-"+image.Name, defaultImageTag, &image)
		if err != nil {
			log.Printf("Image %q failed to build: %v", image.Name, err)
		}
	}
}
