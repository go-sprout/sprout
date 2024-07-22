package reflect

import "github.com/go-sprout/sprout/registry"

type ReflectRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of reflect registry.
func NewRegistry() *ReflectRegistry {
	return &ReflectRegistry{}
}

// Uid returns the unique identifier of the registry.
func (rr *ReflectRegistry) Uid() string {
	return "reflect"
}

// LinkHandler links the handler to the registry at runtime.
func (rr *ReflectRegistry) LinkHandler(fh registry.Handler) {
	rr.handler = &fh
}
