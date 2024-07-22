package uniqueid

import (
	"github.com/go-sprout/sprout/registry"
)

type UniqueIDRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *UniqueIDRegistry {
	return &UniqueIDRegistry{}
}

// Uid returns the unique identifier of the registry.
func (ur *UniqueIDRegistry) Uid() string {
	return "uniqueid"
}

// LinkHandler links the handler to the registry at runtime.
func (ur *UniqueIDRegistry) LinkHandler(fh registry.Handler) {
	ur.handler = &fh
}
