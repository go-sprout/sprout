package sprout

import (
	"errors"
	"fmt"
)

// ErrConvertFailed is an error message when converting a value to another type
// fails. You can create a new error message using NewErrConvertFailed.
var ErrConvertFailed = errors.New("failed to convert")

// ErrRecoverPanic are an utility function to recover panic from a function and
// set the error message to unsure no panic is thrown in the template engine.
//
// This is very useful when you are calling a function that might panic and you
// want to catch the panic and return an error message instead like when you use
// an external package that might panic (e.g.: yaml package).
func ErrRecoverPanic(err *error, errMessage string) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("%s: %v", errMessage, r)
	}
}

// NewErrConvertFailed is an utility function to create a new error message when
// converting a value to another type fails. Th error generated will contain the
// type of the value, the value itself, and the error message.
//
// You can check if the error is an ErrConvertFailed using [errors.Is].
func NewErrConvertFailed(typ string, value any, err error) error {
	return fmt.Errorf("%w: %v to %s: %w", ErrConvertFailed, value, typ, err)
}
