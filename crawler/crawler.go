package crawler

import (
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
	downloaderWorkerCount = 10
	parserWorkerCount     = 5
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

	downloadProcessor := processor.NewDownloadProcessor(100*time.Second, linksStorage)
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
