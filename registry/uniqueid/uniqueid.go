package uniqueid

import "github.com/go-sprout/sprout"

type UniqueIDRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *UniqueIDRegistry {
	return &UniqueIDRegistry{}
}

// UID returns the unique identifier of the registry.
func (ur *UniqueIDRegistry) UID() string {
	return "go-sprout/sprout.uniqueid"
}

// LinkHandler links the handler to the registry at runtime.
func (ur *UniqueIDRegistry) LinkHandler(fh sprout.Handler) error {
	ur.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (ur *UniqueIDRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "uuidv4", ur.Uuidv4)
	sprout.AddFunction(funcsMap, "uuidv7", ur.Uuidv7)
	sprout.AddFunction(funcsMap, "uuidv5", ur.Uuidv5)
	sprout.AddFunction(funcsMap, "uuidv3", ur.Uuidv3)
	sprout.AddFunction(funcsMap, "uuidNil", ur.UuidNil)
	sprout.AddFunction(funcsMap, "isUUID", ur.IsUUID)
	sprout.AddFunction(funcsMap, "uuidVersion", ur.UuidVersion)
	sprout.AddFunction(funcsMap, "uuidTime", ur.UuidTime)
	return nil
}
