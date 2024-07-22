package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/adler32"

	"github.com/go-sprout/sprout/registry"
)

// RegisterFunctions registers all functions of the registry.
func (cr *ChecksumRegistry) RegisterFunctions(funcsMap registry.FunctionMap) {
	registry.AddFunction(funcsMap, "sha1sum", cr.Sha1sum)
	registry.AddFunction(funcsMap, "sha256sum", cr.Sha256sum)
	registry.AddFunction(funcsMap, "adler32sum", cr.Adler32sum)
	registry.AddFunction(funcsMap, "md5sum", cr.Md5sum)
}

// ExampleFunction is a function that does something.
func (cr *ChecksumRegistry) Sha1sum(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (cr *ChecksumRegistry) Sha256sum(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (cr *ChecksumRegistry) Adler32sum(input string) string {
	hash := adler32.Checksum([]byte(input))
	return fmt.Sprintf("%d", hash)
}

func (cr *ChecksumRegistry) Md5sum(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
