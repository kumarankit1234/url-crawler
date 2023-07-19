package queue

import (
	"url-crawler/task"
)

type Queue interface {
	Add(task task.Task)
	Get() task.Task
	Close()
	IsEmpty() bool
	Size() int
}

type queueImpl struct {
	queue chan task.Task
}

func New(size int) Queue {
	return &queueImpl{
		queue: make(chan task.Task, size),
	}
}

func (q *queueImpl) Add(task task.Task) {
	q.queue <- task
}

func (q *queueImpl) Get() task.Task {
	if q.IsEmpty() {
		return nil
	}
	return <-q.queue
}

func (q *queueImpl) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queueImpl) Size() int {
	return len(q.queue)
}

func (q *queueImpl) Close() {
	close(q.queue)
}
