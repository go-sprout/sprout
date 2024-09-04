package maps

import (
	"errors"
	"fmt"

	"dario.cat/mergo"

	"github.com/go-sprout/sprout/deprecated"
	"github.com/go-sprout/sprout/internal/helpers"
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
func (mr *MapsRegistry) Dict(values ...any) map[string]any {
	// Ensure even number of values for key-value pairs
	if len(values)%2 != 0 {
		values = append(values, "")
	}

	// Pre-allocate the map based on half the number of total elements,
	// since we expect every two elements to form a key-value pair.
	dict := make(map[string]any, len(values)/2)

	for i := 0; i < len(values); i += 2 {
		dict[helpers.ToString(values[i])] = values[i+1]
	}

	return dict
}

// Get retrieves the value associated with the specified key from the dictionary.
//
// Parameters:
//
//	key string - the key to look up.
//	dict map[string]any - the dictionary.
//
// Returns:
//
//	any - the value associated with the key, or an empty string if the key does not exist.
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ {"key": "value"} | get "key" }} // Output: "value"
func (mr *MapsRegistry) Get(args ...any) (any, error) {
	// ! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	// ! Due to change in signature, this function still supports the old signature
	// ! to let users transition to the new signature.
	// * Old signature: Get(map[string]any, string)
	// * New signature: Get(string, map[string]any)
	if len(args) != 2 {
		return "", deprecated.ErrArgsCount(2, len(args))
	}

	switch arg0 := args[0].(type) {
	case map[string]any:
		// Old signature
		deprecated.SignatureWarn(mr.handler.Logger(), "get", "{{ get dict key }}", "{{ dict | get key }}")
		return mr.Get(args[1].(string), arg0)
	case string:
		// New signature
		if value, ok := args[1].(map[string]any)[arg0]; ok {
			return value, nil
		}
		return "", nil
	default:
		return "", fmt.Errorf("expected map or string, got %T", arg0)
	}
}

// Set adds or updates a key with a specified value in the dictionary.
//
// Parameters:
//
//	key string - the key to set.
//	value any - the value to associate with the key.
//	dict map[string]any - the dictionary.
//
// Returns:
//
//	map[string]any - the updated dictionary.
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ {"key": "oldValue"} | set "key", "newValue" }} // Output: {"key": "newValue"}
func (mr *MapsRegistry) Set(args ...any) (map[string]any, error) {
	// ! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	// ! Due to change in signature, this function still supports the old signature
	// ! to let users transition to the new signature.
	// * Old signature: Set(map[string]any, string, string)
	// * New signature: Set(string, any, map[string]any)
	if len(args) != 3 {
		return nil, deprecated.ErrArgsCount(3, len(args))
	}

	switch arg0 := args[0].(type) {
	case map[string]any:
		// Old signature
		deprecated.SignatureWarn(mr.handler.Logger(), "set", "{{ set dict key value }}", "{{ dict | set key value }}")
		return mr.Set(args[1].(string), args[2], arg0)
	case string:
		// New signature
		if dict, ok := args[2].(map[string]any); ok {
			dict[arg0] = args[1]
			return dict, nil
		}
		return nil, errors.New("last argument must be a map[string]any")
	default:
		return nil, fmt.Errorf("expected map or string, got %T", arg0)
	}
}

