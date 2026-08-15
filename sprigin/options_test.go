package sprigin

import (
	"bytes"
	htemplate "html/template"
	"io"
	"log/slog"
	"testing"
	ttemplate "text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// bufferedLogger returns a logger writing to buf, so tests can assert on what
// the sprig handler actually logs.
func bufferedLogger(buf *bytes.Buffer) *slog.Logger {
	return slog.New(slog.NewTextHandler(buf, nil))
}

// renderDeprecatedGet executes a template using the old Sprig `get` signature,
// which makes the handler emit a signature deprecation warning.
func renderDeprecatedGet(t *testing.T, funcMap ttemplate.FuncMap) {
	t.Helper()

	tpl, err := ttemplate.New("deprecated-get").Funcs(funcMap).Parse(`{{ get (dict "a" "b") "a" }}`)
	require.NoError(t, err)
	require.NoError(t, tpl.Execute(io.Discard, nil))
}

func TestWithLogger(t *testing.T) {
	t.Run("sets custom logger on handler", func(t *testing.T) {
		var buf bytes.Buffer
		customLogger := bufferedLogger(&buf)

		handler := NewSprigHandler()
		originalLogger := handler.Logger()

		err := WithLogger(customLogger)(handler)

		require.NoError(t, err)
		assert.NotEqual(t, originalLogger, handler.Logger())
		assert.Equal(t, customLogger, handler.Logger())
	})

	t.Run("custom logger receives handler warnings", func(t *testing.T) {
		var buf bytes.Buffer

		handler := NewSprigHandler()
		require.NoError(t, WithLogger(bufferedLogger(&buf))(handler))

		handler.SignatureWarn("get", "{{ get $dict \"key\" }}", "{{ $dict | get \"key\" }}")

		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})

	t.Run("nil logger returns error and keeps default", func(t *testing.T) {
		handler := NewSprigHandler()
		originalLogger := handler.Logger()
		require.NotNil(t, originalLogger)

		err := WithLogger(nil)(handler)

		require.ErrorContains(t, err, "logger is nil")
		assert.Equal(t, originalLogger, handler.Logger())
	})
}

func TestFuncMapWith(t *testing.T) {
	t.Run("deprecation notices are routed to the custom logger", func(t *testing.T) {
		var buf bytes.Buffer

		funcMap := FuncMapWith(WithLogger(bufferedLogger(&buf)))
		renderDeprecatedGet(t, funcMap)

		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})

	t.Run("invalid option is skipped and the func map is still built", func(t *testing.T) {
		funcMap := FuncMapWith(WithLogger(nil))

		assert.Contains(t, funcMap, "get")
	})

	t.Run("without options it matches the legacy func map", func(t *testing.T) {
		withOptions := FuncMapWith()
		legacy := FuncMap()

		assert.Len(t, withOptions, len(legacy))
		for name := range legacy {
			assert.Contains(t, withOptions, name)
		}
	})
}

func TestHermeticFuncMapsWith(t *testing.T) {
	// nonHermeticSample and hermeticSample are taken from nonhermeticFunctions
	// to assert the hermetic variants still filter the func map.
	const nonHermeticSample, hermeticSample = "env", "get"

	t.Run("HermeticTxtFuncMapWith filters and uses the custom logger", func(t *testing.T) {
		var buf bytes.Buffer

		funcMap := HermeticTxtFuncMapWith(WithLogger(bufferedLogger(&buf)))

		assert.Contains(t, funcMap, hermeticSample)
		assert.NotContains(t, funcMap, nonHermeticSample)

		renderDeprecatedGet(t, funcMap)
		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})

	t.Run("HermeticHtmlFuncMapWith filters and uses the custom logger", func(t *testing.T) {
		var buf bytes.Buffer

		funcMap := HermeticHtmlFuncMapWith(WithLogger(bufferedLogger(&buf)))

		assert.Contains(t, funcMap, hermeticSample)
		assert.NotContains(t, funcMap, nonHermeticSample)

		renderDeprecatedGet(t, ttemplate.FuncMap(funcMap))
		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})

	t.Run("TxtFuncMapWith keeps non hermetic functions", func(t *testing.T) {
		var buf bytes.Buffer

		funcMap := TxtFuncMapWith(WithLogger(bufferedLogger(&buf)))

		assert.Contains(t, funcMap, nonHermeticSample)

		renderDeprecatedGet(t, funcMap)
		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})

	t.Run("HtmlFuncMapWith keeps non hermetic functions", func(t *testing.T) {
		var buf bytes.Buffer

		funcMap := HtmlFuncMapWith(WithLogger(bufferedLogger(&buf)))

		assert.Contains(t, funcMap, nonHermeticSample)

		renderDeprecatedGet(t, ttemplate.FuncMap(funcMap))
		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})

	t.Run("GenericFuncMapWith keeps non hermetic functions", func(t *testing.T) {
		var buf bytes.Buffer

		funcMap := GenericFuncMapWith(WithLogger(bufferedLogger(&buf)))

		assert.Contains(t, funcMap, nonHermeticSample)

		renderDeprecatedGet(t, ttemplate.FuncMap(funcMap))
		assert.Contains(t, buf.String(), "The signature of `get` has changed")
	})
}

// TestBackwardCompatibleSignatures pins the signatures of the historical
// entrypoints: they must stay callable without arguments and assignable to
// their original function type, so upgrading sprigin is never a breaking change.
func TestBackwardCompatibleSignatures(t *testing.T) {
	// Conversions, not declarations: they fail to compile if the signature ever
	// gains a parameter, without staticcheck asking to drop the explicit type.
	_ = (func() ttemplate.FuncMap)(FuncMap)
	_ = (func() ttemplate.FuncMap)(TxtFuncMap)
	_ = (func() ttemplate.FuncMap)(HermeticTxtFuncMap)
	_ = (func() htemplate.FuncMap)(HtmlFuncMap)
	_ = (func() htemplate.FuncMap)(HermeticHtmlFuncMap)
	_ = (func() map[string]any)(GenericFuncMap)

	assert.Contains(t, HermeticTxtFuncMap(), "get")
	assert.NotContains(t, HermeticTxtFuncMap(), "env")
	assert.Contains(t, TxtFuncMap(), "env")
}
