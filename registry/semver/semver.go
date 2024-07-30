package semver

import "github.com/go-sprout/sprout"

type SemverRegistry struct {
	handler *sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *SemverRegistry {
	return &SemverRegistry{}
}

// Uid returns the unique identifier of the registry.
func (sr *SemverRegistry) Uid() string {
	return "semver"
}

// LinkHandler links the handler to the registry at runtime.
func (sr *SemverRegistry) LinkHandler(fh sprout.Handler) {
	sr.handler = &fh
}

// RegisterFunctions registers all functions of the registry.
func (br *SemverRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) {
	sprout.AddFunction(funcsMap, "semver", br.Semver)
	sprout.AddFunction(funcsMap, "semverCompare", br.SemverCompare)
}
