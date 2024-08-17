package std

import (
	"fmt"
	"strings"

	"github.com/go-sprout/sprout/internal/helpers"
)

// Hello returns a greeting string.
// It simply returns the string "Hello!" to be used as a test function.
func (sr *StdRegistry) Hello() string {
	return "Hello!"
}

// Default returns the first non-empty value from the given arguments or a
// default value if the argument list is empty or the first element is empty.
// It accepts a default value `defaultValue` of any type and a variadic slice
// `given` of any type. If `given` is not provided or the first element in
// `given` is empty, it returns `defaultValue`.
// Otherwise, it returns the first element of `given`.
// If you want to catch the first non-empty value from a list of values, use
// the `Coalesce` function instead.
//
// Parameters:
//
//	defaultValue any - the default value to return if no valid argument is
//	                   provided or if the first argument is empty.
//	given ...any     - a variadic slice of any type to check the first
//	                   element of it for emptiness.
//
// Returns:
//
//	any - the first element of `given`, or `defaultValue` if `given` is empty
//	      or all values are empty.
//
// Example:
//
//	{{ nil | default "default" }} // Output: "default"
//	{{ "" | default "default" }}  // Output: "default"
//	{{ "first" | default "default" }} // Output: "first"
//	{{ "first" | default "default" "second" }} // Output: "second"
func (sr *StdRegistry) Default(defaultValue any, given ...any) any {
	if len(given) == 0 || helpers.Empty(given[0]) {
		return defaultValue
	}
	return given[0]
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
//	{{ nil | empty }} // Output: true
//	{{ "" | empty }} // Output: true
//	{{ 0 | empty }} // Output: true
//	{{ false | empty }} // Output: true
//	{{ struct{}{} | empty }} // Output: false
func (sr *StdRegistry) Empty(given any) bool {
	return helpers.Empty(given)
}

// All checks if all values in the provided variadic slice are non-empty.
// It returns true only if none of the values are considered empty by the Empty method.
//
// Parameters:
//
//	values ...any - a variadic parameter list of values to be checked.
//
// Returns:
//
//	bool - true if all values are non-empty, false otherwise.
//
// Example:
//
//	{{ 1, "hello", true | all }} // Output: true
//	{{ 1, "", true | all }} // Output: false
func (sr *StdRegistry) All(values ...any) bool {
	for _, val := range values {
		if helpers.Empty(val) {
			return false
		}
	}
	return true
}

// Any checks if any of the provided values are non-empty.
// It returns true if at least one value is non-empty.
//
// Parameters:
//   values ...any - a variadic parameter list of values to be checked.
//
// Returns:
//   bool - true if any value is non-empty, false if all are empty.
//
// Example:
//   {{ "", 0, false | any }} // Output: false
//   {{ "", 0, "text" | any }} // Output: true

func (sr *StdRegistry) Any(values ...any) bool {
	for _, val := range values {
		if !helpers.Empty(val) {
			return true
		}
	}
	return false
}

// Coalesce returns the first non-empty value from the given list.
// If all values are empty, it returns nil.
//
// Parameters:
//   values ...any - a variadic parameter list of values from which the first
//                   non-empty value should be selected.
//
// Returns:
//   any - the first non-empty value, or nil if all values are empty.
//
// Example:
//   {{ nil, "", "first", "second" | coalesce }} // Output: "first"

func (sr *StdRegistry) Coalesce(values ...any) any {
	for _, val := range values {
		if !helpers.Empty(val) {
			return val
		}
	}
	return nil
}

// Ternary mimics the ternary conditional operator found in many programming languages.
// It returns 'trueValue' if 'condition' is true, otherwise 'falseValue'.
//
// Parameters:
//
//	trueValue any - the value to return if 'condition' is true.
//	falseValue any - the value to return if 'condition' is false.
//	condition bool - the condition to evaluate.
//
// Returns:
//
//	any - the result based on the evaluated condition.
//
// Example:
//
//	{{ "yes", "no", true | ternary }} // Output: "yes"
//	{{ "yes", "no", false | ternary }} // Output: "no"
func (sr *StdRegistry) Ternary(trueValue any, falseValue any, condition bool) any {
	if condition {
		return trueValue
	}

	return falseValue
}

// Cat concatenates a series of values into a single string. Each value is
// converted to its string representation and separated by a space. Nil
// values are skipped, and no trailing spaces are added.
//
// Parameters:
//
//	values ...any - a variadic parameter list of values to be concatenated.
//
// Returns:
//
//	string - a single string composed of all non-nil input values separated
//	         by spaces.
//
// Example:
//
//	{{ "Hello", nil, 123, true | cat }} // Output: "Hello 123 true"
func (sr *StdRegistry) Cat(values ...any) string {
	var builder strings.Builder
	for i, item := range values {
		if item == nil {
			continue // Skip nil elements
		}
		if i > 0 {
			builder.WriteRune(' ') // Add space between elements
		}
		// Append the string representation of the item
		builder.WriteString(fmt.Sprint(item))
	}
	// Return the concatenated string without trailing spaces
	return builder.String()
}
