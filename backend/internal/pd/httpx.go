package pd

import (
	"context"

	httpxRunner "github.com/projectdiscovery/httpx/runner"
)

// HTTPXService is a lightweight wrapper around the httpx runner.
type HTTPXService struct{}

// NewHTTPXService creates a new instance of the service.
func NewHTTPXService() *HTTPXService {
	return &HTTPXService{}
}

// Probe verifies that provided hosts are reachable over HTTP/HTTPS. This
// placeholder implementation ensures the httpx dependency is pulled in while
// returning an empty slice.
func (s *HTTPXService) Probe(ctx context.Context, hosts []string) ([]string, error) {
	_ = httpxRunner.Options{}
	return []string{}, nil
}
