package main

import (
	"time"
	"url-crawler/crawler"
)

func main() {
	urlCrawler := crawler.New(crawler.Options{})
	urlCrawler.Start("https://monzo.com")
	for !urlCrawler.IsDone() {
		time.Sleep(1 * time.Second)
	}
	urlCrawler.Stop()
}
