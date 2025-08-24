package pd

import (
	"context"

	subRunner "github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

// SubfinderService wraps the ProjectDiscovery subfinder runner.
type SubfinderService struct{}

// NewSubfinderService returns a new instance.
func NewSubfinderService() *SubfinderService {
	return &SubfinderService{}
}

// Enumerate discovers subdomains for the given domain. For now this function
// simply references the subfinder runner to keep the dependency and returns an
// empty slice.
func (s *SubfinderService) Enumerate(ctx context.Context, domain string) ([]string, error) {
	_ = subRunner.Options{}
	return []string{}, nil
}
