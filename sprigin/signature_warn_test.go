package sprigin

import (
	"bytes"
	"io"
	"testing"
	ttemplate "text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// renderWithLogger executes tmpl with a func map logging to a buffer, and
// returns everything the handler logged while rendering.
func renderWithLogger(t *testing.T, tmpl string) string {
	t.Helper()

	var buf bytes.Buffer
	tpl, err := ttemplate.New("warn").Funcs(FuncMapWith(WithLogger(bufferedLogger(&buf)))).Parse(tmpl)
	require.NoError(t, err)
	require.NoError(t, tpl.Execute(io.Discard, nil))

	return buf.String()
}

func TestSignatureWarn(t *testing.T) {
	var buf bytes.Buffer

	handler := NewSprigHandler()
	require.NoError(t, WithLogger(bufferedLogger(&buf))(handler))

	handler.SignatureWarn("append", sprigSignAppend, sproutSignAppend)

	// slog escapes the quotes of the message, so the assertions use the parts of
	// the signatures that carry no quote.
	logged := buf.String()
	assert.Contains(t, logged, "The signature of `append` has changed")
	assert.Contains(t, logged, "{{ $list = append $list")
	assert.Contains(t, logged, "{{ $list = $list | append")
}

func TestAmbiguousSignatureWarn(t *testing.T) {
	var buf bytes.Buffer

	handler := NewSprigHandler()
	require.NoError(t, WithLogger(bufferedLogger(&buf))(handler))

	handler.AmbiguousSignatureWarn("append", sprigSignAppend, sproutSignAppend)

	logged := buf.String()
	assert.Contains(t, logged, "Template function `append` is ambiguous")
	assert.Contains(t, logged, "match both signatures")
	assert.Contains(t, logged, "use the `sprout` package directly")
	assert.Contains(t, logged, "notice=deprecated")
	// The migration wording must not be used here, the caller may already be on
	// the sprout signature.
	assert.NotContains(t, logged, "has changed")
}

// TestSignatureWarnShowsTheAssignment covers the confusion reported in #180:
// the migration message used to show calls without assignment, which reads as
// if the call alone was enough to accumulate in a `range`.
func TestSignatureWarnShowsTheAssignment(t *testing.T) {
	for _, tc := range []struct {
		name     string
		tmpl     string
		expected string
	}{
		// slog escapes the quotes of the message, so the expectation stops before
		// the quoted value.
		{"append", `{{ $l := list "a" }}{{ $l = append $l "b" }}`, `{{ $list = $list | append`},
		{"prepend", `{{ $l := list "a" }}{{ $l = prepend $l "b" }}`, `{{ $list = $list | prepend`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			logged := renderWithLogger(t, tc.tmpl)

			assert.Contains(t, logged, "has changed")
			assert.Contains(t, logged, tc.expected)
		})
	}
}

// TestAmbiguousCallsWarnAsAmbiguous ensures a call whose arguments match both
// signatures is reported as ambiguous instead of being told to migrate to the
// signature it may already use.
func TestAmbiguousCallsWarnAsAmbiguous(t *testing.T) {
	for _, tc := range []struct {
		name string
		tmpl string
	}{
		{"append two lists", `{{ $a := list "a" }}{{ $b := list "b" }}{{ $a | append $b }}`},
		{"prepend two lists", `{{ $a := list "a" }}{{ $b := list "b" }}{{ $a | prepend $b }}`},
		{"without two lists", `{{ $a := list "a" }}{{ $b := list "b" }}{{ $a | without $b }}`},
		{"set two dicts", `{{ $a := dict "a" 1 }}{{ $b := dict "b" 2 }}{{ set $a "key" $b }}`},
		{"get two dicts", `{{ $a := dict "a" 1 }}{{ $b := dict "b" 2 }}{{ get $a $b }}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			logged := renderWithLogger(t, tc.tmpl)

			assert.Contains(t, logged, "is ambiguous")
			assert.NotContains(t, logged, "has changed")
		})
	}
}

// TestUndecidableCallsDoNotWarn ensures a call with no receiver at all is left
// to the registry error handling, without a misleading migration message.
func TestUndecidableCallsDoNotWarn(t *testing.T) {
	for _, tc := range []struct {
		name string
		tmpl string
	}{
		{"append without a list", `{{ append "a" "b" }}`},
		{"prepend without a list", `{{ prepend "a" "b" }}`},
		{"get without a dict", `{{ get "a" "b" }}`},
		{"hasKey without a dict", `{{ hasKey "a" "b" }}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			logged := renderWithLogger(t, tc.tmpl)

			assert.NotContains(t, logged, "has changed")
			assert.NotContains(t, logged, "is ambiguous")
		})
	}
}

// TestSproutSignatureStaysSilent ensures a template already using the sprout
// signature is not warned about anything.
func TestSproutSignatureStaysSilent(t *testing.T) {
	logged := renderWithLogger(t, `{{ $l := list "a" }}{{ $l = $l | append "b" }}{{ $l | without "a" }}`)

	assert.Empty(t, logged)
}
