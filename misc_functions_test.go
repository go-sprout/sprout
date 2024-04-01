package sprout

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

// TestHelper is a helper function that performs common setup tasks for tests.
func runTemplate(t *testing.T, handler *FunctionHandler, tmplString string) (string, error) {
	tmpl, err := template.New("test").Funcs(FuncMap(WithFunctionHandler(handler))).Parse(tmplString)
	if err != nil {
		assert.FailNow(t, "Failed to parse template", err)
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "test", nil)
	return buf.String(), err
}

// TestHello asserts the Hello method returns the expected greeting.
func TestHello(t *testing.T) {
	handler := NewFunctionHandler()
	expected := "Hello, World!"

	assert.Equal(t, expected, handler.Hello())

	tmplResponse, err := runTemplate(t, handler, `{{hello}}`)
	assert.Nil(t, err)
	assert.Equal(t, expected, tmplResponse)
}
