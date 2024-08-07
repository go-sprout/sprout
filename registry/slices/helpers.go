package slices

import "reflect"

func (sr *SlicesRegistry) inList(haystack []any, needle any) bool {
	for _, h := range haystack {
		if reflect.DeepEqual(needle, h) {
			return true
		}
	}
	return false
}
