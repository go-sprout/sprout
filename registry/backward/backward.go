/**
 * This file lists the functions originally part of the Sprig library that are
 * intentionally excluded from the Sprout library. The exclusions are based on\
 * community decisions and technical evaluations aimed at enhancing security,
 * relevance, and performance in the context of Go templates.
 * Each exclusion is supported by rational and further community discussions
 * can be found on our GitHub issues page.
 *
 * Exclusion Criteria:
 * 1. Crypto functions: Deemed inappropriate for Go templates due to inherent security risks.
 * 2. Irrelevant functions: Omitted because they do not provide utility in the context of Go templates.
 * 3. Deprecated/Insecure: Functions using outdated or insecure standards are excluded.
 * 4. Temporary exclusions: Certain functions are temporarily excluded to prevent breaking changes,
 *    pending the implementation of the new loader feature.
 * 5. Community decision: Choices made by the community are documented and can be discussed at
 *    https://github.com/go-sprout/sprout/issues/1.
 *
 * The Sprout library is an open-source project and welcomes contributions from the community.
 * To discuss existing exclusions or propose new ones, please contribute to the discussions on
 * our GitHub repository.
 */
package backward

import (
	"github.com/go-sprout/sprout"
)

type BackwardCompatibilityRegistry struct {
	handler sprout.Handler // Embedding Handler for shared functionality
}

// NewRegistry creates a new instance of your registry with the embedded Handler.
func NewRegistry() *BackwardCompatibilityRegistry {
	return &BackwardCompatibilityRegistry{}
}

// Uid returns the unique identifier of the registry.
func (bcr *BackwardCompatibilityRegistry) Uid() string {
	return "go-sprout/sprout.backwardcompatibilitywithsprig"
}

// LinkHandler links the handler to the registry at runtime.
func (bcr *BackwardCompatibilityRegistry) LinkHandler(fh sprout.Handler) error {
	bcr.handler = fh
	return nil
}

// RegisterFunctions registers all functions of the registry.
func (bcr *BackwardCompatibilityRegistry) RegisterFunctions(funcsMap sprout.FunctionMap) error {
	sprout.AddFunction(funcsMap, "fail", bcr.Fail)
	sprout.AddFunction(funcsMap, "urlParse", bcr.UrlParse)
	sprout.AddFunction(funcsMap, "urlJoin", bcr.UrlJoin)
	sprout.AddFunction(funcsMap, "getHostByName", bcr.GetHostByName)
	return nil
}
