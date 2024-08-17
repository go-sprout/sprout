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
	return "slices"
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
	sprout.AddFunction(funcsMap, "mustAppend", sr.MustAppend)
	sprout.AddFunction(funcsMap, "mustPrepend", sr.MustPrepend)
	sprout.AddFunction(funcsMap, "mustChunk", sr.MustChunk)
	sprout.AddFunction(funcsMap, "mustUniq", sr.MustUniq)
	sprout.AddFunction(funcsMap, "mustCompact", sr.MustCompact)
	sprout.AddFunction(funcsMap, "mustSlice", sr.MustSlice)
	sprout.AddFunction(funcsMap, "mustHas", sr.MustHas)
	sprout.AddFunction(funcsMap, "mustWithout", sr.MustWithout)
	sprout.AddFunction(funcsMap, "mustRest", sr.MustRest)
	sprout.AddFunction(funcsMap, "mustInitial", sr.MustInitial)
	sprout.AddFunction(funcsMap, "mustFirst", sr.MustFirst)
	sprout.AddFunction(funcsMap, "mustLast", sr.MustLast)
	sprout.AddFunction(funcsMap, "mustReverse", sr.MustReverse)
	return nil
}
