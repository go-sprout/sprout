package numeric

import (
	"github.com/go-sprout/sprout/registry"
)

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
	handler *registry.Handler // Embedding Handler for shared functionality
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
func (nr *NumericRegistry) LinkHandler(fh registry.Handler) {
	nr.handler = &fh
}
