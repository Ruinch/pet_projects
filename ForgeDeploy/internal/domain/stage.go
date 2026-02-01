package domain

import "time"

type StageStatus string

const (
	StagePending StageStatus = "pending"
	StageRunning StageStatus = "running"
	StageFailed  StageStatus = "failed"
	StageSuccess StageStatus = "success"
)

type Stage struct {
	ID         int64
	PipelineID int64
	Name       string
	Status     StageStatus

	Logs       *string
	StartedAt  *time.Time
	FinishedAt *time.Time
}
