package maps

import (
	"errors"
)

func (mr *MapsRegistry) deprecatedGet(dict map[string]any, key string) (any, error) {
	if value, ok := dict[key]; ok {
		return value, nil
	}
	return "", nil
}

func (mr *MapsRegistry) deprecatedSet(dict map[string]any, key string, value any) (map[string]any, error) {
	dict[key] = value
	return dict, nil
}

func (mr *MapsRegistry) deprecatedUnset(dict map[string]any, key string) (map[string]any, error) {
	delete(dict, key)
	return dict, nil
}

func (mr *MapsRegistry) deprecatedHasKey(dict map[string]any, key string) (bool, error) {
	_, ok := dict[key]
	return ok, nil
}

func (mr *MapsRegistry) deprecatedPick(dict map[string]any, keys ...any) (map[string]any, error) {
	// Pre-allocate result map with the size of keys to avoid multiple allocations
	result := make(map[string]any, len(keys))

	for _, k := range keys {
		key, ok := k.(string)
		if !ok {
			return nil, errors.New("all keys must be strings")
		}
		if v, ok := dict[key]; ok {
			result[key] = v
		}
	}
	return result, nil
}

func (mr *MapsRegistry) deprecatedOmit(dict map[string]any, keys ...any) (map[string]any, error) {
	// Pre-allocate result map with the size of the original dictionary to avoid
	// multiple allocations
	result := make(map[string]any, len(dict))

	// Use a map for keys to omit for O(1) lookups
	omit := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		key, ok := k.(string)
		if !ok {
			return nil, errors.New("all keys must be strings")
		}
		omit[key] = struct{}{}
	}

	for key, value := range dict {
		if _, ok := omit[key]; !ok {
			result[key] = value
		}
	}
	return result, nil
}
