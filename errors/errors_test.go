package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testError struct {
	msg    string
	parent *testError
}

func (e *testError) Error() string {
	return e.msg
}

func (e *testError) Err() error {
	return e
}

func (e *testError) Unwrap() error {
	if e.parent == nil {
		return nil
	}
	return e.parent
}

func (e *testError) Is(target error) bool {
	return e == target
}

func (e *testError) Stack() []string {
	return []string{}
}

func TestErrorChaining(t *testing.T) {
	rootErr := New("root error")
	midErr := New("mid error", rootErr)
	finalErr := New("final error", midErr)

	// Test that the final error wraps the mid error
	assert.ErrorIs(t, finalErr, midErr)

	// Test that the final error indirectly wraps the root error
	assert.ErrorIs(t, finalErr, rootErr)

	// Ensure that unwrapping works as expected
	assert.Equal(t, errors.New("mid error"), errors.Unwrap(finalErr))
	assert.Equal(t, errors.New("root error"), errors.Unwrap(midErr))
}

func TestCastNilError(t *testing.T) {
	var err error // nil error
	castedErr := Cast(err)
	assert.Nil(t, castedErr, "Casting a nil error should return nil")
}

func TestCastStandardError(t *testing.T) {
	stdErr := errors.New("standard error")
	castedErr := Cast(stdErr)
	assert.NotNil(t, castedErr, "Casting a standard error should not return nil")
	assert.NotEqual(t, stdErr, castedErr, "Casted error should not be the same as the input error")
	assert.Implements(t, (*Error)(nil), castedErr, "Casted error should implement the Error interface")
	assert.Equal(t, stdErr.Error(), castedErr.Err().Error(), "Casted error message should match the original error message")
}

func TestCastCustomError(t *testing.T) {
	customErr := New("custom error")
	castedErr := Cast(customErr)
	assert.Equal(t, customErr, castedErr, "Casting an error that is already of type Error should return the same error")
}

func TestErrorOutput(t *testing.T) {
	rootErr := New("root error")
	err := New("custom error", rootErr)

	// The error output should contain the stack info and both error messages
	errStr := err.Error()
	assert.Contains(t, errStr, "custom error")

	assert.Contains(t, errStr, "root error")
	assert.Contains(t, errStr, "errors_test.go") // Check for stack info presence
}

func TestErrorCause(t *testing.T) {
	var err *errorStruct
	assert.Nil(t, err.Cause(), "Calling Cause on nil should return nil")

	root := &errorStruct{
		err: errors.New("root cause"),
	}

	middle := &errorStruct{
		prev: root,
		err:  errors.New("middle error"),
	}
	end := &errorStruct{
		prev: middle,
		err:  errors.New("end of chain"),
	}

	assert.Equal(t, root, root.Cause(), "The root of the error chain should be returned as the cause")
	assert.Equal(t, root, middle.Cause(), "The root of the error chain should be returned as the cause")
	assert.Equal(t, root, end.Cause(), "The root of the error chain should be returned as the cause")
}

func TestNilErrorHandling(t *testing.T) {
	// Creating an error with no previous error to simulate edge cases
	err := New("nil previous error")

	// Error should still behave correctly
	assert.NotNil(t, err)
	assert.Nil(t, errors.Unwrap(err))
	assert.Contains(t, err.Error(), "nil previous error")
}

func TestFunctionErrorHandling(t *testing.T) {
	// This test simulates a function failing and being wrapped by our error handling
	funcThatFails := func() (interface{}, error) {
		return nil, New("failure")
	}

	tryFunc := func(f func() (interface{}, error)) func() error {
		return func() error {
			_, err := f()
			return err
		}
	}

	wrappedErr := tryFunc(funcThatFails)()
	assert.NotNil(t, wrappedErr)
	assert.Contains(t, wrappedErr.(Error).Error(), "failure")
}

func TestPreviousAndNextError(t *testing.T) {
	var err *errorStruct
	assert.Nil(t, err.Cause(), "Calling Cause on nil should return nil")

	rootErr := New("root error").(*errorStruct)
	midErr := New("mid error", rootErr).(*errorStruct)
	finalErr := New("final error", midErr).(*errorStruct)

	assert.Equal(t, rootErr, finalErr.Cause())
	assert.Equal(t, midErr, finalErr.Prev())
	assert.Equal(t, rootErr, midErr.Prev())

	cause := finalErr.Cause()
	if assert.Equal(t, midErr, cause.Next()) {
		if assert.Equal(t, finalErr, cause.Next().Next()) {
			assert.Nil(t, cause.Next().Next().Next())
		}
	}
}

func TestErrorIntegrationWithGoError(t *testing.T) {
	goErr := errors.New("standard error")
	err := New("custom error", goErr)

	assert.NotEqual(t, err, goErr)
	assert.ErrorIs(t, err, goErr)
	assert.NotErrorIs(t, goErr, err)

	assert.Equal(t, goErr, errors.Unwrap(err))
	assert.NotEqual(t, err, errors.Unwrap(goErr))

	tErr := &testError{msg: "test error"}
	err = New("custom error", tErr)

	assert.ErrorIs(t, err, tErr)
	assert.NotErrorIs(t, tErr, err)

	assert.Equal(t, tErr, errors.Unwrap(err))
	assert.NotEqual(t, err, errors.Unwrap(tErr))

	assert.NotErrorIs(t, err, errors.New("impossible"))
	assert.NotErrorIs(t, tErr, errors.New("impossible"))
}

func TestStack(t *testing.T) {
	err := New("test error", New("mid error", New("root error")))

	stack := err.Stack()
	assert.NotNil(t, stack)
	assert.NotEmpty(t, stack)

	if assert.Len(t, stack, 3) {
		assert.Contains(t, stack[0], "root error")
		assert.Contains(t, stack[1], "mid error")
		assert.Contains(t, stack[2], "test error")
	}
}
