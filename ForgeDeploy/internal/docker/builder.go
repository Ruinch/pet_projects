package docker

import (
	"bytes"
	"context"
	"os/exec"
)

type BuildResult struct {
	Image string
	Logs  string
}

func BuildImage(ctx context.Context, image, dockerfile, contextDir string) (*BuildResult, error) {
	cmd := exec.CommandContext(
		ctx,
		"docker", "build",
		"-f", dockerfile,
		"-t", image,
		contextDir,
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	return &BuildResult{
		Image: image,
		Logs:  out.String(),
	}, err
}
