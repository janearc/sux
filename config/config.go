package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	AWS struct {
		Region string `yaml:"region"`
	} `yaml:"aws"`
	OpenAI struct {
		APIKey string
	} `yaml:"openai"`
}

func LoadConfig(configFileName string, versionFileName string) (*Config, error) {
	// it seems unlikely that we'll be deployed in docker but this is
	// a good default
	if configFileName == "" {
		configFileName = "/app/config/config.yml"
	}

	file, err := os.Open(configFileName)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to open config file %s", configFileName)
		return nil, err
	}
	defer file.Close()

	var config Config

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		logrus.WithError(err).Fatalf("Failed to decode config file %s", configFileName)
		return nil, err
	}

	// same as above, but for the version file
	if versionFileName == "" {
		versionFileName = "/app/config/version.yml"
	}
	vf, err := os.Open(versionFileName)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to open version file %s", versionFileName)
		return nil, err
	}
	defer vf.Close()

	// Decode the version file
	vfDecoder := yaml.NewDecoder(vf)
	if err := vfDecoder.Decode(&config); err != nil {
		logrus.WithError(err).Fatalf("Failed to decode version file %s", versionFileName)
		return nil, err
	}

	return &config, nil
}
