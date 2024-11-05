package hermetic

import (
	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/checksum"
	"github.com/go-sprout/sprout/registry/conversion"
	"github.com/go-sprout/sprout/registry/encoding"
	"github.com/go-sprout/sprout/registry/filesystem"
	"github.com/go-sprout/sprout/registry/maps"
	"github.com/go-sprout/sprout/registry/numeric"
	"github.com/go-sprout/sprout/registry/reflect"
	"github.com/go-sprout/sprout/registry/regexp"
	"github.com/go-sprout/sprout/registry/semver"
	"github.com/go-sprout/sprout/registry/slices"
	"github.com/go-sprout/sprout/registry/std"
	"github.com/go-sprout/sprout/registry/strings"
	"github.com/go-sprout/sprout/registry/time"
	"github.com/go-sprout/sprout/registry/uniqueid"
)

// hermetic.RegistryGroup is a group of all registries don't depend on external services
// or influenced by the environment where the application is running.
//
// Included registries: checksum, conversion, encoding, filesystem, maps, numeric,
// reflect, regexp, semver, slices, std, strings, time, uniqueid.
func RegistryGroup() *sprout.RegistryGroup {
	return sprout.NewRegistryGroup(
		checksum.NewRegistry(),
		conversion.NewRegistry(),
		encoding.NewRegistry(),
		filesystem.NewRegistry(),
		maps.NewRegistry(),
		numeric.NewRegistry(),
		reflect.NewRegistry(),
		regexp.NewRegistry(),
		semver.NewRegistry(),
		slices.NewRegistry(),
		std.NewRegistry(),
		strings.NewRegistry(),
		time.NewRegistry(),
		uniqueid.NewRegistry(),
	)
}
