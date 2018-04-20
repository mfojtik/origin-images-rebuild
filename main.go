package main

import (
	"flag"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/mfojtik/origin-images-rebuild/pkg/build"
	"github.com/mfojtik/origin-images-rebuild/pkg/sets"
)

type OriginImage struct {
	Config build.BuildConfig
}

var openshiftBinary = []build.BinaryTransport{
	{Source: "openshift", Destination: "/usr/bin/openshift"},
}

const (
	latestTag          = "latest"
	openshiftImageRepo = "openshift"
)

var imageConfig = map[string]OriginImage{
	// Base image
	"origin-control-plane": {
		Config: build.BuildConfig{
			Binaries: []build.BinaryTransport{
				{Source: "openshift", Destination: "/usr/bin/openshift"},
				{Source: "oc", Destination: "/usr/bin/oc"},
				{Source: "hypershift", Destination: "/usr/bin/hypershift"},
				{Source: "hyperkube", Destination: "/usr/bin/hyperkube"},
			},
		},
	},
	// cluster-up images
	"origin-deployer":       {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-sti-builder":    {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-docker-builder": {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-haproxy-router": {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-template-service-broker": {
		Config: build.BuildConfig{Binaries: []build.BinaryTransport{
			{Source: "template-service-broker", Destination: "/usr/bin/template-service-broker"},
		}},
	},
	// other images
	"origin-recycler":              {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-f5-router":             {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-nginx-router":          {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-keepalived-ipfailover": {Config: build.BuildConfig{Binaries: openshiftBinary}},
	"origin-node":                  {Config: build.BuildConfig{Binaries: openshiftBinary}},
}

var minimalImages = []string{
	"origin-control-plane",
	"origin-deployer",
	"origin-sti-builder",
	"origin-docker-builder",
	"origin-haproxy-router",
	"origin-template-service-broker",
}

var additionalImages = []string{
	"origin-recycler",
	"origin-f5-router",
	"origin-nginx-router",
	"origin-keepalived-ipfailover",
	"origin-node",
}

func main() {
	buildAll := flag.Bool("--all", false, "Build all OpenShift images")
	tag := flag.String("--tag", latestTag, "Tag to use for built images")

	flag.Parse()

	imagesToBuild := sets.NewString(minimalImages...)
	if buildAll != nil && *buildAll == true {
		imagesToBuild.Insert(additionalImages...)
	}

	outputImageTag := latestTag
	if tag != nil {
		outputImageTag = *tag
	}

	builder := build.NewImageBuilder(filepath.Join("_output", "local", "bin", "linux"))
	for name, image := range imageConfig {
		if !imagesToBuild.Has(name) {
			continue
		}
		err := builder.Build(openshiftImageRepo+"/"+name, outputImageTag, &image.Config)
		if err != nil {
			log.WithField("image", name+":"+outputImageTag).WithError(err).Infof("failed to build image")
		}
	}
}
