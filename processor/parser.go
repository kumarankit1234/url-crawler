package processor

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"url-crawler/filter"
	"url-crawler/storage"
	task2 "url-crawler/task"
)

type htmlParser struct {
	linksStorage storage.LinksStorage
}

func NewHtmlParser(linksStorage storage.LinksStorage) Processor {
	return &htmlParser{
		linksStorage: linksStorage,
	}
}

func (h *htmlParser) Process(task task2.Task) ([]task2.Task, error) {
	parseTask, ok := task.(task2.ParseTask)
	if !ok {
		return []task2.Task{}, errors.New("not a valid parse task")
	}

	htmlString := parseTask.Blob
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		fmt.Printf("unable to parse html with err %v", err)
		return []task2.Task{}, err
	}

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		links = append(links, link)
	})

	// filter links
	links = h.Filter(links, parseTask.StartingURL)
	if len(links) > 0 {
		fmt.Printf("Processed link %v, and found links %+v\n", parseTask.CurrentURL, links)
	}

	var outputTasks []task2.Task
	for _, link := range links {
		outputTasks = append(outputTasks, task2.DownloadTask{
			URL:         link,
			StartingURL: parseTask.StartingURL,
		})
	}
	return outputTasks, nil
}

func (h *htmlParser) Filter(links []string, startingURL string) []string {
	f := filter.NewFilters(filter.NewValidFilter(), filter.NewSameSubDomainFilter(startingURL), filter.NewAlreadyVisitedFilter(h.linksStorage))
	newLinks := f.Filter(links)
	return newLinks
}
