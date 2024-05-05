package sprout

import (
	"math"
	"reflect"

	"github.com/spf13/cast"
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

// Floor returns the largest integer less than or equal to the provided number.
//
// Parameters:
//
//	num any - the number to floor, expected to be numeric or convertible to float64.
//
// Returns:
//
//	float64 - the floored value.
//
// Example:
//
//	{{ 3.7 | floor }} // Output: 3
func (fh *FunctionHandler) Floor(num any) float64 {
	return math.Floor(cast.ToFloat64(num))
}

// Ceil returns the smallest integer greater than or equal to the provided number.
//
// Parameters:
//
//	num any - the number to ceil, expected to be numeric or convertible to float64.
//
// Returns:
//
//	float64 - the ceiled value.
//
// Example:
//
//	{{ 3.1 | ceil }} // Output: 4
func (fh *FunctionHandler) Ceil(num any) float64 {
	return math.Ceil(cast.ToFloat64(num))
}

// Round rounds a number to a specified precision and rounding threshold.
//
// Parameters:
//
//	num any - the number to round.
//	poww int - the power of ten to which to round.
//	roundOpts ...float64 - optional threshold for rounding up (default is 0.5).
//
// Returns:
//
//	float64 - the rounded number.
//
// Example:
//
//	{{ 3.746, 2, 0.5 | round }} // Output: 3.75
func (fh *FunctionHandler) Round(num any, poww int, roundOpts ...float64) float64 {
	roundOn := 0.5
	if len(roundOpts) > 0 {
		roundOn = roundOpts[0]
	}

	pow := math.Pow(10, float64(poww))
	digit := cast.ToFloat64(num) * pow
	_, div := math.Modf(digit)
	if div >= roundOn {
		return math.Ceil(digit) / pow
	}
	return math.Floor(digit) / pow
}

// Add performs addition on a slice of values.
//
// Parameters:
//
//	values ...any - numbers to add.
//
// Returns:
//
//	any - the sum of the values, converted to the type of the first value.
//
// Example:
//
//	{{ 5, 3.5, 2 | add }} // Output: 10.5
func (fh *FunctionHandler) Add(values ...any) any {
	return operateNumeric(values, func(a, b float64) float64 { return a + b }, 0.0)
}

// Add performs a unary addition operation on a single value.
//
// Parameters:
//
//	x any - the number to add.
//
// Returns:
//
//	any - the sum of the value and 1, converted to the type of the input.
//
// Example:
//
//	{{ 5 | add1 }} // Output: 6
func (fh *FunctionHandler) Add1(x any) any {
	one := reflect.ValueOf(1).Convert(reflect.TypeOf(x)).Interface()
	return fh.Add(x, one)
}

// Sub performs subtraction on a slice of values, starting with the first value.
//
// Parameters:
//
//	values ...any - numbers to subtract from the first number.
//
// Returns:
//
//	any - the result of the subtraction, converted to the type of the first value.
//
// Example:
//
//	{{ 10, 3, 2 | sub }} // Output: 5
func (fh *FunctionHandler) Sub(values ...any) any {
	return operateNumeric(values, func(a, b float64) float64 { return a - b }, 0.0)
}

// MulInt multiplies a sequence of values and returns the result as int64.
//
// Parameters:
//
//	values ...any - numbers to multiply, expected to be numeric or convertible to float64.
//
// Returns:
//
//	int64 - the product of the values.
//
// Example:
//
//	{{ 5, 3, 2 | mulInt }} // Output: 30
func (fh *FunctionHandler) MulInt(values ...any) int64 {
	return cast.ToInt64(
		operateNumeric(values, func(a, b float64) float64 { return a * b }, 1),
	)
}

// Mulf multiplies a sequence of values and returns the result as float64.
//
// Parameters:
//
//	values ...any - numbers to multiply.
//
// Returns:
//
//	any - the product of the values, converted to the type of the first value.
//
// Example:
//
//	{{ 5.5, 2.0, 2.0 | mulf }} // Output: 22.0
func (fh *FunctionHandler) Mulf(values ...any) any {
	return operateNumeric(values, func(a, b float64) float64 { return a * b }, 1.0)
}

// DivInt divides a sequence of values and returns the result as int64.
//
// Parameters:
//
//	values ...any - numbers to divide.
//
// Returns:
//
//	int64 - the quotient of the division.
//
// Example:
//
//	{{ 30, 3, 2 | divInt }} // Output: 5
func (fh *FunctionHandler) DivInt(values ...any) int64 {
	return fh.ToInt64(fh.Divf(values...))
}

// Divf divides a sequence of values, starting with the first value, and returns the result.
//
// Parameters:
//
//	values ...any - numbers to divide.
//
// Returns:
//
//	any - the quotient of the division, converted to the type of the first value.
//
// Example:
//
//	{{ 30.0, 3.0, 2.0 | divf }} // Output: 5.0
func (fh *FunctionHandler) Divf(values ...any) any {
	//FIXME:  Special manipulation to force float operation
	// This is a workaround to ensure that the result is a float to allow
	// BACKWARD COMPATIBILITY with previous versions of Sprig.
	if len(values) > 0 {
		if _, ok := values[0].(float64); !ok {
			values[0] = cast.ToFloat64(values[0])
		}
	}

	return operateNumeric(values, func(a, b float64) float64 { return a / b }, 0.0)
}

// Mod returns the remainder of division of 'x' by 'y'.
//
// Parameters:
//
//	x any, y any - numbers to divide, expected to be numeric or convertible to float64.
//
// Returns:
//
//	any - the remainder, converted to the type of 'x'.
//
// Example:
//
//	{{ 10, 4 | mod }} // Output: 2
func (fh *FunctionHandler) Mod(x, y any) any {
	result := math.Mod(cast.ToFloat64(x), cast.ToFloat64(y))

	// Convert the result to the same type as the input
	return reflect.ValueOf(result).Convert(reflect.TypeOf(x)).Interface()
}

// Min returns the minimum value among the provided arguments.
//
// Parameters:
//
//	a any - the first number to compare.
//	i ...any - additional numbers to compare.
//
// Returns:
//
//	int64 - the smallest number among the inputs.
//
// Example:
//
//	{{ 5, 3, 8, 2 | min }} // Output: 2
func (fh *FunctionHandler) Min(a any, i ...any) int64 {
	aa := fh.ToInt64(a)
	for _, b := range i {
		bb := fh.ToInt64(b)
		if bb < aa {
			aa = bb
		}
	}
	return aa
}

// Minf returns the minimum value among the provided floating-point arguments.
//
// Parameters:
//
//	a any - the first number to compare, expected to be numeric or convertible to float64.
//	i ...any - additional numbers to compare.
//
// Returns:
//
//	float64 - the smallest number among the inputs.
//
// Example:
//
//	{{ 5.2, 3.8, 8.1, 2.6 | minf }} // Output: 2.6
func (fh *FunctionHandler) Minf(a any, i ...any) float64 {
	aa := cast.ToFloat64(a)
	for _, b := range i {
		bb := cast.ToFloat64(b)
		aa = math.Min(aa, bb)
	}
	return aa
}

// Max returns the maximum value among the provided arguments.
//
// Parameters:
//
//	a any - the first number to compare.
//	i ...any - additional numbers to compare.
//
// Returns:
//
//	int64 - the largest number among the inputs.
//
// Example:
//
//	{{ 5, 3, 8, 2 | max }} // Output: 8
func (fh *FunctionHandler) Max(a any, i ...any) int64 {
	aa := fh.ToInt64(a)
	for _, b := range i {
		bb := fh.ToInt64(b)
		if bb > aa {
			aa = bb
		}
	}
	return aa
}

// Maxf returns the maximum value among the provided floating-point arguments.
//
// Parameters:
//
//	a any - the first number to compare, expected to be numeric or convertible to float64.
//	i ...any - additional numbers to compare.
//
// Returns:
//
//	float64 - the largest number among the inputs.
//
// Example:
//
//	{{ 5.2, 3.8, 8.1, 2.6 | maxf }} // Output: 8.1
func (fh *FunctionHandler) Maxf(a any, i ...any) float64 {
	aa := cast.ToFloat64(a)
	for _, b := range i {
		bb := cast.ToFloat64(b)
		aa = math.Max(aa, bb)
	}
	return aa
}
