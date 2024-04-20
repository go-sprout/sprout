package errors

import (
	"log/slog"
	"reflect"
	"slices"
)

type DefaultErrorHandler struct {
	logger     *slog.Logger
	subHandler ErrorHandler

	strategy     ErrorStrategy
	maxErrorKeep int
	errors       []Error
}

type handlingOpts struct {
	previousErr error
}

// WithLogger configures an ErrorHandler to use a specific slog.Logger for
// logging errors. This is useful for integrating the error handler with an
// application's centralized logging system.
func WithLogger(l *slog.Logger) HandlerOption {
	return func(eh ErrorHandler) {
		if deh, ok := eh.(*DefaultErrorHandler); ok {
			deh.logger = l
		}
	}
}

// WithStrategy configures an ErrorHandler to use a specific error handling
// strategy. This allows for custom error handling behaviors to be defined
// within the ErrorHandler.
func WithStrategy(strategy ErrorStrategy) HandlerOption {
	return func(eh ErrorHandler) {
		if deh, ok := eh.(*DefaultErrorHandler); ok {
			deh.strategy = strategy
		}
	}
}

// WithSubHandler configures an ErrorHandler to delegate error handling to
// another ErrorHandler. This allows for compositional error handling strategies
// where errors can be processed by multiple handlers, each configured for
// specific tasks.
// This is useful for handling errors in a layered approach, without losing
// the original error handling strategy. To override the strategy, define your
// own ErrorHandler implementation.
func WithSubHandler(sh ErrorHandler) HandlerOption {
	return func(eh ErrorHandler) {
		if deh, ok := eh.(*DefaultErrorHandler); ok {
			deh.subHandler = sh
		}
	}
}

// WithMaxErrors sets the maximum number of errors to keep in memory before
// resetting the error list. This is useful for managing memory usage when
// errors are generated frequently. The default value is 50.
func WithMaxErrors(max int) HandlerOption {
	return func(eh ErrorHandler) {
		if deh, ok := eh.(*DefaultErrorHandler); ok {
			deh.maxErrorKeep = max
		}
	}
}

// NewErrHandler creates a new ErrorHandler with optional configurations
// provided by ErrHandlerOption. It defaults to using the
// ErrorStrategyTemplateError unless configured otherwise.
func NewErrHandler(opts ...HandlerOption) ErrorHandler {
	eh := &DefaultErrorHandler{
		strategy:     ErrorStrategyTemplateError,
		logger:       slog.New(slog.Default().Handler()),
		maxErrorKeep: 50,
		errors:       make([]Error, 0),
	}

	for _, opt := range opts {
		opt(eh)
	}

	return eh
}

func WithPreviousErr(err error) RuntimeOption {
	return func(i interface{}) {
		if opts, ok := i.(*handlingOpts); ok {
			opts.previousErr = err
		}
	}
}

// Handle checks for the presence of an error and handles it according
// to the configured error handling strategy within the FunctionHandler.
// It also logs the error along with the function and package where it occurred.
//
// This method returns the error and a boolean indicating whether
// the error was present without taking care of error handling.
func (eh *DefaultErrorHandler) Handle(err error, opts ...RuntimeOption) (error, bool) {
	if len(eh.errors) >= eh.maxErrorKeep {
		eh.Reset()
	}

	handlingOpts := &handlingOpts{}
	for _, opt := range opts {
		opt(handlingOpts)
	}

	var prev *errorStruct
	if handlingOpts.previousErr != nil {
		prev = castToErrorStruct(handlingOpts.previousErr)
	}

	var errCasted *errorStruct
	if err != nil {
		var ok bool
		if errCasted, ok = err.(*errorStruct); !ok {
			errCasted = Cast(err).(*errorStruct)
			errCasted = errCasted.SetStacklite(defaultStackliteSkip+1, true)
		}
		errCasted.prev = prev

		if !slices.Contains(eh.errors, errCasted.Cause().(Error)) {
			eh.errors = append(eh.errors, errCasted)
		}
	}
	return eh.err(errCasted)
}

// HandleMessage is a shorthand method for triggering an error with a message.
// It is equivalent to calling ErrIsPresent with an error created from the
// provided message.
func (eh *DefaultErrorHandler) HandleMessage(msg string, opts ...RuntimeOption) (error, bool) {
	return eh.Handle(Cast(New(msg)).(*errorStruct).SetStacklite(defaultStackliteSkip+1, true), opts...)
}

func (eh *DefaultErrorHandler) HasErrors() bool {
	return len(eh.errors) > 0
}

func (eh *DefaultErrorHandler) Errors() []Error {
	return eh.errors
}

func (eh *DefaultErrorHandler) Reset() {
	eh.errors = make([]Error, 0)
}

func (eh *DefaultErrorHandler) err(err Error) (errReturned Error, ok bool) {
	if !reflect.ValueOf(err).IsNil() {
		errReturned = err
		ok = true

		// Log the error if a logger is present.
		if eh.logger != nil {
			if st, okStack := err.(Stackliteable); okStack && st.Stacklite() != nil {
				eh.logger.Error(
					"An error has occured in your template",
					"error", err.Err().Error(),
					"package", st.Stacklite().Package,
					"function", st.Stacklite().Function,
					"file", st.Stacklite().File,
					"line", st.Stacklite().Line,
				)
			} else {
				eh.logger.Error("An error has occured in your template", "error", err.Error())
			}
		}

		// Allow the sub-handler to handle the error. This is useful for
		// chaining error handlers together to perform multiple actions.
		// without losing the original handling strategy.
		if eh.subHandler != nil {
			var e error
			e, ok = eh.subHandler.Handle(err)
			errReturned = Cast(e)
		}

		switch eh.strategy {
		case ErrorStrategyTemplateError:
			return errReturned, ok // Return the error to the caller.
		case ErrorStrategyReturnDefaultValue:
			return nil, ok // Ignore the error and proceed with the default value.
		}
	}

	return err, ok
}
