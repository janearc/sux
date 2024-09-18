package main

import (
	"github.com/janearc/sux/config"
	"github.com/janearc/sux/sux"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {
	log := logrus.New()

	SUX_ROOT := os.Getenv("SUX_ROOT")

	if SUX_ROOT == "" {
		log.Fatalf("SUX_ROOT not defined")
	} else {
		log.Infof("SUX_ROOT: %s", SUX_ROOT)
	}

	cfName := filepath.Join(SUX_ROOT, "config/config.yml")
	vfName := filepath.Join(SUX_ROOT, "config/version.yml")

	config, err := config.LoadConfig(cfName, vfName)

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
