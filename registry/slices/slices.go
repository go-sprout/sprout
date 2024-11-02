package slices

import "github.com/go-sprout/sprout"

type SlicesRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *SlicesRegistry {
	return &SlicesRegistry{}
}

// Uid returns the unique identifier of the registry.
func (sr *SlicesRegistry) Uid() string {
	return "go-sprout/sprout.slices"
}

// LinkHandler links the handler to the registry at runtime.
func (sr *SlicesRegistry) LinkHandler(fh sprout.Handler) error {
	sr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (sr *SlicesRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "list", sr.List)
	sprout.AddFunction(funcsMap, "append", sr.Append)
	sprout.AddFunction(funcsMap, "prepend", sr.Prepend)
	sprout.AddFunction(funcsMap, "concat", sr.Concat)
	sprout.AddFunction(funcsMap, "chunk", sr.Chunk)
	sprout.AddFunction(funcsMap, "uniq", sr.Uniq)
	sprout.AddFunction(funcsMap, "compact", sr.Compact)
	sprout.AddFunction(funcsMap, "flatten", sr.Flatten)
	sprout.AddFunction(funcsMap, "flattenDepth", sr.FlattenDepth)
	sprout.AddFunction(funcsMap, "slice", sr.Slice)
	sprout.AddFunction(funcsMap, "has", sr.Has)
	sprout.AddFunction(funcsMap, "without", sr.Without)
	sprout.AddFunction(funcsMap, "rest", sr.Rest)
	sprout.AddFunction(funcsMap, "initial", sr.Initial)
	sprout.AddFunction(funcsMap, "first", sr.First)
	sprout.AddFunction(funcsMap, "last", sr.Last)
	sprout.AddFunction(funcsMap, "reverse", sr.Reverse)
	sprout.AddFunction(funcsMap, "sortAlpha", sr.SortAlpha)
	sprout.AddFunction(funcsMap, "splitList", sr.SplitList)
	sprout.AddFunction(funcsMap, "strSlice", sr.StrSlice)
	sprout.AddFunction(funcsMap, "until", sr.Until)
	sprout.AddFunction(funcsMap, "untilStep", sr.UntilStep)
	return nil
}

func (sr *SlicesRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "append", "mustAppend")
	sprout.AddAlias(aliasesMap, "prepend", "mustPrepend")
	sprout.AddAlias(aliasesMap, "chunk", "mustChunk")
	sprout.AddAlias(aliasesMap, "uniq", "mustUniq")
	sprout.AddAlias(aliasesMap, "compact", "mustCompact")
	sprout.AddAlias(aliasesMap, "slice", "mustSlice")
	sprout.AddAlias(aliasesMap, "has", "mustHas")
	sprout.AddAlias(aliasesMap, "without", "mustWithout")
	sprout.AddAlias(aliasesMap, "rest", "mustRest")
	sprout.AddAlias(aliasesMap, "initial", "mustInitial")
	sprout.AddAlias(aliasesMap, "first", "mustFirst")
	sprout.AddAlias(aliasesMap, "last", "mustLast")
	sprout.AddAlias(aliasesMap, "reverse", "mustReverse")
	return nil
}

func (sr *SlicesRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustAppend", "please use `append` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustPrepend", "please use `prepend` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustChunk", "please use `chunk` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustUniq", "please use `uniq` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustCompact", "please use `compact` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustSlice", "please use `slice` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustHas", "please use `has` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustWithout", "please use `without` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustRest", "please use `rest` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustInitial", "please use `initial` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustFirst", "please use `first` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustLast", "please use `last` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustReverse", "please use `reverse` instead"))
	return nil
}
