package numeric

import (
	"math"
	"reflect"

	"github.com/spf13/cast"

	"github.com/go-sprout/sprout"
)

// Floor returns the largest integer less than or equal to the provided number.
//
// Parameters:
//
//	num any - the number to floor, expected to be numeric or convertible to float64.
//
// Returns:
//
//	float64 - the floored value.
//	error - an error if the input cannot be converted to float64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: floor].
//
// [Sprout Documentation: floor]: https://docs.atom.codes/sprout/registries/numeric#floor
func (nr *NumericRegistry) Floor(num any) (float64, error) {
	float, err := cast.ToFloat64E(num)
	if err != nil {
		return 0.0, sprout.NewErrConvertFailed("float64", num, err)
	}

	return math.Floor(float), nil
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
//	error - an error if the input cannot be converted to float64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: ceil].
//
// [Sprout Documentation: ceil]: https://docs.atom.codes/sprout/registries/numeric#ceil
func (nr *NumericRegistry) Ceil(num any) (float64, error) {
	float, err := cast.ToFloat64E(num)
	if err != nil {
		return 0.0, sprout.NewErrConvertFailed("float64", num, err)
	}

	return math.Ceil(float), nil
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
// error - an error if the input cannot be converted to float64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: round].
//
// [Sprout Documentation: round]: https://docs.atom.codes/sprout/registries/numeric#round
func (nr *NumericRegistry) Round(num any, poww int, roundOpts ...float64) (float64, error) {
	// ! NEED TO CHANGE PARAMS ORDER
	roundOn := 0.5
	if len(roundOpts) > 0 {
		roundOn = roundOpts[0]
	}

	pow := math.Pow(10, float64(poww))
	float, err := cast.ToFloat64E(num)
	if err != nil {
		return 0.0, sprout.NewErrConvertFailed("float64", num, err)
	}

	digit := float * pow
	_, div := math.Modf(digit)
	if div >= roundOn {
		return math.Ceil(digit) / pow, nil
	}
	return math.Floor(digit) / pow, nil
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: add].
//
// [Sprout Documentation: add]: https://docs.atom.codes/sprout/registries/numeric#add-addf
func (nr *NumericRegistry) Add(values ...any) (any, error) {
	return operateNumeric(values, func(a, b float64) float64 { return a + b }, 0.0)
}

// Add1 performs a unary addition operation on a single value.
//
// Parameters:
//
//	x any - the number to add.
//
// Returns:
//
//	any - the sum of the value and 1, converted to the type of the input.
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: add1].
//
// [Sprout Documentation: add1]: https://docs.atom.codes/sprout/registries/numeric#add1-add1f
func (nr *NumericRegistry) Add1(x any) (any, error) {
	one := reflect.ValueOf(1).Convert(reflect.TypeOf(x)).Interface()
	return nr.Add(x, one)
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: sub].
//
// [Sprout Documentation: sub]: https://docs.atom.codes/sprout/registries/numeric#sub-subf
func (nr *NumericRegistry) Sub(values ...any) (any, error) {
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
//	error - an error if the result cannot be converted to int64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: mul].
//
// [Sprout Documentation: mul]: https://docs.atom.codes/sprout/registries/numeric#mul
func (nr *NumericRegistry) MulInt(values ...any) (int64, error) {
	result, err := operateNumeric(values, func(a, b float64) float64 { return a * b }, 1)
	if err != nil {
		return 0, err
	}

	return cast.ToInt64(result), nil
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: mulf].
//
// [Sprout Documentation: mulf]: https://docs.atom.codes/sprout/registries/numeric#mulf
func (nr *NumericRegistry) Mulf(values ...any) (any, error) {
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
// For an example of this function in a Go template, refer to [Sprout Documentation: div].
//
// [Sprout Documentation: div]: https://docs.atom.codes/sprout/registries/numeric#div
func (nr *NumericRegistry) DivInt(values ...any) (int64, error) {
	result, err := nr.Divf(values...)
	if err != nil {
		return 0, err
	}

	return cast.ToInt64(result), nil
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
// For an example of this function in a Go template, refer to [Sprout Documentation: divf].
//
// [Sprout Documentation: divf]: https://docs.atom.codes/sprout/registries/numeric#divf
func (nr *NumericRegistry) Divf(values ...any) (any, error) {
	// FIXME:  Special manipulation to force float operation
	// This is a workaround to ensure that the result is a float to allow
	// BACKWARDS COMPATIBILITY with previous versions of Sprig.
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: mod].
//
// [Sprout Documentation: mod]: https://docs.atom.codes/sprout/registries/numeric#mod
func (nr *NumericRegistry) Mod(x, y any) (any, error) {
	floatX, err := cast.ToFloat64E(x)
	if err != nil {
		return 0, sprout.NewErrConvertFailed("float64", x, err)
	}

	floatY, err := cast.ToFloat64E(y)
	if err != nil {
		return 0, sprout.NewErrConvertFailed("float64", y, err)
	}

	result := math.Mod(floatX, floatY)

	// Convert the result to the same type as the input
	return reflect.ValueOf(result).Convert(reflect.TypeOf(x)).Interface(), nil
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: min].
//
// [Sprout Documentation: min]: https://docs.atom.codes/sprout/registries/numeric#min
func (nr *NumericRegistry) Min(a any, i ...any) (int64, error) {
	intA, err := cast.ToInt64E(a)
	if err != nil {
		return 0, sprout.NewErrConvertFailed("int64", a, err)
	}

	for _, b := range i {
		intB, err := cast.ToInt64E(b)
		if err != nil {
			return 0, sprout.NewErrConvertFailed("int64", b, err)
		}
		if intB < intA {
			intA = intB
		}
	}
	return intA, nil
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: minf].
//
// [Sprout Documentation: minf]: https://docs.atom.codes/sprout/registries/numeric#minf
func (nr *NumericRegistry) Minf(a any, i ...any) (float64, error) {
	floatA, err := cast.ToFloat64E(a)
	if err != nil {
		return 0, sprout.NewErrConvertFailed("float64", a, err)
	}
	for _, b := range i {
		floatB, err := cast.ToFloat64E(b)
		if err != nil {
			return 0, sprout.NewErrConvertFailed("float64", b, err)
		}
		floatA = math.Min(floatA, floatB)
	}
	return floatA, nil
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: max].
//
// [Sprout Documentation: max]: https://docs.atom.codes/sprout/registries/numeric#max
func (nr *NumericRegistry) Max(a any, i ...any) (int64, error) {
	intA, err := cast.ToInt64E(a)
	if err != nil {
		return 0, sprout.NewErrConvertFailed("int64", a, err)
	}

	for _, b := range i {
		intB, err := cast.ToInt64E(b)
		if err != nil {
			return 0, sprout.NewErrConvertFailed("int64", b, err)
		}

		if intB > intA {
			intA = intB
		}
	}
	return intA, nil
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
//	error - when a value cannot be converted.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: maxf].
//
// [Sprout Documentation: maxf]: https://docs.atom.codes/sprout/registries/numeric#maxf
func (nr *NumericRegistry) Maxf(a any, i ...any) (float64, error) {
	floatA, err := cast.ToFloat64E(a)
	if err != nil {
		return 0, sprout.NewErrConvertFailed("float64", a, err)
	}
	for _, b := range i {
		floatB, err := cast.ToFloat64E(b)
		if err != nil {
			return 0, sprout.NewErrConvertFailed("float64", b, err)
		}
		floatA = math.Max(floatA, floatB)
	}
	return floatA, nil
}
