package checksum

import "github.com/go-sprout/sprout"

type ChecksumRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of the checksum registry.
func NewRegistry() *ChecksumRegistry {
	return &ChecksumRegistry{}
}

// Uid returns the unique identifier of the registry.
func (cr *ChecksumRegistry) Uid() string {
	return "checksum"
}

// LinkHandler links the handler to the registry at runtime.
func (cr *ChecksumRegistry) LinkHandler(fh sprout.Handler) {
	cr.handler = &fh
}

// RegisterFunctions registers all functions of the registry.
func (cr *ChecksumRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "sha1sum", cr.Sha1sum)
	sprout.AddFunction(funcsMap, "sha256sum", cr.Sha256sum)
	sprout.AddFunction(funcsMap, "adler32sum", cr.Adler32sum)
	sprout.AddFunction(funcsMap, "md5sum", cr.Md5sum)
}
