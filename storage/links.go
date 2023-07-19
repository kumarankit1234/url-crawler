package storage

type LinksStorage interface {
	Add(link string)
	IsPresent(link string) bool
	GetAll() []string
}

func NewLinksStorage() LinksStorage {
	return &linksStorageImpl{
		links: []string{},
	}
}

type linksStorageImpl struct {
	links []string
}

func (l *linksStorageImpl) Add(link string) {
	l.links = append(l.links, link)
}

func (l *linksStorageImpl) IsPresent(link string) bool {
	for _, presentLink := range l.links {
		if presentLink == link {
			return true
		}
	}
	return false
}

func (l *linksStorageImpl) GetAll() []string {
	return l.links
}
