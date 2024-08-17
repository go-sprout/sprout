package slices

import "reflect"

// inList checks if a value is in a slice of any type by comparing values.
func (sr *SlicesRegistry) inList(haystack []any, needle any) bool {
	for _, h := range haystack {
		if sr.isComparable(h) && h == needle {
			return true
		}

		if reflect.TypeOf(h) == reflect.TypeOf(needle) && reflect.DeepEqual(needle, h) {
			return true
		}
	}
	return false
}

// isComparable checks if a value is a comparable type.
func (sr *SlicesRegistry) isComparable(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64, string, bool:
		return true
	default:
		return false
	}
}
