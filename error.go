package sprout

import (
	"errors"
	"path"
	"reflect"
	"runtime"
	"strings"
)

// ErrorStrategy defines the strategy for handling errors within FunctionHandler.
// It supports returning default values, panicking, or sending errors to a
// specified channel.
type ErrorStrategy int

const (
	// ErrorStrategyTemplateError indicates that an error should be returned on error
	// following the text/template package behavior (default).
	ErrorStrategyTemplateError ErrorStrategy = iota + 1
	// ErrorStrategyReturnDefaultValue indicates that a default value should be
	// returned on error.
	ErrorStrategyReturnDefaultValue
)

// internalErrorHandler manages error handling within FunctionHandler
// internally. It is not intended for external use.
type internalErrorHandler struct {
	strategy ErrorStrategy
	errChan  chan error
}

// ErrIsPresent checks for the presence of an error and handles it according
// to the configured error handling strategy within the FunctionHandler.
// It also logs the error along with the function and package where it occurred.
//
// This method returns the error and a boolean indicating whether
// the error was present without taking care of error handling.
func (fh *FunctionHandler) ErrIsPresent(err error) (error, bool) {
	return fh.err(fh.errHandler, err)
}

// ErrTrigger is a shorthand method for triggering an error with a message.
// It is equivalent to calling ErrIsPresent with an error created from the
// provided message.
func (fh *FunctionHandler) ErrTrigger(errMessage string) (err error) {
	err, _ = fh.err(fh.errHandler, errors.New(errMessage))
	return
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

// createInternalErrorHandler creates a new error handler with default values.
func createInternalErrorHandler() *internalErrorHandler {
	eh := &internalErrorHandler{
		strategy: ErrorStrategyTemplateError,
		errChan:  make(chan error),
	}

	return eh
}

// err checks for the presence of an error and handles it according
// to the configured error handling strategy within the FunctionHandler.
// It also logs the error along with the function and package where it occurred.
//
// This method returns the error and a boolean indicating whether
// the error was present without taking care of error handling.
func (fh *FunctionHandler) err(eh *internalErrorHandler, err error) (error, bool) {
	if err != nil {
		// Get the caller function to log where the error occurred.
		fc := errFuncCaller(3)

		// Log the error if a logger is present.
		if fh.logger != nil && fc != nil {
			fh.logger.Error("Error caught", "error", err.Error(), "function", fc.Name, "file", fc.File, "line", fc.Line)
		}

		// Send the error to an error channel.
		if eh.errChan != nil {
			eh.errChan <- err
		}

		switch eh.strategy {
		case ErrorStrategyTemplateError:
			return err, true // Return the error to the caller.
		case ErrorStrategyReturnDefaultValue:
			return nil, true // Ignore the error and proceed with the default value.
		}
	}
	return err, err != nil
}

// errFuncCaller uses the runtime package to find the function that called it,
// allowing for detailed logging and error handling. 'skip' levels are bypassed
// to find the actual caller.
// It returns a *runtime.Func representing the caller, or nil if not found.
func errFuncCaller(skip int) *struct {
	Name string
	File string
	Line int
} {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}

	fn := runtime.FuncForPC(pc)
	ss := strings.Split(fn.Name(), ".")

	file, line := fn.FileLine(pc)
	_, file = path.Split(file)
	return &struct {
		Name string
		File string
		Line int
	}{ss[len(ss)-1], file, line}
}
