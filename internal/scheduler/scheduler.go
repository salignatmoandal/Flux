package scheduler

type Job struct {
	ID       string
	Command  string
	Schedule string
	Priority int
}

type Scheduler struct {
	jobs     []Job
	window   string
	strategy string
}

func (s *Scheduler) Schedule(job Job) error {
	// Logique de planification
	return nil
}
