package build

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Images []BuildConfig `yaml:"images"`
}
type BuildConfig struct {
	Name     string            `yaml:"image"`
	Binaries []BinaryTransport `yaml:"binaries"`
}

type BinaryTransport struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

func ReadConfig(path string) (*Config, error) {
	log.Printf("Using %q as build configuration", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := yaml.Unmarshal(bytes, c); err != nil {
		return nil, err
	}
	return c, nil
}
