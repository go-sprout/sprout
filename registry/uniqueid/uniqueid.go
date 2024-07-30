package uniqueid

import "github.com/go-sprout/sprout"

type UniqueIDRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
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
func (ur *UniqueIDRegistry) LinkHandler(fh sprout.Handler) {
	ur.handler = &fh
}

// RegisterFunctions registers all functions of the registry.
func (ur *UniqueIDRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "uuidv4", ur.Uuidv4)
}
