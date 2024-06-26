package sprout

import (
	"fmt"

	"dario.cat/mergo"
)

// Dict creates a dictionary from a list of keys and values.
//
// Parameters:
//
//	values ...any - alternating keys and values.
//
// Returns:
//
//	map[string]any - the created dictionary.
//
// Example:
//
//	{{ dict "key1", "value1", "key2", "value2" }} // Output: {"key1": "value1", "key2": "value2"}
func (fh *FunctionHandler) Dict(values ...any) map[string]any {
	// Ensure even number of values for key-value pairs
	if len(values)%2 != 0 {
		values = append(values, "")
	}

	// Pre-allocate the map based on half the number of total elements,
	// since we expect every two elements to form a key-value pair.
	dict := make(map[string]any, len(values)/2)

	for i := 0; i < len(values); i += 2 {
		dict[fh.ToString(values[i])] = values[i+1]
	}

	return dict
}

// Get retrieves the value associated with the specified key from the dictionary.
//
// Parameters:
//
//	dict map[string]any - the dictionary.
//	key string - the key to look up.
//
// Returns:
//
//	any - the value associated with the key, or an empty string if the key does not exist.
//
// Example:
//
//	{{ get {"key": "value"}, "key" }} // Output: "value"
func (fh *FunctionHandler) Get(dict map[string]any, key string) any {
	if value, ok := dict[key]; ok {
		return value
	}
	return ""
}

// Set adds or updates a key with a specified value in the dictionary.
//
// Parameters:
//
//	dict map[string]any - the dictionary.
//	key string - the key to set.
//	value any - the value to associate with the key.
//
// Returns:
//
//	map[string]any - the updated dictionary.
//
// Example:
//
//	{{ set {"key": "oldValue"}, "key", "newValue" }} // Output: {"key": "newValue"}
func (fh *FunctionHandler) Set(dict map[string]any, key string, value any) map[string]any {
	dict[key] = value
	return dict
}

// Unset removes a key from the dictionary.
//
// Parameters:
//
//	dict map[string]any - the dictionary.
//	key string - the key to remove.
//
// Returns:
//
//	map[string]any - the dictionary after removing the key.
//
// Example:
//
//	{{ {"key": "value"}, "key" | unset }} // Output: {}
func (fh *FunctionHandler) Unset(dict map[string]any, key string) map[string]any {
	delete(dict, key)
	return dict
}

// Keys retrieves all keys from one or more dictionaries.
//
// Parameters:
//
//	dicts ...map[string]any - one or more dictionaries.
//
// Returns:
//
//	[]string - a list of all keys from the dictionaries.
//
// Example:
//
//	{{ keys {"key1": "value1", "key2": "value2"} }} // Output: ["key1", "key2"]
func (fh *FunctionHandler) Keys(dicts ...map[string]any) []string {
	var keyCount int
	for i := range dicts {
		keyCount += len(dicts[i])
	}

	keys := make([]string, 0, keyCount)

	for _, dict := range dicts {
		for key := range dict {
			keys = append(keys, key)
		}
	}

	return keys
}

// Values retrieves all values from a dictionary.
//
// Parameters:
//
//	dict map[string]any - the dictionary.
//
// Returns:
//
//	[]any - a list of all values from the dictionary.
//
// Example:
//
//	{{ values {"key1": "value1", "key2": "value2"} }} // Output: ["value1", "value2"]
func (fh *FunctionHandler) Values(dict map[string]any) []any {
	var values = make([]any, 0, len(dict))
	for _, value := range dict {
		values = append(values, value)
	}

	return values
}

// Pluck extracts values associated with a specified key from a list of dictionaries.
//
// Parameters:
//
//	key string - the key to pluck values for.
//	dicts ...map[string]any - one or more dictionaries.
//
// Returns:
//
//	[]any - a list of values associated with the key from each dictionary.
//
// Example:
//
//	{{ [{"key": "value1"}, {"key": "value2"}] | pluck "key" }} // Output: ["value1", "value2"]
func (fh *FunctionHandler) Pluck(key string, dicts ...map[string]any) []any {
	result := []any{}
	for _, dict := range dicts {
		if val, ok := dict[key]; ok {
			result = append(result, val)
		}
	}
	return result
}

