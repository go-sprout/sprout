package compatibility

import (
	"strings"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/stretchr/testify/assert"

	"github.com/go-sprout/sprout/sprigin"
)

func TestFunctions(t *testing.T) {
	for _, tc := range []struct {
		name  string
		fails bool
		text  string
		data  any
	}{
		{
			name:  "append",
			fails: true, // https://github.com/go-sprout/sprout/issues/162, warning: this test generates warning messages until it overflows the stack
			text:  `{{ append (list "a") (list "b") }}`,
		},
		{
			name:  "derivePassword",
			fails: true, // https://github.com/go-sprout/sprout/pull/157
			text:  `{{ derivePassword 1 "long" "password" "user" "example.com" }}`,
		},
		{
			name: "dig_with_dots_in_keys",
			text: `{{ dig ".key" "default" . }}`,
			data: map[string]any{
				".key": "value",
			},
		},
		{
			name: "trim",
			text: `{{ trim "   hello   " }}`,
		},
		{
			name:  "substr",
			fails: true, // https://github.com/go-sprout/sprout/issues/163
			text:  `{{ substr 1 -1 "abcdef" }}`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.fails {
				t.Skip("skipping known failing test")
			}

			sprigTemplate, err := template.New(tc.name).Funcs(sprig.FuncMap()).Parse(tc.text)
			assert.NoError(t, err)
			var sprigBuilder strings.Builder
			assert.NoError(t, sprigTemplate.Execute(&sprigBuilder, tc.data))

			spriginTemplate, err := template.New(tc.name).Funcs(sprigin.FuncMap()).Parse(tc.text)
			assert.NoError(t, err)
			var spriginBuilder strings.Builder
			assert.NoError(t, spriginTemplate.Execute(&spriginBuilder, tc.data))

			assert.Equal(t, sprigBuilder.String(), spriginBuilder.String())
		})
	}
}
