package regexp

import "github.com/go-sprout/sprout"

type RegexpRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of regexp registry.
func NewRegistry() *RegexpRegistry {
	return &RegexpRegistry{}
}

// Uid returns the unique identifier of the registry.
func (rr *RegexpRegistry) Uid() string {
	return "regexp"
}

// LinkHandler links the handler to the registry at runtime.
func (rr *RegexpRegistry) LinkHandler(fh sprout.Handler) {
	rr.handler = &fh
}
