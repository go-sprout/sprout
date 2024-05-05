package sprout

// FunctionAliasMap is a map that stores a list of aliases for each function.
type FunctionAliasMap map[string][]string

// BACKWARDS COMPATIBILITY
// The following functions are provided for backwards compatibility with the
// original sprig methods. They are not recommended for use in new code.
var bc_registerSprigFuncs = FunctionAliasMap{
	"dateModify":     {"date_modify"},                   //! Deprecated: Should use dateModify instead
	"dateInZone":     {"date_in_zone"},                  //! Deprecated: Should use dateInZone instead
	"mustDateModify": {"must_date_modify"},              //! Deprecated: Should use mustDateModify instead
	"ellipsis":       {"abbrev"},                        //! Deprecated: Should use ellipsis instead
	"ellipsisBoth":   {"abbrevboth"},                    //! Deprecated: Should use ellipsisBoth instead
	"trimAll":        {"trimall"},                       //! Deprecated: Should use trimAll instead
	"int":            {"atoi"},                          //! Deprecated: Should use toInt instead
	"append":         {"push"},                          //! Deprecated: Should use append instead
	"mustAppend":     {"mustPush"},                      //! Deprecated: Should use mustAppend instead
	"list":           {"tuple"},                         // FIXME: with the addition of append/prepend these are no longer immutable.
	"max":            {"biggest"},                       //! Deprecated: Should use max instead
	"toUpper":        {"upper", "toupper", "uppercase"}, //! Deprecated: Should use toUpper instead
	"toLower":        {"lower", "tolower", "lowercase"}, //! Deprecated: Should use toLower instead
	"add":            {"addf"},                          //! Deprecated: Should use add instead
	"add1":           {"add1f"},                         //! Deprecated: Should use add1 instead
	"sub":            {"subf"},                          //! Deprecated: Should use sub instead
	"toTitleCase":    {"title", "titlecase"},            //! Deprecated: Should use toTitleCase instead
	"toCamelCase":    {"camel", "camelcase"},            //! Deprecated: Should use toCamelCase instead
	"toSnakeCase":    {"snake", "snakecase"},            //! Deprecated: Should use toSnakeCase instead
	"toKebabCase":    {"kebab", "kebabcase"},            //! Deprecated: Should use toKebabCase instead
	"swapCase":       {"swapcase"},                      //! Deprecated: Should use swapCase instead
	"base64Encode":   {"b64enc"},                        //! Deprecated: Should use base64Encode instead
	"base64Decode":   {"b64dec"},                        //! Deprecated: Should use base64Decode instead
	"base32Encode":   {"b32enc"},                        //! Deprecated: Should use base32Encode instead
	"base32Decode":   {"b32dec"},                        //! Deprecated: Should use base32Decode instead
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
	// BACKWARDS COMPATIBILITY
	// Register the sprig function aliases
	for originalFunction, aliases := range bc_registerSprigFuncs {
		for _, alias := range aliases {
			fh.funcMap[alias] = fh.funcMap[originalFunction]
		}
	}
	//\ BACKWARDS COMPATIBILITY

	for originalFunction, aliases := range fh.funcsAlias {
		for _, alias := range aliases {
			fh.funcMap[alias] = fh.funcMap[originalFunction]
		}
	}
}
