package cronjob

import "time"

type timeUnit int

const (
	second timeUnit = iota + 1
	minute
	hour
	day
	week
)

func (u timeUnit) String() string {
	return []string{"unknown", "second", "minute", "hour", "day", "week"}[u]
}

type Job struct {
	interval     uint64 // num of unit between job runs
	intervalUnit timeUnit
	err          error                    // err of job
	lastRun      time.Time                // time for job's last run
	nextRun      time.Time                // time for job's next run
	funcs        map[string]interface{}   // key is func name
	funcsParams  map[string][]interface{} // key is func name
}

func (j *Job) NextRun(t time.Time) *Job {
	j.nextRun = t
	return j
}

func (j *Job) setIntervalUnit(unit timeUnit) *Job {
	j.intervalUnit = unit
	return j
}

func (j *Job) setInterval(interval uint64) {
	j.interval = interval
}
