package sux

import (
	"github.com/google/uuid"
	"github.com/janearc/sux/config"
	"net/http"
	"time"
)

type Sux struct {
	sid    Session
	remote Remote
}

type State struct {
	Defined bool
}

type Session struct {
	// unfortunately this is probably going to be bound to aws zones
	sid uuid.UUID

	// this is the aws region
	zone string
}

type Remote struct {
	// this is just a katamari of stuff we need to talk to sr. altman's finest products
	httpClient *http.Client
	cfg        *config.Config
}

type Metadata struct {
	creationDate time.Time
	updatedDate  time.Time
	size         int // in bytes
}

type Thing struct {
	sid     Session
	summary string
	name    string
}
