package errors

// errorChainHandler is an implementation of the errors.ErrorHandler interface
// that captures and stores errors in a chain. It can be used to track all errors
// that occur during the execution of an application.
type errorChainHandler struct {
	*DefaultErrorHandler

	// errors holds the list of errors captured by the handler.
	errors []Error
}

// NewErrorChainHandler creates a new instance of errorChainHandler.
// This handler is useful for applications that need to collect and inspect
// errors rather than immediately handling them.
func NewErrorChainHandler(opts ...ErrHandlerOption) ErrorHandler {
	dh, ok := NewErrHandler(opts...).(*DefaultErrorHandler)
	if !ok {
		// Never happens, but just in case
		panic("error chain handler must be used with the default error handler")
	}

	return &errorChainHandler{
		errors:              make([]Error, 0),
		DefaultErrorHandler: dh,
	}
}

// Handle appends the provided error to the error chain and returns the error.
// This method implements the ErrorHandler interface and allows for flexible error handling strategies.
func (h *errorChainHandler) Handle(err error, opts ...ErrHandlerOption) error {
	e := Cast(err)

	if sl, ok := e.(Stackliteable); ok {
		e = sl.WithFrameSkip(4).(Error)
	}

	h.errors = append(h.errors, e)
	return h.DefaultErrorHandler.Handle(e)
}
