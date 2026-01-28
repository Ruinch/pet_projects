package security

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func ScanImage(ctx context.Context, image string) (string, error) {
	cmd := exec.CommandContext(
		ctx,
		"trivy", "image",
		"--severity", "HIGH,CRITICAL",
		"--exit-code", "1",
		image,
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return out.String(), fmt.Errorf("trivy scan failed: %w", err)
	}

	return out.String(), nil
}
