package all_test

import (
	"testing"

	"github.com/go-sprout/sprout/group/all"
	"github.com/go-sprout/sprout/pesticide"
)

func TestRegistryGroup(t *testing.T) {
	tc := pesticide.GroupTestCase{
		RegistriesUIDs: []string{
			"go-sprout/sprout.checksum",
			"go-sprout/sprout.conversion",
			"go-sprout/sprout.encoding",
			"go-sprout/sprout.env",
			"go-sprout/sprout.filesystem",
			"go-sprout/sprout.maps",
			"go-sprout/sprout.network",
			"go-sprout/sprout.numeric",
			"go-sprout/sprout.random",
			"go-sprout/sprout.reflect",
			"go-sprout/sprout.regexp",
			"go-sprout/sprout.semver",
			"go-sprout/sprout.slices",
			"go-sprout/sprout.std",
			"go-sprout/sprout.strings",
			"go-sprout/sprout.time",
			"go-sprout/sprout.uniqueid",
		},
	}

	pesticide.RunGroupTest(t, all.RegistryGroup(), tc)
}