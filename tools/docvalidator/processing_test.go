package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHandlerFor_Default tests handlerFor when the documentation file has no
// dedicated handler. It verifies that the handler built from all registries is
// returned.
func TestHandlerFor_Default(t *testing.T) {
	assert.Same(t, sproutHandler, handlerFor(filepath.Join("docs", "registries", "example1.md")))
}

// TestHandlerFor_Dedicated tests handlerFor when the documentation file has a
// dedicated handler. It verifies that this handler is returned instead of the
// default one, so the examples are validated against the right registry.
func TestHandlerFor_Dedicated(t *testing.T) {
	file := filepath.Join("docs", "registries", "regex.md")

	handler := handlerFor(file)
	assert.NotSame(t, sproutHandler, handler)
	assert.Same(t, dedicatedHandlers[file], handler)
}

// TestProcessExample_DedicatedHandler tests processExample on a file having a
// dedicated handler. The `regex` registry takes precedence over the deprecated
// `regexp` one, so the main parameter is expected to be the last one.
func TestProcessExample_DedicatedHandler(t *testing.T) {
	example := Example{
		FuncName: "regexSplit",
		File:     filepath.Join("docs", "registries", "regex.md"),
		Code:     `{{ "banana" | regexSplit "a" -1 }}`,
		Expected: "[b n n ]",
	}
	err := processExample(example)
	assert.NoError(t, err)
}

// TestProcessExample_Success tests processExample when the template executes
// successfully and the output matches the expected output.
func TestProcessExample_Success(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ hello }}`,
		Expected: "Hello!",
	}
	err := processExample(example)
	assert.NoError(t, err)
}

func TestProcessExample_Skipped(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ hello }}`,
		Expected: "Hello!",
		Skipped:  true,
	}
	err := processExample(example)
	assert.NoError(t, err)
}

// TestProcessExample_ErrorOnTemplate tests processExample when the template contains a syntax error.
// It verifies that an error is returned indicating a problem with parsing the template.
func TestProcessExample_ErrorOnTemplate(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ errorFunction }}`,
		Expected: "",
	}
	err := processExample(example)
	assert.ErrorContains(t, err, "error parsing template")
}

// TestProcessExample_ErrorExpected tests processExample when the template
// executes with an error and the expected output is an error message.
// It verifies that no error is returned, since the error is expected.
func TestProcessExample_ErrorExpected(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ min }}`,
		Expected: "Error",
	}
	err := processExample(example)
	assert.NoError(t, err)
}

// TestProcessExample_ErrorExceptedButNoError tests processExample when the template
// executes successfully, but the expected output is an error message.
// It verifies that an error is returned indicating that an error was expected but
// the template executed successfully.
func TestProcessExample_ErrorExceptedButNoError(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ hello }}`,
		Expected: "Error",
	}
	err := processExample(example)
	assert.ErrorContains(t, err, "expected an error but the template executed successfully")
}

// TestProcessExample_ErrorUnexpected tests processExample when the template executes with an error and the expected output is not an error message.
// It verifies that an error is returned indicating that an unexpected error occurred.
func TestProcessExample_ErrorUnexpected(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ min }}`,
		Expected: "1",
	}
	err := processExample(example)
	assert.ErrorContains(t, err, "unexpected error during template execution")
}

// TestProcessExample_OutputMismatch tests processExample when the template
// executes successfully, but the actual output does not match the expected output.
// It verifies that an error is returned indicating that the output does not match.
func TestProcessExample_OutputMismatch(t *testing.T) {
	example := Example{
		FuncName: "FunctionName",
		File:     "docs/registries/example1.md",
		Code:     `{{ "Hello" }}`,
		Expected: "Hello, World!",
	}
	err := processExample(example)
	assert.ErrorContains(t, err, "output mismatch")
}
