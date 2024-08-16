package checksum

import "github.com/go-sprout/sprout"

type ChecksumRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
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
func (cr *ChecksumRegistry) LinkHandler(fh sprout.Handler) error {
	cr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (cr *ChecksumRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "sha1sum", cr.SHA1Sum)
	sprout.AddFunction(funcsMap, "sha256sum", cr.SHA256Sum)
	sprout.AddFunction(funcsMap, "adler32sum", cr.Adler32Sum)
	sprout.AddFunction(funcsMap, "md5sum", cr.MD5Sum)
	return nil
}
