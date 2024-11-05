package sprout

import (
	"log/slog"
	"slices"
	gostrings "strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-sprout/sprout/internal/runtime"
)

// Handler is the interface that wraps the basic methods of a handler to manage
// all registries and functions.
// The Handler brick is the main brick of sprout. It is used to configure and
// manage a cross-registry configuration and function management like a global
// logging system, error handling, and more.
// ! This interface is not meant to be implemented by the user but by the
// ! library itself. An user could implement it but it is not recommended.
type Handler interface {
	Logger() *slog.Logger

	// AddRegistry registers a single registry,  into the Handler.
	// This method allows for integrating additional functions into the template
	// processing environment.
	AddRegistry(registry Registry) error

	// RawFunctions returns the map of registered functions without any alias,
	// notices or other additional information. This function is useful for
	// special cases where you need to access raw data from registries.
	//
	// ⚠ To access the function map for the template engine use `Build()` instead.
	RawFunctions() FunctionMap

	// RawAliases returns the map of function aliases managed by the Handler.
	RawAliases() FunctionAliasMap

	// Notices returns the list of function notices managed by the Handler.
	Notices() []FunctionNotice

	// Build retrieves the complete suite of functions and aliases that has been
	// configured within this Handler. This handler is ready to be used with
	// template engines that accept FuncMap, such as html/template or text/template.
	//
	// Build should call AssignAliases and AssignNotices to ensure that all aliases
	// and notices are properly associated with their original functions.
	Build() FunctionMap
}

// DefaultHandler manages function execution with configurable error handling
// and logging.
type DefaultHandler struct {
	logger     *slog.Logger
	registries []Registry
	notices    []FunctionNotice

	wantSafeFuncs bool
	built         bool

	cachedFuncsMap   FunctionMap
	cachedFuncsAlias FunctionAliasMap
}

// RegisterHandler registers a single FunctionRegistry implementation (e.g., a handler)
// into the FunctionHandler's internal function registry. This method allows for integrating
// additional functions into the template processing environment.
// This function prevents duplicate registry registration by checking the UID
// of the registry.
func (dh *DefaultHandler) AddRegistry(reg Registry) error {
	if slices.ContainsFunc(dh.registries, func(r Registry) bool {
		return r.UID() == reg.UID()
	}) {
		return nil
	}

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

	if regNotice, ok := reg.(RegistryWithNotice); ok {
		if err := regNotice.RegisterNotices(&dh.notices); err != nil {
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

// Build retrieves the complete suite of functions and alias that has been configured
// within this Handler. This handler is ready to be used with template engines
// that accept FuncMap, such as html/template or text/template. It will also
// cache the function map for future use to avoid rebuilding the function map
// multiple times, so it is safe to call this method multiple times to retrieve
// the same built function map.
//
// NOTE: This replaces the [github.com/Masterminds/sprig.FuncMap],
// [github.com/Masterminds/sprig.TxtFuncMap] and [github.com/Masterminds/sprig.HtmlFuncMap]
// from sprig
func (dh *DefaultHandler) Build() FunctionMap {
	if dh.built {
		return dh.cachedFuncsMap
	}

	AssignAliases(dh) // Ensure all aliases are processed before returning the registry
	AssignNotices(dh) // Ensure all notices are processed before returning the registry
	if dh.wantSafeFuncs {
		AssignSafeFuncs(dh) // Ensure all functions are wrapped with safe functions
	}

	dh.built = true
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

// RawFunctions returns the map of registered functions managed by the DefaultHandler.
//
// ⚠ This function is for special cases where you need to access the function
// map for the template engine use `Build()` instead.
//
// This function map contains all the functions that have been added to the handler,
// typically for use in templating engines. Each entry in the map associates a function
// name with its corresponding implementation.
func (dh *DefaultHandler) RawFunctions() FunctionMap {
	return dh.cachedFuncsMap
}

// RawAliases returns the map of function aliases managed by the DefaultHandler.
//
// The alias map allows certain functions to be referenced by multiple names. This
// can be useful in templating environments where different names might be preferred
// for the same underlying function. The alias map associates each original function
// name with a list of aliases that can be used interchangeably.
func (dh *DefaultHandler) RawAliases() FunctionAliasMap {
	return dh.cachedFuncsAlias
}

// Notices returns the list of function notices managed by the DefaultHandler.
//
// The notices list contains information about functions that have been deprecated
// or are otherwise subject to special handling. Each notice includes the name of
// the function, a message describing the notice, and the kind of notice (e.g., info
// or deprecated).
func (dh *DefaultHandler) Notices() []FunctionNotice {
	return dh.notices
}

// WithLogger sets the logger used by a DefaultHandler.
func WithLogger(l *slog.Logger) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) error {
		p.logger = l
		return nil
	}
}

// WithHandler updates a DefaultHandler with settings from another DefaultHandler.
// This is useful for copying configurations between handlers.
func WithHandler(new Handler) HandlerOption[*DefaultHandler] {
	return func(fnh *DefaultHandler) error {
		if new == nil {
			return nil
		}

		if fhCast, ok := new.(*DefaultHandler); ok {
			*fnh = *fhCast
		}

		return nil
	}
}

// WithSafeFuncs enables safe function calls in a DefaultHandler. When safe functions
// are enabled, the handler will wrap all functions with a safe wrapper that logs any
// errors that occur during function execution without interrupting the execution of
// the template.
//
// To use a safe function, prepend `safe` to the original function name,
// example: `safeOriginalFuncName` instead of `originalFuncName`.
func WithSafeFuncs(enabled bool) HandlerOption[*DefaultHandler] {
	return func(dh *DefaultHandler) error {
		dh.wantSafeFuncs = enabled
		return nil
	}
}

// AssignSafeFuncs wraps all functions with a safe wrapper that logs any errors
// that occur during function execution. If safe functions are enabled in the
// DefaultHandler, this method will prepend "safe" to the function name and
// create a safe wrapper for each function.
//
// E.G. all functions will have both the original function name and a safe function name:
//
//	originalFuncName -> SafeOriginalFuncName
func AssignSafeFuncs(handler Handler) {
	safeFuncs := make(FunctionMap)
	for funcName, fn := range handler.RawFunctions() {
		safeFuncs[safeFuncName(funcName)] = safeWrapper(handler, funcName, fn)
	}

	for funcName, fn := range safeFuncs {
		handler.RawFunctions()[funcName] = fn
	}
}

// safeWrapper create a safe wrapper function that calls the original function
// and logs any errors that occur during the function call without interrupting
// the execution of the template.
func safeWrapper(handler Handler, functionName string, fn any) wrappedFunction {
	return func(args ...any) (any, error) {
		out, err := runtime.SafeCall(fn, args...)
		if err != nil {
			handler.Logger().With("function", functionName, "error", err).Error("function call failed")
		}
		return out, nil
	}
}

// safeFuncName generates a safe function name by prepending "safe" to the original
// function name and capitalizing the first letter of the function name.
//
// Example:
//
//	originalFuncName -> SafeOriginalFuncName
func safeFuncName(name string) string {
	if name == "" {
		return ""
	}

	var b gostrings.Builder
	b.Grow(len(name) + 4)

	b.WriteString("safe")
	b.WriteString(cases.Title(language.Und, cases.NoLower).String(name))

	return b.String()
}
