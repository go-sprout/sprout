package sprout

// FunctionAliasMap is a map that stores a list of aliases for each function.
type FunctionAliasMap = map[string][]string

// WithAlias returns a FunctionHandlerOption that associates one or more alias
// names with an original function name.
// This allows the function to be called by any of its aliases.
//
// originalFunction specifies the original function name to which aliases will
// be added. aliases is a variadic parameter that takes one or more strings as
// aliases for the original function.
//
// The function does nothing if no aliases are provided.
// If the original function name does not already have associated aliases in
// the FunctionHandler, a new slice of strings is created to hold its aliases.
// Each provided alias is then appended to this slice.
//
// This option must be applied to a FunctionHandler using the FunctionHandler's
// options mechanism for the aliases to take effect.
//
// Example:
//
//	handler := NewFunctionHandler(WithAlias("originalFunc", "alias1", "alias2"))
func WithAlias(originalFunction string, aliases ...string) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		if len(aliases) == 0 {
			return
		}

		if _, ok := p.funcsAlias[originalFunction]; !ok {
			p.funcsAlias[originalFunction] = make([]string, 0)
		}

		p.funcsAlias[originalFunction] = append(p.funcsAlias[originalFunction], aliases...)
	}
}

// WithAliases returns a FunctionHandlerOption that configures multiple aliases
// for function names in a single call. It allows a batch of functions to be
// associated with their respective aliases, facilitating the creation of
// aliases for multiple functions at once.
//
// This option must be applied to a FunctionHandler using the FunctionHandler's
// options mechanism for the aliases to take effect.
// It complements the WithAlias function by providing a means to configure
// multiple aliases in one operation, rather than one at a time.
//
// Example:
//
//	handler := NewFunctionHandler(WithAliases(sprout.FunctionAliasMap{
//	    "originalFunc1": {"alias1_1", "alias1_2"},
//	    "originalFunc2": {"alias2_1", "alias2_2"},
//	}))
func WithAliases(aliases FunctionAliasMap) FunctionHandlerOption {
	return func(p *FunctionHandler) {
		for originalFunction, aliasList := range aliases {
			if _, ok := p.funcsAlias[originalFunction]; !ok {
				p.funcsAlias[originalFunction] = make([]string, 0)
			}

			p.funcsAlias[originalFunction] = append(p.funcsAlias[originalFunction], aliasList...)
		}
	}
}

// registerAliases allows the aliases to be used as references to the original
// functions.
//
// It should be called after all aliases have been added through the WithAlias
// option and before the function map is used to ensure all aliases are properly
// registered.
func (fh *FunctionHandler) registerAliases() {
	for originalFunction, aliases := range fh.funcsAlias {
		for _, alias := range aliases {
			if fn, ok := fh.funcsMap[originalFunction]; ok {
				fh.funcsMap[alias] = fn
			}
		}
	}
}
