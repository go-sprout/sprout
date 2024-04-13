package sprout

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFunctionHandler_DefaultValues(t *testing.T) {
	handler := NewFunctionHandler()

	assert.NotNil(t, handler)
	assert.Equal(t, ErrorStrategyReturnDefaultValue, handler.errHandler.strategy)
	assert.NotNil(t, handler.errHandler.errChan)
	assert.NotNil(t, handler.logger)
}

func TestNewFunctionHandler_CustomValues(t *testing.T) {
	errChan := make(chan error, 1)
	logger := slog.New(&slog.TextHandler{})
	handler := NewFunctionHandler(
		WithErrStrategy(ErrorStrategyTemplateError),
		WithLogger(logger),
		WithErrorChannel(errChan),
	)

	assert.NotNil(t, handler)
	assert.Equal(t, ErrorStrategyTemplateError, handler.errHandler.strategy)
	assert.Equal(t, errChan, handler.errHandler.errChan)
	assert.Equal(t, logger, handler.logger)
}

func TestWithErrHandling(t *testing.T) {
	option := WithErrStrategy(ErrorStrategyTemplateError)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, ErrorStrategyTemplateError, handler.errHandler.strategy)
}

func TestWithLogger(t *testing.T) {
	logger := slog.New(&slog.TextHandler{})
	option := WithLogger(logger)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, logger, handler.logger)
	assert.Equal(t, logger, handler.Logger())
}

func TestWithErrorChannel(t *testing.T) {
	errChan := make(chan error, 1)
	option := WithErrorChannel(errChan)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, errChan, handler.errHandler.errChan)
}

func TestWithParser(t *testing.T) {
	fnHandler := &FunctionHandler{
		errHandler: &internalErrorHandler{
			strategy: ErrorStrategyTemplateError,
			errChan:  make(chan error, 1),
		},
		logger: slog.New(&slog.TextHandler{}),
	}
	option := WithFunctionHandler(fnHandler)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, fnHandler, handler)
}

func TestFuncMap_IncludesHello(t *testing.T) {
	funcMap := FuncMap()

	_, exists := funcMap["hello"]
	assert.True(t, exists)

	helloFunc, ok := funcMap["hello"].(func() string)
	assert.True(t, ok)

	assert.Equal(t, "Hello, World!", helloFunc())
}

// This test ensures backward compatibility by checking if FuncMap (the function mentioned in the comment) exists or needs to be implemented for the test.
func TestFuncMap_BackwardCompatibility(t *testing.T) {
	// Assuming FuncMap() is implemented and returns a template.FuncMap
	// Replace the implementation details as per actual FuncMap function.
	genericMap["TestFuncMap_BackwardCompatibility"] = func() string {
		return "example"
	}

	funcMap := FuncMap()
	exampleFunc, exists := funcMap["TestFuncMap_BackwardCompatibility"]
	assert.True(t, exists)

	result, ok := exampleFunc.(func() string)
	assert.True(t, ok)
	assert.Equal(t, "example", result())

	helloFunc, ok := funcMap["hello"].(func() string)
	assert.True(t, ok)
	assert.Equal(t, "Hello, World!", helloFunc())
}
