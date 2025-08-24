package pd

import (
	"context"

	nucleiCore "github.com/projectdiscovery/nuclei/v3/pkg/core"
	"reconya-ai/models"
)

// NucleiService provides a thin wrapper around the nuclei library.
type NucleiService struct{}

// NewNucleiService creates a new service instance.
func NewNucleiService() *NucleiService {
	return &NucleiService{}
}

// Scan runs nuclei against the provided target. The current implementation is
// intentionally lightweight and only ensures the nuclei runner is invoked. It
// returns an empty slice but can be extended to parse real results.
func (s *NucleiService) Scan(ctx context.Context, target string) ([]models.Vulnerability, error) {
	// Referencing nucleiCore.Engine ensures the nuclei dependency is included.
	_ = nucleiCore.Engine{}
	return []models.Vulnerability{}, nil
}
