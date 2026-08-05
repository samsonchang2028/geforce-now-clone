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
