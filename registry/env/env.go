package env

import (
	"github.com/go-sprout/sprout/registry"
)

type EnvironmentRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of env registry.
func NewRegistry() *EnvironmentRegistry {
	return &EnvironmentRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *EnvironmentRegistry) Uid() string {
	return "env"
}

// LinkHandler links the handler to the registry at runtime.
func (or *EnvironmentRegistry) LinkHandler(fh registry.Handler) {
	or.handler = &fh
}
