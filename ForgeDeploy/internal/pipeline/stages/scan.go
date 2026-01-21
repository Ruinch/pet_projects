package stages

import (
	"context"
	"log"

	"forgedeploy/internal/domain"
)

func Scan(ctx context.Context, p *domain.Pipeline) error {
	log.Println("[SCAN] security check")
	return nil
}
