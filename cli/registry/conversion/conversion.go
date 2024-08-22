// Conversion Registry by sprout.
//
// The Conversion registry includes a collection of functions designed to
// convert one data type to another directly within your templates. This allows
// for seamless type transformations.
//
// Generated on Tue Aug 20 2024, 13:39:15 by sprout/cli
package conversion

import "github.com/go-sprout/sprout"

type ConversionRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewConversionRegistry creates a new instance of the ConversionRegistry registry
func NewConversionRegistry() *ConversionRegistry {
	return &ConversionRegistry{}
}

// Uid returns the unique identifier of the registry.
func (cr *ConversionRegistry) Uid() string {
	return "sprout.conversion"
}

// LinkHandler links the handler to the registry at runtime.
func (cr *ConversionRegistry) LinkHandler(fh sprout.Handler) error {
	cr.handler = fh
	return nil
}

// RegisterFunctions registers all functions in the registry.
func (cr *ConversionRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "toBool", cr.MustToBool)
	sprout.AddFunction(funcsMap, "safeToBool", cr.SafeToBool)
	return nil
}

// RegisterAliases registers all aliases in the registry.
func (cr *ConversionRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasMap, "toBool", "bool", "bobo", "toto")
	return nil
}

// RegisterNotices registers all notices in the registry.
func (cr *ConversionRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("bool", "Use the `toBool` function instead."))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("bobo", "Use the `toBool` function instead."))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("toto", "Use the `toBool` function instead."))
	return nil
}
