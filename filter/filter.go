package filter

import (
	"url-crawler/storage"
	url_checker "url-crawler/url-checker"
)

type Filter func(links []string) []string

type Filters interface {
	Filter(links []string) []string
}

type filtersImpl struct {
	filters []Filter
}

func NewFilters(filters ...Filter) Filters {
	var f []Filter
	for _, filter := range filters {
		f = append(f, filter)
	}
	return &filtersImpl{
		filters: f,
	}
}

func (f *filtersImpl) Filter(links []string) []string {
	for _, filter := range f.filters {
		links = filter(links)
	}
	return links
}

func NewValidFilter() Filter {
	return func(links []string) []string {
		validLinks := []string{}
		for _, link := range links {
			if url_checker.ValidateUrl(link) {
				validLinks = append(validLinks, link)
			}
		}
		return validLinks
	}
}

func NewSameSubDomainFilter(startUrl string) Filter {
	return func(links []string) []string {
		validLinks := []string{}
		for _, link := range links {
			if url_checker.StartsWith(startUrl, link) {
				validLinks = append(validLinks, link)
			}
		}
		return validLinks
	}
}

func NewAlreadyVisitedFilter(linksStorage storage.LinksStorage) Filter {
	return func(links []string) []string {
		validLinks := []string{}
		for _, link := range links {
			if !linksStorage.IsPresent(link) {
				validLinks = append(validLinks, link)
			}
		}
		return validLinks
	}
}
