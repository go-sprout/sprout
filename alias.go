package sprout

// BACKWARDS COMPATIBILITY
// The following functions are provided for backwards compatibility with the
// original sprig methods. They are not recommended for use in new code.
var bc_registerSprigFuncs = map[string][]string{
	"dateModify":     {"date_modify"},      //! Deprecated: Should use dateModify instead
	"dateInZone":     {"date_in_zone"},     //! Deprecated: Should use dateInZone instead
	"mustDateModify": {"must_date_modify"}, //! Deprecated: Should use mustDateModify instead
	"ellipsis":       {"abbrev"},           //! Deprecated: Should use ellipsis instead
	"ellipsisBoth":   {"abbrevboth"},       //! Deprecated: Should use ellipsisBoth instead
	"trimAll":        {"trimall"},          //! Deprecated: Should use trimAll instead
	"int":            {"atoi"},             //! Deprecated: Should use toInt instead
	"append":         {"push"},             //! Deprecated: Should use append instead
	"mustAppend":     {"mustPush"},         //! Deprecated: Should use mustAppend instead
	"list":           {"tuple"},            // FIXME: with the addition of append/prepend these are no longer immutable.
	"max":            {"biggest"},
}

//\ BACKWARDS COMPATIBILITY

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

// registerAliases allows the aliases to be used as references to the original
// functions.
//
// It should be called after all aliases have been added through the WithAlias
// option and before the function map is used to ensure all aliases are properly
// registered.
func (p *FunctionHandler) registerAliases() {
	// BACKWARDS COMPATIBILITY
	// Register the sprig function aliases
	for originalFunction, aliases := range bc_registerSprigFuncs {
		for _, alias := range aliases {
			p.funcMap[alias] = p.funcMap[originalFunction]
		}
	}
	//\ BACKWARDS COMPATIBILITY

	for originalFunction, aliases := range p.funcsAlias {
		for _, alias := range aliases {
			p.funcMap[alias] = p.funcMap[originalFunction]
		}
	}
}
