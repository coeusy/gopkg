package env

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIsOffline(t *testing.T) {
	assert.Equal(t, IsOffline(), true)
}

func TestIsOnline(t *testing.T) {
	assert.Equal(t, IsOnline(), false)
}
