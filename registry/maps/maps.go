package maps

import "github.com/go-sprout/sprout"

type MapsRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of maps registry.
func NewRegistry() *MapsRegistry {
	return &MapsRegistry{}
}

// Uid returns the unique identifier of the registry.
func (mr *MapsRegistry) Uid() string {
	return "time"
}

// LinkHandler links the handler to the registry at runtime.
func (mr *MapsRegistry) LinkHandler(fh sprout.Handler) {
	mr.handler = &fh
}
