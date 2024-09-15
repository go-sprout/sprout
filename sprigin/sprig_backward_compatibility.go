package sprigin

import (
	htemplate "html/template"
	"log/slog"
	gostrings "strings"
	ttemplate "text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/internal/runtime"
	"github.com/go-sprout/sprout/registry/backward"
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
	"github.com/go-sprout/sprout/registry/std"
	"github.com/go-sprout/sprout/registry/strings"
	"github.com/go-sprout/sprout/registry/time"
	"github.com/go-sprout/sprout/registry/uniqueid"
)

// BACKWARDS COMPATIBILITY
// The following functions are provided for backwards compatibility with the
// original sprig methods. They are not recommended for use in new code.
var bc_registerSprigFuncs = sprout.FunctionAliasMap{
	"dateModify":     []string{"date_modify"},                   // ! Deprecated: Should use dateModify instead
	"dateInZone":     []string{"date_in_zone"},                  // ! Deprecated: Should use dateInZone instead
	"mustDateModify": []string{"must_date_modify"},              // ! Deprecated: Should use mustDateModify instead
	"ellipsis":       []string{"abbrev"},                        // ! Deprecated: Should use ellipsis instead
	"ellipsisBoth":   []string{"abbrevboth"},                    // ! Deprecated: Should use ellipsisBoth instead
	"trimAll":        []string{"trimall"},                       // ! Deprecated: Should use trimAll instead
	"append":         []string{"push"},                          // ! Deprecated: Should use append instead
	"mustAppend":     []string{"mustPush"},                      // ! Deprecated: Should use mustAppend instead
	"list":           []string{"tuple"},                         // FIXME: with the addition of append/prepend these are no longer immutable.
	"max":            []string{"biggest"},                       // ! Deprecated: Should use max instead
	"toUpper":        []string{"upper", "toupper", "uppercase"}, // ! Deprecated: Should use toUpper instead
	"toLower":        []string{"lower", "tolower", "lowercase"}, // ! Deprecated: Should use toLower instead
	"add":            []string{"addf"},                          // ! Deprecated: Should use add instead
	"add1":           []string{"add1f"},                         // ! Deprecated: Should use add1 instead
	"sub":            []string{"subf"},                          // ! Deprecated: Should use sub instead
	"toTitleCase":    []string{"title", "titlecase"},            // ! Deprecated: Should use toTitleCase instead
	"toPascalCase":   []string{"camelcase"},                     // ! Deprecated: Should use toPascalCase instead
	"toSnakeCase":    []string{"snake", "snakecase"},            // ! Deprecated: Should use toSnakeCase instead
	"toKebabCase":    []string{"kebab", "kebabcase"},            // ! Deprecated: Should use toKebabCase instead
	"swapCase":       []string{"swapcase"},                      // ! Deprecated: Should use swapCase instead
	"base64Encode":   []string{"b64enc"},                        // ! Deprecated: Should use base64Encode instead
	"base64Decode":   []string{"b64dec"},                        // ! Deprecated: Should use base64Decode instead
	"base32Encode":   []string{"b32enc"},                        // ! Deprecated: Should use base32Encode instead
	"base32Decode":   []string{"b32dec"},                        // ! Deprecated: Should use base32Decode instead
	"pathBase":       []string{"base"},                          // ! Deprecated: Should use pathBase instead
	"pathDir":        []string{"dir"},                           // ! Deprecated: Should use pathDir instead
	"pathExt":        []string{"ext"},                           // ! Deprecated: Should use pathExt instead
	"pathClean":      []string{"clean"},                         // ! Deprecated: Should use pathClean instead
	"pathIsAbs":      []string{"isAbs"},                         // ! Deprecated: Should use pathIsAbs instead
	"expandEnv":      []string{"expandenv"},                     // ! Deprecated: Should use expandEnv instead
	"dateAgo":        []string{"ago"},                           // ! Deprecated: Should use dateAgo instead
	"strSlice":       []string{"toStrings"},                     // ! Deprecated: Should use strSlice instead
	"toInt":          []string{"int", "atoi"},                   // ! Deprecated: Should use toInt instead
	"toInt64":        []string{"int64"},                         // ! Deprecated: Should use toInt64 instead
	"toFloat64":      []string{"float64"},                       // ! Deprecated: Should use toFloat64 instead
	"toOctal":        []string{"toDecimal"},                     // ! Deprecated: Should use toOctal instead
}

