package docker

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func WaitHealthy(ctx context.Context, url string, timeout time.Duration) error {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timeoutCh := time.After(timeout)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timeoutCh:
			return fmt.Errorf("healthcheck timeout")
		case <-ticker.C:
			resp, err := http.Get(url)
			if err == nil && resp.StatusCode == http.StatusOK {
				return nil
			}
		}
	}
}
