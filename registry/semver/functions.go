package semver

import (
	"github.com/Masterminds/semver/v3"
)

// Semver creates a new semantic version object from a given version string.
//
// Parameters:
//
//	value string - the version string to parse into a semantic version object.
//
// Returns:
//
//	*semver.Version - the parsed semantic version object.
//	error - an error if the version string is invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: semver].
//
// [Sprout Documentation: semver]: https://docs.atom.codes/sprout/registries/semver#semver
func (fh *SemverRegistry) Semver(value string) (*semver.Version, error) {
	return semver.NewVersion(value)
}

// SemverCompare checks if a given version string satisfies a specified semantic version constraint.
//
// Parameters:
//
//	constraint string - the version constraint to check against.
//	value string - the version string to validate against the constraint.
//
// Returns:
//
//	bool - true if the version satisfies the constraint, false otherwise.
//	error - an error if either the constraint or version string is invalid.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: semverCompare].
//
// [Sprout Documentation: semverCompare]: https://docs.atom.codes/sprout/registries/semver#semvercompare
func (fh *SemverRegistry) SemverCompare(constraint, value string) (bool, error) {
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		return false, err
	}

	v, err := semver.NewVersion(value)
	if err != nil {
		return false, err
	}

	return c.Check(v), nil
}
