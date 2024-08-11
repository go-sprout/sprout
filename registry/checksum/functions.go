package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/adler32"
)

// Sha1sum calculates the SHA-1 hash of the input string and returns it as a
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
// {{ sha1sum "Hello, World!" }} // Output: 0a0a9f2a6772942557ab5355d76af442f8f65e01
func (cr *ChecksumRegistry) Sha1sum(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

// Sha256sum calculates the SHA-256 hash of the input string and returns it as a
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
// {{ sha256sum "Hello, World!" }} // Output: dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f
func (cr *ChecksumRegistry) Sha256sum(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// Adler32sum calculates the Adler-32 checksum of the input string and returns
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
// {{ adler32sum "Hello, World!" }} // Output: 1f9e046a
func (cr *ChecksumRegistry) Adler32sum(input string) string {
	hash := adler32.Checksum([]byte(input))
	return fmt.Sprintf("%d", hash)
}

// Md5sum calculates the MD5 hash of the input string and returns it as a
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
// {{ md5sum "Hello, World!" }} // Output: 65a8e27d8879283831b664bd8b7f0ad4
func (cr *ChecksumRegistry) Md5sum(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
