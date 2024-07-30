package numeric

import "github.com/go-sprout/sprout"

// numericOperation defines a function type that performs a binary operation on
// two float64 values. It is used to abstract arithmetic operations like
// addition, subtraction, multiplication, or division so that these can be
// applied in a generic function that processes lists of numbers.
//
// Example Usage:
//
//	add := func(a, b float64) float64 { return a + b }
//	result := operateNumeric([]any{1.0, 2.0}, add, 0.0)
//	fmt.Println(result)  // Output: 3.0
type numericOperation func(float64, float64) float64

type NumericRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of numeric registry.
func NewRegistry() *NumericRegistry {
	return &NumericRegistry{}
}

// Uid returns the unique identifier of the registry.
func (nr *NumericRegistry) Uid() string {
	return "numeric"
}

// LinkHandler links the handler to the registry at runtime.
func (nr *NumericRegistry) LinkHandler(fh sprout.Handler) {
	nr.handler = &fh
}

// RegisterFunctions registers all functions of the registry.
func (nr *NumericRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "floor", nr.Floor)
	sprout.AddFunction(funcsMap, "ceil", nr.Ceil)
	sprout.AddFunction(funcsMap, "round", nr.Round)
	sprout.AddFunction(funcsMap, "add", nr.Add)
	sprout.AddFunction(funcsMap, "add1", nr.Add1)
	sprout.AddFunction(funcsMap, "sub", nr.Sub)
	sprout.AddFunction(funcsMap, "mul", nr.MulInt)
	sprout.AddFunction(funcsMap, "mulf", nr.Mulf)
	sprout.AddFunction(funcsMap, "div", nr.DivInt)
	sprout.AddFunction(funcsMap, "divf", nr.Divf)
	sprout.AddFunction(funcsMap, "mod", nr.Mod)
	sprout.AddFunction(funcsMap, "min", nr.Min)
	sprout.AddFunction(funcsMap, "minf", nr.Minf)
	sprout.AddFunction(funcsMap, "max", nr.Max)
	sprout.AddFunction(funcsMap, "maxf", nr.Maxf)
}

func (nr *NumericRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) {
	sprout.AddAlias(aliasMap, "add", "addf")
	sprout.AddAlias(aliasMap, "add1", "add1f")
	sprout.AddAlias(aliasMap, "sub", "subf")
}
