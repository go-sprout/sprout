package maps

import "github.com/go-sprout/sprout"

type MapsRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of maps registry.
func NewRegistry() *MapsRegistry {
	return &MapsRegistry{}
}

// Uid returns the unique identifier of the registry.
func (mr *MapsRegistry) Uid() string {
	return "maps"
}

// LinkHandler links the handler to the registry at runtime.
func (mr *MapsRegistry) LinkHandler(fh sprout.Handler) error {
	mr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (mr *MapsRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "dict", mr.Dict)
	sprout.AddFunction(funcsMap, "get", mr.Get)
	sprout.AddFunction(funcsMap, "set", mr.Set)
	sprout.AddFunction(funcsMap, "unset", mr.Unset)
	sprout.AddFunction(funcsMap, "keys", mr.Keys)
	sprout.AddFunction(funcsMap, "values", mr.Values)
	sprout.AddFunction(funcsMap, "pluck", mr.Pluck)
	sprout.AddFunction(funcsMap, "pick", mr.Pick)
	sprout.AddFunction(funcsMap, "omit", mr.Omit)
	sprout.AddFunction(funcsMap, "dig", mr.Dig)
	sprout.AddFunction(funcsMap, "hasKey", mr.HasKey)
	sprout.AddFunction(funcsMap, "merge", mr.Merge)
	sprout.AddFunction(funcsMap, "mergeOverwrite", mr.MergeOverwrite)
	return nil
}

func (mr *MapsRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "merge", "mustMerge")
	sprout.AddAlias(aliasesMap, "mergeOverwrite", "mustMergeOverwrite")
	return nil
}

func (mr *MapsRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustMerge", "please use `merge` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustMergeOverwrite", "please use `mergeOverwrite` instead"))
	return nil
}
