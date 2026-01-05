package sprigin_test

import (
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"

	"github.com/go-sprout/sprout/sprigin"
)

func TestSprigin(t *testing.T) {
	for _, tc := range []struct {
		name     string
		funcMap  template.FuncMap
		text     string
		data     any
		expected string
	}{
		{
			name:     "bcrypt",
			funcMap:  sprigin.HermeticTxtFuncMap(),
			text:     `{{ bcrypt "abc" | trunc 7 }}`, // Only the first seven bytes of bcrypt output is consistent.
			expected: "$2a$10$",
		},
		{
			name:     "derivePassword",
			funcMap:  sprigin.HermeticTxtFuncMap(),
			text:     `{{ derivePassword 1 "long" "password" "user" "example.com" }}`,
			expected: "ZedaFaxcZaso9*",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tmpl, err := template.New(tc.name).Funcs(tc.funcMap).Parse(tc.text)
			assert.NoError(t, err)
			var builder strings.Builder
			assert.NoError(t, tmpl.Execute(&builder, tc.data))
			assert.Equal(t, tc.expected, builder.String())
		})
	}
}
