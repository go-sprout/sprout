package encoding_test

import (
	"testing"

	"github.com/go-sprout/sprout/pesticide"
	"github.com/go-sprout/sprout/registry/encoding"
)

func TestBase64Encode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base64Encode }}`, Expected: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "Hello World" | base64Encode }}`, Expected: "SGVsbG8gV29ybGQ="},
		{Name: `TestFromVariableInput`, Input: `{{ .V | base64Encode }}`, Expected: "SGVsbG8gV29ybGQ=", Data: map[string]any{"V": "Hello World"}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestBase64Decode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base64Decode }}`, Expected: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "SGVsbG8gV29ybGQ=" | base64Decode }}`, Expected: "Hello World"},
		{Name: "TestFromVariableInput", Input: `{{ .V | base64Decode }}`, Expected: "Hello World", Data: map[string]any{"V": "SGVsbG8gV29ybGQ="}},
		{Name: "TestInvalidInput", Input: `{{ "SGVsbG8gV29ybGQ" | base64Decode }}`, Expected: ""},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestBase32Encode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base32Encode }}`, Expected: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "Hello World" | base32Encode }}`, Expected: "JBSWY3DPEBLW64TMMQ======"},
		{Name: "TestFromVariableInput", Input: `{{ .V | base32Encode }}`, Expected: "JBSWY3DPEBLW64TMMQ======", Data: map[string]any{"V": "Hello World"}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestBase32Decode(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestWithoutInput", Input: `{{ "" | base32Decode }}`, Expected: ""},
		{Name: "TestHelloWorldInput", Input: `{{ "JBSWY3DPEBLW64TMMQ======" | base32Decode }}`, Expected: "Hello World"},
		{Name: "TestFromVariableInput", Input: `{{ .V | base32Decode }}`, Expected: "Hello World", Data: map[string]any{"V": "JBSWY3DPEBLW64TMMQ======"}},
		{Name: "TestInvalidInput", Input: `{{ "JBSWY3DPEBLW64TMMQ" | base32Decode }}`, Expected: ""},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestFromJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | fromJson }}`, Expected: "<no value>"},
		{Name: "TestVariableInput", Input: `{{ .V | fromJson }}`, Expected: "map[foo:55]", Data: map[string]any{"V": `{"foo": 55}`}},
		{Name: "TestAccessField", Input: `{{ (.V | fromJson).foo }}`, Expected: "55", Data: map[string]any{"V": `{"foo": 55}`}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toJson }}`, Expected: "\"\""},
		{Name: "TestVariableInput", Input: `{{ .V | toJson }}`, Expected: "{\"bar\":\"baz\",\"foo\":55}", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToPrettyJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toPrettyJson }}`, Expected: "\"\""},
		{Name: "TestVariableInput", Input: `{{ .V | toPrettyJson }}`, Expected: "{\n  \"bar\": \"baz\",\n  \"foo\": 55\n}", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToRawJson(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toRawJson }}`, Expected: "\"\""},
		{Name: "TestVariableInput", Input: `{{ .V | toRawJson }}`, Expected: "{\"bar\":\"baz\",\"foo\":55}", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestFromYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | fromYaml }}`, Expected: "map[]"},
		{Name: "TestVariableInput", Input: `{{ .V | fromYaml }}`, Expected: "map[bar:map[baz:1] foo:55]", Data: map[string]any{"V": "foo: 55\nbar:\n  baz: 1\n"}},
		{Name: "TestAccessField", Input: `{{ (.V | fromYaml).foo }}`, Expected: "55", Data: map[string]any{"V": "foo: 55"}},
		{Name: "TestInvalidInput", Input: "{{ .V | fromYaml }}", Expected: "<no value>", Data: map[string]any{"V": "foo: :: baz"}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestToYAML(t *testing.T) {
	tc := []pesticide.TestCase{
		{Name: "TestEmptyInput", Input: `{{ "" | toYaml }}`, Expected: "\"\""},
		{Name: "TestVariableInput", Input: `{{ .V | toYaml }}`, Expected: "bar: baz\nfoo: 55", Data: map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	pesticide.RunTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustFromJson(t *testing.T) {
	tc := []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestEmptyInput",
				Input:    `{{ "" | mustFromJson }}`,
				Expected: "",
				Data:     nil,
			},
			ExpectedErr: "unexpected end",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestVariableInput",
				Input:    `{{ .V | mustFromJson }}`,
				Expected: "map[foo:55]",
				Data:     map[string]any{"V": `{"foo": 55}`},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidInput",
				Input:    `{{ .V | mustFromJson }}`,
				Expected: "",
				Data:     map[string]any{"V": "{3}"},
			},
			ExpectedErr: "invalid character '3'",
		},
	}

	pesticide.RunMustTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToJson(t *testing.T) {
	tc := []pesticide.MustTestCase{
		{
			TestCase:    pesticide.TestCase{Name: "TestEmptyInput", Input: `{{ "" | mustToJson }}`, Expected: "\"\""},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestVariableInput",
				Input:    `{{ .V | mustToJson }}`,
				Expected: "{\"bar\":\"baz\",\"foo\":55}",
				Data:     map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidInput",
				Input:    `{{ .V | mustToJson }}`,
				Expected: "",
				Data:     map[string]any{"V": make(chan int)},
			},
			ExpectedErr: "json: unsupported type: chan int",
		},
	}

	pesticide.RunMustTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToPrettyJson(t *testing.T) {
	tc := []pesticide.MustTestCase{
		{
			TestCase:    pesticide.TestCase{Name: "TestEmptyInput", Input: `{{ "" | mustToPrettyJson }}`, Expected: "\"\""},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestVariableInput",
				Input:    `{{ .V | mustToPrettyJson }}`,
				Expected: "{\n  \"bar\": \"baz\",\n  \"foo\": 55\n}",
				Data:     map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidInput",
				Input:    `{{ .V | mustToPrettyJson }}`,
				Expected: "",
				Data:     map[string]any{"V": make(chan int)},
			},
			ExpectedErr: "json: unsupported type: chan int",
		},
	}

	pesticide.RunMustTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToRawJson(t *testing.T) {
	tc := []pesticide.MustTestCase{
		{
			TestCase:    pesticide.TestCase{Name: "TestEmptyInput", Input: `{{ "" | mustToRawJson }}`, Expected: "\"\""},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestVariableInput",
				Input:    `{{ .V | mustToRawJson }}`,
				Expected: "{\"bar\":\"baz\",\"foo\":55}",
				Data:     map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidInput",
				Input:    `{{ .V | mustToRawJson }}`,
				Expected: "",
				Data:     map[string]any{"V": make(chan int)},
			},
			ExpectedErr: "json: unsupported type: chan int",
		},
	}

	pesticide.RunMustTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustFromYAML(t *testing.T) {
	tc := []pesticide.MustTestCase{
		{
			TestCase: pesticide.TestCase{
				Name:     "TestEmptyInput",
				Input:    `{{ "foo: :: baz" | mustFromYaml }}`,
				Expected: "",
				Data:     nil,
			},
			ExpectedErr: "yaml: mapping values are not allowed in this context",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestVariableInput",
				Input:    `{{ .V | mustFromYaml }}`,
				Expected: "map[bar:map[baz:1] foo:55]",
				Data:     map[string]any{"V": "foo: 55\nbar:\n  baz: 1\n"},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidInput",
				Input:    `{{ .V | mustFromYaml }}`,
				Expected: "",
				Data:     map[string]any{"V": ":"},
			},
			ExpectedErr: "did not find expected key",
		},
	}

	pesticide.RunMustTestCases(t, encoding.NewRegistry(), tc)
}

func TestMustToYAML(t *testing.T) {
	tc := []pesticide.MustTestCase{
		{
			TestCase:    pesticide.TestCase{Name: "TestEmptyInput", Input: `{{ "" | mustToYaml }}`, Expected: "\"\""},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestVariableInput",
				Input:    `{{ .V | mustToYaml }}`,
				Expected: "bar: baz\nfoo: 55",
				Data:     map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}},
			},
			ExpectedErr: "",
		},
		{
			TestCase: pesticide.TestCase{
				Name:     "TestInvalidInput",
				Input:    `{{ .V | mustToYaml }}`,
				Expected: "",
				Data:     map[string]any{"V": make(chan int)},
			},
			ExpectedErr: "cannot marshal type: chan int",
		},
	}

	pesticide.RunMustTestCases(t, encoding.NewRegistry(), tc)
}
