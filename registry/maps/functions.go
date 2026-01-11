package maps

import (
	"errors"
	"fmt"

	"dario.cat/mergo"

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
// For an example of this function in a Go template, refer to [Sprout Documentation: dict].
//
// [Sprout Documentation: dict]: https://docs.atom.codes/sprout/registries/maps#dict
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
//	error - error if arguments are invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: get].
//
// [Sprout Documentation: get]: https://docs.atom.codes/sprout/registries/maps#get
func (mr *MapsRegistry) Get(key string, dict map[string]any) (any, error) {
	if value, ok := dict[key]; ok {
		return value, nil
	}
	return "", nil
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
//	error - error if arguments are invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: set].
//
// [Sprout Documentation: set]: https://docs.atom.codes/sprout/registries/maps#set
func (mr *MapsRegistry) Set(key string, value any, dict map[string]any) (map[string]any, error) {
	dict[key] = value
	return dict, nil
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
//	error - error if arguments are invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: unset].
//
// [Sprout Documentation: unset]: https://docs.atom.codes/sprout/registries/maps#unset
func (mr *MapsRegistry) Unset(key string, dict map[string]any) (map[string]any, error) {
	delete(dict, key)
	return dict, nil
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
// For an example of this function in a Go template, refer to [Sprout Documentation: keys].
//
// [Sprout Documentation: keys]: https://docs.atom.codes/sprout/registries/maps#keys
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
// For an example of this function in a Go template, refer to [Sprout Documentation: values].
//
// [Sprout Documentation: values]: https://docs.atom.codes/sprout/registries/maps#values
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
// For an example of this function in a Go template, refer to [Sprout Documentation: pluck].
//
// [Sprout Documentation: pluck]: https://docs.atom.codes/sprout/registries/maps#pluck
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
//	dict map[string]any - the source dictionary (last argument).
//
// Returns:
//
//	map[string]any - a dictionary containing only the picked keys and their values.
//	error - error if arguments are invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: pick].
//
// [Sprout Documentation: pick]: https://docs.atom.codes/sprout/registries/maps#pick
func (mr *MapsRegistry) Pick(args ...any) (map[string]any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("pick requires at least two arguments")
	}

	keys := args[:len(args)-1]
	dict, ok := args[len(args)-1].(map[string]any)
	if !ok {
		return nil, errors.New("last argument must be a map[string]any")
	}

	result := make(map[string]any, len(keys))
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
}

// Omit creates a new dictionary by excluding specified keys from the original dictionary.
//
// Parameters:
//
//	keys ...string - the keys to exclude from the new dictionary.
//	dict map[string]any - the source dictionary (last argument).
//
// Returns:
//
//	map[string]any - a dictionary without the omitted keys.
//	error - error if arguments are invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: omit].
//
// [Sprout Documentation: omit]: https://docs.atom.codes/sprout/registries/maps#omit
func (mr *MapsRegistry) Omit(args ...any) (map[string]any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("omit requires at least two arguments")
	}

	keys := args[:len(args)-1]
	dict, ok := args[len(args)-1].(map[string]any)
	if !ok {
		return nil, errors.New("last argument must be a map[string]any")
	}

	omitSet := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		key, ok := k.(string)
		if !ok {
			return nil, errors.New("all keys must be strings")
		}
		omitSet[key] = struct{}{}
	}

	result := make(map[string]any, len(dict))
	for key, value := range dict {
		if _, ok := omitSet[key]; !ok {
			result[key] = value
		}
	}
	return result, nil
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
// For an example of this function in a Go template, refer to [Sprout Documentation: dig].
//
// [Sprout Documentation: dig]: https://docs.atom.codes/sprout/registries/maps#dig
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
	keys = mr.splitKeys(keys)

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
//	error - error if arguments are invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: hasKey].
//
// [Sprout Documentation: hasKey]: https://docs.atom.codes/sprout/registries/maps#haskey
func (mr *MapsRegistry) HasKey(key string, dict map[string]any) (bool, error) {
	_, ok := dict[key]
	return ok, nil
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
// For an example of this function in a Go template, refer to [Sprout Documentation: merge].
//
// [Sprout Documentation: merge]: https://docs.atom.codes/sprout/registries/maps#merge
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
// For an example of this function in a Go template, refer to [Sprout Documentation: mergeOverwrite].
//
// [Sprout Documentation: mergeOverwrite]: https://docs.atom.codes/sprout/registries/maps#mergeoverwrite
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
