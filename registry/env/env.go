package env

import "github.com/go-sprout/sprout"

type EnvironmentRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of env registry.
func NewRegistry() *EnvironmentRegistry {
	return &EnvironmentRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *EnvironmentRegistry) Uid() string {
	return "env"
}

// LinkHandler links the handler to the registry at runtime.
func (or *EnvironmentRegistry) LinkHandler(fh sprout.Handler) {
	or.handler = &fh
}

// RegisterFunctions registers all functions of the registry.
func (er *EnvironmentRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "env", er.Env)
	sprout.AddFunction(funcsMap, "expandEnv", er.ExpandEnv)
}
