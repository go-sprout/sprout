package reflect

import "github.com/go-sprout/sprout"

type ReflectRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of reflect registry.
func NewRegistry() *ReflectRegistry {
	return &ReflectRegistry{}
}

// Uid returns the unique identifier of the registry.
func (rr *ReflectRegistry) Uid() string {
	return "reflect"
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
	sprout.AddFunction(funcsMap, "deepEqual", rr.DeepEqual)
	sprout.AddFunction(funcsMap, "deepCopy", rr.DeepCopy)
	sprout.AddFunction(funcsMap, "mustDeepCopy", rr.MustDeepCopy)
	sprout.AddFunction(funcsMap, "hasField", rr.HasField)
	return nil
}
