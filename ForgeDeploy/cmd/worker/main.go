package main

import (
	"context"
	"os"

	"forgedeploy/internal/domain"
	"forgedeploy/internal/pipeline"
	"forgedeploy/internal/store"
)

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	db := store.NewPostgres(dsn)
	repo := store.NewPipelineRepo(db)

	engine := pipeline.NewEngine(repo)

	p := &domain.Pipeline{
		ID:        "pipeline-1",
		Repo:      "example/repo",
		CommitSHA: "abc123",
		Status:    "PENDING",
	}

	engine.Run(context.Background(), p)
}
