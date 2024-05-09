package sprout

import (
	htemplate "html/template"
	ttemplate "text/template"
)

// HermeticTxtFuncMap returns a 'text/template'.FuncMap with only repeatable functions.
func HermeticTxtFuncMap(opts ...FunctionHandlerOption) ttemplate.FuncMap {
	r := TxtFuncMap(opts...)
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// HermeticHtmlFuncMap returns an 'html/template'.Funcmap with only repeatable functions.
func HermeticHtmlFuncMap(opts ...FunctionHandlerOption) htemplate.FuncMap {
	r := HtmlFuncMap(opts...)
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// TxtFuncMap returns a 'text/template'.FuncMap
func TxtFuncMap(opts ...FunctionHandlerOption) ttemplate.FuncMap {
	return ttemplate.FuncMap(FuncMap(opts...))
}

// HtmlFuncMap returns an 'html/template'.Funcmap
func HtmlFuncMap(opts ...FunctionHandlerOption) htemplate.FuncMap {
	return htemplate.FuncMap(FuncMap(opts...))
}

// GenericFuncMap returns a copy of the basic function map as a map[string]interface{}.
func GenericFuncMap(opts ...FunctionHandlerOption) map[string]interface{} {
	return FuncMap(opts...)
}

// These functions are not guaranteed to evaluate to the same result for given input, because they
// refer to the environment or global state.
var nonhermeticFunctions = []string{
	// Date functions
	"date",
	"dateInZone",
	"dateModify",
	"now",
	"htmlDate",
	"htmlDateInZone",

	// Strings
	"randAlphaNum",
	"randAlpha",
	"randAscii",
	"randNumeric",
	"randBytes",
	"uuidv4",

	// OS
	"env",
	"expandenv",

	// Network
	"getHostByName",
}
