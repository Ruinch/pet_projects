package domain

import "time"

type PipelineStatus string

const (
	StatusPending PipelineStatus = "PENDING"
	StatusRunning PipelineStatus = "RUNNING"
	StatusFailed  PipelineStatus = "FAILED"
	StatusSuccess PipelineStatus = "SUCCESS"
)

type Pipeline struct {
	ID        string
	Repo      string
	CommitSHA string
	Status    PipelineStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
