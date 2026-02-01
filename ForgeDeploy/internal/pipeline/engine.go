package pipeline

import (
	"context"
	"log"
	"time"

	"forgedeploy/internal/domain"
	"forgedeploy/internal/pipeline/stages"
	"forgedeploy/internal/store"
)

type Engine struct {
	repo store.PipelineRepository
}

func NewEngine(repo store.PipelineRepository) *Engine {
	return &Engine{repo: repo}
}

func (e *Engine) Run(ctx context.Context, p *domain.Pipeline) {
	log.Println("pipeline started:", p.Name)

	// 1. Create pipeline
	if err := e.repo.CreatePipeline(p); err != nil {
		log.Println("failed to create pipeline:", err)
		return
	}

	_ = e.repo.UpdatePipelineStatus(p.ID, domain.PipelineRunning)

	stageDefs := []struct {
		name string
		fn   func(context.Context, *domain.Pipeline) error
	}{
		{"build", stages.Build},
		{"test", stages.Test},
		{"scan", stages.Scan},
		{"deploy", stages.Deploy},
	}

	for _, s := range stageDefs {
		stage := &domain.Stage{
			PipelineID: p.ID,
			Name:       s.name,
			Status:     domain.StagePending,
		}

		_ = e.repo.CreateStage(stage)

		start := time.Now()
		stage.StartedAt = &start
		stage.Status = domain.StageRunning
		_ = e.repo.UpdateStage(stage)

		err := s.fn(ctx, p)

		finish := time.Now()
		stage.FinishedAt = &finish

		if err != nil {
			stage.Status = domain.StageFailed
			msg := err.Error()
			stage.Logs = &msg

			_ = e.repo.UpdateStage(stage)
			_ = e.repo.UpdatePipelineStatus(p.ID, domain.PipelineFailed)

			log.Println("pipeline failed on stage:", s.name, "error:", err)
			return
		}

		stage.Status = domain.StageSuccess
		_ = e.repo.UpdateStage(stage)
	}

	_ = e.repo.UpdatePipelineStatus(p.ID, domain.PipelineSuccess)
	log.Println("pipeline finished:", p.Name)
}
