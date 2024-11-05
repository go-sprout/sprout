package sprout

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewRegistryGroup(t *testing.T) {
	group := NewRegistryGroup()

	require.NotNil(t, group)
	require.NotNil(t, group.Registries)
}

func TestDefaultHandler_AddGroups(t *testing.T) {
	mockRegistry1 := new(MockRegistry)
	mockRegistry1.On("Uid").Return("mockRegistry1")
	mockRegistry1.On("LinkHandler", mock.Anything).Return()
	mockRegistry1.On("RegisterFunctions", mock.Anything).Return()

	mockRegistry2 := new(MockRegistry)
	mockRegistry2.On("Uid").Return("mockRegistry2")
	mockRegistry2.On("LinkHandler", mock.Anything).Return()
	mockRegistry2.On("RegisterFunctions", mock.Anything).Return()

	group1 := NewRegistryGroup(mockRegistry1)
	group2 := NewRegistryGroup(mockRegistry2)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddGroups(group1, group2)
	require.NoError(t, err)

	mockRegistry1.AssertExpectations(t)
	mockRegistry2.AssertExpectations(t)

	require.Len(t, dh.registries, 2, "Both registries should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry1, "First registry should match mockRegistry1")
	assert.Contains(t, dh.registries, mockRegistry2, "Second registry should match mockRegistry2")

	mockRegistry1.AssertCalled(t, "LinkHandler", dh)
	mockRegistry1.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry2.AssertCalled(t, "LinkHandler", dh)
	mockRegistry2.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

func TestDefaultHandler_AddGroups_Error(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.linkHandlerMustCrash = true
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return(errMock)

	group1 := NewRegistryGroup(mockRegistry)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddGroups(group1)
	require.ErrorIs(t, err, errMock, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertNotCalled(t, "RegisterFunctions", mock.Anything)
}

func TestDefaultHandler_AddGroups_MultiplesTimes(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.On("Uid").Return("mockRegistry1")
	mockRegistry.On("LinkHandler", mock.Anything).Return()
	mockRegistry.On("RegisterFunctions", mock.Anything).Return()

	group1 := NewRegistryGroup(mockRegistry, mockRegistry)
	group2 := NewRegistryGroup(mockRegistry, mockRegistry)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := dh.AddGroups(group1, group2)
	require.NoError(t, err)

	err = dh.AddGroups(group1)
	require.NoError(t, err)

	mockRegistry.AssertExpectations(t)

	require.Len(t, dh.registries, 1, "Only one registry should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry, "Registry should match mockRegistry")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry.AssertNumberOfCalls(t, "LinkHandler", 1)
	mockRegistry.AssertNumberOfCalls(t, "RegisterFunctions", 1)
}

func TestWithGroups(t *testing.T) {
	mockRegistry1 := new(MockRegistry)
	mockRegistry1.On("Uid").Return("mockRegistry1")
	mockRegistry1.On("LinkHandler", mock.Anything).Return()
	mockRegistry1.On("RegisterFunctions", mock.Anything).Return()

	mockRegistry2 := new(MockRegistry)
	mockRegistry2.On("Uid").Return("mockRegistry2")
	mockRegistry2.On("LinkHandler", mock.Anything).Return()
	mockRegistry2.On("RegisterFunctions", mock.Anything).Return()

	group1 := NewRegistryGroup(mockRegistry1)
	group2 := NewRegistryGroup(mockRegistry2)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := WithGroups(group1, group2)(dh)
	require.NoError(t, err)

	mockRegistry1.AssertExpectations(t)
	mockRegistry2.AssertExpectations(t)

	require.Len(t, dh.registries, 2, "Both registries should be added to the DefaultHandler")
	assert.Contains(t, dh.registries, mockRegistry1, "First registry should match mockRegistry1")
	assert.Contains(t, dh.registries, mockRegistry2, "Second registry should match mockRegistry2")

	mockRegistry1.AssertCalled(t, "LinkHandler", dh)
	mockRegistry1.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
	mockRegistry2.AssertCalled(t, "LinkHandler", dh)
	mockRegistry2.AssertCalled(t, "RegisterFunctions", dh.cachedFuncsMap)
}

func TestWithGroups_Error(t *testing.T) {
	mockRegistry := new(MockRegistry)
	mockRegistry.linkHandlerMustCrash = true
	mockRegistry.On("Uid").Return("mockRegistry")
	mockRegistry.On("LinkHandler", mock.Anything).Return(errMock)

	group1 := NewRegistryGroup(mockRegistry)

	dh := &DefaultHandler{
		cachedFuncsMap: make(FunctionMap),
	}

	err := WithGroups(group1)(dh)
	require.ErrorIs(t, err, errMock, "Error should match the mock error")

	mockRegistry.AssertCalled(t, "LinkHandler", dh)
	mockRegistry.AssertNotCalled(t, "RegisterFunctions", mock.Anything)
}
