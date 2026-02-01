package store

import "forgedeploy/internal/domain"

// PipelineRepository отвечает за pipelines и stages
// Используется Engine и API
type PipelineRepository interface {
	// pipelines
	CreatePipeline(p *domain.Pipeline) error
	UpdatePipelineStatus(id int64, status domain.PipelineStatus) error
	GetAll() ([]*domain.Pipeline, error)

	// stages
	CreateStage(s *domain.Stage) error
	UpdateStage(s *domain.Stage) error
}
