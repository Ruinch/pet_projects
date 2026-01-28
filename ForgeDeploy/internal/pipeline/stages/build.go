package stages

import (
	"context"
	"fmt"
	"log"

	"forgedeploy/internal/docker"
	"forgedeploy/internal/domain"
)

func Build(ctx context.Context, p *domain.Pipeline) error {
	// image := fmt.Sprintf("forgedeploy/app:%s", p.CommitSHA)
	image := fmt.Sprintf("localhost:5000/forgedeploy/app:%s", p.CommitSHA)
	log.Println("[BUILD] docker image:", image)

	_, err := docker.BuildImage(
		ctx,
		image,
		"docker/app.Dockerfile",
		".",
	)

	return err
}
