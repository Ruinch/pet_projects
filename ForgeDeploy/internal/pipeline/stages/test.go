package stages

import (
	"context"
	"log"

	"forgedeploy/internal/domain"
)

func Test(ctx context.Context, p *domain.Pipeline) error {
	log.Println("[TEST] commit:", p.CommitSHA)
	return nil
}
