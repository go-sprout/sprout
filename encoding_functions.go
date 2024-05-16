package sprout

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"strings"

	"sigs.k8s.io/yaml"
)

// Base64Encode encodes a string into its Base64 representation.
//
// Parameters:
//
//	s string - the string to encode.
//
// Returns:
//
//	string - the Base64 encoded string.
//
// Example:
//
//	{{ "Hello World" | base64Encode }} // Output: "SGVsbG8gV29ybGQ="
func (fh *FunctionHandler) Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode decodes a Base64 encoded string back to its original form.
// Returns an error message if the input is not valid Base64.
//
// Parameters:
//
//	s string - the Base64 encoded string to decode.
//
// Returns:
//
//	string - the decoded string, or an error message if the decoding fails.
//
// Example:
//
//	{{ "SGVsbG8gV29ybGQ=" | base64Decode }} // Output: "Hello World"
func (fh *FunctionHandler) Base64Decode(s string) string {
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// Base32Encode encodes a string into its Base32 representation.
//
// Parameters:
//
//	s string - the string to encode.
//
// Returns:
//
//	string - the Base32 encoded string.
//
// Example:
//
//	{{ "Hello World" | base32Encode }} // Output: "JBSWY3DPEBLW64TMMQQQ===="
func (fh *FunctionHandler) Base32Encode(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

// Base32Decode decodes a Base32 encoded string back to its original form.
// Returns an error message if the input is not valid Base32.
//
// Parameters:
//
//	s string - the Base32 encoded string to decode.
//
// Returns:
//
//	string - the decoded string, or an error message if the decoding fails.
//
// Example:
//
//	{{ "JBSWY3DPEBLW64TMMQQQ====" | base32Decode }} // Output: "Hello World"
func (fh *FunctionHandler) Base32Decode(s string) string {
	bytes, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// FromJson converts a JSON string into a corresponding Go data structure.
//
// Parameters:
//
//	v string - the JSON string to decode.
//
// Returns:
//
//	any - the decoded Go data structure, or nil if the decoding fails.
//
// Example:
//
//	result := fh.FromJson(`{"name":"John", "age":30}`)
//	fmt.Printf("%v\n", result) // Output: map[name:John age:30]
func (fh *FunctionHandler) FromJson(v string) any {
	output, _ := fh.MustFromJson(v)
	return output
}

// ToJson converts a Go data structure into a JSON string.
//
// Parameters:
//
//	v any - the Go data structure to encode.
//
// Returns:
//
//	string - the encoded JSON string.
//
// Example:
//
//	jsonStr := fh.ToJson(map[string]any{"name": "John", "age": 30})
//	fmt.Println(jsonStr) // Output: {"age":30,"name":"John"}
func (fh *FunctionHandler) ToJson(v any) string {
	output, _ := fh.MustToJson(v)
	return output
}

// ToPrettyJson converts a Go data structure into a pretty-printed JSON string.
//
// Parameters:
//
//	v any - the Go data structure to encode.
//
// Returns:
//
//	string - the pretty-printed JSON string.
//
// Example:
//
//	prettyJson := fh.ToPrettyJson(map[string]any{"name": "John", "age": 30})
//	fmt.Println(prettyJson) // Output: {
//	                        //   "age": 30,
//	                        //   "name": "John"
//	                        // }
func (fh *FunctionHandler) ToPrettyJson(v any) string {
	output, _ := fh.MustToPrettyJson(v)
	return output
}

// ToRawJson converts a Go data structure into a JSON string without escaping HTML.
//
// Parameters:
//
//	v any - the Go data structure to encode.
//
// Returns:
//
//	string - the raw JSON string.
//
// Example:
//
//	rawJson := fh.ToRawJson(map[string]any{"content": "<div>Hello World!</div>"})
//	fmt.Println(rawJson) // Output: {"content":"<div>Hello World!</div>"}
func (fh *FunctionHandler) ToRawJson(v any) string {
	output, _ := fh.MustToRawJson(v)
	return output
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
//
// Example:
//
//	{{ "name: John Doe\nage: 30" | fromYAML }} // Output: map[name:John Doe age:30]
func (fh *FunctionHandler) FromYAML(str string) any {
	var m = make(map[string]any)

	if err := yaml.Unmarshal([]byte(str), &m); err != nil {
		return nil
	}

	return m
}

// ToYAML serializes a Go data structure to a YAML string.
//
// Parameters:
//
//	v any - the data structure to serialize.
//
// Returns:
//
//	string - the YAML string representation of the data structure.
//
// Example:
//
//	{{ {"name": "John Doe", "age": 30} | toYAML }} // Output: "name: John Doe\nage: 30\n"
func (fh *FunctionHandler) ToYAML(v any) string {
	result, _ := fh.MustToYAML(v)
	return result
}

// MustFromJson decodes a JSON string into a Go data structure, returning an
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
//	{{ `{"name":"John", "age":30}` | mustFromJson }} // Output: map[name:John age:30], nil
func (fh *FunctionHandler) MustFromJson(v string) (any, error) {
	var output any
	err := json.Unmarshal([]byte(v), &output)
	return output, err
}

// MustToJson encodes a Go data structure into a JSON string, returning an error
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
//	{{ {"name": "John", "age": 30} | mustToJson }} // Output: "{"age":30,"name":"John"}", nil
func (fh *FunctionHandler) MustToJson(v any) (string, error) {
	output, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// MustToPrettyJson encodes a Go data structure into a pretty-printed JSON
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
//	{{ {"name": "John", "age": 30} | mustToPrettyJson }} // Output: "{\n  \"age\": 30,\n  \"name\": \"John\"\n}", nil
func (fh *FunctionHandler) MustToPrettyJson(v any) (string, error) {
	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// MustToRawJson encodes a Go data structure into a JSON string without escaping
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
//	{{ {"content": "<div>Hello World!</div>"} | mustToRawJson }} // Output: "{\"content\":\"<div>Hello World!</div>\"}", nil
func (fh *FunctionHandler) MustToRawJson(v any) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&v)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

// MustFromYaml deserializes a YAML string into a Go data structure, returning
// the result along with any error that occurs.
//
// Parameters:
//
//	v string - the YAML string to deserialize.
//
// Returns:
//
//	any - the Go data structure representing the deserialized YAML content.
//	error - an error if the YAML content cannot be deserialized.
//
// Example:
//
//	{{ "name: John Doe\nage: 30" | mustFromYaml }} // Output: map[name:John Doe age:30], nil
func (fh *FunctionHandler) MustFromYAML(v string) (any, error) {
	var output any
	err := yaml.Unmarshal([]byte(v), &output)
	return output, err
}

// MustToYAML serializes a Go data structure to a YAML string and returns any error that occurs during the serialization.
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
//	{{ {"name": "John Doe", "age": 30} | mustToYAML }} // Output: "name: John Doe\nage: 30\n", nil
func (fh *FunctionHandler) MustToYAML(v any) (string, error) {
	data, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}
