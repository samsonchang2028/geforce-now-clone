package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Job struct {
	ID      int
	Message string
}

func (j Job) Process() {
	fmt.Printf("Processing Job #%d: %s\n", j.ID, j.Message)
	delay := time.Duration(rand.IntN(7)+4) * time.Second
	time.Sleep(delay)
	fmt.Printf("Finished Job #%d\n", j.ID)
}
