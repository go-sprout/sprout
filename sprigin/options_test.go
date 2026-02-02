package sprigin

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithLogger(t *testing.T) {
	t.Run("sets custom logger on handler", func(t *testing.T) {
		var buf bytes.Buffer
		customLogger := slog.New(slog.NewTextHandler(&buf, nil))

		handler := NewSprigHandler()
		originalLogger := handler.Logger()

		err := WithLogger(customLogger)(handler)

		require.NoError(t, err)
		assert.NotEqual(t, originalLogger, handler.Logger())
		assert.Equal(t, customLogger, handler.Logger())
	})

	t.Run("nil logger returns error and keeps default", func(t *testing.T) {
		handler := NewSprigHandler()
		originalLogger := handler.Logger()
		require.NotNil(t, originalLogger)

		err := WithLogger(nil)(handler)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "logger is nil")
		assert.Equal(t, originalLogger, handler.Logger())
	})
}

func TestFuncMapWithLogger(t *testing.T) {
	t.Run("accepts WithLogger option", func(t *testing.T) {
		var buf bytes.Buffer
		customLogger := slog.New(slog.NewTextHandler(&buf, nil))

		funcMap := FuncMap(WithLogger(customLogger))

		assert.NotNil(t, funcMap)
		assert.GreaterOrEqual(t, len(funcMap), sprigFunctionCount)
	})

	t.Run("works without options", func(t *testing.T) {
		funcMap := FuncMap()

		assert.NotNil(t, funcMap)
		assert.GreaterOrEqual(t, len(funcMap), sprigFunctionCount)
	})
}

func TestHermeticFuncMapsWithLogger(t *testing.T) {
	var buf bytes.Buffer
	customLogger := slog.New(slog.NewTextHandler(&buf, nil))

	t.Run("HermeticTxtFuncMap accepts options", func(t *testing.T) {
		funcMap := HermeticTxtFuncMap(WithLogger(customLogger))
		assert.NotNil(t, funcMap)
	})

	t.Run("HermeticHtmlFuncMap accepts options", func(t *testing.T) {
		funcMap := HermeticHtmlFuncMap(WithLogger(customLogger))
		assert.NotNil(t, funcMap)
	})

	t.Run("TxtFuncMap accepts options", func(t *testing.T) {
		funcMap := TxtFuncMap(WithLogger(customLogger))
		assert.NotNil(t, funcMap)
	})

	t.Run("HtmlFuncMap accepts options", func(t *testing.T) {
		funcMap := HtmlFuncMap(WithLogger(customLogger))
		assert.NotNil(t, funcMap)
	})

	t.Run("GenericFuncMap accepts options", func(t *testing.T) {
		funcMap := GenericFuncMap(WithLogger(customLogger))
		assert.NotNil(t, funcMap)
	})
}
