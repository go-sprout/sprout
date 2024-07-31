package sprout

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFunctionHandler_DefaultValues(t *testing.T) {
	handler := NewFunctionHandler()

	assert.NotNil(t, handler)
	assert.Equal(t, ErrHandlingReturnDefaultValue, handler.ErrHandling)
	assert.NotNil(t, handler.errChan)
	assert.NotNil(t, handler.Logger)
}

func TestNewFunctionHandler_CustomValues(t *testing.T) {
	errChan := make(chan error, 1)
	logger := slog.New(&slog.TextHandler{})
	handler := NewFunctionHandler(
		WithErrHandling(ErrHandlingPanic),
		WithLogger(logger),
		WithErrorChannel(errChan),
	)

	assert.NotNil(t, handler)
	assert.Equal(t, ErrHandlingPanic, handler.ErrHandling)
	assert.Equal(t, errChan, handler.errChan)
	assert.Equal(t, logger, handler.Logger())
}

func TestWithErrHandling(t *testing.T) {
	option := WithErrHandling(ErrHandlingPanic)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, ErrHandlingPanic, handler.ErrHandling)
}

func TestWithLogger(t *testing.T) {
	logger := slog.New(&slog.TextHandler{})
	option := WithLogger(logger)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, logger, handler.Logger())
}

func TestWithErrorChannel(t *testing.T) {
	errChan := make(chan error, 1)
	option := WithErrorChannel(errChan)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, errChan, handler.errChan)
}

func TestWithParser(t *testing.T) {
	fnHandler := &DefaultHandler{
		ErrHandling: ErrHandlingErrorChannel,
		logger:      slog.New(&slog.TextHandler{}),
		errChan:     make(chan error, 1),
	}
	option := WithHandler(fnHandler)

	handler := New()
	option(handler) // Apply the option

	assert.Equal(t, fnHandler, handler)
}

func TestWithNilHandler(t *testing.T) {
	fnHandler := &DefaultHandler{
		ErrHandling: ErrHandlingErrorChannel,
		logger:      slog.New(&slog.TextHandler{}),
		errChan:     make(chan error, 1),
	}
	option := WithHandler(nil)

	beforeApply := fnHandler
	option(beforeApply)

	assert.Equal(t, beforeApply, fnHandler)
}
