package processor

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	task2 "url-crawler/task"
	url_checker "url-crawler/url-checker"
)

type downloadProcessor struct {
	downloader Downloader
}

//go:generate mockery --name Downloader
type Downloader interface {
	Get(url string) (resp *http.Response, err error)
}

func NewDownloadProcessor(downloader Downloader) Processor {
	return &downloadProcessor{
		downloader: downloader,
	}
}

func (d *downloadProcessor) Process(task task2.Task) ([]task2.Task, error) {
	downloadTask, ok := task.(task2.DownloadTask)
	if !ok {
		return []task2.Task{}, errors.New("not a valid download task")
	}
	currentUrl := downloadTask.URL
	isValid := url_checker.ValidateUrl(currentUrl)
	if !isValid {
		fmt.Printf("Unable to parse this url %v, with error, not valid url\n", currentUrl)
		return []task2.Task{}, nil
	}

	resp, err := d.downloader.Get(currentUrl)
	if err != nil {
		fmt.Println(err)
		return []task2.Task{}, err
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []task2.Task{}, err
	}
	return []task2.Task{
		task2.ParseTask{
			Blob:        string(resBody),
			StartingURL: downloadTask.StartingURL,
			CurrentURL:  currentUrl,
		},
	}, nil
}
