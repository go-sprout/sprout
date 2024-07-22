package encoding

import (
	"github.com/go-sprout/sprout/registry"
)

type EncodingRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *EncodingRegistry {
	return &EncodingRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *EncodingRegistry) Uid() string {
	return "encoding"
}

// LinkHandler links the handler to the registry at runtime.
func (or *EncodingRegistry) LinkHandler(fh registry.Handler) {
	or.handler = &fh
}
