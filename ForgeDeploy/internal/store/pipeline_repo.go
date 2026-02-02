package store

import "forgedeploy/internal/domain"

// PipelineRepository отвечает за pipelines и stages
// Используется Engine и API
type PipelineRepository interface {
	CreatePipeline(p *domain.Pipeline) error
	UpdatePipelineStatus(id int64, status domain.PipelineStatus) error
	GetAll() ([]*domain.Pipeline, error)
	GetByID(id int64) (*domain.Pipeline, error)

	CreateStage(s *domain.Stage) error
	UpdateStage(s *domain.Stage) error
	GetStages(pipelineID int64) ([]*domain.Stage, error)
}
