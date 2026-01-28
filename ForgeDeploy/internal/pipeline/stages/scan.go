package stages

import (
	"context"
	"fmt"
	"log"

	"forgedeploy/internal/domain"
	"forgedeploy/internal/security"
)

func Scan(ctx context.Context, p *domain.Pipeline) error {
	image := fmt.Sprintf("forgedeploy/app:%s", p.CommitSHA)

	log.Println("[SCAN] trivy image:", image)

	logs, err := security.ScanImage(ctx, image)
	if err != nil {
		log.Println(logs)
		return err
	}

	log.Println("[SCAN] passed")
	return nil
}
