// This package is an example of how to create a new registry for Sprout.
//
// You can use this package as a template to create your own registry and
// replace all instances of `example` with your registry name following the
// conventions. You can see more on [Documentation]
//
// [Documentation]: https://docs.atom.codes/sprout/advanced/how-to-create-a-registry
package example

import (
	"github.com/go-sprout/sprout"
)

type ExampleRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *ExampleRegistry {
	return &ExampleRegistry{}
}

// UID returns the unique identifier of the registry.
func (or *ExampleRegistry) UID() string {
	return "go-sprout/sprout.exampleofregistry" // ! Must be unique and in lowercase, replace `exampleofregistry` with your registry name and `go-sprout/sprout` with your handle name
}

// LinkHandler links the handler to the registry at runtime.
func (or *ExampleRegistry) LinkHandler(fh sprout.Handler) error {
	or.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (or *ExampleRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "example", or.ExampleFunction)
	return nil
}

func (or *ExampleRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) error {
	// Register your alias here if you have any or remove this method if you don't have any
	// You can see more on [Documentation]
	// [Documentation]: https://docs.atom.codes/sprout/features/function-aliases
	return nil
}

func (or *ExampleRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	// Register your notices here if you have any or remove this method if you don't have any
	// You can see more on [Documentation]
	// [Documentation]: https://docs.atom.codes/sprout/features/function-notices
	return nil
}
