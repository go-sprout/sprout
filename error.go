package sprout

import (
	"reflect"
	"runtime"
)

// ErrHandling defines the strategy for handling errors within FunctionHandler.
// It supports returning default values, panicking, or sending errors to a
// specified channel.
type ErrHandling int

const (
	// ErrHandlingReturnDefaultValue indicates that a default value should be
	// returned on error (default).
	ErrHandlingReturnDefaultValue ErrHandling = iota + 1
	// ErrHandlingTemplateError indicates that an error should be returned on error
	// following the text/template package behavior.
	ErrHandlingTemplateError
	// ErrHandlingPanic indicates that a panic should be raised on error.
	ErrHandlingPanic
	// ErrHandlingErrorChannel indicates that errors should be sent to an error
	// channel.
	ErrHandlingErrorChannel
)

// ErrIsPresent checks for the presence of an error and handles it according
// to the configured error handling strategy within the FunctionHandler.
// It also logs the error along with the function and package where it occurred.
//
// This method returns the error and a boolean indicating whether
// the error was present without taking care of error handling.
func (fh *FunctionHandler) ErrIsPresent(err error) (error, bool) {
	if err != nil {
		fc := errFuncCaller(2) // Get the caller function to log where the error occurred.
		if fh.Logger != nil && fc != nil {
			fh.Logger.Error("Error caught", "error", err.Error(), "function", fc.Name(), "package", fc.Entry())
		}

		switch fh.ErrHandling {
		case ErrHandlingPanic:
			panic(err) // Panic if configured to do so.
		case ErrHandlingTemplateError:
			return err, true // Return the error to the caller.
		case ErrHandlingErrorChannel:
			if fh.errChan != nil {
				fh.errChan <- err // Send the error to an error channel.
			}
			return nil, true
		case ErrHandlingReturnDefaultValue:
			return nil, true // Ignore the error and proceed with the default value.
		}
	}
	return err, err != nil
}

// DefaultValueFor returns the zero value for a given type T. This is useful
// for returning default values in functions where an error occurs but
// execution must continue.
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

// errFuncCaller uses the runtime package to find the function that called it,
// allowing for detailed logging and error handling. 'skip' levels are bypassed
// to find the actual caller.
// It returns a *runtime.Func representing the caller, or nil if not found.
func errFuncCaller(skip int) *runtime.Func {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}

	return runtime.FuncForPC(pc)
}
