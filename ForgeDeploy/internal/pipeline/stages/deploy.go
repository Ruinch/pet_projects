package stages

import (
	"context"
	"fmt"
	"log"
	"time"

	"forgedeploy/internal/docker"
	"forgedeploy/internal/domain"
)

func Deploy(ctx context.Context, p *domain.Pipeline) error {
	container := "forgedeploy-app"
	image := fmt.Sprintf("localhost:5000/forgedeploy/app:%s", p.CommitSHA)

	log.Println("[DEPLOY] replacing container")

	// stop old
	docker.StopAndRemove(ctx, container)

	// start new
	if err := docker.RunContainer(ctx, container, image); err != nil {
		return err
	}

	// healthcheck
	if err := docker.WaitHealthy(ctx, "http://localhost:8080/health", 30*time.Second); err != nil {
		return err
	}

	log.Println("[DEPLOY] healthy")
	return nil
}
