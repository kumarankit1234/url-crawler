package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiltersImpl_NewValidFilter(t *testing.T) {
	testCases := []struct {
		input  []string
		output []string
		desc   string
	}{
		{
			input:  []string{"trr", "www.monzo.com", "not_cool", "/abc.adsd.com"},
			output: []string{"www.monzo.com"},
			desc:   "links not valid",
		},
		{
			input:  []string{},
			output: []string{},
			desc:   "empty input",
		},
		{
			input:  []string{"https://ab.com"},
			output: []string{"https://ab.com"},
			desc:   "only valid links",
		},
		{
			input:  []string{"bgg"},
			output: []string{},
			desc:   "only invalid links",
		},
	}
	validFilter := NewValidFilter()
	for _, testCase := range testCases {
		assert.Equal(t, validFilter(testCase.input), testCase.output, testCase.desc)
	}
}

func TestFiltersImpl_NewSubdomainFilter(t *testing.T) {
	testCases := []struct {
		input       []string
		output      []string
		startingURL string
		desc        string
	}{
		{
			input:       []string{"www.monzo.com", "www.monzo.com/abc", "www.web.monzo.com"},
			output:      []string{"www.monzo.com", "www.monzo.com/abc"},
			startingURL: "www.monzo.com",
			desc:        "some links not of same domain",
		},
		{
			input:       []string{},
			output:      []string{},
			startingURL: "www.monzo.com",
			desc:        "empty input",
		},
		{
			input:       []string{"https://ab.com/about"},
			output:      []string{"https://ab.com/about"},
			startingURL: "https://ab.com",
			desc:        "all same subdomain",
		},
		{
			input:       []string{"bgg.com"},
			output:      []string{},
			startingURL: "abc.com",
			desc:        "all other subdomain",
		},
	}
	for _, testCase := range testCases {
		subdomainFilter := NewSameSubDomainFilter(testCase.startingURL)
		assert.Equal(t, subdomainFilter(testCase.input), testCase.output, testCase.desc)
	}
}

func TestFilterImpl_NewAlreadyVisitedFilter(t *testing.T) {
	testCases := []struct {
		input          []string
		output         []string
		alreadyVisited []string
		desc           string
	}{
		{
			input:          []string{"www.abc.com", "www.cdf.com"},
			output:         []string{"www.abc.com"},
			alreadyVisited: []string{"www.cdf.com"},
			desc:           "few already visited",
		},
		{
			input:          []string{"www.abc.com", "www.cdf.com"},
			output:         []string{},
			alreadyVisited: []string{"www.abc.com", "www.cdf.com"},
			desc:           "all visited",
		},
		{
			input:          []string{"www.abc.com", "www.cdf.com"},
			output:         []string{"www.abc.com", "www.cdf.com"},
			alreadyVisited: []string{},
			desc:           "none already visited",
		},
		{
			input:          []string{},
			output:         []string{},
			alreadyVisited: []string{},
			desc:           "empty",
		},
	}

	for _, testCase := range testCases {
		mLinks := &mockLinks{
			visitedLinks: testCase.alreadyVisited,
		}
		alreadyVisitedFilter := NewAlreadyVisitedFilter(mLinks)
		assert.Equal(t, alreadyVisitedFilter(testCase.input), testCase.output, testCase.desc)
	}
}

type mockLinks struct {
	visitedLinks []string
}

func (m *mockLinks) Add(link string) {

}

func (m *mockLinks) IsPresent(link string) bool {
	for _, l := range m.visitedLinks {
		if l == link {
			return true
		}
	}
	return false
}

func (m *mockLinks) GetAll() []string {
	return m.visitedLinks
}
