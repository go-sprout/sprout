package errors_test

import (
	goError "errors"
	"testing"

	"github.com/42atomys/sprout/errors"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationOfErrors(t *testing.T) {
	rootErr := errors.New("root error")
	midErr := errors.New("mid error", rootErr)
	finalErr := errors.New("final error", midErr)

	// Test that the final error wraps the mid error
	assert.ErrorIs(t, finalErr, midErr)

	// Test that the final error indirectly wraps the root error
	assert.ErrorIs(t, finalErr, rootErr)

	// Ensure that unwrapping works as expected
	assert.Equal(t, goError.New("mid error"), goError.Unwrap(finalErr))
	assert.Equal(t, goError.New("root error"), goError.Unwrap(midErr))

	// Ensure a casted error is still the same error
	errUnsupportedCasted := errors.Cast(goError.ErrUnsupported)
	assert.NotNil(t, errUnsupportedCasted)
	assert.ErrorIs(t, errUnsupportedCasted, goError.ErrUnsupported)

	errTest := errors.New("test error", errUnsupportedCasted)
	assert.ErrorIs(t, errTest, errTest)
	assert.True(t, errTest.Is(errTest))
	assert.Equal(t, goError.ErrUnsupported, goError.Unwrap(errTest))

	if assert.Len(t, errTest.Stack(), 2) {
		// standard error
		assert.Contains(t, errTest.Stack()[0], "unsupported operation")
		// augmented error
		assert.Contains(t, errTest.Stack()[1], "errors_integration_test")
		assert.Contains(t, errTest.Stack()[1], "TestIntegrationOfErrors")
		assert.Contains(t, errTest.Stack()[1], "test error")
	}
}
