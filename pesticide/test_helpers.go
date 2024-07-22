/**
 * Pesticide are a package to help you to test your functions on a template
 * engine. It provides a set of test cases that you can use to test your
 * functions.
 * More pesticide for less bugs.
 */
package pesticide

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    string
	Expected string
	Data     map[string]any
}

type MustTestCase struct {
	TestCase
	ExpectedErr string
}

func RunTestCases(t *testing.T, registry registry.Registry, tc []TestCase) {
	t.Helper()
	handler := sprout.NewFunctionHandler()
	_ = handler.AddRegistry(registry)

	for _, test := range tc {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			tmplResponse, err := runTemplate(t, handler, test.Input, test.Data)
			assert.NoError(t, err)
			assert.Equal(t, test.Expected, tmplResponse)
		})
	}
}

func RunMustTestCases(t *testing.T, registry registry.Registry, tc []MustTestCase) {
	t.Helper()
	handler := sprout.NewFunctionHandler()
	_ = handler.AddRegistry(registry)

	for _, test := range tc {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			tmplResponse, err := runTemplate(t, handler, test.Input, test.Data)
			if test.ExpectedErr != "" {
				if assert.Error(t, err) {
					assert.Contains(t, err.Error(), test.ExpectedErr)
				}
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.Expected, tmplResponse)
		})
	}
}

func TestTemplate(t *testing.T, registry registry.Registry, tmplString string, data any) (string, error) {
	t.Helper()
	handler := sprout.NewFunctionHandler()
	_ = handler.AddRegistry(registry)

	return runTemplate(t, handler, tmplString, data)
}

func runTemplate(t *testing.T, handler registry.Handler, tmplString string, data any) (string, error) {
	t.Helper()

	tmpl, err := template.New("test").Funcs(sprout.FuncMap(sprout.WithFunctionHandler(handler))).Parse(tmplString)
	if err != nil {
		assert.FailNow(t, "Failed to parse template", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "test", data)
	return buf.String(), err
}
