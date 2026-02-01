package numeric

import (
	"math"
	"reflect"

	"github.com/spf13/cast"

	"github.com/go-sprout/sprout"
)

// cleanFloatPrecision rounds a float64 to 15 significant decimal digits
// to eliminate floating-point representation noise.
// This fixes issues like 0.1+0.2=0.30000000000000004 becoming exactly 0.3.
func cleanFloatPrecision(v float64) float64 {
	if v == 0 || math.IsNaN(v) || math.IsInf(v, 0) {
		return v
	}

	// Determine magnitude to scale value for rounding
	magnitude := math.Floor(math.Log10(math.Abs(v)))

	// Scale to have significant digits as integers, round, then scale back
	// 15 significant figures is sufficient for float64 while cleaning noise
	const sigFigs = 15
	scale := math.Pow(10, float64(sigFigs)-1-magnitude)

	return math.Round(v*scale) / scale
}

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

	// Clean floating-point precision noise from the result
	result = cleanFloatPrecision(result)

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
