package all

import (
	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/checksum"
	"github.com/go-sprout/sprout/registry/conversion"
	"github.com/go-sprout/sprout/registry/encoding"
	"github.com/go-sprout/sprout/registry/env"
	"github.com/go-sprout/sprout/registry/filesystem"
	"github.com/go-sprout/sprout/registry/maps"
	"github.com/go-sprout/sprout/registry/network"
	"github.com/go-sprout/sprout/registry/numeric"
	"github.com/go-sprout/sprout/registry/random"
	"github.com/go-sprout/sprout/registry/reflect"
	//nolint:staticcheck // kept until v1.2 to not break templates using the sprig signatures, will be swapped for `regex`
	"github.com/go-sprout/sprout/registry/regexp"
	"github.com/go-sprout/sprout/registry/semver"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/std"
	"github.com/go-sprout/sprout/registry/strings"
	"github.com/go-sprout/sprout/registry/time"
	"github.com/go-sprout/sprout/registry/uniqueid"
)

// all.RegistryGroup is a group of all registries available in Sprout excluding
// deprecated and experimental registries.
//
// Included registries: checksum, conversion, encoding, env, filesystem, maps,
// network, numeric, random, reflect, regexp, semver, slices, std, strings, time,
// uniqueid.
//
// The deprecated `regexp` registry is still included to keep this group
// non-breaking, it will be replaced by `regex` in Sprout v1.2. To opt in right
// now, register [github.com/go-sprout/sprout/registry/regex] before this group,
// its functions take precedence over the ones of `regexp`.
func RegistryGroup() *sprout.RegistryGroup {
	return sprout.NewRegistryGroup(
		checksum.NewRegistry(),
		conversion.NewRegistry(),
		encoding.NewRegistry(),
		env.NewRegistry(),
		filesystem.NewRegistry(),
		maps.NewRegistry(),
		network.NewRegistry(),
		numeric.NewRegistry(),
		random.NewRegistry(),
		reflect.NewRegistry(),
		regexp.NewRegistry(), //nolint:staticcheck // see the note above about the v1.2 migration to `regex`
		semver.NewRegistry(),
		slices.NewRegistry(),
		std.NewRegistry(),
		strings.NewRegistry(),
		time.NewRegistry(),
		uniqueid.NewRegistry(),
	)
}
