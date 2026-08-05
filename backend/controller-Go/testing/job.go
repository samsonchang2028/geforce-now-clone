package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID      int
	Message string
}

func (j Job) Process() {
	fmt.Printf("Processing Job #%d: %s\n", j.ID, j.Message)
	time.Sleep(1 * time.Second) // simulate workload
	fmt.Printf("Finished Job #%d\n", j.ID)
}
