package encoding_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/encoding"
)

func TestBase64Encode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base64Encode }}`, ExpectedOutput: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "Hello World" | base64Encode }}`, ExpectedOutput: "SGVsbG8gV29ybGQ="},
		{Name: "TestFromVariableInput", Input: `{{ .V | base64Encode }}`, ExpectedOutput: "SGVsbG8gV29ybGQ=", Data: map[string]any{"V": "Hello World"}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestBase64Decode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base64Decode }}`, ExpectedOutput: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "SGVsbG8gV29ybGQ=" | base64Decode }}`, ExpectedOutput: "Hello World"},
		{Name: "TestFromVariableInput", Input: `{{ .V | base64Decode }}`, ExpectedOutput: "Hello World", Data: map[string]any{"V": "SGVsbG8gV29ybGQ="}},
		{Name: "TestInvalidInput", Input: `{{ "SGVsbG8gV29ybGQ" | base64Decode }}`, ExpectedErr: "base64 decode error"},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestBase32Encode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base32Encode }}`, ExpectedOutput: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "Hello World" | base32Encode }}`, ExpectedOutput: "JBSWY3DPEBLW64TMMQ======"},
		{Name: "TestFromVariableInput", Input: `{{ .V | base32Encode }}`, ExpectedOutput: "JBSWY3DPEBLW64TMMQ======", Data: map[string]any{"V": "Hello World"}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestBase32Decode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base32Decode }}`, ExpectedOutput: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "JBSWY3DPEBLW64TMMQ======" | base32Decode }}`, ExpectedOutput: "Hello World"},
		{Name: "TestFromVariableInput", Input: `{{ .V | base32Decode }}`, ExpectedOutput: "Hello World", Data: map[string]any{"V": "JBSWY3DPEBLW64TMMQ======"}},
		{Name: "TestInvalidInput", Input: `{{ "JBSWY3DPEBLW64TMMQ" | base32Decode }}`, ExpectedErr: "base32 decode error"},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestFromJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | fromJson }}`, ExpectedErr: "json decode error"},
		{Name: "TestVariableInput", Input: `{{ .V | fromJson }}`, ExpectedOutput: "map[foo:55]", Data: map[string]any{"V": `{"foo": 55}`}},
		{Name: "TestAccessField", Input: `{{ (.V | fromJson).foo }}`, ExpectedOutput: "55", Data: map[string]any{"V": `{"foo": 55}`}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toJson }}`, ExpectedOutput: `""`},
		{Name: "TestVariableInput", Input: `{{ .V | toJson }}`, ExpectedOutput: "{\"bar\":\"baz\",\"foo\":55}", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
		{Name: "TestInvalidInput", Input: `{{ .V | toJson }}`, ExpectedErr: "json encode error", Data: map[string]any{"V": make(chan int)}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToPrettyJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toPrettyJson }}`, ExpectedOutput: `""`},
		{Name: "TestVariableInput", Input: `{{ .V | toPrettyJson }}`, ExpectedOutput: "{\n  \"bar\": \"baz\",\n  \"foo\": 55\n}", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToRawJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toRawJson }}`, ExpectedOutput: `""`},
		{Name: "TestVariableInput", Input: `{{ .V | toRawJson }}`, ExpectedOutput: "{\"bar\":\"baz\",\"foo\":55}", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestFromYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | fromYaml }}`, ExpectedOutput: "map[]"},
		{Name: "TestVariableInput", Input: `{{ .V | fromYaml }}`, ExpectedOutput: "map[bar:map[baz:1] foo:55]", Data: map[string]any{"V": "foo: 55\nbar:\n  baz: 1\n"}},
		{Name: "TestAccessField", Input: `{{ (.V | fromYaml).foo }}`, ExpectedOutput: "55", Data: map[string]any{"V": "foo: 55"}},
		{Name: "TestInvalidInput", Input: "{{ .V | fromYaml }}", ExpectedErr: "yaml decode error", Data: map[string]any{"V": "foo: :: baz"}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toYaml }}`, ExpectedOutput: `""`},
		{Name: "TestVariableInput", Input: `{{ .V | toYaml }}`, ExpectedOutput: "bar: baz\nfoo: 55", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
		{Name: "TestInvalidInput", Input: `{{ .V | toYaml }}`, ExpectedErr: "yaml encode error", Data: map[string]any{"V": make(chan int)}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToIndentYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toIndentYaml 8 }}`, ExpectedOutput: `""`},
		{Name: "TestVariableInput", Input: `{{ .V | toIndentYaml 8 }}`, ExpectedOutput: "bar: baz\nfoo:\n        bar: baz\n        baz: bar", Data: map[string]any{"V": map[string]any{"foo": map[string]any{"baz": "bar", "bar": "baz"}, "bar": "baz"}}},
		{Name: "TestInvalidInput", Input: `{{ .V | toIndentYaml 8 }}`, ExpectedErr: "yaml encode error", Data: map[string]any{"V": make(chan int)}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustFromJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestEmptyInput",
			Input:          `{{ "" | mustFromJson }}`,
			ExpectedOutput: "",
			Data:           nil,
			ExpectedErr:    "unexpected end",
		},
		{
			Name:           "TestVariableInput",
			Input:          `{{ .V | mustFromJson }}`,
			ExpectedOutput: "map[foo:55]",
			Data:           map[string]any{"V": `{"foo": 55}`},
			ExpectedErr:    "",
		},
		{
			Name:           "TestInvalidInput",
			Input:          `{{ .V | mustFromJson }}`,
			ExpectedOutput: "",
			Data:           map[string]any{"V": "{3}"},
			ExpectedErr:    "invalid character '3'",
		},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestEmptyInput",
			Input:          `{{ "" | mustToJson }}`,
			ExpectedOutput: `""`,
			ExpectedErr:    "",
		},
		{
			Name:           "TestVariableInput",
			Input:          `{{ .V | mustToJson }}`,
			ExpectedOutput: `{"bar":"baz","foo":55}`,
			Data:           map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			ExpectedErr:    "",
		},
		{
			Name:           "TestInvalidInput",
			Input:          `{{ .V | mustToJson }}`,
			ExpectedOutput: "",
			Data:           map[string]any{"V": make(chan int)},
			ExpectedErr:    "json: unsupported type: chan int",
		},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToPrettyJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestEmptyInput",
			Input:          `{{ "" | mustToPrettyJson }}`,
			ExpectedOutput: `""`,
			ExpectedErr:    "",
		},
		{
			Name:           "TestVariableInput",
			Input:          `{{ .V | mustToPrettyJson }}`,
			ExpectedOutput: "{\n  \"bar\": \"baz\",\n  \"foo\": 55\n}",
			Data:           map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			ExpectedErr:    "",
		},
		{
			Name:           "TestInvalidInput",
			Input:          `{{ .V | mustToPrettyJson }}`,
			ExpectedOutput: "",
			Data:           map[string]any{"V": make(chan int)},
			ExpectedErr:    "json: unsupported type: chan int",
		},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToRawJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestEmptyInput",
			Input:          `{{ "" | mustToRawJson }}`,
			ExpectedOutput: `""`,
			ExpectedErr:    "",
		},
		{
			Name:           "TestVariableInput",
			Input:          `{{ .V | mustToRawJson }}`,
			ExpectedOutput: `{"bar":"baz","foo":55}`,
			Data:           map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			ExpectedErr:    "",
		},
		{
			Name:           "TestInvalidInput",
			Input:          `{{ .V | mustToRawJson }}`,
			ExpectedOutput: "",
			Data:           map[string]any{"V": make(chan int)},
			ExpectedErr:    "json: unsupported type: chan int",
		},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustFromYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestEmptyInput",
			Input:          `{{ "foo: :: baz" | mustFromYaml }}`,
			ExpectedOutput: "",
			Data:           nil,
			ExpectedErr:    "yaml: mapping values are not allowed in this context",
		},
		{
			Name:           "TestVariableInput",
			Input:          `{{ .V | mustFromYaml }}`,
			ExpectedOutput: "map[bar:map[baz:1] foo:55]",
			Data:           map[string]any{"V": "foo: 55\nbar:\n  baz: 1\n"},
			ExpectedErr:    "",
		},
		{
			Name:           "TestInvalidInput",
			Input:          `{{ .V | mustFromYaml }}`,
			ExpectedOutput: "",
			Data:           map[string]any{"V": ":"},
			ExpectedErr:    "did not find expected key",
		},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{
			Name:           "TestEmptyInput",
			Input:          `{{ "" | mustToYaml }}`,
			ExpectedOutput: `""`,
			ExpectedErr:    "",
		},
		{
			Name:           "TestVariableInput",
			Input:          `{{ .V | mustToYaml }}`,
			ExpectedOutput: "bar: baz\nfoo: 55",
			Data:           map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			ExpectedErr:    "",
		},
		{
			Name:           "TestInvalidInput",
			Input:          `{{ .V | mustToYaml }}`,
			ExpectedOutput: "",
			Data:           map[string]any{"V": func() {}},
			ExpectedErr:    "cannot marshal type: func()",
		},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}
