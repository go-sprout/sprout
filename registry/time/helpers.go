package time

import (
	"time"
)

// computeTimeFromFormat returns a time.Time object from the given date.
func computeTimeFromFormat(date any) time.Time {
	switch date := date.(type) {
	case time.Time:
		return date
	case *time.Time:
		return *date
	case int64:
		return time.Unix(date, 0)
	case int:
		return time.Unix(int64(date), 0)
	case int32:
		return time.Unix(int64(date), 0)
	}

	// otherwise, fallback to the current time
	return time.Now().Local()
}
