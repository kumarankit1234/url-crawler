package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"url-crawler/task"
)

func TestQueueImpl_Add(t *testing.T) {
	q := New(10)
	q.Add(task.DownloadTask{URL: "abc.com"})
	assert.Equal(t, 1, q.Size())

	ta := q.Get()
	assert.Equal(t, "abc.com", ta.(task.DownloadTask).URL)
}

func TestQueueImpl_Get_ShouldRemoveElement(t *testing.T) {
	q := New(10)
	q.Add(task.DownloadTask{URL: "abc.com"})
	assert.Equal(t, 1, q.Size())

	ta := q.Get()
	assert.Equal(t, "abc.com", ta.(task.DownloadTask).URL)
	assert.Equal(t, 0, q.Size())
}

func TestQueueImpl_IsEmpty(t *testing.T) {
	q := New(10)
	assert.True(t, q.IsEmpty())

	q.Add(task.DownloadTask{})
	assert.False(t, q.IsEmpty())
}