// Pick creates a new dictionary containing only the specified keys from the original dictionary.
//
// Parameters:
//
//	dict map[string]any - the source dictionary.
//	keys ...string - the keys to include in the new dictionary.
//
// Returns:
//
//	map[string]any - a dictionary containing only the picked keys and their values.
//
// Example:
//
//	{{ pick {"key1": "value1", "key2": "value2", "key3": "value3"}, "key1", "key3" }} // Output: {"key1": "value1", "key3": "value3"}
func (fh *FunctionHandler) Pick(dict map[string]any, keys ...string) map[string]any {
	result := map[string]any{}
	for _, k := range keys {
		if v, ok := dict[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Omit creates a new dictionary by excluding specified keys from the original dictionary.
//
// Parameters:
//
//	dict map[string]any - the source dictionary.
//	keys ...string - the keys to exclude from the new dictionary.
//
// Returns:
//
//	map[string]any - a dictionary without the omitted keys.
//
// Example:
//
//	{{ omit {"key1": "value1", "key2": "value2", "key3": "value3"}, "key2" }} // Output: {"key1": "value1", "key3": "value3"}
func (fh *FunctionHandler) Omit(dict map[string]any, keys ...string) map[string]any {
	result := map[string]any{}

	omit := make(map[string]bool, len(keys))
	for _, key := range keys {
		omit[key] = true
	}

	for key, value := range dict {
		if _, ok := omit[key]; !ok {
			result[key] = value
		}
	}
	return result
}

// Dig navigates through a nested dictionary structure using a sequence of keys
// and returns the value found at the specified path.
//
// Parameters:
//
//	args ...any - a sequence of keys followed by a dictionary as the last argument.
//
// Returns:
//
//	any - the value found at the nested key path or nil if any key in the path is not found.
//	error - an error if there are fewer than three arguments, if the last argument is not a dictionary, or if any key is not a string.
//
// Example:
//
//	{{ dig "user", "profile", "name", {"user": {"profile": {"name": "John Doe"}}} }} // Output: "John Doe", nil
func (fh *FunctionHandler) Dig(args ...any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("dig requires at least two arguments: a sequence of keys and a dictionary")
	}

	dict, ok := args[len(args)-1].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("last argument must be a map[string]any")
	}

	keys, err := fh.parseKeys(args[:len(args)-1])
	if err != nil {
		return nil, err
	}

	return fh.digIntoDict(dict, keys)
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
//	keys, _ := fh.parseKeys([]any{"key1", "key2"})
//	fmt.Println(keys) // Output: ["key1", "key2"]
//
//	keys, err := fh.parseKeys([]any{"key1", 2})
//	fmt.Println(err) // Output: all keys must be strings, got int at position 1
func (fh *FunctionHandler) parseKeys(keySet []any) ([]string, error) {
	keys := make([]string, len(keySet))
	for i, element := range keySet {
		key, ok := element.(string)
		if !ok {
			return nil, fmt.Errorf("all keys must be strings, got %T at position %d", element, i)
		}
		keys[i] = key
	}
	return keys, nil
}

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
func (fh *FunctionHandler) digIntoDict(dict map[string]any, keys []string) (any, error) {
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

// HasKey checks if the specified key exists in the dictionary.
//
// Parameters:
//
//	dict map[string]any - the dictionary to check.
//	key string - the key to look for.
//
// Returns:
//
//	bool - true if the key exists, otherwise false.
//
// Example:
//
//	{{ hasKey {"key": "value"}, "key" }} // Output: true
func (fh *FunctionHandler) HasKey(dict map[string]any, key string) bool {
	_, ok := dict[key]
	return ok
}

// Merge combines multiple source maps into a destination map without
// overwriting existing keys.
//
// Parameters:
//
//	dest map[string]any - the destination map.
//	srcs ...map[string]any - one or more source maps to merge into the destination.
//
// Returns:
//
//	any - the merged destination map.
//
// Example:
//
//	{{ merge {}, {"a": 1}, {"b": 2} }} // Output: {"a": 1, "b": 2}
func (fh *FunctionHandler) Merge(dest map[string]any, srcs ...map[string]any) any {
	result, _ := fh.MustMerge(dest, srcs...)
	return result
}

// MergeOverwrite combines multiple source maps into a destination map,
// overwriting existing keys.
//
// Parameters:
//
//	dest map[string]any - the destination map.
//	srcs ...map[string]any - one or more source maps to merge into the destination, with overwriting.
//
// Returns:
//
//	any - the merged destination map with overwritten values where applicable.
//
// Example:
//
//	{{ mergeOverwrite {}, {"a": 1}, {"a": 2, "b": 3} }} // Output: {"a": 2, "b": 3}
func (fh *FunctionHandler) MergeOverwrite(dest map[string]any, srcs ...map[string]any) any {
	result, _ := fh.MustMergeOverwrite(dest, srcs...)
	return result
}

// MustMerge merges multiple source maps into a destination map without
// overwriting existing keys in the destination.
// If an error occurs during merging, it returns nil and the error.
//
// Parameters:
//
//	dest map[string]any - the destination map to which all source map key-values are added.
//	srcs ...map[string]any - one or more source maps whose key-values are added to the destination.
//
// Returns:
//
//	any - the merged destination map.
//	error - error if the merge fails.
//
// Example:
//
//	{{ mustMerge {}, {"a": 1, "b": 2}, {"b": 3, "c": 4}  }} // Output: {"a": 1, "b": 2, "c": 4}, nil
func (fh *FunctionHandler) MustMerge(dest map[string]any, srcs ...map[string]any) (any, error) {
	for _, src := range srcs {
		if err := mergo.Merge(&dest, src, mergo.WithoutDereference); err != nil {
			// This error is not expected to occur, as we ensure types are correct in
			// the function signature. If it does, it is a bug in the function implementation.
			return nil, err
		}
	}
	return dest, nil
}

// MustMergeOverwrite merges multiple source maps into a destination map,
// overwriting existing keys in the destination.
// If an error occurs during merging, it returns nil and the error.
//
// Parameters:
//
//	dest map[string]any - the destination map to which all source map key-values are added.
//	srcs ...map[string]any - one or more source maps whose key-values are added to the destination, potentially overwriting existing keys.
//
// Returns:
//
//	any - the merged destination map with overwritten values where applicable.
//	error - error if the merge fails.
//
// Example:
//
//	{{ mustMergeOverwrite {}, {"a": 1, "b": 2}, {"b": 3, "c": 4} }} // Output: {"a": 1, "b": 3, "c": 4}, nil
func (fh *FunctionHandler) MustMergeOverwrite(dest map[string]any, srcs ...map[string]any) (any, error) {
	for _, src := range srcs {
		if err := mergo.Merge(&dest, src, mergo.WithOverride, mergo.WithoutDereference); err != nil {
			// This error is not expected to occur, as we ensure types are correct in
			// the function signature. If it does, it is a bug in the function implementation.
			return nil, err
		}
	}
	return dest, nil
}
