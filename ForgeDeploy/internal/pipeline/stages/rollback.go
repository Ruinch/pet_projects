package stages

import (
	"context"
	"log"

	"forgedeploy/internal/docker"
	"forgedeploy/internal/domain"
)

func Rollback(ctx context.Context, p *domain.Pipeline) {
	container := "forgedeploy-app"
	log.Println("[ROLLBACK] stopping container")
	docker.StopAndRemove(ctx, container)
}
