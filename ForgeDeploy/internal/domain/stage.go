package domain

import "time"

type Stage struct {
	ID         int64
	PipelineID int64
	Name       string
	Status     string
	Logs       *string
	StartedAt  *time.Time
	FinishedAt *time.Time
}
