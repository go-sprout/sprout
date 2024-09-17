package network

import (
	"github.com/go-sprout/sprout"
)

type NetworkRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *NetworkRegistry {
	return &NetworkRegistry{}
}

// Uid returns the unique identifier of the registry.
func (nr *NetworkRegistry) Uid() string {
	return "network" // ! Must be unique and in camel case
}

// LinkHandler links the handler to the registry at runtime.
func (nr *NetworkRegistry) LinkHandler(fh sprout.Handler) error {
	nr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (nr *NetworkRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "example", nr.ExampleFunction)
	return nil
}

func (nr *NetworkRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) error {
	// Register your alias here if you have any or remove this method
	return nil
}

func (nr *NetworkRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	// Register your notices here if you have any or remove this method
	return nil
}
