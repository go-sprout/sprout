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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/maps"
	"github.com/go-sprout/sprout/registry/reflect"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/strings"
)

type TestCase struct {
	Name           string
	Input          string
	Data           map[string]any
	ExpectedOutput string
	ExpectedErr    string
}

func RunTestCases(t *testing.T, registry sprout.Registry, tc []TestCase) {
	t.Helper()

	for _, test := range tc {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			tmplResponse, err := runTemplate(t, testHandler(registry), test.Input, test.Data)
			if test.ExpectedErr != "" {
				require.ErrorContains(t, err, test.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, test.ExpectedOutput, tmplResponse)
		})
	}
}

func TestTemplate(t *testing.T, registry sprout.Registry, tmplString string, data any) (string, error) {
	t.Helper()

	return runTemplate(t, testHandler(registry), tmplString, data)
}

// RunTestCasesWithFuncs runs test cases using a custom FuncMap.
// Use this when you need to test functions that aren't registered through
// a standard sprout.Registry (e.g., sprigin compatibility functions).
func RunTestCasesWithFuncs(t *testing.T, funcs template.FuncMap, tc []TestCase) {
	t.Helper()

	for _, test := range tc {
		t.Run(test.Name, func(t *testing.T) {
			t.Helper()

			tmpl, err := template.New("test").Funcs(funcs).Parse(test.Input)
			require.NoError(t, err, "Failed to parse template")

			var buf bytes.Buffer
			err = tmpl.ExecuteTemplate(&buf, "test", test.Data)

			if test.ExpectedErr != "" {
				require.ErrorContains(t, err, test.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, test.ExpectedOutput, buf.String())
		})
	}
}

func runTemplate(t *testing.T, handler sprout.Handler, tmplString string, data any) (string, error) {
	t.Helper()

	tmpl, err := template.New("test").Funcs(handler.Build()).Parse(tmplString)
	require.NoError(t, err, "Failed to parse template")

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "test", data)
	return buf.String(), err
}

func testHandler(registry sprout.Registry) *sprout.DefaultHandler {
	handler := sprout.New()
	_ = handler.AddRegistries(
		strings.NewRegistry(),
		slices.NewRegistry(),
		maps.NewRegistry(),
		reflect.NewRegistry(),
		registry,
	)

	return handler
}
