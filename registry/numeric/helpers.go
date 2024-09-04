package numeric

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/go-sprout/sprout"
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
func operateNumeric(values []any, op numericOperation, initial any) (any, error) {
	if len(values) == 0 {
		return initial, nil
	}

	result, err := cast.ToFloat64E(values[0])
	if err != nil {
		return 0.0, sprout.NewErrConvertFailed("float64", values[0], err)
	}
	for _, value := range values[1:] {
		floatValue, err := cast.ToFloat64E(value)
		if err != nil {
			return 0.0, sprout.NewErrConvertFailed("float64", value, err)
		}
		result = op(result, floatValue)
	}

	// Direct type assertion for common types to avoid reflection overhead
	initialType := reflect.TypeOf(values[0])
	switch initialType.Kind() {
	case reflect.Int:
		return int(result), nil
	case reflect.Float64:
		return result, nil
	default:
		return reflect.ValueOf(result).Convert(initialType).Interface(), nil
	}
}
