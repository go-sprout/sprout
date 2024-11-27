package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExtractExamples tests that extractExamples correctly extracts examples from the markdown content.
func TestExtractExamples(t *testing.T) {
	content := `
### <mark>FunctionName</mark>

Some description.

{% tab title="Template Example" %}
` + "```go" + `
{{ example code }}
// Output: expected output

{{ example code inline }} // Output: expected inline output
` + "```" + `
{% endtab %}

### <mark>SecondFunctionName</mark>

Some description.

{% tab title="Template Example" %}
` + "```go" + `
{{ second example code }}
// Output: second expected output

{{ second example code inline }} // Output: second expected inline output
` + "```" + `
{% endtab %}
`
	filePath := "docs/registries/example1.md"
	expected := []Example{
		{
			FuncName: "FunctionName",
			File:     filePath,
			Code:     "{{ example code }}",
			Expected: "expected output",
		},
		{
			FuncName: "FunctionName",
			File:     filePath,
			Code:     "{{ example code inline }}",
			Expected: "expected inline output",
		},
		{
			FuncName: "SecondFunctionName",
			File:     filePath,
			Code:     "{{ second example code }}",
			Expected: "second expected output",
		},
		{
			FuncName: "SecondFunctionName",
			File:     filePath,
			Code:     "{{ second example code inline }}",
			Expected: "second expected inline output",
		},
	}
	result := extractExamples(content, filePath)
	assert.Equal(t, expected, result)
}

// TestParseCodeBlock tests that parseCodeBlock correctly parses code blocks into individual examples.
func TestParseCodeBlock(t *testing.T) {
	codeBlock := `
{{ example code }}
// Output: expected output

{{ inline example code }} // Output: expected output

{{ failed inline example code }} // Error

{{ failed example code }}
// Error

{{ multiple lines example code }}
{{ multiple lines example code }}
// Output:
expected output

{{ ["1", '1'] }} // Output: ["1", '1']

{{ random code }} // Output(will be different): raaanddoommmmm
`
	funcName := "FunctionName"
	filePath := "docs/registries/example1.md"
	expected := []Example{
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ example code }}",
			Expected: "expected output",
			Skipped:  false,
		},
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ inline example code }}",
			Expected: "expected output",
			Skipped:  false,
		},
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ failed inline example code }}",
			Expected: "Error",
			Skipped:  false,
		},
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ failed example code }}",
			Expected: "Error",
			Skipped:  false,
		},
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ multiple lines example code }}\n{{ multiple lines example code }}",
			Expected: "expected output",
			Skipped:  false,
		},
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ [\"1\", '1'] }}",
			Expected: "[\"1\", '1']",
			Skipped:  false,
		},
		{
			FuncName: funcName,
			File:     filePath,
			Code:     "{{ random code }}",
			Expected: "(will be different)",
			Skipped:  true,
		},
	}

	result := parseCodeBlock(codeBlock, funcName, filePath)
	assert.Equal(t, expected, result)
}
