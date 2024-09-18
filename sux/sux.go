package sux

import "github.com/sirupsen/logrus"

func NewSux() *Sux {
	session, err := NewSession()

	if err != nil {
		logrus.WithError(err).Fatalf("Failed to create session: %v", err)
	}
	return &Sux{
		sid: *session,
		// no remote yet
	}
}

// Sux methods

// MarshalData
// Query
// StorageQuery
