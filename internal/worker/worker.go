package worker

import (
	"alami/pkg/uid"
)

type Worker struct {
	jobCh chan job
}

type job struct {
	id string
	fn func(workerID int) error
}

func worker(workerID int, jobCh chan job) {
	for job := range jobCh {
		job.fn(workerID)
	}
}

func New(maxWorker int) Worker {
	jobCh := make(chan job)

	for i := 1; i <= maxWorker; i++ {
		go worker(i, jobCh)
	}

	return Worker{
		jobCh: jobCh,
	}
}

func (w *Worker) Submit(fn func(workerID int) error) {
	w.jobCh <- job{
		id: uid.GenerateUID(),
		fn: fn,
	}
}
