package store

import "forgedeploy/internal/domain"

type PipelineRepository interface {
	// pipeline
	Create(p *domain.Pipeline) error
	UpdatePipelineStatus(id int64, status string) error
	GetAll() ([]*domain.Pipeline, error)

	// stages
	CreateStage(s *domain.Stage) error
	UpdateStage(s *domain.Stage) error
}
