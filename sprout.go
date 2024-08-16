package sprout

import (
	"log/slog"
	"os"
)

// HandlerOption[Handler] defines a type for functional options that configure
// a typed Handler.
type HandlerOption[T Handler] func(T)

// New creates and returns a new instance of DefaultHandler with optional
// configurations.
//
// The DefaultHandler is initialized with default values, including an error
// handling strategy, a logger, and empty function maps. You can customize the
// DefaultHandler instance by passing in one or more HandlerOption functions,
// which apply specific configurations to the handler.
//
// Example usage:
//
//	logger := slog.New(slog.NewTextHandler(os.Stdout))
//	handler := New(
//	    WithLogger(logger),
//	    WithRegistry(myRegistry),
//	)
//
// In the above example, the DefaultHandler is created with a custom logger and
// a specific registry.
func New(opts ...HandlerOption[*DefaultHandler]) *DefaultHandler {
	dh := &DefaultHandler{
		ErrHandling: ErrHandlingReturnDefaultValue,
		errChan:     make(chan error),
		logger:      slog.New(slog.NewTextHandler(os.Stdout, nil)),
		registries:  make([]Registry, 0),

		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
	}

	for _, opt := range opts {
		opt(dh)
	}

	return dh
}

// DEPRECATED: NewFunctionHandler creates a new function handler with the
// default values. It is deprecated and should not be used. Use `New` instead.
func NewFunctionHandler(opts ...HandlerOption[*DefaultHandler]) *DefaultHandler {
	slog.Warn("NewFunctionHandler are deprecated. Use `New` instead")
	return New(opts...)
}
