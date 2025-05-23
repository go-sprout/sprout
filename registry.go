package sprout

import "text/template"

// FunctionMap is an alias for template.FuncMap, which maps function names
// to functions. This registry is used to register all template functions.
type FunctionMap = template.FuncMap

// Registry is an interface that defines the method to register functions
// within a given Handler.
// The Registry brick are a way to group functions together and register them
// with a Handler. This is useful to split functions into different categories
// and to avoid having a single large file with all functions and optimize the
// performance of the template engine.
// It also allows for easy extension of the template functions by adding a new one.
type Registry interface {
	// UID returns the unique name of the registry. This name is used to identify
	// the registry author and name and prevent duplicate registry registration.
	UID() string
	// LinkHandler links the given Handler to the registry.
	// * This method help you to have access to the main handler and its
	// * functionalities, like the logger, error handling, and more.
	LinkHandler(fh Handler) error
	// RegisterFunctions adds the provided functions into the given function map.
	// This method is called by an Handler to register all functions of a registry.
	RegisterFunctions(fnMap FunctionMap) error
}

type RegistryWithAlias interface {
	// RegisterAliases adds the provided aliases into the given alias map.
	// This method is called by an Handler to register all aliases of a registry.
	RegisterAliases(aliasMap FunctionAliasMap) error
}

type RegistryWithNotice interface {
	// RegisterNotices adds the provided notices into the given notice list.
	// This method is called by an Handler to register all notices of a registry.
	RegisterNotices(notices *[]FunctionNotice) error
}

// AddFunction adds a new function under the specified name to the given registry.
// If the function name already exists in the registry, this method does nothing to
// prevent accidental overwriting of existing registered functions.
func AddFunction(funcsMap FunctionMap, name string, function any) {
	if _, ok := funcsMap[name]; ok {
		return // Prevent overwriting existing functions
	}
	funcsMap[name] = function
}

// AddAlias adds an alias for the original function name. The original function
// name must already exist in the FunctionHandler's function map. If the original
// function name does not exist, the alias is not added.
func AddAlias(aliasMap FunctionAliasMap, originalFunction string, aliases ...string) {
	if len(aliases) == 0 {
		return
	}

	if _, ok := aliasMap[originalFunction]; !ok {
		aliasMap[originalFunction] = make([]string, 0)
	}

	aliasMap[originalFunction] = append(aliasMap[originalFunction], aliases...)
}

// AddNotice adds a new function notice to the given function
func AddNotice(notices *[]FunctionNotice, notice *FunctionNotice) {
	*notices = append(*notices, *notice)
}

// WithRegistries returns a HandlerOption that adds the provided registries to the handler.
// This option simplifies the process of adding multiple sets of functionalities into the
// template engine at once.
//
// Example:
//
//	handler := New(WithRegistries(myRegistry1, myRegistry2, myRegistry3))
func WithRegistries(registries ...Registry) HandlerOption[*DefaultHandler] {
	return func(dh *DefaultHandler) error {
		return dh.AddRegistries(registries...)
	}
}
