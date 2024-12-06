package encoding

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/go-sprout/sprout"
)

// Base64Encode encodes a string into its Base64 representation.
//
// Parameters:
//
//	value string - the string to encode.
//
// Returns:
//
//	string - the Base64 encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: base64Encode].
//
// [Sprout Documentation: base64Encode]: https://docs.atom.codes/sprout/registries/encoding#base64encode
func (er *EncodingRegistry) Base64Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// Base64Decode decodes a Base64 encoded string back to its original form.
// Returns an error message if the input is not valid Base64.
//
// Parameters:
//
//	value string - the Base64 encoded string to decode.
//
// Returns:
//
//	string - the decoded string, or an error message if the decoding fails.
//	error - an error message if the decoding fails.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: base64Decode].
//
// [Sprout Documentation: base64Decode]: https://docs.atom.codes/sprout/registries/encoding#base64decode
func (er *EncodingRegistry) Base64Decode(value string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %w", err)
	}
	return string(bytes), nil
}

// Base32Encode encodes a string into its Base32 representation.
//
// Parameters:
//
//	value string - the string to encode.
//
// Returns:
//
//	string - the Base32 encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: base32Encode].
//
// [Sprout Documentation: base32Encode]: https://docs.atom.codes/sprout/registries/encoding#base32encode
func (er *EncodingRegistry) Base32Encode(value string) string {
	return base32.StdEncoding.EncodeToString([]byte(value))
}

// Base32Decode decodes a Base32 encoded string back to its original form.
// Returns an error message if the input is not valid Base32.
//
// Parameters:
//
//	value string - the Base32 encoded string to decode.
//
// Returns:
//
//	string - the decoded string, or an error message if the decoding fails.
//	error - an error message if the decoding fails.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: base32Decode].
//
// [Sprout Documentation: base32Decode]: https://docs.atom.codes/sprout/registries/encoding#base32decode
func (er *EncodingRegistry) Base32Decode(value string) (string, error) {
	bytes, err := base32.StdEncoding.DecodeString(value)
	if err != nil {
		return "", fmt.Errorf("base32 decode error: %w", err)
	}
	return string(bytes), nil
}

// FromJson decodes a JSON string into a Go data structure, returning an
// error if decoding fails.
//
// Parameters:
//
//	value string - the JSON string to decode.
//
// Returns:
//
//	any - the decoded Go data structure.
//	error - error encountered during decoding, if any.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: fromJson].
//
// [Sprout Documentation: fromJson]: https://docs.atom.codes/sprout/registries/encoding#fromjson
func (er *EncodingRegistry) FromJson(value string) (any, error) {
	var output any
	err := json.Unmarshal([]byte(value), &output)
	if err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	return output, err
}

// ToJson encodes a Go data structure into a JSON string, returning an error
// if encoding fails.
//
// Parameters:
//
//	value any - the Go data structure to encode.
//
// Returns:
//
//	string - the JSON-encoded string.
//	error - error encountered during encoding, if any.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toJson].
//
// [Sprout Documentation: toJson]: https://docs.atom.codes/sprout/registries/encoding#tojson
func (er *EncodingRegistry) ToJson(value any) (string, error) {
	output, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("json encode error: %w", err)
	}
	return string(output), nil
}

// ToPrettyJson encodes a Go data structure into a pretty-printed JSON
// string, returning an error if encoding fails.
//
// Parameters:
//
//	value any - the Go data structure to encode.
//
// Returns:
//
//	string - the pretty-printed JSON string.
//	error - error encountered during encoding, if any.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toPrettyJson].
//
// [Sprout Documentation: toPrettyJson]: https://docs.atom.codes/sprout/registries/encoding#toprettyjson
func (er *EncodingRegistry) ToPrettyJson(value any) (string, error) {
	output, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json encode error: %w", err)
	}
	return string(output), nil
}

// ToRawJson encodes a Go data structure into a JSON string without escaping
// HTML, returning an error if encoding fails.
//
// Parameters:
//
//	value any - the Go data structure to encode.
//
// Returns:
//
//	string - the raw JSON string.
//	error - error encountered during encoding, if any.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toRawJson].
//
// [Sprout Documentation: toRawJson]: https://docs.atom.codes/sprout/registries/encoding#torawjson
func (er *EncodingRegistry) ToRawJson(value any) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&value)
	if err != nil {
		return "", fmt.Errorf("json encode error: %w", err)
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

// FromYAML deserializes a YAML string into a Go map.
//
// Parameters:
//
//	value string - the YAML string to deserialize.
//
// Returns:
//
//	any - a map representing the YAML data. Returns nil if deserialization fails.
//	error - an error message if the YAML content cannot be deserialized.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: fromYaml].
//
// [Sprout Documentation: fromYaml]: https://docs.atom.codes/sprout/registries/encoding#fromyaml
func (er *EncodingRegistry) FromYAML(value string) (any, error) {
	m := make(map[string]any)

	if err := yaml.Unmarshal([]byte(value), &m); err != nil {
		return nil, fmt.Errorf("yaml decode error: %w", err)
	}

	return m, nil
}

// ToYAML serializes a Go data structure to a YAML string and returns any error
// that occurs during the serialization.
//
// Parameters:
//
//	value any - the data structure to serialize.
//
// Returns:
//
//	string - the YAML string representation of the data structure.
//	error - error if the serialization fails.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toYaml].
//
// [Sprout Documentation: toYaml]: https://docs.atom.codes/sprout/registries/encoding#toyaml
func (er *EncodingRegistry) ToYAML(value any) (out string, err error) {
	return er.ToIndentYAML(4, value)
}

// ToIndentYAML serializes a Go data structure to a YAML string and returns any error
// that occurs during the serialization. It allows to set an indentation width.
//
// Parameters:
//
//	value any      - the data structure to serialize.
//	indent int - the indentation
//	omitempty bool - omit empty fields (default: false)
//
// Returns:
//
//	string - the YAML string representation of the data structure.
//	error - error if the serialization fails.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: toIndentYaml].
//
// [Sprout Documentation: toIndentYaml]: https://docs.atom.codes/sprout/registries/encoding#toindentyaml
func (er *EncodingRegistry) ToIndentYAML(indent int, value any) (out string, err error) {
	// recover panic from yaml package
	defer sprout.ErrRecoverPanic(&err, "yaml encode error")

	buf := bytes.Buffer{}
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(indent)

	if err = enc.Encode(&value); err != nil {
		// code unreachable because yaml.Marshal always panic on error and never
		// returns an error, but we still need to handle the error for the sake of
		// consistency. The error message is set by ErrRecoverPanic.
		return "", fmt.Errorf("YAML encode error: %w", err)
	}

	return strings.TrimSuffix(buf.String(), "\n"), nil
}
