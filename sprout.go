package sprout

import (
	"log/slog"

	"github.com/go-sprout/sprout/registry"
)

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

// FunctionHandler manages function execution with configurable error handling
// and logging.
type FunctionHandler struct {
	ErrHandling ErrHandling
	errChan     chan error
	logger      *slog.Logger
	funcsMap    registry.FunctionMap
	funcsAlias  FunctionAliasMap
}

// FunctionHandlerOption defines a type for functional options that configure
// FunctionHandler.
type FunctionHandlerOption func(*FunctionHandler)

// NewFunctionHandler creates a new FunctionHandler with the provided options.
func NewFunctionHandler(opts ...FunctionHandlerOption) *FunctionHandler {
	fnHandler := &FunctionHandler{
		ErrHandling: ErrHandlingReturnDefaultValue,
		errChan:     make(chan error),
		logger:      slog.New(&slog.TextHandler{}),
		funcsMap:    make(registry.FunctionMap),
		funcsAlias:  make(FunctionAliasMap),
	}

	for _, opt := range opts {
		opt(fnHandler)
	}

	return fnHandler
}

// WithErrHandling sets the error handling strategy for a FunctionHandler.
func WithErrHandling(eh ErrHandling) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.ErrHandling = eh
	}
}

// WithLogger sets the logger used by a FunctionHandler.
func WithLogger(l *slog.Logger) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.logger = l
	}
}

// WithErrorChannel sets the error channel for a FunctionHandler.
func WithErrorChannel(ec chan error) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.errChan = ec
	}
}

// WithFunctionHandler updates a FunctionHandler with settings from another FunctionHandler.
// This is useful for copying configurations between handlers.
func WithFunctionHandler(new registry.Handler) FunctionHandlerOption {
	return func(fnh *FunctionHandler) {
		if new == nil {
			return
		}

		if fhCast, ok := new.(*FunctionHandler); ok {
			*fnh = *fhCast
		}
	}
}

// RegisterHandler registers a single FunctionRegistry implementation (e.g., a handler)
// into the FunctionHandler's internal function registry. This method allows for integrating
// additional functions into the template processing environment.
func (fh *FunctionHandler) AddRegistry(registry registry.Registry) error {
	registry.LinkHandler(fh)
	registry.RegisterFunctions(fh.funcsMap)
	return nil
}

// RegisterHandlers registers multiple FunctionRegistry implementations into the
// FunctionHandler's internal function registry. This method simplifies the process
// of adding multiple sets of functionalities into the template engine at once.
func (fh *FunctionHandler) AddRegistries(registries ...registry.Registry) error {
	for _, registry := range registries {
		if err := fh.AddRegistry(registry); err != nil {
			return err
		}
	}
	return nil
}

// Registry retrieves the complete function registry that has been configured
// within this FunctionHandler. This registry is ready to be used with template engines
// that accept FuncMap, such as html/template or text/template.
//
// NOTE: This will replace the `FuncsMap()`, `TxtFuncMap()` and `HtmlFuncMap()` from sprig
func (fh *FunctionHandler) Registry() registry.FunctionMap {
	fh.registerAliases() // Ensure all aliases are processed before returning the registry
	return fh.funcsMap
}

// Logger returns the logger used by a FunctionHandler.
func (fh *FunctionHandler) Logger() *slog.Logger {
	return fh.logger
}
