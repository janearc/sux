package sux

import (
	"github.com/janearc/sux/config"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

func NewSux(cfn string, vfn string, sfn string) *Sux {
	SUX_ROOT := os.Getenv("SUX_ROOT")

	if SUX_ROOT == "" {
		logrus.Fatalf("SUX_ROOT not defined")
	} else {
		logrus.Infof("SUX_ROOT: %s", SUX_ROOT)
	}

	if cfn == "" {
		cfn = filepath.Join(SUX_ROOT, "config/config.yml")
	}
	if vfn == "" {
		vfn = filepath.Join(SUX_ROOT, "config/version.yml")
	}
	if sfn == "" {
		sfn = filepath.Join(SUX_ROOT, "config/secrets.yml")
	}

	c, err := config.LoadConfig(cfn, vfn, sfn)

	if err != nil {
		logrus.WithError(err).Fatalf("Failed to load config: %v", err)
	}

	if c.AWS.Region == "" {
		log.Fatalf("AWS region not defined")
	}

	state := NewState()
	if state.IsDefined() {
		logrus.Info("State is defined")
	} else {
		logrus.Info("State is not defined")
	}

	session, err := NewSession()

	if err != nil {
		logrus.WithError(err).Fatalf("Failed to create session: %v", err)
	}
	return &Sux{
		sid:    *session,
		Log:    logrus.New(),
		config: c,
		// no remote yet
	}
}

// Sux methods

func (s *Sux) GetVersionBuild() string {
	return s.config.Version.Build
}

// MarshalData
// Query
// StorageQuery
