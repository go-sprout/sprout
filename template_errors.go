package sprout

import (
	"reflect"

	"github.com/42atomys/sprout/errors"
)

// errorChainHandler is an implementation of the errors.ErrorHandler interface
// that captures and stores errors in a chain. It can be used to track all errors
// that occur during the execution of an application.
type errorChainHandler struct {
	errors.DefaultErrorHandler

	// errors holds the list of errors captured by the handler.
	errors []errors.Error
}

// NewErrorChainHandler creates a new instance of errorChainHandler.
// This handler is useful for applications that need to collect and inspect
// errors rather than immediately handling them.
func NewErrorChainHandler() errors.ErrorHandler {
	return &errorChainHandler{
		errors: make([]errors.Error, 0),
	}
}

// Handle appends the provided error to the error chain and returns the error.
// This method implements the ErrorHandler interface and allows for flexible error handling strategies.
func (h *errorChainHandler) Handle(err error, opts ...errors.ErrHandlerOption) error {
	h.errors = append(h.errors, errors.Cast(err))
	return err
}

// DefaultValueFor returns the zero value for any type T.
// This utility is particularly useful in situations where functions need to return
// a default value when an error is encountered but execution must continue.
// It supports basic types, slices, maps, and pointers.
func DefaultValueFor[T interface{}](v T) T {
	typeOf := reflect.TypeOf(v)
	if typeOf == nil {
		return v
	}

	value := reflect.Zero(typeOf).Interface().(T)

	if typeOf.Kind() == reflect.Slice {
		value = reflect.MakeSlice(typeOf, 0, 0).Interface().(T)
	} else if typeOf.Kind() == reflect.Map {
		value = reflect.MakeMap(typeOf).Interface().(T)
	} else if typeOf.Kind() == reflect.Ptr {
		value = reflect.New(typeOf.Elem()).Interface().(T)
	}

	return value
}
