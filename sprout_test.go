package sprout

import (
	"log/slog"
	"testing"

	"github.com/42atomys/sprout/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewFunctionHandler_DefaultValues(t *testing.T) {
	handler := NewFunctionHandler()

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.logger)
}

func TestNewFunctionHandler_CustomValues(t *testing.T) {
	logger := slog.New(&slog.TextHandler{})
	handler := NewFunctionHandler(
		WithLogger(logger),
	)

	assert.NotNil(t, handler)
	assert.Equal(t, logger, handler.logger)
}

type testErrHandler struct {
	errors.DefaultErrorHandler
}

var errTestErrhandler = errors.New("test with error handler error")

func (h *testErrHandler) Handle(err error, opts ...errors.ErrHandlerOption) error {
	return errors.Cast(errTestErrhandler, err)
}

func TestWithErrHandling(t *testing.T) {
	option := WithErrHandler(&testErrHandler{})

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.NotNil(t, handler.errHandler)
	assert.Equal(t, errTestErrhandler, handler.errHandler.Handle(nil))
}

func TestWithLogger(t *testing.T) {
	logger := slog.New(&slog.TextHandler{})
	option := WithLogger(logger)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, logger, handler.logger)
	assert.Equal(t, logger, handler.Logger())
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
