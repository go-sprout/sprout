package std

import "github.com/go-sprout/sprout"

type StdRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *StdRegistry {
	return &StdRegistry{}
}

// UID returns the unique identifier of the registry.
func (sr *StdRegistry) UID() string {
	return "go-sprout/sprout.std"
}

// LinkHandler links the handler to the registry at runtime.
func (sr *StdRegistry) LinkHandler(fh sprout.Handler) error {
	sr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (sr *StdRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "hello", sr.Hello)
	sprout.AddFunction(funcsMap, "default", sr.Default)
	sprout.AddFunction(funcsMap, "empty", sr.Empty)
	sprout.AddFunction(funcsMap, "all", sr.All)
	sprout.AddFunction(funcsMap, "any", sr.Any)
	sprout.AddFunction(funcsMap, "coalesce", sr.Coalesce)
	sprout.AddFunction(funcsMap, "ternary", sr.Ternary)
	sprout.AddFunction(funcsMap, "cat", sr.Cat)
	return nil
}
