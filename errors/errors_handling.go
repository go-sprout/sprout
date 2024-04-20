package errors

// ErrorHandler defines an interface for handling errors.
// It offers methods to handle raw errors and string messages as errors,
// with support for custom handler options.
type ErrorHandler interface {

	// Handle processes an error according to implemented strategies and options.
	// It returns the error and a boolean indicating if the error was present. This
	// double return value is useful for checking if an error occurred without
	// taking care of error handling.
	Handle(err error, opts ...RuntimeOption) (error, bool)

	// HandleMessage processes a string message as an error, applying the same
	// strategies and options as Handle if wanted.
	HandleMessage(msg string, opts ...RuntimeOption) (error, bool)
}

// HandlerOption defines a type for functional options that configure
// ErrorHandler instances.
type HandlerOption func(ErrorHandler)

// RuntimeOption defines a type for functional options that configure
// the handling of each handled error.
type RuntimeOption func(interface{})

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
