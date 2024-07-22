package strings

import (
	cryptorand "crypto/rand"
	"math/big"
	mathrand "math/rand"
	"time"

	"github.com/go-sprout/sprout/registry"
)

// caseStyle defines the rules for transforming strings based on capitalization,
// separator insertion, and case enforcement. This struct is typically used to
// configure functions that modify the case and formatting of strings to match
// specific coding or display conventions.
//
// Fields:
//
//	Separator rune - The character used to separate words in the transformed string.
//	                 For example, underscores (_) or hyphens (-).
//
//	CapitalizeNext bool - Determines if the first character of each word should be
//	                      capitalized in the output. Useful for TitleCase or CamelCase.
//
//	ForceLowercase bool - If set to true, all characters in the output are converted
//	                      to lowercase, overriding any capitalization rules.
//
// Usage:
//
// This struct can be used to configure string transformation functions, allowing for
// flexible adaptation to various text formatting needs, such as generating identifiers
// or user-friendly display text.
//
// Example:
//
//		style := caseStyle{
//		    Separator:       '_',
//		    CapitalizeNext:  false,
//		    ForceLowercase:  true,
//		    InsertSeparator: true,
//		}
//	 Use `style` to transform "ExampleText" to "example_text"
type caseStyle struct {
	Separator       rune // Character that separates words.
	CapitalizeFirst bool // Whether to capitalize the first character of the string.
	CapitalizeNext  bool // Whether to capitalize the first character of each word.
	ForceLowercase  bool // Whether to force all characters to lowercase.
	ForceUppercase  bool // Whether to force all characters to uppercase.
}

// randSource is a global variable that provides a source of randomness seeded with
// a cryptographically secure random number. This source is used throughout various
// random generation functions to ensure that randomness is both fast and non-repetitive.
var randSource mathrand.Source

var (
	camelCaseStyle    = caseStyle{Separator: -1, CapitalizeNext: true, CapitalizeFirst: false, ForceLowercase: true}
	kebabCaseStyle    = caseStyle{Separator: '-', ForceLowercase: true}
	pascalCaseStyle   = caseStyle{Separator: -1, CapitalizeFirst: true, CapitalizeNext: true}
	snakeCaseStyle    = caseStyle{Separator: '_', ForceLowercase: true}
	dotCaseStyle      = caseStyle{Separator: '.', ForceLowercase: true}
	pathCaseStyle     = caseStyle{Separator: '/', ForceLowercase: true}
	constantCaseStyle = caseStyle{Separator: '_', ForceUppercase: true}
)

// init is an initialization function that seeds the global random source used
// in random string generation. It retrieves a secure timestamp-based seed from
// crypto/rand and uses it to initialize math/rand's source, ensuring that random
// values are not predictable across program restarts.
func init() {
	index, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(time.Now().UnixNano()))
	randSource = mathrand.NewSource(index.Int64())
}

type StringsRegistry struct {
	handler *registry.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of strings registry.
func NewRegistry() *StringsRegistry {
	return &StringsRegistry{}
}

// Uid returns the unique identifier of the registry.
func (sr *StringsRegistry) Uid() string {
	return "strings"
}

// LinkHandler links the handler to the registry at runtime.
func (sr *StringsRegistry) LinkHandler(fh registry.Handler) {
	sr.handler = &fh
}
