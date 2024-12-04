package example

// ExampleFunction is a function that does something.
//
// Parameters:
//
// Returns:
//
// For an example of this function in a go template, refer to [Sprout Documentation: exampleFunction].
//
// [Sprout Documentation: exampleFunction]: https://docs.atom.codes/sprout/registries/example#examplefunction
func (or *ExampleRegistry) ExampleFunction() (string, error) {
	// Do something with helper
	or.helperFunction()
	return "Example are always better than words", nil
}
