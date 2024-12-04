package slices

import (
	"reflect"
)

// inList checks if the needle is present in the haystack slice.
// It returns true if an element in the haystack is equal to the needle.
// The comparison is performed first by equality if both elements are comparable,
// and then by deep equality if they are of the same type. It returns false otherwise.
// Parameters:
//
//	haystack - list of values to search in
//	needle - value to search for
//
// Returns:
//
//	true if the needle is found in haystack, false otherwise.
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

// isComparable checks if the given value is of a type that can be compared using
// the equality operator (==). It returns true for basic types such as integers,
// floating-point numbers, strings, and booleans, and false otherwise.
//
// Parameters:
//
//	v - the value to check for comparability.
//
// Returns:
//
//	true if the value is a basic comparable type, false otherwise.
func (sr *SlicesRegistry) isComparable(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64, string, bool:
		return true
	default:
		return false
	}
}

// flattenSlice takes a slice or array and recursively flattens it into a
// single-dimensional list of elements. The remainingDeep parameter controls
// the maximum depth of the flattening, with -1 indicating infinite depth and
// 0 indicating no recursion. The function returns a slice of the flattened
// elements.
//
// Parameters:
//
//	val - the slice or array to flatten
//	remainingDeep - the maximum depth of recursion
//
// Returns:
//
//	a single-dimensional list of elements from the input slice or array.
func (sr *SlicesRegistry) flattenSlice(val reflect.Value, remainingDeep int) []any {
	result := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i)

		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}

		if (item.Kind() == reflect.Slice || item.Kind() == reflect.Array) && (remainingDeep > 0 || remainingDeep <= -1) {
			result = append(result, sr.flattenSlice(item, remainingDeep-1)...)
		} else {
			result = append(result, item.Interface())
		}
	}

	return result
}
