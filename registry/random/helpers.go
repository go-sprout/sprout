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
//	size int - the length of the string to generate.
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
func (rr *RandomRegistry) randomString(size int, opts *randomOpts) string {
	if size <= 0 {
		return ""
	}

	if len(opts.withChars) == 0 {
		if opts.withAscii {
			opts.withChars = make([]rune, 95) // 95 printable ASCII characters
			for i := 32; i <= 126; i++ {
				opts.withChars[i-32] = rune(i)
			}
		} else {
			var buffer []rune
			if opts.withLetters {
				buffer = make([]rune, 52) // 26 lowercase + 26 uppercase letters
				for i := 0; i < 26; i++ {
					buffer[i] = 'a' + rune(i)
					buffer[i+26] = 'A' + rune(i)
				}
			}
			if opts.withNumbers {
				if buffer == nil {
					buffer = make([]rune, 10) // 10 digits
				} else {
					buffer = append(buffer, make([]rune, 10)...)
				}
				for i := 0; i < 10; i++ {
					buffer[len(buffer)-10+i] = '0' + rune(i)
				}
			}
			opts.withChars = buffer
		}
	}

	var builder strings.Builder
	builder.Grow(size)

	maxIndex := big.NewInt(int64(len(opts.withChars)))
	for i := 0; i < size; i++ {
		index, _ := cryptorand.Int(cryptorand.Reader, maxIndex)
		builder.WriteRune(opts.withChars[index.Int64()])
	}

	return builder.String()
}
