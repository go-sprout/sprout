package random

import (
	cryptorand "crypto/rand"
	"math/big"
	"strings"
)

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
//	randomStr := rr.randomString(10, opts) // Generates a 10-character alphanumeric string.
func (rr *RandomRegistry) randomString(count int, opts *randomOpts) string {
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
		index, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(len(opts.withChars))))
		builder.WriteRune(opts.withChars[index.Int64()])
	}

	return builder.String()
}
