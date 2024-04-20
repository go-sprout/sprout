package errors

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewErrHandler checks the initialization of an internalErrorHandler with default and custom configurations.
func TestNewErrHandler(t *testing.T) {
	// Test default initialization
	handler := NewErrHandler().(*DefaultErrorHandler)
	assert.NotNil(t, handler.logger, "Logger should be initialized by default")
	assert.Equal(t, ErrorStrategyTemplateError, handler.strategy, "Default strategy should be ErrorStrategyTemplateError")

	// Test custom initialization with options
	customLogger := slog.New(slog.Default().Handler())
	handler = NewErrHandler(WithLogger(customLogger)).(*DefaultErrorHandler)
	assert.Equal(t, customLogger, handler.logger, "Custom logger should be set")
}

// TestDefaultHandler_Handle tests the error handling functionality with different strategies.
func TestDefaultHandler_Handle(t *testing.T) {
	err := Cast(errors.New("test error"))

	// Test with the default error strategy
	handler := NewErrHandler().(*DefaultErrorHandler)
	result, ok := handler.Handle(err)
	assert.True(t, ok, "Error should be returned with the default strategy")
	assert.Equal(t, err, result, "Error should be returned with the default strategy")
	assert.ErrorIs(t, err, result)

	// Test with the return default value strategy
	handler = NewErrHandler(WithStrategy(ErrorStrategyReturnDefaultValue)).(*DefaultErrorHandler)

	result, ok = handler.Handle(err)
	assert.True(t, ok, "Error should be returned with the return default value strategy")
	assert.Nil(t, result, "No error should be returned with return default value strategy")

	// Test with a nil error
	result, ok = handler.Handle(nil)
	assert.False(t, ok, "No error should be returned with a nil error")
	assert.Nil(t, result, "No error should be returned with a nil error")
}

// TestDefaultHandler_HandleMessage tests the handling of error messages.
func TestDefaultHandler_HandleMessage(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	msg := "error message"
	err, ok := handler.HandleMessage(msg)
	assert.True(t, ok, "Error should be returned with the default strategy")
	assert.Contains(t, err.Error(), msg, "Error message should be part of the returned error")

	// Test with a previous error
	prevErr := errors.New("previous error")
	errReturned, ok := handler.Handle(err, WithPreviousErr(prevErr))
	assert.True(t, ok, "Error should be returned with the previous error")
	assert.ErrorIs(t, errReturned, prevErr)

}

// TestDefaultHandler_WithOptions tests the application of options on runtime.
func TestDefaultHandler_WithOptions(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	newLogger := slog.New(slog.Default().Handler())
	WithLogger(newLogger)(handler)
	assert.Equal(t, newLogger, handler.logger, "Logger should be updated with new logger")
}

type testHandler struct {
	DefaultErrorHandler
}

var errTestSubHandler = errors.New("sub handler error")

func (h *testHandler) Handle(err error, opts ...RuntimeOption) (error, bool) {
	return Cast(errTestSubHandler, err), err != nil
}

// TestDefaultHandler_WithSubHandler tests the delegation of error handling to another handler.
func TestDefaultHandler_WithSubHandler(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	subHandler := &testHandler{}
	WithSubHandler(subHandler)(handler)
	assert.Equal(t, subHandler, handler.subHandler, "Sub handler should be set")

	// Test that the sub handler is used for error handling
	err := errors.New("test error")
	result, ok := handler.Handle(err)
	assert.True(t, ok, "Error should be returned with the sub handler")
	assert.ErrorIs(t, result, errTestSubHandler)
}

func TestDefaultHandler_Errors(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	testError := errors.New("test error")
	errReturned, ok := handler.Handle(testError)

	assert.True(t, ok, "The error should be present")
	assert.Equal(t, 1, len(handler.errors), "There should be one error in the chain")
	assert.Equal(t, 1, len(handler.Errors()), "There should be one error in the chain")
	assert.True(t, handler.HasErrors(), "There should be errors in the chain")
	assert.Equal(t, testError, handler.errors[0].Err(), "The error in the chain should match the handled error")
	assert.ErrorIs(t, errReturned, testError, "The returned error should match the handled error")

	handler.Reset()
	assert.Equal(t, 0, len(handler.errors), "There should be no errors in the chain")
	assert.Equal(t, 0, len(handler.Errors()), "There should be no errors in the chain")
	assert.False(t, handler.HasErrors(), "There should be no errors in the chain")
}

func TestDefaultHandler_Handle_WithPreviousErr(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	prevErr := errors.New("previous error")

	errReturned, _ := handler.Handle(errors.New("test error"), WithPreviousErr(prevErr))
	assert.ErrorIs(t, errReturned, prevErr, "Previous error should be set")
}

func TestDefaultHandler_WithStrategy(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	strategy := ErrorStrategyReturnDefaultValue
	WithStrategy(strategy)(handler)
	assert.Equal(t, strategy, handler.strategy, "Strategy should be set")
}

func TestDefaultHandler_Stackline(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	errReturned, ok := handler.Handle(errors.New("test error"))
	assert.True(t, ok, "The error should be present")

	if st, ok := errReturned.(Stackliteable); assert.True(t, ok) {
		assert.Contains(t, st.Stacklite().Package, "errors")
		assert.Contains(t, st.Stacklite().Function, "TestDefaultHandler_Stackline")
		assert.Contains(t, st.Stacklite().File, "default_error_handling_test.go")
		assert.Greater(t, st.Stacklite().Line, 0)
	}

	errReturned, ok = handler.HandleMessage("test error")
	assert.True(t, ok, "The error should be present")

	if st, ok := errReturned.(Stackliteable); assert.True(t, ok) {
		assert.Contains(t, st.Stacklite().Package, "errors")
		assert.Contains(t, st.Stacklite().Function, "TestDefaultHandler_Stackline")
		assert.Contains(t, st.Stacklite().File, "default_error_handling_test.go")
		assert.Greater(t, st.Stacklite().Line, 0)
	}
}
