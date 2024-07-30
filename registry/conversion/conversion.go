package conversion

import "github.com/go-sprout/sprout"

type ConversionRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *ConversionRegistry {
	return &ConversionRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *ConversionRegistry) Uid() string {
	return "conversion"
}

// LinkHandler links the handler to the registry at runtime.
func (or *ConversionRegistry) LinkHandler(fh sprout.Handler) {
	or.handler = &fh
}
