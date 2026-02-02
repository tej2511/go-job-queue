package queue

import (
	"errors"
	"github.com/tej2511/go-job-queue/internal/job"
	"sync"
)

var ErrQueueClosed = errors.New("queue is closed")

type Queue struct {
	jobs   chan job.Job
	closed bool
	mu     sync.Mutex
}

func New(size int) *Queue {
	return &Queue{
		jobs: make(chan job.Job, size),
	}
}

func (q *Queue) Enqueue(j job.Job) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return ErrQueueClosed
	}
	q.jobs <- j
	return nil
}

func (q *Queue) Jobs() <-chan job.Job {
	return q.jobs
}

func (q *Queue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if !q.closed {
		close(q.jobs)
		q.closed = true
	}
}
