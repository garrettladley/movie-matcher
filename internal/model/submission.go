package model

import "time"

type Submission struct {
	Score int
	Time  time.Time `db:"submission_time"`
}
