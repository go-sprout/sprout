package errors

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
