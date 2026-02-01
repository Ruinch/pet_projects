package store

import "forgedeploy/internal/domain"

type PipelineRepository interface {
	Create(p *domain.Pipeline) error
	UpdateStatus(id int64, status string) error

	CreateStage(s *domain.Stage) error
	UpdateStage(s *domain.Stage) error

	GetAll() ([]domain.Pipeline, error)
}
