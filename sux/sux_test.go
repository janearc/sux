package sux_test

import (
	"github.com/janearc/sux/sux"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testNewState(t *testing.T) {
	state := sux.NewState()
	assert.NotNil(t, state)
	assert.True(t, state.IsDefined())
}
