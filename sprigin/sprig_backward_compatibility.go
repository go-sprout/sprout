package sprigin

import (
	"fmt"
	htemplate "html/template"
	"log/slog"
	"strings"
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
	//nolint:staticcheck // sprigin mirrors the sprig API, it must keep the sprig signatures of the `regexp` registry
	"github.com/go-sprout/sprout/registry/regexp"
	"github.com/go-sprout/sprout/registry/semver"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/std"
	rstrings "github.com/go-sprout/sprout/registry/strings"
	rtime "github.com/go-sprout/sprout/registry/time"
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
	logger     *slog.Logger
	registries []sprout.Registry
	notices    []sprout.FunctionNotice

	funcsMap   sprout.FunctionMap
	funcsAlias sprout.FunctionAliasMap
}

func NewSprigHandler() *SprigHandler {
	return &SprigHandler{
		logger:     slog.New(slog.Default().Handler()),
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
	return sh.logger
}

// SignatureWarn logs a warning message about a deprecated function signature.
// This is used to warn users when they use the old Sprig signature instead of
// the new Sprout signature.
//
// oldSign and newSign are shown side by side, so they must differ only by what
// the user has to change. For the functions returning a new value that is
// usually accumulated, both should carry the assignment, otherwise the message
// reads as if the call alone was enough.
func (sh *SprigHandler) SignatureWarn(functionName, oldSign, newSign string) {
	msg := fmt.Sprintf("The signature of `%s` has changed from `%s` to `%s`, please update your code before next upgrade. This change will simplify the usage of the function and respect go/template conventions and allow usage of pipe (`|`).", functionName, oldSign, newSign)

	sh.Logger().With("function", functionName, "notice", "deprecated").Warn(fmt.Sprintf("Template function `%s` is deprecated: %s", functionName, msg))
}

// AmbiguousSignatureWarn logs a warning message when the arguments of a call
// match both the Sprig and the Sprout signatures, which happens when the
// receiver and the value have the same type, e.g. appending a list to a list.
//
// The call cannot be disambiguated from the arguments alone, so the Sprig
// signature is assumed for backward compatibility. Callers already migrated to
// the Sprout signature get their arguments swapped, hence a message of its own
// instead of the migration one, which would tell them to move to the very
// signature they already use.
func (sh *SprigHandler) AmbiguousSignatureWarn(functionName, assumedSign, migratedSign string) {
	msg := fmt.Sprintf("The arguments of `%s` match both signatures, the sprig one `%s` has been assumed. If your template already uses `%s`, its arguments have been swapped: use the `sprout` package directly, where only the sprout signature exists, to remove the ambiguity.", functionName, assumedSign, migratedSign)

	sh.Logger().With("function", functionName, "notice", "ambiguous").Warn(fmt.Sprintf("Template function `%s` is ambiguous: %s", functionName, msg))
}

func (sh *SprigHandler) BreakingWarn(functionName, changeNotice string) {
	sh.Logger().With("function", functionName, "notice", "deprecated").Warn(fmt.Sprintf("Template function `%s` have a significant change in next version : %s", functionName, changeNotice))
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
		rtime.NewRegistry(),
		rstrings.NewRegistry(),
		random.NewRegistry(),
		checksum.NewRegistry(),
		conversion.NewRegistry(),
		numeric.NewRegistry(),
		encoding.NewRegistry(),
		regexp.NewRegistry(), //nolint:staticcheck // sprigin must keep the sprig signatures, it cannot use the `regex` registry
		slices.NewRegistry(),
		maps.NewRegistry(),
		crypto.NewRegistry(),
		filesystem.NewRegistry(),
		env.NewRegistry(),
	)

	// BACKWARDS COMPATIBILITY
	// Override functions to use Sprig's signatures instead of Sprout's new signatures.
	// This is necessary because Sprig and Sprout use different argument orders.
	// The sprig* functions are defined in sprig_functions.go and handle both
	// old Sprig and new Sprout signatures with type-based detection.
	// Deprecation warnings are logged only when old Sprig signature is detected.
	mapsRegistry := maps.NewRegistry()
	_ = mapsRegistry.LinkHandler(sh)
	slicesRegistry := slices.NewRegistry()
	_ = slicesRegistry.LinkHandler(sh)

	// Maps registry overrides
	sh.funcsMap["dig"] = sh.sprigDig()
	sh.funcsMap["get"] = sh.sprigGet(mapsRegistry)
	sh.funcsMap["set"] = sh.sprigSet(mapsRegistry)
	sh.funcsMap["unset"] = sh.sprigUnset(mapsRegistry)
	sh.funcsMap["hasKey"] = sh.sprigHasKey(mapsRegistry)
	sh.funcsMap["pick"] = sh.sprigPick(mapsRegistry)
	sh.funcsMap["omit"] = sh.sprigOmit(mapsRegistry)

	// Slices registry overrides
	sh.funcsMap["append"] = sh.sprigAppend(slicesRegistry)
	sh.funcsMap["prepend"] = sh.sprigPrepend(slicesRegistry)
	sh.funcsMap["slice"] = sh.sprigSlice(slicesRegistry)
	sh.funcsMap["without"] = sh.sprigWithout(slicesRegistry)

	// Encoding registry overrides - return error message like sprig does
	encodingRegistry := encoding.NewRegistry()
	_ = encodingRegistry.LinkHandler(sh)
	sh.funcsMap["base64Decode"] = sh.sprigBase64Decode(encodingRegistry)
	sh.funcsMap["base32Decode"] = sh.sprigBase32Decode(encodingRegistry)

	// Time registry overrides - use local timezone like sprig does
	timeRegistry := rtime.NewRegistry()
	_ = timeRegistry.LinkHandler(sh)
	sh.funcsMap["date"] = sh.sprigDate(timeRegistry)

	// String registry overrides - re-introduce bugs
	stringRegistry := rstrings.NewRegistry()
	_ = stringRegistry.LinkHandler(sh)
	sh.funcsMap["substr"] = sh.sprigSubstr()
	sh.funcsMap["ellipsisBoth"] = sh.sprigAbbrevboth()
	sh.funcsMap["toTitleCase"] = sh.sprigTitle()

	// Reflect registry overrides - re-introduce bugs
	reflectRegistry := reflect.NewRegistry()
	_ = reflectRegistry.LinkHandler(sh)
	sh.funcsMap["kindOf"] = sh.sprigKindOf(reflectRegistry)

	// \ BACKWARDS COMPATIBILITY

	sprout.AssignAliases(sh)
	sprout.AssignNotices(sh)

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

	// BACKWARDS COMPATIBILITY
	// Ensure error handling is consistent with sprig functions
	for funcName, fn := range sh.funcsMap {
		if !strings.HasPrefix(funcName, "must") {
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
	return HermeticTxtFuncMapWith()
}

// HermeticHtmlFuncMap returns an 'html/template'.Funcmap with only repeatable functions.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticHtmlFuncMap() htemplate.FuncMap {
	return HermeticHtmlFuncMapWith()
}

// TxtFuncMap returns a 'text/template'.FuncMap
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func TxtFuncMap() ttemplate.FuncMap {
	return TxtFuncMapWith()
}

// HtmlFuncMap returns an 'html/template'.Funcmap
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func HtmlFuncMap() htemplate.FuncMap {
	return HtmlFuncMapWith()
}

// GenericFuncMap returns a copy of the basic function map as a map[string]any.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARDS COMPATIBILITY ONLY
func GenericFuncMap() map[string]any {
	return GenericFuncMapWith()
}

// FuncMap returns a template.FuncMap for use with text/template or html/template.
// It provides backward compatibility with sprig.FuncMap and integrates
// additional configured functions.
// FOR BACKWARD COMPATIBILITY ONLY
func FuncMap() ttemplate.FuncMap {
	return FuncMapWith()
}

// HermeticTxtFuncMapWith behaves like [HermeticTxtFuncMap] and applies the
// given handler options, such as [WithLogger].
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticTxtFuncMapWith(opts ...sprout.HandlerOption[*SprigHandler]) ttemplate.FuncMap {
	r := TxtFuncMapWith(opts...)
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// HermeticHtmlFuncMapWith behaves like [HermeticHtmlFuncMap] and applies the
// given handler options, such as [WithLogger].
// FOR BACKWARDS COMPATIBILITY ONLY
func HermeticHtmlFuncMapWith(opts ...sprout.HandlerOption[*SprigHandler]) htemplate.FuncMap {
	r := HtmlFuncMapWith(opts...)
	for _, name := range nonhermeticFunctions {
		delete(r, name)
	}
	return r
}

// TxtFuncMapWith behaves like [TxtFuncMap] and applies the given handler
// options, such as [WithLogger].
// FOR BACKWARDS COMPATIBILITY ONLY
func TxtFuncMapWith(opts ...sprout.HandlerOption[*SprigHandler]) ttemplate.FuncMap {
	return ttemplate.FuncMap(FuncMapWith(opts...))
}

// HtmlFuncMapWith behaves like [HtmlFuncMap] and applies the given handler
// options, such as [WithLogger].
// FOR BACKWARDS COMPATIBILITY ONLY
func HtmlFuncMapWith(opts ...sprout.HandlerOption[*SprigHandler]) htemplate.FuncMap {
	return htemplate.FuncMap(FuncMapWith(opts...))
}

// GenericFuncMapWith behaves like [GenericFuncMap] and applies the given
// handler options, such as [WithLogger].
// FOR BACKWARDS COMPATIBILITY ONLY
func GenericFuncMapWith(opts ...sprout.HandlerOption[*SprigHandler]) map[string]any {
	return FuncMapWith(opts...)
}

// FuncMapWith behaves like [FuncMap] and applies the given handler options,
// such as [WithLogger]. An option returning an error is reported through the
// handler logger and skipped, the function map is still built.
// FOR BACKWARD COMPATIBILITY ONLY
func FuncMapWith(opts ...sprout.HandlerOption[*SprigHandler]) ttemplate.FuncMap {
	sprigHandler := NewSprigHandler()

	for _, opt := range opts {
		err := opt(sprigHandler)
		if err != nil {
			sprigHandler.Logger().Error("cannot validate your option", "err", err)
		}
	}

	return sprigHandler.Build()
}
