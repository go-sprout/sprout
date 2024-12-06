package backward

// get retrieves the value associated with the specified key from the given dictionary.
// If the key exists, it returns the corresponding value; otherwise, it returns an empty string.
//
// Parameters:
//
//	dict map[string]any - the dictionary to search for the key.
//	key string - the key whose associated value is to be returned.
//
// Returns:
//
//	any - the value associated with the specified key, or an empty string if the key does not exist.
func (bcr *BackwardCompatibilityRegistry) get(dict map[string]any, key string) any {
	if value, ok := dict[key]; ok {
		return value
	}
	return ""
}
