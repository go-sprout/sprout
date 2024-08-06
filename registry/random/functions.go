package random

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	mathrand "math/rand"
)

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
func (rr *RandomRegistry) RandAlphaNumeric(count int) string {
	return rr.randomString(count, &randomOpts{withLetters: true, withNumbers: true})
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
func (rr *RandomRegistry) RandAlpha(count int) string {
	return rr.randomString(count, &randomOpts{withLetters: true})
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
func (rr *RandomRegistry) RandAscii(count int) string {
	return rr.randomString(count, &randomOpts{withAscii: true})
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
func (rr *RandomRegistry) RandNumeric(count int) string {
	return rr.randomString(count, &randomOpts{withNumbers: true})
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
func (rr *RandomRegistry) RandBytes(count int) (string, error) {
	if count <= 0 {
		return "", nil
	}

	buf := make([]byte, count)
	_, _ = cryptorand.Read(buf)
	return base64.StdEncoding.EncodeToString(buf), nil
}

func (rr *RandomRegistry) RandInt(min, max int) int {
	return mathrand.Intn(max-min) + min
}