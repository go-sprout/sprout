package uniqueid

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Uuidv4 generates a new random UUID (Universally Unique Identifier) version 4.
// This function does not take parameters and returns a string representation
// of a UUID.
//
// Returns:
//
//	string - a new UUID string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidv4].
//
// [Sprout Documentation: uuidv4]: https://docs.atom.codes/sprout/registries/uniqueid#uuidv4
func (ur *UniqueIDRegistry) Uuidv4() string {
	return uuid.New().String()
}

// Uuidv7 generates a new UUID (Universally Unique Identifier) version 7, based
// on the current Unix time in milliseconds. Unlike a version 4, successive
// UUIDs are sortable by generation time.
//
// Returns:
//
//	string - a new UUID string.
//	error - when the random source cannot be read.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidv7].
//
// [Sprout Documentation: uuidv7]: https://docs.atom.codes/sprout/registries/uniqueid#uuidv7
func (ur *UniqueIDRegistry) Uuidv7() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// Uuidv5 generates a UUID (Universally Unique Identifier) version 5, derived
// from a namespace and a name using SHA-1. The same namespace and name always
// produce the same UUID.
//
// Parameters:
//
//	namespace string - one of the predefined namespaces (dns, url, oid, x500) or any valid UUID.
//	name string - the name to derive the UUID from.
//
// Returns:
//
//	string - the UUID string derived from the namespace and the name.
//	error - when the namespace is neither a predefined one nor a valid UUID.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidv5].
//
// [Sprout Documentation: uuidv5]: https://docs.atom.codes/sprout/registries/uniqueid#uuidv5
func (ur *UniqueIDRegistry) Uuidv5(namespace string, name string) (string, error) {
	ns, err := computeNamespace(namespace)
	if err != nil {
		return "", err
	}
	return uuid.NewSHA1(ns, []byte(name)).String(), nil
}

// Uuidv3 generates a UUID (Universally Unique Identifier) version 3, derived
// from a namespace and a name using MD5. The same namespace and name always
// produce the same UUID.
//
// Prefer `uuidv5` for new usages, version 3 only exists for compatibility with
// systems already relying on MD5 derived UUIDs.
//
// Parameters:
//
//	namespace string - one of the predefined namespaces (dns, url, oid, x500) or any valid UUID.
//	name string - the name to derive the UUID from.
//
// Returns:
//
//	string - the UUID string derived from the namespace and the name.
//	error - when the namespace is neither a predefined one nor a valid UUID.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidv3].
//
// [Sprout Documentation: uuidv3]: https://docs.atom.codes/sprout/registries/uniqueid#uuidv3
func (ur *UniqueIDRegistry) Uuidv3(namespace string, name string) (string, error) {
	ns, err := computeNamespace(namespace)
	if err != nil {
		return "", err
	}
	return uuid.NewMD5(ns, []byte(name)).String(), nil
}

// UuidNil returns the nil UUID, the UUID with all its bits set to zero.
//
// Returns:
//
//	string - the nil UUID string.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidNil].
//
// [Sprout Documentation: uuidNil]: https://docs.atom.codes/sprout/registries/uniqueid#uuidnil
func (ur *UniqueIDRegistry) UuidNil() string {
	return uuid.Nil.String()
}

// IsUUID checks if the given value is a valid UUID. On top of the canonical
// form, the urn prefixed form, the braced form and the form without hyphen are
// accepted.
//
// Parameters:
//
//	value string - the value to check.
//
// Returns:
//
//	bool - true if the value is a valid UUID, otherwise false.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: isUUID].
//
// [Sprout Documentation: isUUID]: https://docs.atom.codes/sprout/registries/uniqueid#isuuid
func (ur *UniqueIDRegistry) IsUUID(value string) bool {
	return uuid.Validate(value) == nil
}

// UuidVersion returns the version of the given UUID.
//
// Parameters:
//
//	value string - the UUID to read the version from.
//
// Returns:
//
//	int - the version of the UUID.
//	error - when the value is not a valid UUID.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidVersion].
//
// [Sprout Documentation: uuidVersion]: https://docs.atom.codes/sprout/registries/uniqueid#uuidversion
func (ur *UniqueIDRegistry) UuidVersion(value string) (int, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return 0, err
	}
	return int(id.Version()), nil
}

// UuidTime returns the time embedded in the given UUID. Only the versions
// carrying a timestamp are supported, namely the versions 1, 2, 6 and 7.
//
// Parameters:
//
//	value string - the UUID to read the time from.
//
// Returns:
//
//	time.Time - the time at which the UUID was generated.
//	error - when the value is not a valid UUID or when its version carries no timestamp.
//
// For an example of this function in a Go template, refer to [Sprout Documentation: uuidTime].
//
// [Sprout Documentation: uuidTime]: https://docs.atom.codes/sprout/registries/uniqueid#uuidtime
func (ur *UniqueIDRegistry) UuidTime(value string) (time.Time, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return time.Time{}, err
	}

	switch id.Version() {
	case 1, 2, 6, 7:
		sec, nsec := id.Time().UnixTime()
		return time.Unix(sec, nsec), nil
	default:
		return time.Time{}, fmt.Errorf("uuid version %d does not embed a time", id.Version())
	}
}
