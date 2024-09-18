package sux

import (
	"github.com/janearc/sux/config"
	"net/http"
)

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
