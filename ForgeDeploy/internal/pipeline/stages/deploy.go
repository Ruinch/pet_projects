package stages

import (
	"context"
	"log"

	"forgedeploy/internal/domain"
)

func Deploy(ctx context.Context, p *domain.Pipeline) error {
	log.Println("[DEPLOY] deploying pipeline:", p.ID)
	return nil
}
