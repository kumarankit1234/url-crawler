package workers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"url-crawler/processor/mocks"
	mocks2 "url-crawler/queue/mocks"
)

func TestNewConcurrentWorkers_NotPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		NewConcurrentWorkers(&mocks2.Queue{}, &mocks2.Queue{}, &mocks.Processor{}, 10)
	})
}

func TestNewConcurrentWorkers_NotPanic_Stop(t *testing.T) {
	mockQueue := mocks2.Queue{}
	mockQueue.On("Get").Return(nil)
	workers := NewConcurrentWorkers(&mockQueue, &mocks2.Queue{}, &mocks.Processor{}, 10)

	assert.NotPanics(t, func() {
		workers.Stop()
	})
}
