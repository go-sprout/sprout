package sprout

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

var noopLogger = slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{}))

func TestErrIsPresentNoErrorHandlingConfigured(t *testing.T) {
	handler := &FunctionHandler{logger: noopLogger, errHandler: &internalErrorHandler{strategy: 0}} // Unconfigured
	inputError := errors.New("test error")

	resultErr, resultBool := handler.ErrIsPresent(inputError)

	assert.Equal(t, inputError, resultErr)
	assert.True(t, resultBool)
}

func TestErrIsPresentReturnDefaultValueOnError(t *testing.T) {
	handler := &FunctionHandler{logger: noopLogger, errHandler: &internalErrorHandler{strategy: ErrorStrategyReturnDefaultValue}}
	inputError := errors.New("test error")

	resultErr, resultBool := handler.ErrIsPresent(inputError)

	assert.Nil(t, resultErr)
	assert.True(t, resultBool)
}

func TestErrIsPresentTemplateErrorReturn(t *testing.T) {
	handler := &FunctionHandler{logger: noopLogger, errHandler: &internalErrorHandler{strategy: ErrorStrategyTemplateError}}
	inputError := errors.New("test error")

	resultErr, resultBool := handler.ErrIsPresent(inputError)

	assert.EqualError(t, resultErr, "test error")
	assert.True(t, resultBool)
}

func TestErrIsPresentPanicOnError(t *testing.T) {
	handler := &FunctionHandler{logger: noopLogger, errHandler: &internalErrorHandler{strategy: ErrorStrategyPanic}}
	inputError := errors.New("test error")

	assert.Panics(t, func() {
		err, _ := handler.ErrIsPresent(inputError)
		assert.Nil(t, err)
	}, "The code did not panic")
}

func TestErrIsPresentSendErrorToChannel(t *testing.T) {
	handler := &FunctionHandler{logger: noopLogger, errHandler: &internalErrorHandler{strategy: ErrorStrategyChannel, errChan: make(chan error, 1)}}
	inputError := errors.New("test error")

	_, resultBool := handler.ErrIsPresent(inputError)
	defer close(handler.errHandler.errChan)

	assert.True(t, resultBool)
	assert.Equal(t, inputError, <-handler.errHandler.errChan)
}

func TestErrIsPresentNoErrorPassed(t *testing.T) {
	handler := &FunctionHandler{logger: noopLogger, errHandler: &internalErrorHandler{strategy: ErrorStrategyReturnDefaultValue}}

	resultErr, resultBool := handler.ErrIsPresent(nil)

	assert.Nil(t, resultErr)
	assert.False(t, resultBool)
}

func TestErrIsPresentWithLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{}))
	handler := &FunctionHandler{
		errHandler: &internalErrorHandler{strategy: ErrorStrategyTemplateError},
		logger:     logger,
	}
	inputError := errors.New("test error")

	err, _ := handler.ErrIsPresent(inputError)

	assert.Error(t, err)
	assert.Contains(t, buf.String(), "Error caught")
}

func TestDefaultValueFor(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{"int", 0, 0},
		{"string", "", ""},
		{"bool", false, false},
		{"float64", 0.0, 0.0},
		{"struct", struct{}{}, struct{}{}},
		{"slice", []int{}, []int{}},
		{"map", map[string]string{}, map[string]string{}},
		{"pointer", new(int), new(int)},
		{"nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DefaultValueFor(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestErrFuncCaller(t *testing.T) {
	f := errFuncCaller(1)
	assert.NotNil(t, f)

	name := f.Name
	assert.Contains(t, name, "TestErrFuncCaller")

	// Test that errFuncCaller returns nil when called with a skip value that's too large
	f = errFuncCaller(10000)
	assert.Nil(t, f)

	invalidFunc := func() {
		errFuncCaller(0)
	}

	assert.NotPanics(t, invalidFunc, "The code did not panic")
}
