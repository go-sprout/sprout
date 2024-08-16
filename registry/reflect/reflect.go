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
func (nr *ReflectRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "typeIs", nr.TypeIs)
	sprout.AddFunction(funcsMap, "typeIsLike", nr.TypeIsLike)
	sprout.AddFunction(funcsMap, "typeOf", nr.TypeOf)
	sprout.AddFunction(funcsMap, "kindIs", nr.KindIs)
	sprout.AddFunction(funcsMap, "kindOf", nr.KindOf)
	sprout.AddFunction(funcsMap, "deepEqual", nr.DeepEqual)
	sprout.AddFunction(funcsMap, "deepCopy", nr.DeepCopy)
	sprout.AddFunction(funcsMap, "mustDeepCopy", nr.MustDeepCopy)
	return nil
}
