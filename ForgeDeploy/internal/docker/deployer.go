package docker

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func RunContainer(ctx context.Context, name, image string) error {
	cmd := exec.CommandContext(
		ctx,
		"docker", "run", "-d",
		"--name", name,
		"-p", "8080:8080",
		image,
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker run failed: %w\n%s", err, out.String())
	}
	return nil
}

func StopAndRemove(ctx context.Context, name string) {
	exec.Command("docker", "rm", "-f", name).Run()
}
