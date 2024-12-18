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
//	value any - the value to convert to a boolean. This can be any types reasonably be converted to true or false.
//
// Returns:
//
//	bool - the boolean representation of the value.
//	error - error if the value cannot be converted to a boolean.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toBool].
//
// [Sprout Documentation: toBool]: https://docs.atom.codes/sprout/registries/conversion#tobool
func (cr *ConversionRegistry) ToBool(value any) (bool, error) {
	return cast.ToBoolE(value)
}

// ToInt converts a value to an int using robust type casting.
//
// Parameters:
//
//	value any - the value to convert to an int.
//
// Returns:
//
//	int - the integer representation of the value.
//	error - error if the value cannot be converted to an int.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toInt].
//
// [Sprout Documentation: toInt]: https://docs.atom.codes/sprout/registries/conversion#toint
func (cr *ConversionRegistry) ToInt(value any) (int, error) {
	return cast.ToIntE(value)
}

// ToInt64 converts a value to an int64, accommodating larger integer values.
//
// Parameters:
//
//	value any - the value to convert to an int64.
//
// Returns:
//
//	int64 - the int64 representation of the value.
//	error - error if the value cannot be converted to an int64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toInt64].
//
// [Sprout Documentation: toInt64]: https://docs.atom.codes/sprout/registries/conversion#toint64
func (cr *ConversionRegistry) ToInt64(value any) (int64, error) {
	return cast.ToInt64E(value)
}

// ToUint converts a value to a uint.
//
// Parameters:
//
//	value any - the value to convert to uint. This value can be of any type that is numerically convertible.
//
// Returns:
//
//	uint - the uint representation of the value.
//	error - error if the value cannot be converted to a uint.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toUint].
//
// [Sprout Documentation: toUint]: https://docs.atom.codes/sprout/registries/conversion#touint
func (cr *ConversionRegistry) ToUint(value any) (uint, error) {
	return cast.ToUintE(value)
}

// ToUint64 converts a value to a uint64.
//
// Parameters:
//
//	value any - the value to convert to uint64. This value can be of any type that is numerically convertible.
//
// Returns:
//
//	uint64 - the uint64 representation of the value.
//	error - error if the value cannot be converted to a uint64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toUint64].
//
// [Sprout Documentation: toUint64]: https://docs.atom.codes/sprout/registries/conversion#touint64
func (cr *ConversionRegistry) ToUint64(value any) (uint64, error) {
	return cast.ToUint64E(value)
}

// ToFloat64 converts a value to a float64.
//
// Parameters:
//
//	value any - the value to convert to a float64.
//
// Returns:
//
//	float64 - the float64 representation of the value.
//	error - error if the value cannot be converted to a float64.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toFloat64].
//
// [Sprout Documentation: toFloat64]: https://docs.atom.codes/sprout/registries/conversion#tofloat64
func (cr *ConversionRegistry) ToFloat64(value any) (float64, error) {
	return cast.ToFloat64E(value)
}

// ToOctal parses a string value as an octal (base 8) integer.
//
// Parameters:
//
//	value any - the string representing an octal number.
//
// Returns:
//
//	int64 - the decimal (base 10) representation of the octal value.
//	error - error if the value cannot be converted to an octal number.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toOctal].
//
// [Sprout Documentation: toOctal]: https://docs.atom.codes/sprout/registries/conversion#tooctal
func (cr *ConversionRegistry) ToOctal(value any) (int64, error) {
	result, err := strconv.ParseInt(fmt.Sprint(value), 8, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse octal: %w", err)
	}
	return result, nil
}

// ToString converts a value to a string, handling various types effectively.
//
// Parameters:
//
//	value any - the value to convert to a string.
//
// Returns:
//
//	string - the string representation of the value.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toString].
//
// [Sprout Documentation: toString]: https://docs.atom.codes/sprout/registries/conversion#tostring
func (cr *ConversionRegistry) ToString(value any) string {
	switch value := value.(type) {
	case string:
		return value
	case []byte:
		return string(value)
	case error:
		return value.Error()
	case fmt.Stringer:
		return value.String()
	default:
		return fmt.Sprint(value)
	}
}

// ToDate converts a string to a time.Time object based on a format specification.
//
// Parameters:
//
//	layout string - the date format string.
//	value string - the date string to parse.
//
// Returns:
//
//	time.Time - the parsed date.
//	error - error if the date string does not conform to the format.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toDate].
//
// [Sprout Documentation: toDate]: https://docs.atom.codes/sprout/registries/conversion#todate
func (cr *ConversionRegistry) ToDate(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}

// ToLocalDate converts a string to a time.Time object based on a format specification
// and the local timezone.
//
// Parameters:
//
//	layout string - the date format string.
//	value string - the date string to parse.
//
// Returns:
//
//	time.Time - the parsed date.
//	error - error if the date string does not conform to the format.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toLocalDate].
//
// [Sprout Documentation: toLocalDate]: https://docs.atom.codes/sprout/registries/conversion#tolocaldate
func (cr *ConversionRegistry) ToLocalDate(layout, timezone, value string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation(layout, value, location)
}

// ToDuration converts a value to a time.Duration.
//
// Parameters:
//
//	value any - the value to convert to time.Duration. This value can be a string, int, or another compatible type.
//
// Returns:
//
//	time.Duration - the duration representation of the value.
//	error - error if the value cannot be converted to a duration.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toDuration].
//
// [Sprout Documentation: toDuration]: https://docs.atom.codes/sprout/registries/conversion#toduration
func (cr *ConversionRegistry) ToDuration(value any) (time.Duration, error) {
	return cast.ToDurationE(value)
}
