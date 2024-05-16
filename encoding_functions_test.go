package sprout

import "testing"

func TestBase64Encode(t *testing.T) {
	var tests = testCases{
		{"TestWithoutInput", `{{ "" | base64Encode }}`, "", nil},
		{"TestHelloWorldInput", `{{ "Hello World" | base64Encode }}`, "SGVsbG8gV29ybGQ=", nil},
		{"TestFromVariableInput", `{{ .V | base64Encode }}`, "SGVsbG8gV29ybGQ=", map[string]any{"V": "Hello World"}},
	}

	runTestCases(t, tests)
}

func TestBase64Decode(t *testing.T) {
	var tests = testCases{
		{"TestWithoutInput", `{{ "" | base64Decode }}`, "", nil},
		{"TestHelloWorldInput", `{{ "SGVsbG8gV29ybGQ=" | base64Decode }}`, "Hello World", nil},
		{"TestFromVariableInput", `{{ .V | base64Decode }}`, "Hello World", map[string]any{"V": "SGVsbG8gV29ybGQ="}},
		{"TestInvalidInput", `{{ "SGVsbG8gV29ybGQ" | base64Decode }}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestBase32Encode(t *testing.T) {
	var tests = testCases{
		{"TestWithoutInput", `{{ "" | base32Encode }}`, "", nil},
		{"TestHelloWorldInput", `{{ "Hello World" | base32Encode }}`, "JBSWY3DPEBLW64TMMQ======", nil},
		{"TestFromVariableInput", `{{ .V | base32Encode }}`, "JBSWY3DPEBLW64TMMQ======", map[string]any{"V": "Hello World"}},
	}

	runTestCases(t, tests)
}

func TestBase32Decode(t *testing.T) {
	var tests = testCases{
		{"TestWithoutInput", `{{ "" | base32Decode }}`, "", nil},
		{"TestHelloWorldInput", `{{ "JBSWY3DPEBLW64TMMQ======" | base32Decode }}`, "Hello World", nil},
		{"TestFromVariableInput", `{{ .V | base32Decode }}`, "Hello World", map[string]any{"V": "JBSWY3DPEBLW64TMMQ======"}},
		{"TestInvalidInput", `{{ "JBSWY3DPEBLW64TMMQ" | base32Decode }}`, "", nil},
	}

	runTestCases(t, tests)
}

func TestFromJson(t *testing.T) {
	var tests = testCases{
		{"TestEmptyInput", `{{ "" | fromJson }}`, "<no value>", nil},
		{"TestVariableInput", `{{ .V | fromJson }}`, "map[foo:55]", map[string]any{"V": `{"foo": 55}`}},
		{"TestAccessField", `{{ (.V | fromJson).foo }}`, "55", map[string]any{"V": `{"foo": 55}`}},
	}

	runTestCases(t, tests)
}

func TestToJson(t *testing.T) {
	var tests = testCases{
		{"TestEmptyInput", `{{ "" | toJson }}`, "\"\"", nil},
		{"TestVariableInput", `{{ .V | toJson }}`, "{\"bar\":\"baz\",\"foo\":55}", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	runTestCases(t, tests)
}

func TestToPrettyJson(t *testing.T) {
	var tests = testCases{
		{"TestEmptyInput", `{{ "" | toPrettyJson }}`, "\"\"", nil},
		{"TestVariableInput", `{{ .V | toPrettyJson }}`, "{\n  \"bar\": \"baz\",\n  \"foo\": 55\n}", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	runTestCases(t, tests)
}

func TestToRawJson(t *testing.T) {
	var tests = testCases{
		{"TestEmptyInput", `{{ "" | toRawJson }}`, "\"\"", nil},
		{"TestVariableInput", `{{ .V | toRawJson }}`, "{\"bar\":\"baz\",\"foo\":55}", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	runTestCases(t, tests)
}

func TestFromYAML(t *testing.T) {
	var tests = testCases{
		{"TestEmptyInput", `{{ "" | fromYaml }}`, "map[]", nil},
		{"TestVariableInput", `{{ .V | fromYaml }}`, "map[bar:map[baz:1] foo:55]", map[string]any{"V": "foo: 55\nbar:\n  baz: 1\n"}},
		{"TestAccessField", `{{ (.V | fromYaml).foo }}`, "55", map[string]any{"V": "foo: 55"}},
		{"TestInvalidInput", `{{ .V | fromYaml }}`, "<no value>", map[string]any{"V": "foo: :: baz"}},
	}

	runTestCases(t, tests)
}

func TestToYAML(t *testing.T) {
	var tests = testCases{
		{"TestEmptyInput", `{{ "" | toYaml }}`, "\"\"", nil},
		{"TestVariableInput", `{{ .V | toYaml }}`, "bar: baz\nfoo: 55", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}},
	}

	runTestCases(t, tests)
}

func TestMustToYAML(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestEmptyInput", `{{ "" | mustToYaml }}`, "\"\"", nil}, ""},
		{testCase{"TestVariableInput", `{{ .V | mustToYaml }}`, "bar: baz\nfoo: 55", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}}, ""},
		{testCase{"TestInvalidInput", `{{ .V | mustToYaml }}`, "", map[string]any{"V": make(chan int)}}, "json: unsupported type: chan int"},
	}

	runMustTestCases(t, tests)
}

func TestMustFromJson(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestEmptyInput", `{{ "" | mustFromJson }}`, "", nil}, "unexpected end"},
		{testCase{"TestVariableInput", `{{ .V | mustFromJson }}`, "map[foo:55]", map[string]any{"V": `{"foo": 55}`}}, ""},
		{testCase{"TestInvalidInput", `{{ .V | mustFromJson }}`, "", map[string]any{"V": "{3}"}}, "invalid character '3'"},
	}

	runMustTestCases(t, tests)
}

func TestMustToJson(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestEmptyInput", `{{ "" | mustToJson }}`, "\"\"", nil}, ""},
		{testCase{"TestVariableInput", `{{ .V | mustToJson }}`, "{\"bar\":\"baz\",\"foo\":55}", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}}, ""},
		{testCase{"TestInvalidInput", `{{ .V | mustToJson }}`, "", map[string]any{"V": make(chan int)}}, "json: unsupported type: chan int"},
	}

	runMustTestCases(t, tests)
}

func TestMustToPrettyJson(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestEmptyInput", `{{ "" | mustToPrettyJson }}`, "\"\"", nil}, ""},
		{testCase{"TestVariableInput", `{{ .V | mustToPrettyJson }}`, "{\n  \"bar\": \"baz\",\n  \"foo\": 55\n}", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}}, ""},
		{testCase{"TestInvalidInput", `{{ .V | mustToPrettyJson }}`, "", map[string]any{"V": make(chan int)}}, "json: unsupported type: chan int"},
	}

	runMustTestCases(t, tests)
}

func TestMustToRawJson(t *testing.T) {
	var tests = mustTestCases{
		{testCase{"TestEmptyInput", `{{ "" | mustToRawJson }}`, "\"\"", nil}, ""},
		{testCase{"TestVariableInput", `{{ .V | mustToRawJson }}`, "{\"bar\":\"baz\",\"foo\":55}", map[string]any{"V": map[string]any{"foo": 55, "bar": "baz"}}}, ""},
		{testCase{"TestInvalidInput", `{{ .V | mustToRawJson }}`, "", map[string]any{"V": make(chan int)}}, "json: unsupported type: chan int"},
	}

	runMustTestCases(t, tests)
}
