package example

import (
	"github.com/go-sprout/sprout"
)

// RegisterFunctions registers all functions of the registry.
func (or *ExampleRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "example", or.ExampleFunction)
	return nil
}

// ExampleFunction is a function that does something.
func (or *ExampleRegistry) ExampleFunction() (string, error) {
	// Do something with helper
	or.helperFunction()
	return "Example are always better than words", nil
}
