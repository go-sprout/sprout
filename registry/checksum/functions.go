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

// SHA1Sum calculates the SHA-1 hash of the value string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - value: the string to be hashed.
//
// Returns:
// - the SHA-1 hash of the value string as a hexadecimal encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: sha1Sum].
//
// [Sprout Documentation: sha1Sum]: https://docs.atom.codes/sprout/registries/checksum#sha1sum
func (cr *ChecksumRegistry) SHA1Sum(value string) string {
	hash := sha1.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

// SHA256Sum calculates the SHA-256 hash of the value string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - value: the string to be hashed.
//
// Returns:
// - the SHA-256 hash of the value string as a hexadecimal encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: sha256Sum].
//
// [Sprout Documentation: sha256Sum]: https://docs.atom.codes/sprout/registries/checksum#sha256sum
func (cr *ChecksumRegistry) SHA256Sum(value string) string {
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}

// SHA512Sum calculates the SHA-512 hash of the value string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - value: the string to be hashed.
//
// Returns:
// - the SHA-512 hash of the value string as a hexadecimal encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: sha512Sum].
//
// [Sprout Documentation: sha512Sum]: https://docs.atom.codes/sprout/registries/checksum#sha512sum
func (cr *ChecksumRegistry) SHA512Sum(value string) string {
	hash := sha512.Sum512([]byte(value))
	return hex.EncodeToString(hash[:])
}

// Adler32Sum calculates the Adler-32 checksum of the value string and returns
// it as a hexadecimal encoded string.
//
// Parameters:
// - value: the string to be hashed.
//
// Returns:
// - the Adler-32 checksum of the value string as a hexadecimal encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: adler32Sum].
//
// [Sprout Documentation: adler32Sum]: https://docs.atom.codes/sprout/registries/checksum#adler32sum
func (cr *ChecksumRegistry) Adler32Sum(value string) string {
	hash := adler32.Checksum([]byte(value))
	return fmt.Sprint(hash)
}

// MD5Sum calculates the MD5 hash of the value string and returns it as a
// hexadecimal encoded string.
//
// Parameters:
// - value: the string to be hashed.
//
// Returns:
// - the MD5 hash of the value string as a hexadecimal encoded string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: md5Sum].
//
// [Sprout Documentation: md5Sum]: https://docs.atom.codes/sprout/registries/checksum#md5sum
func (cr *ChecksumRegistry) MD5Sum(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}
