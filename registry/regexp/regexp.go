// Package regexp provides regular expression functions using the historical
// sprig signatures, where the string to work on is not always the last
// parameter.
//
// Deprecated: use [github.com/go-sprout/sprout/registry/regex] instead, where
// the main parameter is always the last one and every function can be used in
// a template pipeline. This registry is kept for backward compatibility and
// will be removed in Sprout v1.2.
package regexp

import "github.com/go-sprout/sprout"

type RegexpRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of regexp registry.
//
// Deprecated: use [github.com/go-sprout/sprout/registry/regex] NewRegistry
// instead. Both registries expose the same function names, so they are mutually
// exclusive: register one or the other, never both.
func NewRegistry() *RegexpRegistry {
	return &RegexpRegistry{}
}

// UID returns the unique identifier of the registry.
func (rr *RegexpRegistry) UID() string {
	return "go-sprout/sprout.regexp"
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
	sprout.AddFunction(funcsMap, "regexFindGroups", rr.RegexFindGroups)
	sprout.AddFunction(funcsMap, "regexFindAllGroups", rr.RegexFindAllGroups)
	sprout.AddFunction(funcsMap, "regexFindNamed", rr.RegexFindNamed)
	sprout.AddFunction(funcsMap, "regexFindAllNamed", rr.RegexFindAllNamed)
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
	// The whole registry is deprecated in favor of the `regex` one, where the
	// main parameter is always the last one. Functions whose signature is left
	// untouched only need the registry to be swapped, the others also need the
	// arguments to be reordered in the templates.
	//
	// These notices are informational, not deprecation warnings: choosing a
	// registry is the library author's call, not the template author's, and the
	// latter cannot act on a warning raised while rendering. Library authors are
	// reached by the `Deprecated:` markers on the package and on NewRegistry.
	sprout.AddNotice(notices, sprout.NewNotice(sprout.NoticeKindInfo, []string{
		"regexFind",
		"regexMatch",
		"regexQuoteMeta",
		"regexFindGroups",
		"regexFindAllGroups",
		"regexFindNamed",
		"regexFindAllNamed",
	}, "the `regexp` registry is deprecated in favor of `regex` and will be removed in v1.2, this function keeps the same signature"))
	sprout.AddNotice(notices, sprout.NewInfoNotice("regexFindAll", "the `regexp` registry is deprecated in favor of `regex` and will be removed in v1.2, where the signature becomes `regexFindAll <regex> <n> <value>`"))
	sprout.AddNotice(notices, sprout.NewInfoNotice("regexSplit", "the `regexp` registry is deprecated in favor of `regex` and will be removed in v1.2, where the signature becomes `regexSplit <regex> <n> <value>`"))
	sprout.AddNotice(notices, sprout.NewInfoNotice("regexReplaceAll", "the `regexp` registry is deprecated in favor of `regex` and will be removed in v1.2, where the signature becomes `regexReplaceAll <regex> <replacedBy> <value>`"))
	sprout.AddNotice(notices, sprout.NewInfoNotice("regexReplaceAllLiteral", "the `regexp` registry is deprecated in favor of `regex` and will be removed in v1.2, where the signature becomes `regexReplaceAllLiteral <regex> <replacedBy> <value>`"))

	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexFind", "please use `regexFind` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexFindAll", "please use `regexFindAll` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexMatch", "please use `regexMatch` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexSplit", "please use `regexSplit` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexReplaceAll", "please use `regexReplaceAll` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRegexReplaceAllLiteral", "please use `regexReplaceAllLiteral` instead"))
	return nil
}
