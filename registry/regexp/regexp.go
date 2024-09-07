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
	return nil
}

func (rr *RegexpRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "regexFind", "mustRegexFind")
	sprout.AddAlias(aliasesMap, "regexFindAll", "mustRegexFindAll")
	sprout.AddAlias(aliasesMap, "regexMatch", "mustRegexMatch")
	sprout.AddAlias(aliasesMap, "regexSplit", "mustRegexSplit")
	sprout.AddAlias(aliasesMap, "regexReplaceAll", "mustRegexReplaceAll")
	sprout.AddAlias(aliasesMap, "regexReplaceAllLiteral", "mustRegexReplaceAllLiteral")
	return nil
}

func (rr *RegexpRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexFind", "please use `regexFind` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexFindAll", "please use `regexFindAll` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexMatch", "please use `regexMatch` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexSplit", "please use `regexSplit` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexReplaceAll", "please use `regexReplaceAll` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexReplaceAllLiteral", "please use `regexReplaceAllLiteral` instead"))
	return nil
}
