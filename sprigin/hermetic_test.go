package sprigin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/env"
	"github.com/go-sprout/sprout/registry/random"
)

// registryFunctionNames returns the template names registered by a registry.
func registryFunctionNames(t *testing.T, registry sprout.Registry) []string {
	t.Helper()

	funcs := sprout.FunctionMap{}
	require.NoError(t, registry.RegisterFunctions(funcs))

	names := make([]string, 0, len(funcs))
	for name := range funcs {
		names = append(names, name)
	}
	return names
}

// TestHermeticFuncMapExcludesNonDeterministicFunctions guards the hermetic func
// maps against a gap that already happened once: nonhermeticFunctions was
// written against the sprig names, so `expandEnv` stayed reachable while its
// `expandenv` alias was filtered out. Deriving the expectation from the
// registries makes any new env or random function fail here until it is
// classified.
func TestHermeticFuncMapExcludesNonDeterministicFunctions(t *testing.T) {
	hermetic := HermeticTxtFuncMap()

	t.Run("no function of the env and random registries", func(t *testing.T) {
		for _, registry := range []sprout.Registry{env.NewRegistry(), random.NewRegistry()} {
			for _, name := range registryFunctionNames(t, registry) {
				assert.NotContains(t, hermetic, name, "`%s` reaches the environment or the random source", name)
			}
		}
	})

	t.Run("no clock or randomness dependent function", func(t *testing.T) {
		for _, name := range []string{"now", "date", "dateAgo", "dateInZone", "htmlDate", "htmlDateInZone", "uuidv4", "uuidv7", "getHostByName"} {
			assert.NotContains(t, hermetic, name, "`%s` does not evaluate to the same result for a given input", name)
		}
	})

	t.Run("deterministic functions are kept", func(t *testing.T) {
		for _, name := range []string{"uuidv5", "uuidv3", "uuidNil", "isUUID", "uuidVersion", "toLocalDate", "fromUnix", "trim", "toJSON"} {
			assert.Contains(t, hermetic, name, "`%s` is deterministic and must stay available", name)
		}
	})
}
