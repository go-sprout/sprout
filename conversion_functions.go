package sprout

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

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
func (fh *FunctionHandler) ToInt(v any) int {
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
func (fh *FunctionHandler) ToInt64(v any) int64 {
	return cast.ToInt64(v)
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
func (fh *FunctionHandler) ToFloat64(v any) float64 {
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
func (fh *FunctionHandler) ToOctal(v any) int64 {
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
func (fh *FunctionHandler) ToString(v any) string {
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
func (fh *FunctionHandler) ToDate(fmt, str string) time.Time {
	result, _ := fh.MustToDate(fmt, str)
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
func (fh *FunctionHandler) ToDuration(v any) time.Duration {
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
func (fh *FunctionHandler) MustToDate(fmt, str string) (time.Time, error) {
	return time.ParseInLocation(fmt, str, time.Local)
}
