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
	customError := errors.New("previous error")
	handler = NewErrHandler(
		WithLogger(customLogger),
		WithPreviousErr(customError),
	).(*DefaultErrorHandler)
	assert.Equal(t, customLogger, handler.logger, "Custom logger should be set")
	assert.Equal(t, customError, handler.previousErr, "Previous error should be set")
}

// TestInternalErrorHandler_Handle tests the error handling functionality with different strategies.
func TestInternalErrorHandler_Handle(t *testing.T) {
	err := Cast(errors.New("test error"))

	// Test with the default error strategy
	handler := NewErrHandler().(*DefaultErrorHandler)
	result := handler.Handle(err)
	assert.Equal(t, err, result, "Error should be returned with the default strategy")
	assert.ErrorIs(t, err, result)

	// Test with the return default value strategy
	handler = NewErrHandler(func(eh ErrorHandler) {
		if ieh, ok := eh.(*DefaultErrorHandler); ok {
			ieh.strategy = ErrorStrategyReturnDefaultValue
		}
	}).(*DefaultErrorHandler)
	result = handler.Handle(err)
	assert.Nil(t, result, "No error should be returned with return default value strategy")

	// Test with a nil error
	result = handler.Handle(nil)
	assert.Nil(t, result, "No error should be returned with a nil error")
}

// TestInternalErrorHandler_HandleMessage tests the handling of error messages.
func TestInternalErrorHandler_HandleMessage(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	msg := "error message"
	err := handler.HandleMessage(msg)
	assert.Contains(t, err.Error(), msg, "Error message should be part of the returned error")

	// Test with a previous error
	prevErr := errors.New("previous error")
	errReturned := handler.Handle(err, WithPreviousErr(prevErr))
	assert.ErrorIs(t, errReturned, prevErr)

}

// TestInternalErrorHandler_WithOptions tests the application of options on runtime.
func TestInternalErrorHandler_WithOptions(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	newLogger := slog.New(slog.Default().Handler())
	WithLogger(newLogger)(handler)
	assert.Equal(t, newLogger, handler.logger, "Logger should be updated with new logger")
}

type testHandler struct {
	DefaultErrorHandler
}

var errTestSubHandler = errors.New("sub handler error")

func (h *testHandler) Handle(err error, opts ...ErrHandlerOption) error {
	return Cast(errTestSubHandler, err)
}

// TestInternalErrorHandler_WithSubHandler tests the delegation of error handling to another handler.
func TestInternalErrorHandler_WithSubHandler(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	subHandler := &testHandler{}
	WithSubHandler(subHandler)(handler)
	assert.Equal(t, subHandler, handler.subHandler, "Sub handler should be set")

	// Test that the sub handler is used for error handling
	err := errors.New("test error")
	result := handler.Handle(err)
	assert.ErrorIs(t, result, errTestSubHandler)
}

func TestInternalErrorHandler_WithPreviousErr(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	prevErr := errors.New("previous error")
	WithPreviousErr(prevErr)(handler)
	assert.Equal(t, prevErr, handler.previousErr, "Previous error should be set")
}

func TestInternalErrorHandler_WithStrategy(t *testing.T) {
	handler := NewErrHandler().(*DefaultErrorHandler)
	strategy := ErrorStrategyReturnDefaultValue
	WithStrategy(strategy)(handler)
	assert.Equal(t, strategy, handler.strategy, "Strategy should be set")
}
