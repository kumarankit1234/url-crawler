package processor

import (
	"url-crawler/task"
)

type Processor interface {
	Process(task task.Task) ([]task.Task, error)
}
