package main

import (
	"github.com/janearc/sux/sux"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	S := sux.NewSux("config/config.yml", "config/version.yml", "config/secrets.yml")
	if S != nil {
		log.Infof("Sux [%s] instantiated", S.GetVersionBuild())
	}
}
