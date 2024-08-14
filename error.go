package sprout

// ErrHandling defines the strategy for handling errors within FunctionHandler.
// It supports returning default values, panicking, or sending errors to a
// specified channel.
type ErrHandling int

const (
	// ErrHandlingReturnDefaultValue indicates that a default value should be
	// returned on error (default).
	ErrHandlingReturnDefaultValue ErrHandling = iota + 1
	// ErrHandlingPanic indicates that a panic should be raised on error.
	ErrHandlingPanic
	// ErrHandlingErrorChannel indicates that errors should be sent to an error
	// channel.
	ErrHandlingErrorChannel
)

// WithErrHandling sets the error handling strategy for a FunctionHandler.
func WithErrHandling(eh ErrHandling) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) {
		p.ErrHandling = eh
	}
}

// WithErrorChannel sets the error channel for a FunctionHandler.
func WithErrorChannel(ec chan error) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) {
		p.errChan = ec
	}
}
