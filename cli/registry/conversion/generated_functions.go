// /!\ DO NOT EDIT THIS FILE /!\
//
// This file is automatically generated, do not edit by hand.
// You can edit the source file at registry/conversion/manual_functions.go and then run `go generate ./...`
// to update this file with the changes. 
//
// Generated on Tue Aug 20 2024, 13:38:59
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
//
// Example:
//
//	{{ "true" | toBool }} // Output: true
func (cr *ConversionRegistry) MustToBool(v any) (bool, error) {
	return cast.ToBoolE(v)
}

// SafeToBool is a wrapper around [MustToBool] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToBool(v any) (bool) {
  var result bool
  
  result, _ = cr.MustToBool(v)
  return result
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
func (cr *ConversionRegistry) MustToInt(v any) (int, error) {
	return cast.ToIntE(v)
}

// SafeToInt is a wrapper around [MustToInt] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToInt(v any) (int) {
  var result int
  
  result, _ = cr.MustToInt(v)
  return result
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
func (cr *ConversionRegistry) MustToInt64(v any) (int64, error) {
	return cast.ToInt64E(v)
}

// SafeToInt64 is a wrapper around [MustToInt64] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToInt64(v any) (int64) {
  var result int64
  
  result, _ = cr.MustToInt64(v)
  return result
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
func (cr *ConversionRegistry) MustToUint(v any) (uint, error) {
	return cast.ToUintE(v)
}

// SafeToUint is a wrapper around [MustToUint] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToUint(v any) (uint) {
  var result uint
  
  result, _ = cr.MustToUint(v)
  return result
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
func (cr *ConversionRegistry) MustToUint64(v any) (uint64, error) {
	return cast.ToUint64E(v)
}

// SafeToUint64 is a wrapper around [MustToUint64] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToUint64(v any) (uint64) {
  var result uint64
  
  result, _ = cr.MustToUint64(v)
  return result
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
func (cr *ConversionRegistry) MustToFloat64(v any) (float64, error) {
	return cast.ToFloat64E(v)
}

// SafeToFloat64 is a wrapper around [MustToFloat64] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToFloat64(v any) (float64) {
  var result float64
  
  result, _ = cr.MustToFloat64(v)
  return result
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
func (cr *ConversionRegistry) MustToOctal(v any) (int64, error) {
	return strconv.ParseInt(fmt.Sprint(v), 8, 64)
}

// SafeToOctal is a wrapper around [MustToOctal] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToOctal(v any) (int64) {
  var result int64
  
  result, _ = cr.MustToOctal(v)
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
func (cr *ConversionRegistry) MustToString(v any) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case error:
		return v.Error(), nil
	case fmt.Stringer:
		return v.String(), nil
	default:
		return fmt.Sprint(v), nil
	}
}

// SafeToString is a wrapper around [MustToString] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToString(v any) (string) {
  var result string
  
  result, _ = cr.MustToString(v)
  return result
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
func (cr *ConversionRegistry) MustToDate(fmt string, str string) (time.Time, error) {
	return time.ParseInLocation(fmt, str, time.Local)
}

// SafeToDate is a wrapper around [MustToDate] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToDate(fmt string, str string) (time.Time) {
  var result time.Time
  
  result, _ = cr.MustToDate(fmt, str)
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
func (cr *ConversionRegistry) MustToDuration(v any) (time.Duration, error) {
	return cast.ToDurationE(v)
}

// SafeToDuration is a wrapper around [MustToDuration] that
// returns the default value when an error is returned by the function and 
// prevents the template from stopping execution. 
func (cr *ConversionRegistry) SafeToDuration(v any) (time.Duration) {
  var result time.Duration
  
  result, _ = cr.MustToDuration(v)
  return result
}
