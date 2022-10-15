package cronjob

import "time"

type Scheduler struct {
	jobs []*Job
	size int
	loc  *time.Location
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		loc: time.Local,
	}
}
