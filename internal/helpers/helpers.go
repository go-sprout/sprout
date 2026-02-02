package helpers

import (
	"fmt"
	"reflect"
)

// strSlice converts a value to a slice of strings, handling various types
// including []string, []any, and other slice types.
//
// Parameters:
//
//	value any - the value to convert to a slice of strings.
//
// Returns:
//
//	[]string - a slice of strings representing the input value.
//
// Example:
//
//	strs := StrSlice([]any{"apple", "banana", "cherry"})
//	fmt.Println(strs) // Output: ["apple", "banana", "cherry"]
func StrSlice(value any) []string {
	if value == nil {
		return []string{}
	}

	// Handle []string type efficiently without reflection.
	if strs, ok := value.([]string); ok {
		return strs
	}

	// For slices of any, convert each element to a string, skipping nil values.
	if interfaces, ok := value.([]any); ok {
		var result []string
		for _, s := range interfaces {
			if s != nil {
				result = append(result, ToString(s))
			}
		}
		return result
	}

	// Use reflection for other slice types to convert them to []string.
	reflectedValue := reflect.ValueOf(value)
	if reflectedValue.Kind() == reflect.Slice || reflectedValue.Kind() == reflect.Array {
		var result []string
		for i := 0; i < reflectedValue.Len(); i++ {
			value := reflectedValue.Index(i).Interface()
			if value != nil {
				result = append(result, ToString(value))
			}
		}
		return result
	}

	// If it's not a slice, array, or nil, return a slice with the string representation of v.
	return []string{ToString(value)}
}

// UntilStep generates a slice of integers from 'start' to 'stop' (exclusive),
// incrementing by 'step'. If 'step' is positive, the sequence increases; if
// negative, it decreases. The function returns an empty slice if the sequence
// does not make logical sense (e.g., positive step when start is greater than
// stop or vice versa).
//
// Parameters:
//
//	start int - the starting point of the sequence.
//	stop int - the endpoint (exclusive) of the sequence.
//	step int - the increment between elements in the sequence.
//
// Returns:
//
//	[]int - a dynamically generated slice of integers based on the input
//	        parameters, or an empty slice if the parameters are inconsistent
//	        with the desired range and step.
//
// Example:
//
//	{{ 0, 10, 2 | untilStep }} // Output: [0 2 4 6 8]
//	{{ 10, 0, -2 | untilStep }} // Output: [10 8 6 4 2]
func UntilStep(start, stop, step int) []int {
	v := []int{}

	if stop < start {
		if step >= 0 {
			return v
		}
		for i := start; i > stop; i += step {
			v = append(v, i)
		}
		return v
	}

	if step <= 0 {
		return v
	}
	for i := start; i < stop; i += step {
		v = append(v, i)
	}
	return v
}

// ToString converts the input value to a string based on its type.
//
// Parameters:
//
//	given any - the value to be converted to a string.
//
// Returns:
//
//	string - the string representation of the input value.
func ToString(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Empty evaluates the emptiness of the provided value 'given'. It returns
// true if 'given' is considered empty based on its type. This method is
// essential for determining the presence or absence of meaningful value
// across various data types.
//
// Parameters:
//
//	given any - the value to be evaluated for emptiness.
//
// Returns:
//
//	bool - true if 'given' is empty, false otherwise.
//
// This method utilizes the reflect package to inspect the type and value of
// 'given'. Depending on the type, it checks for nil pointers, zero-length
// collections (arrays, slices, maps, and strings), zero values of numeric
// types (integers, floats, complex numbers, unsigned ints), and false for
// booleans.
//
// Example:
//
//	Empty(nil) // Output: true
//	Empty("") // Output: true
//	Empty(0) // Output: true
//	Empty(false) // Output: true
//	Empty(struct{}{}) // Output: false
func Empty(given any) bool {
	g := reflect.ValueOf(given)
	if !g.IsValid() {
		return true
	}

	// Basically adapted from text/template.isTrue
	switch g.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return g.Len() == 0
	case reflect.Bool:
		return !g.Bool()
	case reflect.Complex64, reflect.Complex128:
		return g.Complex() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return g.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return g.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return g.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return g.IsNil()
	default:
		return false
	}
}
