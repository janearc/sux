package sux

import (
	"github.com/janearc/sux/config"
	"net/http"
)

type Remote struct {
	// this is just a katamari of stuff we need to talk to sr. altman's finest products
	httpClient *http.Client
	cfg        *config.Config
}

func NewRemote(cfg *config.Config) (*Remote, error) {
	// create a new remote
	remote := &Remote{
		cfg:        cfg,
		httpClient: &http.Client{},
	}

	return remote, nil
}

func GetRemoteUrl(remote *Remote) string {
	return remote.cfg.OpenAI.Url
}
