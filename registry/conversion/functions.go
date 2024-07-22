package conversion

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-sprout/sprout/registry"
	"github.com/spf13/cast"
)

func (cr *ConversionRegistry) RegisterFunctions(funcsMap registry.FunctionMap) {
	registry.AddFunction(funcsMap, "toBool", cr.ToBool)
	registry.AddFunction(funcsMap, "toInt", cr.ToInt)
	registry.AddFunction(funcsMap, "toInt64", cr.ToInt64)
	registry.AddFunction(funcsMap, "toUint", cr.ToUint)
	registry.AddFunction(funcsMap, "toUint64", cr.ToUint64)
	registry.AddFunction(funcsMap, "toFloat64", cr.ToFloat64)
	registry.AddFunction(funcsMap, "toOctal", cr.ToOctal)
	registry.AddFunction(funcsMap, "toString", cr.ToString)
	registry.AddFunction(funcsMap, "toDate", cr.ToDate)
	registry.AddFunction(funcsMap, "toDuration", cr.ToDuration)
	registry.AddFunction(funcsMap, "mustToDate", cr.MustToDate)
}

// ToBool converts a value to a boolean.
//
// Parameters:
//
//	v any - the value to convert to a boolean. This can be any types reasonably be converted to true or false.
//
// Returns:
//
//	bool - the boolean representation of the value.
//
// Example:
//
//	{{ "true" | toBool }} // Output: true
func (cr *ConversionRegistry) ToBool(v any) bool {
	return cast.ToBool(v)
}

// ToInt converts a value to an int using robust type casting.
//
// Parameters:
//
//	v any - the value to convert to an int.
//
// Returns:
//
//	int - the integer representation of the value.
//
// Example:
//
//	{{ "123" | toInt }} // Output: 123
func (cr *ConversionRegistry) ToInt(v any) int {
	return cast.ToInt(v)
}

// ToInt64 converts a value to an int64, accommodating larger integer values.
//
// Parameters:
//
//	v any - the value to convert to an int64.
//
// Returns:
//
//	int64 - the int64 representation of the value.
//
// Example:
//
//	{{ "123456789012" | toInt64 }} // Output: 123456789012
func (cr *ConversionRegistry) ToInt64(v any) int64 {
	return cast.ToInt64(v)
}

// ToUint converts a value to a uint.
//
// Parameters:
//
//	v any - the value to convert to uint. This value can be of any type that is numerically convertible.
//
// Returns:
//
//	uint - the uint representation of the value.
//
// Example:
//
//	{{ "123" | toUint }} // Output: 123
func (cr *ConversionRegistry) ToUint(v any) uint {
	return cast.ToUint(v)
}

// ToUint64 converts a value to a uint64.
//
// Parameters:
//
//	v any - the value to convert to uint64. This value can be of any type that is numerically convertible.
//
// Returns:
//
//	uint64 - the uint64 representation of the value.
//
// Example:
//
//	{{ "123456789012345" | toUint64 }} // Output: 123456789012345
func (cr *ConversionRegistry) ToUint64(v any) uint64 {
	return cast.ToUint64(v)
}

// ToFloat64 converts a value to a float64.
//
// Parameters:
//
//	v any - the value to convert to a float64.
//
// Returns:
//
//	float64 - the float64 representation of the value.
//
// Example:
//
//	{{ "123.456" | toFloat64 }} // Output: 123.456
func (cr *ConversionRegistry) ToFloat64(v any) float64 {
	return cast.ToFloat64(v)
}

// ToOctal parses a string value as an octal (base 8) integer.
//
// Parameters:
//
//	v any - the string representing an octal number.
//
// Returns:
//
//	int64 - the decimal (base 10) representation of the octal value.
//	If parsing fails, returns 0.
//
// Example:
//
//	{{ "123" | toOctal }} // Output: 83 (since "123" in octal is 83 in decimal)
func (cr *ConversionRegistry) ToOctal(v any) int64 {
	result, err := strconv.ParseInt(fmt.Sprint(v), 8, 64)
	if err != nil {
		return 0
	}
	return result
}

// ToString converts a value to a string, handling various types effectively.
//
// Parameters:
//
//	v any - the value to convert to a string.
//
// Returns:
//
//	string - the string representation of the value.
//
// Example:
//
//	{{ 123 | toString }} // Output: "123"
func (cr *ConversionRegistry) ToString(v any) string {
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

// ToDate converts a string to a time.Time object based on a format specification.
//
// Parameters:
//
//	fmt string - the date format string.
//	str string - the date string to parse.
//
// Returns:
//
//	time.Time - the parsed date.
//
// Example:
//
//	{{ "2006-01-02", "2023-05-04" | toDate }} // Output: 2023-05-04 00:00:00 +0000 UTC
func (cr *ConversionRegistry) ToDate(fmt, str string) time.Time {
	result, _ := cr.MustToDate(fmt, str)
	return result
}

// ToDuration converts a value to a time.Duration.
//
// Parameters:
//
//	v any - the value to convert to time.Duration. This value can be a string, int, or another compatible type.
//
// Returns:
//
//	time.Duration - the duration representation of the value.
//
// Example:
//
//	{{ (toDuration "1h30m").Seconds }} // Output: 5400
func (cr *ConversionRegistry) ToDuration(v any) time.Duration {
	return cast.ToDuration(v)
}

// MustToDate tries to parse a string into a time.Time object based on a format,
// returning an error if parsing fails.
//
// Parameters:
//
//	fmt string - the date format string.
//	str string - the date string to parse.
//
// Returns:
//
//	time.Time - the parsed date.
//	error - error if the date string does not conform to the format.
//
// Example:
//
//	{{ "2006-01-02", "2023-05-04" | mustToDate }} // Output: 2023-05-04 00:00:00 +0000 UTC, nil
func (cr *ConversionRegistry) MustToDate(fmt, str string) (time.Time, error) {
	return time.ParseInLocation(fmt, str, time.Local)
}
