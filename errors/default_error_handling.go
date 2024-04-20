package errors

import "log/slog"

type DefaultErrorHandler struct {
	logger     *slog.Logger
	subHandler ErrorHandler

	strategy    ErrorStrategy
	previousErr error
}

// WithLogger configures an ErrorHandler to use a specific slog.Logger for
// logging errors. This is useful for integrating the error handler with an
// application's centralized logging system.
func WithLogger(l *slog.Logger) ErrHandlerOption {
	return func(eh ErrorHandler) {
		if ieh, ok := eh.(*DefaultErrorHandler); ok {
			ieh.logger = l
		}
	}
}

// WithPreviousErr configures an ErrorHandler to record a previous error.
// This can be used to maintain error chains where context about prior errors
// is preserved.
func WithPreviousErr(err error) ErrHandlerOption {
	return func(eh ErrorHandler) {
		if ieh, ok := eh.(*DefaultErrorHandler); ok {
			ieh.previousErr = err
		}
	}
}

// WithStrategy configures an ErrorHandler to use a specific error handling
// strategy. This allows for custom error handling behaviors to be defined
// within the ErrorHandler.
func WithStrategy(strategy ErrorStrategy) ErrHandlerOption {
	return func(eh ErrorHandler) {
		if ieh, ok := eh.(*DefaultErrorHandler); ok {
			ieh.strategy = strategy
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
func WithSubHandler(sh ErrorHandler) ErrHandlerOption {
	return func(eh ErrorHandler) {
		if ieh, ok := eh.(*DefaultErrorHandler); ok {
			ieh.subHandler = sh
		}
	}
}

// NewErrHandler creates a new ErrorHandler with optional configurations
// provided by ErrHandlerOption. It defaults to using the
// ErrorStrategyTemplateError unless configured otherwise.
func NewErrHandler(opts ...ErrHandlerOption) ErrorHandler {
	eh := &DefaultErrorHandler{
		strategy: ErrorStrategyTemplateError,
		logger:   slog.New(slog.Default().Handler()),
	}

	for _, opt := range opts {
		opt(eh)
	}

	return eh
}

// Handle checks for the presence of an error and handles it according
// to the configured error handling strategy within the FunctionHandler.
// It also logs the error along with the function and package where it occurred.
//
// This method returns the error and a boolean indicating whether
// the error was present without taking care of error handling.
func (eh *DefaultErrorHandler) Handle(err error, opts ...ErrHandlerOption) error {
	errReturned, _ := eh.err(Cast(err))
	return errReturned
}

// HandleMessage is a shorthand method for triggering an error with a message.
// It is equivalent to calling ErrIsPresent with an error created from the
// provided message.
func (eh *DefaultErrorHandler) HandleMessage(msg string, opts ...ErrHandlerOption) error {
	return eh.Handle(New(msg, eh.previousErr))
}

func (eh *DefaultErrorHandler) err(err Error) (Error, bool) {
	if err != nil {
		// Log the error if a logger is present.
		if eh.logger != nil {
			eh.logger.Error("An error has occured in your template", "error", err.Error())
		}

		// Allow the sub-handler to handle the error. This is useful for
		// chaining error handlers together to perform multiple actions.
		// without losing the original handling strategy.
		if eh.subHandler != nil {
			err = Cast(eh.subHandler.Handle(err))
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
