package sprout

import (
	htemplate "html/template"
	ttemplate "text/template"

	"github.com/go-sprout/sprout/registry/backward"
	"github.com/go-sprout/sprout/registry/builtin"
	"github.com/go-sprout/sprout/registry/checksum"
	"github.com/go-sprout/sprout/registry/conversion"
	"github.com/go-sprout/sprout/registry/crypto"
	"github.com/go-sprout/sprout/registry/encoding"
	"github.com/go-sprout/sprout/registry/env"
	"github.com/go-sprout/sprout/registry/filesystem"
	"github.com/go-sprout/sprout/registry/maps"
	"github.com/go-sprout/sprout/registry/numeric"
	"github.com/go-sprout/sprout/registry/random"
	"github.com/go-sprout/sprout/registry/reflect"
	"github.com/go-sprout/sprout/registry/regexp"
	"github.com/go-sprout/sprout/registry/semver"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/strings"
	"github.com/go-sprout/sprout/registry/time"
	"github.com/go-sprout/sprout/registry/uniqueid"
)

// These functions are not guaranteed to evaluate to the same result for given input, because they
// refer to the environment or global state.
// FOR BACKWARDS COMPATIBILITY ONLY
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

// HermeticTxtFuncMap returns a 'text/template'.FuncMap with only repeatable functions.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticTxtFuncMap(opts ...FunctionHandlerOption) ttemplate.FuncMap {
	r := TxtFuncMap(opts...)
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// HermeticHtmlFuncMap returns an 'html/template'.Funcmap with only repeatable functions.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticHtmlFuncMap(opts ...FunctionHandlerOption) htemplate.FuncMap {
	r := HtmlFuncMap(opts...)
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// TxtFuncMap returns a 'text/template'.FuncMap
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func TxtFuncMap(opts ...FunctionHandlerOption) ttemplate.FuncMap {
	return ttemplate.FuncMap(FuncMap(opts...))
}

// HtmlFuncMap returns an 'html/template'.Funcmap
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HtmlFuncMap(opts ...FunctionHandlerOption) htemplate.FuncMap {
	return htemplate.FuncMap(FuncMap(opts...))
}

// GenericFuncMap returns a copy of the basic function map as a map[string]interface{}.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func GenericFuncMap(opts ...FunctionHandlerOption) map[string]interface{} {
	return FuncMap(opts...)
}

// FuncMap returns a template.FuncMap for use with text/template or html/template.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARD COMPATIBILITY ONLY
func FuncMap(opts ...FunctionHandlerOption) ttemplate.FuncMap {
	fnHandler := NewFunctionHandler(opts...)

	_ = fnHandler.AddRegistries(
		builtin.NewRegistry(),
		uniqueid.NewRegistry(),
		semver.NewRegistry(),
		backward.NewRegistry(),
		reflect.NewRegistry(),
		time.NewRegistry(),
		strings.NewRegistry(),
		random.NewRegistry(),
		checksum.NewRegistry(),
		conversion.NewRegistry(),
		numeric.NewRegistry(),
		encoding.NewRegistry(),
		regexp.NewRegistry(),
		slices.NewRegistry(),
		maps.NewRegistry(),
		crypto.NewRegistry(),
		filesystem.NewRegistry(),
		env.NewRegistry(),
	)

	// Register aliases for functions
	fnHandler.registerAliases()
	return fnHandler.funcsMap
}
