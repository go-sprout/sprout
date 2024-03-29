/*
package sprout provides template functions for Go.

This package contains a number of utility functions for working with data
inside of Go `html/template` and `text/template` files.

To add these functions, use the `template.Funcs()` method:

	t := templates.New("foo").Funcs(sprout.FuncMap())

Note that you should add the function map before you parse any template files.

	In several cases, sprout reverses the order of arguments from the way they
	appear in the standard library. This is to make it easier to pipe
	arguments into functions.

See http://masterminds.github.io/sprout/ for more detailed documentation on each of the available functions.
*/
package sprout
