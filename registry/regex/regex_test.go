package regex_test

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/group/all"
	"github.com/go-sprout/sprout/registry/regex"
)

// TestRegistryPrecedenceOverRegexp ensures the documented way to opt in the
// `regex` registry while still using a group shipping the deprecated `regexp`
// one works: functions are never overwritten, so the registry registered first
// wins.
func TestRegistryPrecedenceOverRegexp(t *testing.T) {
	handler := sprout.New(
		sprout.WithRegistries(regex.NewRegistry()),
		sprout.WithGroups(all.RegistryGroup()),
	)

	tmpl, err := template.New("test").Funcs(handler.Build()).Parse(`{{ "banana" | regexSplit "a" -1 }}`)
	require.NoError(t, err)

	var buf bytes.Buffer
	require.NoError(t, tmpl.Execute(&buf, nil))
	require.Equal(t, "[b n n ]", buf.String())
}
