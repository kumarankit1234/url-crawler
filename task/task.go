package task

type Task interface{}

type DownloadTask struct {
	URL         string
	StartingURL string
}

type ParseTask struct {
	CurrentURL  string
	Blob        string
	StartingURL string
}
