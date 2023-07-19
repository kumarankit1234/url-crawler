package processor

import (
	"url-crawler/task"
)

//go:generate mockery --name Processor
type Processor interface {
	Process(task task.Task) ([]task.Task, error)
}
