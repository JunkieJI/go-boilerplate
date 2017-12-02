package store

import (
	"testing"

	"github.com/JunkieJI/go-boilerplate/config"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	s := NewStore(config.GetDSN())
	assert.True(t, s.Ping(), "Should be no errors with ping")
}
