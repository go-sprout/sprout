package encoding

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-sprout/sprout"
	"gopkg.in/yaml.v3"
)

// Base64Encode encodes a string into its Base64 representation.
//
// Parameters:
//
//	str string - the string to encode.
//
// Returns:
//
//	string - the Base64 encoded string.
//
// Example:
//
//	{{ "Hello World" | base64Encode }} // Output: "SGVsbG8gV29ybGQ="
func (er *EncodingRegistry) Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode decodes a Base64 encoded string back to its original form.
// Returns an error message if the input is not valid Base64.
//
// Parameters:
//
//	str string - the Base64 encoded string to decode.
//
// Returns:
//
//	string - the decoded string, or an error message if the decoding fails.
//	error - an error message if the decoding fails.
//
// Example:
//
//	{{ "SGVsbG8gV29ybGQ=" | base64Decode }} // Output: "Hello World"
func (er *EncodingRegistry) Base64Decode(str string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %v", err)
	}
	return string(bytes), nil
}

// Base32Encode encodes a string into its Base32 representation.
//
// Parameters:
//
//	str string - the string to encode.
//
// Returns:
//
//	string - the Base32 encoded string.
//
// Example:
//
//	{{ "Hello World" | base32Encode }} // Output: "JBSWY3DPEBLW64TMMQQQ===="
func (er *EncodingRegistry) Base32Encode(str string) string {
	return base32.StdEncoding.EncodeToString([]byte(str))
}

// Base32Decode decodes a Base32 encoded string back to its original form.
// Returns an error message if the input is not valid Base32.
//
// Parameters:
//
//	str string - the Base32 encoded string to decode.
//
// Returns:
//
//	string - the decoded string, or an error message if the decoding fails.
//	error - an error message if the decoding fails.
//
// Example:
//
//	{{ "JBSWY3DPEBLW64TMMQQQ====" | base32Decode }} // Output: "Hello World"
func (er *EncodingRegistry) Base32Decode(str string) (string, error) {
	bytes, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		return "", fmt.Errorf("base32 decode error: %v", err)
	}
	return string(bytes), nil
}

// FromJson decodes a JSON string into a Go data structure, returning an
// error if decoding fails.
//
// Parameters:
//
//	v string - the JSON string to decode.
//
// Returns:
//
//	any - the decoded Go data structure.
//	error - error encountered during decoding, if any.
//
// Example:
//
//	{{ `{"name":"John", "age":30}` | fromJson }} // Output: map[name:John age:30], nil
func (er *EncodingRegistry) FromJson(v string) (any, error) {
	var output any
	err := json.Unmarshal([]byte(v), &output)
	if err != nil {
		return nil, fmt.Errorf("json decode error: %v", err)
	}
	return output, err
}

// ToJson encodes a Go data structure into a JSON string, returning an error
// if encoding fails.
//
// Parameters:
//
//	v any - the Go data structure to encode.
//
// Returns:
//
//	string - the JSON-encoded string.
//	error - error encountered during encoding, if any.
//
// Example:
//
//	{{ {"name": "John", "age": 30} | toJson }} // Output: "{"age":30,"name":"John"}", nil
func (er *EncodingRegistry) ToJson(v any) (string, error) {
	output, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("json encode error: %v", err)
	}
	return string(output), nil
}

// ToPrettyJson encodes a Go data structure into a pretty-printed JSON
// string, returning an error if encoding fails.
//
// Parameters:
//
//	v any - the Go data structure to encode.
//
// Returns:
//
//	string - the pretty-printed JSON string.
//	error - error encountered during encoding, if any.
//
// Example:
//
//	{{ {"name": "John", "age": 30} | toPrettyJson }} // Output: "{\n  \"age\": 30,\n  \"name\": \"John\"\n}", nil
func (er *EncodingRegistry) ToPrettyJson(v any) (string, error) {
	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json encode error: %v", err)
	}
	return string(output), nil
}

// ToRawJson encodes a Go data structure into a JSON string without escaping
// HTML, returning an error if encoding fails.
//
// Parameters:
//
//	v any - the Go data structure to encode.
//
// Returns:
//
//	string - the raw JSON string.
//	error - error encountered during encoding, if any.
//
// Example:
//
//	{{ {"content": "<div>Hello World!</div>"} | toRawJson }} // Output: "{\"content\":\"<div>Hello World!</div>\"}", nil
func (er *EncodingRegistry) ToRawJson(v any) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&v)
	if err != nil {
		return "", fmt.Errorf("json encode error: %v", err)
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

// FromYAML deserializes a YAML string into a Go map.
//
// Parameters:
//
//	str string - the YAML string to deserialize.
//
// Returns:
//
//	any - a map representing the YAML data. Returns nil if deserialization fails.
//	error - an error message if the YAML content cannot be deserialized.
//
// Example:
//
//	{{ "name: John Doe\nage: 30" | fromYaml }} // Output: map[name:John Doe age:30]
func (er *EncodingRegistry) FromYAML(str string) (any, error) {
	m := make(map[string]any)

	if err := yaml.Unmarshal([]byte(str), &m); err != nil {
		return nil, fmt.Errorf("yaml decode error: %v", err)
	}

	return m, nil
}

// ToYAML serializes a Go data structure to a YAML string and returns any error
// that occurs during the serialization.
//
// Parameters:
//
//	v any - the data structure to serialize.
//
// Returns:
//
//	string - the YAML string representation of the data structure.
//	error - error if the serialization fails.
//
// Example:
//
//	{{ $d := dict "name" "John Doe" "age" 30 }}
//	{{ $d | toYaml }} // Output: name: John Doe\nage: 30
func (er *EncodingRegistry) ToYAML(v any) (out string, err error) {
	// recover panic from yaml package
	defer sprout.ErrRecoverPanic(&err, "yaml encode error")

	data, err := yaml.Marshal(v)
	if err != nil {
		// code unreachable because yaml.Marshal always panic on error and never
		// returns an error, but we still need to handle the error for the sake of
		// consistency. The error message is set by ErrRecoverPanic.
		return "", fmt.Errorf("yaml encode error: %v", err)
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}
