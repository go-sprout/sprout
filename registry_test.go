package sprout

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFunction(t *testing.T) {
	funcsMap := make(FunctionMap)

	// Define two different functions
	testFunc1 := func() string { return "Hello, World!" }
	testFunc2 := func() string { return "Should Not Overwrite" }
	testFunc3 := func() string { return "Should Be Defined" }

	// Test adding a new function
	AddFunction(funcsMap, "testFunc", testFunc1)
	assert.Contains(t, funcsMap, "testFunc", "Function 'testFunc' should be added to the FunctionMap")

	// Call the function and check the result
	result := funcsMap["testFunc"].(func() string)()
	assert.Equal(t, "Hello, World!", result, "Function 'testFunc' should return 'Hello, World!'")

	// Test trying to overwrite an existing function
	AddFunction(funcsMap, "testFunc", testFunc2)
	result = funcsMap["testFunc"].(func() string)()
	assert.Equal(t, "Hello, World!", result, "Function 'testFunc' should not be overwritten and should still return 'Hello, World!'")

	// Test adding a new function after the previous one
	AddFunction(funcsMap, "testFunc2", testFunc3)
	result = funcsMap["testFunc2"].(func() string)()
	assert.Equal(t, "Should Be Defined", result, "Function 'testFunc2' should be added and return 'Should Be Defined'")
}

func TestAddAlias(t *testing.T) {
	aliasMap := make(FunctionAliasMap)

	// Test adding zero aliases
	AddAlias(aliasMap, "originalFunc")
	assert.NotContains(t, aliasMap, "originalFunc", "No aliases should be added if none are provided")

	// Test adding aliases for an existing function
	AddAlias(aliasMap, "originalFunc", "alias1", "alias2")
	assert.Contains(t, aliasMap, "originalFunc", "Aliases should be added under 'originalFunc'")
	assert.Equal(t, []string{"alias1", "alias2"}, aliasMap["originalFunc"], "Aliases should be correctly registered")

	// Test adding more aliases to an existing function
	AddAlias(aliasMap, "originalFunc", "alias3")
	assert.Equal(t, []string{"alias1", "alias2", "alias3"}, aliasMap["originalFunc"], "New alias should be added to the existing aliases")

	// Test adding an alias to a function that doesn't exist in the map yet
	AddAlias(aliasMap, "nonExistentFunc", "aliasX")
	assert.Contains(t, aliasMap, "nonExistentFunc", "Aliases should be added under 'nonExistentFunc' even if the function doesn't exist")
}

func TestWithRegistry(t *testing.T) {
	// Define a registry with a function and an alias
	mockRegistry := new(MockRegistry)
	mockRegistry.linkHandlerMustCrash = true
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return(errMock)

	// Create a handler with the registry
	handler := New(WithRegistry(mockRegistry))
	handler.Build()

	// Check that the function and alias are present in the handler
	assert.Contains(t, handler.registries, mockRegistry, "Registry should be added to the handler")
}

func TestWithRegistries(t *testing.T) {
	// Define two registries with functions and aliases
	mockRegistry1 := new(MockRegistry)
	mockRegistry1.On("Uid").Return("mockRegistry1")
	mockRegistry1.On("LinkHandler", mock.Anything).Return(nil)
	mockRegistry1.On("RegisterFunctions", mock.Anything).Return(nil)

	mockRegistry2 := new(MockRegistry)
	mockRegistry2.linkHandlerMustCrash = true
	mockRegistry2.On("Uid").Return("mockRegistry2")
	mockRegistry2.On("LinkHandler", mock.Anything).Return(nil)
	mockRegistry1.On("RegisterFunctions", mock.Anything).Return(nil)

	// Create a handler with the registries
	handler := New(WithRegistries(mockRegistry1, mockRegistry2))
	handler.Build()

	// Check that the functions and aliases are present in the handler
	assert.Contains(t, handler.registries, mockRegistry1, "Registry 1 should be added to the handler")
	assert.Contains(t, handler.registries, mockRegistry2, "Registry 2 should be added to the handler")
}
