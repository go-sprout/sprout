package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/adler32"
)

// SHA1Sum calculates the SHA-1 hash of the input string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - input: the string to be hashed.
//
// Returns:
// - the SHA-1 hash of the input string as a hexadecimal encoded string.
//
// Example:
//
// {{ sha1Sum "Hello, World!" }} // Output: 0a0a9f2a6772942557ab5355d76af442f8f65e01
func (cr *ChecksumRegistry) SHA1Sum(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

// SHA256Sum calculates the SHA-256 hash of the input string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - input: the string to be hashed.
//
// Returns:
// - the SHA-256 hash of the input string as a hexadecimal encoded string.
//
// Example:
//
// {{ sha256Sum "Hello, World!" }} // Output: dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f
func (cr *ChecksumRegistry) SHA256Sum(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// SHA512Sum calculates the SHA-512 hash of the input string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - input: the string to be hashed.
//
// Returns:
// - the SHA-512 hash of the input string as a hexadecimal encoded string.
//
// Example:
//
// {{ sha512Sum "Hello, World!" }} // Output: 374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387
func (cr *ChecksumRegistry) SHA512Sum(input string) string {
	hash := sha512.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}

// Adler32Sum calculates the Adler-32 checksum of the input string and returns
// it as a hexadecimal encoded string.
//
// Parameters:
// - input: the string to be hashed.
//
// Returns:
// - the Adler-32 checksum of the input string as a hexadecimal encoded string.
//
// Example:
//
// {{ adler32Sum "Hello, World!" }} // Output: 1f9e046a
func (cr *ChecksumRegistry) Adler32Sum(input string) string {
	hash := adler32.Checksum([]byte(input))
	return fmt.Sprint(hash)
}

// MD5Sum calculates the MD5 hash of the input string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - input: the string to be hashed.
//
// Returns:
// - the MD5 hash of the input string as a hexadecimal encoded string.
//
// Example:
//
// {{ md5Sum "Hello, World!" }} // Output: 65a8e27d8879283831b664bd8b7f0ad4
func (cr *ChecksumRegistry) MD5Sum(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
