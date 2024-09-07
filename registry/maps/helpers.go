package maps

import (
	"fmt"
	"strings"
)

// digIntoDict navigates through a nested dictionary using a sequence of keys and returns the value found.
//
// Parameters:
//
//	dict map[string]any - the starting dictionary.
//	keys []string - a slice of keys to navigate through the dictionary.
//
// Returns:
//
//	any - the value found at the last key in the sequence.
//	error - an error if a key is not found or if the value at a key is not a dictionary when expected.
func (mr *MapsRegistry) digIntoDict(dict map[string]any, keys []string) (any, error) {
	current := dict
	for i, key := range keys {
		value, exists := current[key]
		if !exists {
			return nil, nil
		}
		if i == len(keys)-1 {
			return value, nil
		}

		nextDict, ok := value.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("value at key %q is not a nested dictionary but %T", key, value)
		}
		current = nextDict
	}
	return nil, fmt.Errorf("unexpected termination of key traversal")
}

// parseKeys converts a slice of any type to a slice of strings, ensuring all elements are strings.
//
// Parameters:
//
//	keySet []any - a slice containing potential keys.
//
// Returns:
//
//	[]string - a slice of strings if all elements in the original slice are strings.
//	error - an error if any element of the original slice is not a string.
//
// Example:
//
//	keys, _ := mr.parseKeys([]any{"key1", "key2"})
//	fmt.Println(keys) // Output: ["key1", "key2"]
//
//	keys, err := mr.parseKeys([]any{"key1", 2})
//	fmt.Println(err) // Output: all keys must be strings, got int at position 1
func (mr *MapsRegistry) parseKeys(keySet []any) ([]string, error) {
	// Calculate the total number of keys needed
	totalKeys := 0
	for i, element := range keySet {
		key, ok := element.(string)
		if !ok {
			return nil, fmt.Errorf("all keys must be strings, got %T at position %d", element, i)
		}
		totalKeys += 1 + strings.Count(key, ".")
	}

	// Preallocate the slice with the exact number of required elements
	keys := make([]string, 0, totalKeys)

	// Now, fill the slice
	for _, element := range keySet {
		key := element.(string)
		if strings.Contains(key, ".") {
			keys = append(keys, strings.Split(key, ".")...)
		} else {
			keys = append(keys, key)
		}
	}

	return keys, nil
}
