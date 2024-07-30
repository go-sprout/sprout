package time

import "github.com/go-sprout/sprout"

type TimeRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *TimeRegistry {
	return &TimeRegistry{}
}

// Uid returns the unique identifier of the registry.
func (tr *TimeRegistry) Uid() string {
	return "time"
}

// LinkHandler links the handler to the registry at runtime.
func (tr *TimeRegistry) LinkHandler(fh sprout.Handler) {
	tr.handler = &fh
}
