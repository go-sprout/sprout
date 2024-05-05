package sprout

import (
	"crypto/rand"
	cryptorand "crypto/rand"
	"encoding/base64"
	"math/big"
	mathrand "math/rand"
	"strings"
	"time"
)

// randSource is a global variable that provides a source of randomness seeded with
// a cryptographically secure random number. This source is used throughout various
// random generation functions to ensure that randomness is both fast and non-repetitive.
var randSource mathrand.Source

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

// init is an initialization function that seeds the global random source used
// in random string generation. It retrieves a secure timestamp-based seed from
// crypto/rand and uses it to initialize math/rand's source, ensuring that random
// values are not predictable across program restarts.
func init() {
	index, _ := cryptorand.Int(rand.Reader, big.NewInt(time.Now().UnixNano()))
	randSource = mathrand.NewSource(index.Int64())
}

// randomString generates a random string of a given length using specified options.
// It supports a flexible character set based on the provided options.
//
// Parameters:
//
//	count int - the length of the string to generate.
//	opts *randomOpts - options specifying character sets to include in the string.
//
// Returns:
//
//	string - the randomly generated string.
//
// Usage:
//
//	opts := &randomOpts{withLetters: true, withNumbers: true}
//	randomStr := fh.randomString(10, opts) // Generates a 10-character alphanumeric string.
func (fh *FunctionHandler) randomString(count int, opts *randomOpts) string {
	if count <= 0 {
		return ""
	}

	if len(opts.withChars) > 0 {
		goto GENERATE
	}

	if opts.withAscii {
		for i := 32; i <= 126; i++ {
			opts.withChars = append(opts.withChars, rune(i))
		}
	}

	if opts.withLetters {
		for i := 'a'; i <= 'z'; i++ {
			opts.withChars = append(opts.withChars, i)
		}
		for i := 'A'; i <= 'Z'; i++ {
			opts.withChars = append(opts.withChars, i)
		}
	}

	if opts.withNumbers {
		for i := '0'; i <= '9'; i++ {
			opts.withChars = append(opts.withChars, i)
		}
	}

GENERATE:
	var builder strings.Builder
	builder.Grow(count)

	for i := 0; i < count; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(opts.withChars))))
		builder.WriteRune(opts.withChars[index.Int64()])
	}

	return builder.String()
}

// RandAlphaNumeric generates a random alphanumeric string of specified length.
//
// Parameters:
//
//	count int - the length of the string to generate.
//
// Returns:
//
//	string - the randomly generated alphanumeric string.
//
// Example:
//
//	{{ 10 | randAlphaNumeric }} // Output: "a1b2c3d4e5" (output will vary)
func (fh *FunctionHandler) RandAlphaNumeric(count int) string {
	return fh.randomString(count, &randomOpts{withLetters: true, withNumbers: true})
}

// RandAlpha generates a random alphabetic string of specified length.
//
// Parameters:
//
//	count int - the length of the string to generate.
//
// Returns:
//
//	string - the randomly generated alphabetic string.
//
// Example:
//
//	{{ 10 | randAlpha }} // Output: "abcdefghij" (output will vary)
func (fh *FunctionHandler) RandAlpha(count int) string {
	return fh.randomString(count, &randomOpts{withLetters: true})
}

// RandAscii generates a random ASCII string (character codes 32 to 126) of specified length.
//
// Parameters:
//
//	count int - the length of the string to generate.
//
// Returns:
//
//	string - the randomly generated ASCII string.
//
// Example:
//
//	{{ 10 | randAscii }} // Output: "}]~>_<:^%" (output will vary)
func (fh *FunctionHandler) RandAscii(count int) string {
	return fh.randomString(count, &randomOpts{withAscii: true})
}

// RandNumeric generates a random numeric string of specified length.
//
// Parameters:
//
//	count int - the length of the string to generate.
//
// Returns:
//
//	string - the randomly generated numeric string.
//
// Example:
//
//	{{ 10 | randNumeric }} // Output: "0123456789" (output will vary)
func (fh *FunctionHandler) RandNumeric(count int) string {
	return fh.randomString(count, &randomOpts{withNumbers: true})
}

// RandBytes generates a random byte array of specified length and returns it as a base64 encoded string.
//
// Parameters:
//
//	count int - the number of bytes to generate.
//
// Returns:
//
//	string - the base64 encoded string of the randomly generated bytes.
//	error - error if the random byte generation fails.
//
// Example:
//
//	{{ 16 | randBytes }} // Output: "c3RhY2thYnVzZSByb2NrcyE=" (output will vary)
func (fh *FunctionHandler) RandBytes(count int) (string, error) {
	if count <= 0 {
		return "", nil
	}

	buf := make([]byte, count)
	if _, err := cryptorand.Read(buf); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf), nil
}

func (fh *FunctionHandler) RandInt(min, max int) int {
	return mathrand.Intn(max-min) + min
}
