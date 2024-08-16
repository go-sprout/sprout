package regexp

import "github.com/go-sprout/sprout"

type RegexpRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of regexp registry.
func NewRegistry() *RegexpRegistry {
	return &RegexpRegistry{}
}

// Uid returns the unique identifier of the registry.
func (rr *RegexpRegistry) Uid() string {
	return "regexp"
}

// LinkHandler links the handler to the registry at runtime.
func (rr *RegexpRegistry) LinkHandler(fh sprout.Handler) error {
	rr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (rr *RegexpRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "regexFind", rr.RegexFind)
	sprout.AddFunction(funcsMap, "regexFindAll", rr.RegexFindAll)
	sprout.AddFunction(funcsMap, "regexMatch", rr.RegexMatch)
	sprout.AddFunction(funcsMap, "regexSplit", rr.RegexSplit)
	sprout.AddFunction(funcsMap, "regexReplaceAll", rr.RegexReplaceAll)
	sprout.AddFunction(funcsMap, "regexReplaceAllLiteral", rr.RegexReplaceAllLiteral)
	sprout.AddFunction(funcsMap, "regexQuoteMeta", rr.RegexQuoteMeta)
	sprout.AddFunction(funcsMap, "mustRegexFind", rr.MustRegexFind)
	sprout.AddFunction(funcsMap, "mustRegexFindAll", rr.MustRegexFindAll)
	sprout.AddFunction(funcsMap, "mustRegexMatch", rr.MustRegexMatch)
	sprout.AddFunction(funcsMap, "mustRegexSplit", rr.MustRegexSplit)
	sprout.AddFunction(funcsMap, "mustRegexReplaceAll", rr.MustRegexReplaceAll)
	sprout.AddFunction(funcsMap, "mustRegexReplaceAllLiteral", rr.MustRegexReplaceAllLiteral)
	return nil
}
