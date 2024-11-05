package checksum

import "github.com/go-sprout/sprout"

type ChecksumRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of the checksum registry.
func NewRegistry() *ChecksumRegistry {
	return &ChecksumRegistry{}
}

// UID returns the unique identifier of the registry.
func (cr *ChecksumRegistry) UID() string {
	return "go-sprout/sprout.checksum"
}

// LinkHandler links the handler to the registry at runtime.
func (cr *ChecksumRegistry) LinkHandler(fh sprout.Handler) error {
	cr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (cr *ChecksumRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "sha1Sum", cr.SHA1Sum)
	sprout.AddFunction(funcsMap, "sha256Sum", cr.SHA256Sum)
	sprout.AddFunction(funcsMap, "sha512Sum", cr.SHA512Sum)
	sprout.AddFunction(funcsMap, "adler32Sum", cr.Adler32Sum)
	sprout.AddFunction(funcsMap, "md5Sum", cr.MD5Sum)
	return nil
}

func (cr *ChecksumRegistry) RegisterAliases(aliasMap sprout.FunctionAliasMap) error {
	sprout.AddAlias(aliasMap, "sha1Sum", "sha1sum")
	sprout.AddAlias(aliasMap, "sha256Sum", "sha256sum")
	sprout.AddAlias(aliasMap, "adler32Sum", "adler32sum")
	sprout.AddAlias(aliasMap, "md5Sum", "md5sum")
	return nil
}

func (cr *ChecksumRegistry) RegisterNotices(notices *[]sprout.FunctionNotice) error {
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("sha1sum", "use `sha1Sum` instead."))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("sha256sum", "use `sha256Sum` instead."))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("adler32sum", "use `adler32Sum` instead."))
	sprout.AddNotice(notices, sprout.NewDeprecatedNotice("md5sum", "use `md5Sum` instead."))
	return nil
}
