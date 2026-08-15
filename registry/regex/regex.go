// Package regex provides regular expression functions where the main parameter
// is always the last one, making every function usable in a template pipeline.
//
// It supersedes the [github.com/go-sprout/sprout/registry/regexp] registry,
// which keeps the historical sprig signatures and is deprecated. Both
// registries expose the same function names, so they are mutually exclusive:
// register one or the other, never both.
package regex

import "github.com/go-sprout/sprout"

type RegexRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of regex registry.
func NewRegistry() *RegexRegistry {
	return &RegexRegistry{}
}

// UID returns the unique identifier of the registry.
func (rr *RegexRegistry) UID() string {
	return "go-sprout/sprout.regex"
}

// LinkHandler links the handler to the registry at runtime.
func (rr *RegexRegistry) LinkHandler(fh sprout.Handler) error {
	rr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (rr *RegexRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "regexFind", rr.RegexFind)
	sprout.AddFunction(funcsMap, "regexFindAll", rr.RegexFindAll)
	sprout.AddFunction(funcsMap, "regexMatch", rr.RegexMatch)
	sprout.AddFunction(funcsMap, "regexSplit", rr.RegexSplit)
	sprout.AddFunction(funcsMap, "regexReplaceAll", rr.RegexReplaceAll)
	sprout.AddFunction(funcsMap, "regexReplaceAllLiteral", rr.RegexReplaceAllLiteral)
	sprout.AddFunction(funcsMap, "regexQuoteMeta", rr.RegexQuoteMeta)
	sprout.AddFunction(funcsMap, "regexFindGroups", rr.RegexFindGroups)
	sprout.AddFunction(funcsMap, "regexFindAllGroups", rr.RegexFindAllGroups)
	sprout.AddFunction(funcsMap, "regexFindNamed", rr.RegexFindNamed)
	sprout.AddFunction(funcsMap, "regexFindAllNamed", rr.RegexFindAllNamed)
	return nil
}
