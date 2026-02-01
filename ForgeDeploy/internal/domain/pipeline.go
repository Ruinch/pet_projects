package domain

import "time"

type PipelineStatus string

const (
	PipelinePending PipelineStatus = "pending"
	PipelineRunning PipelineStatus = "running"
	PipelineSuccess PipelineStatus = "success"
	PipelineFailed  PipelineStatus = "failed"
)

type Pipeline struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	CommitSHA string         `json:"commit_sha"`
	Status    PipelineStatus `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
