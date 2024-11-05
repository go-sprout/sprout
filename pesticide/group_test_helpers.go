package pesticide

import (
	"testing"

	"github.com/go-sprout/sprout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type GroupTestCase struct {
	RegistriesUIDs []string
}

func RunGroupTest(t *testing.T, rg *sprout.RegistryGroup, tc GroupTestCase) {
	t.Helper()

	require.NotNil(t, rg, "RegistryGroup should not be nil")
	require.NotNil(t, rg.Registries, "RegistryGroup.Registries should not be nil")

	require.Len(t, rg.Registries, len(tc.RegistriesUIDs), "The number of registries should match the number of expected registries")

	for _, registry := range rg.Registries {
		assert.Contains(t, tc.RegistriesUIDs, registry.UID(), "Registry UID %s aren't excepted", registry.UID())
	}
}
