package backward

func (bcr *BackwardCompatibilityRegistry) get(dict map[string]any, key string) any {
	if value, ok := dict[key]; ok {
		return value
	}
	return ""
}
