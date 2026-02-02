package queue

import (
	"context"
	"github.com/tej2511/go-job-queue/internal/job"
	"log"
)

type Worker struct {
	ID int
}

func (w *Worker) Start(ctx context.Context, jobs <-chan job.Job) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d stopping\n", w.ID)
			return

		case j, ok := <-jobs:
			if !ok {
				log.Printf("Worker %d: job channel closed, stopping\n", w.ID)
				return
			}

			if err := j.Execute(ctx); err != nil {
				log.Printf("Worker %d failed: %v\n", w.ID, err)

			}
		}
	}
}