// Unset removes a key from the dictionary.
//
// Parameters:
//
//	key string - the key to remove.
//	dict map[string]any - the dictionary.
//
// Returns:
//
//	map[string]any - the dictionary after removing the key.
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ {"key": "value"} | unset "key" }} // Output: {}
func (mr *MapsRegistry) Unset(args ...any) (map[string]any, error) {
	// ! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	// ! Due to change in signature, this function still supports the old signature
	// ! to let users transition to the new signature.
	// * Old signature: Unset(map[string]any, string)
	// * New signature: Unset(string, map[string]any)
	if len(args) != 2 {
		return nil, deprecated.ErrArgsCount(2, len(args))
	}

	switch arg0 := args[0].(type) {
	case map[string]any:
		// Old signature
		deprecated.SignatureWarn(mr.handler.Logger(), "unset", "{{ unset dict key }}", "{{ dict | unset key }}")
		return mr.Unset(args[1].(string), arg0)
	case string:
		// New signature
		if dict, ok := args[1].(map[string]any); ok {
			delete(dict, arg0)
			return dict, nil
		}
		return nil, errors.New("last argument must be a map[string]any")
	default:
		return nil, fmt.Errorf("expected map or string, got %T", arg0)
	}
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
func (mr *MapsRegistry) Keys(dicts ...map[string]any) []string {
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

// Values retrieves all values from one or more dictionaries.
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
func (mr *MapsRegistry) Values(dicts ...map[string]any) []any {
	var keyCount int
	for i := range dicts {
		keyCount += len(dicts[i])
	}

	values := make([]any, 0, keyCount)

	for _, dict := range dicts {
		for _, value := range dict {
			values = append(values, value)
		}
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
//	{{ $d1 := dict "key" "value1"}}
//	{{ $d2 := dict "key" "value2" }}
//	{{ pluck "key"	$d1 $d2 }} // Output: ["value1", "value2"]
func (mr *MapsRegistry) Pluck(key string, dicts ...map[string]any) []any {
	result := make([]any, 0, len(dicts))

	for _, dict := range dicts {
		if val, ok := dict[key]; ok {
			result = append(result, val)
		}
	}
	return result
}

// Pick creates a new dictionary containing only the specified keys from the
// original dictionary.
//
// Parameters:
//
//	keys ...string - the keys to include in the new dictionary.
//	dict map[string]any - the source dictionary.
//
// Returns:
//
//	map[string]any - a dictionary containing only the picked keys and their values.
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
//	{{ $d | pick "key1" "key3" }} // Output: {"key1": "value1", "key3": "value3"}
func (mr *MapsRegistry) Pick(args ...any) (map[string]any, error) {
	// ! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	// ! Due to change in signature, this function still supports the old signature
	// ! to let users transition to the new signature.
	// * Old signature: Pick(map[string]any, ...string)
	// * New signature: Pick(...string, map[string]any)
	if len(args) < 2 {
		return nil, deprecated.ErrArgsCount(2, len(args))
	}

	// Pre-allocate result map with the size of keys to avoid multiple allocations
	result := make(map[string]any, len(args)-1) // Remove the last argument which is the dictionary

	switch arg0 := args[0].(type) {
	case map[string]any:
		// Old signature
		deprecated.SignatureWarn(mr.handler.Logger(), "pick", "{{ pick dict key1 key2 }}", "{{ dict | pick key1 key2 }}")
		return mr.Pick(append(args[1:], args[0])...)
	case string:
		// New signature
		keys := args[:len(args)-1]
		dict, ok := args[len(args)-1].(map[string]any)
		if !ok {
			return nil, errors.New("last argument must be a map[string]any")
		}

		for _, key := range keys {
			keyStr, ok := key.(string)
			if !ok {
				return nil, errors.New("all keys must be strings")
			}
			if value, ok := dict[keyStr]; ok {
				result[keyStr] = value
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("expected map or string, got %T", arg0)
	}
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
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ $d := dict "key1" "value1" "key2" "value2" "key3" "value3" }}
//	{{ omit $d "key1" "key3" }} // Output: {"key2": "value2"}
func (mr *MapsRegistry) Omit(args ...any) (map[string]any, error) {
	// ! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	// ! Due to change in signature, this function still supports the old signature
	// ! to let users transition to the new signature.
	// * Old signature: Omit(map[string]any, ...string)
	// * New signature: Omit(...string, map[string]any)
	if len(args) < 2 {
		return nil, deprecated.ErrArgsCount(2, len(args))
	}

	// Pre-allocate result map with the size of keys to avoid multiple allocations
	result := make(map[string]any, len(args)-1) // Remove the last argument which is the dictionary

	switch arg0 := args[0].(type) {
	case map[string]any:
		// Old signature
		deprecated.SignatureWarn(mr.handler.Logger(), "omit", "{{ omit dict key1 key2 }}", "{{ dict | omit key1 key2 }}")
		return mr.Omit(append(args[1:], args[0])...)
	case string:
		// New signature
		keys := args[:len(args)-1]
		dict, ok := args[len(args)-1].(map[string]any)
		if !ok {
			return nil, errors.New("last argument must be a map[string]any")
		}
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
	default:
		return nil, fmt.Errorf("expected map or string, got %T", arg0)
	}
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
//	{{ dig "user.profile.age", {"user": {"profile": {"name": "John Doe"}}} }} // Output: nil, nil
func (mr *MapsRegistry) Dig(args ...any) (any, error) {
	if len(args) < 2 {
		return nil, errors.New("dig requires at least two arguments: a sequence of keys and a dictionary")
	}

	dict, ok := args[len(args)-1].(map[string]any)
	if !ok {
		return nil, errors.New("last argument must be a map[string]any")
	}

	keys, err := mr.parseKeys(args[:len(args)-1])
	if err != nil {
		return nil, fmt.Errorf("cannot parse keys: %w", err)
	}

	return mr.digIntoDict(dict, keys)
}

// HasKey checks if the specified key exists in the dictionary.
//
// Parameters:
//
//	key string - the key to look for.
//	dict map[string]any - the dictionary to check.
//
// Returns:
//
//	bool - true if the key exists, otherwise false.
//	error - protect against undesired behavior due to migration to new signature.
//
// Example:
//
//	{{ {"key": "value"} | hasKey "key" }} // Output: true
func (mr *MapsRegistry) HasKey(args ...any) (bool, error) {
	// ! BACKWARDS COMPATIBILITY: deprecated in v1.0 and removed in v1.1
	// ! Due to change in signature, this function still supports the old signature
	// ! to let users transition to the new signature.
	// * Old signature: HasKey(dict map[string]any, key string)
	// * New signature: HasKey(key string, map[string]any)
	if len(args) != 2 {
		return false, deprecated.ErrArgsCount(2, len(args))
	}

	switch arg0 := args[0].(type) {
	case map[string]any:
		// Old signature
		deprecated.SignatureWarn(mr.handler.Logger(), "hasKey", "{{ hasKey dict key }}", "{{ dict | hasKey key }}")
		return mr.HasKey(args[1].(string), arg0)
	case string:
		// New signature
		_, ok := args[1].(map[string]any)[arg0]
		return ok, nil
	default:
		return false, fmt.Errorf("expected map or string, got %T", arg0)
	}
}

// Merge merges multiple source maps into a destination map without
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
//	{{ merge {}, {"a": 1, "b": 2}, {"b": 3, "c": 4}  }} // Output: {"a": 1, "b": 2, "c": 4}, nil
func (mr *MapsRegistry) Merge(dest map[string]any, srcs ...map[string]any) (any, error) {
	for _, src := range srcs {
		if err := mergo.Merge(&dest, src, mergo.WithoutDereference); err != nil {
			// This error is not expected to occur, as we ensure types are correct in
			// the function signature. If it does, it is a bug in the function implementation.
			return nil, err
		}
	}
	return dest, nil
}

// MergeOverwrite merges multiple source maps into a destination map,
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
//	{{ mergeOverwrite {}, {"a": 1, "b": 2}, {"b": 3, "c": 4} }} // Output: {"a": 1, "b": 3, "c": 4}, nil
func (mr *MapsRegistry) MergeOverwrite(dest map[string]any, srcs ...map[string]any) (any, error) {
	for _, src := range srcs {
		if err := mergo.Merge(&dest, src, mergo.WithOverride, mergo.WithoutDereference); err != nil {
			// This error is not expected to occur, as we ensure types are correct in
			// the function signature. If it does, it is a bug in the function implementation.
			return nil, err
		}
	}
	return dest, nil
}
