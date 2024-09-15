package sprout

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_DefaultValues(t *testing.T) {
	handler := New()

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.Logger)
}

func TestNew_CustomValues(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	handler := New(
		WithLogger(logger),
	)

	assert.NotNil(t, handler)
	assert.Equal(t, logger, handler.Logger())
}

func TestWithLogger(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	option := WithLogger(logger)

	handler := New()
	require.NoError(t, option(handler)) // Apply the option

	assert.Equal(t, logger, handler.Logger())
}

func TestWithParser(t *testing.T) {
	fnHandler := &DefaultHandler{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	option := WithHandler(fnHandler)

	handler := New()
	require.NoError(t, option(handler)) // Apply the option

	assert.Equal(t, fnHandler, handler)
}

func TestWithNilHandler(t *testing.T) {
	fnHandler := &DefaultHandler{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	option := WithHandler(nil)

	beforeApply := fnHandler
	require.NoError(t, option(beforeApply)) // Apply the option

	assert.Equal(t, beforeApply, fnHandler)
}

func TestWithSafeFuncs(t *testing.T) {
	handler := New(WithSafeFuncs(true))
	assert.True(t, handler.wantSafeFuncs)

	handler.cachedFuncsMap["test"] = func() {}
	funcCount := len(handler.RawFunctions())
	handler.Build()

	assert.Len(t, handler.cachedFuncsMap, funcCount*2)

	var keys []string
	for k := range handler.RawFunctions() {
		keys = append(keys, k)
	}

	assert.Contains(t, keys, "test")
	assert.Contains(t, keys, "safeTest")
}
