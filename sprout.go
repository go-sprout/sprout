package sprout

import (
	"log/slog"
	"text/template"

	"github.com/42atomys/sprout/errors"
)

// FunctionHandler manages function execution with configurable error handling
// and logging.
type FunctionHandler struct {
	errHandler errors.ErrorHandler
	logger     *slog.Logger
	funcMap    template.FuncMap
	funcsAlias FunctionAliasMap
}

// FunctionHandlerOption defines a type for functional options that configure
// FunctionHandler.
type FunctionHandlerOption func(*FunctionHandler)

// NewFunctionHandler creates a new FunctionHandler with the provided options.
func NewFunctionHandler(opts ...FunctionHandlerOption) *FunctionHandler {
	fnHandler := &FunctionHandler{
		errHandler: errors.NewErrHandler(
			errors.WithSubHandler(NewErrorChainHandler()),
		),
		logger:     slog.New(slog.Default().Handler()),
		funcMap:    make(template.FuncMap),
		funcsAlias: make(FunctionAliasMap),
	}

	for _, opt := range opts {
		opt(fnHandler)
	}

	return fnHandler
}

// WithErrHandler sets the error handling strategy for a FunctionHandler.
func WithErrHandler(eh errors.ErrorHandler) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.errHandler = eh
	}
}

// WithLogger sets the logger used by a FunctionHandler and the error handler.
func WithLogger(l *slog.Logger) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		p.logger = l
		errors.WithLogger(l)(p.errHandler)
	}
}

// WithFunctionHandler updates a FunctionHandler with settings from another FunctionHandler.
// This is useful for copying configurations between handlers.
func WithFunctionHandler(new *FunctionHandler) FunctionHandlerOption {
	return func(fnh *FunctionHandler) {
		*fnh = *new
	}
}

// Logger returns the logger used by a FunctionHandler. This is useful for
// logging errors and other information based on the handler's configuration.
func (fnHandler *FunctionHandler) Logger() *slog.Logger {
	return fnHandler.logger
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
