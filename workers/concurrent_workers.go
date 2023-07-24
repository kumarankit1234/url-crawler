package workers

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"url-crawler/processor"
	"url-crawler/queue"
)

type concurrentWorkers struct {
	inputQueue  queue.Queue
	outputQueue queue.Queue
	processor   processor.Processor
	numWorkers  int
	stopChan    chan struct{}
	ongoingTask int64
}

/*
NewConcurrentWorkers is an implementation of Workers
It reads from the input Queue, call the processor method with the input queue task,
and writes the output to the output queue.
The number of worker that can concurrently can be configured.
*/
func NewConcurrentWorkers(inputQ queue.Queue, outputQ queue.Queue, p processor.Processor, numWorkers int) Workers {
	return &concurrentWorkers{
		inputQueue:  inputQ,
		outputQueue: outputQ,
		processor:   p,
		numWorkers:  numWorkers,
		stopChan:    make(chan struct{}, 1),
		ongoingTask: 0,
	}
}

func (c *concurrentWorkers) Start() {
	wg := sync.WaitGroup{}
	for i := 1; i <= c.numWorkers; i++ {
		wg.Add(1)
		go c.work(&wg)
	}
	wg.Wait()
}

func (c *concurrentWorkers) work(wg *sync.WaitGroup) {
	for {
		select {
		case <-c.stopChan:
			fmt.Println("stopping worker")
			wg.Done()
			break
		default:
			if !c.pollItem() {
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func (c *concurrentWorkers) pollItem() bool {
	atomic.AddInt64(&c.ongoingTask, 1)
	defer atomic.AddInt64(&c.ongoingTask, -1)

	task := c.inputQueue.Get()
	if task == nil {
		return false
	}

	outputTasks, err := c.processor.Process(task)
	if err != nil {
		return true
	}
	for _, oTask := range outputTasks {
		c.outputQueue.Add(oTask)
	}

	return true
}

func (c *concurrentWorkers) Stop() {
	c.stopChan <- struct{}{}
}

func (c *concurrentWorkers) IsDone() bool {
	if !c.inputQueue.IsEmpty() {
		return false
	}
	if c.ongoingTask > 0 {
		return false
	}
	return true
}
