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
func (br *BuiltinRegistry) LinkHandler(fh sprout.Handler) error {
	br.handler = &fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (br *BuiltinRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "hello", br.Hello)
	sprout.AddFunction(funcsMap, "default", br.Default)
	sprout.AddFunction(funcsMap, "empty", br.Empty)
	sprout.AddFunction(funcsMap, "all", br.All)
	sprout.AddFunction(funcsMap, "any", br.Any)
	sprout.AddFunction(funcsMap, "coalesce", br.Coalesce)
	sprout.AddFunction(funcsMap, "ternary", br.Ternary)
	sprout.AddFunction(funcsMap, "cat", br.Cat)
	return nil
}
