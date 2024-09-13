package main

import (
	"github.com/janearc/sux/config"
	"github.com/janearc/sux/sux"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	config, err := config.LoadConfig("config/config.yml", "config/version.yml")

	if err != nil {
		log.WithError(err).Fatalf("Failed to load config: %v", err)
	}

	if config.AWS.Region == "" {
		log.Fatalf("AWS region not defined")
	}

	state := sux.NewState()
	if state.IsDefined() {
		log.Info("State is defined")
	} else {
		log.Info("State is not defined")
	}
}
