package numeric

import (
	"reflect"

	"github.com/spf13/cast"
)

// operateNumeric applies a numericOperation to a slice of any type, converting
// to and from float64. The result is converted back to the type of the first
// element in the slice.
//
// Parameters:
//
//	values []any - Slice of numeric values.
//	op numericOperation - Function to apply.
//	initial float64 - Starting value for the operation.
//
// Returns:
//
//	any - Result of the operation, converted to the type of the first slice element.
//
// Example:
//
//	add := func(a, b float64) float64 { return a + b }
//	result := operateNumeric([]any{1.5, 2.5}, add, 0)
//	fmt.Println(result) // Output: 4.0 (type float64 if first element is float64)
func operateNumeric(values []any, op numericOperation, initial any) any {
	if len(values) == 0 {
		return initial
	}

	result := cast.ToFloat64(values[0])
	for _, value := range values[1:] {
		result = op(result, cast.ToFloat64(value))
	}

	initialType := reflect.TypeOf(values[0])
	return reflect.ValueOf(result).Convert(initialType).Interface()
}
