package sprout

import "log/slog"

// Handler is the interface that wraps the basic methods of a handler to manage
// all registries and functions.
// The Handler brick is the main brick of sprout. It is used to configure and
// manage a cross-registry configuration and function management like a global
// logging system, error handling, and more.
// ! This interface is not meant to be implemented by the user but by the
// ! library itself. An user could implement it but it is not recommended.
type Handler interface {
	Logger() *slog.Logger

	AddRegistry(registry Registry) error
	AddRegistries(registries ...Registry) error

	Functions() FunctionMap
	Aliases() FunctionAliasMap

	Build() FunctionMap
}

// DefaultHandler manages function execution with configurable error handling
// and logging.
type DefaultHandler struct {
	ErrHandling      ErrHandling
	errChan          chan error
	logger           *slog.Logger
	registries       []Registry
	cachedFuncsMap   FunctionMap
	cachedFuncsAlias FunctionAliasMap
}

// RegisterHandler registers a single FunctionRegistry implementation (e.g., a handler)
// into the FunctionHandler's internal function registry. This method allows for integrating
// additional functions into the template processing environment.
func (dh *DefaultHandler) AddRegistry(reg Registry) error {
	dh.registries = append(dh.registries, reg)

	if err := reg.LinkHandler(dh); err != nil {
		return err
	}

	if err := reg.RegisterFunctions(dh.cachedFuncsMap); err != nil {
		return err
	}

	if regAlias, ok := reg.(RegistryWithAlias); ok {
		if err := regAlias.RegisterAliases(dh.cachedFuncsAlias); err != nil {
			return err
		}

	}

	return nil
}

// RegisterHandlers registers multiple FunctionRegistry implementations into the
// FunctionHandler's internal function registry. This method simplifies the process
// of adding multiple sets of functionalities into the template engine at once.
func (dh *DefaultHandler) AddRegistries(registries ...Registry) error {
	for _, registry := range registries {
		if err := dh.AddRegistry(registry); err != nil {
			return err
		}
	}
	return nil
}

// Build retrieves the complete suite of functiosn and alias that has been configured
// within this Handler. This handler is ready to be used with template engines
// that accept FuncMap, such as html/template or text/template.
//
// NOTE: This will replace the `FuncsMap()`, `TxtFuncMap()` and `HtmlFuncMap()` from sprig
func (dh *DefaultHandler) Build() FunctionMap {
	AssignAliases(dh) // Ensure all aliases are processed before returning the registry
	return dh.cachedFuncsMap
}

// Logger returns the logger instance associated with the DefaultHandler.
//
// The logger is used for logging information, warnings, and errors that occur
// during the execution of functions managed by the DefaultHandler. By default,
// the logger is initialized with a basic text handler, but it can be customized
// using the WithLogger option when creating a new DefaultHandler.
func (dh *DefaultHandler) Logger() *slog.Logger {
	return dh.logger
}

// Functions returns the map of registered functions managed by the DefaultHandler.
//
// ⚠ This function is for special cases where you need to access the function
// map for the template engine use `Build()` instead.
//
// This function map contains all the functions that have been added to the handler,
// typically for use in templating engines. Each entry in the map associates a function
// name with its corresponding implementation.
func (dh *DefaultHandler) Functions() FunctionMap {
	return dh.cachedFuncsMap
}

// Aliases returns the map of function aliases managed by the DefaultHandler.
//
// The alias map allows certain functions to be referenced by multiple names. This
// can be useful in templating environments where different names might be preferred
// for the same underlying function. The alias map associates each original function
// name with a list of aliases that can be used interchangeably.
func (dh *DefaultHandler) Aliases() FunctionAliasMap {
	return dh.cachedFuncsAlias
}

// WithLogger sets the logger used by a DefaultHandler.
func WithLogger(l *slog.Logger) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) {
		p.logger = l
	}
}

// WithHandler updates a DefaultHandler with settings from another DefaultHandler.
// This is useful for copying configurations between handlers.
func WithHandler(new Handler) HandlerOption[*DefaultHandler] {
	return func(fnh *DefaultHandler) {
		if new == nil {
			return
		}

		if fhCast, ok := new.(*DefaultHandler); ok {
			*fnh = *fhCast
		}
	}
}
