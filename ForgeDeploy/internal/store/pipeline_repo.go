package store

import (
	"context"

	"forgedeploy/internal/domain"
)

type PipelineRepository interface {
	Create(ctx context.Context, p *domain.Pipeline) error
	UpdateStatus(ctx context.Context, id string, status string) error
	SaveStage(
		ctx context.Context,
		pipelineID string,
		stage string,
		status string,
		logs string,
	) error
}
