package sprout

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/mitchellh/copystructure"
)

// Hello returns a greeting string.
// It simply returns the string "Hello!" to be used as a test function.
func (fh *FunctionHandler) Hello() string {
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
func (fh *FunctionHandler) Default(defaultValue any, given ...any) any {
	if fh.Empty(given) || fh.Empty(given[0]) {
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
func (fh *FunctionHandler) Empty(given any) bool {
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
	case reflect.Struct:
		return false
	default:
		return g.IsNil()
	}
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
func (fh *FunctionHandler) All(values ...any) bool {
	for _, val := range values {
		if fh.Empty(val) {
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

func (fh *FunctionHandler) Any(values ...any) bool {
	for _, val := range values {
		if !fh.Empty(val) {
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

func (fh *FunctionHandler) Coalesce(values ...any) any {
	for _, val := range values {
		if !fh.Empty(val) {
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
func (fh *FunctionHandler) Ternary(trueValue any, falseValue any, condition bool) any {
	if condition {
		return trueValue
	}

	return falseValue
}

// Uuidv4 generates a new random UUID (Universally Unique Identifier) version 4.
// This function does not take parameters and returns a string representation
// of a UUID.
//
// Returns:
//
//	string - a new UUID string.
//
// Example:
//
//	{{ uuidv4 }} // Output: "3f0c463e-53f5-4f05-a2ec-3c083aa8f937"
func (fh *FunctionHandler) Uuidv4() string {
	return uuid.New().String()
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
func (fh *FunctionHandler) Cat(values ...any) string {
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

// Until generates a slice of integers from 0 up to but not including 'count'.
// If 'count' is negative, it produces a descending slice from 0 down to 'count',
// inclusive, with a step of -1. The function leverages UntilStep to specify
// the range and step dynamically.
//
// Parameters:
//   count int - the endpoint (exclusive) of the range to generate.
//
// Returns:
//   []int - a slice of integers from 0 to 'count' with the appropriate step
//           depending on whether 'count' is positive or negative.
//
// Example:
//   {{ 5 | until }} // Output: [0 1 2 3 4]
//   {{ -3 | until }} // Output: [0 -1 -2]

func (fh *FunctionHandler) Until(count int) []int {
	step := 1
	if count < 0 {
		step = -1
	}
	return fh.UntilStep(0, count, step)
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
func (fh *FunctionHandler) UntilStep(start, stop, step int) []int {
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

// TypeIs compares the type of 'src' to a target type string 'target'.
// It returns true if the type of 'src' matches the 'target'.
//
// Parameters:
//
//	target string - the string representation of the type to check against.
//	src any - the variable whose type is being checked.
//
// Returns:
//
//	bool - true if 'src' is of type 'target', false otherwise.
//
// Example:
//
//	{{ "int", 42 | typeIs }} // Output: true
func (fh *FunctionHandler) TypeIs(target string, src any) bool {
	return target == fh.TypeOf(src)
}

// TypeIsLike compares the type of 'src' to a target type string 'target',
// including a wildcard '*' prefix option. It returns true if 'src' matches
// 'target' or '*target'. Useful for checking if a variable is of a specific
// type or a pointer to that type.
//
// Parameters:
//
//	target string - the string representation of the type or its wildcard version.
//	src any - the variable whose type is being checked.
//
// Returns:
//
//	bool - true if the type of 'src' matches 'target' or '*'+target, false otherwise.
//
// Example:
//
//	{{ "*int", 42 | typeIsLike }} // Output: true
func (fh *FunctionHandler) TypeIsLike(target string, src any) bool {
	t := fh.TypeOf(src)
	return target == t || "*"+target == t
}

// TypeOf returns the type of 'src' as a string.
//
// Parameters:
//
//	src any - the variable whose type is being determined.
//
// Returns:
//
//	string - the string representation of 'src's type.
//
// Example:
//
//	{{ 42 | typeOf }} // Output: "int"
func (fh *FunctionHandler) TypeOf(src any) string {
	return fmt.Sprintf("%T", src)
}

// KindIs compares the kind of 'src' to a target kind string 'target'.
// It returns true if the kind of 'src' matches the 'target'.
//
// Parameters:
//
//	target string - the string representation of the kind to check against.
//	src any - the variable whose kind is being checked.
//
// Returns:
//
//	bool - true if 'src's kind is 'target', false otherwise.
//
// Example:
//
//	{{ "int", 42 | kindIs }} // Output: true
func (fh *FunctionHandler) KindIs(target string, src any) bool {
	return target == fh.KindOf(src)
}

// KindOf returns the kind of 'src' as a string.
//
// Parameters:
//
//	src any - the variable whose kind is being determined.
//
// Returns:
//
//	string - the string representation of 'src's kind.
//
// Example:
//
//	{{ 42 | kindOf }} // Output: "int"
func (fh *FunctionHandler) KindOf(src any) string {
	return reflect.ValueOf(src).Kind().String()
}

// DeepEqual determines if two variables, 'x' and 'y', are deeply equal.
// It uses reflect.DeepEqual to evaluate equality.
//
// Parameters:
//
//	x, y any - the variables to be compared.
//
// Returns:
//
//	bool - true if 'x' and 'y' are deeply equal, false otherwise.
//
// Example:
//
//	{{ {"a":1}, {"a":1} | deepEqual }} // Output: true
func (fh *FunctionHandler) DeepEqual(x, y any) bool {
	return reflect.DeepEqual(y, x)
}

// DeepCopy performs a deep copy of 'element' and panics if copying fails.
// It relies on MustDeepCopy to perform the copy and handle errors internally.
//
// Parameters:
//
//	element any - the element to be deeply copied.
//
// Returns:
//
//	any - a deep copy of 'element'.
//
// Example:
//
//	{{ {"name":"John"} | deepCopy }} // Output: {"name":"John"}
func (fh *FunctionHandler) DeepCopy(element any) any {
	c, err := fh.MustDeepCopy(element)
	if err != nil {
		return nil
	}

	return c
}

func (fh *FunctionHandler) MustDeepCopy(element any) (any, error) {
	if element == nil {
		return nil, errors.New("element cannot be nil")
	}
	return copystructure.Copy(element)
}

// ! DEPRECATED: This should be removed in the next major version.
//
// Fail creates an error with a specified message and returns a nil pointer
// alongside the created error. This function is typically used to indicate
// failure conditions in functions that return a pointer and an error.
//
// Parameters:
//
//	message string - the error message to be associated with the returned error.
//
// Returns:
//
//	*uint - always returns nil, indicating no value is associated with the failure.
//	error - the error object containing the provided message.
//
// Example:
//
//	{{ "Operation failed" | fail }} // Output: nil, error with "Operation failed"
func (fh *FunctionHandler) Fail(message string) (*uint, error) {
	return nil, errors.New(message)
}
