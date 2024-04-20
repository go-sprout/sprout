package sprout

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewErrorChainHandler tests the creation of a new errorChainHandler.
func TestNewErrorChainHandler(t *testing.T) {
	handler := NewErrorChainHandler()
	assert.NotNil(t, handler, "Handler should not be nil")
	assert.IsType(t, &errorChainHandler{}, handler, "Handler should be of type *errorChainHandler")
}

// TestErrorChainHandler_Handle tests the error handling and chaining functionality.
func TestErrorChainHandler_Handle(t *testing.T) {
	handler := NewErrorChainHandler().(*errorChainHandler)
	testError := errors.New("test error")
	errReturned := handler.Handle(testError)

	assert.Equal(t, 1, len(handler.errors), "There should be one error in the chain")
	assert.Equal(t, testError, handler.errors[0].Err(), "The error in the chain should match the handled error")
	assert.Equal(t, testError, errReturned, "The returned error should match the handled error")
}

// TestDefaultValueFor tests the DefaultValueFor function for various types.
func TestDefaultValueFor(t *testing.T) {
	assert.Equal(t, 0, DefaultValueFor[int](0), "Default value for int should be 0")
	assert.Equal(t, "", DefaultValueFor[string](""), "Default value for string should be an empty string")
	assert.Equal(t, ([]int)(nil), DefaultValueFor[[]int](nil), "Default value for slice should be nil")

	type customStruct struct {
		Field string
	}
	assert.Equal(t, customStruct{}, DefaultValueFor[customStruct](customStruct{}), "Default value for struct should be zero valued struct")
	assert.Nil(t, DefaultValueFor[map[string]int](nil), "Default value for map should be nil")
	assert.Nil(t, DefaultValueFor[*int](nil), "Default value for pointer should be nil")
}

// Running this test will verify that DefaultValueFor correctly returns zero values for specified types.
