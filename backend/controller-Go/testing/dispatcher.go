package main

import (
	"context"
	"sync"
)

type Dispatcher struct {
	WorkerCount int
	JobQueue    chan Job
	workers     []*Worker
	ctx         context.Context
	cancel      context.CancelFunc
	wg          *sync.WaitGroup
}

//creates and initializes a dispatcher

func NewDispatcher(workerCount int, queueSize int) *Dispatcher {
	ctx, cancel := context.WithCancel(context.Background())

	return &Dispatcher{
		WorkerCount: workerCount,
		JobQueue:    make(chan Job, queueSize),
		ctx:         ctx,
		cancel:      cancel,
		wg:          &sync.WaitGroup{},
	}
}

func (d *Dispatcher) Start() {

	for i := 1; i <= d.WorkerCount; i++ {
		worker := &Worker{
			ID:       i,
			JobQueue: d.JobQueue,
			Context:  d.ctx,
		}
		d.workers = append(d.workers, worker)
		worker.Start()
	}
}

// enqueues a job onto the queue
func (d *Dispatcher) Submit(job Job) {
	d.wg.Add(1)
	go func() {
		d.JobQueue <- job // we are putting the job from the input into the go channel buffer
		d.wg.Done()
	}()

}

func (d *Dispatcher) Stop() {
	d.cancel()
	d.wg.Wait()

	close(d.JobQueue)
}
