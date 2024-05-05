package sprout

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	input    string
	expected string
	data     map[string]any
}

type testCases []testCase

func runTestCases(t *testing.T, tc testCases) {
	t.Helper()
	handler := NewFunctionHandler()

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()

			tmplResponse, err := runTemplate(t, handler, test.input, test.data)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, tmplResponse)
		})
	}
}

func runTemplate(t *testing.T, handler *FunctionHandler, tmplString string, data any) (string, error) {
	t.Helper()

	tmpl, err := template.New("test").Funcs(FuncMap(WithFunctionHandler(handler))).Parse(tmplString)
	if err != nil {
		assert.FailNow(t, "Failed to parse template", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "test", data)
	return buf.String(), err
}
