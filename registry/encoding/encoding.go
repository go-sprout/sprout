package encoding

import "github.com/go-sprout/sprout"

type EncodingRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *EncodingRegistry {
	return &EncodingRegistry{}
}

// Uid returns the unique identifier of the registry.
func (or *EncodingRegistry) Uid() string {
	return "encoding"
}

// LinkHandler links the handler to the registry at runtime.
func (or *EncodingRegistry) LinkHandler(fh sprout.Handler) error {
	or.handler = fh
	return nil
}

func (er *EncodingRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "base64Encode", er.Base64Encode)
	sprout.AddFunction(funcsMap, "base64Decode", er.Base64Decode)
	sprout.AddFunction(funcsMap, "base32Encode", er.Base32Encode)
	sprout.AddFunction(funcsMap, "base32Decode", er.Base32Decode)
	sprout.AddFunction(funcsMap, "fromJson", er.FromJson)
	sprout.AddFunction(funcsMap, "toJson", er.ToJson)
	sprout.AddFunction(funcsMap, "toPrettyJson", er.ToPrettyJson)
	sprout.AddFunction(funcsMap, "toRawJson", er.ToRawJson)
	sprout.AddFunction(funcsMap, "fromYaml", er.FromYAML)
	sprout.AddFunction(funcsMap, "toYaml", er.ToYAML)
	sprout.AddFunction(funcsMap, "mustFromJson", er.MustFromJson)
	sprout.AddFunction(funcsMap, "mustToJson", er.MustToJson)
	sprout.AddFunction(funcsMap, "mustToPrettyJson", er.MustToPrettyJson)
	sprout.AddFunction(funcsMap, "mustToRawJson", er.MustToRawJson)
	sprout.AddFunction(funcsMap, "mustFromYaml", er.MustFromYAML)
	sprout.AddFunction(funcsMap, "mustToYaml", er.MustToYAML)
	return nil
}
