package sprout

import (
	"errors"
	"log/slog"
	"testing"

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
	registerNoticesMustCrash bool
}

var errMock = errors.New("mock error")

func (m *MockRegistry) UID() string {
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

type MockRegistryWithNotices struct {
	MockRegistry
}

func (m *MockRegistryWithNotices) RegisterNotices(notices *[]FunctionNotice) error {
	m.Called(notices)
	if m.registerNoticesMustCrash {
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
	mockRegistry.On("UID").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return(errMock)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistries(mockRegistry)
	require.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertNotCalled(t, "RegisterFunctions", mock.Anything)
}

// TestDefaultHandler_AddRegistries_Error tests the AddRegistry method of DefaultHandler when the registry returns an error.
func TestDefaultHandler_AddRegistry_Error_RegisterFuiesctions(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.registerFuncsMustCrash = true
	mockRegistry.On("UID").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return(errMock)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistries(mockRegistry)
	require.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

func TestDefaultHandler_AddRegistry_Error_RegisteriesAliases(t *testing.T) {
	mockRegistry := new(MockRegistryWithAlias)
	mockRegistry.registerAliasesMustCrash = true
	mockRegistry.On("UID").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()
	mockRegistry.On("RegisterAliases", mock.Anything).Return(errMock)

	dh := &DefaultHandler{
		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
	}

	err := dh.AddRegistries(mockRegistry)
	require.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertCalled(t, "RegisterAliases", dh.cachedFuncsAlias)
}

func TestDefaultHandler_AddRegistry_Error_RegisteriesNotices(t *testing.T) {
	mockRegistry := new(MockRegistryWithNotices)
	mockRegistry.registerNoticesMustCrash = true
	mockRegistry.On("UID").Return("mockRegistryWithNotices")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()
	mockRegistry.On("RegisterNotices", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
		notices: []FunctionNotice{
			*NewInfoNotice("", "amazing"),
		},
	}

	err := dh.AddRegistry(mockRegistry)
	require.Error(t, err, "AddRegistry should return an error")
	assert.Equal(t, errMock, err, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertCalled(t, "RegisterNotices", &dh.notices)
}

// TestDefaultHandler_AddRegistry tests the AddRegistry method of DefaultHandler.
func TestDefaultHandler_AddRegistry(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.On("UID").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistry(mockRegistry)
	require.NoError(t, err, "AddRegistry should not return an error")

	require.Len(t, dh.registries, 1, "Registry should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry, "Registry should match the mock registry")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

// TestDefaultHandler_AddRegistries tests the AddRegistries method of DefaultHandler.
func TestDefaultHandler_AddRegistries(t *testing.T) {
	mockRegistry1 := new(MockRegistry)
	mockRegistry1.On("UID").Return("mockRegistry1")
	mockRegistry1.On("LinkHandler", mock.Anything).Return()
	mockRegistry1.On("RegisterFunctions", mock.Anything).Return()

	mockRegistry2 := new(MockRegistry)
	mockRegistry2.On("UID").Return("mockRegistry2")
	mockRegistry2.On("LinkHandler", mock.Anything).Return()
	mockRegistry2.On("RegisterFunctions", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddRegistries(mockRegistry1, mockRegistry2)
	require.NoError(t, err, "AddRegistries should not return an error")

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
	mockRegistry.On("UID").Return("mockRegistryWithAlias")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()
	mockRegistry.On("RegisterAliases", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
	}

	err := dh.AddRegistry(mockRegistry)
	require.NoError(t, err, "AddRegistry should not return an error")

	require.Len(t, dh.registries, 1, "Registry should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry, "Registry should match the mock registry")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertCalled(t, "RegisterAliases", dh.cachedFuncsAlias)
}

func TestDefaultHandler_AddRegistryWithNotices(t *testing.T) {
	mockRegistry := new(MockRegistryWithNotices)
	mockRegistry.On("UID").Return("mockRegistryWithNotices")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()
	mockRegistry.On("RegisterNotices", mock.Anything).Return()

	dh := &DefaultHandler{
		cachedFuncsMap:   make(FunctionMap),
		cachedFuncsAlias: make(FunctionAliasMap),
		notices: []FunctionNotice{
			*NewInfoNotice("", "amazing"),
		},
	}

	err := dh.AddRegistry(mockRegistry)
	require.NoError(t, err, "AddRegistry should not return an error")

	require.Len(t, dh.notices, 1, "Registry should be added to the DefaultHandler")

	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertCalled(t, "RegisterNotices", &dh.notices)
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

// TestDefaultHandler_RawFunctions tests the Functions method of DefaultHandler.
func TestDefaultHandler_RawFunctions(t *testing.T) {
	funcsMap := make(FunctionMap)
	dh := &DefaultHandler{cachedFuncsMap: funcsMap}

	assert.Equal(t, funcsMap, dh.RawFunctions(), "Functions should return the correct FunctionMap")
}

// TestDefaultHandler_Aliases tests the Aliases method of DefaultHandler.
func TestDefaultHandler_Aliases(t *testing.T) {
	aliasesMap := make(FunctionAliasMap)
	dh := &DefaultHandler{cachedFuncsAlias: aliasesMap}

	assert.Equal(t, aliasesMap, dh.RawAliases(), "Aliases should return the correct FunctionAliasMap")
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

	assert.Equal(t, funcsMap, builtFuncsMap, "Build should return the correct FunctionMap")

	builtFuncsMapSecond := dh.Build()
	assert.Equal(t, builtFuncsMap, builtFuncsMapSecond, "Build should return the same FunctionMap on subsequent calls")
}

func TestDefaultHandler_safeWrapper(t *testing.T) {
	loggerHandler := &noticeLoggerHandler{}
	handler := New(WithLogger(slog.New(loggerHandler)))

	fn := func() (any, error) { return nil, errors.New("fail") }
	_, err := fn()
	require.Error(t, err, "fn should return an error")

	safeFn := safeWrapper(handler, "fn", fn)
	_, safeErr := safeFn()
	require.NoError(t, safeErr, "safeFn should not return an error")
	assert.Equal(t, "[ERROR] function call failed\n", loggerHandler.messages.String())
}

func TestSafeFuncName(t *testing.T) {
	assert.Equal(t, "safeFn", safeFuncName("fn"))
	assert.Equal(t, "safeFn", safeFuncName("Fn"))
	assert.Empty(t, safeFuncName(""))
}
