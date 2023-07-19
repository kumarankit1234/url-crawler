package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinksStorageImpl_Add(t *testing.T) {
	l := NewLinksStorage()
	link := "www.abc.com"
	l.Add(link)
	assert.True(t, l.IsPresent("www.abc.com"))
}

func TestLinksStorageImpl_IsPresent(t *testing.T) {
	l := NewLinksStorage()
	firstLink := "www.abc.com"
	secondLink := "www.gh.com"
	l.Add(firstLink)
	assert.True(t, l.IsPresent(firstLink))
	assert.False(t, l.IsPresent(secondLink))

	l.Add(secondLink)
	assert.True(t, l.IsPresent(secondLink))
}
