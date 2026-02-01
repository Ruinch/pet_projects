package stages

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"forgedeploy/internal/docker"
	"forgedeploy/internal/domain"
)

func Build(ctx context.Context, p *domain.Pipeline) error {
	image := fmt.Sprintf("localhost:5000/forgedeploy/app:%s", p.CommitSHA)

	// ✅ ТОЛЬКО env, НИКАКИХ Getwd / Dir / магии
	projectRoot := os.Getenv("PROJECT_ROOT")
	if projectRoot == "" {
		return fmt.Errorf("PROJECT_ROOT environment variable is not set")
	}

	dockerfilePath := filepath.Join(projectRoot, "docker", "app.Dockerfile")

	log.Println("[BUILD] docker image:", image)
	log.Println("[BUILD] project root:", projectRoot)
	log.Println("[BUILD] dockerfile:", dockerfilePath)

	res, err := docker.BuildImage(
		ctx,
		image,
		dockerfilePath,
		projectRoot,
	)

	if err != nil {
		log.Println("[BUILD] docker build failed")
		log.Println(res.Logs)
		return err
	}

	log.Println("[BUILD] docker build success")
	return nil
}
