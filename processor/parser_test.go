package processor

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	mocks2 "url-crawler/storage/mocks"
	"url-crawler/task"
)

func TestHtmlParser_Filter(t *testing.T) {

}

func TestHtmlParser_ProcessInvalidTask(t *testing.T) {
	mLinks := mocks2.LinksStorage{}
	mLinks.On("Add", mock.Anything)
	dProcessor := NewHtmlParser(&mLinks)
	_, err := dProcessor.Process(task.DownloadTask{})
	assert.NotNil(t, err)
}

func TestHtmlParser_ProcessValidHTML(t *testing.T) {
	link := "https://monzo.com/abc"
	html := "<html><a href=" + link + ">link1</a></html"
	mLinks := mocks2.LinksStorage{}
	mLinks.On("Add", mock.Anything)
	mLinks.On("IsPresent", mock.Anything).Return(false)

	dProcessor := NewHtmlParser(&mLinks)
	newTasks, err := dProcessor.Process(task.ParseTask{Blob: html})
	assert.Nil(t, err)
	assert.Equal(t, len(newTasks), 1)
	parsedTask, ok := newTasks[0].(task.DownloadTask)
	assert.True(t, ok)
	assert.Equal(t, parsedTask.URL, "https://monzo.com/abc")
}

func TestHtmlParser_ProcessValidHTML_AlreadyVisited(t *testing.T) {
	link := "https://monzo.com/abc"
	html := "<html><a href=" + link + ">link1</a></html"
	mLinks := mocks2.LinksStorage{}
	mLinks.On("Add", mock.Anything)
	mLinks.On("IsPresent", mock.Anything).Return(true)

	dProcessor := NewHtmlParser(&mLinks)
	newTasks, err := dProcessor.Process(task.ParseTask{Blob: html})
	assert.Nil(t, err)
	assert.Equal(t, len(newTasks), 0)
}
