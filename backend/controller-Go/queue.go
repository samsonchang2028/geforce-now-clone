package main

type Queue struct {
	buffer []Job
}

func (q *Queue) Enqueue(val Job) {
	q.buffer = append(q.buffer, val)
}

func (q *Queue) Dequeue() Job {
	returnVal := q.buffer[0]

	q.buffer = q.buffer[1:]

	return returnVal

}
