package example

import (
	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry"
)

type ExampleRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *ExampleRegistry {
	return &ExampleRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *ExampleRegistry) Uid() string {
	return "example" //! Must be unique and in camel case
}

// LinkHandler links the handler to the registry at runtime.
func (or *ExampleRegistry) LinkHandler(fh registry.Handler) {
	or.handler = &fh
}
