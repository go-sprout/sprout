package conversion

import "github.com/go-sprout/sprout"

type ConversionRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *ConversionRegistry {
	return &ConversionRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *ConversionRegistry) Uid() string {
	return "conversion"
}

// LinkHandler links the handler to the registry at runtime.
func (or *ConversionRegistry) LinkHandler(fh sprout.Handler) error {
	or.handler = fh
	return nil
}

func (cr *ConversionRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "toBool", cr.ToBool)
	sprout.AddFunction(funcsMap, "toInt", cr.ToInt)
	sprout.AddFunction(funcsMap, "toInt64", cr.ToInt64)
	sprout.AddFunction(funcsMap, "toUint", cr.ToUint)
	sprout.AddFunction(funcsMap, "toUint64", cr.ToUint64)
	sprout.AddFunction(funcsMap, "toFloat64", cr.ToFloat64)
	sprout.AddFunction(funcsMap, "toOctal", cr.ToOctal)
	sprout.AddFunction(funcsMap, "toString", cr.ToString)
	sprout.AddFunction(funcsMap, "toDate", cr.ToDate)
	sprout.AddFunction(funcsMap, "toLocalDate", cr.ToLocalDate)
	sprout.AddFunction(funcsMap, "toDuration", cr.ToDuration)
	return nil
}

func (cr *ConversionRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "toDate", "mustToDate")
	return nil
}

func (cr *ConversionRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToDate", "please use `toDate` instead"))
	return nil
}
