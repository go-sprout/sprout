package sprigin

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	helloFunc, ok := funcMap["hello"].(func() string)
	assert.True(t, ok)

	assert.Equal(t, "Hello!", helloFunc())
}
