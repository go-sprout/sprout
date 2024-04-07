package sprout

import (
	"log/slog"
	"text/template"
)

// FunctionHandler manages function execution with configurable error handling
// and logging.
type FunctionHandler struct {
	ErrHandling ErrHandling
	errChan     chan error
	Logger      *slog.Logger
	funcMap     template.FuncMap
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
		Logger:      slog.New(&slog.TextHandler{}),
		funcMap:     make(template.FuncMap),
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
		p.Logger = l
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
func WithFunctionHandler(new *FunctionHandler) FunctionHandlerOption {
	return func(fnh *FunctionHandler) {
		*fnh = *new
	}
}

// FuncMap returns a template.FuncMap for use with text/template or html/template.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARD COMPATIBILITY ONLY
func FuncMap(opts ...FunctionHandlerOption) template.FuncMap {
	fnHandler := NewFunctionHandler(opts...)

	// BACKWARD COMPATIBILITY
	// Fallback to FuncMap() to get the unmigrated functions
	for k, v := range TxtFuncMap() {
		fnHandler.funcMap[k] = v
	}

	// Added migrated functions
	fnHandler.funcMap["hello"] = fnHandler.Hello

	// Register aliases for functions
	fnHandler.registerAliases()
	return fnHandler.funcMap
}
