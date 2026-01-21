package pipeline

import (
	"context"
	"log"

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

func (e *Engine) Run(ctx context.Context, p *domain.Pipeline) error {
	log.Println("pipeline started:", p.ID)

	_ = e.repo.Create(ctx, p)
	_ = e.repo.UpdateStatus(ctx, p.ID, "RUNNING")

	runStage := func(name string, fn func(context.Context, *domain.Pipeline) error) error {
		err := fn(ctx, p)
		status := "SUCCESS"
		logs := "ok"

		if err != nil {
			status = "FAILED"
			logs = err.Error()
		}

		_ = e.repo.SaveStage(ctx, p.ID, name, status, logs)
		return err
	}

	if err := runStage("BUILD", stages.Build); err != nil {
		e.repo.UpdateStatus(ctx, p.ID, "FAILED")
		stages.Rollback(ctx, p)
		return err
	}

	if err := runStage("TEST", stages.Test); err != nil {
		e.repo.UpdateStatus(ctx, p.ID, "FAILED")
		stages.Rollback(ctx, p)
		return err
	}

	if err := runStage("SCAN", stages.Scan); err != nil {
		e.repo.UpdateStatus(ctx, p.ID, "FAILED")
		stages.Rollback(ctx, p)
		return err
	}

	if err := runStage("DEPLOY", stages.Deploy); err != nil {
		e.repo.UpdateStatus(ctx, p.ID, "FAILED")
		stages.Rollback(ctx, p)
		return err
	}

	e.repo.UpdateStatus(ctx, p.ID, "SUCCESS")
	log.Println("pipeline finished:", p.ID)
	return nil
}
