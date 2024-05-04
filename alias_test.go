package sprout

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestWithAlias checks that aliases are correctly added to a function.
func TestWithAlias(t *testing.T) {
	handler := NewFunctionHandler()
	originalFunc := "originalFunc"
	alias1 := "alias1"
	alias2 := "alias2"

	// Apply the WithAlias option with two aliases.
	WithAlias(originalFunc, alias1, alias2)(handler)

	// Check that the aliases were added.
	assert.Contains(t, handler.funcsAlias, originalFunc)
	assert.Contains(t, handler.funcsAlias[originalFunc], alias1)
	assert.Contains(t, handler.funcsAlias[originalFunc], alias2)
	assert.Len(t, handler.funcsAlias[originalFunc], 2, "there should be exactly 2 aliases")
}

func TestWithAlias_Empty(t *testing.T) {
	handler := NewFunctionHandler()
	originalFunc := "originalFunc"

	// Apply the WithAlias option with no aliases.
	WithAlias(originalFunc)(handler)

	// Check that no aliases were added.
	assert.NotContains(t, handler.funcsAlias, originalFunc)
}

func TestWithAliases(t *testing.T) {
	handler := NewFunctionHandler()
	originalFunc1 := "originalFunc1"
	alias1 := "alias1"
	alias2 := "alias2"
	originalFunc2 := "originalFunc2"
	alias3 := "alias3"

	// Apply the WithAliases option with two sets of aliases.
	WithAliases(FunctionAliasMap{
		originalFunc1: {alias1, alias2},
		originalFunc2: {alias3},
	})(handler)

	// Check that the aliases were added.
	assert.Contains(t, handler.funcsAlias, originalFunc1)
	assert.Contains(t, handler.funcsAlias[originalFunc1], alias1)
	assert.Contains(t, handler.funcsAlias[originalFunc1], alias2)
	assert.Len(t, handler.funcsAlias[originalFunc1], 2, "there should be exactly 2 aliases")

	assert.Contains(t, handler.funcsAlias, originalFunc2)
	assert.Contains(t, handler.funcsAlias[originalFunc2], alias3)
	assert.Len(t, handler.funcsAlias[originalFunc2], 1, "there should be exactly 1 alias")
}

// TestRegisterAliases checks that aliases are correctly registered in the function map.
func TestRegisterAliases(t *testing.T) {
	handler := NewFunctionHandler()
	originalFunc := "originalFunc"
	alias1 := "alias1"
	alias2 := "alias2"

	// Mock a function for originalFunc and add it to funcMap.
	mockFunc := func() {}
	handler.funcMap[originalFunc] = mockFunc

	// Apply the WithAlias option and then register the aliases.
	WithAlias(originalFunc, alias1, alias2)(handler)
	handler.registerAliases()

	// Check that the aliases are mapped to the same function as the original function in funcMap.
	assert.True(t, reflect.ValueOf(handler.funcMap[originalFunc]).Pointer() == reflect.ValueOf(handler.funcMap[alias1]).Pointer())
	assert.True(t, reflect.ValueOf(handler.funcMap[originalFunc]).Pointer() == reflect.ValueOf(handler.funcMap[alias2]).Pointer())
}

func TestAliasesInTemplate(t *testing.T) {
	handler := NewFunctionHandler()
	originalFuncName := "originalFunc"
	alias1 := "alias1"
	alias2 := "alias2"

	// Mock a function for originalFunc and add it to funcMap.
	mockFunc := func() string { return "cheese" }
	handler.funcMap[originalFuncName] = mockFunc

	// Apply the WithAlias option and then register the aliases.
	WithAlias(originalFuncName, alias1, alias2)(handler)

	// Create a template with the aliases.
	result, err := runTemplate(t, handler, `{{originalFunc}} {{alias1}} {{alias2}}`, nil)
	assert.NoError(t, err)
	assert.Equal(t, "cheese cheese cheese", result)
}
