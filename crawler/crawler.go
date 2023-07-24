package crawler

import (
	"net/http"
	"time"
	"url-crawler/processor"
	"url-crawler/queue"
	"url-crawler/storage"
	"url-crawler/task"
	"url-crawler/workers"
)

const (
	downloaderQueueSize   = 100000
	parserQueueSize       = 100000
	downloaderWorkerCount = 20
	parserWorkerCount     = 10
)

type Crawler interface {
	Start(url string)
	IsDone() bool
	Stop()
}

type Options struct {
	DownloaderQueueSize   int
	ParserQueueSize       int
	DownloaderWorkerCount int
	ParserWorkerCount     int
}

type crawlerImpl struct {
	downloaderQueue queue.Queue
	parserQueue     queue.Queue
	downloadWorkers workers.Workers
	parseWorkers    workers.Workers
	linksStorage    storage.LinksStorage
}

/*
New is used to initialize a crawler
It internally uses two queues
Downloader Queue is the queue where an url is put to be downloaded. It is a FIFO queue
Download workers reads the url from the queue, downloads the html and put it in Parser Queue
Parser workers reads the html from the parser queue, parse the html to get all links,
filter the required links and put it to the downloader queue.
*/
func New(options Options) Crawler {
	dqSize := downloaderQueueSize
	if options.DownloaderQueueSize != 0 {
		dqSize = options.DownloaderQueueSize
	}

	pqSize := parserQueueSize
	if options.ParserQueueSize != 0 {
		pqSize = options.ParserQueueSize
	}

	dWorkerCount := downloaderWorkerCount
	if options.DownloaderWorkerCount != 0 {
		dWorkerCount = options.DownloaderWorkerCount
	}

	pWorkerCount := parserWorkerCount
	if options.ParserWorkerCount != 0 {
		pWorkerCount = options.ParserWorkerCount
	}

	downloaderQueue := queue.New(dqSize)
	parserQueue := queue.New(pqSize)

	linksStorage := storage.NewLinksStorage()

	downloadClient := &http.Client{
		Timeout: 100 * time.Second,
	}
	downloadProcessor := processor.NewDownloadProcessor(downloadClient)
	downloadWorkers := workers.NewConcurrentWorkers(downloaderQueue, parserQueue, downloadProcessor, dWorkerCount)

	parseProcessor := processor.NewHtmlParser(linksStorage)
	parseWorkers := workers.NewConcurrentWorkers(parserQueue, downloaderQueue, parseProcessor, pWorkerCount)
	go downloadWorkers.Start()
	go parseWorkers.Start()

	return &crawlerImpl{
		downloaderQueue: downloaderQueue,
		parserQueue:     parserQueue,
		downloadWorkers: downloadWorkers,
		parseWorkers:    parseWorkers,
		linksStorage:    linksStorage,
	}

}

func (c *crawlerImpl) Start(url string) {
	c.linksStorage.Add(url)
	c.downloaderQueue.Add(task.DownloadTask{
		URL:         url,
		StartingURL: url,
	})
}

func (c *crawlerImpl) Stop() {
	c.downloadWorkers.Stop()
	c.parseWorkers.Stop()
	c.downloaderQueue.Close()
	c.parserQueue.Close()
}

func (c *crawlerImpl) IsDone() bool {
	return c.downloadWorkers.IsDone() && c.parseWorkers.IsDone()
}