// \ BACKWARDS COMPATIBILITY

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

type SprigHandler struct {
	registries []sprout.Registry
	notices    []sprout.FunctionNotice

	funcsMap   sprout.FunctionMap
	funcsAlias sprout.FunctionAliasMap
}

func NewSprigHandler() *SprigHandler {
	return &SprigHandler{
		registries: make([]sprout.Registry, 0),
		funcsMap:   make(sprout.FunctionMap),
		funcsAlias: make(sprout.FunctionAliasMap),
	}
}

func (sh *SprigHandler) AddRegistry(registry sprout.Registry) error {
	sh.registries = append(sh.registries, registry)

	_ = registry.LinkHandler(sh)
	_ = registry.RegisterFunctions(sh.funcsMap)

	if regAlias, ok := registry.(sprout.RegistryWithAlias); ok {
		_ = regAlias.RegisterAliases(sh.funcsAlias)
	}

	if regNotice, ok := registry.(sprout.RegistryWithNotice); ok {
		_ = regNotice.RegisterNotices(&sh.notices)
	}

	return nil
}

func (sh *SprigHandler) AddRegistries(registries ...sprout.Registry) error {
	for _, registry := range registries {
		_ = sh.AddRegistry(registry)
	}
	return nil
}

func (sh *SprigHandler) Logger() *slog.Logger {
	return slog.New(slog.Default().Handler())
}

func (sh *SprigHandler) RawFunctions() sprout.FunctionMap {
	return sh.funcsMap
}

func (sh *SprigHandler) RawAliases() sprout.FunctionAliasMap {
	return sh.funcsAlias
}

func (sh *SprigHandler) Notices() []sprout.FunctionNotice {
	return sh.notices
}

func (sh *SprigHandler) Build() sprout.FunctionMap {
	_ = sh.AddRegistries(
		std.NewRegistry(),
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
	// BACKWARDS COMPATIBILITY
	// Register the sprig function aliases
	for originalFunction, aliases := range bc_registerSprigFuncs {
		for _, alias := range aliases {
			if fn, ok := sh.funcsMap[originalFunction]; ok {
				sh.funcsMap[alias] = fn
				sh.notices = append(sh.notices, *sprout.NewDeprecatedNotice(alias, "please use `"+originalFunction+"` instead"))
			}
		}
	}
	// \ BACKWARDS COMPATIBILITY

	sprout.AssignAliases(sh)
	sprout.AssignNotices(sh)

	// BACKWARDS COMPATIBILITY
	// Ensure error handling is consistent with sprig functions
	for funcName, fn := range sh.funcsMap {
		if !gostrings.HasPrefix(funcName, "must") {
			sh.funcsMap[funcName] = func(args ...any) (any, error) {
				out, _ := runtime.SafeCall(fn, args...)
				return out, nil
			}
		}
	}
	// \ BACKWARDS COMPATIBILITY

	return sh.funcsMap
}

// HermeticTxtFuncMap returns a 'text/template'.FuncMap with only repeatable functions.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticTxtFuncMap() ttemplate.FuncMap {
	r := TxtFuncMap()
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// HermeticHtmlFuncMap returns an 'html/template'.Funcmap with only repeatable functions.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticHtmlFuncMap() htemplate.FuncMap {
	r := HtmlFuncMap()
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// TxtFuncMap returns a 'text/template'.FuncMap
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func TxtFuncMap() ttemplate.FuncMap {
	return ttemplate.FuncMap(FuncMap())
}

// HtmlFuncMap returns an 'html/template'.Funcmap
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HtmlFuncMap() htemplate.FuncMap {
	return htemplate.FuncMap(FuncMap())
}

// GenericFuncMap returns a copy of the basic function map as a map[string]any.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func GenericFuncMap() map[string]any {
	return FuncMap()
}

// FuncMap returns a template.FuncMap for use with text/template or html/template.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARD COMPATIBILITY ONLY
func FuncMap() ttemplate.FuncMap {
	sprigHandler := NewSprigHandler()

	return sprigHandler.Build()
}
