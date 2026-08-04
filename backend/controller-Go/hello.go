package main

import (
	"errors"
	"fmt"
)

// properties of a queue

// First in first out

// popping from the ehad

// enq  = insert at the back of the queue
// deq = remove the thing at the head

type Queue struct {
	items []any
}

func (q *Queue) enqueue(input any) {

	q.items = append(q.items, input)
}

func (q *Queue) dequeue() (any, error) {

	if len(q.items) == 0 {
		return 0, errors.New("queue is empty, cannot dequeue")
	}

	front := q.items[0]

	q.items = q.items[1:]

	return front, nil

}

func main() {

	queue := Queue{}

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(7)
	fmt.Println(queue.items)
	queue.dequeue()
	queue.dequeue()
	fmt.Println(queue.items)

}
