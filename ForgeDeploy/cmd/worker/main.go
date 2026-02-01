package main

import (
	"context"
	"log"
	"os"

	"forgedeploy/internal/domain"
	"forgedeploy/internal/pipeline"
	"forgedeploy/internal/store"
)

func main() {
	ctx := context.Background()

	db := store.NewPostgres(os.Getenv("POSTGRES_DSN"))
	repo := store.NewPipelineRepoPostgres(db)

	engine := pipeline.NewEngine(repo)

	if err := store.ApplyMigrations(db); err != nil {
		log.Fatal("failed to apply migrations:", err)
	}

	p := &domain.Pipeline{
		Name:      "pipeline-1",
		CommitSHA: "abc123",
		Status:    domain.PipelinePending,
	}

	engine.Run(ctx, p)
}
