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
	sprout.AddFunction(funcsMap, "fromJSON", er.FromJson)
	sprout.AddFunction(funcsMap, "toJSON", er.ToJson)
	sprout.AddFunction(funcsMap, "toPrettyJSON", er.ToPrettyJson)
	sprout.AddFunction(funcsMap, "toRawJSON", er.ToRawJson)
	sprout.AddFunction(funcsMap, "fromYAML", er.FromYAML)
	sprout.AddFunction(funcsMap, "toYAML", er.ToYAML)
	sprout.AddFunction(funcsMap, "toIndentYAML", er.ToIndentYAML)
	return nil
}

func (er *EncodingRegistry) RegisterAliases(aliasesMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasesMap, "fromJSON", "fromJson", "mustFromJson")
	sprout.AddAlias(aliasesMap, "toJSON", "toJson", "mustToJson")
	sprout.AddAlias(aliasesMap, "toPrettyJSON", "toPrettyJson", "mustToPrettyJson")
	sprout.AddAlias(aliasesMap, "toRawJSON", "toRawJson", "mustToRawJson")
	sprout.AddAlias(aliasesMap, "fromYAML", "fromYaml", "mustFromYaml")
	sprout.AddAlias(aliasesMap, "toYAML", "toYaml", "mustToYaml")
	sprout.AddAlias(aliasesMap, "toIndentYAML", "toIndentYaml")
	return nil
}

func (er *EncodingRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustFromJson", "please use `fromJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToJson", "please use `toJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToPrettyJson", "please use `toPrettyJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToRawJson", "please use `toRawJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustFromYaml", "please use `fromYAML` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("mustToYaml", "please use `toYAML` instead"))

	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("fromJson", "please use `fromJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("toJson", "please use `toJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("toPrettyJson", "please use `toPrettyJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("toRawJson", "please use `toRawJSON` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("fromYaml", "please use `fromYAML` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("toYaml", "please use `toYAML` instead"))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("toIndentYaml", "please use `toIndentYAML` instead"))

	return nil
}
