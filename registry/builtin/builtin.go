package builtin

import "github.com/go-sprout/sprout"

type BuiltinRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *BuiltinRegistry {
	return &BuiltinRegistry{}
}

// Uid returns the unique identifier of the registry.
func (br *BuiltinRegistry) Uid() string {
	return "builtin"
}

// LinkHandler links the handler to the registry at runtime.
func (br *BuiltinRegistry) LinkHandler(fh sprout.Handler) {
	br.handler = &fh
}
