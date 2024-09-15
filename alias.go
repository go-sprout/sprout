package sprout

// FunctionAliasMap is a map that stores a list of aliases for each function.
type FunctionAliasMap = map[string][]string

// AssignAliases assigns all aliases defined in the handler to their original
// functions. This function is used to ensure that all aliases are properly
// associated with their original functions in the handler instance.
//
// It should be called after all functions and aliases have been added and
// inside the Build function in case of using a custom handler.
func AssignAliases(h Handler) {
	for originalName, aliases := range h.RawAliases() {
		_, exists := h.RawFunctions()[originalName]
		if !exists {
			continue
		}

		for _, alias := range aliases {
			if fn, ok := h.RawFunctions()[originalName]; ok {
				h.RawFunctions()[alias] = fn
			}
		}
	}
}

// WithAlias returns a FunctionOption[**DefaultHandler] that associates one or more alias
// names with an original function name.
// This allows the function to be called by any of its aliases.
//
// originalFunction specifies the original function name to which aliases will
// be added. aliases is a variadic parameter that takes one or more strings as
// aliases for the original function.
//
// The function does nothing if no aliases are provided.
// If the original function name does not already have associated aliases in
// the DefaultHandler, a new slice of strings is created to hold its aliases.
// Each provided alias is then appended to this slice.
//
// This option must be applied to a DefaultHandler using the DefaultHandler's
// options mechanism for the aliases to take effect.
//
// Example:
//
//	handler := New(WithAlias("originalFunc", "alias1", "alias2"))
func WithAlias(originalFunction string, aliases ...string) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) {
		if len(aliases) == 0 {
			return
		}

		if _, ok := p.cachedFuncsAlias[originalFunction]; !ok {
			p.cachedFuncsAlias[originalFunction] = make([]string, 0)
		}

		p.cachedFuncsAlias[originalFunction] = append(p.cachedFuncsAlias[originalFunction], aliases...)
	}
}

// WithAliases returns a FunctionOption[**DefaultHandler] that configures multiple aliases
// for function names in a single call. It allows a batch of functions to be
// associated with their respective aliases, facilitating the creation of
// aliases for multiple functions at once.
//
// This option must be applied to a DefaultHandler using the DefaultHandler's
// options mechanism for the aliases to take effect.
// It complements the WithAlias function by providing a means to configure
// multiple aliases in one operation, rather than one at a time.
//
// Example:
//
//	handler := New(WithAliases(sprout.FunctionAliasMap{
//	    "originalFunc1": {"alias1_1", "alias1_2"},
//	    "originalFunc2": {"alias2_1", "alias2_2"},
//	}))
func WithAliases(aliases FunctionAliasMap) HandlerOption[*DefaultHandler] {
	return func(p *DefaultHandler) {
		for originalFunction, aliasList := range aliases {
			if _, ok := p.cachedFuncsAlias[originalFunction]; !ok {
				p.cachedFuncsAlias[originalFunction] = make([]string, 0)
			}

			p.cachedFuncsAlias[originalFunction] = append(p.cachedFuncsAlias[originalFunction], aliasList...)
		}
	}
}
