package sux

import "github.com/google/uuid"

// session constructor and methods

func NewSession() (*Session, error) {
	// create a new session
	return &Session{}, nil
}

func (s *Session) ToUUID() uuid.UUID {
	return s.sid
}
