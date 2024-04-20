package sprout

import (
	"log/slog"
	"reflect"
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
		errHandler: errors.NewErrHandler(),
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

// ErrHandler returns the error handler used by a FunctionHandler. This is useful
// for handling errors and logging based on the handler's configuration.
func (fnHandler *FunctionHandler) ErrHandler() errors.ErrorHandler {
	return fnHandler.errHandler
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

// DefaultValueFor returns the zero value for any type T.
// This utility is particularly useful in situations where functions need to return
// a default value when an error is encountered but execution must continue.
// It supports basic types, slices, maps, and pointers.
func DefaultValueFor[T interface{}](v T) T {
	typeOf := reflect.TypeOf(v)
	if typeOf == nil {
		return v
	}

	value := reflect.Zero(typeOf).Interface().(T)

	switch typeOf.Kind() {
	case reflect.Slice:
		value = reflect.MakeSlice(typeOf, 0, 0).Interface().(T)
	case reflect.Map:
		value = reflect.MakeMap(typeOf).Interface().(T)
	}

	return value
}
