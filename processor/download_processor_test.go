package processor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"strings"
	"testing"
	"url-crawler/processor/mocks"
	"url-crawler/task"
)

func TestDownloadProcessor_ProcessInvalidTask(t *testing.T) {
	mDownloader := mocks.Downloader{}
	mDownloader.On("Get", mock.Anything).Return("")
	dProcessor := NewDownloadProcessor(&mDownloader)
	_, err := dProcessor.Process(task.ParseTask{})
	assert.NotNil(t, err)
}

func TestDownloadProcessor_ProcessInvalidURL(t *testing.T) {
	mDownloader := mocks.Downloader{}
	mDownloader.On("Get", mock.Anything).Return("")
	dProcessor := NewDownloadProcessor(&mDownloader)
	newTasks, err := dProcessor.Process(task.DownloadTask{URL: "abc/hu"})
	assert.Nil(t, err)
	assert.Equal(t, len(newTasks), 0)
}

func TestDownloadProcessor_ProcessValidURL(t *testing.T) {
	html := "<html><a href='https://monzo.com/abc'>link1</a></html"
	mDownloader := mocks.Downloader{}
	downloadRes := &http.Response{
		Body: io.NopCloser(strings.NewReader(html)),
	}
	mDownloader.On("Get", mock.Anything).Return(downloadRes, nil)
	dProcessor := NewDownloadProcessor(&mDownloader)
	newTasks, err := dProcessor.Process(task.DownloadTask{URL: "https://monzo.com"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(newTasks))
	parsedTask, ok := newTasks[0].(task.ParseTask)
	assert.True(t, ok)
	assert.Equal(t, parsedTask.Blob, html)
}
