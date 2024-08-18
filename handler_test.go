package sprout

import (
	"errors"
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock implementations for the Registry and RegistryWithAlias interfaces

type MockRegistry struct {
	mock.Mock
	linkHandlerMustCrash     bool
	registerFuncsMustCrash   bool
	registerAliasesMustCrash bool
}

var errMock = errors.New("mock error")

func (m *MockRegistry) Uid() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockRegistry) LinkHandler(fh Handler) error {
	m.Called(fh)
	if m.linkHandlerMustCrash {
		return errMock
	}
	return nil
}

func (m *MockRegistry) RegisterFunctions(fnMap FunctionMap) error {
	m.Called(fnMap)
	if m.registerFuncsMustCrash {
		return errMock
	}
	return nil
}

type MockRegistryWithAlias struct {
	MockRegistry
}

func (m *MockRegistryWithAlias) RegisterAliases(aliasMap FunctionAliasMap) error {
	m.Called(aliasMap)
	if m.registerAliasesMustCrash {
		return errMock
	}
	return nil
}

// TestDefaultHandler_Logger tests the Logger method of DefaultHandler.
func TestDefaultHandler_Logger(t *testing.T) {
	logger := slog.New(&slog.TextHandler{})
	dh := &DefaultHandler{logger: logger}

	assert.Equal(t, logger, dh.Logger(), "Logger should return the initialized logger")
}

// TestDefaultHandler_AddRegistries_Error tests the AddRegistry method of DefaultHandler when the registry returns an error.
func TestDefaultHandler_AddRegistries_Error(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.linkHandlerMustCrash = true
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return(errMock)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistries(mockRegistry)
	assert.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertNotCalled(t, "RegisterFunctions", mock.Anything)
}

// TestDefaultHandler_AddRegistries_Error tests the AddRegistry method of DefaultHandler when the registry returns an error.
func TestDefaultHandler_AddRegistry_Error_RegisterFuiesctions(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.registerFuncsMustCrash = true
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return(errMock)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistries(mockRegistry)
	assert.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

// TestDefaultHandler_AddRegistries_Error tests the AddRegistry method of DefaultHandler when the registry returns an error.
func TestDefaultHandler_AddRegistry_Error_Registeriesliases(t *testing.T) {
	mockRegistry := new(MockRegistryWithAlias)
	mockRegistry.registerAliasesMustCrash = true
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()
	mockRegistry.On("RegisterAliases", mock.Anything).Return(errMock)

	dh := &DefaultHandler{
		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
	}

	err := dh.AddRegistries(mockRegistry)
	assert.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertCalled(t, "RegisterAliases", dh.cachedFuncsAlias)
}

// TestDefaultHandler_AddRegistry tests the AddRegistry method of DefaultHandler.
func TestDefaultHandler_AddRegistry(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistry(mockRegistry)
	assert.NoError(t, err, "AddRegistry should not return an error")

	require.Len(t, dh.registries, 1, "Registry should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry, "Registry should match the mock registry")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

// TestDefaultHandler_AddRegistries tests the AddRegistries method of DefaultHandler.
func TestDefaultHandler_AddRegistries(t *testing.T) {
	mockRegistry1 := new(MockRegistry)
	mockRegistry1.On("Uid").Return("mockRegistry1")
	mockRegistry1.On("LinkHandler", mock.Anything).Return()
	mockRegistry1.On("RegisterFunctions", mock.Anything).Return()

	mockRegistry2 := new(MockRegistry)
	mockRegistry2.On("Uid").Return("mockRegistry2")
	mockRegistry2.On("LinkHandler", mock.Anything).Return()
	mockRegistry2.On("RegisterFunctions", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistries(mockRegistry1, mockRegistry2)
	assert.NoError(t, err, "AddRegistries should not return an error")

	require.Len(t, dh.registries, 2, "Both registries should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry1, "First registry should match mockRegistry1")
	assert.Contains(t, dh.registries, mockRegistry2, "Second registry should match mockRegistry2")

	mockRegistry1.AssertCalled(t, "LinkHandler", dh)
	mockRegistry1.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry2.AssertCalled(t, "LinkHandler", dh)
	mockRegistry2.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

// TestDefaultHandler_AddRegistryWithAlias tests AddRegistry when the registry also implements RegistryWithAlias.
func TestDefaultHandler_AddRegistryWithAlias(t *testing.T) {
	mockRegistry := new(MockRegistryWithAlias)
	mockRegistry.On("Uid").Return("mockRegistryWithAlias")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()
	mockRegistry.On("RegisterAliases", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
	}

	err := dh.AddRegistry(mockRegistry)
	assert.NoError(t, err, "AddRegistry should not return an error")

	require.Len(t, dh.registries, 1, "Registry should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry, "Registry should match the mock registry")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertCalled(t, "RegisterAliases", dh.cachedFuncsAlias)
}

func TestDefaultHandler_Registries(t *testing.T) {
	dh := &DefaultHandler{
		registries: []Registry{
			new(MockRegistry),
			new(MockRegistry),
		},
	}

	assert.Len(t, dh.registries, 2, "Registries should return the correct number of registries")
}

// TestDefaultHandler_Functions tests the Functions method of DefaultHandler.
func TestDefaultHandler_Functions(t *testing.T) {
	funcsMap := make(FunctionMap)
	dh := &DefaultHandler{cachedFuncsMap: funcsMap}

	assert.Equal(t, funcsMap, dh.Functions(), "Functions should return the correct FunctionMap")
}

// TestDefaultHandler_Aliases tests the Aliases method of DefaultHandler.
func TestDefaultHandler_Aliases(t *testing.T) {
	aliasesMap := make(FunctionAliasMap)
	dh := &DefaultHandler{cachedFuncsAlias: aliasesMap}

	assert.Equal(t, aliasesMap, dh.Aliases(), "Aliases should return the correct FunctionAliasMap")
}

// TestDefaultHandler_Build tests the Build method of DefaultHandler.
func TestDefaultHandler_Build(t *testing.T) {
	funcsMap := make(FunctionMap)
	aliasesMap := make(FunctionAliasMap)

	dh := &DefaultHandler{
		cachedFuncsMap:   funcsMap,
		cachedFuncsAlias: aliasesMap,
	}

	// Ensure aliases are correctly processed before returning the registry
	AssignAliases(dh)

	builtFuncsMap := dh.Build()

	// TODO: Ensure the test test the value of
	assert.Len(t, funcsMap, len(builtFuncsMap), "Build should return the correct FunctionMap")
}
