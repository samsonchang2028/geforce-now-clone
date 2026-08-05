package main

import (
	"context"
	"fmt"
)

type Worker struct {
	ID       int
	JobQueue <-chan Job      // read only channel for jobs chan<- Job is write only chan Job
	Context  context.Context //used for cancel
}

func (w *Worker) Start() {
	go func() {
		fmt.Printf("Worker #%d started \n", w.ID)
		for {
			select {
			case <-w.Context.Done():
				fmt.Printf("worker #%d: stopping.. \n", w.ID)

			case job, ok := <-w.JobQueue:
				if !ok {
					fmt.Printf("Worker #%d: job queue closed... \n", w.ID)
					return
				}
				job.Process()
			}

		}
	}()
}
