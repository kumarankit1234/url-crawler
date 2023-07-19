package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew_NoPanic(t *testing.T) {
	assert.NotPanics(t, func() { New(Options{}) })
	c := New(Options{})
	assert.NotPanics(t, func() {
		c.Start("www.abc.com")
	})

	assert.NotPanics(t, c.Stop)
}

func TestStop(t *testing.T) {
	c := New(Options{})
	assert.True(t, c.IsDone())
}
