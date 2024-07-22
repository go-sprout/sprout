package semver

import (
	"github.com/go-sprout/sprout/registry"
)

type SemverRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *SemverRegistry {
	return &SemverRegistry{}
}

// Uid returns the unique identifier of the registry.
func (sr *SemverRegistry) Uid() string {
	return "semver"
}

// LinkHandler links the handler to the registry at runtime.
func (sr *SemverRegistry) LinkHandler(fh registry.Handler) {
	sr.handler = &fh
}
