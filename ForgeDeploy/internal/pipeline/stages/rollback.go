package stages

import (
	"context"
	"log"

	"forgedeploy/internal/domain"
)

func Rollback(ctx context.Context, p *domain.Pipeline) {
	log.Println("[ROLLBACK] pipeline:", p.ID)
}
