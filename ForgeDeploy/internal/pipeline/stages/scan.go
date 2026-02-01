package stages

import (
	"context"
	"fmt"
	"log"

	"forgedeploy/internal/domain"
	"forgedeploy/internal/security"
)

func Scan(ctx context.Context, p *domain.Pipeline) error {
	image := fmt.Sprintf("localhost:5000/forgedeploy/app:%s", p.CommitSHA)

	log.Println("[SCAN] trivy image:", image)

	if err := security.ScanImage(ctx, image); err != nil {
		log.Println("[SCAN] error:", err) // 🔥 ВАЖНО
		return err
	}

	log.Println("[SCAN] passed")
	return nil
}
