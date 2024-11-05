package reflect

import "github.com/go-sprout/sprout"

type ReflectRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of reflect registry.
func NewRegistry() *ReflectRegistry {
	return &ReflectRegistry{}
}

// UID returns the unique identifier of the registry.
func (rr *ReflectRegistry) UID() string {
	return "go-sprout/sprout.reflect"
}

// LinkHandler links the handler to the registry at runtime.
func (rr *ReflectRegistry) LinkHandler(fh sprout.Handler) error {
	rr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (rr *ReflectRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "typeIs", rr.TypeIs)
	sprout.AddFunction(funcsMap, "typeIsLike", rr.TypeIsLike)
	sprout.AddFunction(funcsMap, "typeOf", rr.TypeOf)
	sprout.AddFunction(funcsMap, "kindIs", rr.KindIs)
	sprout.AddFunction(funcsMap, "kindOf", rr.KindOf)
	sprout.AddFunction(funcsMap, "hasField", rr.HasField)
	sprout.AddFunction(funcsMap, "deepEqual", rr.DeepEqual)
	sprout.AddFunction(funcsMap, "deepCopy", rr.DeepCopy)
	return nil
}

func (rr *ReflectRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "deepCopy", "mustDeepCopy")
	return nil
}

func (rr *ReflectRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustDeepCopy", "please use `deepCopy` instead"))
	return nil
}
