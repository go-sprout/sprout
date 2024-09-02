package conversion

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

// ToBool converts a value to a boolean.
//
// Parameters:
//
//	v any - the value to convert to a boolean. This can be any types reasonably be converted to true or false.
//
// Returns:
//
//	bool - the boolean representation of the value.
//	error - error if the value cannot be converted to a boolean.
//
// Example:
//
//	{{ "true" | toBool }} // Output: true
func (cr *ConversionRegistry) ToBool(v any) (bool, error) {
	return cast.ToBoolE(v)
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
//	error - error if the value cannot be converted to an int.
//
// Example:
//
//	{{ "123" | toInt }} // Output: 123
func (cr *ConversionRegistry) ToInt(v any) (int, error) {
	return cast.ToIntE(v)
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
//	error - error if the value cannot be converted to an int64.
//
// Example:
//
//	{{ "123456789012" | toInt64 }} // Output: 123456789012
func (cr *ConversionRegistry) ToInt64(v any) (int64, error) {
	return cast.ToInt64E(v)
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
//	error - error if the value cannot be converted to a uint.
//
// Example:
//
//	{{ "123" | toUint }} // Output: 123
func (cr *ConversionRegistry) ToUint(v any) (uint, error) {
	return cast.ToUintE(v)
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
//	error - error if the value cannot be converted to a uint64.
//
// Example:
//
//	{{ "123456789012345" | toUint64 }} // Output: 123456789012345
func (cr *ConversionRegistry) ToUint64(v any) (uint64, error) {
	return cast.ToUint64E(v)
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
//	error - error if the value cannot be converted to a float64.
//
// Example:
//
//	{{ "123.456" | toFloat64 }} // Output: 123.456
func (cr *ConversionRegistry) ToFloat64(v any) (float64, error) {
	return cast.ToFloat64E(v)
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
//	error - error if the value cannot be converted to an octal number.
//
// Example:
//
//	{{ "123" | toOctal }} // Output: 83 (since "123" in octal is 83 in decimal)
func (cr *ConversionRegistry) ToOctal(v any) (int64, error) {
	result, err := strconv.ParseInt(fmt.Sprint(v), 8, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse octal: %w", err)
	}
	return result, nil
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
		return fmt.Sprint(v)
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
//	error - error if the date string does not conform to the format.
//
// Example:
//
//	{{ "2006-01-02", "2023-05-04" | toDate }} // Output: 2023-05-04 00:00:00 +0000 UTC
func (cr *ConversionRegistry) ToDate(fmt, str string) (time.Time, error) {
	return time.ParseInLocation(fmt, str, time.Local)
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
//	error - error if the value cannot be converted to a duration.
//
// Example:
//
//	{{ (toDuration "1h30m").Seconds }} // Output: 5400
func (cr *ConversionRegistry) ToDuration(v any) (time.Duration, error) {
	return cast.ToDurationE(v)
}
