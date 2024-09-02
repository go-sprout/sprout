package time

import (
	"strconv"
	"strings"
	"time"
)

// Date formats a given date or current time into a specified format string.
//
// Parameters:
//
//	fmt string - the format string.
//	date any - the date to format or the current time if not a date type.
//
// Returns:
//
//	string - the formatted date.
//	error - when the timezone is invalid or the date is not in a valid format.
//
// Example:
//
//	{{ "2023-05-04T15:04:05Z" | date "Jan 2, 2006" }} // Output: "May 4, 2023"
func (tr *TimeRegistry) Date(fmt string, date any) (string, error) {
	return tr.DateInZone(fmt, date, "Local")
}

// DateInZone formats a given date or current time into a specified format string in a specified timezone.
//
// Parameters:
//
//	fmt string - the format string.
//	date any - the date to format, in various acceptable formats.
//	zone string - the timezone name.
//
// Returns:
//
//	string - the formatted date.
//	error - when the timezone is invalid or the date is not in a valid format.
//
// Example:
//
//	{{ dateInZone "Jan 2, 2006", "2023-05-04T15:04:05Z", "UTC" }} // Output: "May 4, 2023"
//
// TODO: Change signature
func (tr *TimeRegistry) DateInZone(fmt string, date any, zone string) (string, error) {
	var t time.Time
	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case *time.Time:
		t = *date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	case int32:
		t = time.Unix(int64(date), 0)
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		return t.In(time.UTC).Format(fmt), err
	}

	return t.In(loc).Format(fmt), nil
}

// Duration converts seconds into a human-readable duration string.
//
// Parameters:
//
//	sec any - the duration in seconds.
//
// Returns:
//
//	string - the human-readable duration.
//
// Example:
//
//	{{ 3661 | duration }} // Output: "1h1m1s"
func (tr *TimeRegistry) Duration(sec any) string {
	var n int64
	switch value := sec.(type) {
	case string:
		n, _ = strconv.ParseInt(value, 10, 64)
	case int64:
		n = value
	case int:
		n = int64(value)
	case float64:
		n = int64(value)
	case float32:
		n = int64(value)
	default:
		n = 0
	}
	return (time.Duration(n) * time.Second).String()
}

// DateAgo calculates how much time has passed since the given date.
//
// Parameters:
//
//	date any - the starting date for the calculation.
//
// Returns:
//
//	string - a human-readable string describing how long ago the date was.
//
// Example:
//
//	{{ "2023-05-04T15:04:05Z" | dateAgo }} // Output: "4m"
func (tr *TimeRegistry) DateAgo(date any) string {
	var t time.Time

	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case *time.Time:
		t = *date
	case int64:
		t = time.Unix(date, 0)
	case int32:
		t = time.Unix(int64(date), 0)
	case int:
		t = time.Unix(int64(date), 0)
	}
	// Drop resolution to seconds
	duration := time.Since(t).Round(time.Second)
	return duration.String()
}

// Now returns the current time.
//
// Returns:
//
//	time.Time - the current time.
//
// Example:
//
//	{{ now }} // Output: "2023-05-07T15:04:05Z"
func (tr *TimeRegistry) Now() time.Time {
	return time.Now()
}

// UnixEpoch returns the Unix epoch timestamp of a given date.
//
// Parameters:
//
//	date time.Time - the date to convert to a Unix timestamp.
//
// Returns:
//
//	string - the Unix timestamp as a string.
//
// Example:
//
//	{{ now | unixEpoch }} // Output: "1683306245"
func (tr *TimeRegistry) UnixEpoch(date time.Time) string {
	return strconv.FormatInt(date.Unix(), 10)
}

// DateModify adjusts a given date by a specified duration. If the duration
// format is incorrect, it returns an error.
//
// Parameters:
//   fmt string - the duration string to add to the date, such as "2h" for two hours.
//   date time.Time - the date to modify.
//
// Returns:
//   time.Time - the modified date after adding the duration
//	 error - an error if the duration format is incorrect
//
// Example:
//   {{ "2024-05-04T15:04:05Z" | dateModify "48h" }} // Outputs the date two days later

func (tr *TimeRegistry) DateModify(fmt string, date time.Time) (time.Time, error) {
	d, err := time.ParseDuration(fmt)
	if err != nil {
		return time.Time{}, err
	}
	return date.Add(d), nil
}

// DurationRound rounds a duration to the nearest significant unit, such as years or seconds.
//
// Parameters:
//
//	duration any - the duration to round.
//
// Returns:
//
//	string - the rounded duration.
//
// Example:
//
//	{{ "3600s" | durationRound }} // Output: "1h"
func (tr *TimeRegistry) DurationRound(duration any) string {
	var d time.Duration

	switch duration := duration.(type) {
	case string:
		d, _ = time.ParseDuration(duration)
	case int64:
		d = time.Duration(duration)
	case time.Duration:
		d = duration
	case time.Time:
		d = time.Since(duration)
	default:
		d = 0
	}

	u := uint64(d)
	neg := d < 0
	if neg {
		u = -u
	}

	if u == 0 {
		return "0s"
	}

	const (
		year   = uint64(time.Hour) * 24 * 365
		month  = uint64(time.Hour) * 24 * 30
		day    = uint64(time.Hour) * 24
		hour   = uint64(time.Hour)
		minute = uint64(time.Minute)
		second = uint64(time.Second)
	)

	var b strings.Builder
	b.Grow(3)

	if neg {
		b.WriteByte('-')
	}

	switch {
	case u > year:
		b.WriteString(strconv.FormatUint(u/year, 10))
		b.WriteRune('y')
	case u > month:
		b.WriteString(strconv.FormatUint(u/month, 10))
		b.WriteString("mo")
	case u > day:
		b.WriteString(strconv.FormatUint(u/day, 10))
		b.WriteRune('d')
	case u > hour:
		b.WriteString(strconv.FormatUint(u/hour, 10))
		b.WriteRune('h')
	case u > minute:
		b.WriteString(strconv.FormatUint(u/minute, 10))
		b.WriteRune('m')
	case u > second:
		b.WriteString(strconv.FormatUint(u/second, 10))
		b.WriteRune('s')
	}
	return b.String()
}

// HtmlDate formats a date into a standard HTML date format (YYYY-MM-DD).
//
// Parameters:
//
//	date any - the date to format.
//
// Returns:
//
//	string - the formatted date in HTML format.
//	error - when the timezone is invalid or the date is not in a valid format.
//
// Example:
//
//	{{ "2023-05-04T15:04:05Z" | htmlDate }} // Output: "2023-05-04"
func (tr *TimeRegistry) HtmlDate(date any) (string, error) {
	return tr.DateInZone("2006-01-02", date, "Local")
}

// HtmlDateInZone formats a date into a standard HTML date format (YYYY-MM-DD) in a specified timezone.
//
// Parameters:
//
//	date any - the date to format.
//	zone string - the timezone name.
//
// Returns:
//
//	string - the formatted date in HTML format.
//	error - when the timezone is invalid or the date is not in a valid format.
//
// Example:
//
//	{{ "2023-05-04T15:04:05Z", "UTC" | htmlDateInZone }} // Output: "2023-05-04"
//
// TODO: Change signature
func (tr *TimeRegistry) HtmlDateInZone(date any, zone string) (string, error) {
	return tr.DateInZone("2006-01-02", date, zone)
}
