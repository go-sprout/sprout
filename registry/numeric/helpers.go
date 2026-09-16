package numeric

import (
	"errors"
	"math"
	"reflect"
	"slices"

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

// toFloatSlice converts the arguments given to a statistical function into a
// slice of float64.
//
// A single slice or array argument is expanded, so that both `sum .Values` and
// `sum 1 2 3` are accepted. Templates have no way to spread a slice into a
// variadic call, so the expanded form is the only way these functions can be
// used on data.
func toFloatSlice(values []any) ([]float64, error) {
	if len(values) == 1 && values[0] != nil {
		if rv := reflect.ValueOf(values[0]); rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
			expanded := make([]any, rv.Len())
			for i := range expanded {
				expanded[i] = rv.Index(i).Interface()
			}
			values = expanded
		}
	}

	out := make([]float64, len(values))
	for i, value := range values {
		floatValue, err := cast.ToFloat64E(value)
		if err != nil {
			return nil, sprout.NewErrConvertFailed("float64", value, err)
		}
		out[i] = floatValue
	}
	return out, nil
}

// sumFloats returns the total of the given values, cleaned of the
// representation noise that repeated float64 addition accumulates, matching
// what operateNumeric does for add.
func sumFloats(values []float64) float64 {
	var total float64
	for _, value := range values {
		total += value
	}
	return cleanFloatPrecision(total)
}

// meanFloats returns the arithmetic mean of the given values.
func meanFloats(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("cannot compute the mean of no values")
	}
	return cleanFloatPrecision(sumFloats(values) / float64(len(values))), nil
}

// medianFloats returns the middle value, or the mean of the two middle values
// when the count is even.
func medianFloats(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("cannot compute the median of no values")
	}

	sorted := make([]float64, len(values))
	copy(sorted, values)
	slices.Sort(sorted)

	middle := len(sorted) / 2
	if len(sorted)%2 == 1 {
		return sorted[middle], nil
	}
	return cleanFloatPrecision((sorted[middle-1] + sorted[middle]) / 2), nil
}

// modeFloats returns the most frequent value. Ties are resolved in favour of
// the value that appears first, so the result is stable for a given input.
func modeFloats(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("cannot compute the mode of no values")
	}

	occurrences := make(map[float64]int, len(values))
	mode, highest := values[0], 0
	for _, value := range values {
		occurrences[value]++
		if occurrences[value] > highest {
			mode, highest = value, occurrences[value]
		}
	}
	return mode, nil
}

// spreadFloats returns the distance between the largest and smallest value.
func spreadFloats(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("cannot compute the spread of no values")
	}

	lowest, highest := values[0], values[0]
	for _, value := range values[1:] {
		lowest, highest = math.Min(lowest, value), math.Max(highest, value)
	}
	return cleanFloatPrecision(highest - lowest), nil
}
