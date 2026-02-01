package security

import (
	"context"
	"os"
	"os/exec"
)

func ScanImage(ctx context.Context, image string) error {
	cmd := exec.CommandContext(
		ctx,
		"docker", "run", "--rm",
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"aquasec/trivy:latest",
		"image",
		"--severity", "CRITICAL",
		"--exit-code", "0",
		image,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
