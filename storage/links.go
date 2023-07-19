package storage

import "sync"

type LinksStorage interface {
	Add(link string)
	IsPresent(link string) bool
}

func NewLinksStorage() LinksStorage {
	return &linksStorageImpl{}
}

type linksStorageImpl struct {
	linkMap sync.Map
}

func (l *linksStorageImpl) Add(link string) {
	l.linkMap.Store(link, true)
}

func (l *linksStorageImpl) IsPresent(link string) bool {
	entry, found := l.linkMap.Load(link)
	if found && entry.(bool) {
		return true
	}
	return false
}
