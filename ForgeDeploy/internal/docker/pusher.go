package docker

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func PushImage(ctx context.Context, image string) (string, error) {
	cmd := exec.CommandContext(ctx, "docker", "push", image)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return out.String(), fmt.Errorf("docker push failed: %w", err)
	}

	return out.String(), nil
}
