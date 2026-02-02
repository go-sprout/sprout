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
	keys := make([]string, 0, len(keySet))

	// Convert each key to string
	for i, element := range keySet {
		key, ok := element.(string)
		if !ok {
			return nil, fmt.Errorf("all keys must be strings, got %T at position %d", element, i)
		}
		keys = append(keys, key)
	}

	return keys, nil
}

// splitKeysWithEscapes breaks up nested keys into simple key elements, splitting
// on unescaped dots only. Escape sequences are resolved in the output.
//
// Escape sequences:
//   - \. → literal dot (not a path separator)
//   - \\ → literal backslash
//   - \x (other) → error
//
// Parameters:
//
//	keySet []string - a slice with simple and nested keys, possibly containing escape sequences.
//
// Returns:
//
//	[]string - a slice containing only simple keys with escape sequences resolved.
//	error - an error if an invalid escape sequence is encountered.
//
// Example:
//
//	keys, _ := mr.splitKeysWithEscapes([]string{"a", "nested.key"})
//	fmt.Printf("%q\n", keys) // Output: ["a" "nested" "key"]
//
//	keys, _ := mr.splitKeysWithEscapes([]string{"a\\.b"})
//	fmt.Printf("%q\n", keys) // Output: ["a.b"]
//
//	keys, _ := mr.splitKeysWithEscapes([]string{"a\\\\b"})
//	fmt.Printf("%q\n", keys) // Output: ["a\\b"]
func (mr *MapsRegistry) splitKeysWithEscapes(keySet []string) ([]string, error) {
	var result []string

	for _, key := range keySet {
		// Fast path: if no backslash, use simple split
		if !strings.Contains(key, "\\") {
			parts := strings.Split(key, ".")
			// Validate no empty segments (consecutive, leading, or trailing dots)
			for _, part := range parts {
				if part == "" && len(parts) > 1 {
					return nil, fmt.Errorf("empty key segment in path %q (consecutive or leading/trailing dots)", key)
				}
			}
			result = append(result, parts...)
			continue
		}

		// Slow path: handle escape sequences
		parts, err := mr.splitKeyOnUnescapedDots(key)
		if err != nil {
			return nil, err
		}
		result = append(result, parts...)
	}

	return result, nil
}

// splitKeyOnUnescapedDots splits a key on unescaped dots and resolves escape sequences.
// Empty segments (from consecutive dots like "a..b") are preserved intentionally,
// as empty strings are valid map keys in Go.
//
// Parameters:
//
//	key string - a key that may contain escape sequences.
//
// Returns:
//
//	[]string - the key split on unescaped dots, with escapes resolved.
//	error - an error if an invalid escape sequence is found or if the key contains empty segments.
func (mr *MapsRegistry) splitKeyOnUnescapedDots(key string) ([]string, error) {
	var parts []string
	var segment strings.Builder
	lastWasDot := true // Start true to detect leading dot

	runes := []rune(key)
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Handle unescaped dot as path separator
		if r == '.' {
			if lastWasDot {
				return nil, fmt.Errorf("empty key segment in path %q (consecutive or leading/trailing dots)", key)
			}
			parts = append(parts, segment.String())
			segment.Reset()
			lastWasDot = true
			continue
		}

		lastWasDot = false

		// Handle non-escape characters
		if r != '\\' {
			segment.WriteRune(r)
			continue
		}

		// Handle escape sequences
		if i+1 >= len(runes) {
			return nil, fmt.Errorf("invalid escape sequence: trailing backslash in key %q", key)
		}

		next := runes[i+1]
		switch next {
		case '.', '\\':
			segment.WriteRune(next)
		default:
			return nil, fmt.Errorf("invalid escape sequence: \\%c in key %q", next, key)
		}
		i++
	}

	// Check for trailing dot
	if lastWasDot && len(runes) > 0 {
		return nil, fmt.Errorf("empty key segment in path %q (consecutive or leading/trailing dots)", key)
	}

	parts = append(parts, segment.String())
	return parts, nil
}
