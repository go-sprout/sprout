package sprigin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sprigFunctionCount = 203

func TestSprig_backward_compatibility(t *testing.T) {
	gfm := GenericFuncMap()
	assert.NotNil(t, gfm)
	assert.GreaterOrEqual(t, len(gfm), sprigFunctionCount)

	hfm := HtmlFuncMap()
	assert.NotNil(t, hfm)
	assert.GreaterOrEqual(t, len(hfm), sprigFunctionCount)

	tfm := TxtFuncMap()
	assert.NotNil(t, tfm)
	assert.GreaterOrEqual(t, len(tfm), sprigFunctionCount)

	hhfm := HermeticHtmlFuncMap()
	assert.NotNil(t, hhfm)
	assert.Equal(t, len(hhfm), len(gfm)-len(nonhermeticFunctions))

	htfm := HermeticTxtFuncMap()
	assert.NotNil(t, htfm)
	assert.Equal(t, len(htfm), len(gfm)-len(nonhermeticFunctions))
}

func TestFuncMap_IncludesHello(t *testing.T) {
	funcMap := FuncMap()

	_, exists := funcMap["hello"]
	assert.True(t, exists)

	helloFunc, ok := funcMap["hello"].(func(...any) (any, error))
	assert.True(t, ok)

	result, err := helloFunc()
	require.NoError(t, err)
	assert.Equal(t, "Hello!", result)
}

func TestSprigHandler(t *testing.T) {
	handler := NewSprigHandler()

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.Logger())

	handler.Build()

	assert.GreaterOrEqual(t, len(handler.RawFunctions()), sprigFunctionCount)
	assert.Len(t, handler.RawAliases(), 37) // Hardcoded for backward compatibility

	assert.Len(t, handler.registries, 18) // Hardcoded for backward compatibility

	registriesUids := []string{}
	for _, registry := range handler.registries {
		registriesUids = append(registriesUids, registry.Uid())
	}

	assert.ElementsMatch(t, registriesUids, []string{
		"go-sprout/sprout.std",
		"go-sprout/sprout.uniqueid",
		"go-sprout/sprout.semver",
		"go-sprout/sprout.backwardcompatibilitywithsprig",
		"go-sprout/sprout.reflect",
		"go-sprout/sprout.time",
		"go-sprout/sprout.strings",
		"go-sprout/sprout.random",
		"go-sprout/sprout.checksum",
		"go-sprout/sprout.conversion",
		"go-sprout/sprout.numeric",
		"go-sprout/sprout.encoding",
		"go-sprout/sprout.regexp",
		"go-sprout/sprout.slices",
		"go-sprout/sprout.maps",
		"go-sprout/sprout.crypto",
		"go-sprout/sprout.filesystem",
		"go-sprout/sprout.env",
	})
}
