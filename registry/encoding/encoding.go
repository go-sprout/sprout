package encoding

import "github.com/go-sprout/sprout"

type EncodingRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of conversion registry.
func NewRegistry() *EncodingRegistry {
	return &EncodingRegistry{}
}

// UID returns the unique identifier of the registry.
func (or *EncodingRegistry) UID() string {
	return "go-sprout/sprout.encoding"
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
	return nil
}

func (er *EncodingRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "fromJson", "mustFromJson")
	sprout.AddAlias(aliasesMap, "toJson", "mustToJson")
	sprout.AddAlias(aliasesMap, "toPrettyJson", "mustToPrettyJson")
	sprout.AddAlias(aliasesMap, "toRawJson", "mustToRawJson")
	sprout.AddAlias(aliasesMap, "fromYaml", "mustFromYaml")
	sprout.AddAlias(aliasesMap, "toYaml", "mustToYaml")
	return nil
}

func (er *EncodingRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustFromJson", "please use `fromJson` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToJson", "please use `toJson` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToPrettyJson", "please use `toPrettyJson` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToRawJson", "please use `toRawJson` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustFromYaml", "please use `fromYaml` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToYaml", "please use `toYaml` instead"))
	return nil
}
