package random

import (
	cryptorand "crypto/rand"
	"math/big"
	mathrand "math/rand"
	"time"

	"github.com/go-sprout/sprout"
)

// randomOpts defines options for generating random strings. These options specify
// which character sets to include in the random generation process. When you provide
// a set of chars with `withChars`, the other options are ignored.
//
// Fields:
//
//	withLetters bool - Includes lowercase and uppercase alphabetic characters if set to true.
//	withNumbers bool - Includes numeric characters (0-9) if set to true.
//	withAscii bool - Includes all printable ASCII characters (from space to tilde) if set to true.
//	withChars []rune - Allows for specifying an explicit list of characters to include in the generation.
//
// Usage:
//
//	opts := randomOpts{
//	    withLetters: true,
//	    withNumbers: true,
//	    withAscii: false,
//	    withChars: nil,
//	}
//
// Use these options in a random string generation function to produce a string
// consisting only of alphabetic and numeric characters.
type randomOpts struct {
	withLetters bool
	withNumbers bool
	withAscii   bool
	withChars   []rune
}

// randSource is a global variable that provides a source of randomness seeded with
// a cryptographically secure random number. This source is used throughout various
// random generation functions to ensure that randomness is both fast and non-repetitive.
var randSource mathrand.Source

// init is an initialization function that seeds the global random source used
// in random string generation. It retrieves a secure timestamp-based seed from
// crypto/rand and uses it to initialize math/rand's source, ensuring that random
// values are not predictable across program restarts.
func init() {
	index, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(time.Now().UnixNano()))
	randSource = mathrand.NewSource(index.Int64())
}

type RandomRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of regexp registry.
func NewRegistry() *RandomRegistry {
	return &RandomRegistry{}
}

// Uid returns the unique identifier of the registry.
func (rr *RandomRegistry) Uid() string {
	return "random"
}

// LinkHandler links the handler to the registry at runtime.
func (rr *RandomRegistry) LinkHandler(fh sprout.Handler) {
	rr.handler = &fh
}

// RegisterFunctions registers all functions of the registry.
func (rr *RandomRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "randAlphaNum", rr.RandAlphaNumeric)
	sprout.AddFunction(funcsMap, "randAlpha", rr.RandAlpha)
	sprout.AddFunction(funcsMap, "randAscii", rr.RandAscii)
	sprout.AddFunction(funcsMap, "randNumeric", rr.RandNumeric)
	sprout.AddFunction(funcsMap, "randBytes", rr.RandBytes)
	sprout.AddFunction(funcsMap, "randInt", rr.RandInt)
}
