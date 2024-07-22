package slices

import (
	"github.com/go-sprout/sprout/registry"
)

type SlicesRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *SlicesRegistry {
	return &SlicesRegistry{}
}

// Uid returns the unique identifier of the registry.
func (sr *SlicesRegistry) Uid() string {
	return "slices"
}

// LinkHandler links the handler to the registry at runtime.
func (sr *SlicesRegistry) LinkHandler(fh registry.Handler) {
	sr.handler = &fh
}
